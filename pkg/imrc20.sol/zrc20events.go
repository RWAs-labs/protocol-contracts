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

// MRC20EventsMetaData contains all meta data concerning the MRC20Events contract.
var MRC20EventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedGateway\",\"inputs\":[{\"name\":\"gateway\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedProtocolFlatFee\",\"inputs\":[{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedSystemContract\",\"inputs\":[{\"name\":\"systemContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MRC20EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use MRC20EventsMetaData.ABI instead.
var MRC20EventsABI = MRC20EventsMetaData.ABI

// MRC20Events is an auto generated Go binding around an Ethereum contract.
type MRC20Events struct {
	MRC20EventsCaller     // Read-only binding to the contract
	MRC20EventsTransactor // Write-only binding to the contract
	MRC20EventsFilterer   // Log filterer for contract events
}

// MRC20EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MRC20EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MRC20EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MRC20EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MRC20EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MRC20EventsSession struct {
	Contract     *MRC20Events      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MRC20EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MRC20EventsCallerSession struct {
	Contract *MRC20EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MRC20EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MRC20EventsTransactorSession struct {
	Contract     *MRC20EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MRC20EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MRC20EventsRaw struct {
	Contract *MRC20Events // Generic contract binding to access the raw methods on
}

// MRC20EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MRC20EventsCallerRaw struct {
	Contract *MRC20EventsCaller // Generic read-only contract binding to access the raw methods on
}

// MRC20EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MRC20EventsTransactorRaw struct {
	Contract *MRC20EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMRC20Events creates a new instance of MRC20Events, bound to a specific deployed contract.
func NewMRC20Events(address common.Address, backend bind.ContractBackend) (*MRC20Events, error) {
	contract, err := bindMRC20Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MRC20Events{MRC20EventsCaller: MRC20EventsCaller{contract: contract}, MRC20EventsTransactor: MRC20EventsTransactor{contract: contract}, MRC20EventsFilterer: MRC20EventsFilterer{contract: contract}}, nil
}

// NewMRC20EventsCaller creates a new read-only instance of MRC20Events, bound to a specific deployed contract.
func NewMRC20EventsCaller(address common.Address, caller bind.ContractCaller) (*MRC20EventsCaller, error) {
	contract, err := bindMRC20Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsCaller{contract: contract}, nil
}

// NewMRC20EventsTransactor creates a new write-only instance of MRC20Events, bound to a specific deployed contract.
func NewMRC20EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*MRC20EventsTransactor, error) {
	contract, err := bindMRC20Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsTransactor{contract: contract}, nil
}

// NewMRC20EventsFilterer creates a new log filterer instance of MRC20Events, bound to a specific deployed contract.
func NewMRC20EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*MRC20EventsFilterer, error) {
	contract, err := bindMRC20Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsFilterer{contract: contract}, nil
}

// bindMRC20Events binds a generic wrapper to an already deployed contract.
func bindMRC20Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MRC20EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MRC20Events *MRC20EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MRC20Events.Contract.MRC20EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MRC20Events *MRC20EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MRC20Events.Contract.MRC20EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MRC20Events *MRC20EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MRC20Events.Contract.MRC20EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MRC20Events *MRC20EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MRC20Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MRC20Events *MRC20EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MRC20Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MRC20Events *MRC20EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MRC20Events.Contract.contract.Transact(opts, method, params...)
}

// MRC20EventsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MRC20Events contract.
type MRC20EventsApprovalIterator struct {
	Event *MRC20EventsApproval // Event containing the contract specifics and raw log

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
func (it *MRC20EventsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsApproval)
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
		it.Event = new(MRC20EventsApproval)
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
func (it *MRC20EventsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsApproval represents a Approval event raised by the MRC20Events contract.
type MRC20EventsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MRC20EventsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsApprovalIterator{contract: _MRC20Events.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MRC20EventsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsApproval)
				if err := _MRC20Events.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MRC20Events *MRC20EventsFilterer) ParseApproval(log types.Log) (*MRC20EventsApproval, error) {
	event := new(MRC20EventsApproval)
	if err := _MRC20Events.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MRC20Events contract.
type MRC20EventsDepositIterator struct {
	Event *MRC20EventsDeposit // Event containing the contract specifics and raw log

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
func (it *MRC20EventsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsDeposit)
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
		it.Event = new(MRC20EventsDeposit)
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
func (it *MRC20EventsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsDeposit represents a Deposit event raised by the MRC20Events contract.
type MRC20EventsDeposit struct {
	From  []byte
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) FilterDeposit(opts *bind.FilterOpts, to []common.Address) (*MRC20EventsDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsDepositIterator{contract: _MRC20Events.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MRC20EventsDeposit, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "Deposit", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsDeposit)
				if err := _MRC20Events.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab3.
//
// Solidity: event Deposit(bytes from, address indexed to, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) ParseDeposit(log types.Log) (*MRC20EventsDeposit, error) {
	event := new(MRC20EventsDeposit)
	if err := _MRC20Events.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MRC20Events contract.
type MRC20EventsTransferIterator struct {
	Event *MRC20EventsTransfer // Event containing the contract specifics and raw log

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
func (it *MRC20EventsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsTransfer)
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
		it.Event = new(MRC20EventsTransfer)
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
func (it *MRC20EventsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsTransfer represents a Transfer event raised by the MRC20Events contract.
type MRC20EventsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MRC20EventsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsTransferIterator{contract: _MRC20Events.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MRC20Events *MRC20EventsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MRC20EventsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsTransfer)
				if err := _MRC20Events.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MRC20Events *MRC20EventsFilterer) ParseTransfer(log types.Log) (*MRC20EventsTransfer, error) {
	event := new(MRC20EventsTransfer)
	if err := _MRC20Events.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsUpdatedGasLimitIterator is returned from FilterUpdatedGasLimit and is used to iterate over the raw logs and unpacked data for UpdatedGasLimit events raised by the MRC20Events contract.
type MRC20EventsUpdatedGasLimitIterator struct {
	Event *MRC20EventsUpdatedGasLimit // Event containing the contract specifics and raw log

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
func (it *MRC20EventsUpdatedGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsUpdatedGasLimit)
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
		it.Event = new(MRC20EventsUpdatedGasLimit)
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
func (it *MRC20EventsUpdatedGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsUpdatedGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsUpdatedGasLimit represents a UpdatedGasLimit event raised by the MRC20Events contract.
type MRC20EventsUpdatedGasLimit struct {
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGasLimit is a free log retrieval operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_MRC20Events *MRC20EventsFilterer) FilterUpdatedGasLimit(opts *bind.FilterOpts) (*MRC20EventsUpdatedGasLimitIterator, error) {

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return &MRC20EventsUpdatedGasLimitIterator{contract: _MRC20Events.contract, event: "UpdatedGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdatedGasLimit is a free log subscription operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_MRC20Events *MRC20EventsFilterer) WatchUpdatedGasLimit(opts *bind.WatchOpts, sink chan<- *MRC20EventsUpdatedGasLimit) (event.Subscription, error) {

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "UpdatedGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsUpdatedGasLimit)
				if err := _MRC20Events.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
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

// ParseUpdatedGasLimit is a log parse operation binding the contract event 0xff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a.
//
// Solidity: event UpdatedGasLimit(uint256 gasLimit)
func (_MRC20Events *MRC20EventsFilterer) ParseUpdatedGasLimit(log types.Log) (*MRC20EventsUpdatedGasLimit, error) {
	event := new(MRC20EventsUpdatedGasLimit)
	if err := _MRC20Events.contract.UnpackLog(event, "UpdatedGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsUpdatedGatewayIterator is returned from FilterUpdatedGateway and is used to iterate over the raw logs and unpacked data for UpdatedGateway events raised by the MRC20Events contract.
type MRC20EventsUpdatedGatewayIterator struct {
	Event *MRC20EventsUpdatedGateway // Event containing the contract specifics and raw log

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
func (it *MRC20EventsUpdatedGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsUpdatedGateway)
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
		it.Event = new(MRC20EventsUpdatedGateway)
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
func (it *MRC20EventsUpdatedGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsUpdatedGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsUpdatedGateway represents a UpdatedGateway event raised by the MRC20Events contract.
type MRC20EventsUpdatedGateway struct {
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatedGateway is a free log retrieval operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_MRC20Events *MRC20EventsFilterer) FilterUpdatedGateway(opts *bind.FilterOpts) (*MRC20EventsUpdatedGatewayIterator, error) {

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return &MRC20EventsUpdatedGatewayIterator{contract: _MRC20Events.contract, event: "UpdatedGateway", logs: logs, sub: sub}, nil
}

// WatchUpdatedGateway is a free log subscription operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_MRC20Events *MRC20EventsFilterer) WatchUpdatedGateway(opts *bind.WatchOpts, sink chan<- *MRC20EventsUpdatedGateway) (event.Subscription, error) {

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "UpdatedGateway")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsUpdatedGateway)
				if err := _MRC20Events.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
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

// ParseUpdatedGateway is a log parse operation binding the contract event 0x88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a387.
//
// Solidity: event UpdatedGateway(address gateway)
func (_MRC20Events *MRC20EventsFilterer) ParseUpdatedGateway(log types.Log) (*MRC20EventsUpdatedGateway, error) {
	event := new(MRC20EventsUpdatedGateway)
	if err := _MRC20Events.contract.UnpackLog(event, "UpdatedGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsUpdatedProtocolFlatFeeIterator is returned from FilterUpdatedProtocolFlatFee and is used to iterate over the raw logs and unpacked data for UpdatedProtocolFlatFee events raised by the MRC20Events contract.
type MRC20EventsUpdatedProtocolFlatFeeIterator struct {
	Event *MRC20EventsUpdatedProtocolFlatFee // Event containing the contract specifics and raw log

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
func (it *MRC20EventsUpdatedProtocolFlatFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsUpdatedProtocolFlatFee)
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
		it.Event = new(MRC20EventsUpdatedProtocolFlatFee)
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
func (it *MRC20EventsUpdatedProtocolFlatFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsUpdatedProtocolFlatFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsUpdatedProtocolFlatFee represents a UpdatedProtocolFlatFee event raised by the MRC20Events contract.
type MRC20EventsUpdatedProtocolFlatFee struct {
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedProtocolFlatFee is a free log retrieval operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) FilterUpdatedProtocolFlatFee(opts *bind.FilterOpts) (*MRC20EventsUpdatedProtocolFlatFeeIterator, error) {

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return &MRC20EventsUpdatedProtocolFlatFeeIterator{contract: _MRC20Events.contract, event: "UpdatedProtocolFlatFee", logs: logs, sub: sub}, nil
}

// WatchUpdatedProtocolFlatFee is a free log subscription operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) WatchUpdatedProtocolFlatFee(opts *bind.WatchOpts, sink chan<- *MRC20EventsUpdatedProtocolFlatFee) (event.Subscription, error) {

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "UpdatedProtocolFlatFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsUpdatedProtocolFlatFee)
				if err := _MRC20Events.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
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

// ParseUpdatedProtocolFlatFee is a log parse operation binding the contract event 0xef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f.
//
// Solidity: event UpdatedProtocolFlatFee(uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) ParseUpdatedProtocolFlatFee(log types.Log) (*MRC20EventsUpdatedProtocolFlatFee, error) {
	event := new(MRC20EventsUpdatedProtocolFlatFee)
	if err := _MRC20Events.contract.UnpackLog(event, "UpdatedProtocolFlatFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsUpdatedSystemContractIterator is returned from FilterUpdatedSystemContract and is used to iterate over the raw logs and unpacked data for UpdatedSystemContract events raised by the MRC20Events contract.
type MRC20EventsUpdatedSystemContractIterator struct {
	Event *MRC20EventsUpdatedSystemContract // Event containing the contract specifics and raw log

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
func (it *MRC20EventsUpdatedSystemContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsUpdatedSystemContract)
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
		it.Event = new(MRC20EventsUpdatedSystemContract)
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
func (it *MRC20EventsUpdatedSystemContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsUpdatedSystemContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsUpdatedSystemContract represents a UpdatedSystemContract event raised by the MRC20Events contract.
type MRC20EventsUpdatedSystemContract struct {
	SystemContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdatedSystemContract is a free log retrieval operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_MRC20Events *MRC20EventsFilterer) FilterUpdatedSystemContract(opts *bind.FilterOpts) (*MRC20EventsUpdatedSystemContractIterator, error) {

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return &MRC20EventsUpdatedSystemContractIterator{contract: _MRC20Events.contract, event: "UpdatedSystemContract", logs: logs, sub: sub}, nil
}

// WatchUpdatedSystemContract is a free log subscription operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_MRC20Events *MRC20EventsFilterer) WatchUpdatedSystemContract(opts *bind.WatchOpts, sink chan<- *MRC20EventsUpdatedSystemContract) (event.Subscription, error) {

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "UpdatedSystemContract")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsUpdatedSystemContract)
				if err := _MRC20Events.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
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

// ParseUpdatedSystemContract is a log parse operation binding the contract event 0xd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae.
//
// Solidity: event UpdatedSystemContract(address systemContract)
func (_MRC20Events *MRC20EventsFilterer) ParseUpdatedSystemContract(log types.Log) (*MRC20EventsUpdatedSystemContract, error) {
	event := new(MRC20EventsUpdatedSystemContract)
	if err := _MRC20Events.contract.UnpackLog(event, "UpdatedSystemContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MRC20EventsWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the MRC20Events contract.
type MRC20EventsWithdrawalIterator struct {
	Event *MRC20EventsWithdrawal // Event containing the contract specifics and raw log

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
func (it *MRC20EventsWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MRC20EventsWithdrawal)
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
		it.Event = new(MRC20EventsWithdrawal)
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
func (it *MRC20EventsWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MRC20EventsWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MRC20EventsWithdrawal represents a Withdrawal event raised by the MRC20Events contract.
type MRC20EventsWithdrawal struct {
	From            common.Address
	To              []byte
	Value           *big.Int
	GasFee          *big.Int
	ProtocolFlatFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) FilterWithdrawal(opts *bind.FilterOpts, from []common.Address) (*MRC20EventsWithdrawalIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20Events.contract.FilterLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return &MRC20EventsWithdrawalIterator{contract: _MRC20Events.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *MRC20EventsWithdrawal, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MRC20Events.contract.WatchLogs(opts, "Withdrawal", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MRC20EventsWithdrawal)
				if err := _MRC20Events.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d955.
//
// Solidity: event Withdrawal(address indexed from, bytes to, uint256 value, uint256 gasFee, uint256 protocolFlatFee)
func (_MRC20Events *MRC20EventsFilterer) ParseWithdrawal(log types.Log) (*MRC20EventsWithdrawal, error) {
	event := new(MRC20EventsWithdrawal)
	if err := _MRC20Events.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
