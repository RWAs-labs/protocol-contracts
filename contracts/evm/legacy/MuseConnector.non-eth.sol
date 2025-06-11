// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "./MuseConnector.base.sol";
import "./MuseInterfaces.sol";
import "./MuseNonEthInterface.sol";

/**
 * @dev Non ETH implementation of MuseConnector.
 * This contract manages interactions between TSS and different chains.
 * This version is for every chain but Etherum network because in the other chains we mint and burn and in Etherum we
 * lock and unlock
 */
contract MuseConnectorNonEth is MuseConnectorBase {
    uint256 public maxSupply = 2 ** 256 - 1;

    event MaxSupplyUpdated(address callerAddress, uint256 newMaxSupply);

    constructor(
        address museTokenAddress_,
        address tssAddress_,
        address tssAddressUpdater_,
        address pauserAddress_
    )
        MuseConnectorBase(museTokenAddress_, tssAddress_, tssAddressUpdater_, pauserAddress_)
    { }

    function getLockedAmount() external view returns (uint256) {
        return MuseNonEthInterface(museToken).balanceOf(address(this));
    }

    function setMaxSupply(uint256 maxSupply_) external onlyTssAddress {
        maxSupply = maxSupply_;
        emit MaxSupplyUpdated(msg.sender, maxSupply_);
    }

    /**
     * @dev Entry point to send data to protocol
     * This call burn the token and emit an event with all the data needed by the protocol
     */
    function send(MuseInterfaces.SendInput calldata input) external override whenNotPaused {
        MuseNonEthInterface(museToken).burnFrom(msg.sender, input.museValueAndGas);

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
     * Transfer the Muse tokens to destination and calls onMuseMessage if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
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
        if (museValue + MuseNonEthInterface(museToken).totalSupply() > maxSupply) revert ExceedsMaxSupply(maxSupply);
        MuseNonEthInterface(museToken).mint(destinationAddress, museValue, internalSendHash);

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
     * Transfer the Muse tokens to destination and calls onMuseRevert if it's needed.
     * To perform the transfer mint new tokens, validating first the maxSupply allowed in the current chain.
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
        if (remainingMuseValue + MuseNonEthInterface(museToken).totalSupply() > maxSupply) {
            revert ExceedsMaxSupply(maxSupply);
        }
        MuseNonEthInterface(museToken).mint(museTxSenderAddress, remainingMuseValue, internalSendHash);

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
