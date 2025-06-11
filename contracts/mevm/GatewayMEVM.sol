// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

import { CallOptions, IGatewayMEVM } from "./interfaces/IGatewayMEVM.sol";

import { INotSupportedMethods } from "../../contracts/Errors.sol";
import { AbortContext, Abortable, RevertContext, RevertOptions, Revertable } from "../../contracts/Revert.sol";

import { IMRC20 } from "./interfaces/IMRC20.sol";
import "./interfaces/IWMUSE.sol";
import { MessageContext, UniversalContract } from "./interfaces/UniversalContract.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

/// @title GatewayMEVM.sol
/// @notice The GatewayMEVM.sol contract is the endpoint to call smart contracts on omnichain.
/// @dev The contract doesn't hold any funds and should never have active allowances.
contract GatewayMEVM is
    IGatewayMEVM,
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    INotSupportedMethods
{
    /// @notice Error indicating a zero address was provided.
    error ZeroAddress();

    /// @notice The constant address of the protocol
    address public constant PROTOCOL_ADDRESS = 0x735b14BB79463307AAcBED86DAf3322B1e6226aB;

    /// @notice The address of the Muse token.
    address public museToken;

    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /// @notice Max size of message + revertOptions revert message.
    uint256 public constant MAX_MESSAGE_SIZE = 2048;

    /// @notice Minimum gas limit for a call.
    uint256 public constant MIN_GAS_LIMIT = 100_000;

    /// @dev Only protocol address allowed modifier.
    modifier onlyProtocol() {
        if (msg.sender != PROTOCOL_ADDRESS) {
            revert CallerIsNotProtocol();
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /// @notice Initialize with address of muse token and admin account set as DEFAULT_ADMIN_ROLE.
    /// @dev Using admin to authorize upgrades and pause.
    function initialize(address museToken_, address admin_) public initializer {
        if (museToken_ == address(0) || admin_ == address(0)) {
            revert ZeroAddress();
        }
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(PAUSER_ROLE, admin_);
        museToken = museToken_;
    }

    /// @dev Authorizes the upgrade of the contract.
    /// @param newImplementation The address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @dev Receive function to receive MUSE from WETH9.withdraw().
    receive() external payable whenNotPaused {
        if (msg.sender != museToken && msg.sender != PROTOCOL_ADDRESS) revert OnlyWMUSEOrProtocol();
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Helper function to safely execute transferFrom
    /// @param mrc20 The MRC20.sol token address
    /// @param from The sender address
    /// @param to The recipient address
    /// @param amount The amount to transfer
    /// @return True if the transfer was successful, false otherwise.
    function _safeTransferFrom(address mrc20, address from, address to, uint256 amount) private returns (bool) {
        try IMRC20(mrc20).transferFrom(from, to, amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    // @notice Helper function to safely burn MRC20.sol tokens
    // @param mrc20 The MRC20.sol token address
    // @param amount The amount to burn
    // @return True if the burn was successful, false otherwise
    function _safeBurn(address mrc20, uint256 amount) private returns (bool) {
        try IMRC20(mrc20).burn(amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    // @notice Helper function to safely deposit
    // @param mrc20 The MRC20.sol token address
    // @param amount The target address to receive the deposited tokens
    // @param amount The amount to deposit
    // @return True if the deposit was successful, false otherwise
    function _safeDeposit(address mrc20, address target, uint256 amount) private returns (bool) {
        try IMRC20(mrc20).deposit(target, amount) returns (bool success) {
            return success;
        } catch {
            return false;
        }
    }

    /// @dev Private function to withdraw MRC20.sol tokens.
    /// @param amount The amount of tokens to withdraw.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @return The gas fee for the withdrawal.
    function _withdrawMRC20(uint256 amount, address mrc20) private returns (uint256) {
        // Use gas limit from mrc20
        return _withdrawMRC20WithGasLimit(amount, mrc20, IMRC20(mrc20).GAS_LIMIT());
    }

    /// @dev Private function to withdraw MRC20.sol tokens with gas limit.
    /// @param amount The amount of tokens to withdraw.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param gasLimit Gas limit.
    /// @return The gas fee for the withdrawal.
    function _withdrawMRC20WithGasLimit(uint256 amount, address mrc20, uint256 gasLimit) private returns (uint256) {
        (address gasMRC20, uint256 gasFee) = IMRC20(mrc20).withdrawGasFeeWithGasLimit(gasLimit);

        if (!_safeTransferFrom(gasMRC20, msg.sender, PROTOCOL_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed(gasMRC20, PROTOCOL_ADDRESS, gasFee);
        }

        if (!_safeTransferFrom(mrc20, msg.sender, address(this), amount)) {
            revert MRC20TransferFailed(mrc20, msg.sender, address(this), amount);
        }

        if (!_safeBurn(mrc20, amount)) {
            revert MRC20BurnFailed(mrc20, amount);
        }

        return gasFee;
    }

    /// @dev Private function to transfer MUSE tokens.
    /// @param amount The amount of tokens to transfer.
    /// @param to The address to transfer the tokens to.
    function _transferMUSE(uint256 amount, address to) private {
        if (!_safeTransferFrom(museToken, msg.sender, address(this), amount)) {
            revert FailedMuseSent(address(this), amount);
        }
        try IWETH9(museToken).withdraw(amount) {
            (bool sent,) = to.call{ value: amount }("");
            if (!sent) revert FailedMuseSent(to, amount);
        } catch {
            revert FailedMuseSent(to, amount);
        }
    }

    /// @notice Withdraw MRC20.sol tokens to an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param revertOptions Revert options.
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        address mrc20,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) {
            revert MessageSizeExceeded(revertOptions.revertMessage.length, MAX_MESSAGE_SIZE);
        }

        uint256 gasFee = _withdrawMRC20(amount, mrc20);
        emit Withdrawn(
            msg.sender,
            0,
            receiver,
            mrc20,
            amount,
            gasFee,
            IMRC20(mrc20).PROTOCOL_FLAT_FEE(),
            "",
            CallOptions({ gasLimit: IMRC20(mrc20).GAS_LIMIT(), isArbitraryCall: true }),
            revertOptions
        );
    }

    /// @notice Withdraw MRC20.sol tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        address mrc20,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) {
            revert MessageSizeExceeded(message.length + revertOptions.revertMessage.length, MAX_MESSAGE_SIZE);
        }

        uint256 gasFee = _withdrawMRC20WithGasLimit(amount, mrc20, callOptions.gasLimit);
        emit WithdrawnAndCalled(
            msg.sender,
            0,
            receiver,
            mrc20,
            amount,
            gasFee,
            IMRC20(mrc20).PROTOCOL_FLAT_FEE(),
            message,
            callOptions,
            revertOptions
        );
    }

    /// @notice Withdraw MUSE tokens to an external chain.
    //// @param receiver The receiver address on the external chain.
    //// @param amount The amount of tokens to withdraw.
    //// @param revertOptions Revert options.
    function withdraw(
        bytes memory, /*receiver*/
        uint256, /*amount*/
        uint256, /*chainId*/
        RevertOptions calldata /*revertOptions*/
    )
        external
        view
        whenNotPaused
    {
        // TODO: remove error and comment out code once MUSE supported back
        // MUSE is not currently supported for withdraws
        revert MUSENotSupported();

        // if (receiver.length == 0) revert ZeroAddress();
        // if (amount == 0) revert InsufficientMuseAmount();
        // if (revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        // _transferMUSE(amount, PROTOCOL_ADDRESS);
        // emit Withdrawn(
        //     msg.sender,
        //     chainId,
        //     receiver,
        //     address(museToken),
        //     amount,
        //     0,
        //     0,
        //     "",
        //     CallOptions({ gasLimit: 0, isArbitraryCall: true }),
        //     revertOptions
        // );
    }

    /// @notice Withdraw MUSE tokens and call a smart contract on an external chain.
    //// @param receiver The receiver address on the external chain.
    //// @param amount The amount of tokens to withdraw.
    //// @param chainId Chain id of the external chain.
    //// @param message The calldata to pass to the contract call.
    //// @param callOptions Call options including gas limit and arbirtrary call flag.
    //// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory, /*receiver*/
        uint256, /*amount*/
        uint256, /*chainId*/
        bytes calldata, /*message*/
        CallOptions calldata, /*callOptions*/
        RevertOptions calldata /*revertOptions*/
    )
        external
        view
        whenNotPaused
    {
        // TODO: remove error and comment out code once MUSE supported back
        // MUSE is not currently supported for withdraws
        revert MUSENotSupported();

        // if (receiver.length == 0) revert ZeroAddress();
        // if (amount == 0) revert InsufficientMuseAmount();
        // if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        // if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) revert MessageSizeExceeded();

        // _transferMUSE(amount, PROTOCOL_ADDRESS);
        // emit WithdrawnAndCalled(
        //     msg.sender, chainId, receiver, address(museToken), amount, 0, 0, message, callOptions, revertOptions
        // );
    }

    /// @notice Call a smart contract on an external chain without asset transfer.
    /// @param receiver The receiver address on the external chain.
    /// @param mrc20 Address of mrc20 to pay fees.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function call(
        bytes memory receiver,
        address mrc20,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        whenNotPaused
    {
        if (callOptions.gasLimit < MIN_GAS_LIMIT) revert InsufficientGasLimit();
        if (message.length + revertOptions.revertMessage.length > MAX_MESSAGE_SIZE) {
            revert MessageSizeExceeded(message.length + revertOptions.revertMessage.length, MAX_MESSAGE_SIZE);
        }

        _call(receiver, mrc20, message, callOptions, revertOptions);
    }

    function _call(
        bytes memory receiver,
        address mrc20,
        bytes calldata message,
        CallOptions memory callOptions,
        RevertOptions memory revertOptions
    )
        private
    {
        if (receiver.length == 0) revert ZeroAddress();

        (address gasMRC20, uint256 gasFee) = IMRC20(mrc20).withdrawGasFeeWithGasLimit(callOptions.gasLimit);
        if (!_safeTransferFrom(gasMRC20, msg.sender, PROTOCOL_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed(gasMRC20, PROTOCOL_ADDRESS, gasFee);
        }

        emit Called(msg.sender, mrc20, receiver, message, callOptions, revertOptions);
    }

    /// @notice Deposit foreign coins into MRC20.sol.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(address mrc20, uint256 amount, address target) external onlyProtocol whenNotPaused {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();

        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!_safeDeposit(mrc20, target, amount)) {
            revert MRC20DepositFailed(mrc20, target, amount);
        }
    }

    /// @notice Execute a user-specified contract on MEVM.
    /// @param context The context of the cross-chain call.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function execute(
        MessageContext calldata context,
        address mrc20,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();

        UniversalContract(target).onCall(context, mrc20, amount, message);
    }

    /// @notice Deposit foreign coins into MRC20.sol and call a user-specified contract on MEVM.
    /// @param context The context of the cross-chain call.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        MessageContext calldata context,
        address mrc20,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!_safeDeposit(mrc20, target, amount)) {
            revert MRC20DepositFailed(mrc20, target, amount);
        }

        UniversalContract(target).onCall(context, mrc20, amount, message);
    }

    /// @notice Deposit MUSE and call a user-specified contract on MEVM.
    /// @param context The context of the cross-chain call.
    /// @param amount The amount of tokens to transfer.
    /// @param target The target contract to call.
    /// @param message The calldata to pass to the contract call.
    function depositAndCall(
        MessageContext calldata context,
        uint256 amount,
        address target,
        bytes calldata message
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMuseAmount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        _transferMUSE(amount, target);
        UniversalContract(target).onCall(context, museToken, amount, message);
    }

    /// @notice Revert a user-specified contract on MEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(
        address target,
        RevertContext calldata revertContext
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (target == address(0)) revert ZeroAddress();

        Revertable(target).onRevert(revertContext);
    }

    /// @notice Deposit foreign coins into MRC20.sol and revert a user-specified contract on MEVM.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param amount The amount of tokens to revert.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function depositAndRevert(
        address mrc20,
        uint256 amount,
        address target,
        RevertContext calldata revertContext
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!_safeDeposit(mrc20, target, amount)) {
            revert MRC20DepositFailed(mrc20, target, amount);
        }

        Revertable(target).onRevert(revertContext);
    }

    /// @notice Call onAbort on a user-specified contract on MEVM.
    /// this function doesn't deposit the asset to the target contract. This operation is done directly by the protocol.
    /// the assets are deposited to the target contract even if onAbort reverts.
    /// @param target The target contract to call.
    /// @param abortContext Abort context to pass to onAbort.
    function executeAbort(
        address target,
        AbortContext calldata abortContext
    )
        external
        nonReentrant
        onlyProtocol
        whenNotPaused
    {
        if (target == address(0)) revert ZeroAddress();
        Abortable(target).onAbort(abortContext);
    }
}
