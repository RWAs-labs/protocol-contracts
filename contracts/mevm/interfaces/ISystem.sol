// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

/// @title ISystem
/// @notice Interface for the System contract.
/// @dev Defines functions for system contract callable by fungible module.
interface ISystem {
    function FUNGIBLE_MODULE_ADDRESS() external view returns (address);

    function wMuseContractAddress() external view returns (address);

    function uniswapv2FactoryAddress() external view returns (address);

    function gasPriceByChainId(uint256 chainID) external view returns (uint256);

    function gasCoinMRC20ByChainId(uint256 chainID) external view returns (address);

    function gasMusePoolByChainId(uint256 chainID) external view returns (address);
}
