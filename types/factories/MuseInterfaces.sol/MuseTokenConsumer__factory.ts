/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  MuseTokenConsumer,
  MuseTokenConsumerInterface,
} from "../../MuseInterfaces.sol/MuseTokenConsumer";

const _abi = [
  {
    type: "function",
    name: "getEthFromMuse",
    inputs: [
      {
        name: "destinationAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "minAmountOut",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "museTokenAmount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "getTokenFromMuse",
    inputs: [
      {
        name: "destinationAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "minAmountOut",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "outputToken",
        type: "address",
        internalType: "address",
      },
      {
        name: "museTokenAmount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "getMuseFromEth",
    inputs: [
      {
        name: "destinationAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "minAmountOut",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "getMuseFromToken",
    inputs: [
      {
        name: "destinationAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "minAmountOut",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "inputToken",
        type: "address",
        internalType: "address",
      },
      {
        name: "inputTokenAmount",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    outputs: [
      {
        name: "",
        type: "uint256",
        internalType: "uint256",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "hasMuseLiquidity",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "event",
    name: "EthExchangedForMuse",
    inputs: [
      {
        name: "amountIn",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "amountOut",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "TokenExchangedForMuse",
    inputs: [
      {
        name: "token",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "amountIn",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "amountOut",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "MuseExchangedForEth",
    inputs: [
      {
        name: "amountIn",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "amountOut",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "MuseExchangedForToken",
    inputs: [
      {
        name: "token",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "amountIn",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "amountOut",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
] as const;

export class MuseTokenConsumer__factory {
  static readonly abi = _abi;
  static createInterface(): MuseTokenConsumerInterface {
    return new Interface(_abi) as MuseTokenConsumerInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): MuseTokenConsumer {
    return new Contract(address, _abi, runner) as unknown as MuseTokenConsumer;
  }
}
