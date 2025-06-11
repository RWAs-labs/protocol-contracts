// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

import "../../../contracts/Revert.sol";
import "./UniversalContract.sol";

/// @title IGatewayMEVMEvents
/// @notice Interface for the events emitted by the GatewayMEVM.sol contract.
interface IGatewayMEVMEvents {
    /// @notice Emitted when a cross-chain call is made.
    /// @param sender The address of the sender.
    /// @param mrc20 Address of mrc20 to pay fees.
    /// @param receiver The receiver address on the external chain.
    /// @param message The calldata passed to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event Called(
        address indexed sender,
        address indexed mrc20,
        bytes receiver,
        bytes message,
        CallOptions callOptions,
        RevertOptions revertOptions
    );

    /// @notice Emitted when a withdrawal is made.
    /// @param sender The address from which the tokens are withdrawn.
    /// @param chainId Chain id of external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param value The amount of tokens withdrawn.
    /// @param gasfee The gas fee for the withdrawal.
    /// @param protocolFlatFee The protocol flat fee for the withdrawal.
    /// @param message The calldata passed with the withdraw. No longer used. Kept to maintain compatibility.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event Withdrawn(
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

    /// @notice Emitted when a withdraw and call is made.
    /// @param sender The address from which the tokens are withdrawn.
    /// @param chainId Chain id of external chain.
    /// @param receiver The receiver address on the external chain.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param value The amount of tokens withdrawn.
    /// @param gasfee The gas fee for the withdrawal.
    /// @param protocolFlatFee The protocol flat fee for the withdrawal.
    /// @param message The calldata passed to the contract call.
    /// @param callOptions Call options including gas limit and arbirtrary call flag.
    /// @param revertOptions Revert options.
    event WithdrawnAndCalled(
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
}

/// @title IGatewayMEVMErrors
/// @notice Interface for the errors used in the GatewayMEVM.sol contract.
interface IGatewayMEVMErrors {
    /// @notice Error indicating a withdrawal failure.
    /// @param token The address of the token that failed to withdraw.
    /// @param recipient The address that was supposed to receive the tokens.
    /// @param amount The amount of tokens that failed to withdraw.
    error WithdrawalFailed(address token, address recipient, uint256 amount);

    /// @notice Error indicating an insufficient MRC20.sol token amount.
    error InsufficientMRC20Amount();

    /// @notice Error indicating an insufficient muse amount.
    error InsufficientMuseAmount();

    /// @notice Error indicating a failure to burn MRC20.sol tokens.
    /// @param mrc20 The address of the MRC20.sol token that failed to burn.
    /// @param amount The amount of tokens that failed to burn.
    error MRC20BurnFailed(address mrc20, uint256 amount);

    /// @notice Error indicating a failure to transfer MRC20.sol tokens.
    /// @param mrc20 The address of the MRC20.sol token that failed to transfer.
    /// @param from The address sending the tokens.
    /// @param to The address receiving the tokens.
    /// @param amount The amount of tokens that failed to transfer.
    error MRC20TransferFailed(address mrc20, address from, address to, uint256 amount);

    /// @notice Error indicating a failure to deposit MRC20.sol tokens.
    /// @param mrc20 The address of the MRC20.sol token that failed to deposit.
    /// @param to The address that was supposed to receive the deposit.
    /// @param amount The amount of tokens that failed to deposit.
    error MRC20DepositFailed(address mrc20, address to, uint256 amount);

    /// @notice Error indicating a failure to transfer gas fee.
    /// @param token The address of the token used for gas fee.
    /// @param to The address that was supposed to receive the gas fee.
    /// @param amount The amount of gas fee that failed to transfer.
    error GasFeeTransferFailed(address token, address to, uint256 amount);

    /// @notice Error indicating that the caller is not the protocol account.
    error CallerIsNotProtocol();

    /// @notice Error indicating an invalid target address.
    error InvalidTarget();

    /// @notice Error indicating a failure to send MUSE tokens.
    /// @param recipient The address that was supposed to receive the MUSE tokens.
    /// @param amount The amount of MUSE tokens that failed to send.
    error FailedMuseSent(address recipient, uint256 amount);

    /// @notice Error indicating that only WMUSE or the protocol address can call the function.
    error OnlyWMUSEOrProtocol();

    /// @notice Error indicating an insufficient gas limit.
    error InsufficientGasLimit();

    /// @notice Error indicating message size exceeded in external functions.
    /// @param provided The size of the message that was provided.
    /// @param maximum The maximum allowed message size.
    error MessageSizeExceeded(uint256 provided, uint256 maximum);
}

/// @title IGatewayMEVM.sol
/// @notice Interface for the GatewayMEVM.sol contract.
/// @dev Defines functions for cross-chain interactions and token handling.
interface IGatewayMEVM is IGatewayMEVMErrors, IGatewayMEVMEvents {
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
        external;

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
        external;

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
        external;

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
        external;

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
        external;

    /// @notice Deposit foreign coins into MRC20.sol.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param amount The amount of tokens to deposit.
    /// @param target The target address to receive the deposited tokens.
    function deposit(address mrc20, uint256 amount, address target) external;

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
        external;

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
        external;

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
        external;

    /// @notice Revert a user-specified contract on MEVM.
    /// @param target The target contract to call.
    /// @param revertContext Revert context to pass to onRevert.
    function executeRevert(address target, RevertContext calldata revertContext) external;

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
        external;
}

/// @notice CallOptions struct passed to call and withdrawAndCall functions.
/// @param gasLimit Gas limit.
/// @param isArbitraryCall Indicates if call should be arbitrary or authenticated.
struct CallOptions {
    uint256 gasLimit;
    bool isArbitraryCall;
}
