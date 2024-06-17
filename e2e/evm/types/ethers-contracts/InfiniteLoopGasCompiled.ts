/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
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
} from "./common";

export interface InfiniteLoopGasCompiledInterface extends Interface {
  getFunction(nameOrSignature: "counter" | "forever"): FunctionFragment;

  encodeFunctionData(functionFragment: "counter", values?: undefined): string;
  encodeFunctionData(functionFragment: "forever", values?: undefined): string;

  decodeFunctionResult(functionFragment: "counter", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "forever", data: BytesLike): Result;
}

export interface InfiniteLoopGasCompiled extends BaseContract {
  connect(runner?: ContractRunner | null): InfiniteLoopGasCompiled;
  waitForDeployment(): Promise<this>;

  interface: InfiniteLoopGasCompiledInterface;

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

  counter: TypedContractMethod<[], [bigint], "view">;

  forever: TypedContractMethod<[], [void], "nonpayable">;

  getFunction<T extends ContractMethod = ContractMethod>(
    key: string | FunctionFragment
  ): T;

  getFunction(
    nameOrSignature: "counter"
  ): TypedContractMethod<[], [bigint], "view">;
  getFunction(
    nameOrSignature: "forever"
  ): TypedContractMethod<[], [void], "nonpayable">;

  filters: {};
}