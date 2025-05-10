// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package museinterfaces

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

// MuseConnectorMetaData contains all meta data concerning the MuseConnector contract.
var MuseConnectorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"send\",\"inputs\":[{\"name\":\"input\",\"type\":\"tuple\",\"internalType\":\"structMuseInterfaces.SendInput\",\"components\":[{\"name\":\"destinationChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destinationAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destinationGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"museValueAndGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"museParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// MuseConnectorABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseConnectorMetaData.ABI instead.
var MuseConnectorABI = MuseConnectorMetaData.ABI

// MuseConnector is an auto generated Go binding around an Ethereum contract.
type MuseConnector struct {
	MuseConnectorCaller     // Read-only binding to the contract
	MuseConnectorTransactor // Write-only binding to the contract
	MuseConnectorFilterer   // Log filterer for contract events
}

// MuseConnectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseConnectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseConnectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseConnectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseConnectorSession struct {
	Contract     *MuseConnector    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuseConnectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseConnectorCallerSession struct {
	Contract *MuseConnectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MuseConnectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseConnectorTransactorSession struct {
	Contract     *MuseConnectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MuseConnectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseConnectorRaw struct {
	Contract *MuseConnector // Generic contract binding to access the raw methods on
}

// MuseConnectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseConnectorCallerRaw struct {
	Contract *MuseConnectorCaller // Generic read-only contract binding to access the raw methods on
}

// MuseConnectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseConnectorTransactorRaw struct {
	Contract *MuseConnectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseConnector creates a new instance of MuseConnector, bound to a specific deployed contract.
func NewMuseConnector(address common.Address, backend bind.ContractBackend) (*MuseConnector, error) {
	contract, err := bindMuseConnector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseConnector{MuseConnectorCaller: MuseConnectorCaller{contract: contract}, MuseConnectorTransactor: MuseConnectorTransactor{contract: contract}, MuseConnectorFilterer: MuseConnectorFilterer{contract: contract}}, nil
}

// NewMuseConnectorCaller creates a new read-only instance of MuseConnector, bound to a specific deployed contract.
func NewMuseConnectorCaller(address common.Address, caller bind.ContractCaller) (*MuseConnectorCaller, error) {
	contract, err := bindMuseConnector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorCaller{contract: contract}, nil
}

// NewMuseConnectorTransactor creates a new write-only instance of MuseConnector, bound to a specific deployed contract.
func NewMuseConnectorTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseConnectorTransactor, error) {
	contract, err := bindMuseConnector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorTransactor{contract: contract}, nil
}

// NewMuseConnectorFilterer creates a new log filterer instance of MuseConnector, bound to a specific deployed contract.
func NewMuseConnectorFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseConnectorFilterer, error) {
	contract, err := bindMuseConnector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorFilterer{contract: contract}, nil
}

// bindMuseConnector binds a generic wrapper to an already deployed contract.
func bindMuseConnector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseConnectorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnector *MuseConnectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnector.Contract.MuseConnectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnector *MuseConnectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnector.Contract.MuseConnectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnector *MuseConnectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnector.Contract.MuseConnectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnector *MuseConnectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnector *MuseConnectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnector *MuseConnectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnector.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnector *MuseConnectorTransactor) Send(opts *bind.TransactOpts, input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnector.contract.Transact(opts, "send", input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnector *MuseConnectorSession) Send(input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnector.Contract.Send(&_MuseConnector.TransactOpts, input)
}

// Send is a paid mutator transaction binding the contract method 0xec026901.
//
// Solidity: function send((uint256,bytes,uint256,bytes,uint256,bytes) input) returns()
func (_MuseConnector *MuseConnectorTransactorSession) Send(input MuseInterfacesSendInput) (*types.Transaction, error) {
	return _MuseConnector.Contract.Send(&_MuseConnector.TransactOpts, input)
}
