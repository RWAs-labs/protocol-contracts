// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package muse

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

// MuseNonEthMetaData contains all meta data concerning the MuseNonEth contract.
var MuseNonEthMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tssAddressUpdater_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"connectorAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceTssAddressUpdater\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tssAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tssAddressUpdater\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateTssAndConnectorAddresses\",\"inputs\":[{\"name\":\"tssAddress_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"connectorAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Burnt\",\"inputs\":[{\"name\":\"burnee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ConnectorAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newConnectorAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"mintee\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"internalSendHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TSSAddressUpdaterUpdated\",\"inputs\":[{\"name\":\"callerAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newTssUpdaterAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotConnector\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTss\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssOrUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"CallerIsNotTssUpdater\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MuseTransferError\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b506040516112a63803806112a683398101604081905261002f91610110565b604051806040016040528060048152602001634d75736560e01b815250604051806040016040528060048152602001634d55534560e01b815250816003908161007891906101e2565b50600461008582826101e2565b5050506001600160a01b03821615806100a557506001600160a01b038116155b156100c35760405163e6c4247b60e01b815260040160405180910390fd5b600680546001600160a01b039384166001600160a01b031991821617909155600780549290931691161790556102a0565b80516001600160a01b038116811461010b57600080fd5b919050565b6000806040838503121561012357600080fd5b61012c836100f4565b915061013a602084016100f4565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c9082168061016d57607f821691505b60208210810361018d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156101dd57806000526020600020601f840160051c810160208510156101ba5750805b601f840160051c820191505b818110156101da57600081556001016101c6565b50505b505050565b81516001600160401b038111156101fb576101fb610143565b61020f816102098454610159565b84610193565b6020601f821160018114610243576000831561022b5750848201515b600019600385901b1c1916600184901b1784556101da565b600084815260208120601f198516915b828110156102735787850151825560209485019460019092019101610253565b50848210156102915786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b610ff7806102af6000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806342966c68116100b257806379cc679011610081578063a9059cbb11610066578063a9059cbb1461028e578063bff9662a146102a1578063dd62ed3e146102c157600080fd5b806379cc67901461027357806395d89b411461028657600080fd5b806342966c68146102025780635b1125911461021557806370a0823114610235578063779e3b631461026b57600080fd5b80631e458bee116100ee5780631e458bee1461018857806323b872dd1461019b578063313ce567146101ae578063328a01d0146101bd57600080fd5b806306fdde0314610120578063095ea7b31461013e57806315d57fd41461016157806318160ddd14610176575b600080fd5b610128610307565b6040516101359190610d97565b60405180910390f35b61015161014c366004610e2c565b610399565b6040519015158152602001610135565b61017461016f366004610e56565b6103b3565b005b6002545b604051908152602001610135565b610174610196366004610e89565b61057e565b6101516101a9366004610ebc565b610631565b60405160128152602001610135565b6007546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610135565b610174610210366004610ef9565b610655565b6006546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a610243366004610f12565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b610174610662565b610174610281366004610e2c565b610786565b610128610837565b61015161029c366004610e2c565b610846565b6005546101dd9073ffffffffffffffffffffffffffffffffffffffff1681565b61017a6102cf366004610e56565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461031690610f34565b80601f016020809104026020016040519081016040528092919081815260200182805461034290610f34565b801561038f5780601f106103645761010080835404028352916020019161038f565b820191906000526020600020905b81548152906001019060200180831161037257829003601f168201915b5050505050905090565b6000336103a7818585610854565b60019150505b92915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633148015906103f3575060065473ffffffffffffffffffffffffffffffffffffffff163314155b15610431576040517fcdfcef970000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82161580610468575073ffffffffffffffffffffffffffffffffffffffff8116155b1561049f576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6006805473ffffffffffffffffffffffffffffffffffffffff8481167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316811790935560058054918516919092161790556040805133815260208101929092527fe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff910160405180910390a16040805133815273ffffffffffffffffffffffffffffffffffffffff831660208201527f1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c910160405180910390a15050565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105d1576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6105db8383610866565b808373ffffffffffffffffffffffffffffffffffffffff167fc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb8460405161062491815260200190565b60405180910390a3505050565b60003361063f8582856108c6565b61064a858585610995565b506001949350505050565b61065f3382610a40565b50565b60075473ffffffffffffffffffffffffffffffffffffffff1633146106b5576040517fe700765e000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b60065473ffffffffffffffffffffffffffffffffffffffff16610704576040517fe6c4247b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600654600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90921691821790556040805133815260208101929092527f5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0910160405180910390a1565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107d9576040517f3fe32fba000000000000000000000000000000000000000000000000000000008152336004820152602401610428565b6107e38282610a9c565b8173ffffffffffffffffffffffffffffffffffffffff167f919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b18260405161082b91815260200190565b60405180910390a25050565b60606004805461031690610f34565b6000336103a7818585610995565b6108618383836001610ab1565b505050565b73ffffffffffffffffffffffffffffffffffffffff82166108b6576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c260008383610bf9565b5050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461098f5781811015610980576040517ffb8f41b200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024810182905260448101839052606401610428565b61098f84848484036000610ab1565b50505050565b73ffffffffffffffffffffffffffffffffffffffff83166109e5576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8216610a35576040517fec442f0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b610861838383610bf9565b73ffffffffffffffffffffffffffffffffffffffff8216610a90576040517f96c6fd1e00000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b6108c282600083610bf9565b610aa78233836108c6565b6108c28282610a40565b73ffffffffffffffffffffffffffffffffffffffff8416610b01576040517fe602df0500000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8316610b51576040517f94280d6200000000000000000000000000000000000000000000000000000000815260006004820152602401610428565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152600160209081526040808320938716835292905220829055801561098f578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610beb91815260200190565b60405180910390a350505050565b73ffffffffffffffffffffffffffffffffffffffff8316610c31578060026000828254610c269190610f87565b90915550610ce39050565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610cb7576040517fe450d38c00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851660048201526024810182905260448101839052606401610428565b73ffffffffffffffffffffffffffffffffffffffff841660009081526020819052604090209082900390555b73ffffffffffffffffffffffffffffffffffffffff8216610d0c57600280548290039055610d38565b73ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604090208054820190555b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161062491815260200190565b602081526000825180602084015260005b81811015610dc55760208186018101516040868401015201610da8565b5060006040828501015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e2757600080fd5b919050565b60008060408385031215610e3f57600080fd5b610e4883610e03565b946020939093013593505050565b60008060408385031215610e6957600080fd5b610e7283610e03565b9150610e8060208401610e03565b90509250929050565b600080600060608486031215610e9e57600080fd5b610ea784610e03565b95602085013595506040909401359392505050565b600080600060608486031215610ed157600080fd5b610eda84610e03565b9250610ee860208501610e03565b929592945050506040919091013590565b600060208284031215610f0b57600080fd5b5035919050565b600060208284031215610f2457600080fd5b610f2d82610e03565b9392505050565b600181811c90821680610f4857607f821691505b602082108103610f81577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b808201808211156103ad577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220019aa29a7bddde82747dc2dbf333975b1079ccd42dd8a4b19f930198a4ca6e7e64736f6c634300081d0033",
}

// MuseNonEthABI is the input ABI used to generate the binding from.
// Deprecated: Use MuseNonEthMetaData.ABI instead.
var MuseNonEthABI = MuseNonEthMetaData.ABI

// MuseNonEthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MuseNonEthMetaData.Bin instead.
var MuseNonEthBin = MuseNonEthMetaData.Bin

// DeployMuseNonEth deploys a new Ethereum contract, binding an instance of MuseNonEth to it.
func DeployMuseNonEth(auth *bind.TransactOpts, backend bind.ContractBackend, tssAddress_ common.Address, tssAddressUpdater_ common.Address) (common.Address, *types.Transaction, *MuseNonEth, error) {
	parsed, err := MuseNonEthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MuseNonEthBin), backend, tssAddress_, tssAddressUpdater_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MuseNonEth{MuseNonEthCaller: MuseNonEthCaller{contract: contract}, MuseNonEthTransactor: MuseNonEthTransactor{contract: contract}, MuseNonEthFilterer: MuseNonEthFilterer{contract: contract}}, nil
}

// MuseNonEth is an auto generated Go binding around an Ethereum contract.
type MuseNonEth struct {
	MuseNonEthCaller     // Read-only binding to the contract
	MuseNonEthTransactor // Write-only binding to the contract
	MuseNonEthFilterer   // Log filterer for contract events
}

// MuseNonEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuseNonEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseNonEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuseNonEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseNonEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuseNonEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuseNonEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuseNonEthSession struct {
	Contract     *MuseNonEth       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuseNonEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuseNonEthCallerSession struct {
	Contract *MuseNonEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MuseNonEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuseNonEthTransactorSession struct {
	Contract     *MuseNonEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MuseNonEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuseNonEthRaw struct {
	Contract *MuseNonEth // Generic contract binding to access the raw methods on
}

// MuseNonEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuseNonEthCallerRaw struct {
	Contract *MuseNonEthCaller // Generic read-only contract binding to access the raw methods on
}

// MuseNonEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuseNonEthTransactorRaw struct {
	Contract *MuseNonEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuseNonEth creates a new instance of MuseNonEth, bound to a specific deployed contract.
func NewMuseNonEth(address common.Address, backend bind.ContractBackend) (*MuseNonEth, error) {
	contract, err := bindMuseNonEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuseNonEth{MuseNonEthCaller: MuseNonEthCaller{contract: contract}, MuseNonEthTransactor: MuseNonEthTransactor{contract: contract}, MuseNonEthFilterer: MuseNonEthFilterer{contract: contract}}, nil
}

// NewMuseNonEthCaller creates a new read-only instance of MuseNonEth, bound to a specific deployed contract.
func NewMuseNonEthCaller(address common.Address, caller bind.ContractCaller) (*MuseNonEthCaller, error) {
	contract, err := bindMuseNonEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthCaller{contract: contract}, nil
}

// NewMuseNonEthTransactor creates a new write-only instance of MuseNonEth, bound to a specific deployed contract.
func NewMuseNonEthTransactor(address common.Address, transactor bind.ContractTransactor) (*MuseNonEthTransactor, error) {
	contract, err := bindMuseNonEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthTransactor{contract: contract}, nil
}

// NewMuseNonEthFilterer creates a new log filterer instance of MuseNonEth, bound to a specific deployed contract.
func NewMuseNonEthFilterer(address common.Address, filterer bind.ContractFilterer) (*MuseNonEthFilterer, error) {
	contract, err := bindMuseNonEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthFilterer{contract: contract}, nil
}

// bindMuseNonEth binds a generic wrapper to an already deployed contract.
func bindMuseNonEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuseNonEthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseNonEth *MuseNonEthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseNonEth.Contract.MuseNonEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseNonEth *MuseNonEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseNonEth.Contract.MuseNonEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseNonEth *MuseNonEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseNonEth.Contract.MuseNonEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuseNonEth *MuseNonEthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuseNonEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuseNonEth *MuseNonEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseNonEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuseNonEth *MuseNonEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuseNonEth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MuseNonEth *MuseNonEthCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MuseNonEth *MuseNonEthSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MuseNonEth.Contract.Allowance(&_MuseNonEth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MuseNonEth *MuseNonEthCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MuseNonEth.Contract.Allowance(&_MuseNonEth.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MuseNonEth *MuseNonEthCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MuseNonEth *MuseNonEthSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MuseNonEth.Contract.BalanceOf(&_MuseNonEth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MuseNonEth *MuseNonEthCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MuseNonEth.Contract.BalanceOf(&_MuseNonEth.CallOpts, account)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_MuseNonEth *MuseNonEthCaller) ConnectorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "connectorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_MuseNonEth *MuseNonEthSession) ConnectorAddress() (common.Address, error) {
	return _MuseNonEth.Contract.ConnectorAddress(&_MuseNonEth.CallOpts)
}

// ConnectorAddress is a free data retrieval call binding the contract method 0xbff9662a.
//
// Solidity: function connectorAddress() view returns(address)
func (_MuseNonEth *MuseNonEthCallerSession) ConnectorAddress() (common.Address, error) {
	return _MuseNonEth.Contract.ConnectorAddress(&_MuseNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MuseNonEth *MuseNonEthCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MuseNonEth *MuseNonEthSession) Decimals() (uint8, error) {
	return _MuseNonEth.Contract.Decimals(&_MuseNonEth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MuseNonEth *MuseNonEthCallerSession) Decimals() (uint8, error) {
	return _MuseNonEth.Contract.Decimals(&_MuseNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MuseNonEth *MuseNonEthCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MuseNonEth *MuseNonEthSession) Name() (string, error) {
	return _MuseNonEth.Contract.Name(&_MuseNonEth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MuseNonEth *MuseNonEthCallerSession) Name() (string, error) {
	return _MuseNonEth.Contract.Name(&_MuseNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MuseNonEth *MuseNonEthCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MuseNonEth *MuseNonEthSession) Symbol() (string, error) {
	return _MuseNonEth.Contract.Symbol(&_MuseNonEth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MuseNonEth *MuseNonEthCallerSession) Symbol() (string, error) {
	return _MuseNonEth.Contract.Symbol(&_MuseNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MuseNonEth *MuseNonEthCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MuseNonEth *MuseNonEthSession) TotalSupply() (*big.Int, error) {
	return _MuseNonEth.Contract.TotalSupply(&_MuseNonEth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MuseNonEth *MuseNonEthCallerSession) TotalSupply() (*big.Int, error) {
	return _MuseNonEth.Contract.TotalSupply(&_MuseNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseNonEth *MuseNonEthCaller) TssAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "tssAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseNonEth *MuseNonEthSession) TssAddress() (common.Address, error) {
	return _MuseNonEth.Contract.TssAddress(&_MuseNonEth.CallOpts)
}

// TssAddress is a free data retrieval call binding the contract method 0x5b112591.
//
// Solidity: function tssAddress() view returns(address)
func (_MuseNonEth *MuseNonEthCallerSession) TssAddress() (common.Address, error) {
	return _MuseNonEth.Contract.TssAddress(&_MuseNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_MuseNonEth *MuseNonEthCaller) TssAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuseNonEth.contract.Call(opts, &out, "tssAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_MuseNonEth *MuseNonEthSession) TssAddressUpdater() (common.Address, error) {
	return _MuseNonEth.Contract.TssAddressUpdater(&_MuseNonEth.CallOpts)
}

// TssAddressUpdater is a free data retrieval call binding the contract method 0x328a01d0.
//
// Solidity: function tssAddressUpdater() view returns(address)
func (_MuseNonEth *MuseNonEthCallerSession) TssAddressUpdater() (common.Address, error) {
	return _MuseNonEth.Contract.TssAddressUpdater(&_MuseNonEth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Approve(&_MuseNonEth.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Approve(&_MuseNonEth.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MuseNonEth *MuseNonEthTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MuseNonEth *MuseNonEthSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Burn(&_MuseNonEth.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_MuseNonEth *MuseNonEthTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Burn(&_MuseNonEth.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MuseNonEth *MuseNonEthTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MuseNonEth *MuseNonEthSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.BurnFrom(&_MuseNonEth.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_MuseNonEth *MuseNonEthTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.BurnFrom(&_MuseNonEth.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_MuseNonEth *MuseNonEthTransactor) Mint(opts *bind.TransactOpts, mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "mint", mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_MuseNonEth *MuseNonEthSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Mint(&_MuseNonEth.TransactOpts, mintee, value, internalSendHash)
}

// Mint is a paid mutator transaction binding the contract method 0x1e458bee.
//
// Solidity: function mint(address mintee, uint256 value, bytes32 internalSendHash) returns()
func (_MuseNonEth *MuseNonEthTransactorSession) Mint(mintee common.Address, value *big.Int, internalSendHash [32]byte) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Mint(&_MuseNonEth.TransactOpts, mintee, value, internalSendHash)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_MuseNonEth *MuseNonEthTransactor) RenounceTssAddressUpdater(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "renounceTssAddressUpdater")
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_MuseNonEth *MuseNonEthSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _MuseNonEth.Contract.RenounceTssAddressUpdater(&_MuseNonEth.TransactOpts)
}

// RenounceTssAddressUpdater is a paid mutator transaction binding the contract method 0x779e3b63.
//
// Solidity: function renounceTssAddressUpdater() returns()
func (_MuseNonEth *MuseNonEthTransactorSession) RenounceTssAddressUpdater() (*types.Transaction, error) {
	return _MuseNonEth.Contract.RenounceTssAddressUpdater(&_MuseNonEth.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Transfer(&_MuseNonEth.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.Transfer(&_MuseNonEth.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.TransferFrom(&_MuseNonEth.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_MuseNonEth *MuseNonEthTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MuseNonEth.Contract.TransferFrom(&_MuseNonEth.TransactOpts, from, to, value)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_MuseNonEth *MuseNonEthTransactor) UpdateTssAndConnectorAddresses(opts *bind.TransactOpts, tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _MuseNonEth.contract.Transact(opts, "updateTssAndConnectorAddresses", tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_MuseNonEth *MuseNonEthSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _MuseNonEth.Contract.UpdateTssAndConnectorAddresses(&_MuseNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// UpdateTssAndConnectorAddresses is a paid mutator transaction binding the contract method 0x15d57fd4.
//
// Solidity: function updateTssAndConnectorAddresses(address tssAddress_, address connectorAddress_) returns()
func (_MuseNonEth *MuseNonEthTransactorSession) UpdateTssAndConnectorAddresses(tssAddress_ common.Address, connectorAddress_ common.Address) (*types.Transaction, error) {
	return _MuseNonEth.Contract.UpdateTssAndConnectorAddresses(&_MuseNonEth.TransactOpts, tssAddress_, connectorAddress_)
}

// MuseNonEthApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MuseNonEth contract.
type MuseNonEthApprovalIterator struct {
	Event *MuseNonEthApproval // Event containing the contract specifics and raw log

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
func (it *MuseNonEthApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthApproval)
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
		it.Event = new(MuseNonEthApproval)
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
func (it *MuseNonEthApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthApproval represents a Approval event raised by the MuseNonEth contract.
type MuseNonEthApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MuseNonEth *MuseNonEthFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MuseNonEthApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthApprovalIterator{contract: _MuseNonEth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MuseNonEth *MuseNonEthFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MuseNonEthApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthApproval)
				if err := _MuseNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MuseNonEth *MuseNonEthFilterer) ParseApproval(log types.Log) (*MuseNonEthApproval, error) {
	event := new(MuseNonEthApproval)
	if err := _MuseNonEth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthBurntIterator is returned from FilterBurnt and is used to iterate over the raw logs and unpacked data for Burnt events raised by the MuseNonEth contract.
type MuseNonEthBurntIterator struct {
	Event *MuseNonEthBurnt // Event containing the contract specifics and raw log

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
func (it *MuseNonEthBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthBurnt)
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
		it.Event = new(MuseNonEthBurnt)
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
func (it *MuseNonEthBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthBurnt represents a Burnt event raised by the MuseNonEth contract.
type MuseNonEthBurnt struct {
	Burnee common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurnt is a free log retrieval operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_MuseNonEth *MuseNonEthFilterer) FilterBurnt(opts *bind.FilterOpts, burnee []common.Address) (*MuseNonEthBurntIterator, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthBurntIterator{contract: _MuseNonEth.contract, event: "Burnt", logs: logs, sub: sub}, nil
}

// WatchBurnt is a free log subscription operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_MuseNonEth *MuseNonEthFilterer) WatchBurnt(opts *bind.WatchOpts, sink chan<- *MuseNonEthBurnt, burnee []common.Address) (event.Subscription, error) {

	var burneeRule []interface{}
	for _, burneeItem := range burnee {
		burneeRule = append(burneeRule, burneeItem)
	}

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "Burnt", burneeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthBurnt)
				if err := _MuseNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
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

// ParseBurnt is a log parse operation binding the contract event 0x919f7e2092ffcc9d09f599be18d8152860b0c054df788a33bc549cdd9d0f15b1.
//
// Solidity: event Burnt(address indexed burnee, uint256 amount)
func (_MuseNonEth *MuseNonEthFilterer) ParseBurnt(log types.Log) (*MuseNonEthBurnt, error) {
	event := new(MuseNonEthBurnt)
	if err := _MuseNonEth.contract.UnpackLog(event, "Burnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthConnectorAddressUpdatedIterator is returned from FilterConnectorAddressUpdated and is used to iterate over the raw logs and unpacked data for ConnectorAddressUpdated events raised by the MuseNonEth contract.
type MuseNonEthConnectorAddressUpdatedIterator struct {
	Event *MuseNonEthConnectorAddressUpdated // Event containing the contract specifics and raw log

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
func (it *MuseNonEthConnectorAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthConnectorAddressUpdated)
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
		it.Event = new(MuseNonEthConnectorAddressUpdated)
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
func (it *MuseNonEthConnectorAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthConnectorAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthConnectorAddressUpdated represents a ConnectorAddressUpdated event raised by the MuseNonEth contract.
type MuseNonEthConnectorAddressUpdated struct {
	CallerAddress       common.Address
	NewConnectorAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConnectorAddressUpdated is a free log retrieval operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_MuseNonEth *MuseNonEthFilterer) FilterConnectorAddressUpdated(opts *bind.FilterOpts) (*MuseNonEthConnectorAddressUpdatedIterator, error) {

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "ConnectorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &MuseNonEthConnectorAddressUpdatedIterator{contract: _MuseNonEth.contract, event: "ConnectorAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchConnectorAddressUpdated is a free log subscription operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_MuseNonEth *MuseNonEthFilterer) WatchConnectorAddressUpdated(opts *bind.WatchOpts, sink chan<- *MuseNonEthConnectorAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "ConnectorAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthConnectorAddressUpdated)
				if err := _MuseNonEth.contract.UnpackLog(event, "ConnectorAddressUpdated", log); err != nil {
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

// ParseConnectorAddressUpdated is a log parse operation binding the contract event 0x1b9352454524a57a51f24f67dc66d898f616922cd1f7a12d73570ece12b1975c.
//
// Solidity: event ConnectorAddressUpdated(address callerAddress, address newConnectorAddress)
func (_MuseNonEth *MuseNonEthFilterer) ParseConnectorAddressUpdated(log types.Log) (*MuseNonEthConnectorAddressUpdated, error) {
	event := new(MuseNonEthConnectorAddressUpdated)
	if err := _MuseNonEth.contract.UnpackLog(event, "ConnectorAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the MuseNonEth contract.
type MuseNonEthMintedIterator struct {
	Event *MuseNonEthMinted // Event containing the contract specifics and raw log

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
func (it *MuseNonEthMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthMinted)
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
		it.Event = new(MuseNonEthMinted)
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
func (it *MuseNonEthMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthMinted represents a Minted event raised by the MuseNonEth contract.
type MuseNonEthMinted struct {
	Mintee           common.Address
	Amount           *big.Int
	InternalSendHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_MuseNonEth *MuseNonEthFilterer) FilterMinted(opts *bind.FilterOpts, mintee []common.Address, internalSendHash [][32]byte) (*MuseNonEthMintedIterator, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthMintedIterator{contract: _MuseNonEth.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_MuseNonEth *MuseNonEthFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MuseNonEthMinted, mintee []common.Address, internalSendHash [][32]byte) (event.Subscription, error) {

	var minteeRule []interface{}
	for _, minteeItem := range mintee {
		minteeRule = append(minteeRule, minteeItem)
	}

	var internalSendHashRule []interface{}
	for _, internalSendHashItem := range internalSendHash {
		internalSendHashRule = append(internalSendHashRule, internalSendHashItem)
	}

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "Minted", minteeRule, internalSendHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthMinted)
				if err := _MuseNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xc263b302aec62d29105026245f19e16f8e0137066ccd4a8bd941f716bd4096bb.
//
// Solidity: event Minted(address indexed mintee, uint256 amount, bytes32 indexed internalSendHash)
func (_MuseNonEth *MuseNonEthFilterer) ParseMinted(log types.Log) (*MuseNonEthMinted, error) {
	event := new(MuseNonEthMinted)
	if err := _MuseNonEth.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthTSSAddressUpdatedIterator is returned from FilterTSSAddressUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdated events raised by the MuseNonEth contract.
type MuseNonEthTSSAddressUpdatedIterator struct {
	Event *MuseNonEthTSSAddressUpdated // Event containing the contract specifics and raw log

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
func (it *MuseNonEthTSSAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthTSSAddressUpdated)
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
		it.Event = new(MuseNonEthTSSAddressUpdated)
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
func (it *MuseNonEthTSSAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthTSSAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthTSSAddressUpdated represents a TSSAddressUpdated event raised by the MuseNonEth contract.
type MuseNonEthTSSAddressUpdated struct {
	CallerAddress common.Address
	NewTssAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdated is a free log retrieval operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_MuseNonEth *MuseNonEthFilterer) FilterTSSAddressUpdated(opts *bind.FilterOpts) (*MuseNonEthTSSAddressUpdatedIterator, error) {

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &MuseNonEthTSSAddressUpdatedIterator{contract: _MuseNonEth.contract, event: "TSSAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdated is a free log subscription operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_MuseNonEth *MuseNonEthFilterer) WatchTSSAddressUpdated(opts *bind.WatchOpts, sink chan<- *MuseNonEthTSSAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "TSSAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthTSSAddressUpdated)
				if err := _MuseNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
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

// ParseTSSAddressUpdated is a log parse operation binding the contract event 0xe79965b5c67dcfb2cf5fe152715e4a7256cee62a3d5dd8484fd8a8539eb8beff.
//
// Solidity: event TSSAddressUpdated(address callerAddress, address newTssAddress)
func (_MuseNonEth *MuseNonEthFilterer) ParseTSSAddressUpdated(log types.Log) (*MuseNonEthTSSAddressUpdated, error) {
	event := new(MuseNonEthTSSAddressUpdated)
	if err := _MuseNonEth.contract.UnpackLog(event, "TSSAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthTSSAddressUpdaterUpdatedIterator is returned from FilterTSSAddressUpdaterUpdated and is used to iterate over the raw logs and unpacked data for TSSAddressUpdaterUpdated events raised by the MuseNonEth contract.
type MuseNonEthTSSAddressUpdaterUpdatedIterator struct {
	Event *MuseNonEthTSSAddressUpdaterUpdated // Event containing the contract specifics and raw log

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
func (it *MuseNonEthTSSAddressUpdaterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthTSSAddressUpdaterUpdated)
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
		it.Event = new(MuseNonEthTSSAddressUpdaterUpdated)
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
func (it *MuseNonEthTSSAddressUpdaterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthTSSAddressUpdaterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthTSSAddressUpdaterUpdated represents a TSSAddressUpdaterUpdated event raised by the MuseNonEth contract.
type MuseNonEthTSSAddressUpdaterUpdated struct {
	CallerAddress        common.Address
	NewTssUpdaterAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTSSAddressUpdaterUpdated is a free log retrieval operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_MuseNonEth *MuseNonEthFilterer) FilterTSSAddressUpdaterUpdated(opts *bind.FilterOpts) (*MuseNonEthTSSAddressUpdaterUpdatedIterator, error) {

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return &MuseNonEthTSSAddressUpdaterUpdatedIterator{contract: _MuseNonEth.contract, event: "TSSAddressUpdaterUpdated", logs: logs, sub: sub}, nil
}

// WatchTSSAddressUpdaterUpdated is a free log subscription operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_MuseNonEth *MuseNonEthFilterer) WatchTSSAddressUpdaterUpdated(opts *bind.WatchOpts, sink chan<- *MuseNonEthTSSAddressUpdaterUpdated) (event.Subscription, error) {

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "TSSAddressUpdaterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthTSSAddressUpdaterUpdated)
				if err := _MuseNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
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

// ParseTSSAddressUpdaterUpdated is a log parse operation binding the contract event 0x5104c9abdc7d111c2aeb4ce890ac70274a4be2ee83f46a62551be5e6ebc82dd0.
//
// Solidity: event TSSAddressUpdaterUpdated(address callerAddress, address newTssUpdaterAddress)
func (_MuseNonEth *MuseNonEthFilterer) ParseTSSAddressUpdaterUpdated(log types.Log) (*MuseNonEthTSSAddressUpdaterUpdated, error) {
	event := new(MuseNonEthTSSAddressUpdaterUpdated)
	if err := _MuseNonEth.contract.UnpackLog(event, "TSSAddressUpdaterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MuseNonEthTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MuseNonEth contract.
type MuseNonEthTransferIterator struct {
	Event *MuseNonEthTransfer // Event containing the contract specifics and raw log

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
func (it *MuseNonEthTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MuseNonEthTransfer)
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
		it.Event = new(MuseNonEthTransfer)
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
func (it *MuseNonEthTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MuseNonEthTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MuseNonEthTransfer represents a Transfer event raised by the MuseNonEth contract.
type MuseNonEthTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MuseNonEth *MuseNonEthFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MuseNonEthTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseNonEth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MuseNonEthTransferIterator{contract: _MuseNonEth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MuseNonEth *MuseNonEthFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MuseNonEthTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MuseNonEth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MuseNonEthTransfer)
				if err := _MuseNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MuseNonEth *MuseNonEthFilterer) ParseTransfer(log types.Log) (*MuseNonEthTransfer, error) {
	event := new(MuseNonEthTransfer)
	if err := _MuseNonEth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
