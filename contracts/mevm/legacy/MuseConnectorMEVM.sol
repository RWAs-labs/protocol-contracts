// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "../interfaces/IWMUSE.sol";

interface MuseInterfaces {
    /**
     * @dev Use SendInput to interact with the Connector: connector.send(SendInput)
     */
    struct SendInput {
        /// @dev Chain id of the destination chain. More about chain ids
        /// https://docs.musechain.com/learn/glossary#chain-id
        uint256 destinationChainId;
        /// @dev Address receiving the message on the destination chain (expressed in bytes since it can be non-EVM)
        bytes destinationAddress;
        /// @dev Gas limit for the destination chain's transaction
        uint256 destinationGasLimit;
        /// @dev An encoded, arbitrary message to be parsed by the destination contract
        bytes message;
        /// @dev MUSE to be sent cross-chain + MuseChain gas fees + destination chain gas fees (expressed in MUSE)
        uint256 museValueAndGas;
        /// @dev Optional parameters for the MuseChain protocol
        bytes museParams;
    }

    /**
     * @dev Our Connector calls onMuseMessage with this struct as argument
     */
    struct MuseMessage {
        bytes museTxSenderAddress;
        uint256 sourceChainId;
        address destinationAddress;
        /// @dev Remaining MUSE from museValueAndGas after subtracting MuseChain gas fees and destination gas fees
        uint256 museValue;
        bytes message;
    }

    /**
     * @dev Our Connector calls onMuseRevert with this struct as argument
     */
    struct MuseRevert {
        address museTxSenderAddress;
        uint256 sourceChainId;
        bytes destinationAddress;
        uint256 destinationChainId;
        /// @dev Equals to: museValueAndGas - MuseChain gas fees - destination chain gas fees - source chain revert tx
        /// gas fees
        uint256 remainingMuseValue;
        bytes message;
    }
}

interface MuseReceiver {
    /**
     * @dev onMuseMessage is called when a cross-chain message reaches a contract
     */
    function onMuseMessage(MuseInterfaces.MuseMessage calldata museMessage) external;

    /**
     * @dev onMuseRevert is called when a cross-chain message reverts.
     * It's useful to rollback to the original state
     */
    function onMuseRevert(MuseInterfaces.MuseRevert calldata museRevert) external;
}

contract MuseConnectorMEVM {
    /// @notice WMUSE token address.
    address public wmuse;

    /// @notice Contract custom errors.
    error OnlyWMUSEOrFungible();
    error WMUSETransferFailed();
    error OnlyFungibleModule();
    error FailedMuseSent();
    error WrongValue();

    /// @notice Fungible module address.
    address public constant FUNGIBLE_MODULE_ADDRESS = payable(0x735b14BB79463307AAcBED86DAf3322B1e6226aB);

    event SetWMUSE(address wmuse_);

    event MuseSent(
        address sourceTxOriginAddress,
        address indexed museTxSenderAddress,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 museValueAndGas,
        uint256 destinationGasLimit,
        bytes message,
        bytes museParams
    );

    event MuseReceived(
        bytes museTxSenderAddress,
        uint256 indexed sourceChainId,
        address indexed destinationAddress,
        uint256 museValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    event MuseReverted(
        address museTxSenderAddress,
        uint256 sourceChainId,
        uint256 indexed destinationChainId,
        bytes destinationAddress,
        uint256 remainingMuseValue,
        bytes message,
        bytes32 indexed internalSendHash
    );

    /**
     * @dev Modifier to restrict actions to fungible module.
     */
    modifier onlyFungibleModule() {
        if (msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyFungibleModule();
        _;
    }

    constructor(address wmuse_) {
        wmuse = wmuse_;
    }

    /// @dev Receive function to receive MUSE from WETH9.withdraw().
    receive() external payable {
        if (msg.sender != wmuse && msg.sender != FUNGIBLE_MODULE_ADDRESS) revert OnlyWMUSEOrFungible();
    }

    function setWmuseAddress(address wmuse_) external onlyFungibleModule {
        wmuse = wmuse_;
        emit SetWMUSE(wmuse_);
    }

    /**
     * @dev Sends MUSE and bytes messages (to execute it) crosschain.
     * @param input, SendInput struct, checkout above.
     */
    function send(MuseInterfaces.SendInput calldata input) external {
        // Transfer wmuse to "fungible" module, which will be burnt by the protocol post processing via hooks.
        if (!IWETH9(wmuse).transferFrom(msg.sender, address(this), input.museValueAndGas)) revert WMUSETransferFailed();
        IWETH9(wmuse).withdraw(input.museValueAndGas);
        (bool sent,) = FUNGIBLE_MODULE_ADDRESS.call{ value: input.museValueAndGas }("");
        if (!sent) revert FailedMuseSent();
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
     * This method can be called only by Fungible Module.
     * Transfer the Muse tokens to destination and calls onMuseMessage if it's needed.
     * To perform the transfer wrap the new tokens
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
        payable
        onlyFungibleModule
    {
        if (msg.value != museValue) revert WrongValue();
        IWETH9(wmuse).deposit{ value: museValue }();
        if (!IWETH9(wmuse).transferFrom(address(this), destinationAddress, museValue)) revert WMUSETransferFailed();

        if (message.length > 0) {
            MuseReceiver(destinationAddress).onMuseMessage(
                MuseInterfaces.MuseMessage(museTxSenderAddress, sourceChainId, destinationAddress, museValue, message)
            );
        }

        emit MuseReceived(museTxSenderAddress, sourceChainId, destinationAddress, museValue, message, internalSendHash);
    }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by Fungible Module.
     * Transfer the Muse tokens to destination and calls onMuseRevert if it's needed.
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
        payable
        onlyFungibleModule
    {
        if (msg.value != remainingMuseValue) revert WrongValue();
        IWETH9(wmuse).deposit{ value: remainingMuseValue }();
        if (!IWETH9(wmuse).transferFrom(address(this), museTxSenderAddress, remainingMuseValue)) {
            revert WMUSETransferFailed();
        }

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
