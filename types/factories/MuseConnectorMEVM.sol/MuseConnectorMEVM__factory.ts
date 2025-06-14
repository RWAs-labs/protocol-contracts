/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../../common";
import type {
  MuseConnectorMEVM,
  MuseConnectorMEVMInterface,
} from "../../MuseConnectorMEVM.sol/MuseConnectorMEVM";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "wmuse_",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "FUNGIBLE_MODULE_ADDRESS",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "onReceive",
    inputs: [
      {
        name: "museTxSenderAddress",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "sourceChainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "destinationAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "museValue",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "internalSendHash",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "onRevert",
    inputs: [
      {
        name: "museTxSenderAddress",
        type: "address",
        internalType: "address",
      },
      {
        name: "sourceChainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "destinationAddress",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "destinationChainId",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "remainingMuseValue",
        type: "uint256",
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        internalType: "bytes",
      },
      {
        name: "internalSendHash",
        type: "bytes32",
        internalType: "bytes32",
      },
    ],
    outputs: [],
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "send",
    inputs: [
      {
        name: "input",
        type: "tuple",
        internalType: "struct MuseInterfaces.sol.SendInput",
        components: [
          {
            name: "destinationChainId",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "destinationAddress",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "destinationGasLimit",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "message",
            type: "bytes",
            internalType: "bytes",
          },
          {
            name: "museValueAndGas",
            type: "uint256",
            internalType: "uint256",
          },
          {
            name: "museParams",
            type: "bytes",
            internalType: "bytes",
          },
        ],
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setWmuseAddress",
    inputs: [
      {
        name: "wmuse_",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "wmuse",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "event",
    name: "SetWMUSE",
    inputs: [
      {
        name: "wmuse_",
        type: "address",
        indexed: false,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "MuseReceived",
    inputs: [
      {
        name: "museTxSenderAddress",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "sourceChainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "destinationAddress",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "museValue",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "internalSendHash",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "MuseReverted",
    inputs: [
      {
        name: "museTxSenderAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "sourceChainId",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "destinationChainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "destinationAddress",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "remainingMuseValue",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "internalSendHash",
        type: "bytes32",
        indexed: true,
        internalType: "bytes32",
      },
    ],
    anonymous: false,
  },
  {
    type: "event",
    name: "MuseSent",
    inputs: [
      {
        name: "sourceTxOriginAddress",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "museTxSenderAddress",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "destinationChainId",
        type: "uint256",
        indexed: true,
        internalType: "uint256",
      },
      {
        name: "destinationAddress",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "museValueAndGas",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "destinationGasLimit",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "message",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
      {
        name: "museParams",
        type: "bytes",
        indexed: false,
        internalType: "bytes",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "FailedMuseSent",
    inputs: [],
  },
  {
    type: "error",
    name: "OnlyFungibleModule",
    inputs: [],
  },
  {
    type: "error",
    name: "OnlyWMUSEOrFungible",
    inputs: [],
  },
  {
    type: "error",
    name: "WMUSETransferFailed",
    inputs: [],
  },
  {
    type: "error",
    name: "WrongValue",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x6080604052348015600f57600080fd5b5060405161122f38038061122f833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6111a28061008d6000396000f3fe6080604052600436106100685760003560e01c8063942a5e1611610043578063942a5e1614610178578063eb3bacbd1461018b578063ec026901146101ab57600080fd5b8062173d46146100e757806329dd214d1461013d5780633ce4a5bc1461015057600080fd5b366100e25760005473ffffffffffffffffffffffffffffffffffffffff1633148015906100a957503373735b14bb79463307aacbed86daf3322b1e6226ab14155b156100e0576040517f229930b200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b3480156100f357600080fd5b506000546101149073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100e061014b366004610c5e565b6101cb565b34801561015c57600080fd5b5061011473735b14bb79463307aacbed86daf3322b1e6226ab81565b6100e0610186366004610cfe565b610547565b34801561019757600080fd5b506100e06101a6366004610da8565b6108b4565b3480156101b757600080fd5b506100e06101c6366004610dca565b61097a565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610218576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833414610251576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156102b957600080fd5b505af11580156102cd573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8a81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af1158015610352573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103769190610e05565b6103ac576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81156104e5578473ffffffffffffffffffffffffffffffffffffffff16633749c51a6040518060a001604052808b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018b905273ffffffffffffffffffffffffffffffffffffffff8a16604080840191909152606083018a90528051601f890183900483028101830190915287815260809092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526104b29190600401610e8b565b600060405180830381600087803b1580156104cc57600080fd5b505af11580156104e0573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff1302855733b40d8acb467ee990b6d56c05c80e28ebcabfa6e6f3f57cb50d6988b8b898989604051610535959493929190610f68565b60405180910390a45050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610594576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8334146105cd576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561063557600080fd5b505af1158015610649573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8e81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af11580156106ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f29190610e05565b610728576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8115610864578873ffffffffffffffffffffffffffffffffffffffff16633ff0693c6040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018a905260408083018a90528051601f890183900483028101830190915287815260609092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526108319190600401610fa1565b600060405180830381600087803b15801561084b57600080fd5b505af115801561085f573d6000803e3d6000fd5b505050505b80857f521fb0b407c2eb9b1375530e9b9a569889992140a688bc076aa72c1712012c888b8b8b8b8a8a8a6040516108a19796959493929190611036565b60405180910390a3505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610901576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f7325870b05f8f3412c318a35fc6a74feca51ea15811ec7a257676ca4db9d41769060200160405180910390a150565b6000546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526080830135604482015273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af11580156109fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1e9190610e05565b610a54576040517fa8c6fd4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526080830135600482015273ffffffffffffffffffffffffffffffffffffffff90911690632e1a7d4d90602401600060405180830381600087803b158015610ac357600080fd5b505af1158015610ad7573d6000803e3d6000fd5b50506040516000925073735b14bb79463307aacbed86daf3322b1e6226ab91506080840135908381818185875af1925050503d8060008114610b35576040519150601f19603f3d011682016040523d82523d6000602084013e610b3a565b606091505b5050905080610b75576040517fc7ffc47b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135337f7ec1c94701e09b1652f3e1d307e60c4b9ebf99aff8c2079fd1d8c585e031c4e432610ba76020870187611093565b60808801356040890135610bbe60608b018b611093565b610bcb60a08d018d611093565b604051610be0999897969594939291906110f8565b60405180910390a35050565b60008083601f840112610bfe57600080fd5b50813567ffffffffffffffff811115610c1657600080fd5b602083019150836020828501011115610c2e57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c5957600080fd5b919050565b60008060008060008060008060c0898b031215610c7a57600080fd5b883567ffffffffffffffff811115610c9157600080fd5b610c9d8b828c01610bec565b90995097505060208901359550610cb660408a01610c35565b945060608901359350608089013567ffffffffffffffff811115610cd957600080fd5b610ce58b828c01610bec565b999c989b50969995989497949560a00135949350505050565b600080600080600080600080600060e08a8c031215610d1c57600080fd5b610d258a610c35565b985060208a0135975060408a013567ffffffffffffffff811115610d4857600080fd5b610d548c828d01610bec565b90985096505060608a0135945060808a0135935060a08a013567ffffffffffffffff811115610d8257600080fd5b610d8e8c828d01610bec565b9a9d999c50979a9699959894979660c00135949350505050565b600060208284031215610dba57600080fd5b610dc382610c35565b9392505050565b600060208284031215610ddc57600080fd5b813567ffffffffffffffff811115610df357600080fd5b820160c08185031215610dc357600080fd5b600060208284031215610e1757600080fd5b81518015158114610dc357600080fd5b6000815180845260005b81811015610e4d57602081850181015186830182015201610e31565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000825160a06020840152610ea760c0840182610e27565b90506020840151604084015273ffffffffffffffffffffffffffffffffffffffff60408501511660608401526060840151608084015260808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160a0850152610f168282610e27565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b606081526000610f7c606083018789610f1f565b8560208401528281036040840152610f95818587610f1f565b98975050505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201526000604083015160c06060840152610fe760e0840182610e27565b905060608401516080840152608084015160a084015260a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160c0850152610f168282610e27565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a06040820152600061106c60a083018789610f1f565b8560608401528281036080840152611085818587610f1f565b9a9950505050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126110c857600080fd5b83018035915067ffffffffffffffff8211156110e357600080fd5b602001915036819003821315610c2e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16815260c06020820152600061112860c083018a8c610f1f565b8860408401528760608401528281036080840152611147818789610f1f565b905082810360a084015261115c818587610f1f565b9c9b50505050505050505050505056fea2646970667358221220bd61bc74d11634cf7d576167c47c99d49f303c9531df4335bc6539a2f0a0260264736f6c634300081a0033";

type MuseConnectorMEVMConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: MuseConnectorMEVMConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class MuseConnectorMEVM__factory extends ContractFactory {
  constructor(...args: MuseConnectorMEVMConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    wmuse_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(wmuse_, overrides || {});
  }
  override deploy(
    wmuse_: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(wmuse_, overrides || {}) as Promise<
      MuseConnectorMEVM & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): MuseConnectorMEVM__factory {
    return super.connect(runner) as MuseConnectorMEVM__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): MuseConnectorMEVMInterface {
    return new Interface(_abi) as MuseConnectorMEVMInterface;
  }
  static connect(
    address: string,
    runner?: ContractRunner | null
  ): MuseConnectorMEVM {
    return new Contract(address, _abi, runner) as unknown as MuseConnectorMEVM;
  }
}
