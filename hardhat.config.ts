import "@nomicfoundation/hardhat-toolbox";
import "./tasks/addresses";
import { getHardhatConfigNetworks } from "@rwas-labs/networks";
import { HardhatUserConfig } from "hardhat/config";

const config: HardhatUserConfig = {
  networks: {
    ...getHardhatConfigNetworks(),
  },
  solidity: "0.8.29",
};

export default config;
