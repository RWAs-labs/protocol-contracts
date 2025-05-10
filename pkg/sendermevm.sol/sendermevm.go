// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sendermevm

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

// SenderMEVMMetaData contains all meta data concerning the SenderMEVM contract.
var SenderMEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawAndCallReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"num\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ApprovalFailed\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b50604051610aad380380610aad833981016040819052602c916050565b600080546001600160a01b0319166001600160a01b0392909216919091179055607e565b600060208284031215606157600080fd5b81516001600160a01b0381168114607757600080fd5b9392505050565b610a208061008d6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630abd890514610046578063116191b61461005b5780637a34d8bb146100a4575b600080fd5b61005961005436600461065c565b6100b7565b005b60005461007b9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100596100b23660046106fc565b61032f565b60008383836040516024016100ce939291906107f8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f970000000000000000000000000000000000000000000000000000000017905260005490915073ffffffffffffffffffffffffffffffffffffffff8087169163095ea7b3911661017789620186a0610822565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156101e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020b9190610862565b610241576040517f8164f84200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160a08101825261032180825260016020808401919091528284019190915282518082018452600080825260608401919091526080830181905283518085018552620186a081529182018190525492517f7b15118b0000000000000000000000000000000000000000000000000000000081529192909173ffffffffffffffffffffffffffffffffffffffff90911690637b15118b906102f2908c908c908c90899088908a906004016108fb565b600060405180830381600087803b15801561030c57600080fd5b505af1158015610320573d6000803e3d6000fd5b50505050505050505050505050565b6000838383604051602401610346939291906107f8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fe04d4f9700000000000000000000000000000000000000000000000000000000179052815160a0810183526103218082526001828401528184015282518083018452600080825260608301919091526080820181905283518085018552620186a0808252938101829052905493517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff94851660048201526024810193909352939450929188169063095ea7b3906044016020604051808303816000875af1158015610480573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104a49190610862565b506000546040517f06cb898300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff909116906306cb898390610503908b908b90889087908990600401610976565b600060405180830381600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261057f57600080fd5b81356020830160008067ffffffffffffffff8411156105a0576105a061053f565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156105ed576105ed61053f565b60405283815290508082840187101561060557600080fd5b838360208301376000602085830101528094505050505092915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461064657600080fd5b919050565b801515811461065957600080fd5b50565b60008060008060008060c0878903121561067557600080fd5b863567ffffffffffffffff81111561068c57600080fd5b61069889828a0161056e565b965050602087013594506106ae60408801610622565b9350606087013567ffffffffffffffff8111156106ca57600080fd5b6106d689828a0161056e565b9350506080870135915060a08701356106ee8161064b565b809150509295509295509295565b600080600080600060a0868803121561071457600080fd5b853567ffffffffffffffff81111561072b57600080fd5b6107378882890161056e565b95505061074660208701610622565b9350604086013567ffffffffffffffff81111561076257600080fd5b61076e8882890161056e565b9350506060860135915060808601356107868161064b565b809150509295509295909350565b6000815180845260005b818110156107ba5760208185018101518683018201520161079e565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60608152600061080b6060830186610794565b602083019490945250901515604090910152919050565b8082018082111561085c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60006020828403121561087457600080fd5b815161087f8161064b565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff815116825260208101511515602083015273ffffffffffffffffffffffffffffffffffffffff60408201511660408301526000606082015160a060608501526108e760a0850182610794565b608093840151949093019390935250919050565b60e08152600061090e60e0830189610794565b87602084015273ffffffffffffffffffffffffffffffffffffffff8716604084015282810360608401526109428187610794565b855160808501526020860151151560a0850152905082810360c08401526109698185610886565b9998505050505050505050565b60c08152600061098960c0830188610794565b73ffffffffffffffffffffffffffffffffffffffff8716602084015282810360408401526109b78187610794565b85516060850152602086015115156080850152905082810360a08401526109de8185610886565b9897505050505050505056fea2646970667358221220c1ba37a424a9e81ccacad2835dec6ee8d25b2a59262b3b18a21601b0b7356cc764736f6c634300081a0033",
}

// SenderMEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use SenderMEVMMetaData.ABI instead.
var SenderMEVMABI = SenderMEVMMetaData.ABI

// SenderMEVMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SenderMEVMMetaData.Bin instead.
var SenderMEVMBin = SenderMEVMMetaData.Bin

// DeploySenderMEVM deploys a new Ethereum contract, binding an instance of SenderMEVM to it.
func DeploySenderMEVM(auth *bind.TransactOpts, backend bind.ContractBackend, gateway_ common.Address) (common.Address, *types.Transaction, *SenderMEVM, error) {
	parsed, err := SenderMEVMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SenderMEVMBin), backend, gateway_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SenderMEVM{SenderMEVMCaller: SenderMEVMCaller{contract: contract}, SenderMEVMTransactor: SenderMEVMTransactor{contract: contract}, SenderMEVMFilterer: SenderMEVMFilterer{contract: contract}}, nil
}

// SenderMEVM is an auto generated Go binding around an Ethereum contract.
type SenderMEVM struct {
	SenderMEVMCaller     // Read-only binding to the contract
	SenderMEVMTransactor // Write-only binding to the contract
	SenderMEVMFilterer   // Log filterer for contract events
}

// SenderMEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type SenderMEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderMEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SenderMEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderMEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SenderMEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SenderMEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SenderMEVMSession struct {
	Contract     *SenderMEVM       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SenderMEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SenderMEVMCallerSession struct {
	Contract *SenderMEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SenderMEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SenderMEVMTransactorSession struct {
	Contract     *SenderMEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SenderMEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type SenderMEVMRaw struct {
	Contract *SenderMEVM // Generic contract binding to access the raw methods on
}

// SenderMEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SenderMEVMCallerRaw struct {
	Contract *SenderMEVMCaller // Generic read-only contract binding to access the raw methods on
}

// SenderMEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SenderMEVMTransactorRaw struct {
	Contract *SenderMEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSenderMEVM creates a new instance of SenderMEVM, bound to a specific deployed contract.
func NewSenderMEVM(address common.Address, backend bind.ContractBackend) (*SenderMEVM, error) {
	contract, err := bindSenderMEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SenderMEVM{SenderMEVMCaller: SenderMEVMCaller{contract: contract}, SenderMEVMTransactor: SenderMEVMTransactor{contract: contract}, SenderMEVMFilterer: SenderMEVMFilterer{contract: contract}}, nil
}

// NewSenderMEVMCaller creates a new read-only instance of SenderMEVM, bound to a specific deployed contract.
func NewSenderMEVMCaller(address common.Address, caller bind.ContractCaller) (*SenderMEVMCaller, error) {
	contract, err := bindSenderMEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SenderMEVMCaller{contract: contract}, nil
}

// NewSenderMEVMTransactor creates a new write-only instance of SenderMEVM, bound to a specific deployed contract.
func NewSenderMEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*SenderMEVMTransactor, error) {
	contract, err := bindSenderMEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SenderMEVMTransactor{contract: contract}, nil
}

// NewSenderMEVMFilterer creates a new log filterer instance of SenderMEVM, bound to a specific deployed contract.
func NewSenderMEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*SenderMEVMFilterer, error) {
	contract, err := bindSenderMEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SenderMEVMFilterer{contract: contract}, nil
}

// bindSenderMEVM binds a generic wrapper to an already deployed contract.
func bindSenderMEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SenderMEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderMEVM *SenderMEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderMEVM.Contract.SenderMEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderMEVM *SenderMEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderMEVM.Contract.SenderMEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderMEVM *SenderMEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderMEVM.Contract.SenderMEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SenderMEVM *SenderMEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SenderMEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SenderMEVM *SenderMEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SenderMEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SenderMEVM *SenderMEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SenderMEVM.Contract.contract.Transact(opts, method, params...)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderMEVM *SenderMEVMCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SenderMEVM.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderMEVM *SenderMEVMSession) Gateway() (common.Address, error) {
	return _SenderMEVM.Contract.Gateway(&_SenderMEVM.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_SenderMEVM *SenderMEVMCallerSession) Gateway() (common.Address, error) {
	return _SenderMEVM.Contract.Gateway(&_SenderMEVM.CallOpts)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMTransactor) CallReceiver(opts *bind.TransactOpts, receiver []byte, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.contract.Transact(opts, "callReceiver", receiver, mrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMSession) CallReceiver(receiver []byte, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.Contract.CallReceiver(&_SenderMEVM.TransactOpts, receiver, mrc20, str, num, flag)
}

// CallReceiver is a paid mutator transaction binding the contract method 0x7a34d8bb.
//
// Solidity: function callReceiver(bytes receiver, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMTransactorSession) CallReceiver(receiver []byte, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.Contract.CallReceiver(&_SenderMEVM.TransactOpts, receiver, mrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMTransactor) WithdrawAndCallReceiver(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.contract.Transact(opts, "withdrawAndCallReceiver", receiver, amount, mrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.Contract.WithdrawAndCallReceiver(&_SenderMEVM.TransactOpts, receiver, amount, mrc20, str, num, flag)
}

// WithdrawAndCallReceiver is a paid mutator transaction binding the contract method 0x0abd8905.
//
// Solidity: function withdrawAndCallReceiver(bytes receiver, uint256 amount, address mrc20, string str, uint256 num, bool flag) returns()
func (_SenderMEVM *SenderMEVMTransactorSession) WithdrawAndCallReceiver(receiver []byte, amount *big.Int, mrc20 common.Address, str string, num *big.Int, flag bool) (*types.Transaction, error) {
	return _SenderMEVM.Contract.WithdrawAndCallReceiver(&_SenderMEVM.TransactOpts, receiver, amount, mrc20, str, num, flag)
}
