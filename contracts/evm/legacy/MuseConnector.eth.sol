// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./ConnectorErrors.sol";
import "./MuseConnector.base.sol";
import "./MuseInterfaces.sol";

/**
 * @dev ETH implementation of MuseConnector.
 * This contract manages interactions between TSS and different chains.
 * This version is only for Ethereum network because in the other chains we mint and burn and in this one we lock and
 * unlock.
 */
contract MuseConnectorEth is MuseConnectorBase {
    constructor(
        address museToken_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    )
        MuseConnectorBase(museToken_, tssAddress_, tssAddressUpdater_, pauserAddress_)
    { }

    function getLockedAmount() external view returns (uint256) {
        return IERC20(museToken).balanceOf(address(this));
    }

    /**
     * @dev Entrypoint to send data through MuseChain
     * This call locks the token on the contract and emits an event with all the data needed by the protocol.
     */
    function send(MuseInterfaces.SendInput calldata input) external override whenNotPaused {
        bool success = IERC20(museToken).transferFrom(msg.sender, address(this), input.museValueAndGas);
        if (!success) revert MuseTransferError();

        emit MuseSent(
            tx.origin,
            msg.sender,
            input.destinationChainId,
            input.destinationAddress,
            input.museValueAndGas,
            input.destinationGasLimit,
            input.message,
            input.museParams
        );
    }

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by TSS.
     * Transfers the Muse tokens to destination and calls onMuseMessage if it's needed.
     */
    function onReceive(
        bytes calldata museTxSenderAddress,
        uint256 sourceChainId,
        address destinationAddress,
        uint256 museValue,
        bytes calldata message,
        bytes32 internalSendHash
    )
        external
        override
        onlyTssAddress
    {
        bool success = IERC20(museToken).transfer(destinationAddress, museValue);
        if (!success) revert MuseTransferError();

        if (message.length > 0) {
            MuseReceiver(destinationAddress).onMuseMessage(
                MuseInterfaces.MuseMessage(museTxSenderAddress, sourceChainId, destinationAddress, museValue, message)
            );
        }

        emit MuseReceived(museTxSenderAddress, sourceChainId, destinationAddress, museValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by TSS.
     * Transfers the Muse tokens to destination and calls onMuseRevert if it's needed.
     */
    function onRevert(
        address museTxSenderAddress,
        uint256 sourceChainId,
        bytes calldata destinationAddress,
        uint256 destinationChainId,
        uint256 remainingMuseValue,
        bytes calldata message,
        bytes32 internalSendHash
    )
        external
        override
        whenNotPaused
        onlyTssAddress
    {
        bool success = IERC20(museToken).transfer(museTxSenderAddress, remainingMuseValue);
        if (!success) revert MuseTransferError();

        if (message.length > 0) {
            MuseReceiver(museTxSenderAddress).onMuseRevert(
                MuseInterfaces.MuseRevert(
                    museTxSenderAddress,
                    sourceChainId,
                    destinationAddress,
                    destinationChainId,
                    remainingMuseValue,
                    message
                )
            );
        }

        emit MuseReverted(
            museTxSenderAddress,
            sourceChainId,
            destinationChainId,
            destinationAddress,
            remainingMuseValue,
            message,
            internalSendHash
        );
    }
}
