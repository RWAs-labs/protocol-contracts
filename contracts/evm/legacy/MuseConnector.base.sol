// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/Pausable.sol";

import "./ConnectorErrors.sol";
import "./MuseInterfaces.sol";

/**
 * @dev Main abstraction of MuseConnector.
 * This contract manages interactions between TSS and different chains.
 * There's an instance of this contract on each chain supported by MuseChain.
 */
contract MuseConnectorBase is ConnectorErrors, Pausable {
    address public immutable museToken;

    /**
     * @dev Multisig contract to pause incoming transactions.
     * The responsibility of pausing outgoing transactions is left to the protocol for more flexibility.
     */
    address public pauserAddress;

    /**
     * @dev Collectively held by MuseChain validators.
     */
    address public tssAddress;

    /**
     * @dev This address will start pointing to a multisig contract, then it will become the TSS address itself.
     */
    address public tssAddressUpdater;

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

    event TSSAddressUpdated(address callerAddress, address newTssAddress);

    event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress);

    event PauserAddressUpdated(address callerAddress, address newTssAddress);

    /**
     * @dev Constructor requires initial addresses.
     * museToken address is the only immutable one, while others can be updated.
     */
    constructor(address museToken_, address tssAddress_, address tssAddressUpdater_, address pauserAddress_) {
        if (
            museToken_ == address(0) || tssAddress_ == address(0) || tssAddressUpdater_ == address(0)
                || pauserAddress_ == address(0)
        ) {
            revert MuseCommonErrors.InvalidAddress();
        }

        museToken = museToken_;
        tssAddress = tssAddress_;
        tssAddressUpdater = tssAddressUpdater_;
        pauserAddress = pauserAddress_;
    }

    /**
     * @dev Modifier to restrict actions to pauser address.
     */
    modifier onlyPauser() {
        if (msg.sender != pauserAddress) revert CallerIsNotPauser(msg.sender);
        _;
    }

    /**
     * @dev Modifier to restrict actions to TSS address.
     */
    modifier onlyTssAddress() {
        if (msg.sender != tssAddress) revert CallerIsNotTss(msg.sender);
        _;
    }

    /**
     * @dev Modifier to restrict actions to TSS updater address.
     */
    modifier onlyTssUpdater() {
        if (msg.sender != tssAddressUpdater) revert CallerIsNotTssUpdater(msg.sender);
        _;
    }

    /**
     * @dev Update the pauser address. The only address allowed to do that is the current pauser.
     */
    function updatePauserAddress(address pauserAddress_) external onlyPauser {
        if (pauserAddress_ == address(0)) revert MuseCommonErrors.InvalidAddress();

        pauserAddress = pauserAddress_;

        emit PauserAddressUpdated(msg.sender, pauserAddress_);
    }

    /**
     * @dev Update the TSS address. The address can be updated by the TSS updater or the TSS address itself.
     */
    function updateTssAddress(address tssAddress_) external {
        if (msg.sender != tssAddress && msg.sender != tssAddressUpdater) revert CallerIsNotTssOrUpdater(msg.sender);
        if (tssAddress_ == address(0)) revert MuseCommonErrors.InvalidAddress();

        tssAddress = tssAddress_;

        emit TSSAddressUpdated(msg.sender, tssAddress_);
    }

    /**
     * @dev Changes the ownership of tssAddressUpdater to be the one held by the MuseChain TSS Signer nodes.
     */
    function renounceTssAddressUpdater() external onlyTssUpdater {
        if (tssAddress == address(0)) revert MuseCommonErrors.InvalidAddress();

        tssAddressUpdater = tssAddress;
        emit TSSAddressUpdaterUpdated(msg.sender, tssAddressUpdater);
    }

    /**
     * @dev Pause the input (send) transactions.
     */
    function pause() external onlyPauser {
        _pause();
    }

    /**
     * @dev Unpause the contract to allow transactions again.
     */
    function unpause() external onlyPauser {
        _unpause();
    }

    /**
     * @dev Entrypoint to send data and value through MuseChain.
     */
    function send(MuseInterfaces.SendInput calldata input) external virtual { }

    /**
     * @dev Handler to receive data from other chain.
     * This method can be called only by TSS. Access validation is in implementation.
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
        virtual
    { }

    /**
     * @dev Handler to receive errors from other chain.
     * This method can be called only by TSS. Access validation is in implementation.
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
        virtual
    { }
}
