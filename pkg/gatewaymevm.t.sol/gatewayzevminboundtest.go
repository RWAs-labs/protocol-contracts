// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewaymevm

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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// GatewayMEVMInboundTestMetaData contains all meta data concerning the GatewayMEVMInboundTest contract.
var GatewayMEVMInboundTestMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"testCallWithCallOpts\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testCallWithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testUpgradeAndWithdrawMRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20FailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20FailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSE\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEFailsIfAmountIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEFailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEFailsIfNoBalance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEFailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEWithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIfMessageSizeExceeded\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIfNoBalanceForGasFee\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIfNoBalanceForTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIfReceiverIsZeroAddress\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20FailsIsAmountIs0\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20WithCallOptsWithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20WithMessage\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20WithMessageFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Called\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnAndCalled\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawnV2\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"mrc20\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasfee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"protocolFlatFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"message\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"callOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structCallOptions\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isArbitraryCall\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"revertOptions\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structRevertOptions\",\"components\":[{\"name\":\"revertAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callOnRevert\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"abortAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"revertMessage\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"onRevertGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedMuseSent\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"GasFeeTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMRC20Amount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientMuseAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LowBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MessageSizeExceeded\",\"inputs\":[{\"name\":\"provided\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximum\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OnlyWMUSEOrProtocol\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MUSENotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MRC20BurnFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20DepositFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MRC20TransferFailed\",\"inputs\":[{\"name\":\"mrc20\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
	Bin: "0x6080604052600c8054600160ff199182168117909255601f80549091169091179055348015602c57600080fd5b5062011315806200003e6000396000f3fe608060405234801561001057600080fd5b50600436106103415760003560e01c80636dfcbc50116101bd578063c20049f4116100f9578063e51c6388116100a2578063eb7a2fac1161007c578063eb7a2fac1461050a578063fa7626d414610512578063fbc611c81461051f578063fdad0ad01461052757600080fd5b8063e51c6388146104fa578063e804a40614610378578063ea37902f1461050257600080fd5b8063dc749dd7116100d3578063dc749dd7146104e2578063dde7e967146104ea578063e20c9f71146104f257600080fd5b8063c20049f4146104ca578063c946d7c0146104d2578063ceccfab3146104da57600080fd5b8063b0464fdc11610166578063ba414fa611610140578063ba414fa61461049a578063ba800c91146104b2578063ba9adeef146104ba578063bed3e813146104c257600080fd5b8063b0464fdc14610482578063b51ac0711461048a578063b5508aa91461049257600080fd5b806385226c811161019757806385226c8114610458578063916a17c61461046d578063a721b2d3146103f357600080fd5b80636dfcbc50146104405780637ae697301461044857806383ababa91461045057600080fd5b806336431b3f1161028c5780635d72228f116102355780636221b5091161020f5780636221b5091461041357806364002a1f1461041b57806366d9a9a0146104235780636d6ce0d01461043857600080fd5b80635d72228f146103fb5780635efe72a9146104035780636198fb191461040b57600080fd5b806342752d411161026657806342752d41146103e357806343181437146103eb5780634ffab9de146103f357600080fd5b806336431b3f146103cb5780633e5e3c23146103d35780633f7286f4146103db57600080fd5b80631e63d2b9116102ee57806321aeb18c116102c857806321aeb18c146103a6578063264b524c146103ae5780632ade3880146103b657600080fd5b80631e63d2b9146103785780631ed7831c1461038057806320dee15f1461039e57600080fd5b80631238212c1161031f5780631238212c1461036057806314759766146103685780631b9641bf1461037057600080fd5b806305b9f046146103465780630a9254e4146103505780630b5ad28d14610358575b600080fd5b61034e61052f565b005b61034e6106d9565b61034e61125d565b61034e6113c5565b61034e61177e565b61034e611995565b61034e611af7565b610388611efe565b604051610395919061b1e4565b60405180910390f35b61034e611f60565b61034e61234e565b61034e61246c565b6103be612772565b604051610395919061b280565b61034e6128b4565b6103886129f5565b610388612a55565b61034e612ab5565b61034e612c1e565b61034e612d3c565b61034e613142565b61034e613494565b61034e613537565b61034e6136e8565b61034e6138d6565b61042b613a9d565b604051610395919061b3e6565b61034e613c0a565b61034e613efa565b61034e6142e4565b61034e6144b2565b6104606145c4565b604051610395919061b484565b610475614694565b604051610395919061b4fb565b61047561477a565b61034e614860565b61046061494a565b6104a2614a1a565b6040519015158152602001610395565b61034e614aee565b61034e614c04565b61034e614d77565b61034e614edb565b61034e61505d565b61034e6152e1565b61034e61572b565b61034e6157ce565b610388615c0b565b61034e615c6b565b61034e6160f5565b61034e6162b4565b601f546104a29060ff1681565b61034e61639e565b61034e61672b565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529190526020810180516001600160e01b0316630427d73b60e51b17905290506105856001620186a061b5c1565b602c55604051630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156105f557600080fd5b505af1158015610609573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506306cb898391506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106a4926001600160a01b03909116908690602c9060289060040161b710565b600060405180830381600087803b1580156106be57600080fd5b505af11580156106d2573d6000803e3d6000fd5b5050505050565b602580547fffffffffffffffffffffffff000000000000000000000000000000000000000090811630179091556026805490911661123417905560405161071f9061b0f3565b604051809103906000f08015801561073b573d6000803e3d6000fd5b50602280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155604080518082018252600f81527f476174657761795a45564d2e736f6c00000000000000000000000000000000006020820152602554915160248101939093529216604482015261080a919060640160408051601f198184030181529190526020810180516001600160e01b03167f485cc95500000000000000000000000000000000000000000000000000000000179052616a16565b601f80547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b039384168102919091179182905560208054919092049092167fffffffffffffffffffffffff000000000000000000000000000000000000000090921682178155604080517f2722feee0000000000000000000000000000000000000000000000000000000081529051632722feee926004808401939192918290030181865afa1580156108cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f0919061b796565b602780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03929092169190911790556040516109349061b101565b604051809103906000f080158015610950573d6000803e3d6000fd5b50602480547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283161781556027546040517f06447d5600000000000000000000000000000000000000000000000000000000815292166004830152737109709ecfa91a80626ff3989d68f67f5b1dd12d916306447d569101600060405180830381600087803b1580156109ec57600080fd5b505af1158015610a00573d6000803e3d6000fd5b505050506000806000604051610a159061b10f565b6001600160a01b03938416815291831660208301529091166040820152606001604051809103906000f080158015610a51573d6000803e3d6000fd5b50602380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03928316908117909155602054604051601293600193849360009391921690610aa79061b11d565b610ab69695949392919061b7b1565b604051809103906000f080158015610ad2573d6000803e3d6000fd5b50602180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b039283169081179091556023546040517fee2815ba0000000000000000000000000000000000000000000000000000000081526001600482015260248101929092529091169063ee2815ba90604401600060405180830381600087803b158015610b6957600080fd5b505af1158015610b7d573d6000803e3d6000fd5b50506023546040517fa7cb050700000000000000000000000000000000000000000000000000000000815260016004820181905260248201526001600160a01b03909116925063a7cb05079150604401600060405180830381600087803b158015610be757600080fd5b505af1158015610bfb573d6000803e3d6000fd5b50506027546040517fc88a5e6d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152633b9aca006024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d925063c88a5e6d9150604401600060405180830381600087803b158015610c7b57600080fd5b505af1158015610c8f573d6000803e3d6000fd5b50505050602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015610ce457600080fd5b505af1158015610cf8573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015610d6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d90919061b8a6565b506021546025546040517f47e7ef240000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116906347e7ef24906044016020604051808303816000875af1158015610e02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e26919061b8a6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610e8557600080fd5b505af1158015610e99573d6000803e3d6000fd5b50506025546040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d92506306447d569150602401600060405180830381600087803b158015610f0f57600080fd5b505af1158015610f23573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0391821660048201526305f5e10060248201529116925063095ea7b391506044016020604051808303816000875af1158015610f98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbc919061b8a6565b50602260009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0600a6040518263ffffffff1660e01b81526004016000604051808303818588803b15801561100e57600080fd5b505af1158015611022573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600a60248201529116935063095ea7b3925060440190506020604051808303816000875af1158015611096573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ba919061b8a6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561111957600080fd5b505af115801561112d573d6000803e3d6000fd5b50506040805160a08101825261032180825260016020808401918252838501928352845190810190945260008085526060840185905260808401528251602880549251151574010000000000000000000000000000000000000000027fffffffffffffffffffffff0000000000000000000000000000000000000000009093166001600160a01b0392831617929092178255915160298054919093167fffffffffffffffffffffffff000000000000000000000000000000000000000091909116179091559093509150602a90611204908261b93e565b506080919091015160039091015560408051808201909152620186a080825260016020909201829052602c55602d80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790526000602c5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561131557600080fd5b505af1158015611329573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b91506034015b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106a4926001916001600160a01b0316908790602c9060289060040161b9fd565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015611416573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143a919061ba6f565b6025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b1580156114af57600080fd5b505af11580156114c3573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015611535573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611559919061b8a6565b506026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b15801561160757600080fd5b505af115801561161b573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084526000602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526116c793919289926001600160a01b039091169188919060289060040161ba88565b600060405180830381600087803b1580156116e157600080fd5b505af11580156116f5573d6000803e3d6000fd5b50506021546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611748573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176c919061ba6f565b90506117788382616a35565b50505050565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b179052602554602654925193945060019384936001600160a01b03928316937f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9361182b939216910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815290829052602254611863926001600160a01b0390911690899060009081908b90602c9060289061badc565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024015b600060405180830381600087803b1580156118d857600080fd5b505af11580156118ec573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052858486602c60286040518763ffffffff1660e01b81526004016119609695949392919061bb50565b60006040518083038186803b15801561197857600080fd5b505afa15801561198c573d6000803e3d6000fd5b50505050505050565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015611a4857600080fd5b505af1158015611a5c573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526106a4926000916001600160a01b0316908790602c9060289060040161b9fd565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015611b48573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6c919061ba6f565b6026546040516001600160a01b03909116602482015290915060009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150620186a0908190737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b158015611c4b57600080fd5b505af1158015611c5f573d6000803e3d6000fd5b505060255460265460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507fd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b03909216918b9189918491634d8943bb916004808201926020929091908290030181865afa158015611d31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d55919061ba6f565b6040805180820182528a8152600160208201529051611d7d9695949392918d9160289061bb81565b60405180910390a3602080546026546040516001600160a01b0392831693637b15118b93611dc59316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f1981840301815260215483830183528684526001602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152611e319391928b926001600160a01b03909116918a919060289060040161ba88565b600060405180830381600087803b158015611e4b57600080fd5b505af1158015611e5f573d6000803e3d6000fd5b50506021546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015611eb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ed6919061ba6f565b9050611ef683611ee6888861b5c1565b611ef0919061b5c1565b82616a35565b505050505050565b60606016805480602002602001604051908101604052809291908181526020018280548015611f5657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611f38575b5050505050905090565b601f54604080518082018252601a81527f476174657761795a45564d55706772616465546573742e736f6c000000000000602080830191909152825190810190925260008252602554611fc5936001600160a01b036101009091048116939116616a8d565b601f546021546025546040516370a0823160e01b81526001600160a01b03918216600482015261010090930481169260019260009216906370a0823190602401602060405180830381865afa158015612022573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612046919061ba6f565b6040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b0385166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156120d357600080fd5b505af11580156120e7573d6000803e3d6000fd5b505060255460265460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d979060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa1580156121ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121de919061ba6f565b604080518082018252600081526001602082015290516122069695949392919060289061bbe3565b60405180910390a360265460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b03841690637c0dcb5f9060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16835261229a9287916001600160a01b03169060289060040161bc65565b600060405180830381600087803b1580156122b457600080fd5b505af11580156122c8573d6000803e3d6000fd5b50506021546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561231b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061233f919061ba6f565b9050611778611ef0848461b5c1565b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b1580156123ba57600080fd5b505af11580156123ce573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef191506034015b6040516020818303038152906040526000600160286040518563ffffffff1660e01b8152600401612440949392919061bc9f565b60006040518083038186803b15801561245857600080fd5b505afa158015611778573d6000803e3d6000fd5b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa1580156124ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124ee919061ba6f565b6124f990600161bcd0565b67ffffffffffffffff8111156125115761251161b8c8565b6040519080825280601f01601f19166020018201604052801561253b576020820181803683370190505b50602a90612549908261b93e565b5060006028600201805461255c9061b5d4565b905090506000602060009054906101000a90046001600160a01b03166001600160a01b03166397d340f56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156125b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125d9919061ba6f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916126809160040161bce3565b600060405180830381600087803b15801561269a57600080fd5b505af11580156126ae573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352612744926002916001600160a01b03169060289060040161bc65565b600060405180830381600087803b15801561275e57600080fd5b505af1158015611ef6573d6000803e3d6000fd5b6060601e805480602002602001604051908101604052809291908181526020016000905b828210156128ab57600084815260208082206040805180820182526002870290920180546001600160a01b03168352600181018054835181870281018701909452808452939591948681019491929084015b828210156128945783829060005260206000200180546128079061b5d4565b80601f01602080910402602001604051908101604052809291908181526020018280546128339061b5d4565b80156128805780601f1061285557610100808354040283529160200191612880565b820191906000526020600020905b81548152906001019060200180831161286357829003601f168201915b5050505050815260200190600101906127e8565b505050508152505081526020019060010190612796565b50505050905090565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561296757600080fd5b505af115801561297b573d6000803e3d6000fd5b505060208054604080516000808252602154606083018452620186a09583019586528284019190915291517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b95506106a49491936001931691889160289060040161ba88565b60606018805480602002602001604051908101604052809291908181526020018280548015611f56576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611f38575050505050905090565b60606017805480602002602001604051908101604052809291908181526020018280548015611f56576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611f38575050505050905090565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612b6857600080fd5b505af1158015612b7c573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae6391506034016040516020818303038152906040526000600185602c60286040518763ffffffff1660e01b8152600401612bf29695949392919061bb50565b60006040518083038186803b158015612c0a57600080fd5b505afa1580156106d2573d6000803e3d6000fd5b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015612cd157600080fd5b505af1158015612ce5573d6000803e3d6000fd5b5050602080546040516001600160a01b039091169350632810ae639250015b60405160208183030381529060405260018085602c60286040518763ffffffff1660e01b8152600401612bf29695949392919061bb50565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015612d9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dc3919061ba6f565b612dcd919061bcf6565b67ffffffffffffffff811115612de557612de561b8c8565b6040519080825280601f01601f191660200182016040528015612e0f576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015612e74573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e98919061ba6f565b612ea2919061bcf6565b612ead90600161bcd0565b67ffffffffffffffff811115612ec557612ec561b8c8565b6040519080825280601f01601f191660200182016040528015612eef576020820181803683370190505b50602a90612efd908261b93e565b50600060286002018054612f109061b5d4565b8351612f1c925061bcd0565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b03909216926397d340f5926004808401938290030181865afa158015612f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fa4919061ba6f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161304b9160040161bce3565b600060405180830381600087803b15801561306557600080fd5b505af1158015613079573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352613114926001916001600160a01b0316908990602c9060289060040161b9fd565b600060405180830381600087803b15801561312e57600080fd5b505af115801561198c573d6000803e3d6000fd5b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015613193573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131b7919061ba6f565b6025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561322c57600080fd5b505af1158015613240573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af11580156132b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132d6919061b8a6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561333557600080fd5b505af1158015613349573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526133de9287916001600160a01b03169060289060040161bc65565b600060405180830381600087803b1580156133f857600080fd5b505af115801561340c573d6000803e3d6000fd5b50506021546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561345f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613483919061ba6f565b905061348f8282616a35565b505050565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b1790526000602c5551630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016105db565b6026546040516001600160a01b03909116602482015260009060440160408051601f19818403018152918152602080830180516001600160e01b0316630427d73b60e51b1790525490517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b15801561360c57600080fd5b505af1158015613620573d6000803e3d6000fd5b505060215460255460265460405160609190911b6bffffffffffffffffffffffff191660208201526001600160a01b039283169450911691507f306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e49060340160408051601f19818403018152908290526136a0918690602c9060289061bd31565b60405180910390a3602080546026546040516001600160a01b03928316936306cb8983936106479316910160609190911b6bffffffffffffffffffffffff1916815260140190565b6022546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015613739573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061375d919061ba6f565b6022546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af11580156137ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137ee919061b8a6565b506000600190507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561385357600080fd5b505af1158015613867573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052858460286040518563ffffffff1660e01b8152600401611960949392919061bc9f565b6026546040516001600160a01b03909116602482015260019060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260255490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561399657600080fd5b505af11580156139aa573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015613a1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a40919061b8a6565b506000602c55604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e0906024016118be565b6060601b805480602002602001604051908101604052809291908181526020016000905b828210156128ab5783829060005260206000209060020201604051806040016040529081600082018054613af49061b5d4565b80601f0160208091040260200160405190810160405280929190818152602001828054613b209061b5d4565b8015613b6d5780601f10613b4257610100808354040283529160200191613b6d565b820191906000526020600020905b815481529060010190602001808311613b5057829003601f168201915b5050505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015613bf257602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b03191681526020019060040190602082600301049283019260010382029150808411613bb45790505b50505050508152505081526020019060010190613ac1565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015613c5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c7f919061ba6f565b6025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015613cf457600080fd5b505af1158015613d08573d6000803e3d6000fd5b50506021546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015613d7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d9e919061b8a6565b506026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905280517ff48448140000000000000000000000000000000000000000000000000000000081529051919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f48448149160048082019260009290919082900301818387803b158015613e4c57600080fd5b505af1158015613e60573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526116c79288916001600160a01b0316908790602c9060289060040161b9fd565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa158015613f5d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f81919061ba6f565b613f8b919061bcf6565b67ffffffffffffffff811115613fa357613fa361b8c8565b6040519080825280601f01601f191660200182016040528015613fcd576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015614032573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614056919061ba6f565b614060919061bcf6565b61406b90600161bcd0565b67ffffffffffffffff8111156140835761408361b8c8565b6040519080825280601f01601f1916602001820160405280156140ad576020820181803683370190505b50602a906140bb908261b93e565b506000602860020180546140ce9061b5d4565b83516140da925061bcd0565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506000936001600160a01b03909216926397d340f5926004808401938290030181865afa15801561413e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614162919061ba6f565b6040805160248101859052604480820184905282518083039091018152606490910182526020810180516001600160e01b03167fcd6f4e6d0000000000000000000000000000000000000000000000000000000017905290517ff28dceb3000000000000000000000000000000000000000000000000000000008152919250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916142099160040161bce3565b600060405180830381600087803b15801561422357600080fd5b505af1158015614237573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a084526000602085015291517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b1681526131149391926001926001600160a01b03909116918a919060289060040161ba88565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516001600160a01b03909216926397d340f5926004808401938290030181865afa158015614342573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614366919061ba6f565b61437190600161bcd0565b67ffffffffffffffff8111156143895761438961b8c8565b6040519080825280601f01601f1916602001820160405280156143b3576020820181803683370190505b50602a906143c1908261b93e565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561442e57600080fd5b505af1158015614442573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef1915060340160405160208183030381529060405260018060286040518563ffffffff1660e01b8152600401612440949392919061bc9f565b604051630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561451e57600080fd5b505af1158015614532573d6000803e3d6000fd5b5050602080546040805160008152928301908190526021547f7c0dcb5f000000000000000000000000000000000000000000000000000000009091526001600160a01b039182169450637c0dcb5f935061459692916001911660286024840161bc65565b600060405180830381600087803b1580156145b057600080fd5b505af1158015611778573d6000803e3d6000fd5b6060601a805480602002602001604051908101604052809291908181526020016000905b828210156128ab5783829060005260206000200180546146079061b5d4565b80601f01602080910402602001604051908101604052809291908181526020018280546146339061b5d4565b80156146805780601f1061465557610100808354040283529160200191614680565b820191906000526020600020905b81548152906001019060200180831161466357829003601f168201915b5050505050815260200190600101906145e8565b6060601d805480602002602001604051908101604052809291908181526020016000905b828210156128ab5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561476257602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b031916815260200190600401906020826003010492830192600103820291508084116147245790505b505050505081525050815260200190600101906146b8565b6060601c805480602002602001604051908101604052809291908181526020016000905b828210156128ab5760008481526020908190206040805180820182526002860290920180546001600160a01b0316835260018101805483518187028101870190945280845293949193858301939283018282801561484857602002820191906000526020600020906000905b82829054906101000a900460e01b6001600160e01b0319168152602001906004019060208260030104928301926001038202915080841161480a5790505b5050505050815250508152602001906001019061479e565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561491357600080fd5b505af1158015614927573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506306cb8983925001610647565b60606019805480602002602001604051908101604052809291908181526020016000905b828210156128ab57838290600052602060002001805461498d9061b5d4565b80601f01602080910402602001604051908101604052809291908181526020018280546149b99061b5d4565b8015614a065780601f106149db57610100808354040283529160200191614a06565b820191906000526020600020905b8154815290600101906020018083116149e957829003601f168201915b50505050508152602001906001019061496e565b60085460009060ff1615614a32575060085460ff1690565b6040517f667f9d70000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d600482018190527f6661696c65640000000000000000000000000000000000000000000000000000602483015260009163667f9d7090604401602060405180830381865afa158015614ac3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ae7919061ba6f565b1415905090565b604051630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614b5a57600080fd5b505af1158015614b6e573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b168352614596926000916001600160a01b03169060289060040161bc65565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f5d67094f000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614cb757600080fd5b505af1158015614ccb573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f198184030181526021548383018352620186a0845260006020850181905292517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b1681526106a4949293926001600160a01b0390921691889160289060040161ba88565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614e2a57600080fd5b505af1158015614e3e573d6000803e3d6000fd5b5050602080546026546040805160609290921b6bffffffffffffffffffffffff1916938201939093528251808203601401815260215460748301855260006034840181815260549094015293517f7b15118b0000000000000000000000000000000000000000000000000000000081526001600160a01b039384169650637b15118b95506106a4949193600193921691889160289060040161ba88565b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527f60ee1247000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b158015614f8e57600080fd5b505af1158015614fa2573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637b15118b915060340160408051601f19818403018152602154838301909252916001916001600160a01b03169086908061501385620186a061b5c1565b815260006020909101526040517fffffffff0000000000000000000000000000000000000000000000000000000060e088901b1681526106a495949392919060289060040161ba88565b60208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290516000936002936001600160a01b0316926397d340f592600480830193928290030181865afa1580156150c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906150e4919061ba6f565b6150ee919061bcf6565b67ffffffffffffffff8111156151065761510661b8c8565b6040519080825280601f01601f191660200182016040528015615130576020820181803683370190505b5060208054604080517f97d340f500000000000000000000000000000000000000000000000000000000815290519394506002936001600160a01b03909216926397d340f5926004808401938290030181865afa158015615195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906151b9919061ba6f565b6151c3919061bcf6565b6151ce90600161bcd0565b67ffffffffffffffff8111156151e6576151e661b8c8565b6040519080825280601f01601f191660200182016040528015615210576020820181803683370190505b50602a9061521e908261b93e565b50604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561528b57600080fd5b505af115801561529f573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401612d04565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615332573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615356919061ba6f565b6021546040517fa9059cbb0000000000000000000000000000000000000000000000000000000081526101236004820152602481018390529192506001600160a01b03169063a9059cbb906044016020604051808303816000875af11580156153c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906153e7919061b8a6565b506027546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561545a57600080fd5b505af115801561546e573d6000803e3d6000fd5b50506021546040517ff687d12a000000000000000000000000000000000000000000000000000000008152600a60048201526001600160a01b03909116925063f687d12a9150602401600060405180830381600087803b1580156154d157600080fd5b505af11580156154e5573d6000803e3d6000fd5b50506021546040517ffc5fecd5000000000000000000000000000000000000000000000000000000008152600a6004820152600093508392506001600160a01b039091169063fc5fecd5906024016040805180830381865afa15801561554f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615573919061bd80565b602754604080516001600160a01b0385811660248301529092166044830152606480830184905281518084039091018152608490920181526020820180516001600160e01b03167f6670111200000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152929450909250737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb3916156309160040161bce3565b600060405180830381600087803b15801561564a57600080fd5b505af115801561565e573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526156f39289916001600160a01b03169060289060040161bc65565b600060405180830381600087803b15801561570d57600080fd5b505af1158015615721573d6000803e3d6000fd5b5050505050505050565b604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561579757600080fd5b505af11580156157ab573d6000803e3d6000fd5b5050602080546040516001600160a01b0390911693506397a1cef192500161240c565b6022546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa15801561581f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615843919061ba6f565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015615895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906158b9919061ba6f565b6027546025546040517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152929350163190600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b15801561593757600080fd5b505af115801561594b573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af11580156159bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906159e1919061b8a6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615a4057600080fd5b505af1158015615a54573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052878460286040518563ffffffff1660e01b8152600401615ac3949392919061bc9f565b60006040518083038186803b158015615adb57600080fd5b505afa158015615aef573d6000803e3d6000fd5b50506022546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa158015615b42573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615b66919061ba6f565b9050615b728582616a35565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa158015615bc3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615be7919061ba6f565b9050615bf38582616a35565b60275461198c9085906001600160a01b031631616a35565b60606015805480602002602001604051908101604052809291908181526020018280548015611f56576020028201919060005260206000209081546001600160a01b03168152600190910190602001808311611f38575050505050905090565b6022546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa158015615cbc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ce0919061ba6f565b6022546020546040516370a0823160e01b81526001600160a01b0391821660048201529293506000929116906370a0823190602401602060405180830381865afa158015615d32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615d56919061ba6f565b6027546026546040516001600160a01b03918216602482015292935016319060009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905260255490517fca669fa70000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152909150600190737109709ecfa91a80626ff3989d68f67f5b1dd12d9063ca669fa790602401600060405180830381600087803b158015615e1c57600080fd5b505af1158015615e30573d6000803e3d6000fd5b50506022546020546040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b039182166004820152600060248201529116925063095ea7b391506044016020604051808303816000875af1158015615ea2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ec6919061b8a6565b507f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b031663f48448146040518163ffffffff1660e01b8152600401600060405180830381600087803b158015615f2557600080fd5b505af1158015615f39573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250632810ae639150603401604051602081830303815290604052888486602c60286040518763ffffffff1660e01b8152600401615fad9695949392919061bb50565b60006040518083038186803b158015615fc557600080fd5b505afa158015615fd9573d6000803e3d6000fd5b50506022546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa15801561602c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616050919061ba6f565b905061605c8682616a35565b6022546020546040516370a0823160e01b81526001600160a01b03918216600482015260009291909116906370a0823190602401602060405180830381865afa1580156160ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906160d1919061ba6f565b90506160dd8682616a35565b6027546157219086906001600160a01b031631616a35565b60255460265460405160609190911b6bffffffffffffffffffffffff19166020820152600191829182916001600160a01b0316907f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f19818403018152602254838301835260008085526001602086015292516161919492936001600160a01b03909216928a929182919060289061bbe3565b60405180910390a3604051630618f58760e51b81527fe4dd681d000000000000000000000000000000000000000000000000000000006004820152737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561620557600080fd5b505af1158015616219573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b031692506397a1cef19150603401604051602081830303815290604052848460286040518563ffffffff1660e01b8152600401616288949392919061bc9f565b60006040518083038186803b1580156162a057600080fd5b505afa158015611ef6573d6000803e3d6000fd5b6026546040516001600160a01b03909116602482015260009060440160408051601f198184030181529181526020820180516001600160e01b0316630427d73b60e51b17905251630618f58760e51b81527fd92e233d000000000000000000000000000000000000000000000000000000006004820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c31eb0e090602401600060405180830381600087803b15801561636757600080fd5b505af115801561637b573d6000803e3d6000fd5b5050602080546040516001600160a01b039091169350637b15118b925001611367565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260019260009216906370a0823190602401602060405180830381865afa1580156163ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616413919061ba6f565b6020546040517f81bad6f3000000000000000000000000000000000000000000000000000000008152600160048201819052602482018190526044820181905260648201526001600160a01b039091166084820152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d906381bad6f39060a401600060405180830381600087803b1580156164a457600080fd5b505af11580156164b8573d6000803e3d6000fd5b505060255460265460405160609190911b6bffffffffffffffffffffffff19166020820152600093506001600160a01b0390911691507f07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c9060340160408051601f198184030181528282526021547f4d8943bb000000000000000000000000000000000000000000000000000000008452915190926001600160a01b039092169188916000918491634d8943bb916004808201926020929091908290030181865afa15801561658b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906165af919061ba6f565b604080518082018252600081526001602082015290516165d79695949392919060289061bbe3565b60405180910390a3602080546026546040516001600160a01b0392831693637c0dcb5f9361661f9316910160609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526166779287916001600160a01b03169060289060040161bc65565b600060405180830381600087803b15801561669157600080fd5b505af11580156166a5573d6000803e3d6000fd5b50506021546025546040516370a0823160e01b81526001600160a01b03918216600482015260009450911691506370a0823190602401602060405180830381865afa1580156166f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061671c919061ba6f565b905061348f611ef0848461b5c1565b6021546025546040516370a0823160e01b81526001600160a01b03918216600482015260029260009216906370a0823190602401602060405180830381865afa15801561677c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906167a0919061ba6f565b6021549091506001600160a01b031663a9059cbb6101236167c260018561b5c1565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1681526001600160a01b03909216600483015260248201526044016020604051808303816000875af1158015616825573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616849919061b8a6565b5060215460255460208054604080516001600160a01b039586166024820181905294861660448201819052959092166064830181905260848084018990528251808503909101815260a4909301825292820180516001600160e01b03167f489ca9b700000000000000000000000000000000000000000000000000000000179052517ff28dceb3000000000000000000000000000000000000000000000000000000008152929392737109709ecfa91a80626ff3989d68f67f5b1dd12d9163f28dceb39161691a919060040161bce3565b600060405180830381600087803b15801561693457600080fd5b505af1158015616948573d6000803e3d6000fd5b50506020805460265460405160609190911b6bffffffffffffffffffffffff1916928101929092526001600160a01b03169250637c0dcb5f915060340160408051601f19818403018152908290526021547fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1683526169dd928a916001600160a01b03169060289060040161bc65565b600060405180830381600087803b1580156169f757600080fd5b505af1158015616a0b573d6000803e3d6000fd5b505050505050505050565b6000616a2061b12b565b616a2b848483616aa2565b9150505b92915050565b6040517f98296c540000000000000000000000000000000000000000000000000000000081526004810183905260248101829052737109709ecfa91a80626ff3989d68f67f5b1dd12d906398296c5490604401616288565b616a9561b12b565b6106d28585858486616b1d565b600080616aaf8584616c1d565b9050616b126040518060400160405280601d81526020017f4552433139363750726f78792e736f6c3a4552433139363750726f78790000008152508286604051602001616afd92919061bdac565b60405160208183030381529060405285616c29565b9150505b9392505050565b6040517f06447d560000000000000000000000000000000000000000000000000000000081526001600160a01b03821660048201528190737109709ecfa91a80626ff3989d68f67f5b1dd12d9081906306447d5690602401600060405180830381600087803b158015616b8f57600080fd5b505af1925050508015616ba0575060015b616bb557616bb087878787616c57565b61198c565b616bc187878787616c57565b806001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015616bfc57600080fd5b505af1158015616c10573d6000803e3d6000fd5b5050505050505050505050565b6000616b168383616c70565b60c08101515160009015616c4d57616c4684848460c00151616c8b565b9050616b16565b616c468484616e31565b6000616c638483616f1c565b90506106d2858285616f28565b6000616c7c83836172f2565b616b1683836020015184616c29565b600080616c96617302565b90506000616ca486836173d5565b90506000616cbb826060015183602001518561787b565b90506000616ccb83838989617a8d565b90506000616cd88261890a565b602081015181519192509060030b15616d4b57898260400151604051602001616d0292919061bdce565b60408051601f19818403018152908290527f08c379a0000000000000000000000000000000000000000000000000000000008252616d429160040161bce3565b60405180910390fd5b6000616d8e6040518060400160405280601581526020017f4465706c6f79656420746f20616464726573733a200000000000000000000000815250836001618ad9565b6040517fc6ce059d000000000000000000000000000000000000000000000000000000008152909150737109709ecfa91a80626ff3989d68f67f5b1dd12d9063c6ce059d90616de190849060040161bce3565b602060405180830381865afa158015616dfe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616e22919061b796565b9b9a5050505050505050505050565b6040517f8d1cc9250000000000000000000000000000000000000000000000000000000081526000908190737109709ecfa91a80626ff3989d68f67f5b1dd12d90638d1cc92590616e8690879060040161bce3565b600060405180830381865afa158015616ea3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052616ecb919081019061bf08565b90506000616ef98285604051602001616ee592919061bf3d565b604051602081830303815290604052618cd9565b90506001600160a01b038116616a2b578484604051602001616d0292919061bf6c565b6000616c7c8383618cec565b6040517f667f9d700000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61036024820152737109709ecfa91a80626ff3989d68f67f5b1dd12d90600090829063667f9d7090604401602060405180830381865afa158015616fc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190616fe8919061ba6f565b90508061718f576000616ffa86618cf8565b604080518082018252600581527f352e302e3000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150617085905b60408051808201825260008082526020918201528151808301909252845182528085019082015290618ddb565b80617091575060008451115b15617114576040517f4f1ef2860000000000000000000000000000000000000000000000000000000081526001600160a01b03871690634f1ef286906170dd908890889060040161bdac565b600060405180830381600087803b1580156170f757600080fd5b505af115801561710b573d6000803e3d6000fd5b50505050617189565b6040517f3659cfe60000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152871690633659cfe690602401600060405180830381600087803b15801561717057600080fd5b505af1158015617184573d6000803e3d6000fd5b505050505b506106d2565b80600061719b82618cf8565b604080518082018252600581527f352e302e30000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201529091506171fd90617058565b80617209575060008551115b1561728e576040517f9623609d0000000000000000000000000000000000000000000000000000000081526001600160a01b03831690639623609d90617257908a908a908a9060040161c017565b600060405180830381600087803b15801561727157600080fd5b505af1158015617285573d6000803e3d6000fd5b5050505061198c565b6040517f99a88ec40000000000000000000000000000000000000000000000000000000081526001600160a01b03888116600483015287811660248301528316906399a88ec490604401600060405180830381600087803b158015616bfc57600080fd5b6172fe82826000618def565b5050565b604080518082018252600381527f6f75740000000000000000000000000000000000000000000000000000000000602082015290517fd145736c000000000000000000000000000000000000000000000000000000008152606091737109709ecfa91a80626ff3989d68f67f5b1dd12d91829063d145736c9061738990849060040161c048565b600060405180830381865afa1580156173a6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526173ce919081019061c08f565b9250505090565b6174076040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6000737109709ecfa91a80626ff3989d68f67f5b1dd12d90506174526040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b61745b85618ef2565b6020820152600061746b866192d7565b90506000836001600160a01b031663d930a0e66040518163ffffffff1660e01b8152600401600060405180830381865afa1580156174ad573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526174d5919081019061c08f565b868385602001516040516020016174ef949392919061c0d8565b60408051601f19818403018152908290527f60f9bb1100000000000000000000000000000000000000000000000000000000825291506000906001600160a01b038616906360f9bb119061754790859060040161bce3565b600060405180830381865afa158015617564573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261758c919081019061c08f565b6040517fdb4235f60000000000000000000000000000000000000000000000000000000081529091506001600160a01b0386169063db4235f6906175d490849060040161c1dc565b602060405180830381865afa1580156175f1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190617615919061b8a6565b61762a5781604051602001616d02919061c22e565b6040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061766f90849060040161c2c0565b600060405180830381865afa15801561768c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526176b4919081019061c08f565b84526040517fdb4235f60000000000000000000000000000000000000000000000000000000081526001600160a01b0386169063db4235f6906176fb90849060040161c312565b602060405180830381865afa158015617718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061773c919061b8a6565b156177d1576040517f49c4fac80000000000000000000000000000000000000000000000000000000081526001600160a01b038616906349c4fac89061778690849060040161c312565b600060405180830381865afa1580156177a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526177cb919081019061c08f565b60408501525b846001600160a01b03166349c4fac88286600001516040516020016177f6919061c364565b6040516020818303038152906040526040518363ffffffff1660e01b815260040161782292919061c3d0565b600060405180830381865afa15801561783f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052617867919081019061c08f565b606085015250608083015250949350505050565b60408051600480825260a0820190925260609160009190816020015b60608152602001906001900390816178975790505090506040518060400160405280600481526020017f6772657000000000000000000000000000000000000000000000000000000000815250816000815181106178f7576178f761c3f5565b60200260200101819052506040518060400160405280600381526020017f2d726c00000000000000000000000000000000000000000000000000000000008152508160018151811061794b5761794b61c3f5565b602002602001018190525084604051602001617967919061c424565b604051602081830303815290604052816002815181106179895761798961c3f5565b6020026020010181905250826040516020016179a5919061c490565b604051602081830303815290604052816003815181106179c7576179c761c3f5565b602002602001018190525060006179dd8261890a565b602080820151604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000008185019081528251808401845260008082529086015282518084019093529051825292810192909252919250617a6e906040805180820182526000808252602091820152815180830190925284518252808501908201529061955a565b617a835785604051602001616d02919061c4d1565b9695505050505050565b60a0810151604080518082018252600080825260209182015281518083019092528251808352928101910152606090737109709ecfa91a80626ff3989d68f67f5b1dd12d9015617add565b511590565b617c5157826020015115617b99576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605860248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b6970566572696679536f757260648201527f6365436f646560206f7074696f6e206973206074727565600000000000000000608482015260a401616d42565b8260c0015115617c51576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f54686520606c6963656e73655479706560206f7074696f6e2063616e6e6f742060448201527f62652075736564207768656e207468652060736b69704c6963656e736554797060648201527f6560206f7074696f6e2069732060747275656000000000000000000000000000608482015260a401616d42565b6040805160ff8082526120008201909252600091816020015b6060815260200190600190039081617c6a57905050905060006040518060400160405280600381526020017f6e70780000000000000000000000000000000000000000000000000000000000815250828280617cc59061c562565b935060ff1681518110617cda57617cda61c3f5565b60200260200101819052506040518060400160405280600d81526020017f302e302e312d616c7068612e3700000000000000000000000000000000000000815250604051602001617d2b919061c581565b604051602081830303815290604052828280617d469061c562565b935060ff1681518110617d5b57617d5b61c3f5565b60200260200101819052506040518060400160405280600681526020017f6465706c6f790000000000000000000000000000000000000000000000000000815250828280617da89061c562565b935060ff1681518110617dbd57617dbd61c3f5565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e74726163744e616d65000000000000000000000000000000000000815250828280617e0a9061c562565b935060ff1681518110617e1f57617e1f61c3f5565b60200260200101819052508760200151828280617e3b9061c562565b935060ff1681518110617e5057617e5061c3f5565b60200260200101819052506040518060400160405280600e81526020017f2d2d636f6e747261637450617468000000000000000000000000000000000000815250828280617e9d9061c562565b935060ff1681518110617eb257617eb261c3f5565b602090810291909101015287518282617eca8161c562565b935060ff1681518110617edf57617edf61c3f5565b60200260200101819052506040518060400160405280600981526020017f2d2d636861696e49640000000000000000000000000000000000000000000000815250828280617f2c9061c562565b935060ff1681518110617f4157617f4161c3f5565b6020026020010181905250617f55466195bb565b8282617f608161c562565b935060ff1681518110617f7557617f7561c3f5565b60200260200101819052506040518060400160405280600f81526020017f2d2d6275696c64496e666f46696c650000000000000000000000000000000000815250828280617fc29061c562565b935060ff1681518110617fd757617fd761c3f5565b602002602001018190525086828280617fef9061c562565b935060ff16815181106180045761800461c3f5565b602090810291909101015285511561812b5760408051808201909152601581527f2d2d636f6e7374727563746f7242797465636f64650000000000000000000000602082015282826180558161c562565b935060ff168151811061806a5761806a61c3f5565b60209081029190910101526040517f71aad10d0000000000000000000000000000000000000000000000000000000081526001600160a01b038416906371aad10d906180ba90899060040161bce3565b600060405180830381865afa1580156180d7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526180ff919081019061c08f565b828261810a8161c562565b935060ff168151811061811f5761811f61c3f5565b60200260200101819052505b8460200151156181fb5760408051808201909152601281527f2d2d766572696679536f75726365436f64650000000000000000000000000000602082015282826181748161c562565b935060ff16815181106181895761818961c3f5565b60200260200101819052506040518060400160405280600581526020017f66616c73650000000000000000000000000000000000000000000000000000008152508282806181d69061c562565b935060ff16815181106181eb576181eb61c3f5565b60200260200101819052506183c2565b618233617ad88660a0015160408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6182c65760408051808201909152600d81527f2d2d6c6963656e73655479706500000000000000000000000000000000000000602082015282826182768161c562565b935060ff168151811061828b5761828b61c3f5565b60200260200101819052508460a001516040516020016182ab919061c424565b6040516020818303038152906040528282806181d69061c562565b8460c0015115801561830957506040808901518151808301835260008082526020918201528251808401909352815183529081019082015261830790511590565b155b156183c25760408051808201909152600d81527f2d2d6c6963656e736554797065000000000000000000000000000000000000006020820152828261834d8161c562565b935060ff16815181106183625761836261c3f5565b60200260200101819052506183768861965b565b604051602001618386919061c424565b6040516020818303038152906040528282806183a19061c562565b935060ff16815181106183b6576183b661c3f5565b60200260200101819052505b604080860151815180830183526000808252602091820152825180840190935281518352908101908201526183f690511590565b61848b5760408051808201909152600b81527f2d2d72656c617965724964000000000000000000000000000000000000000000602082015282826184398161c562565b935060ff168151811061844e5761844e61c3f5565b6020026020010181905250846040015182828061846a9061c562565b935060ff168151811061847f5761847f61c3f5565b60200260200101819052505b6060850151156185ac5760408051808201909152600681527f2d2d73616c740000000000000000000000000000000000000000000000000000602082015282826184d48161c562565b935060ff16815181106184e9576184e961c3f5565b602090810291909101015260608501516040517fb11a19e800000000000000000000000000000000000000000000000000000000815260048101919091526001600160a01b0384169063b11a19e890602401600060405180830381865afa158015618558573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618580919081019061c08f565b828261858b8161c562565b935060ff16815181106185a0576185a061c3f5565b60200260200101819052505b60e085015151156186535760408051808201909152600a81527f2d2d6761734c696d697400000000000000000000000000000000000000000000602082015282826185f68161c562565b935060ff168151811061860b5761860b61c3f5565b60200260200101819052506186278560e00151600001516195bb565b82826186328161c562565b935060ff16815181106186475761864761c3f5565b60200260200101819052505b60e085015160200151156186fd5760408051808201909152600a81527f2d2d676173507269636500000000000000000000000000000000000000000000602082015282826186a08161c562565b935060ff16815181106186b5576186b561c3f5565b60200260200101819052506186d18560e00151602001516195bb565b82826186dc8161c562565b935060ff16815181106186f1576186f161c3f5565b60200260200101819052505b60e085015160400151156187a75760408051808201909152600e81527f2d2d6d61784665655065724761730000000000000000000000000000000000006020820152828261874a8161c562565b935060ff168151811061875f5761875f61c3f5565b602002602001018190525061877b8560e00151604001516195bb565b82826187868161c562565b935060ff168151811061879b5761879b61c3f5565b60200260200101819052505b60e085015160600151156188515760408051808201909152601681527f2d2d6d61785072696f7269747946656550657247617300000000000000000000602082015282826187f48161c562565b935060ff16815181106188095761880961c3f5565b60200260200101819052506188258560e00151606001516195bb565b82826188308161c562565b935060ff16815181106188455761884561c3f5565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561886f5761886f61b8c8565b6040519080825280602002602001820160405280156188a257816020015b606081526020019060019003908161888d5790505b50905060005b8260ff168160ff1610156188fb57838160ff16815181106188cb576188cb61c3f5565b6020026020010151828260ff16815181106188e8576188e861c3f5565b60209081029190910101526001016188a8565b5093505050505b949350505050565b6189316040518060600160405280600060030b815260200160608152602001606081525090565b60408051808201825260048082527f6261736800000000000000000000000000000000000000000000000000000000602083015291517fd145736c000000000000000000000000000000000000000000000000000000008152737109709ecfa91a80626ff3989d68f67f5b1dd12d92600091849163d145736c916189b79186910161c5ec565b600060405180830381865afa1580156189d4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526189fc919081019061c08f565b90506000618a0a868361a14a565b90506000846001600160a01b031663f45c1ce7836040518263ffffffff1660e01b8152600401618a3a919061b484565b6000604051808303816000875af1158015618a59573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052618a81919081019061c633565b805190915060030b15801590618a9a5750602081015151155b8015618aa95750604081015151155b15617a835781600081518110618ac157618ac161c3f5565b6020026020010151604051602001616d02919061c6e9565b60606000618b0e8560408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600080825260209182015281518083019092528651825280870190820152909150618b459082905b9061a29f565b15618ca2576000618bc282618bbc84618bb6618b888a60408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925282518252918201519181019190915290565b9061a2c6565b9061a328565b604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150618c2690829061a29f565b15618c9057604080518082018252600181527f0a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618c8d905b829061a3ad565b90505b618c998161a3d3565b92505050616b16565b8215618cbb578484604051602001616d0292919061c8d5565b5050604080516020810190915260008152616b16565b509392505050565b6000808251602084016000f09392505050565b6172fe82826001618def565b60408051600481526024810182526020810180516001600160e01b03167fad3cb1cc00000000000000000000000000000000000000000000000000000000179052905160609160009182916001600160a01b03861691618d58919061c97c565b600060405180830381855afa9150503d8060008114618d93576040519150601f19603f3d011682016040523d82523d6000602084013e618d98565b606091505b5091509150818015618dab575060208151115b15618dc45780806020019051810190618902919061c08f565b505060408051602081019091526000815292915050565b6000618de7838361a43c565b159392505050565b8160a0015115618dfe57505050565b6000618e0b84848461a517565b90506000618e188261890a565b602081015181519192509060030b158015618eb45750604080518082018252600781527f535543434553530000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152618eb490604080518082018252600080825260209182015281518083019092528451825280850190820152618b3f565b15618ec157505050505050565b60408201515115618ee1578160400151604051602001616d02919061c998565b80604051602001616d02919061c9f6565b60606000618f278360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152909150618f8c905b829061955a565b15618ffb57604080518082018252600481527f2e736f6c0000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b1690618ff690839061aab2565b61a3d3565b604080518082018252600181527f3a000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261905d905b829061ab3c565b60010361912a57604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526190c390618c86565b50604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b1690618ff6905b839061a3ad565b604080518082018252600581527f2e6a736f6e0000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261918990618f85565b156192c057604080518082018252600181527f2f000000000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201819052845180860190955292518452830152906191f190839061abd6565b905060008160018351619204919061b5c1565b815181106192145761921461c3f5565b602002602001015190506192b7618ff661928a6040518060400160405280600581526020017f2e6a736f6e00000000000000000000000000000000000000000000000000000081525060408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b6040805180820182526000808252602091820152815180830190925285518252808601908201529061aab2565b95945050505050565b82604051602001616d02919061ca61565b50919050565b6060600061930c8360408051808201825260008082526020918201528151808301909252825182529182019181019190915290565b604080518082018252600481527f2e736f6c000000000000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015290915061936e90618f85565b1561937c57616b168161a3d3565b604080518082018252600181527f3a00000000000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526193db90619056565b60010361944557604080518082018252600181527f3a0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152616b1690618ff690619123565b604080518082018252600581527f2e6a736f6e000000000000000000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526194a490618f85565b156192c057604080518082018252600181527f2f0000000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082018190528451808601909552925184528301529061950c90839061abd6565b9050600181511115619548578060028251619527919061b5c1565b815181106195375761953761c3f5565b602002602001015192505050919050565b5082604051602001616d02919061ca61565b80518251600091111561956f57506000616a2f565b815183516020850151600092916195859161bcd0565b61958f919061b5c1565b9050826020015181036195a6576001915050616a2f565b82516020840151819020912014905092915050565b606060006195c88361ac7b565b600101905060008167ffffffffffffffff8111156195e8576195e861b8c8565b6040519080825280601f01601f191660200182016040528015619612576020820181803683370190505b5090508181016020015b600019017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461961c57509392505050565b604081810151815180830183526000808252602091820181905283518085018552835181529282018383015283518085018552600a81527f554e4c4943454e53454400000000000000000000000000000000000000000000818401908152855180870187528381528401929092528451808601909552518452908301526060916196e7905b8290618ddb565b1561972757505060408051808201909152600481527f4e6f6e65000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600981527f556e6c6963656e7365000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619786906196e0565b156197c657505060408051808201909152600981527f556e6c6963656e736500000000000000000000000000000000000000000000006020820152919050565b604080518082018252600381527f4d4954000000000000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619825906196e0565b1561986557505060408051808201909152600381527f4d495400000000000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d322e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526198c4906196e0565b806199295750604080518082018252601081527f47504c2d322e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619929906196e0565b1561996957505060408051808201909152600981527f474e552047504c763200000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f47504c2d332e302d6f6e6c790000000000000000000000000000000000000000602080830191825283518085018552600080825290820152835180850190945291518352908201526199c8906196e0565b80619a2d5750604080518082018252601081527f47504c2d332e302d6f722d6c617465720000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619a2d906196e0565b15619a6d57505060408051808201909152600981527f474e552047504c763300000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d322e312d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619acc906196e0565b80619b315750604080518082018252601181527f4c47504c2d322e312d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619b31906196e0565b15619b7157505060408051808201909152600c81527f474e55204c47504c76322e3100000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4c47504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619bd0906196e0565b80619c355750604080518082018252601181527f4c47504c2d332e302d6f722d6c6174657200000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619c35906196e0565b15619c7557505060408051808201909152600a81527f474e55204c47504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d322d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619cd4906196e0565b15619d1457505060408051808201909152600c81527f4253442d322d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600c81527f4253442d332d436c61757365000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619d73906196e0565b15619db357505060408051808201909152600c81527f4253442d332d436c6175736500000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4d504c2d322e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619e12906196e0565b15619e5257505060408051808201909152600781527f4d504c2d322e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600781527f4f534c2d332e300000000000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619eb1906196e0565b15619ef157505060408051808201909152600781527f4f534c2d332e30000000000000000000000000000000000000000000000000006020820152919050565b604080518082018252600a81527f4170616368652d322e300000000000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619f50906196e0565b15619f9057505060408051808201909152600a81527f4170616368652d322e30000000000000000000000000000000000000000000006020820152919050565b604080518082018252600d81527f4147504c2d332e302d6f6e6c790000000000000000000000000000000000000060208083019182528351808501855260008082529082015283518085019094529151835290820152619fef906196e0565b8061a0545750604080518082018252601181527f4147504c2d332e302d6f722d6c617465720000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a054906196e0565b1561a09457505060408051808201909152600a81527f474e55204147504c7633000000000000000000000000000000000000000000006020820152919050565b604080518082018252600881527f4255534c2d312e310000000000000000000000000000000000000000000000006020808301918252835180850185526000808252908201528351808501909452915183529082015261a0f3906196e0565b1561a13357505060408051808201909152600781527f42534c20312e31000000000000000000000000000000000000000000000000006020820152919050565b60408084015184519151616d02929060200161cb3f565b60608060005b845181101561a1d5578185828151811061a16c5761a16c61c3f5565b602002602001015160405160200161a18592919061bf3d565b60405160208183030381529060405291506001855161a1a4919061b5c1565b811461a1cd578160405160200161a1bb919061cca8565b60405160208183030381529060405291505b60010161a150565b5060408051600380825260808201909252600091816020015b606081526020019060019003908161a1ee579050509050838160008151811061a2195761a21961c3f5565b60200260200101819052506040518060400160405280600281526020017f2d630000000000000000000000000000000000000000000000000000000000008152508160018151811061a26d5761a26d61c3f5565b6020026020010181905250818160028151811061a28c5761a28c61c3f5565b6020908102919091010152949350505050565b602080830151835183519284015160009361a2bd929184919061ad5d565b14159392505050565b6040805180820190915260008082526020820152600061a2f8846000015185602001518560000151866020015161ae6e565b905083602001518161a30a919061b5c1565b8451859061a31990839061b5c1565b90525060208401525090919050565b604080518082019091526000808252602082015281518351101561a34d575081616a2f565b602080830151908401516001911461a3745750815160208481015190840151829020919020145b801561a3a55782518451859061a38b90839061b5c1565b905250825160208501805161a3a190839061bcd0565b9052505b509192915050565b604080518082019091526000808252602082015261a3cc83838361af8e565b5092915050565b60606000826000015167ffffffffffffffff81111561a3f45761a3f461b8c8565b6040519080825280601f01601f19166020018201604052801561a41e576020820181803683370190505b509050600060208201905061a3cc818560200151866000015161b039565b815181516000919081111561a44f575081515b6020808501519084015160005b8381101561a508578251825180821461a4d857600019602087101561a4b75760018461a48989602061b5c1565b61a493919061bcd0565b61a49e90600861cce9565b61a4a990600261cde7565b61a4b3919061b5c1565b1990505b818116838216818103911461a4d5579750616a2f9650505050505050565b50505b61a4e360208661bcd0565b945061a4f060208561bcd0565b9350505060208161a501919061bcd0565b905061a45c565b5084518651617a83919061cdf3565b6060600061a523617302565b6040805160ff808252612000820190925291925060009190816020015b606081526020019060019003908161a54057905050905060006040518060400160405280600381526020017f6e7078000000000000000000000000000000000000000000000000000000000081525082828061a59b9061c562565b935060ff168151811061a5b05761a5b061c3f5565b60200260200101819052506040518060400160405280600781526020017f5e312e33322e330000000000000000000000000000000000000000000000000081525060405160200161a601919061ce13565b60405160208183030381529060405282828061a61c9061c562565b935060ff168151811061a6315761a63161c3f5565b60200260200101819052506040518060400160405280600881526020017f76616c696461746500000000000000000000000000000000000000000000000081525082828061a67e9061c562565b935060ff168151811061a6935761a69361c3f5565b60200260200101819052508260405160200161a6af919061c490565b60405160208183030381529060405282828061a6ca9061c562565b935060ff168151811061a6df5761a6df61c3f5565b60200260200101819052506040518060400160405280600a81526020017f2d2d636f6e74726163740000000000000000000000000000000000000000000081525082828061a72c9061c562565b935060ff168151811061a7415761a74161c3f5565b602002602001018190525061a756878461b0b3565b828261a7618161c562565b935060ff168151811061a7765761a77661c3f5565b60209081029190910101528551511561a8225760408051808201909152600b81527f2d2d7265666572656e63650000000000000000000000000000000000000000006020820152828261a7c88161c562565b935060ff168151811061a7dd5761a7dd61c3f5565b602002602001018190525061a7f686600001518461b0b3565b828261a8018161c562565b935060ff168151811061a8165761a81661c3f5565b60200260200101819052505b85608001511561a8905760408051808201909152601881527f2d2d756e73616665536b697053746f72616765436865636b00000000000000006020820152828261a86b8161c562565b935060ff168151811061a8805761a88061c3f5565b602002602001018190525061a8f6565b841561a8f65760408051808201909152601281527f2d2d726571756972655265666572656e636500000000000000000000000000006020820152828261a8d58161c562565b935060ff168151811061a8ea5761a8ea61c3f5565b60200260200101819052505b6040860151511561a9925760408051808201909152600d81527f2d2d756e73616665416c6c6f77000000000000000000000000000000000000006020820152828261a9408161c562565b935060ff168151811061a9555761a95561c3f5565b6020026020010181905250856040015182828061a9719061c562565b935060ff168151811061a9865761a98661c3f5565b60200260200101819052505b85606001511561a9fc5760408051808201909152601481527f2d2d756e73616665416c6c6f7752656e616d65730000000000000000000000006020820152828261a9db8161c562565b935060ff168151811061a9f05761a9f061c3f5565b60200260200101819052505b60008160ff1667ffffffffffffffff81111561aa1a5761aa1a61b8c8565b60405190808252806020026020018201604052801561aa4d57816020015b606081526020019060019003908161aa385790505b50905060005b8260ff168160ff16101561aaa657838160ff168151811061aa765761aa7661c3f5565b6020026020010151828260ff168151811061aa935761aa9361c3f5565b602090810291909101015260010161aa53565b50979650505050505050565b604080518082019091526000808252602082015281518351101561aad7575081616a2f565b8151835160208501516000929161aaed9161bcd0565b61aaf7919061b5c1565b6020840151909150600190821461ab18575082516020840151819020908220145b801561ab335783518551869061ab2f90839061b5c1565b9052505b50929392505050565b600080826000015161ab60856000015186602001518660000151876020015161ae6e565b61ab6a919061bcd0565b90505b8351602085015161ab7e919061bcd0565b811161a3cc578161ab8e8161ce58565b925050826000015161abc585602001518361aba9919061b5c1565b865161abb5919061b5c1565b838660000151876020015161ae6e565b61abcf919061bcd0565b905061ab6d565b6060600061abe4848461ab3c565b61abef90600161bcd0565b67ffffffffffffffff81111561ac075761ac0761b8c8565b60405190808252806020026020018201604052801561ac3a57816020015b606081526020019060019003908161ac255790505b50905060005b8151811015618cd15761ac56618ff6868661a3ad565b82828151811061ac685761ac6861c3f5565b602090810291909101015260010161ac40565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061acc4577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061acf0576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061ad0e57662386f26fc10000830492506010015b6305f5e100831061ad26576305f5e100830492506008015b612710831061ad3a57612710830492506004015b6064831061ad4c576064830492506002015b600a8310616a2f5760010192915050565b60008085841161ae64576020841161ae10576000841561ada857600161ad8486602061b5c1565b61ad8f90600861cce9565b61ad9a90600261cde7565b61ada4919061b5c1565b1990505b835181168561adb7898961bcd0565b61adc1919061b5c1565b805190935082165b81811461adfb5787841161ade35787945050505050618902565b8361aded8161ce72565b94505082845116905061adc9565b61ae05878561bcd0565b945050505050618902565b83832061ae1d858861b5c1565b61ae27908761bcd0565b91505b85821061ae625784822080820361ae4f5761ae45868461bcd0565b9350505050618902565b61ae5a60018461b5c1565b92505061ae2a565b505b5092949350505050565b6000838186851161af79576020851161af28576000851561aeba57600161ae9687602061b5c1565b61aea190600861cce9565b61aeac90600261cde7565b61aeb6919061b5c1565b1990505b8451811660008761aecb8b8b61bcd0565b61aed5919061b5c1565b855190915083165b82811461af1a5781861061af025761aef58b8b61bcd0565b9650505050505050618902565b8561af0c8161ce58565b96505083865116905061aedd565b859650505050505050618902565b508383206000905b61af3a868961b5c1565b821161af775785832080820361af565783945050505050618902565b61af6160018561bcd0565b935050818061af6f9061ce58565b92505061af30565b505b61af83878761bcd0565b979650505050505050565b6040805180820190915260008082526020820152600061afc0856000015186602001518660000151876020015161ae6e565b60208087018051918601919091525190915061afdc908261b5c1565b83528451602086015161afef919061bcd0565b810361affe576000855261b030565b8351835161b00c919061bcd0565b8551869061b01b90839061b5c1565b905250835161b02a908261bcd0565b60208601525b50909392505050565b6020811061b071578151835261b05060208461bcd0565b925061b05d60208361bcd0565b915061b06a60208261b5c1565b905061b039565b600019811561b0a057600161b08783602061b5c1565b61b0939061010061cde7565b61b09d919061b5c1565b90505b9151835183169219169190911790915250565b6060600061b0c184846173d5565b805160208083015160405193945061b0db9390910161ce89565b60405160208183030381529060405291505092915050565b610b67806200cee283390190565b6107b6806200da4983390190565b61106f806200e1ff83390190565b612072806200f26e83390190565b6040518060e0016040528060608152602001606081526020016060815260200160001515815260200160001515815260200160001515815260200161b16e61b173565b905290565b6040518061010001604052806000151581526020016000151581526020016060815260200160008019168152602001606081526020016060815260200160001515815260200161b16e6040518060800160405280600081526020016000815260200160008152602001600081525090565b602080825282518282018190526000918401906040840190835b8181101561b2255783516001600160a01b031683526020938401939092019160010161b1fe565b509095945050505050565b60005b8381101561b24b57818101518382015260200161b233565b50506000910152565b6000815180845261b26c81602086016020860161b230565b601f01601f19169290920160200192915050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b37c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516001600160a01b03168652602090810151604082880181905281519088018190529101906060600582901b88018101919088019060005b8181101561b362577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa08a850301835261b34c84865161b254565b602095860195909450929092019160010161b312565b50919750505060209485019492909201915060010161b2a8565b50929695505050505050565b600081518084526020840193506020830160005b8281101561b3dc5781517fffffffff000000000000000000000000000000000000000000000000000000001686526020958601959091019060010161b39c565b5093949350505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b37c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0878603018452815180516040875261b452604088018261b254565b905060208201519150868103602088015261b46d818361b388565b96505050602093840193919091019060010161b40e565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b37c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845261b4e685835161b254565b9450602093840193919091019060010161b4ac565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561b37c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc087860301845281516001600160a01b038151168652602081015190506040602087015261b57c604087018261b388565b955050602093840193919091019060010161b523565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115616a2f57616a2f61b592565b600181811c9082168061b5e857607f821691505b6020821081036192d1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600081546001600160a01b038116845260ff8160a01c1615156020850152506001600160a01b0360018301541660408401526002820160a060608501526000815461b66b8161b5d4565b8060a0880152600182166000811461b68a576001811461b6c45761b6f8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00831660c089015260c082151560051b890101935061b6f8565b84600052602060002060005b8381101561b6ef5781548a820160c0015260019091019060200161b6d0565b890160c0019450505b50505060038401546080860152809250505092915050565b60c08152600061b72360c083018861b254565b6001600160a01b0387166020840152828103604084015261b744818761b254565b85546060850152600186015460ff1615156080850152905082810360a084015261b76e818561b621565b98975050505050505050565b80516001600160a01b038116811461b79157600080fd5b919050565b60006020828403121561b7a857600080fd5b616b168261b77a565b610100815260056101008201527f544f4b454e000000000000000000000000000000000000000000000000000000610120820152610140602082015260036101408201527f544b4e000000000000000000000000000000000000000000000000000000000061016082015260006101808201905060ff881660408301528660608301526003861061b86b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8560808301528460a083015261b88c60c08301856001600160a01b03169052565b6001600160a01b03831660e0830152979650505050505050565b60006020828403121561b8b857600080fd5b81518015158114616b1657600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b601f82111561348f57806000526020600020601f840160051c8101602085101561b91e5750805b601f840160051c820191505b818110156106d2576000815560010161b92a565b815167ffffffffffffffff81111561b9585761b95861b8c8565b61b96c8161b966845461b5d4565b8461b8f7565b6020601f82116001811461b9a0576000831561b9885750848201515b600019600385901b1c1916600184901b1784556106d2565b600084815260208120601f198516915b8281101561b9d0578785015182556020948501946001909201910161b9b0565b508482101561b9ee5786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60e08152600061ba1060e083018961b254565b8760208401526001600160a01b0387166040840152828103606084015261ba37818761b254565b85546080850152600186015460ff16151560a085015290505b82810360c084015261ba62818561b621565b9998505050505050505050565b60006020828403121561ba8157600080fd5b5051919050565b60e08152600061ba9b60e083018961b254565b8760208401526001600160a01b0387166040840152828103606084015261bac2818761b254565b855160808501526020860151151560a0850152905061ba50565b6101208152600061baf161012083018b61b254565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261bb24818761b254565b855460c0850152600186015460ff16151560e085015290505b828103610100840152616e22818561b621565b60e08152600061bb6360e083018961b254565b876020840152866040840152828103606084015261ba37818761b254565b6101208152600061bb9661012083018b61b254565b6001600160a01b038a16602084015288604084015287606084015286608084015282810360a084015261bbc9818761b254565b855160c08501526020860151151560e0850152905061bb3d565b6101208152600061bbf861012083018a61b254565b6001600160a01b03891660208401528760408401528660608401528560808401528281038060a08501526000825261bc3f60c0850187805182526020908101511515910152565b602081016101008501525061bc57602082018561b621565b9a9950505050505050505050565b60808152600061bc78608083018761b254565b8560208401526001600160a01b0385166040840152828103606084015261af83818561b621565b60808152600061bcb2608083018761b254565b856020840152846040840152828103606084015261af83818561b621565b80820180821115616a2f57616a2f61b592565b602081526000616b16602083018461b254565b60008261bd2c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60a08152600061bd4460a083018761b254565b828103602084015261bd56818761b254565b85546040850152600186015460ff16151560608501529050828103608084015261af83818561b621565b6000806040838503121561bd9357600080fd5b61bd9c8361b77a565b6020939093015192949293505050565b6001600160a01b0383168152604060208201526000618902604083018461b254565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161be0681601a85016020880161b230565b7f3a20000000000000000000000000000000000000000000000000000000000000601a91840191820152835161be4381601c84016020880161b230565b01601c01949350505050565b6040516060810167ffffffffffffffff8111828210171561be725761be7261b8c8565b60405290565b60008067ffffffffffffffff84111561be935761be9361b8c8565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff8211171561bec25761bec261b8c8565b60405283815290508082840185101561beda57600080fd5b618cd184602083018561b230565b600082601f83011261bef957600080fd5b616b168383516020850161be78565b60006020828403121561bf1a57600080fd5b815167ffffffffffffffff81111561bf3157600080fd5b616a2b8482850161bee8565b6000835161bf4f81846020880161b230565b83519083019061bf6381836020880161b230565b01949350505050565b7f4661696c656420746f206465706c6f7920636f6e74726163742000000000000081526000835161bfa481601a85016020880161b230565b7f207573696e6720636f6e7374727563746f722064617461202200000000000000601a91840191820152835161bfe181603384016020880161b230565b7f220000000000000000000000000000000000000000000000000000000000000060339290910191820152603401949350505050565b6001600160a01b03841681526001600160a01b03831660208201526060604082015260006192b7606083018461b254565b60408152600b60408201527f464f554e4452595f4f55540000000000000000000000000000000000000000006060820152608060208201526000616b16608083018461b254565b60006020828403121561c0a157600080fd5b815167ffffffffffffffff81111561c0b857600080fd5b8201601f8101841361c0c957600080fd5b616a2b8482516020840161be78565b6000855161c0ea818460208a0161b230565b7f2f00000000000000000000000000000000000000000000000000000000000000908301908152855161c124816001840160208a0161b230565b7f2f0000000000000000000000000000000000000000000000000000000000000060019290910191820152845161c16281600284016020890161b230565b6001818301019150507f2f000000000000000000000000000000000000000000000000000000000000006001820152835161c1a481600284016020880161b230565b7f2e6a736f6e000000000000000000000000000000000000000000000000000000600292909101918201526007019695505050505050565b60408152600061c1ef604083018461b254565b8281036020840152600481527f2e6173740000000000000000000000000000000000000000000000000000000060208201526040810191505092915050565b7f436f756c64206e6f742066696e642041535420696e206172746966616374200081526000825161c26681601f85016020870161b230565b7f2e205365742060617374203d20747275656020696e20666f756e6472792e746f601f9390910192830152507f6d6c000000000000000000000000000000000000000000000000000000000000603f820152604101919050565b60408152600061c2d3604083018461b254565b8281036020840152601181527f2e6173742e6162736f6c7574655061746800000000000000000000000000000060208201526040810191505092915050565b60408152600061c325604083018461b254565b8281036020840152600c81527f2e6173742e6c6963656e7365000000000000000000000000000000000000000060208201526040810191505092915050565b7f2e6d657461646174612e736f75726365732e5b2700000000000000000000000081526000825161c39c81601485016020870161b230565b7f275d2e6b656363616b32353600000000000000000000000000000000000000006014939091019283015250602001919050565b60408152600061c3e3604083018561b254565b8281036020840152616b12818561b254565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f220000000000000000000000000000000000000000000000000000000000000081526000825161c45c81600185016020870161b230565b7f22000000000000000000000000000000000000000000000000000000000000006001939091019283015250600201919050565b6000825161c4a281846020870161b230565b7f2f6275696c642d696e666f000000000000000000000000000000000000000000920191825250600b01919050565b7f436f756c64206e6f742066696e64206275696c642d696e666f2066696c65207781527f697468206d61746368696e6720736f7572636520636f6465206861736820666f60208201527f7220636f6e74726163742000000000000000000000000000000000000000000060408201526000825161c55581604b85016020870161b230565b91909101604b0192915050565b600060ff821660ff810361c5785761c57861b592565b60010192915050565b7f406f70656e7a657070656c696e2f646566656e6465722d6465706c6f792d636c81527f69656e742d636c6940000000000000000000000000000000000000000000000060208201526000825161c5df81602985016020870161b230565b9190910160290192915050565b60408152601660408201527f4f50454e5a455050454c494e5f424153485f50415448000000000000000000006060820152608060208201526000616b16608083018461b254565b60006020828403121561c64557600080fd5b815167ffffffffffffffff81111561c65c57600080fd5b82016060818503121561c66e57600080fd5b61c67661be4f565b81518060030b811461c68757600080fd5b8152602082015167ffffffffffffffff81111561c6a357600080fd5b61c6af8682850161bee8565b602083015250604082015167ffffffffffffffff81111561c6cf57600080fd5b61c6db8682850161bee8565b604083015250949350505050565b7f4661696c656420746f2072756e206261736820636f6d6d616e6420776974682081527f220000000000000000000000000000000000000000000000000000000000000060208201526000825161c74781602185016020870161b230565b7f222e20496620796f7520617265207573696e672057696e646f77732c2073657460219390910192830152507f20746865204f50454e5a455050454c494e5f424153485f5041544820656e766960418201527f726f6e6d656e74207661726961626c6520746f207468652066756c6c7920717560618201527f616c69666965642070617468206f66207468652062617368206578656375746160818201527f626c652e20466f72206578616d706c652c20696620796f75206172652075736960a18201527f6e672047697420666f722057696e646f77732c206164642074686520666f6c6c60c18201527f6f77696e67206c696e6520696e20746865202e656e762066696c65206f66207960e18201527f6f75722070726f6a65637420287573696e6720666f727761726420736c6173686101018201527f6573293a0a4f50454e5a455050454c494e5f424153485f504154483d22433a2f6101218201527f50726f6772616d2046696c65732f4769742f62696e2f6261736822000000000061014182015261015c01919050565b7f4661696c656420746f2066696e64206c696e652077697468207072656669782081527f270000000000000000000000000000000000000000000000000000000000000060208201526000835161c93381602185016020880161b230565b7f2720696e206f75747075743a2000000000000000000000000000000000000000602191840191820152835161c97081602e84016020880161b230565b01602e01949350505050565b6000825161c98e81846020870161b230565b9190910192915050565b7f4661696c656420746f2072756e2075706772616465207361666574792076616c81527f69646174696f6e3a20000000000000000000000000000000000000000000000060208201526000825161c5df81602985016020870161b230565b7f55706772616465207361666574792076616c69646174696f6e206661696c656481527f3a0a00000000000000000000000000000000000000000000000000000000000060208201526000825161ca5481602285016020870161b230565b9190910160220192915050565b7f436f6e7472616374206e616d652000000000000000000000000000000000000081526000825161ca9981600e85016020870161b230565b7f206d75737420626520696e2074686520666f726d6174204d79436f6e74726163600e9390910192830152507f742e736f6c3a4d79436f6e7472616374206f72204d79436f6e74726163742e73602e8201527f6f6c206f72206f75742f4d79436f6e74726163742e736f6c2f4d79436f6e7472604e8201527f6163742e6a736f6e000000000000000000000000000000000000000000000000606e820152607601919050565b7f53504458206c6963656e7365206964656e74696669657220000000000000000081526000835161cb7781601885016020880161b230565b7f20696e2000000000000000000000000000000000000000000000000000000000601891840191820152835161cbb481601c84016020880161b230565b7f20646f6573206e6f74206c6f6f6b206c696b65206120737570706f7274656420601c92909101918201527f6c6963656e736520666f7220626c6f636b206578706c6f726572207665726966603c8201527f69636174696f6e2e205573652074686520606c6963656e73655479706560206f605c8201527f7074696f6e20746f20737065636966792061206c6963656e736520747970652c607c8201527f206f7220736574207468652060736b69704c6963656e73655479706560206f70609c8201527f74696f6e20746f2060747275656020746f20736b69702e00000000000000000060bc82015260d301949350505050565b6000825161ccba81846020870161b230565b7f2000000000000000000000000000000000000000000000000000000000000000920191825250600101919050565b8082028115828204841417616a2f57616a2f61b592565b6001815b600184111561cd3b5780850481111561cd1f5761cd1f61b592565b600184161561cd2d57908102905b60019390931c92800261cd04565b935093915050565b60008261cd5257506001616a2f565b8161cd5f57506000616a2f565b816001811461cd75576002811461cd7f5761cd9b565b6001915050616a2f565b60ff84111561cd905761cd9061b592565b50506001821b616a2f565b5060208310610133831016604e8410600b841016171561cdbe575081810a616a2f565b61cdcb600019848461cd00565b806000190482111561cddf5761cddf61b592565b029392505050565b6000616b16838361cd43565b818103600083128015838313168383128216171561a3cc5761a3cc61b592565b7f406f70656e7a657070656c696e2f75706772616465732d636f7265400000000081526000825161ce4b81601c85016020870161b230565b91909101601c0192915050565b6000600019820361ce6b5761ce6b61b592565b5060010190565b60008161ce815761ce8161b592565b506000190190565b6000835161ce9b81846020880161b230565b7f3a00000000000000000000000000000000000000000000000000000000000000908301908152835161ced581600184016020880161b230565b0160010194935050505056fe60c0604052600d60809081526c2bb930b83832b21022ba3432b960991b60a05260009061002c9082610114565b506040805180820190915260048152630ae8aa8960e31b60208201526001906100559082610114565b506002805460ff1916601217905534801561006f57600080fd5b506101d2565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061009f57607f821691505b6020821081036100bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561010f57806000526020600020601f840160051c810160208510156100ec5750805b601f840160051c820191505b8181101561010c57600081556001016100f8565b50505b505050565b81516001600160401b0381111561012d5761012d610075565b6101418161013b845461008b565b846100c5565b6020601f821160018114610175576000831561015d5750848201515b600019600385901b1c1916600184901b17845561010c565b600084815260208120601f198516915b828110156101a55787850151825560209485019460019092019101610185565b50848210156101c35786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610986806101e16000396000f3fe6080604052600436106100c05760003560e01c8063313ce56711610074578063a9059cbb1161004e578063a9059cbb146101fa578063d0e30db01461021a578063dd62ed3e1461022257600080fd5b8063313ce5671461018c57806370a08231146101b857806395d89b41146101e557600080fd5b806318160ddd116100a557806318160ddd1461012f57806323b872dd1461014c5780632e1a7d4d1461016c57600080fd5b806306fdde03146100d4578063095ea7b3146100ff57600080fd5b366100cf576100cd61025a565b005b600080fd5b3480156100e057600080fd5b506100e96102b5565b6040516100f69190610745565b60405180910390f35b34801561010b57600080fd5b5061011f61011a3660046107da565b610343565b60405190151581526020016100f6565b34801561013b57600080fd5b50475b6040519081526020016100f6565b34801561015857600080fd5b5061011f610167366004610804565b6103bd565b34801561017857600080fd5b506100cd610187366004610841565b610647565b34801561019857600080fd5b506002546101a69060ff1681565b60405160ff90911681526020016100f6565b3480156101c457600080fd5b5061013e6101d336600461085a565b60036020526000908152604090205481565b3480156101f157600080fd5b506100e9610724565b34801561020657600080fd5b5061011f6102153660046107da565b610731565b6100cd61025a565b34801561022e57600080fd5b5061013e61023d366004610875565b600460209081526000928352604080842090915290825290205481565b33600090815260036020526040812080543492906102799084906108d7565b909155505060405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2565b600080546102c2906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102ee906108ea565b801561033b5780601f106103105761010080835404028352916020019161033b565b820191906000526020600020905b81548152906001019060200180831161031e57829003601f168201915b505050505081565b33600081815260046020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906103ab9086815260200190565b60405180910390a35060015b92915050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081205482111561042b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600060248201526044015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841633148015906104a1575073ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091529020547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff14155b156105605773ffffffffffffffffffffffffffffffffffffffff8416600090815260046020908152604080832033845290915290205482111561051a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b73ffffffffffffffffffffffffffffffffffffffff841660009081526004602090815260408083203384529091528120805484929061055a90849061093d565b90915550505b73ffffffffffffffffffffffffffffffffffffffff84166000908152600360205260408120805484929061059590849061093d565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080548492906105cf9084906108d7565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161063591815260200190565b60405180910390a35060019392505050565b3360009081526003602052604090205481111561069a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526020600482015260006024820152604401610422565b33600090815260036020526040812080548392906106b990849061093d565b9091555050604051339082156108fc029083906000818181858888f193505050501580156106eb573d6000803e3d6000fd5b5060405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b659060200160405180910390a250565b600180546102c2906108ea565b600061073e3384846103bd565b9392505050565b602081526000825180602084015260005b818110156107735760208186018101516040868401015201610756565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107d557600080fd5b919050565b600080604083850312156107ed57600080fd5b6107f6836107b1565b946020939093013593505050565b60008060006060848603121561081957600080fd5b610822846107b1565b9250610830602085016107b1565b929592945050506040919091013590565b60006020828403121561085357600080fd5b5035919050565b60006020828403121561086c57600080fd5b61073e826107b1565b6000806040838503121561088857600080fd5b610891836107b1565b915061089f602084016107b1565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156103b7576103b76108a8565b600181811c908216806108fe57607f821691505b602082108103610937577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b818103818111156103b7576103b76108a856fea2646970667358221220bb5b9dcef0ba90bcdefcbd63f71b1df95b50e29550a7456c69c6b9ff9dcdd20e64736f6c634300081a00336080604052348015600f57600080fd5b506107978061001f6000396000f3fe6080604052600436106100355760003560e01c80632d4cfb7e1461003e5780635bcfd6161461005e578063c9028a361461007e57005b3661003c57005b005b34801561004a57600080fd5b5061003c610059366004610182565b61009e565b34801561006a57600080fd5b5061003c6100793660046101ed565b6100d8565b34801561008a57600080fd5b5061003c6100993660046102aa565b610153565b7f7d77d89ce7d68a5bf49d11b7fd7a992caa1c107a4c09b324ada48ee9c21b3db7816040516100cd9190610399565b60405180910390a150565b606081156100ef576100ec8284018461049f565b90505b7fcdc8ee677dc5ebe680fb18cebda5e26ba5ea1f0ba504a47e2a9a2ecb476dc98e61011a8780610595565b61012a60408a0160208b016105fa565b8960400135338660405161014396959493929190610615565b60405180910390a1505050505050565b7fd75bb509c8f32a725aac99ac5c4541060dbfb889a3aca8314d6f00395618c4c4816040516100cd91906106d7565b60006020828403121561019457600080fd5b813567ffffffffffffffff8111156101ab57600080fd5b820160c081850312156101bd57600080fd5b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146101e857600080fd5b919050565b60008060008060006080868803121561020557600080fd5b853567ffffffffffffffff81111561021c57600080fd5b86016060818903121561022e57600080fd5b945061023c602087016101c4565b935060408601359250606086013567ffffffffffffffff81111561025f57600080fd5b8601601f8101881361027057600080fd5b803567ffffffffffffffff81111561028757600080fd5b88602082840101111561029957600080fd5b959894975092955050506020019190565b6000602082840312156102bc57600080fd5b813567ffffffffffffffff8111156102d357600080fd5b8201608081850312156101bd57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261031a57600080fd5b830160208101925035905067ffffffffffffffff81111561033a57600080fd5b80360382131561034957600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6020815260006103a983846102e5565b60c060208501526103be60e085018284610350565b91505073ffffffffffffffffffffffffffffffffffffffff6103e2602086016101c4565b16604084015260006040850135905080606085015250606084013580151580821461040c57600080fd5b80608086015250506000608085013590508060a08501525061043160a08501856102e5565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08584030160c0860152610466838284610350565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156104b157600080fd5b813567ffffffffffffffff8111156104c857600080fd5b8201601f810184136104d957600080fd5b803567ffffffffffffffff8111156104f3576104f3610470565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8501160116810181811067ffffffffffffffff8211171561055f5761055f610470565b60405281815282820160200186101561057757600080fd5b81602084016020830137600091810160200191909152949350505050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126105ca57600080fd5b83018035915067ffffffffffffffff8211156105e557600080fd5b60200191503681900382131561034957600080fd5b60006020828403121561060c57600080fd5b6101bd826101c4565b60a08152600061062960a08301888a610350565b73ffffffffffffffffffffffffffffffffffffffff8716602084015285604084015273ffffffffffffffffffffffffffffffffffffffff851660608401528281036080840152835180825260005b8181101561069357602081870181015184830182015201610677565b5060006020828401015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011683010192505050979650505050505050565b6020815273ffffffffffffffffffffffffffffffffffffffff6106f9836101c4565b16602082015273ffffffffffffffffffffffffffffffffffffffff610720602084016101c4565b166040820152600080604084013590508060608401525061074460608401846102e5565b60808085015261075860a085018284610350565b9594505050505056fea2646970667358221220121f02e9d998ca4c5df318c6995b5a93c9477c42c6e34b1b54eafd621fd74daf64736f6c634300081a003360c060405234801561001057600080fd5b5060405161106f38038061106f83398101604081905261002f916100db565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461006357604051632b2add3d60e01b815260040160405180910390fd5b600380546001600160a01b0319166001600160a01b0385811691909117909155828116608052811660a0526040517f80699e81136d69cb8367ad52a994e25c722a86da654b561d0c14b61a777e7ac590600090a150505061011e565b80516001600160a01b03811681146100d657600080fd5b919050565b6000806000606084860312156100f057600080fd5b6100f9846100bf565b9250610107602085016100bf565b9150610115604085016100bf565b90509250925092565b60805160a051610f2561014a60003960006101e50152600081816102b9015261045b0152610f256000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806397770dff11610097578063c63585cc11610066578063c63585cc14610273578063d7fd7afb14610286578063d936a012146102b4578063ee2815ba146102db57600080fd5b806397770dff1461021a578063a7cb05071461022d578063c39aca3714610240578063c62178ac1461025357600080fd5b8063513a9c05116100d3578063513a9c051461018a578063569541b9146101c0578063842da36d146101e057806391dd645f1461020757600080fd5b80630be15547146100fa5780631f0e251b1461015a5780633ce4a5bc1461016f575b600080fd5b610130610108366004610bd1565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61016d610168366004610c13565b6102ee565b005b61013073735b14bb79463307aacbed86daf3322b1e6226ab81565b610130610198366004610bd1565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b6003546101309073ffffffffffffffffffffffffffffffffffffffff1681565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d610215366004610c35565b610402565b61016d610228366004610c13565b610526565b61016d61023b366004610c61565b610633565b61016d61024e366004610c83565b6106ce565b6004546101309073ffffffffffffffffffffffffffffffffffffffff1681565b610130610281366004610d53565b6108cd565b6102a6610294366004610bd1565b60006020819052908152604090205481565b604051908152602001610151565b6101307f000000000000000000000000000000000000000000000000000000000000000081565b61016d6102e9366004610c35565b610a02565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461033b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610388576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f3ade88e3922d64780e1bf4460d364c2970b69da813f9c0c07a1c187b5647636c906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461044f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600354600090610497907f00000000000000000000000000000000000000000000000000000000000000009073ffffffffffffffffffffffffffffffffffffffff16846108cd565b60008481526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251878152918201529192507f0ecec485166da6139b13bb7e033e9446e2d35348e80ebf1180d4afe2dba1704e910160405180910390a1505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610573576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166105c0576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fdba79d534382d1a8ae108e4c8ecb27c6ae42ab8b91d44eedf88bd329f3868d5e906020016103f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610680576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000828152602081815260409182902083905581518481529081018390527f49f492222906ac486c3c1401fa545626df1f0c0e5a77a05597ea2ed66af9850d91015b60405180910390a15050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461071b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff831673735b14bb79463307aacbed86daf3322b1e6226ab1480610768575073ffffffffffffffffffffffffffffffffffffffff831630145b1561079f576040517f82d5d76a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f47e7ef2400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018690528616906347e7ef24906044016020604051808303816000875af1158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190610d96565b506040517fde43156e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84169063de43156e906108939089908990899088908890600401610e01565b600060405180830381600087803b1580156108ad57600080fd5b505af11580156108c1573d6000803e3d6000fd5b50505050505050505050565b60008060006108dc8585610ad3565b6040517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000606084811b8216602084015283901b16603482015291935091508690604801604051602081830303815290604052805190602001206040516020016109c29291907fff00000000000000000000000000000000000000000000000000000000000000815260609290921b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600183015260158201527f96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f603582015260550190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209695505050505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610a4f576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526001602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091558251858152918201527fd1b36d30f6248e97c473b4d1348ca164a4ef6759022f54a58ec200326c39c45d91016106c2565b6000808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610b3b576040517fcb1e7cfe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1610610b75578284610b78565b83835b909250905073ffffffffffffffffffffffffffffffffffffffff8216610bca576040517f78b507da00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600060208284031215610be357600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c0e57600080fd5b919050565b600060208284031215610c2557600080fd5b610c2e82610bea565b9392505050565b60008060408385031215610c4857600080fd5b82359150610c5860208401610bea565b90509250929050565b60008060408385031215610c7457600080fd5b50508035926020909101359150565b60008060008060008060a08789031215610c9c57600080fd5b863567ffffffffffffffff811115610cb357600080fd5b87016060818a031215610cc557600080fd5b9550610cd360208801610bea565b945060408701359350610ce860608801610bea565b9250608087013567ffffffffffffffff811115610d0457600080fd5b8701601f81018913610d1557600080fd5b803567ffffffffffffffff811115610d2c57600080fd5b896020828401011115610d3e57600080fd5b60208201935080925050509295509295509295565b600080600060608486031215610d6857600080fd5b610d7184610bea565b9250610d7f60208501610bea565b9150610d8d60408501610bea565b90509250925092565b600060208284031215610da857600080fd5b81518015158114610c2e57600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60808152600086357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1883603018112610e3957600080fd5b870160208101903567ffffffffffffffff811115610e5657600080fd5b803603821315610e6557600080fd5b60606080850152610e7a60e085018284610db8565b91505073ffffffffffffffffffffffffffffffffffffffff610e9e60208a01610bea565b1660a0840152604088013560c084015273ffffffffffffffffffffffffffffffffffffffff871660208401528560408401528281036060840152610ee3818587610db8565b9897505050505050505056fea26469706673582212204dd385be01c68a79b7e2cc165deaf12752160a02e9f9b52233e4ef8c420d7b3c64736f6c634300081a003360c060405234801561001057600080fd5b5060405161207238038061207283398101604081905261002f916101f0565b6001600160a01b038216158061004c57506001600160a01b038116155b1561006a5760405163d92e233d60e01b815260040160405180910390fd5b60066100768982610342565b5060076100838882610342565b506008805460ff191660ff881617905560808590528360028111156100aa576100aa610400565b60a08160028111156100be576100be610400565b905250600192909255600080546001600160a01b039283166001600160a01b0319909116179055600880549190921661010002610100600160a81b0319909116179055506104169350505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261013357600080fd5b81516001600160401b0381111561014c5761014c61010c565b604051601f8201601f19908116603f011681016001600160401b038111828210171561017a5761017a61010c565b60405281815283820160200185101561019257600080fd5b60005b828110156101b157602081860181015183830182015201610195565b506000918101602001919091529392505050565b8051600381106101d457600080fd5b919050565b80516001600160a01b03811681146101d457600080fd5b600080600080600080600080610100898b03121561020d57600080fd5b88516001600160401b0381111561022357600080fd5b61022f8b828c01610122565b60208b015190995090506001600160401b0381111561024d57600080fd5b6102598b828c01610122565b975050604089015160ff8116811461027057600080fd5b60608a0151909650945061028660808a016101c5565b60a08a0151909450925061029c60c08a016101d9565b91506102aa60e08a016101d9565b90509295985092959890939650565b600181811c908216806102cd57607f821691505b6020821081036102ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561033d57806000526020600020601f840160051c8101602085101561031a5750805b601f840160051c820191505b8181101561033a5760008155600101610326565b50505b505050565b81516001600160401b0381111561035b5761035b61010c565b61036f8161036984546102b9565b846102f3565b6020601f8211600181146103a3576000831561038b5750848201515b600019600385901b1c1916600184901b17845561033a565b600084815260208120601f198516915b828110156103d357878501518255602094850194600190920191016103b3565b50848210156103f15786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052602160045260246000fd5b60805160a051611c1b61045760003960006103440152600081816102f001528181610bdc01528181610ce201528181610efe01526110040152611c1b6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c806395d89b41116100f9578063ccc7759911610097578063eddeb12311610071578063eddeb12314610461578063f2441b3214610474578063f687d12a14610494578063fc5fecd5146104a757600080fd5b8063ccc77599146103d4578063d9eeebed146103e7578063dd62ed3e1461041b57600080fd5b8063b84c8246116100d3578063b84c824614610386578063c47f00271461039b578063c7012626146103ae578063c835d7cc146103c157600080fd5b806395d89b4114610337578063a3413d031461033f578063a9059cbb1461037357600080fd5b80633ce4a5bc116101665780634d8943bb116101405780634d8943bb146102ac57806370a08231146102b557806385e1f4d0146102eb5780638b851b951461031257600080fd5b80633ce4a5bc1461024657806342966c681461028657806347e7ef241461029957600080fd5b806318160ddd1161019757806318160ddd1461021657806323b872dd1461021e578063313ce5671461023157600080fd5b806306fdde03146101be578063091d2788146101dc578063095ea7b3146101f3575b600080fd5b6101c66104ba565b6040516101d39190611648565b60405180910390f35b6101e560015481565b6040519081526020016101d3565b610206610201366004611687565b61054c565b60405190151581526020016101d3565b6005546101e5565b61020661022c3660046116b3565b610563565b60085460405160ff90911681526020016101d3565b61026173735b14bb79463307aacbed86daf3322b1e6226ab81565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b6102066102943660046116f4565b6105fa565b6102066102a7366004611687565b61060e565b6101e560025481565b6101e56102c336600461170d565b73ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6101e57f000000000000000000000000000000000000000000000000000000000000000081565b60085461026190610100900473ffffffffffffffffffffffffffffffffffffffff1681565b6101c6610767565b6103667f000000000000000000000000000000000000000000000000000000000000000081565b6040516101d3919061172a565b610206610381366004611687565b610776565b610399610394366004611832565b610783565b005b6103996103a9366004611832565b6107e0565b6102066103bc366004611883565b610839565b6103996103cf36600461170d565b610988565b6103996103e236600461170d565b610a9c565b6103ef610bb0565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526020830191909152016101d3565b6101e56104293660046118dc565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260046020908152604080832093909416825291909152205490565b61039961046f3660046116f4565b610dce565b6000546102619073ffffffffffffffffffffffffffffffffffffffff1681565b6103996104a23660046116f4565b610e50565b6103ef6104b53660046116f4565b610ed2565b6060600680546104c990611915565b80601f01602080910402602001604051908101604052809291908181526020018280546104f590611915565b80156105425780601f1061051757610100808354040283529160200191610542565b820191906000526020600020905b81548152906001019060200180831161052557829003601f168201915b5050505050905090565b60006105593384846110ee565b5060015b92915050565b60006105708484846111f7565b73ffffffffffffffffffffffffffffffffffffffff84166000908152600460209081526040808320338452909152902054828110156105db576040517f10bad14700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6105ef85336105ea8685611997565b6110ee565b506001949350505050565b600061060633836113b2565b506001919050565b60003373735b14bb79463307aacbed86daf3322b1e6226ab1480159061064c575060005473ffffffffffffffffffffffffffffffffffffffff163314155b80156106755750600854610100900473ffffffffffffffffffffffffffffffffffffffff163314155b156106ac576040517fddb5de5e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106b683836114f4565b6040517f735b14bb79463307aacbed86daf3322b1e6226ab000000000000000000000000602082015273ffffffffffffffffffffffffffffffffffffffff8416907f67fc7bdaed5b0ec550d8706b87d60568ab70c6b781263c70101d54cd1564aab390603401604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526107569186906119aa565b60405180910390a250600192915050565b6060600780546104c990611915565b60006105593384846111f7565b3373735b14bb79463307aacbed86daf3322b1e6226ab146107d0576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076107dc8282611a1b565b5050565b3373735b14bb79463307aacbed86daf3322b1e6226ab1461082d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60066107dc8282611a1b565b6000806000610846610bb0565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015273735b14bb79463307aacbed86daf3322b1e6226ab602482015260448101829052919350915073ffffffffffffffffffffffffffffffffffffffff8316906323b872dd906064016020604051808303816000875af11580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc9190611b34565b610932576040517f0a7cd6d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61093c33856113b2565b60025460405133917f9ffbffc04a397460ee1dbe8c9503e098090567d6b7f4b3c02a8617d800b6d9559161097591899189918791611b56565b60405180910390a2506001949350505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab146109d5576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610a22576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527fd55614e962c5fd6ece71614f6348d702468a997a394dd5e5c1677950226d97ae906020015b60405180910390a150565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610ae9576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b36576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffff0000000000000000000000000000000000000000ff1661010073ffffffffffffffffffffffffffffffffffffffff8416908102919091179091556040519081527f88815d964e380677e86d817e7d65dea59cb7b4c3b5b7a0c8ec7ea4a74f90a38790602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610c43573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c679190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610cb6576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015610d45573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d699190611ba2565b905080600003610da5576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060025460015483610db89190611bbb565b610dc29190611bd2565b92959294509192505050565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e1b576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028190556040518181527fef13af88e424b5d15f49c77758542c1938b08b8b95b91ed0751f98ba99000d8f90602001610a91565b3373735b14bb79463307aacbed86daf3322b1e6226ab14610e9d576040517f2b2add3d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018190556040518181527fff5788270f43bfc1ca41c503606d2594aa3023a1a7547de403a3e2f146a4a80a90602001610a91565b600080546040517f0be155470000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060048201528291829173ffffffffffffffffffffffffffffffffffffffff90911690630be1554790602401602060405180830381865afa158015610f65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f899190611b85565b905073ffffffffffffffffffffffffffffffffffffffff8116610fd8576040517f78fff39600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080546040517fd7fd7afb0000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000600482015273ffffffffffffffffffffffffffffffffffffffff9091169063d7fd7afb90602401602060405180830381865afa158015611067573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108b9190611ba2565b9050806000036110c7576040517fe661aed000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000906110d78784611bbb565b6110e19190611bd2565b9296929550919350505050565b73ffffffffffffffffffffffffffffffffffffffff831661113b576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611188576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff83811660008181526004602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316611244576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216611291576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040902054818110156112f1576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112fb8282611997565b73ffffffffffffffffffffffffffffffffffffffff808616600090815260036020526040808220939093559085168152908120805484929061133e908490611bd2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516113a491815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff82166113ff576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260409020548181101561145f576040517ffe382aa700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114698282611997565b73ffffffffffffffffffffffffffffffffffffffff8416600090815260036020526040812091909155600580548492906114a4908490611997565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020016111ea565b73ffffffffffffffffffffffffffffffffffffffff8216611541576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600560008282546115539190611bd2565b909155505073ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805483929061158d908490611bd2565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b6000815180845260005b8181101561160a576020818501810151868301820152016115ee565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061165b60208301846115e4565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461168457600080fd5b50565b6000806040838503121561169a57600080fd5b82356116a581611662565b946020939093013593505050565b6000806000606084860312156116c857600080fd5b83356116d381611662565b925060208401356116e381611662565b929592945050506040919091013590565b60006020828403121561170657600080fd5b5035919050565b60006020828403121561171f57600080fd5b813561165b81611662565b6020810160038310611765577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008067ffffffffffffffff8411156117b5576117b561176b565b506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f85018116603f0116810181811067ffffffffffffffff821117156118025761180261176b565b60405283815290508082840185101561181a57600080fd5b83836020830137600060208583010152509392505050565b60006020828403121561184457600080fd5b813567ffffffffffffffff81111561185b57600080fd5b8201601f8101841361186c57600080fd5b61187b8482356020840161179a565b949350505050565b6000806040838503121561189657600080fd5b823567ffffffffffffffff8111156118ad57600080fd5b8301601f810185136118be57600080fd5b6118cd8582356020840161179a565b95602094909401359450505050565b600080604083850312156118ef57600080fd5b82356118fa81611662565b9150602083013561190a81611662565b809150509250929050565b600181811c9082168061192957607f821691505b602082108103611962577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561055d5761055d611968565b6040815260006119bd60408301856115e4565b90508260208301529392505050565b601f821115611a1657806000526020600020601f840160051c810160208510156119f35750805b601f840160051c820191505b81811015611a1357600081556001016119ff565b50505b505050565b815167ffffffffffffffff811115611a3557611a3561176b565b611a4981611a438454611915565b846119cc565b6020601f821160018114611a9b5760008315611a655750848201515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600385901b1c1916600184901b178455611a13565b6000848152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08516915b82811015611ae95787850151825560209485019460019092019101611ac9565b5084821015611b2557868401517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600387901b60f8161c191681555b50505050600190811b01905550565b600060208284031215611b4657600080fd5b8151801515811461165b57600080fd5b608081526000611b6960808301876115e4565b6020830195909552506040810192909252606090910152919050565b600060208284031215611b9757600080fd5b815161165b81611662565b600060208284031215611bb457600080fd5b5051919050565b808202811582820484141761055d5761055d611968565b8082018082111561055d5761055d61196856fea26469706673582212208307d60e253f5034856b93df00e3e2f46b06f9765d906dbd93ee947935fc608764736f6c634300081a0033a264697066735822122038f3eebdef427300426bcddf5020778554c5a0d2965348857f3e4b934f466e4d64736f6c634300081a0033",
}

// GatewayMEVMInboundTestABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMEVMInboundTestMetaData.ABI instead.
var GatewayMEVMInboundTestABI = GatewayMEVMInboundTestMetaData.ABI

// GatewayMEVMInboundTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMEVMInboundTestMetaData.Bin instead.
var GatewayMEVMInboundTestBin = GatewayMEVMInboundTestMetaData.Bin

// DeployGatewayMEVMInboundTest deploys a new Ethereum contract, binding an instance of GatewayMEVMInboundTest to it.
func DeployGatewayMEVMInboundTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayMEVMInboundTest, error) {
	parsed, err := GatewayMEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayMEVMInboundTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayMEVMInboundTest{GatewayMEVMInboundTestCaller: GatewayMEVMInboundTestCaller{contract: contract}, GatewayMEVMInboundTestTransactor: GatewayMEVMInboundTestTransactor{contract: contract}, GatewayMEVMInboundTestFilterer: GatewayMEVMInboundTestFilterer{contract: contract}}, nil
}

// GatewayMEVMInboundTest is an auto generated Go binding around an Ethereum contract.
type GatewayMEVMInboundTest struct {
	GatewayMEVMInboundTestCaller     // Read-only binding to the contract
	GatewayMEVMInboundTestTransactor // Write-only binding to the contract
	GatewayMEVMInboundTestFilterer   // Log filterer for contract events
}

// GatewayMEVMInboundTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayMEVMInboundTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMInboundTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayMEVMInboundTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMInboundTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayMEVMInboundTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayMEVMInboundTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayMEVMInboundTestSession struct {
	Contract     *GatewayMEVMInboundTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GatewayMEVMInboundTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayMEVMInboundTestCallerSession struct {
	Contract *GatewayMEVMInboundTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GatewayMEVMInboundTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayMEVMInboundTestTransactorSession struct {
	Contract     *GatewayMEVMInboundTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GatewayMEVMInboundTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayMEVMInboundTestRaw struct {
	Contract *GatewayMEVMInboundTest // Generic contract binding to access the raw methods on
}

// GatewayMEVMInboundTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayMEVMInboundTestCallerRaw struct {
	Contract *GatewayMEVMInboundTestCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayMEVMInboundTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayMEVMInboundTestTransactorRaw struct {
	Contract *GatewayMEVMInboundTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayMEVMInboundTest creates a new instance of GatewayMEVMInboundTest, bound to a specific deployed contract.
func NewGatewayMEVMInboundTest(address common.Address, backend bind.ContractBackend) (*GatewayMEVMInboundTest, error) {
	contract, err := bindGatewayMEVMInboundTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTest{GatewayMEVMInboundTestCaller: GatewayMEVMInboundTestCaller{contract: contract}, GatewayMEVMInboundTestTransactor: GatewayMEVMInboundTestTransactor{contract: contract}, GatewayMEVMInboundTestFilterer: GatewayMEVMInboundTestFilterer{contract: contract}}, nil
}

// NewGatewayMEVMInboundTestCaller creates a new read-only instance of GatewayMEVMInboundTest, bound to a specific deployed contract.
func NewGatewayMEVMInboundTestCaller(address common.Address, caller bind.ContractCaller) (*GatewayMEVMInboundTestCaller, error) {
	contract, err := bindGatewayMEVMInboundTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestCaller{contract: contract}, nil
}

// NewGatewayMEVMInboundTestTransactor creates a new write-only instance of GatewayMEVMInboundTest, bound to a specific deployed contract.
func NewGatewayMEVMInboundTestTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayMEVMInboundTestTransactor, error) {
	contract, err := bindGatewayMEVMInboundTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestTransactor{contract: contract}, nil
}

// NewGatewayMEVMInboundTestFilterer creates a new log filterer instance of GatewayMEVMInboundTest, bound to a specific deployed contract.
func NewGatewayMEVMInboundTestFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayMEVMInboundTestFilterer, error) {
	contract, err := bindGatewayMEVMInboundTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestFilterer{contract: contract}, nil
}

// bindGatewayMEVMInboundTest binds a generic wrapper to an already deployed contract.
func bindGatewayMEVMInboundTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMEVMInboundTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMEVMInboundTest.Contract.GatewayMEVMInboundTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.GatewayMEVMInboundTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.GatewayMEVMInboundTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayMEVMInboundTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) ISTEST() (bool, error) {
	return _GatewayMEVMInboundTest.Contract.ISTEST(&_GatewayMEVMInboundTest.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) ISTEST() (bool, error) {
	return _GatewayMEVMInboundTest.Contract.ISTEST(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) ExcludeArtifacts() ([]string, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeArtifacts(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeContracts(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeContracts(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeSenders(&_GatewayMEVMInboundTest.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.ExcludeSenders(&_GatewayMEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) Failed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "failed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) Failed() (bool, error) {
	return _GatewayMEVMInboundTest.Contract.Failed(&_GatewayMEVMInboundTest.CallOpts)
}

// Failed is a free data retrieval call binding the contract method 0xba414fa6.
//
// Solidity: function failed() view returns(bool)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) Failed() (bool, error) {
	return _GatewayMEVMInboundTest.Contract.Failed(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayMEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _GatewayMEVMInboundTest.Contract.TargetArtifactSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetArtifacts() ([]string, error) {
	return _GatewayMEVMInboundTest.Contract.TargetArtifacts(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetArtifacts() ([]string, error) {
	return _GatewayMEVMInboundTest.Contract.TargetArtifacts(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetContracts() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.TargetContracts(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetContracts() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.TargetContracts(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayMEVMInboundTest.Contract.TargetInterfaces(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _GatewayMEVMInboundTest.Contract.TargetInterfaces(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayMEVMInboundTest.Contract.TargetSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _GatewayMEVMInboundTest.Contract.TargetSelectors(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GatewayMEVMInboundTest.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TargetSenders() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.TargetSenders(&_GatewayMEVMInboundTest.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestCallerSession) TargetSenders() ([]common.Address, error) {
	return _GatewayMEVMInboundTest.Contract.TargetSenders(&_GatewayMEVMInboundTest.CallOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) SetUp() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.SetUp(&_GatewayMEVMInboundTest.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) SetUp() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.SetUp(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestCallWithCallOpts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testCallWithCallOpts")
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOpts is a paid mutator transaction binding the contract method 0x6198fb19.
//
// Solidity: function testCallWithCallOpts() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestCallWithCallOpts() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOpts(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsBelowMin")
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0x05b9f046.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsBelowMin(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestCallWithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfGasLimitIsZero")
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x5efe72a9.
//
// Solidity: function testCallWithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestCallWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfMessageSizeExceeded")
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x6dfcbc50.
//
// Solidity: function testCallWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestCallWithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testCallWithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestCallWithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xb51ac071.
//
// Solidity: function testCallWithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestCallWithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestCallWithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawMRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestUpgradeAndWithdrawMRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testUpgradeAndWithdrawMRC20")
}

// TestUpgradeAndWithdrawMRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestUpgradeAndWithdrawMRC20() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestUpgradeAndWithdrawMRC20(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestUpgradeAndWithdrawMRC20 is a paid mutator transaction binding the contract method 0x20dee15f.
//
// Solidity: function testUpgradeAndWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestUpgradeAndWithdrawMRC20() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestUpgradeAndWithdrawMRC20(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress")
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x43181437.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x42752d41.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xc946d7c0.
//
// Solidity: function testWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMUSEWithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20FailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20FailsIfAmountIsZero")
}

// TestWithdrawAndCallMRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfAmountIsZero is a paid mutator transaction binding the contract method 0xba9adeef.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20FailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin")
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin is a paid mutator transaction binding the contract method 0xc20049f4.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfGasLimitIsBelowMin(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20FailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20FailsIfGasLimitIsZero")
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0xbed3e813.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20FailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x4ffab9de.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x36431b3f.
//
// Solidity: function testWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20FailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero")
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x1b9641bf.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero")
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x0b5ad28d.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded")
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0xa721b2d3.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress")
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xeb7a2fac.
//
// Solidity: function testWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawAndCallMRC20WithCallOptsFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSE is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawMUSE() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSE")
}

// TestWithdrawMUSE is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawMUSE() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSE() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSE(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSE is a paid mutator transaction binding the contract method 0xea37902f.
//
// Solidity: function testWithdrawMUSE() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSE() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSE(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawMUSEFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEFailsIfAmountIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEFailsIfAmountIsZero")
}

// TestWithdrawMUSEFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawMUSEFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfAmountIsZero is a paid mutator transaction binding the contract method 0x21aeb18c.
//
// Solidity: function testWithdrawMUSEFailsIfAmountIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEFailsIfAmountIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfAmountIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawMUSEFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEFailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEFailsIfMessageSizeExceeded")
}

// TestWithdrawMUSEFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawMUSEFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x7ae69730.
//
// Solidity: function testWithdrawMUSEFailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEFailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawMUSEFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEFailsIfNoAllowance")
}

// TestWithdrawMUSEFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawMUSEFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xdde7e967.
//
// Solidity: function testWithdrawMUSEFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawMUSEFailsIfNoBalance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEFailsIfNoBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEFailsIfNoBalance")
}

// TestWithdrawMUSEFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawMUSEFailsIfNoBalance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfNoBalance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfNoBalance is a paid mutator transaction binding the contract method 0x6221b509.
//
// Solidity: function testWithdrawMUSEFailsIfNoBalance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEFailsIfNoBalance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfNoBalance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawMUSEFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEFailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEFailsIfReceiverIsZeroAddress")
}

// TestWithdrawMUSEFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawMUSEFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEFailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0xdc749dd7.
//
// Solidity: function testWithdrawMUSEFailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEFailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEFailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEWithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEWithCallOptsWithMessage")
}

// TestWithdrawMUSEWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessage is a paid mutator transaction binding the contract method 0x14759766.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEWithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero")
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero is a paid mutator transaction binding the contract method 0x64002a1f.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessageFailsIfGasLimitIsZero(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance")
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0xe51c6388.
//
// Solidity: function testWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMUSEWithCallOptsWithMessageFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20")
}

// TestWithdrawMRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20 is a paid mutator transaction binding the contract method 0xfbc611c8.
//
// Solidity: function testWithdrawMRC20() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIfMessageSizeExceeded(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIfMessageSizeExceeded")
}

// TestWithdrawMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfMessageSizeExceeded is a paid mutator transaction binding the contract method 0x264b524c.
//
// Solidity: function testWithdrawMRC20FailsIfMessageSizeExceeded() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIfMessageSizeExceeded() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfMessageSizeExceeded(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawMRC20FailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIfNoAllowance")
}

// TestWithdrawMRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawMRC20FailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoAllowance is a paid mutator transaction binding the contract method 0x5d72228f.
//
// Solidity: function testWithdrawMRC20FailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIfNoBalanceForGasFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIfNoBalanceForGasFee")
}

// TestWithdrawMRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoBalanceForGasFee(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoBalanceForGasFee is a paid mutator transaction binding the contract method 0xceccfab3.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForGasFee() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIfNoBalanceForGasFee() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoBalanceForGasFee(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIfNoBalanceForTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIfNoBalanceForTransfer")
}

// TestWithdrawMRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoBalanceForTransfer(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfNoBalanceForTransfer is a paid mutator transaction binding the contract method 0xfdad0ad0.
//
// Solidity: function testWithdrawMRC20FailsIfNoBalanceForTransfer() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIfNoBalanceForTransfer() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfNoBalanceForTransfer(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIfReceiverIsZeroAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIfReceiverIsZeroAddress")
}

// TestWithdrawMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIfReceiverIsZeroAddress is a paid mutator transaction binding the contract method 0x83ababa9.
//
// Solidity: function testWithdrawMRC20FailsIfReceiverIsZeroAddress() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIfReceiverIsZeroAddress() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIfReceiverIsZeroAddress(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawMRC20FailsIsAmountIs0() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20FailsIsAmountIs0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20FailsIsAmountIs0")
}

// TestWithdrawMRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawMRC20FailsIsAmountIs0() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIsAmountIs0(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20FailsIsAmountIs0 is a paid mutator transaction binding the contract method 0xba800c91.
//
// Solidity: function testWithdrawMRC20FailsIsAmountIs0() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20FailsIsAmountIs0() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20FailsIsAmountIs0(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawMRC20WithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20WithCallOptsWithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20WithCallOptsWithMessage")
}

// TestWithdrawMRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawMRC20WithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithCallOptsWithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithCallOptsWithMessage is a paid mutator transaction binding the contract method 0xe804a406.
//
// Solidity: function testWithdrawMRC20WithCallOptsWithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20WithCallOptsWithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithCallOptsWithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawMRC20WithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20WithMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20WithMessage")
}

// TestWithdrawMRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawMRC20WithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20WithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessage is a paid mutator transaction binding the contract method 0x1e63d2b9.
//
// Solidity: function testWithdrawMRC20WithMessage() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20WithMessage() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessage(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawMRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20WithMessageFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20WithMessageFailsIfNoAllowance")
}

// TestWithdrawMRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawMRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessageFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessageFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x1238212c.
//
// Solidity: function testWithdrawMRC20WithMessageFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20WithMessageFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessageFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactor) TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.contract.Transact(opts, "testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance")
}

// TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestSession) TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance is a paid mutator transaction binding the contract method 0x6d6ce0d0.
//
// Solidity: function testWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() returns()
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestTransactorSession) TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance() (*types.Transaction, error) {
	return _GatewayMEVMInboundTest.Contract.TestWithdrawMRC20WithMessageWithCallOptsFailsIfNoAllowance(&_GatewayMEVMInboundTest.TransactOpts)
}

// GatewayMEVMInboundTestCalledIterator is returned from FilterCalled and is used to iterate over the raw logs and unpacked data for Called events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestCalledIterator struct {
	Event *GatewayMEVMInboundTestCalled // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestCalled)
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
		it.Event = new(GatewayMEVMInboundTestCalled)
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
func (it *GatewayMEVMInboundTestCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestCalled represents a Called event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestCalled struct {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterCalled(opts *bind.FilterOpts, sender []common.Address, mrc20 []common.Address) (*GatewayMEVMInboundTestCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestCalledIterator{contract: _GatewayMEVMInboundTest.contract, event: "Called", logs: logs, sub: sub}, nil
}

// WatchCalled is a free log subscription operation binding the contract event 0x306ee13f48319a123b222c69908e44dcf91abffc20cacc502e3cf5a4ff23e0e4.
//
// Solidity: event Called(address indexed sender, address indexed mrc20, bytes receiver, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchCalled(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestCalled, sender []common.Address, mrc20 []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var mrc20Rule []interface{}
	for _, mrc20Item := range mrc20 {
		mrc20Rule = append(mrc20Rule, mrc20Item)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "Called", senderRule, mrc20Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestCalled)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseCalled(log types.Log) (*GatewayMEVMInboundTestCalled, error) {
	event := new(GatewayMEVMInboundTestCalled)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "Called", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawnIterator struct {
	Event *GatewayMEVMInboundTestWithdrawn // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestWithdrawn)
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
		it.Event = new(GatewayMEVMInboundTestWithdrawn)
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
func (it *GatewayMEVMInboundTestWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestWithdrawn represents a Withdrawn event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawn struct {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterWithdrawn(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMInboundTestWithdrawnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestWithdrawnIterator{contract: _GatewayMEVMInboundTest.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x07bf64173efd8f3dfb9e4eb3834bab9d5b85a3d89a1c6425797329de0668502c.
//
// Solidity: event Withdrawn(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestWithdrawn, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "Withdrawn", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestWithdrawn)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseWithdrawn(log types.Log) (*GatewayMEVMInboundTestWithdrawn, error) {
	event := new(GatewayMEVMInboundTestWithdrawn)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestWithdrawnAndCalledIterator is returned from FilterWithdrawnAndCalled and is used to iterate over the raw logs and unpacked data for WithdrawnAndCalled events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawnAndCalledIterator struct {
	Event *GatewayMEVMInboundTestWithdrawnAndCalled // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestWithdrawnAndCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestWithdrawnAndCalled)
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
		it.Event = new(GatewayMEVMInboundTestWithdrawnAndCalled)
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
func (it *GatewayMEVMInboundTestWithdrawnAndCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestWithdrawnAndCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestWithdrawnAndCalled represents a WithdrawnAndCalled event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawnAndCalled struct {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterWithdrawnAndCalled(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMInboundTestWithdrawnAndCalledIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestWithdrawnAndCalledIterator{contract: _GatewayMEVMInboundTest.contract, event: "WithdrawnAndCalled", logs: logs, sub: sub}, nil
}

// WatchWithdrawnAndCalled is a free log subscription operation binding the contract event 0xd90f94752d2b12f364f4a2237ebe1aff24ba6127585376bf4935f6a7be17dd2a.
//
// Solidity: event WithdrawnAndCalled(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchWithdrawnAndCalled(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestWithdrawnAndCalled, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "WithdrawnAndCalled", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestWithdrawnAndCalled)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseWithdrawnAndCalled(log types.Log) (*GatewayMEVMInboundTestWithdrawnAndCalled, error) {
	event := new(GatewayMEVMInboundTestWithdrawnAndCalled)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "WithdrawnAndCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestWithdrawnV2Iterator is returned from FilterWithdrawnV2 and is used to iterate over the raw logs and unpacked data for WithdrawnV2 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawnV2Iterator struct {
	Event *GatewayMEVMInboundTestWithdrawnV2 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestWithdrawnV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestWithdrawnV2)
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
		it.Event = new(GatewayMEVMInboundTestWithdrawnV2)
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
func (it *GatewayMEVMInboundTestWithdrawnV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestWithdrawnV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestWithdrawnV2 represents a WithdrawnV2 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestWithdrawnV2 struct {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterWithdrawnV2(opts *bind.FilterOpts, sender []common.Address, chainId []*big.Int) (*GatewayMEVMInboundTestWithdrawnV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestWithdrawnV2Iterator{contract: _GatewayMEVMInboundTest.contract, event: "WithdrawnV2", logs: logs, sub: sub}, nil
}

// WatchWithdrawnV2 is a free log subscription operation binding the contract event 0x5d7cd8ae449a6b25de63f10534ddd17d8dd3e79c7aa5f28964b7a7c760258d97.
//
// Solidity: event WithdrawnV2(address indexed sender, uint256 indexed chainId, bytes receiver, address mrc20, uint256 value, uint256 gasfee, uint256 protocolFlatFee, bytes message, (uint256,bool) callOptions, (address,bool,address,bytes,uint256) revertOptions)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchWithdrawnV2(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestWithdrawnV2, sender []common.Address, chainId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "WithdrawnV2", senderRule, chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestWithdrawnV2)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
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
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseWithdrawnV2(log types.Log) (*GatewayMEVMInboundTestWithdrawnV2, error) {
	event := new(GatewayMEVMInboundTestWithdrawnV2)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "WithdrawnV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogIterator struct {
	Event *GatewayMEVMInboundTestLog // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLog)
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
		it.Event = new(GatewayMEVMInboundTestLog)
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
func (it *GatewayMEVMInboundTestLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLog represents a Log event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLog(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogIterator{contract: _GatewayMEVMInboundTest.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLog) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLog)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLog(log types.Log) (*GatewayMEVMInboundTestLog, error) {
	event := new(GatewayMEVMInboundTestLog)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogAddressIterator struct {
	Event *GatewayMEVMInboundTestLogAddress // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogAddress)
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
		it.Event = new(GatewayMEVMInboundTestLogAddress)
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
func (it *GatewayMEVMInboundTestLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogAddress represents a LogAddress event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogAddress(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogAddressIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogAddressIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogAddress)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogAddress(log types.Log) (*GatewayMEVMInboundTestLogAddress, error) {
	event := new(GatewayMEVMInboundTestLogAddress)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArrayIterator struct {
	Event *GatewayMEVMInboundTestLogArray // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogArray)
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
		it.Event = new(GatewayMEVMInboundTestLogArray)
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
func (it *GatewayMEVMInboundTestLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogArray represents a LogArray event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogArray(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogArrayIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogArrayIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogArray) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogArray)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogArray(log types.Log) (*GatewayMEVMInboundTestLogArray, error) {
	event := new(GatewayMEVMInboundTestLogArray)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArray0Iterator struct {
	Event *GatewayMEVMInboundTestLogArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogArray0)
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
		it.Event = new(GatewayMEVMInboundTestLogArray0)
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
func (it *GatewayMEVMInboundTestLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogArray0 represents a LogArray0 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogArray0(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogArray0Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogArray0Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogArray0)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogArray0(log types.Log) (*GatewayMEVMInboundTestLogArray0, error) {
	event := new(GatewayMEVMInboundTestLogArray0)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArray1Iterator struct {
	Event *GatewayMEVMInboundTestLogArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogArray1)
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
		it.Event = new(GatewayMEVMInboundTestLogArray1)
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
func (it *GatewayMEVMInboundTestLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogArray1 represents a LogArray1 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogArray1(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogArray1Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogArray1Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogArray1)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogArray1(log types.Log) (*GatewayMEVMInboundTestLogArray1, error) {
	event := new(GatewayMEVMInboundTestLogArray1)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogBytesIterator struct {
	Event *GatewayMEVMInboundTestLogBytes // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogBytes)
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
		it.Event = new(GatewayMEVMInboundTestLogBytes)
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
func (it *GatewayMEVMInboundTestLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogBytes represents a LogBytes event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogBytes(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogBytesIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogBytesIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogBytes)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogBytes(log types.Log) (*GatewayMEVMInboundTestLogBytes, error) {
	event := new(GatewayMEVMInboundTestLogBytes)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogBytes32Iterator struct {
	Event *GatewayMEVMInboundTestLogBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogBytes32)
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
		it.Event = new(GatewayMEVMInboundTestLogBytes32)
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
func (it *GatewayMEVMInboundTestLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogBytes32 represents a LogBytes32 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogBytes32Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogBytes32Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogBytes32)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogBytes32(log types.Log) (*GatewayMEVMInboundTestLogBytes32, error) {
	event := new(GatewayMEVMInboundTestLogBytes32)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogIntIterator struct {
	Event *GatewayMEVMInboundTestLogInt // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogInt)
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
		it.Event = new(GatewayMEVMInboundTestLogInt)
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
func (it *GatewayMEVMInboundTestLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogInt represents a LogInt event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogInt(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogIntIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogIntIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogInt) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogInt)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogInt(log types.Log) (*GatewayMEVMInboundTestLogInt, error) {
	event := new(GatewayMEVMInboundTestLogInt)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedAddressIterator struct {
	Event *GatewayMEVMInboundTestLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedAddress)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedAddress)
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
func (it *GatewayMEVMInboundTestLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedAddress represents a LogNamedAddress event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedAddressIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedAddressIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedAddress)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedAddress(log types.Log) (*GatewayMEVMInboundTestLogNamedAddress, error) {
	event := new(GatewayMEVMInboundTestLogNamedAddress)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArrayIterator struct {
	Event *GatewayMEVMInboundTestLogNamedArray // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedArray)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedArray)
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
func (it *GatewayMEVMInboundTestLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedArray represents a LogNamedArray event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedArrayIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedArrayIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedArray)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedArray(log types.Log) (*GatewayMEVMInboundTestLogNamedArray, error) {
	event := new(GatewayMEVMInboundTestLogNamedArray)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArray0Iterator struct {
	Event *GatewayMEVMInboundTestLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedArray0)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedArray0)
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
func (it *GatewayMEVMInboundTestLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedArray0 represents a LogNamedArray0 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedArray0Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedArray0Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedArray0)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedArray0(log types.Log) (*GatewayMEVMInboundTestLogNamedArray0, error) {
	event := new(GatewayMEVMInboundTestLogNamedArray0)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArray1Iterator struct {
	Event *GatewayMEVMInboundTestLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedArray1)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedArray1)
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
func (it *GatewayMEVMInboundTestLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedArray1 represents a LogNamedArray1 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedArray1Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedArray1Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedArray1)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedArray1(log types.Log) (*GatewayMEVMInboundTestLogNamedArray1, error) {
	event := new(GatewayMEVMInboundTestLogNamedArray1)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedBytesIterator struct {
	Event *GatewayMEVMInboundTestLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedBytes)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedBytes)
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
func (it *GatewayMEVMInboundTestLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedBytes represents a LogNamedBytes event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedBytesIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedBytesIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedBytes)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedBytes(log types.Log) (*GatewayMEVMInboundTestLogNamedBytes, error) {
	event := new(GatewayMEVMInboundTestLogNamedBytes)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedBytes32Iterator struct {
	Event *GatewayMEVMInboundTestLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedBytes32)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedBytes32)
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
func (it *GatewayMEVMInboundTestLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedBytes32 represents a LogNamedBytes32 event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedBytes32Iterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedBytes32Iterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedBytes32)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedBytes32(log types.Log) (*GatewayMEVMInboundTestLogNamedBytes32, error) {
	event := new(GatewayMEVMInboundTestLogNamedBytes32)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedDecimalIntIterator struct {
	Event *GatewayMEVMInboundTestLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedDecimalInt)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedDecimalInt)
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
func (it *GatewayMEVMInboundTestLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedDecimalIntIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedDecimalIntIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedDecimalInt)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedDecimalInt(log types.Log) (*GatewayMEVMInboundTestLogNamedDecimalInt, error) {
	event := new(GatewayMEVMInboundTestLogNamedDecimalInt)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedDecimalUintIterator struct {
	Event *GatewayMEVMInboundTestLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedDecimalUint)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedDecimalUint)
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
func (it *GatewayMEVMInboundTestLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedDecimalUintIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedDecimalUintIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedDecimalUint)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedDecimalUint(log types.Log) (*GatewayMEVMInboundTestLogNamedDecimalUint, error) {
	event := new(GatewayMEVMInboundTestLogNamedDecimalUint)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedIntIterator struct {
	Event *GatewayMEVMInboundTestLogNamedInt // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedInt)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedInt)
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
func (it *GatewayMEVMInboundTestLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedInt represents a LogNamedInt event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedIntIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedIntIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedInt)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedInt(log types.Log) (*GatewayMEVMInboundTestLogNamedInt, error) {
	event := new(GatewayMEVMInboundTestLogNamedInt)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedStringIterator struct {
	Event *GatewayMEVMInboundTestLogNamedString // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedString)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedString)
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
func (it *GatewayMEVMInboundTestLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedString represents a LogNamedString event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedStringIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedStringIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedString) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedString)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedString(log types.Log) (*GatewayMEVMInboundTestLogNamedString, error) {
	event := new(GatewayMEVMInboundTestLogNamedString)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedUintIterator struct {
	Event *GatewayMEVMInboundTestLogNamedUint // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogNamedUint)
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
		it.Event = new(GatewayMEVMInboundTestLogNamedUint)
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
func (it *GatewayMEVMInboundTestLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogNamedUint represents a LogNamedUint event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogNamedUintIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogNamedUintIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogNamedUint)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogNamedUint(log types.Log) (*GatewayMEVMInboundTestLogNamedUint, error) {
	event := new(GatewayMEVMInboundTestLogNamedUint)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogStringIterator struct {
	Event *GatewayMEVMInboundTestLogString // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogString)
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
		it.Event = new(GatewayMEVMInboundTestLogString)
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
func (it *GatewayMEVMInboundTestLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogString represents a LogString event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogString(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogStringIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogStringIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogString) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogString)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogString(log types.Log) (*GatewayMEVMInboundTestLogString, error) {
	event := new(GatewayMEVMInboundTestLogString)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogUintIterator struct {
	Event *GatewayMEVMInboundTestLogUint // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogUint)
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
		it.Event = new(GatewayMEVMInboundTestLogUint)
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
func (it *GatewayMEVMInboundTestLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogUint represents a LogUint event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogUint(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogUintIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogUintIterator{contract: _GatewayMEVMInboundTest.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogUint) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogUint)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogUint(log types.Log) (*GatewayMEVMInboundTestLogUint, error) {
	event := new(GatewayMEVMInboundTestLogUint)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayMEVMInboundTestLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogsIterator struct {
	Event *GatewayMEVMInboundTestLogs // Event containing the contract specifics and raw log

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
func (it *GatewayMEVMInboundTestLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayMEVMInboundTestLogs)
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
		it.Event = new(GatewayMEVMInboundTestLogs)
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
func (it *GatewayMEVMInboundTestLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayMEVMInboundTestLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayMEVMInboundTestLogs represents a Logs event raised by the GatewayMEVMInboundTest contract.
type GatewayMEVMInboundTestLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) FilterLogs(opts *bind.FilterOpts) (*GatewayMEVMInboundTestLogsIterator, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &GatewayMEVMInboundTestLogsIterator{contract: _GatewayMEVMInboundTest.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *GatewayMEVMInboundTestLogs) (event.Subscription, error) {

	logs, sub, err := _GatewayMEVMInboundTest.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayMEVMInboundTestLogs)
				if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_GatewayMEVMInboundTest *GatewayMEVMInboundTestFilterer) ParseLogs(log types.Log) (*GatewayMEVMInboundTestLogs, error) {
	event := new(GatewayMEVMInboundTestLogs)
	if err := _GatewayMEVMInboundTest.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
