// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewaymevmupgradetest

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

// GatewayMEVMUpgradeTestMetaData contains all meta data concerning the GatewayMEVMUpgradeTest contract.
var GatewayMEVMUpgradeTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_MESSAGE_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROTOCOL_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"call\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndCall\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositAndRevert\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"context\",\"type\":\"tuple\",\"internalType\":\"structMessageContext\",\"components\":[{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"senderEVM\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeRevert\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertContext\",\"type\":\"tuple\",\"internalType\":\"structRevertContext\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"asset\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"museToken_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawAndCall\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"museToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedMuseSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMuseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWMUSEOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20BurnFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20DepositFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20TransferFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516144ed6100fd60003960008181612d3d01528181612d660152612f3701526144ed6000f3fe6080604052600436106101e75760003560e01c806352d1902d116101025780639d4ba46511610095578063c39aca3711610064578063c39aca37146106a2578063d547741f146106c2578063e63ab1e9146106e2578063f45346dc1461071657600080fd5b80639d4ba465146105f7578063a217fddf14610617578063ad3cb1cc1461062c578063bcf7f32b1461068257600080fd5b80638456cb59116100d15780638456cb591461054757806391d148541461055c57806397a1cef1146105c157806397d340f5146105e157600080fd5b806352d1902d146104bb5780635c975abb146104d05780637b15118b146105075780637c0dcb5f1461052757600080fd5b80632722feee1161017a5780633b283933116101495780633b283933146104535780633f4ba83a14610473578063485cc955146104885780634f1ef286146104a857600080fd5b80632722feee146103cb5780632810ae63146103f35780632f2ff15d1461041357806336568abe1461043357600080fd5b80631cb5ea75116101b65780631cb5ea75146102f657806321501a951461031657806321e093b114610336578063248a9ca31461036e57600080fd5b806301ffc9a714610261578063048ae42c1461029657806306cb8983146102b6578063184b0793146102d657600080fd5b3661025c576101f4610736565b6000546001600160a01b0316331480159061022357503373735b14bb79463307aacbed86daf3322b1e6226ab14155b1561025a576040517fb3af013700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b600080fd5b34801561026d57600080fd5b5061028161027c366004613480565b610794565b60405190151581526020015b60405180910390f35b3480156102a257600080fd5b5061025a6102b1366004613614565b61082d565b3480156102c257600080fd5b5061025a6102d13660046136e6565b610a49565b3480156102e257600080fd5b5061025a6102f13660046137b6565b610b4d565b34801561030257600080fd5b5061025a610311366004613806565b610c3c565b34801561032257600080fd5b5061025a6103313660046138b4565b610d12565b34801561034257600080fd5b50600054610356906001600160a01b031681565b6040516001600160a01b03909116815260200161028d565b34801561037a57600080fd5b506103bd610389366004613940565b60009081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b60405190815260200161028d565b3480156103d757600080fd5b5061035673735b14bb79463307aacbed86daf3322b1e6226ab81565b3480156103ff57600080fd5b5061025a61040e366004613959565b610ec6565b34801561041f57600080fd5b5061025a61042e3660046139fe565b61106f565b34801561043f57600080fd5b5061025a61044e3660046139fe565b6110b9565b34801561045f57600080fd5b5061025a61046e366004613a2e565b61110a565b34801561047f57600080fd5b5061025a61128c565b34801561049457600080fd5b5061025a6104a3366004613ac1565b6112c1565b61025a6104b6366004613aef565b6114fd565b3480156104c757600080fd5b506103bd61151c565b3480156104dc57600080fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610281565b34801561051357600080fd5b5061025a610522366004613b35565b61154b565b34801561053357600080fd5b5061025a610542366004613ba7565b611711565b34801561055357600080fd5b5061025a611935565b34801561056857600080fd5b506102816105773660046139fe565b60009182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156105cd57600080fd5b5061025a6105dc366004613c2c565b611967565b3480156105ed57600080fd5b506103bd61040081565b34801561060357600080fd5b5061025a610612366004613c90565b611adb565b34801561062357600080fd5b506103bd600081565b34801561063857600080fd5b506106756040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b60405161028d9190613d5e565b34801561068e57600080fd5b5061025a61069d366004613d71565b611d68565b3480156106ae57600080fd5b5061025a6106bd366004613d71565b611e7f565b3480156106ce57600080fd5b5061025a6106dd3660046139fe565b61208f565b3480156106ee57600080fd5b506103bd7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b34801561072257600080fd5b5061025a610731366004613e0f565b6120d3565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1615610792576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061082757507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b6108356122e3565b61083d610736565b865160000361085f5760405163d92e233d60e01b815260040160405180910390fd5b85600003610899576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816000036108d3576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006108e26060830183613e51565b6108ed915085613eb6565b905061040081111561093b576040517fcd6f4e6d0000000000000000000000000000000000000000000000000000000081526004810182905261040060248201526044015b60405180910390fd5b6000610948888886612364565b90506000336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8b8a8c868d6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109dd9190613ef0565b8d8d60405180604001604052808f8152602001600115158152508d604051610a0d99989796959493929190614045565b60405180910390a35050610a4060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50505050505050565b610a516122e3565b610a59610736565b8135600003610a94576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610aa36060830183613e51565b610aae915085613eb6565b9050610400811115610af7576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b610b1b87878787610b0d368990038901896140c7565b610b168861411f565b6126c3565b50610b4560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610b9a576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ba2610736565b6001600160a01b038216610bc95760405163d92e233d60e01b815260040160405180910390fd5b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690610c0e9084906004016141c7565b600060405180830381600087803b158015610c2857600080fd5b505af1158015610b45573d6000803e3d6000fd5b610c446122e3565b610c4c610736565b81600003610c86576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610c956060830183613e51565b610ca0915085613eb6565b9050610400811115610ce9576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b610b1b8787878760405180604001604052808981526020016001151581525087610b169061411f565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610d5f576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610d67610736565b6001600160a01b038316610d8e5760405163d92e233d60e01b815260040160405180910390fd5b83600003610dc8576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480610dfb57506001600160a01b03831630145b15610e32576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610e3c84846128d6565b6000546040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b0380861692635bcfd61692610e8d928a921690899088908890600401614237565b600060405180830381600087803b158015610ea757600080fd5b505af1158015610ebb573d6000803e3d6000fd5b505050505050505050565b610ece6122e3565b610ed6610736565b8651600003610ef85760405163d92e233d60e01b815260040160405180910390fd5b85600003610f32576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8135600003610f6d576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610f7c6060830183613e51565b610f87915085613eb6565b9050610400811115610fd0576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b610fee8773735b14bb79463307aacbed86daf3322b1e6226ab6128d6565b60008054604051889233927f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9261103d928e926001600160a01b0316918e919081908e908e908e908e906142d3565b60405180910390a350610a4060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546110a981612abf565b6110b38383612ac9565b50505050565b6001600160a01b03811633146110fb576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111058282612bb6565b505050565b6111126122e3565b61111a610736565b855160000361113c5760405163d92e233d60e01b815260040160405180910390fd5b84600003611176576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006111856060830183613e51565b611190915084613eb6565b90506104008111156111d9576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b6111f78673735b14bb79463307aacbed86daf3322b1e6226ab6128d6565b60008054604080518082018252838152600160208201529051889333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9361125a938e936001600160a01b03909316928e92909182918e918e91908e90614045565b60405180910390a350610b4560017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6112b681612abf565b6112be612c7a565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff1660008115801561130c5750825b905060008267ffffffffffffffff1660011480156113295750303b155b905081158015611337575080155b1561136e576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016600117855583156113cf5784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff16680100000000000000001785555b6001600160a01b03871615806113ec57506001600160a01b038616155b1561140a5760405163d92e233d60e01b815260040160405180910390fd5b611412612d0a565b61141a612d0a565b611422612d12565b61142a612d22565b611435600087612ac9565b506114607f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a87612ac9565b50600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0389161790558315610a405784547fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b611505612d32565b61150e82612e02565b6115188282612e0d565b5050565b6000611526612f2c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b6115536122e3565b61155b610736565b865160000361157d5760405163d92e233d60e01b815260040160405180910390fd5b856000036115b7576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81356000036115f2576040517f60ee124700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006116016060830183613e51565b61160c915085613eb6565b9050610400811115611655576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b600061166388888635612364565b90506000336001600160a01b03167f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c8b8a8c868d6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116d4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f89190613ef0565b8d8d8d8d604051610a0d999897969594939291906142d3565b6117196122e3565b611721610736565b83516000036117435760405163d92e233d60e01b815260040160405180910390fd5b8260000361177d576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061178c6060830183613e51565b9150506104008111156117d6576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b60006117e28585612f8e565b90506000336001600160a01b03167f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97888789868a6001600160a01b0316634d8943bb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611853573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118779190613ef0565b60405180604001604052808d6001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118e49190613ef0565b81526001602090910152604051611902969594939291908d9061432b565b60405180910390a350506110b360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61195f81612abf565b6112be612ffc565b61196f6122e3565b611977610736565b83516000036119995760405163d92e233d60e01b815260040160405180910390fd5b826000036119d3576040517f19c08f4900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006119e26060830183613e51565b915050610400811115611a2c576040517fcd6f4e6d000000000000000000000000000000000000000000000000000000008152600481018290526104006024820152604401610932565b611a4a8473735b14bb79463307aacbed86daf3322b1e6226ab6128d6565b60008054604080518082018252838152600160208201529051869333937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c93611aa9938c936001600160a01b03909316928c9290918291908c9061432b565b60405180910390a3506110b360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611b28576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b30610736565b6001600160a01b0384161580611b4d57506001600160a01b038216155b15611b6b5760405163d92e233d60e01b815260040160405180910390fd5b82600003611ba5576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03821673735b14bb79463307aacbed86daf3322b1e6226ab1480611bd857506001600160a01b03821630145b15611c0f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038381166004830152602482018590528516906347e7ef24906044016020604051808303816000875af1158015611c77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c9b91906143ad565b611ceb576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b0380861660048301528316602482015260448101849052606401610932565b6040517fc9028a360000000000000000000000000000000000000000000000000000000081526001600160a01b0383169063c9028a3690611d309084906004016141c7565b600060405180830381600087803b158015611d4a57600080fd5b505af1158015611d5e573d6000803e3d6000fd5b5050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611db5576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611dbd610736565b6001600160a01b0385161580611dda57506001600160a01b038316155b15611df85760405163d92e233d60e01b815260040160405180910390fd5b6040517f5bcfd6160000000000000000000000000000000000000000000000000000000081526001600160a01b03841690635bcfd61690611e459089908990899088908890600401614237565b600060405180830381600087803b158015611e5f57600080fd5b505af1158015611e73573d6000803e3d6000fd5b50505050505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14611ecc576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611ed4610736565b6001600160a01b0385161580611ef157506001600160a01b038316155b15611f0f5760405163d92e233d60e01b815260040160405180910390fd5b83600003611f49576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03831673735b14bb79463307aacbed86daf3322b1e6226ab1480611f7c57506001600160a01b03831630145b15611fb3576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af115801561201b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061203f91906143ad565b611df8576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b0380871660048301528416602482015260448101859052606401610932565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546120c981612abf565b6110b38383612bb6565b3373735b14bb79463307aacbed86daf3322b1e6226ab14612120576040517f42c0407e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612128610736565b6001600160a01b038316158061214557506001600160a01b038116155b156121635760405163d92e233d60e01b815260040160405180910390fd5b8160000361219d576040517f5d67094f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160a01b03811673735b14bb79463307aacbed86daf3322b1e6226ab14806121d057506001600160a01b03811630145b15612207576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b038281166004830152602482018490528416906347e7ef24906044016020604051808303816000875af115801561226f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061229391906143ad565b611105576040517f40a143ba0000000000000000000000000000000000000000000000000000000081526001600160a01b0380851660048301528216602482015260448101839052606401610932565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080547ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0161235e576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6000806000846001600160a01b031663fc5fecd5856040518263ffffffff1660e01b815260040161239791815260200190565b6040805180830381865afa1580156123b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123d791906143ca565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af115801561245c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061248091906143ad565b6124e2576040517f667011120000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052606401610932565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018790526001600160a01b038616906323b872dd906064016020604051808303816000875af115801561254e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061257291906143ad565b6125c6576040517f489ca9b70000000000000000000000000000000000000000000000000000000081526001600160a01b038616600482015233602482015230604482015260648101879052608401610932565b6040517f42966c68000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b038616906342966c68906024016020604051808303816000875af1158015612626573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061264a91906143ad565b612692576040517f7112ae770000000000000000000000000000000000000000000000000000000081526001600160a01b038616600482015260248101879052604401610932565b9150505b9392505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b85516000036126e55760405163d92e233d60e01b815260040160405180910390fd5b81516040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600481019190915260009081906001600160a01b0388169063fc5fecd5906024016040805180830381865afa15801561274a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061276e91906143ca565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab60248201526044810182905291935091506001600160a01b038316906323b872dd906064016020604051808303816000875af11580156127f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061281791906143ad565b612879576040517f667011120000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052606401610932565b866001600160a01b0316336001600160a01b03167f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e48a898989896040516128c49594939291906143f8565b60405180910390a35050505050505050565b6000546040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490526001600160a01b03909116906323b872dd906064016020604051808303816000875af1158015612946573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061296a91906143ad565b6129a9576040517ff279af7e00000000000000000000000000000000000000000000000000000000815230600482015260248101839052604401610932565b6000546040517f2e1a7d4d000000000000000000000000000000000000000000000000000000008152600481018490526001600160a01b0390911690632e1a7d4d90602401600060405180830381600087803b158015612a0857600080fd5b505af1158015612a1c573d6000803e3d6000fd5b505050506000816001600160a01b03168360405160006040518083038185875af1925050503d8060008114612a6d576040519150601f19603f3d011682016040523d82523d6000602084013e612a72565b606091505b5050905080611105576040517ff279af7e0000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260248101849052604401610932565b6112be8133613075565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612bac576000848152602082815260408083206001600160a01b0387168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055612b623390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610827565b6000915050610827565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612bac576000848152602082815260408083206001600160a01b038716808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610827565b612c82613102565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b61079261315d565b612d1a61315d565b6107926131c4565b612d2a61315d565b610792613215565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612dcb57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612dbf7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15610792576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061151881612abf565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e85575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252612e8291810190613ef0565b60015b612ec6576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610932565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612f22576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610932565b611105838361321d565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610792576040517fe07c8dba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006126968383846001600160a01b031663091d27886040518163ffffffff1660e01b8152600401602060405180830381865afa158015612fd3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff79190613ef0565b612364565b613004610736565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612cec565b60008281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611518576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401610932565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16610792576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16610792576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6131cc61315d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b61269d61315d565b61322682613273565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561326b57611105828261331b565b611518613391565b806001600160a01b03163b6000036132c2576040517f4c9c8ce30000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610932565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b6060600080846001600160a01b031684604051613338919061449b565b600060405180830381855af49150503d8060008114613373576040519150601f19603f3d011682016040523d82523d6000602084013e613378565b606091505b50915091506133888583836133c9565b95945050505050565b3415610792576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826133de576133d98261343e565b612696565b81511580156133f557506001600160a01b0384163b155b15613437576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610932565b5080612696565b80511561344e5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006020828403121561349257600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461269657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261350257600080fd5b813567ffffffffffffffff81111561351c5761351c6134c2565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff81118282101715613569576135696134c2565b60405281815283820160200185101561358157600080fd5b816020850160208301376000918101602001919091529392505050565b6001600160a01b03811681146112be57600080fd5b60008083601f8401126135c557600080fd5b50813567ffffffffffffffff8111156135dd57600080fd5b6020830191508360208285010111156135f557600080fd5b9250929050565b600060a0828403121561360e57600080fd5b50919050565b600080600080600080600060c0888a03121561362f57600080fd5b873567ffffffffffffffff81111561364657600080fd5b6136528a828b016134f1565b97505060208801359550604088013561366a8161359e565b9450606088013567ffffffffffffffff81111561368657600080fd5b6136928a828b016135b3565b9095509350506080880135915060a088013567ffffffffffffffff8111156136b957600080fd5b6136c58a828b016135fc565b91505092959891949750929550565b60006040828403121561360e57600080fd5b60008060008060008060c087890312156136ff57600080fd5b863567ffffffffffffffff81111561371657600080fd5b61372289828a016134f1565b96505060208701356137338161359e565b9450604087013567ffffffffffffffff81111561374f57600080fd5b61375b89828a016135b3565b909550935061376f905088606089016136d4565b915060a087013567ffffffffffffffff81111561378b57600080fd5b61379789828a016135fc565b9150509295509295509295565b60006080828403121561360e57600080fd5b600080604083850312156137c957600080fd5b82356137d48161359e565b9150602083013567ffffffffffffffff8111156137f057600080fd5b6137fc858286016137a4565b9150509250929050565b60008060008060008060a0878903121561381f57600080fd5b863567ffffffffffffffff81111561383657600080fd5b61384289828a016134f1565b96505060208701356138538161359e565b9450604087013567ffffffffffffffff81111561386f57600080fd5b61387b89828a016135b3565b90955093505060608701359150608087013567ffffffffffffffff81111561378b57600080fd5b60006060828403121561360e57600080fd5b6000806000806000608086880312156138cc57600080fd5b853567ffffffffffffffff8111156138e357600080fd5b6138ef888289016138a2565b9550506020860135935060408601356139078161359e565b9250606086013567ffffffffffffffff81111561392357600080fd5b61392f888289016135b3565b969995985093965092949392505050565b60006020828403121561395257600080fd5b5035919050565b600080600080600080600060e0888a03121561397457600080fd5b873567ffffffffffffffff81111561398b57600080fd5b6139978a828b016134f1565b9750506020880135955060408801359450606088013567ffffffffffffffff8111156139c257600080fd5b6139ce8a828b016135b3565b90955093506139e290508960808a016136d4565b915060c088013567ffffffffffffffff8111156136b957600080fd5b60008060408385031215613a1157600080fd5b823591506020830135613a238161359e565b809150509250929050565b60008060008060008060a08789031215613a4757600080fd5b863567ffffffffffffffff811115613a5e57600080fd5b613a6a89828a016134f1565b9650506020870135945060408701359350606087013567ffffffffffffffff811115613a9557600080fd5b613aa189828a016135b3565b909450925050608087013567ffffffffffffffff81111561378b57600080fd5b60008060408385031215613ad457600080fd5b8235613adf8161359e565b91506020830135613a238161359e565b60008060408385031215613b0257600080fd5b8235613b0d8161359e565b9150602083013567ffffffffffffffff811115613b2957600080fd5b6137fc858286016134f1565b600080600080600080600060e0888a031215613b5057600080fd5b873567ffffffffffffffff811115613b6757600080fd5b613b738a828b016134f1565b975050602088013595506040880135613b8b8161359e565b9450606088013567ffffffffffffffff8111156139c257600080fd5b60008060008060808587031215613bbd57600080fd5b843567ffffffffffffffff811115613bd457600080fd5b613be0878288016134f1565b945050602085013592506040850135613bf88161359e565b9150606085013567ffffffffffffffff811115613c1457600080fd5b613c20878288016135fc565b91505092959194509250565b60008060008060808587031215613c4257600080fd5b843567ffffffffffffffff811115613c5957600080fd5b613c65878288016134f1565b9450506020850135925060408501359150606085013567ffffffffffffffff811115613c1457600080fd5b60008060008060808587031215613ca657600080fd5b8435613cb18161359e565b9350602085013592506040850135613cc88161359e565b9150606085013567ffffffffffffffff811115613ce457600080fd5b613c20878288016137a4565b60005b83811015613d0b578181015183820152602001613cf3565b50506000910152565b60008151808452613d2c816020860160208601613cf0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006126966020830184613d14565b60008060008060008060a08789031215613d8a57600080fd5b863567ffffffffffffffff811115613da157600080fd5b613dad89828a016138a2565b9650506020870135613dbe8161359e565b9450604087013593506060870135613dd58161359e565b9250608087013567ffffffffffffffff811115613df157600080fd5b613dfd89828a016135b3565b979a9699509497509295939492505050565b600080600060608486031215613e2457600080fd5b8335613e2f8161359e565b9250602084013591506040840135613e468161359e565b809150509250925092565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613e8657600080fd5b83018035915067ffffffffffffffff821115613ea157600080fd5b6020019150368190038213156135f557600080fd5b80820180821115610827577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060208284031215613f0257600080fd5b5051919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b80151581146112be57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613f9557600080fd5b830160208101925035905067ffffffffffffffff811115613fb557600080fd5b8036038213156135f557600080fd5b60008135613fd18161359e565b6001600160a01b031683526020820135613fea81613f52565b151560208401526040820135613fff8161359e565b6001600160a01b0316604084015261401a6060830183613f60565b60a0606086015261402f60a086018284613f09565b6080948501359590940194909452509092915050565b6101208152600061405a61012083018c613d14565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a084015261408e818789613f09565b855160c08501526020860151151560e085015290505b8281036101008401526140b78185613fc4565b9c9b505050505050505050505050565b600060408284031280156140da57600080fd5b506040805190810167ffffffffffffffff811182821017156140fe576140fe6134c2565b60405282358152602083013561411381613f52565b60208201529392505050565b600060a0823603121561413157600080fd5b60405160a0810167ffffffffffffffff81118282101715614154576141546134c2565b60405282356141628161359e565b8152602083013561417281613f52565b602082015260408301356141858161359e565b6040820152606083013567ffffffffffffffff8111156141a457600080fd5b6141b0368286016134f1565b606083015250608092830135928101929092525090565b60208152600082356141d88161359e565b6001600160a01b03811660208401525060208301356141f68161359e565b6001600160a01b038116604084015250600060408401359050806060840152506142236060840184613f60565b60808085015261338860a085018284613f09565b6080815260006142478788613f60565b6060608085015261425c60e085018284613f09565b915050602088013561426d8161359e565b6001600160a01b0390811660a085015260408981013560c08601529088166020850152830186905282810360608401526142a8818587613f09565b98975050505050505050565b8035825260208101356142c681613f52565b8015156020840152505050565b610120815260006142e861012083018c613d14565b6001600160a01b038b16602084015289604084015288606084015287608084015282810360a084015261431c818789613f09565b90506140a460c08401866142b4565b6101208152600061434061012083018a613d14565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501526000825261438760c0850187805182526020908101511515910152565b602081016101008501525061439f6020820185613fc4565b9a9950505050505050505050565b6000602082840312156143bf57600080fd5b815161269681613f52565b600080604083850312156143dd57600080fd5b82516143e88161359e565b6020939093015192949293505050565b60a08152600061440b60a0830188613d14565b828103602084015261441e818789613f09565b85516040850152602086015115156060850152905082810360808401526001600160a01b0384511681526020840151151560208201526001600160a01b036040850151166040820152606084015160a0606083015261448060a0830182613d14565b90506080850151608083015280925050509695505050505050565b600082516144ad818460208701613cf0565b919091019291505056fea2646970667358221220e2759414fb451a6141a3224349c40dd37732e5c7d6cb8c2da06fd97f2901452164736f6c634300081a0033",
}

// GatewayMEVMUpgradeTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMEVMUpgradeTestMetaData.ABI instead.
var GatewayMEVMUpgradeTestABI = GatewayMEVMUpgradeTestMetaData.ABI

// GatewayMEVMUpgradeTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMEVMUpgradeTestMetaData.Bin instead.
var GatewayMEVMUpgradeTestBin = GatewayMEVMUpgradeTestMetaData.Bin

// DeployGatewayMEVMUpgradeTest deploys a new Ethereum contract, binding an instance of GatewayMEVMUpgradeTest to it.
func DeployGatewayMEVMUpgradeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayMEVMUpgradeTest, error) {
	parsed, err := GatewayMEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayMEVMUpgradeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayMEVMUpgradeTest{GatewayMEVMUpgradeTestCaller: GatewayMEVMUpgradeTestCaller{contract: contract}, GatewayMEVMUpgradeTestTransactor: GatewayMEVMUpgradeTestTransactor{contract: contract}, GatewayMEVMUpgradeTestFilterer: GatewayMEVMUpgradeTestFilterer{contract: contract}}, nil
}

// GatewayMEVMUpgradeTest is an auto generated Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTest struct {
	GatewayMEVMUpgradeTestCaller     // Read-only binding to the contract
	GatewayMEVMUpgradeTestTransactor // Write-only binding to the contract
	GatewayMEVMUpgradeTestFilterer   // Log filterer for contract events
}

// GatewayMEVMUpgradeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMUpgradeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMUpgradeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayMEVMUpgradeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMUpgradeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayMEVMUpgradeTestSession struct {
	Contract     *GatewayMEVMUpgradeTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayMEVMUpgradeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayMEVMUpgradeTestCallerSession struct {
	Contract *GatewayMEVMUpgradeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayMEVMUpgradeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayMEVMUpgradeTestTransactorSession struct {
	Contract     *GatewayMEVMUpgradeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayMEVMUpgradeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTestRaw struct {
	Contract *GatewayMEVMUpgradeTest // Generic contract binding to access the raw methods on
}

// GatewayMEVMUpgradeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTestCallerRaw struct {
	Contract *GatewayMEVMUpgradeTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayMEVMUpgradeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayMEVMUpgradeTestTransactorRaw struct {
	Contract *GatewayMEVMUpgradeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayMEVMUpgradeTest creates a new instance of GatewayMEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayMEVMUpgradeTest(address common.Address, backend bind.ContractBackend) (*GatewayMEVMUpgradeTest, error) {
	contract, err := bindGatewayMEVMUpgradeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTest{GatewayMEVMUpgradeTestCaller: GatewayMEVMUpgradeTestCaller{contract: contract}, GatewayMEVMUpgradeTestTransactor: GatewayMEVMUpgradeTestTransactor{contract: contract}, GatewayMEVMUpgradeTestFilterer: GatewayMEVMUpgradeTestFilterer{contract: contract}}, nil
}

// NewGatewayMEVMUpgradeTestCaller creates a new read-only instance of GatewayMEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayMEVMUpgradeTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayMEVMUpgradeTestCaller, error) {
	contract, err := bindGatewayMEVMUpgradeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestCaller{contract: contract}, nil
}

// NewGatewayMEVMUpgradeTestTransactor creates a new write-only instance of GatewayMEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayMEVMUpgradeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayMEVMUpgradeTestTransactor, error) {
	contract, err := bindGatewayMEVMUpgradeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestTransactor{contract: contract}, nil
}

// NewGatewayMEVMUpgradeTestFilterer creates a new log filterer instance of GatewayMEVMUpgradeTest, bound to a specific deployed contract.
func NewGatewayMEVMUpgradeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayMEVMUpgradeTestFilterer, error) {
	contract, err := bindGatewayMEVMUpgradeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestFilterer{contract: contract}, nil
}

// bindGatewayMEVMUpgradeTest binds a generic wrapper to an already deployed contract.
func bindGatewayMEVMUpgradeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMEVMUpgradeTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMEVMUpgradeTest.Contract.GatewayMEVMUpgradeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.GatewayMEVMUpgradeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.GatewayMEVMUpgradeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMEVMUpgradeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.DEFAULTADMINROLE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) MAXMESSAGESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "MAX_MESSAGE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayMEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// MAXMESSAGESIZE is a free data retrieval call binding the contract method 0x97d340f5.
//
// Solidity: function MAX_MESSAGE_SIZE() view returns(uint256)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) MAXMESSAGESIZE() (*big.Int, error) {
	return _GatewayMEVMUpgradeTest.Contract.MAXMESSAGESIZE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) PAUSERROLE() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.PAUSERROLE(&_GatewayMEVMUpgradeTest.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) PROTOCOLADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "PROTOCOL_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayMEVMUpgradeTest.Contract.PROTOCOLADDRESS(&_GatewayMEVMUpgradeTest.CallOpts)
}

// PROTOCOLADDRESS is a free data retrieval call binding the contract method 0x2722feee.
//
// Solidity: function PROTOCOL_ADDRESS() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) PROTOCOLADDRESS() (common.Address, error) {
	return _GatewayMEVMUpgradeTest.Contract.PROTOCOLADDRESS(&_GatewayMEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayMEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayMEVMUpgradeTest.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _GatewayMEVMUpgradeTest.Contract.UPGRADEINTERFACEVERSION(&_GatewayMEVMUpgradeTest.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayMEVMUpgradeTest.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.GetRoleAdmin(&_GatewayMEVMUpgradeTest.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.HasRole(&_GatewayMEVMUpgradeTest.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.HasRole(&_GatewayMEVMUpgradeTest.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Paused() (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.Paused(&_GatewayMEVMUpgradeTest.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) Paused() (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.Paused(&_GatewayMEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayMEVMUpgradeTest.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GatewayMEVMUpgradeTest.Contract.ProxiableUUID(&_GatewayMEVMUpgradeTest.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.SupportsInterface(&_GatewayMEVMUpgradeTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GatewayMEVMUpgradeTest.Contract.SupportsInterface(&_GatewayMEVMUpgradeTest.CallOpts, interfaceId)
}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCaller) MuseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMUpgradeTest.contract.Call(opts, &out, "museToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) MuseToken() (common.Address, error) {
	return _GatewayMEVMUpgradeTest.Contract.MuseToken(&_GatewayMEVMUpgradeTest.CallOpts)
}

// MuseToken is a free data retrieval call binding the contract method 0x21e093b1.
//
// Solidity: function museToken() view returns(address)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestCallerSession) MuseToken() (common.Address, error) {
	return _GatewayMEVMUpgradeTest.Contract.MuseToken(&_GatewayMEVMUpgradeTest.CallOpts)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Call(opts *bind.TransactOpts, receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "call", receiver, mrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Call(receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Call(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, mrc20, message, callOptions, revertOptions)
}

// Call is a paid mutator transaction binding the contract method 0x06cb8983.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Call(receiver []byte, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Call(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, mrc20, message, callOptions, revertOptions)
}

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Call0(opts *bind.TransactOpts, receiver []byte, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "call0", receiver, mrc20, message, gasLimit, revertOptions)
}

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Call0(receiver []byte, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Call0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, mrc20, message, gasLimit, revertOptions)
}

// Call0 is a paid mutator transaction binding the contract method 0x1cb5ea75.
//
// Solidity: function call(bytes receiver, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Call0(receiver []byte, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Call0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, mrc20, message, gasLimit, revertOptions)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Deposit(opts *bind.TransactOpts, mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "deposit", mrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Deposit(mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Deposit(&_GatewayMEVMUpgradeTest.TransactOpts, mrc20, amount, target)
}

// Deposit is a paid mutator transaction binding the contract method 0xf45346dc.
//
// Solidity: function deposit(address mrc20, uint256 amount, address target) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Deposit(mrc20 common.Address, amount *big.Int, target common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Deposit(&_GatewayMEVMUpgradeTest.TransactOpts, mrc20, amount, target)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) DepositAndCall(opts *bind.TransactOpts, context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "depositAndCall", context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, context, amount, target, message)
}

// DepositAndCall is a paid mutator transaction binding the contract method 0x21501a95.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) DepositAndCall(context MessageContext, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, context, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) DepositAndCall0(opts *bind.TransactOpts, context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "depositAndCall0", context, mrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) DepositAndCall0(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayMEVMUpgradeTest.TransactOpts, context, mrc20, amount, target, message)
}

// DepositAndCall0 is a paid mutator transaction binding the contract method 0xc39aca37.
//
// Solidity: function depositAndCall((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) DepositAndCall0(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndCall0(&_GatewayMEVMUpgradeTest.TransactOpts, context, mrc20, amount, target, message)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) DepositAndRevert(opts *bind.TransactOpts, mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "depositAndRevert", mrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) DepositAndRevert(mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayMEVMUpgradeTest.TransactOpts, mrc20, amount, target, revertContext)
}

// DepositAndRevert is a paid mutator transaction binding the contract method 0x9d4ba465.
//
// Solidity: function depositAndRevert(address mrc20, uint256 amount, address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) DepositAndRevert(mrc20 common.Address, amount *big.Int, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.DepositAndRevert(&_GatewayMEVMUpgradeTest.TransactOpts, mrc20, amount, target, revertContext)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Execute(opts *bind.TransactOpts, context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "execute", context, mrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Execute(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Execute(&_GatewayMEVMUpgradeTest.TransactOpts, context, mrc20, amount, target, message)
}

// Execute is a paid mutator transaction binding the contract method 0xbcf7f32b.
//
// Solidity: function execute((bytes,address,uint256) context, address mrc20, uint256 amount, address target, bytes message) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Execute(context MessageContext, mrc20 common.Address, amount *big.Int, target common.Address, message []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Execute(&_GatewayMEVMUpgradeTest.TransactOpts, context, mrc20, amount, target, message)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) ExecuteRevert(opts *bind.TransactOpts, target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "executeRevert", target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayMEVMUpgradeTest.TransactOpts, target, revertContext)
}

// ExecuteRevert is a paid mutator transaction binding the contract method 0x184b0793.
//
// Solidity: function executeRevert(address target, (address,address,uint256,bytes) revertContext) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) ExecuteRevert(target common.Address, revertContext RevertContext) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.ExecuteRevert(&_GatewayMEVMUpgradeTest.TransactOpts, target, revertContext)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.GrantRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.GrantRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address museToken_, address admin_) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Initialize(opts *bind.TransactOpts, museToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "initialize", museToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address museToken_, address admin_) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Initialize(museToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Initialize(&_GatewayMEVMUpgradeTest.TransactOpts, museToken_, admin_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address museToken_, address admin_) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Initialize(museToken_ common.Address, admin_ common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Initialize(&_GatewayMEVMUpgradeTest.TransactOpts, museToken_, admin_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Pause() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Pause(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Pause() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Pause(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.RenounceRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.RenounceRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.RevokeRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.RevokeRole(&_GatewayMEVMUpgradeTest.TransactOpts, role, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Unpause() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Unpause(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Unpause() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Unpause(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.UpgradeToAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Withdraw(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdraw", receiver, amount, mrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Withdraw(receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Withdraw(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, revertOptions)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7c0dcb5f.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, address mrc20, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Withdraw(receiver []byte, amount *big.Int, mrc20 common.Address, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Withdraw(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Withdraw0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdraw0", receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Withdraw0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x97a1cef1.
//
// Solidity: function withdraw(bytes receiver, uint256 amount, uint256 chainId, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Withdraw0(receiver []byte, amount *big.Int, chainId *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Withdraw0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) WithdrawAndCall(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall", receiver, amount, mrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) WithdrawAndCall(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall is a paid mutator transaction binding the contract method 0x048ae42c.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, uint256 gasLimit, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) WithdrawAndCall(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, gasLimit *big.Int, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, message, gasLimit, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) WithdrawAndCall0(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall0", receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall0 is a paid mutator transaction binding the contract method 0x2810ae63.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) WithdrawAndCall0(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall0(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, callOptions, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) WithdrawAndCall1(opts *bind.TransactOpts, receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall1", receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) WithdrawAndCall1(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall1 is a paid mutator transaction binding the contract method 0x3b283933.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, uint256 chainId, bytes message, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) WithdrawAndCall1(receiver []byte, amount *big.Int, chainId *big.Int, message []byte, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall1(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, chainId, message, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) WithdrawAndCall2(opts *bind.TransactOpts, receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.Transact(opts, "withdrawAndCall2", receiver, amount, mrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) WithdrawAndCall2(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, message, callOptions, revertOptions)
}

// WithdrawAndCall2 is a paid mutator transaction binding the contract method 0x7b15118b.
//
// Solidity: function withdrawAndCall(bytes receiver, uint256 amount, address mrc20, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions) returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) WithdrawAndCall2(receiver []byte, amount *big.Int, mrc20 common.Address, message []byte, callOptions CallOptions, revertOptions RevertOptions) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.WithdrawAndCall2(&_GatewayMEVMUpgradeTest.TransactOpts, receiver, amount, mrc20, message, callOptions, revertOptions)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestSession) Receive() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Receive(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestTransactorSession) Receive() (*types.Transaction, error) {
	return _GatewayMEVMUpgradeTest.Contract.Receive(&_GatewayMEVMUpgradeTest.TransactOpts)
}

// GatewayMEVMUpgradeTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestCalledIterator struct {
	Event *GatewayMEVMUpgradeTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestCalled)
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
		it.Event = new(GatewayMEVMUpgradeTestCalled)
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
func (it *GatewayMEVMUpgradeTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestCalled represents a Called event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestCalled struct {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, mrc20 []common.Address) (*GatewayMEVMUpgradeTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestCalledIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestCalled, sender []common.Address, mrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestCalled)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseCalled(log types.Log) (*GatewayMEVMUpgradeTestCalled, error) {
	event := new(GatewayMEVMUpgradeTestCalled)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestInitializedIterator struct {
	Event *GatewayMEVMUpgradeTestInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestInitialized)
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
		it.Event = new(GatewayMEVMUpgradeTestInitialized)
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
func (it *GatewayMEVMUpgradeTestInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestInitialized represents a Initialized event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayMEVMUpgradeTestInitializedIterator, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestInitializedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestInitialized)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseInitialized(log types.Log) (*GatewayMEVMUpgradeTestInitialized, error) {
	event := new(GatewayMEVMUpgradeTestInitialized)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestPausedIterator struct {
	Event *GatewayMEVMUpgradeTestPaused // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestPaused)
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
		it.Event = new(GatewayMEVMUpgradeTestPaused)
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
func (it *GatewayMEVMUpgradeTestPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestPaused represents a Paused event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayMEVMUpgradeTestPausedIterator, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestPausedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestPaused) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestPaused)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParsePaused(log types.Log) (*GatewayMEVMUpgradeTestPaused, error) {
	event := new(GatewayMEVMUpgradeTestPaused)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleAdminChangedIterator struct {
	Event *GatewayMEVMUpgradeTestRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestRoleAdminChanged)
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
		it.Event = new(GatewayMEVMUpgradeTestRoleAdminChanged)
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
func (it *GatewayMEVMUpgradeTestRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestRoleAdminChanged represents a RoleAdminChanged event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*GatewayMEVMUpgradeTestRoleAdminChangedIterator, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestRoleAdminChangedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestRoleAdminChanged)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseRoleAdminChanged(log types.Log) (*GatewayMEVMUpgradeTestRoleAdminChanged, error) {
	event := new(GatewayMEVMUpgradeTestRoleAdminChanged)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleGrantedIterator struct {
	Event *GatewayMEVMUpgradeTestRoleGranted // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestRoleGranted)
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
		it.Event = new(GatewayMEVMUpgradeTestRoleGranted)
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
func (it *GatewayMEVMUpgradeTestRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestRoleGranted represents a RoleGranted event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayMEVMUpgradeTestRoleGrantedIterator, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestRoleGrantedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestRoleGranted)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseRoleGranted(log types.Log) (*GatewayMEVMUpgradeTestRoleGranted, error) {
	event := new(GatewayMEVMUpgradeTestRoleGranted)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleRevokedIterator struct {
	Event *GatewayMEVMUpgradeTestRoleRevoked // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestRoleRevoked)
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
		it.Event = new(GatewayMEVMUpgradeTestRoleRevoked)
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
func (it *GatewayMEVMUpgradeTestRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestRoleRevoked represents a RoleRevoked event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*GatewayMEVMUpgradeTestRoleRevokedIterator, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestRoleRevokedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestRoleRevoked)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseRoleRevoked(log types.Log) (*GatewayMEVMUpgradeTestRoleRevoked, error) {
	event := new(GatewayMEVMUpgradeTestRoleRevoked)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestUnpausedIterator struct {
	Event *GatewayMEVMUpgradeTestUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestUnpaused)
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
		it.Event = new(GatewayMEVMUpgradeTestUnpaused)
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
func (it *GatewayMEVMUpgradeTestUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestUnpaused represents a Unpaused event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayMEVMUpgradeTestUnpausedIterator, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestUnpausedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestUnpaused) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestUnpaused)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseUnpaused(log types.Log) (*GatewayMEVMUpgradeTestUnpaused, error) {
	event := new(GatewayMEVMUpgradeTestUnpaused)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestUpgradedIterator struct {
	Event *GatewayMEVMUpgradeTestUpgraded // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestUpgraded)
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
		it.Event = new(GatewayMEVMUpgradeTestUpgraded)
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
func (it *GatewayMEVMUpgradeTestUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestUpgraded represents a Upgraded event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GatewayMEVMUpgradeTestUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestUpgradedIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestUpgraded)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseUpgraded(log types.Log) (*GatewayMEVMUpgradeTestUpgraded, error) {
	event := new(GatewayMEVMUpgradeTestUpgraded)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawnIterator struct {
	Event *GatewayMEVMUpgradeTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestWithdrawn)
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
		it.Event = new(GatewayMEVMUpgradeTestWithdrawn)
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
func (it *GatewayMEVMUpgradeTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestWithdrawn represents a Withdrawn event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawn struct {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMUpgradeTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestWithdrawnIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestWithdrawn)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseWithdrawn(log types.Log) (*GatewayMEVMUpgradeTestWithdrawn, error) {
	event := new(GatewayMEVMUpgradeTestWithdrawn)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawnAndCalledIterator struct {
	Event *GatewayMEVMUpgradeTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestWithdrawnAndCalled)
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
		it.Event = new(GatewayMEVMUpgradeTestWithdrawnAndCalled)
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
func (it *GatewayMEVMUpgradeTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawnAndCalled struct {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMUpgradeTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestWithdrawnAndCalledIterator{contract: _GatewayMEVMUpgradeTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestWithdrawnAndCalled)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayMEVMUpgradeTestWithdrawnAndCalled, error) {
	event := new(GatewayMEVMUpgradeTestWithdrawnAndCalled)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMUpgradeTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawnV2Iterator struct {
	Event *GatewayMEVMUpgradeTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMUpgradeTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMUpgradeTestWithdrawnV2)
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
		it.Event = new(GatewayMEVMUpgradeTestWithdrawnV2)
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
func (it *GatewayMEVMUpgradeTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMUpgradeTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMUpgradeTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayMEVMUpgradeTest contract.
type GatewayMEVMUpgradeTestWithdrawnV2 struct {
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

// FilterWithdrawnV2 is a free log retrieval operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMUpgradeTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMUpgradeTestWithdrawnV2Iterator{contract: _GatewayMEVMUpgradeTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayMEVMUpgradeTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMUpgradeTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMUpgradeTestWithdrawnV2)
				if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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

// ParseWithdrawnV2 is a log parse operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMUpgradeTest *GatewayMEVMUpgradeTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayMEVMUpgradeTestWithdrawnV2, error) {
	event := new(GatewayMEVMUpgradeTestWithdrawnV2)
	if err := _GatewayMEVMUpgradeTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
