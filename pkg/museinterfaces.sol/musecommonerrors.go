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

// MuseCommonErrorsMetaData contains all meta data concerning the MuseCommonErrors contract.
var MuseCommonErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]}]",
}

// MuseCommonErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseCommonErrorsMetaData.ABI instead.
var MuseCommonErrorsABI = MuseCommonErrorsMetaData.ABI

// MuseCommonErrors is an auto generated Go binding around an Ethereum contract.
type MuseCommonErrors struct {
	MuseCommonErrorsCaller     // Read-only binding to the contract
	MuseCommonErrorsTransactor // Write-only binding to the contract
	MuseCommonErrorsFilterer   // Log filterer for contract events
}

// MuseCommonErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseCommonErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseCommonErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseCommonErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseCommonErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseCommonErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseCommonErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseCommonErrorsSession struct {
	Contract     *MuseCommonErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuseCommonErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseCommonErrorsCallerSession struct {
	Contract *MuseCommonErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MuseCommonErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseCommonErrorsTransactorSession struct {
	Contract     *MuseCommonErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MuseCommonErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseCommonErrorsRaw struct {
	Contract *MuseCommonErrors // Generic contract binding to access the raw methods on
}

// MuseCommonErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseCommonErrorsCallerRaw struct {
	Contract *MuseCommonErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// MuseCommonErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseCommonErrorsTransactorRaw struct {
	Contract *MuseCommonErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseCommonErrors creates a new instance of MuseCommonErrors, bound to a specific deployed contract.
func NewMuseCommonErrors(address common.Address, backend bind.ContractBackend) (*MuseCommonErrors, error) {
	contract, err := bindMuseCommonErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseCommonErrors{MuseCommonErrorsCaller: MuseCommonErrorsCaller{contract: contract}, MuseCommonErrorsTransactor: MuseCommonErrorsTransactor{contract: contract}, MuseCommonErrorsFilterer: MuseCommonErrorsFilterer{contract: contract}}, nil
}

// NewMuseCommonErrorsCaller creates a new read-only instance of MuseCommonErrors, bound to a specific deployed contract.
func NewMuseCommonErrorsCaller(address common.Address, caller bind.ContractCaller) (*MuseCommonErrorsCaller, error) {
	contract, err := bindMuseCommonErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseCommonErrorsCaller{contract: contract}, nil
}

// NewMuseCommonErrorsTransactor creates a new write-only instance of MuseCommonErrors, bound to a specific deployed contract.
func NewMuseCommonErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseCommonErrorsTransactor, error) {
	contract, err := bindMuseCommonErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseCommonErrorsTransactor{contract: contract}, nil
}

// NewMuseCommonErrorsFilterer creates a new log filterer instance of MuseCommonErrors, bound to a specific deployed contract.
func NewMuseCommonErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseCommonErrorsFilterer, error) {
	contract, err := bindMuseCommonErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseCommonErrorsFilterer{contract: contract}, nil
}

// bindMuseCommonErrors binds a generic wrapper to an already deployed contract.
func bindMuseCommonErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseCommonErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseCommonErrors *MuseCommonErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseCommonErrors.Contract.MuseCommonErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseCommonErrors *MuseCommonErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseCommonErrors.Contract.MuseCommonErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseCommonErrors *MuseCommonErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseCommonErrors.Contract.MuseCommonErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseCommonErrors *MuseCommonErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseCommonErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseCommonErrors *MuseCommonErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseCommonErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseCommonErrors *MuseCommonErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseCommonErrors.Contract.contract.Transact(opts, method, params...)
}
