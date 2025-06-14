// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { CallOptions, IGatewayMEVM } from "../../../contracts/mevm/interfaces/IGatewayMEVM.sol";

import { RevertContext, RevertOptions, Revertable } from "../../../contracts/Revert.sol";
import "../../../contracts/mevm/interfaces/IWMUSE.sol";
import { IMRC20 } from "../../../contracts/mevm/interfaces/IMRC20.sol";
import { MessageContext, UniversalContract } from "../../../contracts/mevm/interfaces/UniversalContract.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

/// @title GatewayMEVMUpgradeTest
/// @notice Modified GatewayMEVM.sol contract for testing upgrades
/// @dev The only difference is in event naming
/// @custom:oz-upgrades-from GatewayMEVM.sol
contract GatewayMEVMUpgradeTest is
    IGatewayMEVM,
    Initializable,
    AccessControlUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable
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
    uint256 public constant MAX_MESSAGE_SIZE = 1024;

    /// @dev Modified event for testing upgrade.
    event WithdrawnV2(
        address indexed sender,
        uint256 indexed chainId,
        bytes receiver,
        address mrc20,
        uint256 value,
        uint256 gasfee,
        uint256 protocolFlatFee,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

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
        if (!IMRC20(gasMRC20).transferFrom(msg.sender, PROTOCOL_ADDRESS, gasFee)) {
            revert GasFeeTransferFailed(gasMRC20, PROTOCOL_ADDRESS, gasFee);
        }

        if (!IMRC20(mrc20).transferFrom(msg.sender, address(this), amount)) {
            revert MRC20TransferFailed(mrc20, msg.sender, address(this), amount);
        }

        if (!IMRC20(mrc20).burn(amount)) revert MRC20BurnFailed(mrc20, amount);

        return gasFee;
    }

    /// @dev Private function to transfer MUSE tokens.
    /// @param amount The amount of tokens to transfer.
    /// @param to The address to transfer the tokens to.
    function _transferMUSE(uint256 amount, address to) private {
        if (!IWETH9(museToken).transferFrom(msg.sender, address(this), amount)) {
            revert FailedMuseSent(address(this), amount);
        }
        IWETH9(museToken).withdraw(amount);
        (bool sent,) = to.call{ value: amount }("");
        if (!sent) revert FailedMuseSent(to, amount);
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
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        uint256 messageSize = revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        uint256 gasFee = _withdrawMRC20(amount, mrc20);
        emit WithdrawnV2(
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
    /// @param gasLimit Gas limit.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        address mrc20,
        bytes calldata message,
        uint256 gasLimit,
        RevertOptions calldata revertOptions
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (gasLimit == 0) revert InsufficientGasLimit();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        uint256 gasFee = _withdrawMRC20WithGasLimit(amount, mrc20, gasLimit);
        emit Withdrawn(
            msg.sender,
            0,
            receiver,
            mrc20,
            amount,
            gasFee,
            IMRC20(mrc20).PROTOCOL_FLAT_FEE(),
            message,
            CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }),
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
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (callOptions.gasLimit == 0) revert InsufficientGasLimit();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        uint256 gasFee = _withdrawMRC20WithGasLimit(amount, mrc20, callOptions.gasLimit);
        emit Withdrawn(
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
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param revertOptions Revert options.
    function withdraw(
        bytes memory receiver,
        uint256 amount,
        uint256 chainId,
        RevertOptions calldata revertOptions
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMuseAmount();
        uint256 messageSize = revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        _transferMUSE(amount, PROTOCOL_ADDRESS);
        emit Withdrawn(
            msg.sender,
            chainId,
            receiver,
            address(museToken),
            amount,
            0,
            0,
            "",
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
    }

    /// @notice Withdraw MUSE tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param chainId Chain id of the external chain.
    /// @param message The calldata to pass to the contract call.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        uint256 chainId,
        bytes calldata message,
        RevertOptions calldata revertOptions
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMuseAmount();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        _transferMUSE(amount, PROTOCOL_ADDRESS);
        emit Withdrawn(
            msg.sender,
            chainId,
            receiver,
            address(museToken),
            amount,
            0,
            0,
            message,
            CallOptions({ gasLimit: 0, isArbitraryCall: true }),
            revertOptions
        );
    }

    /// @notice Withdraw MUSE tokens and call a smart contract on an external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param chainId Chain id of the external chain.
    /// @param message The calldata to pass to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    function withdrawAndCall(
        bytes memory receiver,
        uint256 amount,
        uint256 chainId,
        bytes calldata message,
        CallOptions calldata callOptions,
        RevertOptions calldata revertOptions
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (receiver.length == 0) revert ZeroAddress();
        if (amount == 0) revert InsufficientMuseAmount();
        if (callOptions.gasLimit == 0) revert InsufficientGasLimit();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        _transferMUSE(amount, PROTOCOL_ADDRESS);
        emit Withdrawn(
            msg.sender, chainId, receiver, address(museToken), amount, 0, 0, message, callOptions, revertOptions
        );
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
        nonReentrant
        whenNotPaused
    {
        if (callOptions.gasLimit == 0) revert InsufficientGasLimit();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        _call(receiver, mrc20, message, callOptions, revertOptions);
    }

    /// @notice Call a smart contract on an external chain without asset transfer.
    /// @param receiver The receiver address on the external chain.
    /// @param mrc20 Address of mrc20 to pay fees.
    /// @param message The calldata to pass to the contract call.
    /// @param gasLimit Gas limit.
    /// @param revertOptions Revert options.
    function call(
        bytes memory receiver,
        address mrc20,
        bytes calldata message,
        uint256 gasLimit,
        RevertOptions calldata revertOptions
    )
        external
        nonReentrant
        whenNotPaused
    {
        if (gasLimit == 0) revert InsufficientGasLimit();
        uint256 messageSize = message.length + revertOptions.revertMessage.length;
        if (messageSize > MAX_MESSAGE_SIZE) revert MessageSizeExceeded(messageSize, MAX_MESSAGE_SIZE);

        _call(receiver, mrc20, message, CallOptions({ gasLimit: gasLimit, isArbitraryCall: true }), revertOptions);
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
        if (!IMRC20(gasMRC20).transferFrom(msg.sender, PROTOCOL_ADDRESS, gasFee)) {
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

        if (!IMRC20(mrc20).deposit(target, amount)) revert MRC20DepositFailed(mrc20, target, amount);
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
        onlyProtocol
        whenNotPaused
    {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!IMRC20(mrc20).deposit(target, amount)) revert MRC20DepositFailed(mrc20, target, amount);
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
    function executeRevert(address target, RevertContext calldata revertContext) external onlyProtocol whenNotPaused {
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
        onlyProtocol
        whenNotPaused
    {
        if (mrc20 == address(0) || target == address(0)) revert ZeroAddress();
        if (amount == 0) revert InsufficientMRC20Amount();
        if (target == PROTOCOL_ADDRESS || target == address(this)) revert InvalidTarget();

        if (!IMRC20(mrc20).deposit(target, amount)) revert MRC20DepositFailed(mrc20, target, amount);
        Revertable(target).onRevert(revertContext);
    }
}
