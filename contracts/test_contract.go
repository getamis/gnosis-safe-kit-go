// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestContractABI is the input ABI used to generate the binding from.
const TestContractABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"name\":\"callAssert\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"ok\",\"type\":\"bool\"}],\"name\":\"callRequire\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"callRevert\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// TestContractBin is the compiled bytecode used for deploying new contracts.
var TestContractBin = "0x608060405234801561001057600080fd5b5060e68061001f6000396000f3fe60806040526004361060305760003560e01c80634e79a584146035578063663720b9146062578063ed32929914608f575b600080fd5b606060048036036020811015604957600080fd5b810190808035151590602001909291905050506097565b005b608d60048036036020811015607657600080fd5b8101908080351515906020019092919050505060a3565b005b609560ac565b005b8060a057600080fd5b50565b8060a957fe5b50565b600080fdfea265627a7a7231582086fc181478e9a3e29bb2aed20a494161419698a3743af662533c121352bb626464736f6c63430005110032"

// DeployTestContract deploys a new Ethereum contract, binding an instance of TestContract to it.
func DeployTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// TestContract is an auto generated Go binding around an Ethereum contract.
type TestContract struct {
	TestContractCaller     // Read-only binding to the contract
	TestContractTransactor // Write-only binding to the contract
	TestContractFilterer   // Log filterer for contract events
}

// TestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractSession struct {
	Contract     *TestContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractCallerSession struct {
	Contract *TestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractTransactorSession struct {
	Contract     *TestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractRaw struct {
	Contract *TestContract // Generic contract binding to access the raw methods on
}

// TestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractCallerRaw struct {
	Contract *TestContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractTransactorRaw struct {
	Contract *TestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContract creates a new instance of TestContract, bound to a specific deployed contract.
func NewTestContract(address common.Address, backend bind.ContractBackend) (*TestContract, error) {
	contract, err := bindTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// NewTestContractCaller creates a new read-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractCaller(address common.Address, caller bind.ContractCaller) (*TestContractCaller, error) {
	contract, err := bindTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractCaller{contract: contract}, nil
}

// NewTestContractTransactor creates a new write-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractTransactor, error) {
	contract, err := bindTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractTransactor{contract: contract}, nil
}

// NewTestContractFilterer creates a new log filterer instance of TestContract, bound to a specific deployed contract.
func NewTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestContractFilterer, error) {
	contract, err := bindTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestContractFilterer{contract: contract}, nil
}

// bindTestContract binds a generic wrapper to an already deployed contract.
func bindTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.TestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transact(opts, method, params...)
}

// CallAssert is a paid mutator transaction binding the contract method 0x663720b9.
//
// Solidity: function callAssert(bool ok) payable returns()
func (_TestContract *TestContractTransactor) CallAssert(opts *bind.TransactOpts, ok bool) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "callAssert", ok)
}

// CallAssert is a paid mutator transaction binding the contract method 0x663720b9.
//
// Solidity: function callAssert(bool ok) payable returns()
func (_TestContract *TestContractSession) CallAssert(ok bool) (*types.Transaction, error) {
	return _TestContract.Contract.CallAssert(&_TestContract.TransactOpts, ok)
}

// CallAssert is a paid mutator transaction binding the contract method 0x663720b9.
//
// Solidity: function callAssert(bool ok) payable returns()
func (_TestContract *TestContractTransactorSession) CallAssert(ok bool) (*types.Transaction, error) {
	return _TestContract.Contract.CallAssert(&_TestContract.TransactOpts, ok)
}

// CallRequire is a paid mutator transaction binding the contract method 0x4e79a584.
//
// Solidity: function callRequire(bool ok) payable returns()
func (_TestContract *TestContractTransactor) CallRequire(opts *bind.TransactOpts, ok bool) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "callRequire", ok)
}

// CallRequire is a paid mutator transaction binding the contract method 0x4e79a584.
//
// Solidity: function callRequire(bool ok) payable returns()
func (_TestContract *TestContractSession) CallRequire(ok bool) (*types.Transaction, error) {
	return _TestContract.Contract.CallRequire(&_TestContract.TransactOpts, ok)
}

// CallRequire is a paid mutator transaction binding the contract method 0x4e79a584.
//
// Solidity: function callRequire(bool ok) payable returns()
func (_TestContract *TestContractTransactorSession) CallRequire(ok bool) (*types.Transaction, error) {
	return _TestContract.Contract.CallRequire(&_TestContract.TransactOpts, ok)
}

// CallRevert is a paid mutator transaction binding the contract method 0xed329299.
//
// Solidity: function callRevert() payable returns()
func (_TestContract *TestContractTransactor) CallRevert(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "callRevert")
}

// CallRevert is a paid mutator transaction binding the contract method 0xed329299.
//
// Solidity: function callRevert() payable returns()
func (_TestContract *TestContractSession) CallRevert() (*types.Transaction, error) {
	return _TestContract.Contract.CallRevert(&_TestContract.TransactOpts)
}

// CallRevert is a paid mutator transaction binding the contract method 0xed329299.
//
// Solidity: function callRevert() payable returns()
func (_TestContract *TestContractTransactorSession) CallRevert() (*types.Transaction, error) {
	return _TestContract.Contract.CallRevert(&_TestContract.TransactOpts)
}
