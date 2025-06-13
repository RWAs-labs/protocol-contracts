// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalcontract

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

// MContext is an auto generated low-level Go binding around an user-defined struct.
type MContext struct {
	Origin  []byte
	Sender  common.Address
	ChainID *big.Int
}

// MContractMetaData contains all meta data concerning the MContract contract.
var MContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"onCrossChainCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structmContext\",\"components\":[{\"name\":\"origin\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// MContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MContractMetaData.ABI instead.
var MContractABI = MContractMetaData.ABI

// MContract is an auto generated Go binding around an Ethereum contract.
type MContract struct {
	MContractCaller     // Read-only binding to the contract
	MContractTransactor // Write-only binding to the contract
	MContractFilterer   // Log filterer for contract events
}

// MContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MContractSession struct {
	Contract     *MContract        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MContractCallerSession struct {
	Contract *MContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MContractTransactorSession struct {
	Contract     *MContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MContractRaw struct {
	Contract *MContract // Generic contract binding to access the raw methods on
}

// MContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MContractCallerRaw struct {
	Contract *MContractCaller // Generic read-only contract binding to access the raw methods on
}

// MContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MContractTransactorRaw struct {
	Contract *MContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMContract creates a new instance of MContract, bound to a specific deployed contract.
func NewMContract(address common.Address, backend bind.ContractBackend) (*MContract, error) {
	contract, err := bindMContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MContract{MContractCaller: MContractCaller{contract: contract}, MContractTransactor: MContractTransactor{contract: contract}, MContractFilterer: MContractFilterer{contract: contract}}, nil
}

// NewMContractCaller creates a new read-only instance of MContract, bound to a specific deployed contract.
func NewMContractCaller(address common.Address, caller bind.ContractCaller) (*MContractCaller, error) {
	contract, err := bindMContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MContractCaller{contract: contract}, nil
}

// NewMContractTransactor creates a new write-only instance of MContract, bound to a specific deployed contract.
func NewMContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MContractTransactor, error) {
	contract, err := bindMContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MContractTransactor{contract: contract}, nil
}

// NewMContractFilterer creates a new log filterer instance of MContract, bound to a specific deployed contract.
func NewMContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MContractFilterer, error) {
	contract, err := bindMContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MContractFilterer{contract: contract}, nil
}

// bindMContract binds a generic wrapper to an already deployed contract.
func bindMContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MContract *MContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MContract.Contract.MContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MContract *MContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MContract.Contract.MContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MContract *MContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MContract.Contract.MContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MContract *MContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MContract *MContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MContract *MContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MContract.Contract.contract.Transact(opts, method, params...)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address mrc20, uint256 amount, bytes message) returns()
func (_MContract *MContractTransactor) OnCrossChainCall(opts *bind.TransactOpts, context MContext, mrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _MContract.contract.Transact(opts, "onCrossChainCall", context, mrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address mrc20, uint256 amount, bytes message) returns()
func (_MContract *MContractSession) OnCrossChainCall(context MContext, mrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _MContract.Contract.OnCrossChainCall(&_MContract.TransactOpts, context, mrc20, amount, message)
}

// OnCrossChainCall is a paid mutator transaction binding the contract method 0xde43156e.
//
// Solidity: function onCrossChainCall((bytes,address,uint256) context, address mrc20, uint256 amount, bytes message) returns()
func (_MContract *MContractTransactorSession) OnCrossChainCall(context MContext, mrc20 common.Address, amount *big.Int, message []byte) (*types.Transaction, error) {
	return _MContract.Contract.OnCrossChainCall(&_MContract.TransactOpts, context, mrc20, amount, message)
}
