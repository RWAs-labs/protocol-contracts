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

// MuseTokenConsumerMetaData contains all meta data concerning the MuseTokenConsumer contract.
var MuseTokenConsumerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getEthFromMuse\",\"inputs\":[{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"museTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTokenFromMuse\",\"inputs\":[{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"museTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMuseFromEth\",\"inputs\":[{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getMuseFromToken\",\"inputs\":[{\"name\":\"destinationAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"minAmountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasMuseLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EthExchangedForMuse\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenExchangedForMuse\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MuseExchangedForEth\",\"inputs\":[{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MuseExchangedForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MuseTokenConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseTokenConsumerMetaData.ABI instead.
var MuseTokenConsumerABI = MuseTokenConsumerMetaData.ABI

// MuseTokenConsumer is an auto generated Go binding around an Ethereum contract.
type MuseTokenConsumer struct {
	MuseTokenConsumerCaller     // Read-only binding to the contract
	MuseTokenConsumerTransactor // Write-only binding to the contract
	MuseTokenConsumerFilterer   // Log filterer for contract events
}

// MuseTokenConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseTokenConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseTokenConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseTokenConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseTokenConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseTokenConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseTokenConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseTokenConsumerSession struct {
	Contract     *MuseTokenConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MuseTokenConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseTokenConsumerCallerSession struct {
	Contract *MuseTokenConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MuseTokenConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseTokenConsumerTransactorSession struct {
	Contract     *MuseTokenConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MuseTokenConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseTokenConsumerRaw struct {
	Contract *MuseTokenConsumer // Generic contract binding to access the raw methods on
}

// MuseTokenConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseTokenConsumerCallerRaw struct {
	Contract *MuseTokenConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// MuseTokenConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseTokenConsumerTransactorRaw struct {
	Contract *MuseTokenConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseTokenConsumer creates a new instance of MuseTokenConsumer, bound to a specific deployed contract.
func NewMuseTokenConsumer(address common.Address, backend bind.ContractBackend) (*MuseTokenConsumer, error) {
	contract, err := bindMuseTokenConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumer{MuseTokenConsumerCaller: MuseTokenConsumerCaller{contract: contract}, MuseTokenConsumerTransactor: MuseTokenConsumerTransactor{contract: contract}, MuseTokenConsumerFilterer: MuseTokenConsumerFilterer{contract: contract}}, nil
}

// NewMuseTokenConsumerCaller creates a new read-only instance of MuseTokenConsumer, bound to a specific deployed contract.
func NewMuseTokenConsumerCaller(address common.Address, caller bind.ContractCaller) (*MuseTokenConsumerCaller, error) {
	contract, err := bindMuseTokenConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerCaller{contract: contract}, nil
}

// NewMuseTokenConsumerTransactor creates a new write-only instance of MuseTokenConsumer, bound to a specific deployed contract.
func NewMuseTokenConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseTokenConsumerTransactor, error) {
	contract, err := bindMuseTokenConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerTransactor{contract: contract}, nil
}

// NewMuseTokenConsumerFilterer creates a new log filterer instance of MuseTokenConsumer, bound to a specific deployed contract.
func NewMuseTokenConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseTokenConsumerFilterer, error) {
	contract, err := bindMuseTokenConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerFilterer{contract: contract}, nil
}

// bindMuseTokenConsumer binds a generic wrapper to an already deployed contract.
func bindMuseTokenConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseTokenConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseTokenConsumer *MuseTokenConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseTokenConsumer.Contract.MuseTokenConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseTokenConsumer *MuseTokenConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.MuseTokenConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseTokenConsumer *MuseTokenConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.MuseTokenConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseTokenConsumer *MuseTokenConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseTokenConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseTokenConsumer *MuseTokenConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseTokenConsumer *MuseTokenConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.contract.Transact(opts, method, params...)
}

// HasMuseLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasMuseLiquidity() view returns(bool)
func (_MuseTokenConsumer *MuseTokenConsumerCaller) HasMuseLiquidity(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MuseTokenConsumer.contract.Call(opts, &out, "hasMuseLiquidity")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMuseLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasMuseLiquidity() view returns(bool)
func (_MuseTokenConsumer *MuseTokenConsumerSession) HasMuseLiquidity() (bool, error) {
	return _MuseTokenConsumer.Contract.HasMuseLiquidity(&_MuseTokenConsumer.CallOpts)
}

// HasMuseLiquidity is a free data retrieval call binding the contract method 0x80801f84.
//
// Solidity: function hasMuseLiquidity() view returns(bool)
func (_MuseTokenConsumer *MuseTokenConsumerCallerSession) HasMuseLiquidity() (bool, error) {
	return _MuseTokenConsumer.Contract.HasMuseLiquidity(&_MuseTokenConsumer.CallOpts)
}

// GetEthFromMuse is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromMuse(address destinationAddress, uint256 minAmountOut, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactor) GetEthFromMuse(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.contract.Transact(opts, "getEthFromMuse", destinationAddress, minAmountOut, museTokenAmount)
}

// GetEthFromMuse is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromMuse(address destinationAddress, uint256 minAmountOut, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerSession) GetEthFromMuse(destinationAddress common.Address, minAmountOut *big.Int, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetEthFromMuse(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, museTokenAmount)
}

// GetEthFromMuse is a paid mutator transaction binding the contract method 0x54c49a2a.
//
// Solidity: function getEthFromMuse(address destinationAddress, uint256 minAmountOut, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactorSession) GetEthFromMuse(destinationAddress common.Address, minAmountOut *big.Int, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetEthFromMuse(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, museTokenAmount)
}

// GetTokenFromMuse is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromMuse(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactor) GetTokenFromMuse(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.contract.Transact(opts, "getTokenFromMuse", destinationAddress, minAmountOut, outputToken, museTokenAmount)
}

// GetTokenFromMuse is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromMuse(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerSession) GetTokenFromMuse(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetTokenFromMuse(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, museTokenAmount)
}

// GetTokenFromMuse is a paid mutator transaction binding the contract method 0xa53fb10b.
//
// Solidity: function getTokenFromMuse(address destinationAddress, uint256 minAmountOut, address outputToken, uint256 museTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactorSession) GetTokenFromMuse(destinationAddress common.Address, minAmountOut *big.Int, outputToken common.Address, museTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetTokenFromMuse(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, outputToken, museTokenAmount)
}

// GetMuseFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getMuseFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactor) GetMuseFromEth(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.contract.Transact(opts, "getMuseFromEth", destinationAddress, minAmountOut)
}

// GetMuseFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getMuseFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerSession) GetMuseFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetMuseFromEth(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetMuseFromEth is a paid mutator transaction binding the contract method 0x013b2ff8.
//
// Solidity: function getMuseFromEth(address destinationAddress, uint256 minAmountOut) payable returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactorSession) GetMuseFromEth(destinationAddress common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetMuseFromEth(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut)
}

// GetMuseFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getMuseFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactor) GetMuseFromToken(opts *bind.TransactOpts, destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.contract.Transact(opts, "getMuseFromToken", destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetMuseFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getMuseFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerSession) GetMuseFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetMuseFromToken(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// GetMuseFromToken is a paid mutator transaction binding the contract method 0x2405620a.
//
// Solidity: function getMuseFromToken(address destinationAddress, uint256 minAmountOut, address inputToken, uint256 inputTokenAmount) returns(uint256)
func (_MuseTokenConsumer *MuseTokenConsumerTransactorSession) GetMuseFromToken(destinationAddress common.Address, minAmountOut *big.Int, inputToken common.Address, inputTokenAmount *big.Int) (*types.Transaction, error) {
	return _MuseTokenConsumer.Contract.GetMuseFromToken(&_MuseTokenConsumer.TransactOpts, destinationAddress, minAmountOut, inputToken, inputTokenAmount)
}

// MuseTokenConsumerEthExchangedForMuseIterator is returned from FilterEthExchangedForMuse and is used to iterate over the raw logs and unpacked data for EthExchangedForMuse events raised by the MuseTokenConsumer contract.
type MuseTokenConsumerEthExchangedForMuseIterator struct {
	Event *MuseTokenConsumerEthExchangedForMuse // Event containing the contract specifics and raw log

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
func (it *MuseTokenConsumerEthExchangedForMuseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseTokenConsumerEthExchangedForMuse)
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
		it.Event = new(MuseTokenConsumerEthExchangedForMuse)
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
func (it *MuseTokenConsumerEthExchangedForMuseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseTokenConsumerEthExchangedForMuseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseTokenConsumerEthExchangedForMuse represents a EthExchangedForMuse event raised by the MuseTokenConsumer contract.
type MuseTokenConsumerEthExchangedForMuse struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEthExchangedForMuse is a free log retrieval operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForMuse(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) FilterEthExchangedForMuse(opts *bind.FilterOpts) (*MuseTokenConsumerEthExchangedForMuseIterator, error) {

	logs, sub, err := _MuseTokenConsumer.contract.FilterLogs(opts, "EthExchangedForMuse")
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerEthExchangedForMuseIterator{contract: _MuseTokenConsumer.contract, event: "EthExchangedForMuse", logs: logs, sub: sub}, nil
}

// WatchEthExchangedForMuse is a free log subscription operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForMuse(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) WatchEthExchangedForMuse(opts *bind.WatchOpts, sink chan<- *MuseTokenConsumerEthExchangedForMuse) (event.Subscription, error) {

	logs, sub, err := _MuseTokenConsumer.contract.WatchLogs(opts, "EthExchangedForMuse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseTokenConsumerEthExchangedForMuse)
				if err := _MuseTokenConsumer.contract.UnpackLog(event, "EthExchangedForMuse", log); err != nil {
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

// ParseEthExchangedForMuse is a log parse operation binding the contract event 0x87890b0a30955b1db16cc894fbe24779ae05d9f337bfe8b6dfc0809b5bf9da11.
//
// Solidity: event EthExchangedForMuse(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) ParseEthExchangedForMuse(log types.Log) (*MuseTokenConsumerEthExchangedForMuse, error) {
	event := new(MuseTokenConsumerEthExchangedForMuse)
	if err := _MuseTokenConsumer.contract.UnpackLog(event, "EthExchangedForMuse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseTokenConsumerTokenExchangedForMuseIterator is returned from FilterTokenExchangedForMuse and is used to iterate over the raw logs and unpacked data for TokenExchangedForMuse events raised by the MuseTokenConsumer contract.
type MuseTokenConsumerTokenExchangedForMuseIterator struct {
	Event *MuseTokenConsumerTokenExchangedForMuse // Event containing the contract specifics and raw log

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
func (it *MuseTokenConsumerTokenExchangedForMuseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseTokenConsumerTokenExchangedForMuse)
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
		it.Event = new(MuseTokenConsumerTokenExchangedForMuse)
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
func (it *MuseTokenConsumerTokenExchangedForMuseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseTokenConsumerTokenExchangedForMuseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseTokenConsumerTokenExchangedForMuse represents a TokenExchangedForMuse event raised by the MuseTokenConsumer contract.
type MuseTokenConsumerTokenExchangedForMuse struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenExchangedForMuse is a free log retrieval operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForMuse(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) FilterTokenExchangedForMuse(opts *bind.FilterOpts) (*MuseTokenConsumerTokenExchangedForMuseIterator, error) {

	logs, sub, err := _MuseTokenConsumer.contract.FilterLogs(opts, "TokenExchangedForMuse")
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerTokenExchangedForMuseIterator{contract: _MuseTokenConsumer.contract, event: "TokenExchangedForMuse", logs: logs, sub: sub}, nil
}

// WatchTokenExchangedForMuse is a free log subscription operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForMuse(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) WatchTokenExchangedForMuse(opts *bind.WatchOpts, sink chan<- *MuseTokenConsumerTokenExchangedForMuse) (event.Subscription, error) {

	logs, sub, err := _MuseTokenConsumer.contract.WatchLogs(opts, "TokenExchangedForMuse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseTokenConsumerTokenExchangedForMuse)
				if err := _MuseTokenConsumer.contract.UnpackLog(event, "TokenExchangedForMuse", log); err != nil {
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

// ParseTokenExchangedForMuse is a log parse operation binding the contract event 0x017190d3d99ee6d8dd0604ef1e71ff9802810c6485f57c9b2ed6169848dd119f.
//
// Solidity: event TokenExchangedForMuse(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) ParseTokenExchangedForMuse(log types.Log) (*MuseTokenConsumerTokenExchangedForMuse, error) {
	event := new(MuseTokenConsumerTokenExchangedForMuse)
	if err := _MuseTokenConsumer.contract.UnpackLog(event, "TokenExchangedForMuse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseTokenConsumerMuseExchangedForEthIterator is returned from FilterMuseExchangedForEth and is used to iterate over the raw logs and unpacked data for MuseExchangedForEth events raised by the MuseTokenConsumer contract.
type MuseTokenConsumerMuseExchangedForEthIterator struct {
	Event *MuseTokenConsumerMuseExchangedForEth // Event containing the contract specifics and raw log

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
func (it *MuseTokenConsumerMuseExchangedForEthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseTokenConsumerMuseExchangedForEth)
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
		it.Event = new(MuseTokenConsumerMuseExchangedForEth)
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
func (it *MuseTokenConsumerMuseExchangedForEthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseTokenConsumerMuseExchangedForEthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseTokenConsumerMuseExchangedForEth represents a MuseExchangedForEth event raised by the MuseTokenConsumer contract.
type MuseTokenConsumerMuseExchangedForEth struct {
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMuseExchangedForEth is a free log retrieval operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event MuseExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) FilterMuseExchangedForEth(opts *bind.FilterOpts) (*MuseTokenConsumerMuseExchangedForEthIterator, error) {

	logs, sub, err := _MuseTokenConsumer.contract.FilterLogs(opts, "MuseExchangedForEth")
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerMuseExchangedForEthIterator{contract: _MuseTokenConsumer.contract, event: "MuseExchangedForEth", logs: logs, sub: sub}, nil
}

// WatchMuseExchangedForEth is a free log subscription operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event MuseExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) WatchMuseExchangedForEth(opts *bind.WatchOpts, sink chan<- *MuseTokenConsumerMuseExchangedForEth) (event.Subscription, error) {

	logs, sub, err := _MuseTokenConsumer.contract.WatchLogs(opts, "MuseExchangedForEth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseTokenConsumerMuseExchangedForEth)
				if err := _MuseTokenConsumer.contract.UnpackLog(event, "MuseExchangedForEth", log); err != nil {
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

// ParseMuseExchangedForEth is a log parse operation binding the contract event 0x74e171117e91660f493740924d8bad0caf48dc4fbccb767fb05935397a2c17ae.
//
// Solidity: event MuseExchangedForEth(uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) ParseMuseExchangedForEth(log types.Log) (*MuseTokenConsumerMuseExchangedForEth, error) {
	event := new(MuseTokenConsumerMuseExchangedForEth)
	if err := _MuseTokenConsumer.contract.UnpackLog(event, "MuseExchangedForEth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseTokenConsumerMuseExchangedForTokenIterator is returned from FilterMuseExchangedForToken and is used to iterate over the raw logs and unpacked data for MuseExchangedForToken events raised by the MuseTokenConsumer contract.
type MuseTokenConsumerMuseExchangedForTokenIterator struct {
	Event *MuseTokenConsumerMuseExchangedForToken // Event containing the contract specifics and raw log

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
func (it *MuseTokenConsumerMuseExchangedForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseTokenConsumerMuseExchangedForToken)
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
		it.Event = new(MuseTokenConsumerMuseExchangedForToken)
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
func (it *MuseTokenConsumerMuseExchangedForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseTokenConsumerMuseExchangedForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseTokenConsumerMuseExchangedForToken represents a MuseExchangedForToken event raised by the MuseTokenConsumer contract.
type MuseTokenConsumerMuseExchangedForToken struct {
	Token     common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMuseExchangedForToken is a free log retrieval operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event MuseExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) FilterMuseExchangedForToken(opts *bind.FilterOpts) (*MuseTokenConsumerMuseExchangedForTokenIterator, error) {

	logs, sub, err := _MuseTokenConsumer.contract.FilterLogs(opts, "MuseExchangedForToken")
	if err != nil {
		return nil, err
	}
	return &MuseTokenConsumerMuseExchangedForTokenIterator{contract: _MuseTokenConsumer.contract, event: "MuseExchangedForToken", logs: logs, sub: sub}, nil
}

// WatchMuseExchangedForToken is a free log subscription operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event MuseExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) WatchMuseExchangedForToken(opts *bind.WatchOpts, sink chan<- *MuseTokenConsumerMuseExchangedForToken) (event.Subscription, error) {

	logs, sub, err := _MuseTokenConsumer.contract.WatchLogs(opts, "MuseExchangedForToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseTokenConsumerMuseExchangedForToken)
				if err := _MuseTokenConsumer.contract.UnpackLog(event, "MuseExchangedForToken", log); err != nil {
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

// ParseMuseExchangedForToken is a log parse operation binding the contract event 0x0a7cb8f6e1d29e616c1209a3f418c17b3a9137005377f6dd072217b1ede2573b.
//
// Solidity: event MuseExchangedForToken(address token, uint256 amountIn, uint256 amountOut)
func (_MuseTokenConsumer *MuseTokenConsumerFilterer) ParseMuseExchangedForToken(log types.Log) (*MuseTokenConsumerMuseExchangedForToken, error) {
	event := new(MuseTokenConsumerMuseExchangedForToken)
	if err := _MuseTokenConsumer.contract.UnpackLog(event, "MuseExchangedForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
