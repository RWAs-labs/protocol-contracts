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

// CallOptions is an auto generated low-level Go binding around an user-defined struct.
type CallOptions struct {
	GasLimit        *big.Int
	IsArbitraryCall bool
}

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender    []byte
	SenderEVM common.Address
	ChainID   *big.Int
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// IGatewayMEVMMetaData contains all meta data concerning the IGatewayMEVM contract.
var IGatewayMEVMMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedMuseSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMuseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MRC20BurnFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20DepositFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20TransferFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWMUSEOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// IGatewayMEVMABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayMEVMMetaData.ABI instead.
var IGatewayMEVMABI = IGatewayMEVMMetaData.ABI

// IGatewayMEVM is an auto generated Go binding around an Ethereum contract.
type IGatewayMEVM struct {
	IGatewayMEVMCaller     // Read-only binding to the contract
	IGatewayMEVMTransactor // Write-only binding to the contract
	IGatewayMEVMFilterer   // Log filterer for contract events
}

// IGatewayMEVMCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayMEVMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayMEVMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayMEVMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayMEVMSession struct {
	Contract     *IGatewayMEVM     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGatewayMEVMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayMEVMCallerSession struct {
	Contract *IGatewayMEVMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IGatewayMEVMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayMEVMTransactorSession struct {
	Contract     *IGatewayMEVMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IGatewayMEVMRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayMEVMRaw struct {
	Contract *IGatewayMEVM // Generic contract binding to access the raw methods on
}

// IGatewayMEVMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayMEVMCallerRaw struct {
	Contract *IGatewayMEVMCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayMEVMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayMEVMTransactorRaw struct {
	Contract *IGatewayMEVMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayMEVM creates a new instance of IGatewayMEVM, bound to a specific deployed contract.
func NewIGatewayMEVM(address common.Address, backend bind.ContractBackend) (*IGatewayMEVM, error) {
	contract, err := bindIGatewayMEVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVM{IGatewayMEVMCaller: IGatewayMEVMCaller{contract: contract}, IGatewayMEVMTransactor: IGatewayMEVMTransactor{contract: contract}, IGatewayMEVMFilterer: IGatewayMEVMFilterer{contract: contract}}, nil
}

// NewIGatewayMEVMCaller creates a new read-only instance of IGatewayMEVM, bound to a specific deployed contract.
func NewIGatewayMEVMCaller(address common.Address, caller bind.ContractCaller) (*IGatewayMEVMCaller, error) {
	contract, err := bindIGatewayMEVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMCaller{contract: contract}, nil
}

// NewIGatewayMEVMTransactor creates a new write-only instance of IGatewayMEVM, bound to a specific deployed contract.
func NewIGatewayMEVMTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayMEVMTransactor, error) {
	contract, err := bindIGatewayMEVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMTransactor{contract: contract}, nil
}

// NewIGatewayMEVMFilterer creates a new log filterer instance of IGatewayMEVM, bound to a specific deployed contract.
func NewIGatewayMEVMFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayMEVMFilterer, error) {
	contract, err := bindIGatewayMEVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMFilterer{contract: contract}, nil
}

// bindIGatewayMEVM binds a generic wrapper to an already deployed contract.
func bindIGatewayMEVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayMEVMMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVM *IGatewayMEVMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVM.Contract.IGatewayMEVMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVM *IGatewayMEVMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.IGatewayMEVMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVM *IGatewayMEVMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.IGatewayMEVMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVM *IGatewayMEVMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVM *IGatewayMEVMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVM *IGatewayMEVMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) Call(opts *bind.TransactOpts, receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "call", receiver, mrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) Call(receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Call(&_IGatewayMEVM.TransactOpts, receiver, mrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) Call(receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Call(&_IGatewayMEVM.TransactOpts, receiver, mrc20, message, callOptions, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) Deposit(opts *bind.TransactOpts, mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "deposit", mrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) Deposit(mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Deposit(&_IGatewayMEVM.TransactOpts, mrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) Deposit(mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Deposit(&_IGatewayMEVM.TransactOpts, mrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndCall(&_IGatewayMEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndCall(&_IGatewayMEVM.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "depositAndCall0", context, mrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) DepositAndCall0(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndCall0(&_IGatewayMEVM.TransactOpts, context, mrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) DepositAndCall0(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndCall0(&_IGatewayMEVM.TransactOpts, context, mrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) DepositAndRevert(opts *bind.TransactOpts, mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "depositAndRevert", mrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) DepositAndRevert(mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndRevert(&_IGatewayMEVM.TransactOpts, mrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) DepositAndRevert(mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.DepositAndRevert(&_IGatewayMEVM.TransactOpts, mrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) Execute(opts *bind.TransactOpts, context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "execute", context, mrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) Execute(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Execute(&_IGatewayMEVM.TransactOpts, context, mrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) Execute(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Execute(&_IGatewayMEVM.TransactOpts, context, mrc20, amount, target, message)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.ExecuteRevert(&_IGatewayMEVM.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.ExecuteRevert(&_IGatewayMEVM.TransactOpts, target, revertContext)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "withdraw", receiver, amount, mrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) Withdraw(receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Withdraw(&_IGatewayMEVM.TransactOpts, receiver, amount, mrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) Withdraw(receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Withdraw(&_IGatewayMEVM.TransactOpts, receiver, amount, mrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Withdraw0(&_IGatewayMEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.Withdraw0(&_IGatewayMEVM.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "withdrawAndCall", receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.WithdrawAndCall(&_IGatewayMEVM.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.WithdrawAndCall(&_IGatewayMEVM.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.contract.Transact(opts, "withdrawAndCall0", receiver, amount, mrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMSession) WithdrawAndCall0(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.WithdrawAndCall0(&_IGatewayMEVM.TransactOpts, receiver, amount, mrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_IGatewayMEVM *IGatewayMEVMTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _IGatewayMEVM.Contract.WithdrawAndCall0(&_IGatewayMEVM.TransactOpts, receiver, amount, mrc20, message, callOptions, revertOptions)
}

// IGatewayMEVMCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the IGatewayMEVM contract.
type IGatewayMEVMCalledIterator struct {
	Event *IGatewayMEVMCalled // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMCalled)
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
		it.Event = new(IGatewayMEVMCalled)
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
func (it *IGatewayMEVMCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMCalled represents a Called event raised by the IGatewayMEVM contract.
type IGatewayMEVMCalled struct {
	Sender        common.Address
	Mrc20         common.Address
	Receiver      []byte
	Message       []byte
	CallOptions   CallOptions
	RevertOptions RevertOptions
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCalled is a free log retrieval operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, mrc20 []common.Address) (*IGatewayMEVMCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _IGatewayMEVM.contract.FilterLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMCalledIterator{contract: _IGatewayMEVM.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMCalled, sender []common.Address, mrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _IGatewayMEVM.contract.WatchLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMCalled)
				if err := _IGatewayMEVM.contract.UnpackLog(event, "Called", log); err != nil {
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

// ParseCalled is a log parse operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) ParseCalled(log types.Log) (*IGatewayMEVMCalled, error) {
	event := new(IGatewayMEVMCalled)
	if err := _IGatewayMEVM.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayMEVMWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IGatewayMEVM contract.
type IGatewayMEVMWithdrawnIterator struct {
	Event *IGatewayMEVMWithdrawn // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMWithdrawn)
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
		it.Event = new(IGatewayMEVMWithdrawn)
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
func (it *IGatewayMEVMWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMWithdrawn represents a Withdrawn event raised by the IGatewayMEVM contract.
type IGatewayMEVMWithdrawn struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Mrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*IGatewayMEVMWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVM.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMWithdrawnIterator{contract: _IGatewayMEVM.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVM.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMWithdrawn)
				if err := _IGatewayMEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) ParseWithdrawn(log types.Log) (*IGatewayMEVMWithdrawn, error) {
	event := new(IGatewayMEVMWithdrawn)
	if err := _IGatewayMEVM.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayMEVMWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IGatewayMEVM contract.
type IGatewayMEVMWithdrawnAndCalledIterator struct {
	Event *IGatewayMEVMWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMWithdrawnAndCalled)
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
		it.Event = new(IGatewayMEVMWithdrawnAndCalled)
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
func (it *IGatewayMEVMWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IGatewayMEVM contract.
type IGatewayMEVMWithdrawnAndCalled struct {
	Sender          common.Address
	ChainId         *big.Int
	Receiver        []byte
	Mrc20           common.Address
	Value           *big.Int
	Gasfee          *big.Int
	ProtocolFlatFee *big.Int
	Message         []byte
	CallOptions     CallOptions
	RevertOptions   RevertOptions
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*IGatewayMEVMWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVM.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMWithdrawnAndCalledIterator{contract: _IGatewayMEVM.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVM.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMWithdrawnAndCalled)
				if err := _IGatewayMEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVM *IGatewayMEVMFilterer) ParseWithdrawnAndCalled(log types.Log) (*IGatewayMEVMWithdrawnAndCalled, error) {
	event := new(IGatewayMEVMWithdrawnAndCalled)
	if err := _IGatewayMEVM.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
