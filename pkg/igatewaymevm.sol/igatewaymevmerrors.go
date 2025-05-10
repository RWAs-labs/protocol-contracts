// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package igatewaymevm

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

// IGatewayMEVMErrorsMetaData contains all meta data concerning the IGatewayMEVMErrors contract.
var IGatewayMEVMErrorsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedMuseSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMuseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWMUSEOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20BurnFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20DepositFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20TransferFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// IGatewayMEVMErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayMEVMErrorsMetaData.ABI instead.
var IGatewayMEVMErrorsABI = IGatewayMEVMErrorsMetaData.ABI

// IGatewayMEVMErrors is an auto generated Go binding around an Ethereum contract.
type IGatewayMEVMErrors struct {
	IGatewayMEVMErrorsCaller     // Read-only binding to the contract
	IGatewayMEVMErrorsTransactor // Write-only binding to the contract
	IGatewayMEVMErrorsFilterer   // Log filterer for contract events
}

// IGatewayMEVMErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayMEVMErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayMEVMErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayMEVMErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayMEVMErrorsSession struct {
	Contract     *IGatewayMEVMErrors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGatewayMEVMErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayMEVMErrorsCallerSession struct {
	Contract *IGatewayMEVMErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IGatewayMEVMErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayMEVMErrorsTransactorSession struct {
	Contract     *IGatewayMEVMErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IGatewayMEVMErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayMEVMErrorsRaw struct {
	Contract *IGatewayMEVMErrors // Generic contract binding to access the raw methods on
}

// IGatewayMEVMErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayMEVMErrorsCallerRaw struct {
	Contract *IGatewayMEVMErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayMEVMErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayMEVMErrorsTransactorRaw struct {
	Contract *IGatewayMEVMErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayMEVMErrors creates a new instance of IGatewayMEVMErrors, bound to a specific deployed contract.
func NewIGatewayMEVMErrors(address common.Address, backend bind.ContractBackend) (*IGatewayMEVMErrors, error) {
	contract, err := bindIGatewayMEVMErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMErrors{IGatewayMEVMErrorsCaller: IGatewayMEVMErrorsCaller{contract: contract}, IGatewayMEVMErrorsTransactor: IGatewayMEVMErrorsTransactor{contract: contract}, IGatewayMEVMErrorsFilterer: IGatewayMEVMErrorsFilterer{contract: contract}}, nil
}

// NewIGatewayMEVMErrorsCaller creates a new read-only instance of IGatewayMEVMErrors, bound to a specific deployed contract.
func NewIGatewayMEVMErrorsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayMEVMErrorsCaller, error) {
	contract, err := bindIGatewayMEVMErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMErrorsCaller{contract: contract}, nil
}

// NewIGatewayMEVMErrorsTransactor creates a new write-only instance of IGatewayMEVMErrors, bound to a specific deployed contract.
func NewIGatewayMEVMErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayMEVMErrorsTransactor, error) {
	contract, err := bindIGatewayMEVMErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMErrorsTransactor{contract: contract}, nil
}

// NewIGatewayMEVMErrorsFilterer creates a new log filterer instance of IGatewayMEVMErrors, bound to a specific deployed contract.
func NewIGatewayMEVMErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayMEVMErrorsFilterer, error) {
	contract, err := bindIGatewayMEVMErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMErrorsFilterer{contract: contract}, nil
}

// bindIGatewayMEVMErrors binds a generic wrapper to an already deployed contract.
func bindIGatewayMEVMErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayMEVMErrorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVMErrors.Contract.IGatewayMEVMErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVMErrors.Contract.IGatewayMEVMErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVMErrors.Contract.IGatewayMEVMErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVMErrors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVMErrors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVMErrors *IGatewayMEVMErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVMErrors.Contract.contract.Transact(opts, method, params...)
}
