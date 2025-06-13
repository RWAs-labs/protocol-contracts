// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import { IGatewayMEVM } from "../../contracts/mevm/interfaces/IGatewayMEVM.sol";
import { CallOptions, RevertOptions } from "../../contracts/mevm/interfaces/IGatewayMEVM.sol";
import "../../contracts/mevm/interfaces/IMRC20.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/// @title SenderMEVM
/// @notice This contract is used just for testing.
/// @dev Provides functions to call a receiver on EVM and to withdraw and call a receiver on EVM.
contract SenderMEVM {
    /// @notice The address of the gateway contract.
    address public gateway;

    /// @notice Error indicating that the approval of tokens failed.
    error ApprovalFailed();

    constructor(address gateway_) {
        gateway = gateway_;
    }

    /// @notice Call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param mrc20 Address of mrc20 to pay fees.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Encodes the function call and passes it to the gateway.
    function callReceiver(bytes memory receiver, address mrc20, string memory str, uint256 num, bool flag) external {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        RevertOptions memory revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        CallOptions memory callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        IMRC20(mrc20).approve(gateway, 100_000);

        // Pass encoded call to gateway
        IGatewayMEVM(gateway).call(receiver, mrc20, message, callOptions, revertOptions);
    }

    /// @notice Withdraw and call a receiver on EVM.
    /// @param receiver The address of the receiver on the external chain.
    /// @param amount The amount of tokens to withdraw.
    /// @param mrc20 The address of the MRC20.sol token.
    /// @param str A string parameter to pass to the receiver's function.
    /// @param num A numeric parameter to pass to the receiver's function.
    /// @param flag A boolean parameter to pass to the receiver's function.
    /// @dev Approves the gateway to withdraw tokens and encodes the function call to pass to the gateway.
    function withdrawAndCallReceiver(
        bytes memory receiver,
        uint256 amount,
        address mrc20,
        string memory str,
        uint256 num,
        bool flag
    )
        external
    {
        // Encode the function call to the receiver's receivePayable method
        bytes memory message = abi.encodeWithSignature("receivePayable(string,uint256,bool)", str, num, flag);

        // Approve gateway to withdraw
        if (!IMRC20(mrc20).approve(gateway, amount + 100_000)) revert ApprovalFailed();

        RevertOptions memory revertOptions = RevertOptions({
            revertAddress: address(0x321),
            callOnRevert: true,
            abortAddress: address(0x321),
            revertMessage: "",
            onRevertGasLimit: 0
        });

        CallOptions memory callOptions = CallOptions({ gasLimit: 100_000, isArbitraryCall: false });

        // Pass encoded call to gateway
        IGatewayMEVM(gateway).withdrawAndCall(receiver, amount, mrc20, message, callOptions, revertOptions);
    }
}
