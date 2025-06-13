// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/console.sol";
import "forge-std/Script.sol";
import "contracts/mevm/GatewayMEVM.sol";

contract UpgradeSimulation is Script {
    function run() external {
        // Forked environment
        address payable proxyAddress = payable(vm.envAddress("PROXY_ADDRESS"));
        address adminAddress = vm.envAddress("ADMIN_ADDRESS");

        // Deploy the new implementation contract
        bytes32 implSalt = keccak256("GatewayMEVM");
        address gatewayImpl = address(new GatewayMEVM{salt: implSalt}());

        GatewayMEVM proxy = GatewayMEVM(proxyAddress);

        // Get the current state
        address museToken = proxy.museToken();

        // Simulate the upgrade
        vm.prank(adminAddress);
        proxy.upgradeToAndCall(gatewayImpl, "");

        // After upgrade, verify that the state is intact
        require(museToken == proxy.museToken(), "museToken address mismatch");

        console.log("Upgraded contract state:");
        console.log("museToken address:", proxy.museToken());
    }
}