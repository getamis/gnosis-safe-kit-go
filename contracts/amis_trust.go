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

// AMISTrustABI is the input ABI used to generate the binding from.
const AMISTrustABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"contractAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumEnum.Operation\",\"name\":\"operation\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"safeTxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"gasToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"refundReceiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"execTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// AMISTrustBin is the compiled bytecode used for deploying new contracts.
var AMISTrustBin = "0x608060405234801561001057600080fd5b506105c0806100206000396000f3fe60806040526004361061001e5760003560e01c806368bb7c6914610023575b600080fd5b6101a5600480360361016081101561003a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803590602001906401000000008111156100a157600080fd5b8201836020820111156100b357600080fd5b803590602001918460018302840111640100000000831117156100d557600080fd5b9091929391929390803560ff169060200190929190803590602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561016157600080fd5b82018360208201111561017357600080fd5b8035906020019184600183028401116401000000008311171561019557600080fd5b90919293919293905050506101bf565b604051808215151515815260200191505060405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff168e73ffffffffffffffffffffffffffffffffffffffff161415610263576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f496e76616c696420636f6e74726163742061646472657373000000000000000081525060200191505060405180910390fd5b60008e90508e73ffffffffffffffffffffffffffffffffffffffff163460405180600001905060006040518083038185875af1925050503d80600081146102c6576040519150601f19603f3d011682016040523d82523d6000602084013e6102cb565b606091505b50508092505081610344576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f4661696c656420746f2073656e642076616c756500000000000000000000000081525060200191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16636a7612028f8f8f8f8f8f8f8f8f8f8f8f6040518d63ffffffff1660e01b8152600401808d73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018c8152602001806020018a60018111156103c757fe5b60ff1681526020018981526020018881526020018781526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200183810383528d8d82818152602001925080828437600081840152601f19601f8201169050808301925050508381038252858582818152602001925080828437600081840152601f19601f8201169050808301925050509e505050505050505050505050505050602060405180830381600087803b1580156104c957600080fd5b505af11580156104dd573d6000803e3d6000fd5b505050506040513d60208110156104f357600080fd5b8101908080519060200190929190505050915081610579576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f4661696c656420746f2065786563207472616e61736374696f6e00000000000081525060200191505060405180910390fd5b509d9c5050505050505050505050505056fea265627a7a723158200f2b8d1f10c78c93c2a62f6d6c9754b02db29a00e46bdafb4ede50bd22ffb2c964736f6c63430005110032"

// DeployAMISTrust deploys a new Ethereum contract, binding an instance of AMISTrust to it.
func DeployAMISTrust(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AMISTrust, error) {
	parsed, err := abi.JSON(strings.NewReader(AMISTrustABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AMISTrustBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AMISTrust{AMISTrustCaller: AMISTrustCaller{contract: contract}, AMISTrustTransactor: AMISTrustTransactor{contract: contract}, AMISTrustFilterer: AMISTrustFilterer{contract: contract}}, nil
}

// AMISTrust is an auto generated Go binding around an Ethereum contract.
type AMISTrust struct {
	AMISTrustCaller     // Read-only binding to the contract
	AMISTrustTransactor // Write-only binding to the contract
	AMISTrustFilterer   // Log filterer for contract events
}

// AMISTrustCaller is an auto generated read-only Go binding around an Ethereum contract.
type AMISTrustCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMISTrustTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AMISTrustTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMISTrustFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AMISTrustFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AMISTrustSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AMISTrustSession struct {
	Contract     *AMISTrust        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AMISTrustCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AMISTrustCallerSession struct {
	Contract *AMISTrustCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AMISTrustTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AMISTrustTransactorSession struct {
	Contract     *AMISTrustTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AMISTrustRaw is an auto generated low-level Go binding around an Ethereum contract.
type AMISTrustRaw struct {
	Contract *AMISTrust // Generic contract binding to access the raw methods on
}

// AMISTrustCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AMISTrustCallerRaw struct {
	Contract *AMISTrustCaller // Generic read-only contract binding to access the raw methods on
}

// AMISTrustTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AMISTrustTransactorRaw struct {
	Contract *AMISTrustTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAMISTrust creates a new instance of AMISTrust, bound to a specific deployed contract.
func NewAMISTrust(address common.Address, backend bind.ContractBackend) (*AMISTrust, error) {
	contract, err := bindAMISTrust(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AMISTrust{AMISTrustCaller: AMISTrustCaller{contract: contract}, AMISTrustTransactor: AMISTrustTransactor{contract: contract}, AMISTrustFilterer: AMISTrustFilterer{contract: contract}}, nil
}

// NewAMISTrustCaller creates a new read-only instance of AMISTrust, bound to a specific deployed contract.
func NewAMISTrustCaller(address common.Address, caller bind.ContractCaller) (*AMISTrustCaller, error) {
	contract, err := bindAMISTrust(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AMISTrustCaller{contract: contract}, nil
}

// NewAMISTrustTransactor creates a new write-only instance of AMISTrust, bound to a specific deployed contract.
func NewAMISTrustTransactor(address common.Address, transactor bind.ContractTransactor) (*AMISTrustTransactor, error) {
	contract, err := bindAMISTrust(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AMISTrustTransactor{contract: contract}, nil
}

// NewAMISTrustFilterer creates a new log filterer instance of AMISTrust, bound to a specific deployed contract.
func NewAMISTrustFilterer(address common.Address, filterer bind.ContractFilterer) (*AMISTrustFilterer, error) {
	contract, err := bindAMISTrust(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AMISTrustFilterer{contract: contract}, nil
}

// bindAMISTrust binds a generic wrapper to an already deployed contract.
func bindAMISTrust(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AMISTrustABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMISTrust *AMISTrustRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMISTrust.Contract.AMISTrustCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMISTrust *AMISTrustRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMISTrust.Contract.AMISTrustTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMISTrust *AMISTrustRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMISTrust.Contract.AMISTrustTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AMISTrust *AMISTrustCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AMISTrust.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AMISTrust *AMISTrustTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AMISTrust.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AMISTrust *AMISTrustTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AMISTrust.Contract.contract.Transact(opts, method, params...)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x68bb7c69.
//
// Solidity: function execTransaction(address contractAddr, address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_AMISTrust *AMISTrustTransactor) ExecTransaction(opts *bind.TransactOpts, contractAddr common.Address, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _AMISTrust.contract.Transact(opts, "execTransaction", contractAddr, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x68bb7c69.
//
// Solidity: function execTransaction(address contractAddr, address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_AMISTrust *AMISTrustSession) ExecTransaction(contractAddr common.Address, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _AMISTrust.Contract.ExecTransaction(&_AMISTrust.TransactOpts, contractAddr, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}

// ExecTransaction is a paid mutator transaction binding the contract method 0x68bb7c69.
//
// Solidity: function execTransaction(address contractAddr, address to, uint256 value, bytes data, uint8 operation, uint256 safeTxGas, uint256 baseGas, uint256 gasPrice, address gasToken, address refundReceiver, bytes signatures) payable returns(bool success)
func (_AMISTrust *AMISTrustTransactorSession) ExecTransaction(contractAddr common.Address, to common.Address, value *big.Int, data []byte, operation uint8, safeTxGas *big.Int, baseGas *big.Int, gasPrice *big.Int, gasToken common.Address, refundReceiver common.Address, signatures []byte) (*types.Transaction, error) {
	return _AMISTrust.Contract.ExecTransaction(&_AMISTrust.TransactOpts, contractAddr, to, value, data, operation, safeTxGas, baseGas, gasPrice, gasToken, refundReceiver, signatures)
}
