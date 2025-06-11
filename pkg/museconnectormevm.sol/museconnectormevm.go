// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package museconnectormevm

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MuseInterfacesSendInput is an auto generated low-level Go binding around an user-defined struct.
type MuseInterfacesSendInput struct {
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	DestinationGasLimit *big.Int
	Message             []byte
	MuseValueAndGas     *big.Int
	MuseParams          []byte
}

// MuseConnectorMEVMMetaData contains all meta data concerning the MuseConnectorMEVM contract.
var MuseConnectorMEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"wmuse_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"FUNGIBLE_MODULE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onReceive\",\"inputs\":[{\"name\":\"museTxSenderAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"museValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"onRevert\",\"inputs\":[{\"name\":\"museTxSenderAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingMuseValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structMuseInterfaces.SendInput\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"museValueAndGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"museParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWmuseAddress\",\"inputs\":[{\"name\":\"wmuse_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"wmuse\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MuseReceived\",\"inputs\":[{\"name\":\"museTxSenderAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"museValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MuseReverted\",\"inputs\":[{\"name\":\"museTxSenderAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sourceChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"remainingMuseValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MuseSent\",\"inputs\":[{\"name\":\"sourceTxOriginAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"museTxSenderAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"museValueAndGas\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"museParams\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetWMUSE\",\"inputs\":[{\"name\":\"wmuse_\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedMuseSent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyFungibleModule\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWMUSEOrFungible\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WMUSETransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongValue\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051611230380380611230833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b6111a38061008d6000396000f3fe6080604052600436106100695760003560e01c80636e4b3d52116100435780636e4b3d5214610179578063942a5e1614610199578063ec026901146101ac57600080fd5b806329dd214d146100e85780633ce4a5bc146100fb57806369a646371461014c57600080fd5b366100e35760005473ffffffffffffffffffffffffffffffffffffffff1633148015906100aa57503373735b14bb79463307aacbed86daf3322b1e6226ab14155b156100e1576040517f1c068dfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b6100e16100f6366004610c5f565b6101cc565b34801561010757600080fd5b5061012373735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561015857600080fd5b506000546101239073ffffffffffffffffffffffffffffffffffffffff1681565b34801561018557600080fd5b506100e1610194366004610cff565b610548565b6100e16101a7366004610d21565b61060e565b3480156101b857600080fd5b506100e16101c7366004610dcb565b61097b565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610219576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833414610252576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156102ba57600080fd5b505af11580156102ce573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8a81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af1158015610353573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103779190610e06565b6103ad576040517f14b1f9f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81156104e6578473ffffffffffffffffffffffffffffffffffffffff16630a4076556040518060a001604052808b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018b905273ffffffffffffffffffffffffffffffffffffffff8a16604080840191909152606083018a90528051601f890183900483028101830190915287815260809092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526104b39190600401610e8c565b600060405180830381600087803b1580156104cd57600080fd5b505af11580156104e1573d6000803e3d6000fd5b505050505b808573ffffffffffffffffffffffffffffffffffffffff16877ff430afc8ea78e97aef2646b9aee7afd691fcb57b04b185e1430e279a518a8b3c8b8b898989604051610536959493929190610f69565b60405180910390a45050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610595576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f702607b065ed0b40b11595136269ccf48c3f5786b60c546289e812cb2d947b679060200160405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461065b576040517fea02b3f300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b833414610694576040517f98d4901c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b1580156106fc57600080fd5b505af1158015610710573d6000803e3d6000fd5b50506000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8e81166024830152604482018a905290911693506323b872dd925060640190506020604051808303816000875af1158015610795573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107b99190610e06565b6107ef576040517f14b1f9f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b811561092b578873ffffffffffffffffffffffffffffffffffffffff166311b0d8956040518060c001604052808c73ffffffffffffffffffffffffffffffffffffffff1681526020018b81526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018a905260408083018a90528051601f890183900483028101830190915287815260609092019190889088908190840183828082843760009201919091525050509152506040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526108f89190600401610fa2565b600060405180830381600087803b15801561091257600080fd5b505af1158015610926573d6000803e3d6000fd5b505050505b80857f5d921d2b8faf864052d4535b06cf528da89b7227158f2b251b54189d2e77411c8b8b8b8b8a8a8a6040516109689796959493929190611037565b60405180910390a3505050505050505050565b6000546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526080830135604482015273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af11580156109fb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1f9190610e06565b610a55576040517f14b1f9f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526080830135600482015273ffffffffffffffffffffffffffffffffffffffff90911690632e1a7d4d90602401600060405180830381600087803b158015610ac457600080fd5b505af1158015610ad8573d6000803e3d6000fd5b50506040516000925073735b14bb79463307aacbed86daf3322b1e6226ab91506080840135908381818185875af1925050503d8060008114610b36576040519150601f19603f3d011682016040523d82523d6000602084013e610b3b565b606091505b5050905080610b76576040517f3a1caf4a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135337f9bf48436256611f46f7d8bbefc3abe994fdf7da86f3829eb6c73f8dd5d9ef27d32610ba86020870187611094565b60808801356040890135610bbf60608b018b611094565b610bcc60a08d018d611094565b604051610be1999897969594939291906110f9565b60405180910390a35050565b60008083601f840112610bff57600080fd5b50813567ffffffffffffffff811115610c1757600080fd5b602083019150836020828501011115610c2f57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c5a57600080fd5b919050565b60008060008060008060008060c0898b031215610c7b57600080fd5b883567ffffffffffffffff811115610c9257600080fd5b610c9e8b828c01610bed565b90995097505060208901359550610cb760408a01610c36565b945060608901359350608089013567ffffffffffffffff811115610cda57600080fd5b610ce68b828c01610bed565b999c989b50969995989497949560a00135949350505050565b600060208284031215610d1157600080fd5b610d1a82610c36565b9392505050565b600080600080600080600080600060e08a8c031215610d3f57600080fd5b610d488a610c36565b985060208a0135975060408a013567ffffffffffffffff811115610d6b57600080fd5b610d778c828d01610bed565b90985096505060608a0135945060808a0135935060a08a013567ffffffffffffffff811115610da557600080fd5b610db18c828d01610bed565b9a9d999c50979a9699959894979660c00135949350505050565b600060208284031215610ddd57600080fd5b813567ffffffffffffffff811115610df457600080fd5b820160c08185031215610d1a57600080fd5b600060208284031215610e1857600080fd5b81518015158114610d1a57600080fd5b6000815180845260005b81811015610e4e57602081850181015186830182015201610e32565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000825160a06020840152610ea860c0840182610e28565b90506020840151604084015273ffffffffffffffffffffffffffffffffffffffff60408501511660608401526060840151608084015260808401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160a0850152610f178282610e28565b95945050505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b606081526000610f7d606083018789610f20565b8560208401528281036040840152610f96818587610f20565b98975050505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff8251166020820152602082015160408201526000604083015160c06060840152610fe860e0840182610e28565b905060608401516080840152608084015160a084015260a08401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08483030160c0850152610f178282610e28565b73ffffffffffffffffffffffffffffffffffffffff8816815286602082015260a06040820152600061106d60a083018789610f20565b8560608401528281036080840152611086818587610f20565b9a9950505050505050505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126110c957600080fd5b83018035915067ffffffffffffffff8211156110e457600080fd5b602001915036819003821315610c2f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8a16815260c06020820152600061112960c083018a8c610f20565b8860408401528760608401528281036080840152611148818789610f20565b905082810360a084015261115d818587610f20565b9c9b50505050505050505050505056fea264697066735822122057b3906b9c4a14331cb65faa9345c66574dfae396c004154901d907e963343e664736f6c634300081d0033",
}

// MuseConnectorMEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseConnectorMEVMMetaData.ABI instead.
var MuseConnectorMEVMABI = MuseConnectorMEVMMetaData.ABI

// MuseConnectorMEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MuseConnectorMEVMMetaData.Bin instead.
var MuseConnectorMEVMBin = MuseConnectorMEVMMetaData.Bin

// DeployMuseConnectorMEVM deploys a new Ethereum contract, binding an instance of MuseConnectorMEVM to it.
func DeployMuseConnectorMEVM(auth *bind.TransactOpts, backend bind.ContractBackend, wmuse_ common.Address) (common.Address, *types.Transaction, *MuseConnectorMEVM, error) {
	parsed, err := MuseConnectorMEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MuseConnectorMEVMBin), backend, wmuse_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MuseConnectorMEVM{MuseConnectorMEVMCaller: MuseConnectorMEVMCaller{contract: contract}, MuseConnectorMEVMTransactor: MuseConnectorMEVMTransactor{contract: contract}, MuseConnectorMEVMFilterer: MuseConnectorMEVMFilterer{contract: contract}}, nil
}

// MuseConnectorMEVM is an auto generated Go binding around an Ethereum contract.
type MuseConnectorMEVM struct {
	MuseConnectorMEVMCaller     // Read-only binding to the contract
	MuseConnectorMEVMTransactor // Write-only binding to the contract
	MuseConnectorMEVMFilterer   // Log filterer for contract events
}

// MuseConnectorMEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseConnectorMEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorMEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseConnectorMEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorMEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseConnectorMEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorMEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseConnectorMEVMSession struct {
	Contract     *MuseConnectorMEVM // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MuseConnectorMEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseConnectorMEVMCallerSession struct {
	Contract *MuseConnectorMEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MuseConnectorMEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseConnectorMEVMTransactorSession struct {
	Contract     *MuseConnectorMEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MuseConnectorMEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseConnectorMEVMRaw struct {
	Contract *MuseConnectorMEVM // Generic contract binding to access the raw methods on
}

// MuseConnectorMEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseConnectorMEVMCallerRaw struct {
	Contract *MuseConnectorMEVMCaller // Generic read-only contract binding to access the raw methods on
}

// MuseConnectorMEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseConnectorMEVMTransactorRaw struct {
	Contract *MuseConnectorMEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseConnectorMEVM creates a new instance of MuseConnectorMEVM, bound to a specific deployed contract.
func NewMuseConnectorMEVM(address common.Address, backend bind.ContractBackend) (*MuseConnectorMEVM, error) {
	contract, err := bindMuseConnectorMEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVM{MuseConnectorMEVMCaller: MuseConnectorMEVMCaller{contract: contract}, MuseConnectorMEVMTransactor: MuseConnectorMEVMTransactor{contract: contract}, MuseConnectorMEVMFilterer: MuseConnectorMEVMFilterer{contract: contract}}, nil
}

// NewMuseConnectorMEVMCaller creates a new read-only instance of MuseConnectorMEVM, bound to a specific deployed contract.
func NewMuseConnectorMEVMCaller(address common.Address, caller bind.ContractCaller) (*MuseConnectorMEVMCaller, error) {
	contract, err := bindMuseConnectorMEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMCaller{contract: contract}, nil
}

// NewMuseConnectorMEVMTransactor creates a new write-only instance of MuseConnectorMEVM, bound to a specific deployed contract.
func NewMuseConnectorMEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseConnectorMEVMTransactor, error) {
	contract, err := bindMuseConnectorMEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMTransactor{contract: contract}, nil
}

// NewMuseConnectorMEVMFilterer creates a new log filterer instance of MuseConnectorMEVM, bound to a specific deployed contract.
func NewMuseConnectorMEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseConnectorMEVMFilterer, error) {
	contract, err := bindMuseConnectorMEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMFilterer{contract: contract}, nil
}

// bindMuseConnectorMEVM binds a generic wrapper to an already deployed contract.
func bindMuseConnectorMEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseConnectorMEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnectorMEVM *MuseConnectorMEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnectorMEVM.Contract.MuseConnectorMEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnectorMEVM *MuseConnectorMEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.MuseConnectorMEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnectorMEVM *MuseConnectorMEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.MuseConnectorMEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnectorMEVM *MuseConnectorMEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnectorMEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.contract.Transact(opts, method, params...)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMCaller) FUNGIBLEMODULEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseConnectorMEVM.contract.Call(opts, &out, "FUNGIBLE_MODULE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _MuseConnectorMEVM.Contract.FUNGIBLEMODULEADDRESS(&_MuseConnectorMEVM.CallOpts)
}

// FUNGIBLEMODULEADDRESS is a free data retrieval call binding the contract method 0x3ce4a5bc.
//
// Solidity: function FUNGIBLE_MODULE_ADDRESS() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMCallerSession) FUNGIBLEMODULEADDRESS() (common.Address, error) {
	return _MuseConnectorMEVM.Contract.FUNGIBLEMODULEADDRESS(&_MuseConnectorMEVM.CallOpts)
}

// Wmuse is a free data retrieval call binding the contract method 0x69a64637.
//
// Solidity: function wmuse() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMCaller) Wmuse(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseConnectorMEVM.contract.Call(opts, &out, "wmuse")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wmuse is a free data retrieval call binding the contract method 0x69a64637.
//
// Solidity: function wmuse() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) Wmuse() (common.Address, error) {
	return _MuseConnectorMEVM.Contract.Wmuse(&_MuseConnectorMEVM.CallOpts)
}

// Wmuse is a free data retrieval call binding the contract method 0x69a64637.
//
// Solidity: function wmuse() view returns(address)
func (_MuseConnectorMEVM *MuseConnectorMEVMCallerSession) Wmuse() (common.Address, error) {
	return _MuseConnectorMEVM.Contract.Wmuse(&_MuseConnectorMEVM.CallOpts)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes museTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 museValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactor) OnReceive(opts *bind.TransactOpts, museTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, museValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.contract.Transact(opts, "onReceive", museTxSenderAddress, sourceChainId, destinationAddress, museValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes museTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 museValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) OnReceive(museTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, museValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.OnReceive(&_MuseConnectorMEVM.TransactOpts, museTxSenderAddress, sourceChainId, destinationAddress, museValue, message, internalSendHash)
}

// OnReceive is a paid mutator transaction binding the contract method 0x29dd214d.
//
// Solidity: function onReceive(bytes museTxSenderAddress, uint256 sourceChainId, address destinationAddress, uint256 museValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorSession) OnReceive(museTxSenderAddress []byte, sourceChainId *big.Int, destinationAddress common.Address, museValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.OnReceive(&_MuseConnectorMEVM.TransactOpts, museTxSenderAddress, sourceChainId, destinationAddress, museValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address museTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingMuseValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactor) OnRevert(opts *bind.TransactOpts, museTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingMuseValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.contract.Transact(opts, "onRevert", museTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingMuseValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address museTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingMuseValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) OnRevert(museTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingMuseValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.OnRevert(&_MuseConnectorMEVM.TransactOpts, museTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingMuseValue, message, internalSendHash)
}

// OnRevert is a paid mutator transaction binding the contract method 0x942a5e16.
//
// Solidity: function onRevert(address museTxSenderAddress, uint256 sourceChainId, bytes destinationAddress, uint256 destinationChainId, uint256 remainingMuseValue, bytes message, bytes32 internalSendHash) payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorSession) OnRevert(museTxSenderAddress common.Address, sourceChainId *big.Int, destinationAddress []byte, destinationChainId *big.Int, remainingMuseValue *big.Int, message []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.OnRevert(&_MuseConnectorMEVM.TransactOpts, museTxSenderAddress, sourceChainId, destinationAddress, destinationChainId, remainingMuseValue, message, internalSendHash)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactor) Send(opts *bind.TransactOpts, input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnectorMEVM.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) Send(input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.Send(&_MuseConnectorMEVM.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorSession) Send(input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.Send(&_MuseConnectorMEVM.TransactOpts, input)
}

// SetWmuseAddress is a paid mutator transaction binding the contract method 0x6e4b3d52.
//
// Solidity: function setWmuseAddress(address wmuse_) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactor) SetWmuseAddress(opts *bind.TransactOpts, wmuse_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorMEVM.contract.Transact(opts, "setWmuseAddress", wmuse_)
}

// SetWmuseAddress is a paid mutator transaction binding the contract method 0x6e4b3d52.
//
// Solidity: function setWmuseAddress(address wmuse_) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) SetWmuseAddress(wmuse_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.SetWmuseAddress(&_MuseConnectorMEVM.TransactOpts, wmuse_)
}

// SetWmuseAddress is a paid mutator transaction binding the contract method 0x6e4b3d52.
//
// Solidity: function setWmuseAddress(address wmuse_) returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorSession) SetWmuseAddress(wmuse_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.SetWmuseAddress(&_MuseConnectorMEVM.TransactOpts, wmuse_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorMEVM.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMSession) Receive() (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.Receive(&_MuseConnectorMEVM.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MuseConnectorMEVM *MuseConnectorMEVMTransactorSession) Receive() (*types.Transaction, error) {
	return _MuseConnectorMEVM.Contract.Receive(&_MuseConnectorMEVM.TransactOpts)
}

// MuseConnectorMEVMMuseReceivedIterator is returned from FilterMuseReceived and is used to iterate over the raw logs and unpacked data for MuseReceived events raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseReceivedIterator struct {
	Event *MuseConnectorMEVMMuseReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MuseConnectorMEVMMuseReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorMEVMMuseReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MuseConnectorMEVMMuseReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MuseConnectorMEVMMuseReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorMEVMMuseReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorMEVMMuseReceived represents a MuseReceived event raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseReceived struct {
	MuseTxSenderAddress []byte
	SourceChainId       *big.Int
	DestinationAddress  common.Address
	MuseValue           *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMuseReceived is a free log retrieval operation binding the contract event 0xf430afc8ea78e97aef2646b9aee7afd691fcb57b04b185e1430e279a518a8b3c.
//
// Solidity: event MuseReceived(bytes museTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 museValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) FilterMuseReceived(opts *bind.FilterOpts, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (*MuseConnectorMEVMMuseReceivedIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.FilterLogs(opts, "MuseReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMMuseReceivedIterator{contract: _MuseConnectorMEVM.contract, event: "MuseReceived", logs: logs, sub: sub}, nil
}

// WatchMuseReceived is a free log subscription operation binding the contract event 0xf430afc8ea78e97aef2646b9aee7afd691fcb57b04b185e1430e279a518a8b3c.
//
// Solidity: event MuseReceived(bytes museTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 museValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) WatchMuseReceived(opts *bind.WatchOpts, sink chan<- *MuseConnectorMEVMMuseReceived, sourceChainId []*big.Int, destinationAddress []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var destinationAddressRule []interface{}
	for _, destinationAddressItem := range destinationAddress {
		destinationAddressRule = append(destinationAddressRule, destinationAddressItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.WatchLogs(opts, "MuseReceived", sourceChainIdRule, destinationAddressRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorMEVMMuseReceived)
				if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMuseReceived is a log parse operation binding the contract event 0xf430afc8ea78e97aef2646b9aee7afd691fcb57b04b185e1430e279a518a8b3c.
//
// Solidity: event MuseReceived(bytes museTxSenderAddress, uint256 indexed sourceChainId, address indexed destinationAddress, uint256 museValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) ParseMuseReceived(log types.Log) (*MuseConnectorMEVMMuseReceived, error) {
	event := new(MuseConnectorMEVMMuseReceived)
	if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorMEVMMuseRevertedIterator is returned from FilterMuseReverted and is used to iterate over the raw logs and unpacked data for MuseReverted events raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseRevertedIterator struct {
	Event *MuseConnectorMEVMMuseReverted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MuseConnectorMEVMMuseRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorMEVMMuseReverted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MuseConnectorMEVMMuseReverted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MuseConnectorMEVMMuseRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorMEVMMuseRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorMEVMMuseReverted represents a MuseReverted event raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseReverted struct {
	MuseTxSenderAddress common.Address
	SourceChainId       *big.Int
	DestinationChainId  *big.Int
	DestinationAddress  []byte
	RemainingMuseValue  *big.Int
	Message             []byte
	InternalSendHash    [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMuseReverted is a free log retrieval operation binding the contract event 0x5d921d2b8faf864052d4535b06cf528da89b7227158f2b251b54189d2e77411c.
//
// Solidity: event MuseReverted(address museTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingMuseValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) FilterMuseReverted(opts *bind.FilterOpts, destinationChainId []*big.Int, internalSendHash [][32]byte) (*MuseConnectorMEVMMuseRevertedIterator, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.FilterLogs(opts, "MuseReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMMuseRevertedIterator{contract: _MuseConnectorMEVM.contract, event: "MuseReverted", logs: logs, sub: sub}, nil
}

// WatchMuseReverted is a free log subscription operation binding the contract event 0x5d921d2b8faf864052d4535b06cf528da89b7227158f2b251b54189d2e77411c.
//
// Solidity: event MuseReverted(address museTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingMuseValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) WatchMuseReverted(opts *bind.WatchOpts, sink chan<- *MuseConnectorMEVMMuseReverted, destinationChainId []*big.Int, internalSendHash [][32]byte) (event.Subscription, error) {

	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.WatchLogs(opts, "MuseReverted", destinationChainIdRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorMEVMMuseReverted)
				if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseReverted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMuseReverted is a log parse operation binding the contract event 0x5d921d2b8faf864052d4535b06cf528da89b7227158f2b251b54189d2e77411c.
//
// Solidity: event MuseReverted(address museTxSenderAddress, uint256 sourceChainId, uint256 indexed destinationChainId, bytes destinationAddress, uint256 remainingMuseValue, bytes message, bytes32 indexed internalSendHash)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) ParseMuseReverted(log types.Log) (*MuseConnectorMEVMMuseReverted, error) {
	event := new(MuseConnectorMEVMMuseReverted)
	if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorMEVMMuseSentIterator is returned from FilterMuseSent and is used to iterate over the raw logs and unpacked data for MuseSent events raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseSentIterator struct {
	Event *MuseConnectorMEVMMuseSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MuseConnectorMEVMMuseSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorMEVMMuseSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MuseConnectorMEVMMuseSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MuseConnectorMEVMMuseSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorMEVMMuseSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorMEVMMuseSent represents a MuseSent event raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMMuseSent struct {
	SourceTxOriginAddress common.Address
	MuseTxSenderAddress   common.Address
	DestinationChainId    *big.Int
	DestinationAddress    []byte
	MuseValueAndGas       *big.Int
	DestinationGasLimit   *big.Int
	Message               []byte
	MuseParams            []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMuseSent is a free log retrieval operation binding the contract event 0x9bf48436256611f46f7d8bbefc3abe994fdf7da86f3829eb6c73f8dd5d9ef27d.
//
// Solidity: event MuseSent(address sourceTxOriginAddress, address indexed museTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 museValueAndGas, uint256 destinationGasLimit, bytes message, bytes museParams)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) FilterMuseSent(opts *bind.FilterOpts, museTxSenderAddress []common.Address, destinationChainId []*big.Int) (*MuseConnectorMEVMMuseSentIterator, error) {

	var museTxSenderAddressRule []interface{}
	for _, museTxSenderAddressItem := range museTxSenderAddress {
		museTxSenderAddressRule = append(museTxSenderAddressRule, museTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.FilterLogs(opts, "MuseSent", museTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMMuseSentIterator{contract: _MuseConnectorMEVM.contract, event: "MuseSent", logs: logs, sub: sub}, nil
}

// WatchMuseSent is a free log subscription operation binding the contract event 0x9bf48436256611f46f7d8bbefc3abe994fdf7da86f3829eb6c73f8dd5d9ef27d.
//
// Solidity: event MuseSent(address sourceTxOriginAddress, address indexed museTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 museValueAndGas, uint256 destinationGasLimit, bytes message, bytes museParams)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) WatchMuseSent(opts *bind.WatchOpts, sink chan<- *MuseConnectorMEVMMuseSent, museTxSenderAddress []common.Address, destinationChainId []*big.Int) (event.Subscription, error) {

	var museTxSenderAddressRule []interface{}
	for _, museTxSenderAddressItem := range museTxSenderAddress {
		museTxSenderAddressRule = append(museTxSenderAddressRule, museTxSenderAddressItem)
	}
	var destinationChainIdRule []interface{}
	for _, destinationChainIdItem := range destinationChainId {
		destinationChainIdRule = append(destinationChainIdRule, destinationChainIdItem)
	}

	logs, sub, err := _MuseConnectorMEVM.contract.WatchLogs(opts, "MuseSent", museTxSenderAddressRule, destinationChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorMEVMMuseSent)
				if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMuseSent is a log parse operation binding the contract event 0x9bf48436256611f46f7d8bbefc3abe994fdf7da86f3829eb6c73f8dd5d9ef27d.
//
// Solidity: event MuseSent(address sourceTxOriginAddress, address indexed museTxSenderAddress, uint256 indexed destinationChainId, bytes destinationAddress, uint256 museValueAndGas, uint256 destinationGasLimit, bytes message, bytes museParams)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) ParseMuseSent(log types.Log) (*MuseConnectorMEVMMuseSent, error) {
	event := new(MuseConnectorMEVMMuseSent)
	if err := _MuseConnectorMEVM.contract.UnpackLog(event, "MuseSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorMEVMSetWMUSEIterator is returned from FilterSetWMUSE and is used to iterate over the raw logs and unpacked data for SetWMUSE events raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMSetWMUSEIterator struct {
	Event *MuseConnectorMEVMSetWMUSE // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MuseConnectorMEVMSetWMUSEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorMEVMSetWMUSE)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MuseConnectorMEVMSetWMUSE)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MuseConnectorMEVMSetWMUSEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorMEVMSetWMUSEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorMEVMSetWMUSE represents a SetWMUSE event raised by the MuseConnectorMEVM contract.
type MuseConnectorMEVMSetWMUSE struct {
	Wmuse common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetWMUSE is a free log retrieval operation binding the contract event 0x702607b065ed0b40b11595136269ccf48c3f5786b60c546289e812cb2d947b67.
//
// Solidity: event SetWMUSE(address wmuse_)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) FilterSetWMUSE(opts *bind.FilterOpts) (*MuseConnectorMEVMSetWMUSEIterator, error) {

	logs, sub, err := _MuseConnectorMEVM.contract.FilterLogs(opts, "SetWMUSE")
	if err != nil {
		return nil, err
	}
	return &MuseConnectorMEVMSetWMUSEIterator{contract: _MuseConnectorMEVM.contract, event: "SetWMUSE", logs: logs, sub: sub}, nil
}

// WatchSetWMUSE is a free log subscription operation binding the contract event 0x702607b065ed0b40b11595136269ccf48c3f5786b60c546289e812cb2d947b67.
//
// Solidity: event SetWMUSE(address wmuse_)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) WatchSetWMUSE(opts *bind.WatchOpts, sink chan<- *MuseConnectorMEVMSetWMUSE) (event.Subscription, error) {

	logs, sub, err := _MuseConnectorMEVM.contract.WatchLogs(opts, "SetWMUSE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorMEVMSetWMUSE)
				if err := _MuseConnectorMEVM.contract.UnpackLog(event, "SetWMUSE", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetWMUSE is a log parse operation binding the contract event 0x702607b065ed0b40b11595136269ccf48c3f5786b60c546289e812cb2d947b67.
//
// Solidity: event SetWMUSE(address wmuse_)
func (_MuseConnectorMEVM *MuseConnectorMEVMFilterer) ParseSetWMUSE(log types.Log) (*MuseConnectorMEVMSetWMUSE, error) {
	event := new(MuseConnectorMEVMSetWMUSE)
	if err := _MuseConnectorMEVM.contract.UnpackLog(event, "SetWMUSE", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
