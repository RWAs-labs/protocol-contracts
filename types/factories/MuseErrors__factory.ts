/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type { MuseErrors, MuseErrorsInterface } from "../MuseErrors";

const _abi = [
  {
    type: "error",
    name: "CallerIsNotConnector",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "CallerIsNotTss",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "CallerIsNotTssOrUpdater",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "CallerIsNotTssUpdater",
    inputs: [
      {
        name: "caller",
        type: "address",
        internalType: "address",
      },
    ],
  },
  {
    type: "error",
    name: "InvalidAddress",
    inputs: [],
  },
  {
    type: "error",
    name: "MuseTransferError",
    inputs: [],
  },
] as const;

export class MuseErrors__factory {
  static readonly abi = _abi;
  static createInterface(): MuseErrorsInterface {
    return new Interface(_abi) as MuseErrorsInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): MuseErrors {
    return new Contract(address, _abi, runner) as unknown as MuseErrors;
  }
}
