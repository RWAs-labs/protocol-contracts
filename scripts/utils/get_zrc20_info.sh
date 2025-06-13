#!/bin/bash

# This script retrieves various attributes of a MRC20.sol contract on a blockchain.
# It takes the contract address and an RPC URL as inputs and calls several 
# functions on the contract to obtain key information. These attributes 
# are typically required when a clone of the contract is to be deployed.
#
# Usage:
#   ./get_mrc20_info.sh <contract_address> <rpc_url>

# Check if cast is installed
if ! command -v cast &> /dev/null
then
    echo "cast CLI could not be found. Please install it first."
    exit 1
fi

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <contract_address> <rpc_url>"
    exit 1
fi

# Get the contract address and RPC URL from the script arguments
CONTRACT_ADDRESS=$1
RPC_URL=$2

# Run the cast commands and capture the outputs
SYSTEM_CONTRACT=$(cast call "$CONTRACT_ADDRESS" "SYSTEM_CONTRACT_ADDRESS()(address)" --rpc-url "$RPC_URL" | tr -d '\n')
MRC20_NAME=$(cast call "$CONTRACT_ADDRESS" "name()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
MRC20_SYMBOL=$(cast call "$CONTRACT_ADDRESS" "symbol()(string)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/"//g')
MRC20_DECIMALS=$(cast call "$CONTRACT_ADDRESS" "decimals()(uint8)" --rpc-url "$RPC_URL" | tr -d '\n')
MRC20_CHAIN_ID=$(cast call "$CONTRACT_ADDRESS" "CHAIN_ID()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/\[.*\]//g')
MRC20_COIN_TYPE=$(cast call "$CONTRACT_ADDRESS" "COIN_TYPE()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n')
MRC20_GAS_LIMIT=$(cast call "$CONTRACT_ADDRESS" "GAS_LIMIT()(uint256)" --rpc-url "$RPC_URL" | tr -d '\n' | sed 's/\[.*\]//g')

# Output the final result string
echo "SYSTEM_CONTRACT=$SYSTEM_CONTRACT"
echo "MRC20_NAME=$MRC20_NAME"
echo "MRC20_SYMBOL=$MRC20_SYMBOL"
echo "MRC20_DECIMALS=$MRC20_DECIMALS"
echo "MRC20_CHAIN_ID=$MRC20_CHAIN_ID"
echo "MRC20_COIN_TYPE=$MRC20_COIN_TYPE"
echo "MRC20_GAS_LIMIT=$MRC20_GAS_LIMIT"