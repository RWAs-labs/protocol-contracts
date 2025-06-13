// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "test/utils/TestERC20.sol";

// This is just to deploy test erc20 tokens for testing deployments
contract DeployTestERC20 is Script {
    function run() external {
        bytes32 erc20Salt = keccak256("TestERC20");

        vm.startBroadcast();

        TestERC20 muse = new TestERC20{salt: erc20Salt}("muse", "MUSE");
        require(address(muse) != address(0), "deployment failed");

        address expectedAddr = vm.computeCreate2Address(
            erc20Salt,
            hashInitCode(type(TestERC20).creationCode, abi.encode("muse", "MUSE"))
        );

        require(expectedAddr == address(muse), "muse address doesn't match expected address");

        vm.stopBroadcast();
    }
}