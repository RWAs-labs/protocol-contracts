// SPDX-License-Identifier: MIT
pragma solidity 0.8.29;

/// @title INotSupportedMethods
/// @notice Interface for contracts that with non supported methods.
interface INotSupportedMethods {
    error MUSENotSupported();
    error CallOnRevertNotSupported();
}
