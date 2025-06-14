/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumberish,
  BytesLike,
  FunctionFragment,
  Result,
  Interface,
  ContractRunner,
  ContractMethod,
  Listener,
} from "ethers";
import type {
  TypedContractEvent,
  TypedDeferredTopicFilter,
  TypedEventLog,
  TypedListener,
  TypedContractMethod,
} from "../common";

export declare namespace MuseInterfaces {
  export type SendInputStruct = {
    destinationChainId: BigNumberish;
    destinationAddress: BytesLike;
    destinationGasLimit: BigNumberish;
    message: BytesLike;
    museValueAndGas: BigNumberish;
    museParams: BytesLike;
  };

  export type SendInputStructOutput = [
    destinationChainId: bigint,
    destinationAddress: string,
    destinationGasLimit: bigint,
    message: string,
    museValueAndGas: bigint,
    museParams: string
  ] & {
    destinationChainId: bigint;
    destinationAddress: string;
    destinationGasLimit: bigint;
    message: string;
    museValueAndGas: bigint;
    museParams: string;
  };
}

export interface MuseConnectorInterface extends Interface {
  getFunction(nameOrSignature: "send"): FunctionFragment;

  encodeFunctionData(
    functionFragment: "send",
    values: [MuseInterfaces.SendInputStruct]
  ): string;

  decodeFunctionResult(functionFragment: "send", data: BytesLike): Result;
}

export interface MuseConnector extends BaseContract {
  connect(runner?: ContractRunner | null): MuseConnector;
  waitForDeployment(): Promise<this>;

  interface: MuseConnectorInterface;

  queryFilter<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;
  queryFilter<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEventLog<TCEvent>>>;

  on<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  on<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  once<TCEvent extends TypedContractEvent>(
    event: TCEvent,
    listener: TypedListener<TCEvent>
  ): Promise<this>;
  once<TCEvent extends TypedContractEvent>(
    filter: TypedDeferredTopicFilter<TCEvent>,
    listener: TypedListener<TCEvent>
  ): Promise<this>;

  listeners<TCEvent extends TypedContractEvent>(
    event: TCEvent
  ): Promise<Array<TypedListener<TCEvent>>>;
  listeners(eventName?: string): Promise<Array<Listener>>;
  removeAllListeners<TCEvent extends TypedContractEvent>(
    event?: TCEvent
  ): Promise<this>;

  send: TypedContractMethod<
    [input: MuseInterfaces.SendInputStruct],
    [void],
    "nonpayable"
  >;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "send"
  ): TypedContractMethod<
    [input: MuseInterfaces.SendInputStruct],
    [void],
    "nonpayable"
  >;

  filters: {};
}
