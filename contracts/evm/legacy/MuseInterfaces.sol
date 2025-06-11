// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

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

interface MuseConnector {
    /**
     * @dev Sending value and data cross-chain is as easy as calling connector.send(SendInput)
     */
    function send(MuseInterfaces.SendInput calldata input) external;
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

/**
 * @dev MuseTokenConsumer makes it easier to handle the following situations:
 *   - Getting Muse using native coin (to pay for destination gas while using `connector.send`)
 *   - Getting Muse using a token (to pay for destination gas while using `connector.send`)
 *   - Getting native coin using Muse (to return unused destination gas when `onMuseRevert` is executed)
 *   - Getting a token using Muse (to return unused destination gas when `onMuseRevert` is executed)
 * @dev The interface can be implemented using different strategies, like UniswapV2, UniswapV3, etc
 */
interface MuseTokenConsumer {
    event EthExchangedForMuse(uint256 amountIn, uint256 amountOut);
    event TokenExchangedForMuse(address token, uint256 amountIn, uint256 amountOut);
    event MuseExchangedForEth(uint256 amountIn, uint256 amountOut);
    event MuseExchangedForToken(address token, uint256 amountIn, uint256 amountOut);

    function getMuseFromEth(address destinationAddress, uint256 minAmountOut) external payable returns (uint256);

    function getMuseFromToken(
        address destinationAddress,
        uint256 minAmountOut,
        address inputToken,
        uint256 inputTokenAmount
    )
        external
        returns (uint256);

    function getEthFromMuse(
        address destinationAddress,
        uint256 minAmountOut,
        uint256 museTokenAmount
    )
        external
        returns (uint256);

    function getTokenFromMuse(
        address destinationAddress,
        uint256 minAmountOut,
        address outputToken,
        uint256 museTokenAmount
    )
        external
        returns (uint256);

    function hasMuseLiquidity() external view returns (bool);
}

interface MuseCommonErrors {
    error InvalidAddress();
}
