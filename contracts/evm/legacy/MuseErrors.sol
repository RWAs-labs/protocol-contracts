// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @dev Common custom errors
 */
interface MuseErrors {
    // @dev Thrown when caller is not the address defined as TSS address
    error CallerIsNotTss(address caller);

    // @dev Thrown when caller is not the address defined as connector address
    error CallerIsNotConnector(address caller);

    // @dev Thrown when caller is not the address defined as TSS Updater address
    error CallerIsNotTssUpdater(address caller);

    // @dev Thrown when caller is not the address defined as TSS or TSS Updater address
    error CallerIsNotTssOrUpdater(address caller);

    // @dev Thrown when a contract receives an invalid address param, mostly zero address validation
    error InvalidAddress();

    // @dev Thrown when Muse can't be transferred for some reason
    error MuseTransferError();
}
