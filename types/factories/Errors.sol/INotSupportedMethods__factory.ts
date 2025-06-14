/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Interface, type ContractRunner } from "ethers";
import type {
  INotSupportedMethods,
  INotSupportedMethodsInterface,
} from "../../Errors.sol/INotSupportedMethods";

const _abi = [
  {
    type: "error",
    name: "CallOnRevertNotSupported",
    inputs: [],
  },
  {
    type: "error",
    name: "MUSENotSupported",
    inputs: [],
  },
] as const;

export class INotSupportedMethods__factory {
  static readonly abi = _abi;
  static createInterface(): INotSupportedMethodsInterface {
    return new Interface(_abi) as INotSupportedMethodsInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): INotSupportedMethods {
    return new Contract(
      address,
      _abi,
      runner
    ) as unknown as INotSupportedMethods;
  }
}
