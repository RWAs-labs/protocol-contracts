// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import { RevertContext } from "../../contracts/Revert.sol";
import {
    IGatewayEVM,
    IGatewayEVMErrors,
    IGatewayEVMEvents,
    MessageContext
} from "../../contracts/evm/interfaces/IGatewayEVM.sol";
import "../../contracts/evm/interfaces/IMuseConnector.sol";

/// @title MuseConnectorBase.sol
/// @notice Abstract base contract for MuseConnector.
/// @dev This contract implements basic functionality for handling tokens and interacting with the Gateway contract.
abstract contract MuseConnectorBase is
    Initializable,
    UUPSUpgradeable,
    IMuseConnectorEvents,
    ReentrancyGuardUpgradeable,
    PausableUpgradeable,
    AccessControlUpgradeable
{
    using SafeERC20 for IERC20;

    /// @notice Error indicating that a zero address was provided.
    error ZeroAddress();

    /// @notice The Gateway contract used for executing cross-chain calls.
    IGatewayEVM public gateway;
    /// @notice The address of the Muse token.
    address public museToken;
    /// @notice The address of the TSS (Threshold Signature Scheme) contract.
    address public tssAddress;

    /// @notice New role identifier for withdrawer role.
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");
    /// @notice New role identifier for pauser role.
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
    /// @notice New role identifier for tss role.
    bytes32 public constant TSS_ROLE = keccak256("TSS_ROLE");

    /// @notice Initializer for MuseConnectors.
    /// @dev Set admin as default admin and pauser, and tssAddress as tss role.
    function initialize(
        address gateway_,
        address museToken_,
        address tssAddress_,
        address admin_
    )
        public
        virtual
        initializer
    {
        if (gateway_ == address(0) || museToken_ == address(0) || tssAddress_ == address(0) || admin_ == address(0)) {
            revert ZeroAddress();
        }

        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __AccessControl_init();
        __Pausable_init();

        gateway = IGatewayEVM(gateway_);
        museToken = museToken_;
        tssAddress = tssAddress_;

        _grantRole(DEFAULT_ADMIN_ROLE, admin_);
        _grantRole(WITHDRAWER_ROLE, tssAddress_);
        _grantRole(TSS_ROLE, tssAddress_);
        _grantRole(PAUSER_ROLE, admin_);
        _grantRole(PAUSER_ROLE, tssAddress_);
    }

    /// @dev Authorizes the upgrade of the contract, sender must be owner.
    /// @param newImplementation Address of the new implementation.
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(DEFAULT_ADMIN_ROLE) { }

    /// @notice Update tss address
    /// @param newTSSAddress new tss address
    function updateTSSAddress(address newTSSAddress) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (newTSSAddress == address(0)) revert ZeroAddress();

        _revokeRole(WITHDRAWER_ROLE, tssAddress);
        _revokeRole(TSS_ROLE, tssAddress);

        _grantRole(WITHDRAWER_ROLE, newTSSAddress);
        _grantRole(TSS_ROLE, newTSSAddress);

        emit UpdatedMuseConnectorTSSAddress(tssAddress, newTSSAddress);

        tssAddress = newTSSAddress;
    }

    /// @notice Pause contract.
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /// @notice Unpause contract.
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /// @notice Withdraw tokens to a specified address.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    function withdraw(address to, uint256 amount, bytes32 internalSendHash) external virtual;

    /// @notice Withdraw tokens and call a contract through Gateway.
    /// @param messageContext Message context containing sender.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    function withdrawAndCall(
        MessageContext calldata messageContext,
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash
    )
        external
        virtual;

    /// @notice Withdraw tokens and call a contract with a revert callback through Gateway.
    /// @param to The address to withdraw tokens to.
    /// @param amount The amount of tokens to withdraw.
    /// @param data The calldata to pass to the contract call.
    /// @param internalSendHash A hash used for internal tracking of the transaction.
    /// @param revertContext Revert context to pass to onRevert.
    function withdrawAndRevert(
        address to,
        uint256 amount,
        bytes calldata data,
        bytes32 internalSendHash,
        RevertContext calldata revertContext
    )
        external
        virtual;

    /// @notice Handle received tokens.
    /// @param amount The amount of tokens received.
    function receiveTokens(uint256 amount) external virtual;
}
