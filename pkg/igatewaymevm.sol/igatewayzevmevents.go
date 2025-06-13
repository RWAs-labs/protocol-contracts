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

// RevertOptions is an auto generated low-level Go binding around an user-defined struct.
type RevertOptions struct {
	RevertAddress    common.Address
	CallOnRevert     bool
	AbortAddress     common.Address
	RevertMessage    []byte
	OnRevertGasLimit *big.Int
}

// IGatewayMEVMEventsMetaData contains all meta data concerning the IGatewayMEVMEvents contract.
var IGatewayMEVMEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false}]",
}

// IGatewayMEVMEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IGatewayMEVMEventsMetaData.ABI instead.
var IGatewayMEVMEventsABI = IGatewayMEVMEventsMetaData.ABI

// IGatewayMEVMEvents is an auto generated Go binding around an Ethereum contract.
type IGatewayMEVMEvents struct {
	IGatewayMEVMEventsCaller     // Read-only binding to the contract
	IGatewayMEVMEventsTransactor // Write-only binding to the contract
	IGatewayMEVMEventsFilterer   // Log filterer for contract events
}

// IGatewayMEVMEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGatewayMEVMEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGatewayMEVMEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGatewayMEVMEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGatewayMEVMEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGatewayMEVMEventsSession struct {
	Contract     *IGatewayMEVMEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IGatewayMEVMEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGatewayMEVMEventsCallerSession struct {
	Contract *IGatewayMEVMEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IGatewayMEVMEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGatewayMEVMEventsTransactorSession struct {
	Contract     *IGatewayMEVMEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IGatewayMEVMEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGatewayMEVMEventsRaw struct {
	Contract *IGatewayMEVMEvents // Generic contract binding to access the raw methods on
}

// IGatewayMEVMEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGatewayMEVMEventsCallerRaw struct {
	Contract *IGatewayMEVMEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IGatewayMEVMEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGatewayMEVMEventsTransactorRaw struct {
	Contract *IGatewayMEVMEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGatewayMEVMEvents creates a new instance of IGatewayMEVMEvents, bound to a specific deployed contract.
func NewIGatewayMEVMEvents(address common.Address, backend bind.ContractBackend) (*IGatewayMEVMEvents, error) {
	contract, err := bindIGatewayMEVMEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEvents{IGatewayMEVMEventsCaller: IGatewayMEVMEventsCaller{contract: contract}, IGatewayMEVMEventsTransactor: IGatewayMEVMEventsTransactor{contract: contract}, IGatewayMEVMEventsFilterer: IGatewayMEVMEventsFilterer{contract: contract}}, nil
}

// NewIGatewayMEVMEventsCaller creates a new read-only instance of IGatewayMEVMEvents, bound to a specific deployed contract.
func NewIGatewayMEVMEventsCaller(address common.Address, caller bind.ContractCaller) (*IGatewayMEVMEventsCaller, error) {
	contract, err := bindIGatewayMEVMEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsCaller{contract: contract}, nil
}

// NewIGatewayMEVMEventsTransactor creates a new write-only instance of IGatewayMEVMEvents, bound to a specific deployed contract.
func NewIGatewayMEVMEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IGatewayMEVMEventsTransactor, error) {
	contract, err := bindIGatewayMEVMEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsTransactor{contract: contract}, nil
}

// NewIGatewayMEVMEventsFilterer creates a new log filterer instance of IGatewayMEVMEvents, bound to a specific deployed contract.
func NewIGatewayMEVMEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IGatewayMEVMEventsFilterer, error) {
	contract, err := bindIGatewayMEVMEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsFilterer{contract: contract}, nil
}

// bindIGatewayMEVMEvents binds a generic wrapper to an already deployed contract.
func bindIGatewayMEVMEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGatewayMEVMEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVMEvents.Contract.IGatewayMEVMEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVMEvents.Contract.IGatewayMEVMEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVMEvents.Contract.IGatewayMEVMEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGatewayMEVMEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGatewayMEVMEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGatewayMEVMEvents *IGatewayMEVMEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGatewayMEVMEvents.Contract.contract.Transact(opts, method, params...)
}

// IGatewayMEVMEventsCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsCalledIterator struct {
	Event *IGatewayMEVMEventsCalled // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMEventsCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMEventsCalled)
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
		it.Event = new(IGatewayMEVMEventsCalled)
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
func (it *IGatewayMEVMEventsCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMEventsCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMEventsCalled represents a Called event raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsCalled struct {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, mrc20 []common.Address) (*IGatewayMEVMEventsCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.FilterLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsCalledIterator{contract: _IGatewayMEVMEvents.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMEventsCalled, sender []common.Address, mrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.WatchLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMEventsCalled)
				if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) ParseCalled(log types.Log) (*IGatewayMEVMEventsCalled, error) {
	event := new(IGatewayMEVMEventsCalled)
	if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayMEVMEventsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsWithdrawnIterator struct {
	Event *IGatewayMEVMEventsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMEventsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMEventsWithdrawn)
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
		it.Event = new(IGatewayMEVMEventsWithdrawn)
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
func (it *IGatewayMEVMEventsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMEventsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMEventsWithdrawn represents a Withdrawn event raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsWithdrawn struct {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*IGatewayMEVMEventsWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsWithdrawnIterator{contract: _IGatewayMEVMEvents.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMEventsWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMEventsWithdrawn)
				if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) ParseWithdrawn(log types.Log) (*IGatewayMEVMEventsWithdrawn, error) {
	event := new(IGatewayMEVMEventsWithdrawn)
	if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGatewayMEVMEventsWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsWithdrawnAndCalledIterator struct {
	Event *IGatewayMEVMEventsWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IGatewayMEVMEventsWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGatewayMEVMEventsWithdrawnAndCalled)
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
		it.Event = new(IGatewayMEVMEventsWithdrawnAndCalled)
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
func (it *IGatewayMEVMEventsWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGatewayMEVMEventsWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGatewayMEVMEventsWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IGatewayMEVMEvents contract.
type IGatewayMEVMEventsWithdrawnAndCalled struct {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*IGatewayMEVMEventsWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &IGatewayMEVMEventsWithdrawnAndCalledIterator{contract: _IGatewayMEVMEvents.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IGatewayMEVMEventsWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _IGatewayMEVMEvents.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGatewayMEVMEventsWithdrawnAndCalled)
				if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_IGatewayMEVMEvents *IGatewayMEVMEventsFilterer) ParseWithdrawnAndCalled(log types.Log) (*IGatewayMEVMEventsWithdrawnAndCalled, error) {
	event := new(IGatewayMEVMEventsWithdrawnAndCalled)
	if err := _IGatewayMEVMEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
