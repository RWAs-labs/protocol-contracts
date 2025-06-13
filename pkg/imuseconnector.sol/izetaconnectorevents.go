// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imuseconnector

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

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// IMuseConnectorEventsMetaData contains all meta data concerning the IMuseConnectorEvents contract.
var IMuseConnectorEventsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"UpdatedMuseConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false}]",
}

// IMuseConnectorEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use IMuseConnectorEventsMetaData.ABI instead.
var IMuseConnectorEventsABI = IMuseConnectorEventsMetaData.ABI

// IMuseConnectorEvents is an auto generated Go binding around an Ethereum contract.
type IMuseConnectorEvents struct {
	IMuseConnectorEventsCaller     // Read-only binding to the contract
	IMuseConnectorEventsTransactor // Write-only binding to the contract
	IMuseConnectorEventsFilterer   // Log filterer for contract events
}

// IMuseConnectorEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMuseConnectorEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseConnectorEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMuseConnectorEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseConnectorEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMuseConnectorEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMuseConnectorEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMuseConnectorEventsSession struct {
	Contract     *IMuseConnectorEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IMuseConnectorEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMuseConnectorEventsCallerSession struct {
	Contract *IMuseConnectorEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IMuseConnectorEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMuseConnectorEventsTransactorSession struct {
	Contract     *IMuseConnectorEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IMuseConnectorEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMuseConnectorEventsRaw struct {
	Contract *IMuseConnectorEvents // Generic contract binding to access the raw methods on
}

// IMuseConnectorEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMuseConnectorEventsCallerRaw struct {
	Contract *IMuseConnectorEventsCaller // Generic read-only contract binding to access the raw methods on
}

// IMuseConnectorEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMuseConnectorEventsTransactorRaw struct {
	Contract *IMuseConnectorEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMuseConnectorEvents creates a new instance of IMuseConnectorEvents, bound to a specific deployed contract.
func NewIMuseConnectorEvents(address common.Address, backend bind.ContractBackend) (*IMuseConnectorEvents, error) {
	contract, err := bindIMuseConnectorEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEvents{IMuseConnectorEventsCaller: IMuseConnectorEventsCaller{contract: contract}, IMuseConnectorEventsTransactor: IMuseConnectorEventsTransactor{contract: contract}, IMuseConnectorEventsFilterer: IMuseConnectorEventsFilterer{contract: contract}}, nil
}

// NewIMuseConnectorEventsCaller creates a new read-only instance of IMuseConnectorEvents, bound to a specific deployed contract.
func NewIMuseConnectorEventsCaller(address common.Address, caller bind.ContractCaller) (*IMuseConnectorEventsCaller, error) {
	contract, err := bindIMuseConnectorEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsCaller{contract: contract}, nil
}

// NewIMuseConnectorEventsTransactor creates a new write-only instance of IMuseConnectorEvents, bound to a specific deployed contract.
func NewIMuseConnectorEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*IMuseConnectorEventsTransactor, error) {
	contract, err := bindIMuseConnectorEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsTransactor{contract: contract}, nil
}

// NewIMuseConnectorEventsFilterer creates a new log filterer instance of IMuseConnectorEvents, bound to a specific deployed contract.
func NewIMuseConnectorEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*IMuseConnectorEventsFilterer, error) {
	contract, err := bindIMuseConnectorEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsFilterer{contract: contract}, nil
}

// bindIMuseConnectorEvents binds a generic wrapper to an already deployed contract.
func bindIMuseConnectorEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMuseConnectorEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuseConnectorEvents *IMuseConnectorEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuseConnectorEvents.Contract.IMuseConnectorEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuseConnectorEvents *IMuseConnectorEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuseConnectorEvents.Contract.IMuseConnectorEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuseConnectorEvents *IMuseConnectorEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuseConnectorEvents.Contract.IMuseConnectorEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMuseConnectorEvents *IMuseConnectorEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMuseConnectorEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMuseConnectorEvents *IMuseConnectorEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMuseConnectorEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMuseConnectorEvents *IMuseConnectorEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMuseConnectorEvents.Contract.contract.Transact(opts, method, params...)
}

// IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator is returned from FilterUpdatedMuseConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedMuseConnectorTSSAddress events raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator struct {
	Event *IMuseConnectorEventsUpdatedMuseConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseConnectorEventsUpdatedMuseConnectorTSSAddress)
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
		it.Event = new(IMuseConnectorEventsUpdatedMuseConnectorTSSAddress)
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
func (it *IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseConnectorEventsUpdatedMuseConnectorTSSAddress represents a UpdatedMuseConnectorTSSAddress event raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsUpdatedMuseConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMuseConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedMuseConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) FilterUpdatedMuseConnectorTSSAddress(opts *bind.FilterOpts) (*IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator, error) {

	logs, sub, err := _IMuseConnectorEvents.contract.FilterLogs(opts, "UpdatedMuseConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsUpdatedMuseConnectorTSSAddressIterator{contract: _IMuseConnectorEvents.contract, event: "UpdatedMuseConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedMuseConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedMuseConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) WatchUpdatedMuseConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *IMuseConnectorEventsUpdatedMuseConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _IMuseConnectorEvents.contract.WatchLogs(opts, "UpdatedMuseConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseConnectorEventsUpdatedMuseConnectorTSSAddress)
				if err := _IMuseConnectorEvents.contract.UnpackLog(event, "UpdatedMuseConnectorTSSAddress", log); err != nil {
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

// ParseUpdatedMuseConnectorTSSAddress is a log parse operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedMuseConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) ParseUpdatedMuseConnectorTSSAddress(log types.Log) (*IMuseConnectorEventsUpdatedMuseConnectorTSSAddress, error) {
	event := new(IMuseConnectorEventsUpdatedMuseConnectorTSSAddress)
	if err := _IMuseConnectorEvents.contract.UnpackLog(event, "UpdatedMuseConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMuseConnectorEventsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawnIterator struct {
	Event *IMuseConnectorEventsWithdrawn // Event containing the contract specifics and raw log

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
func (it *IMuseConnectorEventsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseConnectorEventsWithdrawn)
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
		it.Event = new(IMuseConnectorEventsWithdrawn)
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
func (it *IMuseConnectorEventsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseConnectorEventsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseConnectorEventsWithdrawn represents a Withdrawn event raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*IMuseConnectorEventsWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsWithdrawnIterator{contract: _IMuseConnectorEvents.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *IMuseConnectorEventsWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseConnectorEventsWithdrawn)
				if err := _IMuseConnectorEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) ParseWithdrawn(log types.Log) (*IMuseConnectorEventsWithdrawn, error) {
	event := new(IMuseConnectorEventsWithdrawn)
	if err := _IMuseConnectorEvents.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMuseConnectorEventsWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawnAndCalledIterator struct {
	Event *IMuseConnectorEventsWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *IMuseConnectorEventsWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseConnectorEventsWithdrawnAndCalled)
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
		it.Event = new(IMuseConnectorEventsWithdrawnAndCalled)
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
func (it *IMuseConnectorEventsWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseConnectorEventsWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseConnectorEventsWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*IMuseConnectorEventsWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsWithdrawnAndCalledIterator{contract: _IMuseConnectorEvents.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *IMuseConnectorEventsWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseConnectorEventsWithdrawnAndCalled)
				if err := _IMuseConnectorEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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

// ParseWithdrawnAndCalled is a log parse operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) ParseWithdrawnAndCalled(log types.Log) (*IMuseConnectorEventsWithdrawnAndCalled, error) {
	event := new(IMuseConnectorEventsWithdrawnAndCalled)
	if err := _IMuseConnectorEvents.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMuseConnectorEventsWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawnAndRevertedIterator struct {
	Event *IMuseConnectorEventsWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *IMuseConnectorEventsWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMuseConnectorEventsWithdrawnAndReverted)
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
		it.Event = new(IMuseConnectorEventsWithdrawnAndReverted)
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
func (it *IMuseConnectorEventsWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMuseConnectorEventsWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMuseConnectorEventsWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the IMuseConnectorEvents contract.
type IMuseConnectorEventsWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*IMuseConnectorEventsWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &IMuseConnectorEventsWithdrawnAndRevertedIterator{contract: _IMuseConnectorEvents.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *IMuseConnectorEventsWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMuseConnectorEvents.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMuseConnectorEventsWithdrawnAndReverted)
				if err := _IMuseConnectorEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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

// ParseWithdrawnAndReverted is a log parse operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_IMuseConnectorEvents *IMuseConnectorEventsFilterer) ParseWithdrawnAndReverted(log types.Log) (*IMuseConnectorEventsWithdrawnAndReverted, error) {
	event := new(IMuseConnectorEventsWithdrawnAndReverted)
	if err := _IMuseConnectorEvents.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
