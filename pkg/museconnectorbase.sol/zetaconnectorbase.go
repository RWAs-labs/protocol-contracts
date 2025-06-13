// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package museconnectorbase

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

// MessageContext is an auto generated low-level Go binding around an user-defined struct.
type MessageContext struct {
	Sender common.Address
}

// RevertContext is an auto generated low-level Go binding around an user-defined struct.
type RevertContext struct {
	Sender        common.Address
	Asset         common.Address
	Amount        *big.Int
	RevertMessage []byte
}

// MuseConnectorBaseMetaData contains all meta data concerning the MuseConnectorBase contract.
var MuseConnectorBaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TSS_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gateway\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIGatewayEVM\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"gateway_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"museToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveTokens\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateTSSAddress\",\"inputs\":[{\"name\":\"newTSSAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"messageContext\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndRevert\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"museToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatedMuseConnectorTSSAddress\",\"inputs\":[{\"name\":\"oldTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTSSAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndReverted\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// MuseConnectorBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseConnectorBaseMetaData.ABI instead.
var MuseConnectorBaseABI = MuseConnectorBaseMetaData.ABI

// MuseConnectorBase is an auto generated Go binding around an Ethereum contract.
type MuseConnectorBase struct {
	MuseConnectorBaseCaller     // Read-only binding to the contract
	MuseConnectorBaseTransactor // Write-only binding to the contract
	MuseConnectorBaseFilterer   // Log filterer for contract events
}

// MuseConnectorBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseConnectorBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseConnectorBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseConnectorBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseConnectorBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseConnectorBaseSession struct {
	Contract     *MuseConnectorBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MuseConnectorBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseConnectorBaseCallerSession struct {
	Contract *MuseConnectorBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MuseConnectorBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseConnectorBaseTransactorSession struct {
	Contract     *MuseConnectorBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MuseConnectorBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseConnectorBaseRaw struct {
	Contract *MuseConnectorBase // Generic contract binding to access the raw methods on
}

// MuseConnectorBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseConnectorBaseCallerRaw struct {
	Contract *MuseConnectorBaseCaller // Generic read-only contract binding to access the raw methods on
}

// MuseConnectorBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseConnectorBaseTransactorRaw struct {
	Contract *MuseConnectorBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseConnectorBase creates a new instance of MuseConnectorBase, bound to a specific deployed contract.
func NewMuseConnectorBase(address common.Address, backend bind.ContractBackend) (*MuseConnectorBase, error) {
	contract, err := bindMuseConnectorBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBase{MuseConnectorBaseCaller: MuseConnectorBaseCaller{contract: contract}, MuseConnectorBaseTransactor: MuseConnectorBaseTransactor{contract: contract}, MuseConnectorBaseFilterer: MuseConnectorBaseFilterer{contract: contract}}, nil
}

// NewMuseConnectorBaseCaller creates a new read-only instance of MuseConnectorBase, bound to a specific deployed contract.
func NewMuseConnectorBaseCaller(address common.Address, caller bind.ContractCaller) (*MuseConnectorBaseCaller, error) {
	contract, err := bindMuseConnectorBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseCaller{contract: contract}, nil
}

// NewMuseConnectorBaseTransactor creates a new write-only instance of MuseConnectorBase, bound to a specific deployed contract.
func NewMuseConnectorBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseConnectorBaseTransactor, error) {
	contract, err := bindMuseConnectorBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseTransactor{contract: contract}, nil
}

// NewMuseConnectorBaseFilterer creates a new log filterer instance of MuseConnectorBase, bound to a specific deployed contract.
func NewMuseConnectorBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseConnectorBaseFilterer, error) {
	contract, err := bindMuseConnectorBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseFilterer{contract: contract}, nil
}

// bindMuseConnectorBase binds a generic wrapper to an already deployed contract.
func bindMuseConnectorBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseConnectorBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnectorBase *MuseConnectorBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnectorBase.Contract.MuseConnectorBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnectorBase *MuseConnectorBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.MuseConnectorBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnectorBase *MuseConnectorBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.MuseConnectorBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseConnectorBase *MuseConnectorBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseConnectorBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseConnectorBase *MuseConnectorBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseConnectorBase *MuseConnectorBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.DEFAULTADMINROLE(&_MuseConnectorBase.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.DEFAULTADMINROLE(&_MuseConnectorBase.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) PAUSERROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.PAUSERROLE(&_MuseConnectorBase.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) PAUSERROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.PAUSERROLE(&_MuseConnectorBase.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) TSSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "TSS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) TSSROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.TSSROLE(&_MuseConnectorBase.CallOpts)
}

// TSSROLE is a free data retrieval call binding the contract method 0xa783c789.
//
// Solidity: function TSS_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) TSSROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.TSSROLE(&_MuseConnectorBase.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MuseConnectorBase *MuseConnectorBaseCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MuseConnectorBase *MuseConnectorBaseSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MuseConnectorBase.Contract.UPGRADEINTERFACEVERSION(&_MuseConnectorBase.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _MuseConnectorBase.Contract.UPGRADEINTERFACEVERSION(&_MuseConnectorBase.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) WITHDRAWERROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.WITHDRAWERROLE(&_MuseConnectorBase.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _MuseConnectorBase.Contract.WITHDRAWERROLE(&_MuseConnectorBase.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseSession) Gateway() (common.Address, error) {
	return _MuseConnectorBase.Contract.Gateway(&_MuseConnectorBase.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) Gateway() (common.Address, error) {
	return _MuseConnectorBase.Contract.Gateway(&_MuseConnectorBase.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MuseConnectorBase.Contract.GetRoleAdmin(&_MuseConnectorBase.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MuseConnectorBase.Contract.GetRoleAdmin(&_MuseConnectorBase.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MuseConnectorBase.Contract.HasRole(&_MuseConnectorBase.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MuseConnectorBase.Contract.HasRole(&_MuseConnectorBase.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseSession) Paused() (bool, error) {
	return _MuseConnectorBase.Contract.Paused(&_MuseConnectorBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) Paused() (bool, error) {
	return _MuseConnectorBase.Contract.Paused(&_MuseConnectorBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseSession) ProxiableUUID() ([32]byte, error) {
	return _MuseConnectorBase.Contract.ProxiableUUID(&_MuseConnectorBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MuseConnectorBase.Contract.ProxiableUUID(&_MuseConnectorBase.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MuseConnectorBase.Contract.SupportsInterface(&_MuseConnectorBase.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MuseConnectorBase.Contract.SupportsInterface(&_MuseConnectorBase.CallOpts, interfaceId)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseSession) TssAddress() (common.Address, error) {
	return _MuseConnectorBase.Contract.TssAddress(&_MuseConnectorBase.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) TssAddress() (common.Address, error) {
	return _MuseConnectorBase.Contract.TssAddress(&_MuseConnectorBase.CallOpts)
}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCaller) MuseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseConnectorBase.contract.Call(opts, &out, "museToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseSession) MuseToken() (common.Address, error) {
	return _MuseConnectorBase.Contract.MuseToken(&_MuseConnectorBase.CallOpts)
}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_MuseConnectorBase *MuseConnectorBaseCallerSession) MuseToken() (common.Address, error) {
	return _MuseConnectorBase.Contract.MuseToken(&_MuseConnectorBase.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.GrantRole(&_MuseConnectorBase.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.GrantRole(&_MuseConnectorBase.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address museToken_, address tssAddress_, address admin_) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) Initialize(opts *bind.TransactOpts, gateway_ common.Address, museToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "initialize", gateway_, museToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address museToken_, address tssAddress_, address admin_) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) Initialize(gateway_ common.Address, museToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Initialize(&_MuseConnectorBase.TransactOpts, gateway_, museToken_, tssAddress_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address gateway_, address museToken_, address tssAddress_, address admin_) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) Initialize(gateway_ common.Address, museToken_ common.Address, tssAddress_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Initialize(&_MuseConnectorBase.TransactOpts, gateway_, museToken_, tssAddress_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) Pause() (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Pause(&_MuseConnectorBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Pause(&_MuseConnectorBase.TransactOpts)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) ReceiveTokens(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "receiveTokens", amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.ReceiveTokens(&_MuseConnectorBase.TransactOpts, amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x743e0c9b.
//
// Solidity: function receiveTokens(uint256 amount) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) ReceiveTokens(amount *big.Int) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.ReceiveTokens(&_MuseConnectorBase.TransactOpts, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.RenounceRole(&_MuseConnectorBase.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.RenounceRole(&_MuseConnectorBase.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.RevokeRole(&_MuseConnectorBase.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.RevokeRole(&_MuseConnectorBase.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) Unpause() (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Unpause(&_MuseConnectorBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Unpause(&_MuseConnectorBase.TransactOpts)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) UpdateTSSAddress(opts *bind.TransactOpts, newTSSAddress common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "updateTSSAddress", newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.UpdateTSSAddress(&_MuseConnectorBase.TransactOpts, newTSSAddress)
}

// UpdateTSSAddress is a paid mutator transaction binding the contract method 0x950837aa.
//
// Solidity: function updateTSSAddress(address newTSSAddress) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) UpdateTSSAddress(newTSSAddress common.Address) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.UpdateTSSAddress(&_MuseConnectorBase.TransactOpts, newTSSAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.UpgradeToAndCall(&_MuseConnectorBase.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.UpgradeToAndCall(&_MuseConnectorBase.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "withdraw", to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Withdraw(&_MuseConnectorBase.TransactOpts, to, amount, internalSendHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x106e6290.
//
// Solidity: function withdraw(address to, uint256 amount, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) Withdraw(to common.Address, amount *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.Withdraw(&_MuseConnectorBase.TransactOpts, to, amount, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) WithdrawAndCall(opts *bind.TransactOpts, messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "withdrawAndCall", messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.WithdrawAndCall(&_MuseConnectorBase.TransactOpts, messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x6fb9a7af.
//
// Solidity: function withdrawAndCall((address) messageContext, address to, uint256 amount, bytes data, bytes32 internalSendHash) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) WithdrawAndCall(messageContext MessageContext, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.WithdrawAndCall(&_MuseConnectorBase.TransactOpts, messageContext, to, amount, data, internalSendHash)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactor) WithdrawAndRevert(opts *bind.TransactOpts, to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _MuseConnectorBase.contract.Transact(opts, "withdrawAndRevert", to, amount, data, internalSendHash, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_MuseConnectorBase *MuseConnectorBaseSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.WithdrawAndRevert(&_MuseConnectorBase.TransactOpts, to, amount, data, internalSendHash, revertContext)
}

// WithdrawAndRevert is a paid mutator transaction binding the contract method 0x6f8728ad.
//
// Solidity: function withdrawAndRevert(address to, uint256 amount, bytes data, bytes32 internalSendHash, (address,address,uint256,bytes) revertContext) returns()
func (_MuseConnectorBase *MuseConnectorBaseTransactorSession) WithdrawAndRevert(to common.Address, amount *big.Int, data []byte, internalSendHash [32]byte, revertContext RevertContext) (*types.Transaction, error) {
	return _MuseConnectorBase.Contract.WithdrawAndRevert(&_MuseConnectorBase.TransactOpts, to, amount, data, internalSendHash, revertContext)
}

// MuseConnectorBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MuseConnectorBase contract.
type MuseConnectorBaseInitializedIterator struct {
	Event *MuseConnectorBaseInitialized // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseInitialized)
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
		it.Event = new(MuseConnectorBaseInitialized)
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
func (it *MuseConnectorBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseInitialized represents a Initialized event raised by the MuseConnectorBase contract.
type MuseConnectorBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*MuseConnectorBaseInitializedIterator, error) {

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseInitializedIterator{contract: _MuseConnectorBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseInitialized)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseInitialized(log types.Log) (*MuseConnectorBaseInitialized, error) {
	event := new(MuseConnectorBaseInitialized)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MuseConnectorBase contract.
type MuseConnectorBasePausedIterator struct {
	Event *MuseConnectorBasePaused // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBasePaused)
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
		it.Event = new(MuseConnectorBasePaused)
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
func (it *MuseConnectorBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBasePaused represents a Paused event raised by the MuseConnectorBase contract.
type MuseConnectorBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*MuseConnectorBasePausedIterator, error) {

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBasePausedIterator{contract: _MuseConnectorBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MuseConnectorBasePaused) (event.Subscription, error) {

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBasePaused)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParsePaused(log types.Log) (*MuseConnectorBasePaused, error) {
	event := new(MuseConnectorBasePaused)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleAdminChangedIterator struct {
	Event *MuseConnectorBaseRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseRoleAdminChanged)
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
		it.Event = new(MuseConnectorBaseRoleAdminChanged)
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
func (it *MuseConnectorBaseRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseRoleAdminChanged represents a RoleAdminChanged event raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MuseConnectorBaseRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseRoleAdminChangedIterator{contract: _MuseConnectorBase.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseRoleAdminChanged)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseRoleAdminChanged(log types.Log) (*MuseConnectorBaseRoleAdminChanged, error) {
	event := new(MuseConnectorBaseRoleAdminChanged)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleGrantedIterator struct {
	Event *MuseConnectorBaseRoleGranted // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseRoleGranted)
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
		it.Event = new(MuseConnectorBaseRoleGranted)
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
func (it *MuseConnectorBaseRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseRoleGranted represents a RoleGranted event raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MuseConnectorBaseRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseRoleGrantedIterator{contract: _MuseConnectorBase.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseRoleGranted)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseRoleGranted(log types.Log) (*MuseConnectorBaseRoleGranted, error) {
	event := new(MuseConnectorBaseRoleGranted)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleRevokedIterator struct {
	Event *MuseConnectorBaseRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseRoleRevoked)
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
		it.Event = new(MuseConnectorBaseRoleRevoked)
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
func (it *MuseConnectorBaseRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseRoleRevoked represents a RoleRevoked event raised by the MuseConnectorBase contract.
type MuseConnectorBaseRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MuseConnectorBaseRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseRoleRevokedIterator{contract: _MuseConnectorBase.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseRoleRevoked)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseRoleRevoked(log types.Log) (*MuseConnectorBaseRoleRevoked, error) {
	event := new(MuseConnectorBaseRoleRevoked)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MuseConnectorBase contract.
type MuseConnectorBaseUnpausedIterator struct {
	Event *MuseConnectorBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseUnpaused)
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
		it.Event = new(MuseConnectorBaseUnpaused)
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
func (it *MuseConnectorBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseUnpaused represents a Unpaused event raised by the MuseConnectorBase contract.
type MuseConnectorBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MuseConnectorBaseUnpausedIterator, error) {

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseUnpausedIterator{contract: _MuseConnectorBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseUnpaused)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseUnpaused(log types.Log) (*MuseConnectorBaseUnpaused, error) {
	event := new(MuseConnectorBaseUnpaused)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator is returned from FilterUpdatedMuseConnectorTSSAddress and is used to iterate over the raw logs and unpacked data for UpdatedMuseConnectorTSSAddress events raised by the MuseConnectorBase contract.
type MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator struct {
	Event *MuseConnectorBaseUpdatedMuseConnectorTSSAddress // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseUpdatedMuseConnectorTSSAddress)
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
		it.Event = new(MuseConnectorBaseUpdatedMuseConnectorTSSAddress)
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
func (it *MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseUpdatedMuseConnectorTSSAddress represents a UpdatedMuseConnectorTSSAddress event raised by the MuseConnectorBase contract.
type MuseConnectorBaseUpdatedMuseConnectorTSSAddress struct {
	OldTSSAddress common.Address
	NewTSSAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMuseConnectorTSSAddress is a free log retrieval operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedMuseConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterUpdatedMuseConnectorTSSAddress(opts *bind.FilterOpts) (*MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator, error) {

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "UpdatedMuseConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseUpdatedMuseConnectorTSSAddressIterator{contract: _MuseConnectorBase.contract, event: "UpdatedMuseConnectorTSSAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedMuseConnectorTSSAddress is a free log subscription operation binding the contract event 0x33770ab682353c17917ad3e667f05905fc8dda00671ef1ed33bef9bc8db0323e.
//
// Solidity: event UpdatedMuseConnectorTSSAddress(address oldTSSAddress, address newTSSAddress)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchUpdatedMuseConnectorTSSAddress(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseUpdatedMuseConnectorTSSAddress) (event.Subscription, error) {

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "UpdatedMuseConnectorTSSAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseUpdatedMuseConnectorTSSAddress)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "UpdatedMuseConnectorTSSAddress", log); err != nil {
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
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseUpdatedMuseConnectorTSSAddress(log types.Log) (*MuseConnectorBaseUpdatedMuseConnectorTSSAddress, error) {
	event := new(MuseConnectorBaseUpdatedMuseConnectorTSSAddress)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "UpdatedMuseConnectorTSSAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MuseConnectorBase contract.
type MuseConnectorBaseUpgradedIterator struct {
	Event *MuseConnectorBaseUpgraded // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseUpgraded)
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
		it.Event = new(MuseConnectorBaseUpgraded)
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
func (it *MuseConnectorBaseUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseUpgraded represents a Upgraded event raised by the MuseConnectorBase contract.
type MuseConnectorBaseUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MuseConnectorBaseUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseUpgradedIterator{contract: _MuseConnectorBase.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseUpgraded)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseUpgraded(log types.Log) (*MuseConnectorBaseUpgraded, error) {
	event := new(MuseConnectorBaseUpgraded)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawnIterator struct {
	Event *MuseConnectorBaseWithdrawn // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseWithdrawn)
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
		it.Event = new(MuseConnectorBaseWithdrawn)
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
func (it *MuseConnectorBaseWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseWithdrawn represents a Withdrawn event raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterWithdrawn(opts *bind.FilterOpts, to []common.Address) (*MuseConnectorBaseWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseWithdrawnIterator{contract: _MuseConnectorBase.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed to, uint256 amount)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "Withdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseWithdrawn)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseWithdrawn(log types.Log) (*MuseConnectorBaseWithdrawn, error) {
	event := new(MuseConnectorBaseWithdrawn)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawnAndCalledIterator struct {
	Event *MuseConnectorBaseWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseWithdrawnAndCalled)
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
		it.Event = new(MuseConnectorBaseWithdrawnAndCalled)
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
func (it *MuseConnectorBaseWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawnAndCalled struct {
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndCalled is a free log retrieval operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, to []common.Address) (*MuseConnectorBaseWithdrawnAndCalledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseWithdrawnAndCalledIterator{contract: _MuseConnectorBase.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0x23b9573b29ff81f01c7aa1968188e1cb7d5858b08582e111fdaf386d9ef9bd8d.
//
// Solidity: event WithdrawnAndCalled(address indexed to, uint256 amount, bytes data)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseWithdrawnAndCalled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "WithdrawnAndCalled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseWithdrawnAndCalled)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseWithdrawnAndCalled(log types.Log) (*MuseConnectorBaseWithdrawnAndCalled, error) {
	event := new(MuseConnectorBaseWithdrawnAndCalled)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseConnectorBaseWithdrawnAndRevertedIterator is returned from FilterWithdrawnAndReverted and is used to iterate over the raw logs and unpacked data for WithdrawnAndReverted events raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawnAndRevertedIterator struct {
	Event *MuseConnectorBaseWithdrawnAndReverted // Event containing the contract specifics and raw log

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
func (it *MuseConnectorBaseWithdrawnAndRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseConnectorBaseWithdrawnAndReverted)
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
		it.Event = new(MuseConnectorBaseWithdrawnAndReverted)
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
func (it *MuseConnectorBaseWithdrawnAndRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseConnectorBaseWithdrawnAndRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseConnectorBaseWithdrawnAndReverted represents a WithdrawnAndReverted event raised by the MuseConnectorBase contract.
type MuseConnectorBaseWithdrawnAndReverted struct {
	To            common.Address
	Amount        *big.Int
	Data          []byte
	RevertContext RevertContext
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnAndReverted is a free log retrieval operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) FilterWithdrawnAndReverted(opts *bind.FilterOpts, to []common.Address) (*MuseConnectorBaseWithdrawnAndRevertedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.FilterLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return &MuseConnectorBaseWithdrawnAndRevertedIterator{contract: _MuseConnectorBase.contract, event: "WithdrawnAndReverted", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndReverted is a free log subscription operation binding the contract event 0x5272d2fee39bff41b2e763562526315906046373ce08a7bacf76c3080d731ff0.
//
// Solidity: event WithdrawnAndReverted(address indexed to, uint256 amount, bytes data, (address,address,uint256,bytes) revertContext)
func (_MuseConnectorBase *MuseConnectorBaseFilterer) WatchWithdrawnAndReverted(opts *bind.WatchOpts, sink chan<- *MuseConnectorBaseWithdrawnAndReverted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseConnectorBase.contract.WatchLogs(opts, "WithdrawnAndReverted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseConnectorBaseWithdrawnAndReverted)
				if err := _MuseConnectorBase.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
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
func (_MuseConnectorBase *MuseConnectorBaseFilterer) ParseWithdrawnAndReverted(log types.Log) (*MuseConnectorBaseWithdrawnAndReverted, error) {
	event := new(MuseConnectorBaseWithdrawnAndReverted)
	if err := _MuseConnectorBase.contract.UnpackLog(event, "WithdrawnAndReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
