// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imrc20

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

// IMRC20MetadataMetaData contains all meta data concerning the IMRC20Metadata contract.
var IMRC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_FLAT_FEE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"newName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"newSymbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawGasFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawGasFeeWithGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// IMRC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IMRC20MetadataMetaData.ABI instead.
var IMRC20MetadataABI = IMRC20MetadataMetaData.ABI

// IMRC20Metadata is an auto generated Go binding around an Ethereum contract.
type IMRC20Metadata struct {
	IMRC20MetadataCaller     // Read-only binding to the contract
	IMRC20MetadataTransactor // Write-only binding to the contract
	IMRC20MetadataFilterer   // Log filterer for contract events
}

// IMRC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMRC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMRC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMRC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMRC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMRC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMRC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMRC20MetadataSession struct {
	Contract     *IMRC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMRC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMRC20MetadataCallerSession struct {
	Contract *IMRC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IMRC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMRC20MetadataTransactorSession struct {
	Contract     *IMRC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IMRC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMRC20MetadataRaw struct {
	Contract *IMRC20Metadata // Generic contract binding to access the raw methods on
}

// IMRC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMRC20MetadataCallerRaw struct {
	Contract *IMRC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IMRC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMRC20MetadataTransactorRaw struct {
	Contract *IMRC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMRC20Metadata creates a new instance of IMRC20Metadata, bound to a specific deployed contract.
func NewIMRC20Metadata(address common.Address, backend bind.ContractBackend) (*IMRC20Metadata, error) {
	contract, err := bindIMRC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMRC20Metadata{IMRC20MetadataCaller: IMRC20MetadataCaller{contract: contract}, IMRC20MetadataTransactor: IMRC20MetadataTransactor{contract: contract}, IMRC20MetadataFilterer: IMRC20MetadataFilterer{contract: contract}}, nil
}

// NewIMRC20MetadataCaller creates a new read-only instance of IMRC20Metadata, bound to a specific deployed contract.
func NewIMRC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IMRC20MetadataCaller, error) {
	contract, err := bindIMRC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMRC20MetadataCaller{contract: contract}, nil
}

// NewIMRC20MetadataTransactor creates a new write-only instance of IMRC20Metadata, bound to a specific deployed contract.
func NewIMRC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IMRC20MetadataTransactor, error) {
	contract, err := bindIMRC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMRC20MetadataTransactor{contract: contract}, nil
}

// NewIMRC20MetadataFilterer creates a new log filterer instance of IMRC20Metadata, bound to a specific deployed contract.
func NewIMRC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IMRC20MetadataFilterer, error) {
	contract, err := bindIMRC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMRC20MetadataFilterer{contract: contract}, nil
}

// bindIMRC20Metadata binds a generic wrapper to an already deployed contract.
func bindIMRC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMRC20MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMRC20Metadata *IMRC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMRC20Metadata.Contract.IMRC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMRC20Metadata *IMRC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.IMRC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMRC20Metadata *IMRC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.IMRC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMRC20Metadata *IMRC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMRC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMRC20Metadata *IMRC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMRC20Metadata *IMRC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) GASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) GASLIMIT() (*big.Int, error) {
	return _IMRC20Metadata.Contract.GASLIMIT(&_IMRC20Metadata.CallOpts)
}

// GASLIMIT is a free data retrieval call binding the contract method 0x091d2788.
//
// Solidity: function GAS_LIMIT() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) GASLIMIT() (*big.Int, error) {
	return _IMRC20Metadata.Contract.GASLIMIT(&_IMRC20Metadata.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) PROTOCOLFLATFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "PROTOCOL_FLAT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IMRC20Metadata.Contract.PROTOCOLFLATFEE(&_IMRC20Metadata.CallOpts)
}

// PROTOCOLFLATFEE is a free data retrieval call binding the contract method 0x4d8943bb.
//
// Solidity: function PROTOCOL_FLAT_FEE() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) PROTOCOLFLATFEE() (*big.Int, error) {
	return _IMRC20Metadata.Contract.PROTOCOLFLATFEE(&_IMRC20Metadata.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMRC20Metadata.Contract.Allowance(&_IMRC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMRC20Metadata.Contract.Allowance(&_IMRC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMRC20Metadata.Contract.BalanceOf(&_IMRC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMRC20Metadata.Contract.BalanceOf(&_IMRC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IMRC20Metadata *IMRC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IMRC20Metadata *IMRC20MetadataSession) Decimals() (uint8, error) {
	return _IMRC20Metadata.Contract.Decimals(&_IMRC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IMRC20Metadata.Contract.Decimals(&_IMRC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataSession) Name() (string, error) {
	return _IMRC20Metadata.Contract.Name(&_IMRC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) Name() (string, error) {
	return _IMRC20Metadata.Contract.Name(&_IMRC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataSession) Symbol() (string, error) {
	return _IMRC20Metadata.Contract.Symbol(&_IMRC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) Symbol() (string, error) {
	return _IMRC20Metadata.Contract.Symbol(&_IMRC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IMRC20Metadata.Contract.TotalSupply(&_IMRC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IMRC20Metadata.Contract.TotalSupply(&_IMRC20Metadata.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) WithdrawGasFee(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "withdrawGasFee")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IMRC20Metadata.Contract.WithdrawGasFee(&_IMRC20Metadata.CallOpts)
}

// WithdrawGasFee is a free data retrieval call binding the contract method 0xd9eeebed.
//
// Solidity: function withdrawGasFee() view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) WithdrawGasFee() (common.Address, *big.Int, error) {
	return _IMRC20Metadata.Contract.WithdrawGasFee(&_IMRC20Metadata.CallOpts)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataCaller) WithdrawGasFeeWithGasLimit(opts *bind.CallOpts, gasLimit *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _IMRC20Metadata.contract.Call(opts, &out, "withdrawGasFeeWithGasLimit", gasLimit)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IMRC20Metadata.Contract.WithdrawGasFeeWithGasLimit(&_IMRC20Metadata.CallOpts, gasLimit)
}

// WithdrawGasFeeWithGasLimit is a free data retrieval call binding the contract method 0xfc5fecd5.
//
// Solidity: function withdrawGasFeeWithGasLimit(uint256 gasLimit) view returns(address, uint256)
func (_IMRC20Metadata *IMRC20MetadataCallerSession) WithdrawGasFeeWithGasLimit(gasLimit *big.Int) (common.Address, *big.Int, error) {
	return _IMRC20Metadata.Contract.WithdrawGasFeeWithGasLimit(&_IMRC20Metadata.CallOpts, gasLimit)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Approve(&_IMRC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Approve(&_IMRC20Metadata.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Burn(&_IMRC20Metadata.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Burn(&_IMRC20Metadata.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) Deposit(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "deposit", to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Deposit(&_IMRC20Metadata.TransactOpts, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) Deposit(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Deposit(&_IMRC20Metadata.TransactOpts, to, amount)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IMRC20Metadata *IMRC20MetadataTransactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IMRC20Metadata *IMRC20MetadataSession) SetName(newName string) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.SetName(&_IMRC20Metadata.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.SetName(&_IMRC20Metadata.TransactOpts, newName)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IMRC20Metadata *IMRC20MetadataTransactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IMRC20Metadata *IMRC20MetadataSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.SetSymbol(&_IMRC20Metadata.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.SetSymbol(&_IMRC20Metadata.TransactOpts, newSymbol)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Transfer(&_IMRC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Transfer(&_IMRC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.TransferFrom(&_IMRC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.TransferFrom(&_IMRC20Metadata.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactor) Withdraw(opts *bind.TransactOpts, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Withdraw(&_IMRC20Metadata.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc7012626.
//
// Solidity: function withdraw(bytes to, uint256 amount) returns(bool)
func (_IMRC20Metadata *IMRC20MetadataTransactorSession) Withdraw(to []byte, amount *big.Int) (*types.Transaction, error) {
	return _IMRC20Metadata.Contract.Withdraw(&_IMRC20Metadata.TransactOpts, to, amount)
}
