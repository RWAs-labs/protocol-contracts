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

// MuseInterfacesMetaData contains all meta data concerning the MuseInterfaces contract.
var MuseInterfacesMetaData = &bind.MetaData{
	ABI: "[]",
}

// MuseInterfacesABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseInterfacesMetaData.ABI instead.
var MuseInterfacesABI = MuseInterfacesMetaData.ABI

// MuseInterfaces is an auto generated Go binding around an Ethereum contract.
type MuseInterfaces struct {
	MuseInterfacesCaller     // Read-only binding to the contract
	MuseInterfacesTransactor // Write-only binding to the contract
	MuseInterfacesFilterer   // Log filterer for contract events
}

// MuseInterfacesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseInterfacesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseInterfacesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseInterfacesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseInterfacesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseInterfacesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseInterfacesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseInterfacesSession struct {
	Contract     *MuseInterfaces   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuseInterfacesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseInterfacesCallerSession struct {
	Contract *MuseInterfacesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MuseInterfacesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseInterfacesTransactorSession struct {
	Contract     *MuseInterfacesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MuseInterfacesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseInterfacesRaw struct {
	Contract *MuseInterfaces // Generic contract binding to access the raw methods on
}

// MuseInterfacesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseInterfacesCallerRaw struct {
	Contract *MuseInterfacesCaller // Generic read-only contract binding to access the raw methods on
}

// MuseInterfacesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseInterfacesTransactorRaw struct {
	Contract *MuseInterfacesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseInterfaces creates a new instance of MuseInterfaces, bound to a specific deployed contract.
func NewMuseInterfaces(address common.Address, backend bind.ContractBackend) (*MuseInterfaces, error) {
	contract, err := bindMuseInterfaces(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseInterfaces{MuseInterfacesCaller: MuseInterfacesCaller{contract: contract}, MuseInterfacesTransactor: MuseInterfacesTransactor{contract: contract}, MuseInterfacesFilterer: MuseInterfacesFilterer{contract: contract}}, nil
}

// NewMuseInterfacesCaller creates a new read-only instance of MuseInterfaces, bound to a specific deployed contract.
func NewMuseInterfacesCaller(address common.Address, caller bind.ContractCaller) (*MuseInterfacesCaller, error) {
	contract, err := bindMuseInterfaces(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseInterfacesCaller{contract: contract}, nil
}

// NewMuseInterfacesTransactor creates a new write-only instance of MuseInterfaces, bound to a specific deployed contract.
func NewMuseInterfacesTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseInterfacesTransactor, error) {
	contract, err := bindMuseInterfaces(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseInterfacesTransactor{contract: contract}, nil
}

// NewMuseInterfacesFilterer creates a new log filterer instance of MuseInterfaces, bound to a specific deployed contract.
func NewMuseInterfacesFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseInterfacesFilterer, error) {
	contract, err := bindMuseInterfaces(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseInterfacesFilterer{contract: contract}, nil
}

// bindMuseInterfaces binds a generic wrapper to an already deployed contract.
func bindMuseInterfaces(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseInterfacesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseInterfaces *MuseInterfacesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseInterfaces.Contract.MuseInterfacesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseInterfaces *MuseInterfacesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseInterfaces.Contract.MuseInterfacesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseInterfaces *MuseInterfacesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseInterfaces.Contract.MuseInterfacesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseInterfaces *MuseInterfacesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseInterfaces.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseInterfaces *MuseInterfacesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseInterfaces.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseInterfaces *MuseInterfacesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseInterfaces.Contract.contract.Transact(opts, method, params...)
}
