// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package muse

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

// MuseErrorsMetaData contains all meta data concerning the MuseErrors contract.
var MuseErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"CallerIsNotConnector\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTss\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssOrUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MuseTransferError\",\"inputs\":[]}]",
}

// MuseErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseErrorsMetaData.ABI instead.
var MuseErrorsABI = MuseErrorsMetaData.ABI

// MuseErrors is an auto generated Go binding around an Ethereum contract.
type MuseErrors struct {
	MuseErrorsCaller     // Read-only binding to the contract
	MuseErrorsTransactor // Write-only binding to the contract
	MuseErrorsFilterer   // Log filterer for contract events
}

// MuseErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseErrorsSession struct {
	Contract     *MuseErrors       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuseErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseErrorsCallerSession struct {
	Contract *MuseErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MuseErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseErrorsTransactorSession struct {
	Contract     *MuseErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MuseErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseErrorsRaw struct {
	Contract *MuseErrors // Generic contract binding to access the raw methods on
}

// MuseErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseErrorsCallerRaw struct {
	Contract *MuseErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// MuseErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseErrorsTransactorRaw struct {
	Contract *MuseErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseErrors creates a new instance of MuseErrors, bound to a specific deployed contract.
func NewMuseErrors(address common.Address, backend bind.ContractBackend) (*MuseErrors, error) {
	contract, err := bindMuseErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseErrors{MuseErrorsCaller: MuseErrorsCaller{contract: contract}, MuseErrorsTransactor: MuseErrorsTransactor{contract: contract}, MuseErrorsFilterer: MuseErrorsFilterer{contract: contract}}, nil
}

// NewMuseErrorsCaller creates a new read-only instance of MuseErrors, bound to a specific deployed contract.
func NewMuseErrorsCaller(address common.Address, caller bind.ContractCaller) (*MuseErrorsCaller, error) {
	contract, err := bindMuseErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseErrorsCaller{contract: contract}, nil
}

// NewMuseErrorsTransactor creates a new write-only instance of MuseErrors, bound to a specific deployed contract.
func NewMuseErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseErrorsTransactor, error) {
	contract, err := bindMuseErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseErrorsTransactor{contract: contract}, nil
}

// NewMuseErrorsFilterer creates a new log filterer instance of MuseErrors, bound to a specific deployed contract.
func NewMuseErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseErrorsFilterer, error) {
	contract, err := bindMuseErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseErrorsFilterer{contract: contract}, nil
}

// bindMuseErrors binds a generic wrapper to an already deployed contract.
func bindMuseErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseErrors *MuseErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseErrors.Contract.MuseErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseErrors *MuseErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseErrors.Contract.MuseErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseErrors *MuseErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseErrors.Contract.MuseErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseErrors *MuseErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseErrors *MuseErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseErrors *MuseErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseErrors.Contract.contract.Transact(opts, method, params...)
}
