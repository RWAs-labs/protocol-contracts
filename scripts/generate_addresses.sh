#!/bin/bash

echo "Generating protocol addresses..."

npx hardhat addresses --network muse_testnet > ./data/addresses.testnet.json
npx hardhat addresses --network muse_mainnet > ./data/addresses.mainnet.json

echo "Generating protocol addresses types..."

npx ts-node scripts/generate_addresses_types.ts > ./lib/types.ts