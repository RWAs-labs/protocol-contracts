// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imusenonethnew

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

// IMuseNonEthNewMetaData contains all meta data concerning the IMuseNonEthNew contract.
var IMuseNonEthNewMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// IMuseNonEthNewABI is the input ABI used to generate the binding from.
// Deprecated: Use IMuseNonEthNewMetaData.ABI instead.
var IMuseNonEthNewABI = IMuseNonEthNewMetaData.ABI

// IMuseNonEthNew is an auto generated Go binding around an Ethereum contract.
type IMuseNonEthNew struct {
	IMuseNonEthNewCaller     // Read-only binding to the contract
	IMuseNonEthNewTransactor // Write-only binding to the contract
	IMuseNonEthNewFilterer   // Log filterer for contract events
}

// IMuseNonEthNewCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMuseNonEthNewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseNonEthNewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMuseNonEthNewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseNonEthNewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMuseNonEthNewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseNonEthNewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMuseNonEthNewSession struct {
	Contract     *IMuseNonEthNew   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMuseNonEthNewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMuseNonEthNewCallerSession struct {
	Contract *IMuseNonEthNewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IMuseNonEthNewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMuseNonEthNewTransactorSession struct {
	Contract     *IMuseNonEthNewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IMuseNonEthNewRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMuseNonEthNewRaw struct {
	Contract *IMuseNonEthNew // Generic contract binding to access the raw methods on
}

// IMuseNonEthNewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMuseNonEthNewCallerRaw struct {
	Contract *IMuseNonEthNewCaller // Generic read-only contract binding to access the raw methods on
}

// IMuseNonEthNewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMuseNonEthNewTransactorRaw struct {
	Contract *IMuseNonEthNewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMuseNonEthNew creates a new instance of IMuseNonEthNew, bound to a specific deployed contract.
func NewIMuseNonEthNew(address common.Address, backend bind.ContractBackend) (*IMuseNonEthNew, error) {
	contract, err := bindIMuseNonEthNew(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNew{IMuseNonEthNewCaller: IMuseNonEthNewCaller{contract: contract}, IMuseNonEthNewTransactor: IMuseNonEthNewTransactor{contract: contract}, IMuseNonEthNewFilterer: IMuseNonEthNewFilterer{contract: contract}}, nil
}

// NewIMuseNonEthNewCaller creates a new read-only instance of IMuseNonEthNew, bound to a specific deployed contract.
func NewIMuseNonEthNewCaller(address common.Address, caller bind.ContractCaller) (*IMuseNonEthNewCaller, error) {
	contract, err := bindIMuseNonEthNew(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNewCaller{contract: contract}, nil
}

// NewIMuseNonEthNewTransactor creates a new write-only instance of IMuseNonEthNew, bound to a specific deployed contract.
func NewIMuseNonEthNewTransactor(address common.Address, transactor bind.ContractTransactor) (*IMuseNonEthNewTransactor, error) {
	contract, err := bindIMuseNonEthNew(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNewTransactor{contract: contract}, nil
}

// NewIMuseNonEthNewFilterer creates a new log filterer instance of IMuseNonEthNew, bound to a specific deployed contract.
func NewIMuseNonEthNewFilterer(address common.Address, filterer bind.ContractFilterer) (*IMuseNonEthNewFilterer, error) {
	contract, err := bindIMuseNonEthNew(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNewFilterer{contract: contract}, nil
}

// bindIMuseNonEthNew binds a generic wrapper to an already deployed contract.
func bindIMuseNonEthNew(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMuseNonEthNewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuseNonEthNew *IMuseNonEthNewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuseNonEthNew.Contract.IMuseNonEthNewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuseNonEthNew *IMuseNonEthNewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.IMuseNonEthNewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuseNonEthNew *IMuseNonEthNewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.IMuseNonEthNewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuseNonEthNew *IMuseNonEthNewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuseNonEthNew.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuseNonEthNew *IMuseNonEthNewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuseNonEthNew *IMuseNonEthNewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMuseNonEthNew.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMuseNonEthNew.Contract.Allowance(&_IMuseNonEthNew.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMuseNonEthNew.Contract.Allowance(&_IMuseNonEthNew.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMuseNonEthNew.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMuseNonEthNew.Contract.BalanceOf(&_IMuseNonEthNew.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMuseNonEthNew.Contract.BalanceOf(&_IMuseNonEthNew.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMuseNonEthNew.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewSession) TotalSupply() (*big.Int, error) {
	return _IMuseNonEthNew.Contract.TotalSupply(&_IMuseNonEthNew.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMuseNonEthNew *IMuseNonEthNewCallerSession) TotalSupply() (*big.Int, error) {
	return _IMuseNonEthNew.Contract.TotalSupply(&_IMuseNonEthNew.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Approve(&_IMuseNonEthNew.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Approve(&_IMuseNonEthNew.TransactOpts, spender, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IMuseNonEthNew *IMuseNonEthNewTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IMuseNonEthNew *IMuseNonEthNewSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.BurnFrom(&_IMuseNonEthNew.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_IMuseNonEthNew *IMuseNonEthNewTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.BurnFrom(&_IMuseNonEthNew.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IMuseNonEthNew *IMuseNonEthNewTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IMuseNonEthNew.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IMuseNonEthNew *IMuseNonEthNewSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Mint(&_IMuseNonEthNew.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_IMuseNonEthNew *IMuseNonEthNewTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Mint(&_IMuseNonEthNew.TransactOpts, mintee, value, internalSendHash)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Transfer(&_IMuseNonEthNew.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.Transfer(&_IMuseNonEthNew.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.TransferFrom(&_IMuseNonEthNew.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IMuseNonEthNew *IMuseNonEthNewTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IMuseNonEthNew.Contract.TransferFrom(&_IMuseNonEthNew.TransactOpts, from, to, value)
}

// IMuseNonEthNewApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IMuseNonEthNew contract.
type IMuseNonEthNewApprovalIterator struct {
	Event *IMuseNonEthNewApproval // Event containing the contract specifics and raw log

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
func (it *IMuseNonEthNewApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseNonEthNewApproval)
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
		it.Event = new(IMuseNonEthNewApproval)
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
func (it *IMuseNonEthNewApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseNonEthNewApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseNonEthNewApproval represents a Approval event raised by the IMuseNonEthNew contract.
type IMuseNonEthNewApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IMuseNonEthNewApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMuseNonEthNew.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNewApprovalIterator{contract: _IMuseNonEthNew.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IMuseNonEthNewApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMuseNonEthNew.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseNonEthNewApproval)
				if err := _IMuseNonEthNew.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) ParseApproval(log types.Log) (*IMuseNonEthNewApproval, error) {
	event := new(IMuseNonEthNewApproval)
	if err := _IMuseNonEthNew.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMuseNonEthNewTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IMuseNonEthNew contract.
type IMuseNonEthNewTransferIterator struct {
	Event *IMuseNonEthNewTransfer // Event containing the contract specifics and raw log

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
func (it *IMuseNonEthNewTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseNonEthNewTransfer)
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
		it.Event = new(IMuseNonEthNewTransfer)
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
func (it *IMuseNonEthNewTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseNonEthNewTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseNonEthNewTransfer represents a Transfer event raised by the IMuseNonEthNew contract.
type IMuseNonEthNewTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IMuseNonEthNewTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseNonEthNew.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IMuseNonEthNewTransferIterator{contract: _IMuseNonEthNew.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IMuseNonEthNewTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseNonEthNew.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseNonEthNewTransfer)
				if err := _IMuseNonEthNew.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMuseNonEthNew *IMuseNonEthNewFilterer) ParseTransfer(log types.Log) (*IMuseNonEthNewTransfer, error) {
	event := new(IMuseNonEthNewTransfer)
	if err := _IMuseNonEthNew.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
