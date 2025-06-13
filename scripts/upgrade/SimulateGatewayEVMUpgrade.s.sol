// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";
import "forge-std/Script.sol";
import "contracts/evm/GatewayEVM.sol";

contract UpgradeSimulation is Script {
    function run() external {
        // Forked environment
        address proxyAddress = vm.envAddress("PROXY_ADDRESS");
        address adminAddress = vm.envAddress("ADMIN_ADDRESS");

        // Deploy the new implementation contract
        bytes32 implSalt = keccak256("GatewayEVM");
        address gatewayImpl = address(new GatewayEVM{salt: implSalt}());

        GatewayEVM proxy = GatewayEVM(proxyAddress);

        // Get the current state
        address custody = proxy.custody();
        address tssAddress = proxy.tssAddress();
        address museConnector = proxy.museConnector();
        address museToken = proxy.museToken();

        // Simulate the upgrade
        vm.prank(adminAddress);
        proxy.upgradeToAndCall(gatewayImpl, "");

        // After upgrade, verify that the state is intact
        require(custody == proxy.custody(), "custody address mismatch");
        require(tssAddress == proxy.tssAddress(), "tss address mismatch");
        require(museConnector == proxy.museConnector(), "museConnector address mismatch");
        require(museToken == proxy.museToken(), "museToken address mismatch");

        console.log("Upgraded contract state:");
        console.log("custody address:", proxy.custody());
        console.log("tss address:", proxy.tssAddress());
        console.log("museConnector address:", proxy.museConnector());
        console.log("museToken address:", proxy.museToken());
    }
}