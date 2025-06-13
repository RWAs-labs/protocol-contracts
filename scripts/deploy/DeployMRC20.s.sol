// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import "forge-std/Script.sol";
import "../../contracts/mevm/MRC20.sol";
import "../../contracts/mevm/interfaces/IMRC20.sol";

contract DeployMRC20 is Script {
    function run() external {
        address gateway = vm.envAddress("GATEWAY_PROXY_MEVM");
        address systemContract = vm.envAddress("SYSTEM_CONTRACT");
        string memory name = vm.envString("MRC20_NAME");
        string memory symbol = vm.envString("MRC20_SYMBOL");
        uint8 decimals = uint8(vm.envUint("MRC20_DECIMALS"));
        uint chainID = vm.envUint("MRC20_CHAIN_ID");
        CoinType coinType = CoinType(vm.envUint("MRC20_COIN_TYPE"));
        uint gasLimit = vm.envUint("MRC20_GAS_LIMIT");

        vm.startBroadcast();

        MRC20 mrc20 = new MRC20(
            name,
            symbol,
            decimals,
            chainID,
            coinType,
            gasLimit,
            systemContract,
            gateway
        );

        require(address(mrc20) != address(0), "deployment failed");

        require(keccak256(abi.encodePacked(mrc20.name())) == keccak256(abi.encodePacked(name)), "name not set");
        require(keccak256(abi.encodePacked(mrc20.symbol())) == keccak256(abi.encodePacked(symbol)), "symbol not set");
        require(mrc20.decimals() == decimals, "decimals not set");
        require(mrc20.CHAIN_ID() == chainID, "chain id not set");
        require(mrc20.GAS_LIMIT() == gasLimit, "gas limit not set");
        require(mrc20.COIN_TYPE() == coinType, "coin type not set");
        require(mrc20.gatewayAddress() == gateway, "gateway not set");
        require(mrc20.SYSTEM_CONTRACT_ADDRESS() == systemContract, "system contract not set");

        vm.stopBroadcast();
    }
}
