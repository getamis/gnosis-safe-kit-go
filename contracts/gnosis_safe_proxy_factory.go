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

// GnosisSafeProxyFactoryABI is the input ABI used to generate the binding from.
const GnosisSafeProxyFactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreation\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mastercopy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"calculateCreateProxyWithNonceAddress\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"masterCopy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createProxy\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mastercopy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"},{\"internalType\":\"contractIProxyCreationCallback\",\"name\":\"callback\",\"type\":\"address\"}],\"name\":\"createProxyWithCallback\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mastercopy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"createProxyWithNonce\",\"outputs\":[{\"internalType\":\"contractGnosisSafeProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyCreationCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxyRuntimeCode\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// GnosisSafeProxyFactoryBin is the compiled bytecode used for deploying new contracts.
var GnosisSafeProxyFactoryBin = "0x608060405234801561001057600080fd5b50610f73806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631688f0b9146100675780632500510e1461018c57806353e5d9351461026f57806361b69abd146102f2578063addacc0f1461040d578063d18af54d14610490575b600080fd5b61014a6004803603606081101561007d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156100ba57600080fd5b8201836020820111156100cc57600080fd5b803590602001918460018302840111640100000000831117156100ee57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506105d5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61022d600480360360608110156101a257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156101df57600080fd5b8201836020820111156101f157600080fd5b8035906020019184600183028401116401000000008311171561021357600080fd5b909192939192939080359060200190929190505050610674565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6102776107b7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102b757808201518184015260208101905061029c565b50505050905090810190601f1680156102e45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103cb6004803603604081101561030857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561034557600080fd5b82018360208201111561035757600080fd5b8035906020019184600183028401116401000000008311171561037957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506107e2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104156108d5565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561045557808201518184015260208101905061043a565b50505050905090810190601f1680156104825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610593600480360360808110156104a657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156104e357600080fd5b8201836020820111156104f557600080fd5b8035906020019184600183028401116401000000008311171561051757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610900565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b60006105e2848484610af1565b905060008351111561060a5760008060008551602087016000865af1141561060957600080fd5b5b7fa38789425dbeee0239e16ff2d2567e31720127fbc6430758c1a4efc6aef29f8081604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a19392505050565b60006106c58585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505084610af1565b905080604051602001808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b81526014019150506040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561077c578082015181840152602081019050610761565b50505050905090810190601f1680156107a95780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b6060604051806020016107c990610c94565b6020820181038252601f19601f82011660405250905090565b6000826040516107f190610c94565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610843573d6000803e3d6000fd5b50905060008251111561086c5760008060008451602086016000865af1141561086b57600080fd5b5b7fa38789425dbeee0239e16ff2d2567e31720127fbc6430758c1a4efc6aef29f8081604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a192915050565b6060604051806020016108e790610ca1565b6020820181038252601f19601f82011660405250905090565b6000808383604051602001808381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b8152601401925050506040516020818303038152906040528051906020012060001c90506109718686836105d5565b9150600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610ae8578273ffffffffffffffffffffffffffffffffffffffff16631e52b518838888886040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015610a80578082015181840152602081019050610a65565b50505050905090810190601f168015610aad5780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b158015610acf57600080fd5b505af1158015610ae3573d6000803e3d6000fd5b505050505b50949350505050565b6000808380519060200120836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209050606060405180602001610b3e90610c94565b6020820181038252601f19601f820116604052508673ffffffffffffffffffffffffffffffffffffffff166040516020018083805190602001908083835b60208310610b9f5780518252602082019150602081019050602083039250610b7c565b6001836020036101000a038019825116818451168082178552505050505050905001828152602001925050506040516020818303038152906040529050818151826020016000f59250600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c8b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f437265617465322063616c6c206661696c65640000000000000000000000000081525060200191505060405180910390fd5b50509392505050565b6101e780610cae83390190565b60aa80610e958339019056fe608060405234801561001057600080fd5b506040516101e73803806101e78339818101604052602081101561003357600080fd5b8101908080519060200190929190505050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156100ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806101c36024913960400191505060405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505060aa806101196000396000f3fe608060405273ffffffffffffffffffffffffffffffffffffffff600054167fa619486e0000000000000000000000000000000000000000000000000000000060003514156050578060005260206000f35b3660008037600080366000845af43d6000803e60008114156070573d6000fd5b3d6000f3fea265627a7a72315820c33e359396c174274739d4a77974186536e9fe5c386e2bafae3b244631fd719c64736f6c63430005110032496e76616c6964206d617374657220636f707920616464726573732070726f7669646564608060405273ffffffffffffffffffffffffffffffffffffffff600054167fa619486e0000000000000000000000000000000000000000000000000000000060003514156050578060005260206000f35b3660008037600080366000845af43d6000803e60008114156070573d6000fd5b3d6000f3fea265627a7a72315820c33e359396c174274739d4a77974186536e9fe5c386e2bafae3b244631fd719c64736f6c63430005110032a265627a7a7231582016d1a04585f22be555bf01750000d6068e90441eba1c02155b984673ffa1c4c964736f6c63430005110032"

// DeployGnosisSafeProxyFactory deploys a new Ethereum contract, binding an instance of GnosisSafeProxyFactory to it.
func DeployGnosisSafeProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GnosisSafeProxyFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeProxyFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GnosisSafeProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GnosisSafeProxyFactory{GnosisSafeProxyFactoryCaller: GnosisSafeProxyFactoryCaller{contract: contract}, GnosisSafeProxyFactoryTransactor: GnosisSafeProxyFactoryTransactor{contract: contract}, GnosisSafeProxyFactoryFilterer: GnosisSafeProxyFactoryFilterer{contract: contract}}, nil
}

// GnosisSafeProxyFactory is an auto generated Go binding around an Ethereum contract.
type GnosisSafeProxyFactory struct {
	GnosisSafeProxyFactoryCaller     // Read-only binding to the contract
	GnosisSafeProxyFactoryTransactor // Write-only binding to the contract
	GnosisSafeProxyFactoryFilterer   // Log filterer for contract events
}

// GnosisSafeProxyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GnosisSafeProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GnosisSafeProxyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GnosisSafeProxyFactorySession struct {
	Contract     *GnosisSafeProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GnosisSafeProxyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GnosisSafeProxyFactoryCallerSession struct {
	Contract *GnosisSafeProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// GnosisSafeProxyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GnosisSafeProxyFactoryTransactorSession struct {
	Contract     *GnosisSafeProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// GnosisSafeProxyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryRaw struct {
	Contract *GnosisSafeProxyFactory // Generic contract binding to access the raw methods on
}

// GnosisSafeProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryCallerRaw struct {
	Contract *GnosisSafeProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GnosisSafeProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GnosisSafeProxyFactoryTransactorRaw struct {
	Contract *GnosisSafeProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGnosisSafeProxyFactory creates a new instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactory(address common.Address, backend bind.ContractBackend) (*GnosisSafeProxyFactory, error) {
	contract, err := bindGnosisSafeProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactory{GnosisSafeProxyFactoryCaller: GnosisSafeProxyFactoryCaller{contract: contract}, GnosisSafeProxyFactoryTransactor: GnosisSafeProxyFactoryTransactor{contract: contract}, GnosisSafeProxyFactoryFilterer: GnosisSafeProxyFactoryFilterer{contract: contract}}, nil
}

// NewGnosisSafeProxyFactoryCaller creates a new read-only instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*GnosisSafeProxyFactoryCaller, error) {
	contract, err := bindGnosisSafeProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryCaller{contract: contract}, nil
}

// NewGnosisSafeProxyFactoryTransactor creates a new write-only instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GnosisSafeProxyFactoryTransactor, error) {
	contract, err := bindGnosisSafeProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryTransactor{contract: contract}, nil
}

// NewGnosisSafeProxyFactoryFilterer creates a new log filterer instance of GnosisSafeProxyFactory, bound to a specific deployed contract.
func NewGnosisSafeProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GnosisSafeProxyFactoryFilterer, error) {
	contract, err := bindGnosisSafeProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryFilterer{contract: contract}, nil
}

// bindGnosisSafeProxyFactory binds a generic wrapper to an already deployed contract.
func bindGnosisSafeProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GnosisSafeProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.GnosisSafeProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GnosisSafeProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCaller) ProxyCreationCode(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GnosisSafeProxyFactory.contract.Call(opts, out, "proxyCreationCode")
	return *ret0, err
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) ProxyCreationCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyCreationCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyCreationCode is a free data retrieval call binding the contract method 0x53e5d935.
//
// Solidity: function proxyCreationCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerSession) ProxyCreationCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyCreationCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCaller) ProxyRuntimeCode(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _GnosisSafeProxyFactory.contract.Call(opts, out, "proxyRuntimeCode")
	return *ret0, err
}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) ProxyRuntimeCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyRuntimeCode(&_GnosisSafeProxyFactory.CallOpts)
}

// ProxyRuntimeCode is a free data retrieval call binding the contract method 0xaddacc0f.
//
// Solidity: function proxyRuntimeCode() pure returns(bytes)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryCallerSession) ProxyRuntimeCode() ([]byte, error) {
	return _GnosisSafeProxyFactory.Contract.ProxyRuntimeCode(&_GnosisSafeProxyFactory.CallOpts)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CalculateCreateProxyWithNonceAddress(opts *bind.TransactOpts, _mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "calculateCreateProxyWithNonceAddress", _mastercopy, initializer, saltNonce)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CalculateCreateProxyWithNonceAddress(_mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CalculateCreateProxyWithNonceAddress(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce)
}

// CalculateCreateProxyWithNonceAddress is a paid mutator transaction binding the contract method 0x2500510e.
//
// Solidity: function calculateCreateProxyWithNonceAddress(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CalculateCreateProxyWithNonceAddress(_mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CalculateCreateProxyWithNonceAddress(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address masterCopy, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, masterCopy common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxy", masterCopy, data)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address masterCopy, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxy(masterCopy common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxy(&_GnosisSafeProxyFactory.TransactOpts, masterCopy, data)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x61b69abd.
//
// Solidity: function createProxy(address masterCopy, bytes data) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxy(masterCopy common.Address, data []byte) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxy(&_GnosisSafeProxyFactory.TransactOpts, masterCopy, data)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _mastercopy, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxyWithCallback(opts *bind.TransactOpts, _mastercopy common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxyWithCallback", _mastercopy, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _mastercopy, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxyWithCallback(_mastercopy common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithCallback(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce, callback)
}

// CreateProxyWithCallback is a paid mutator transaction binding the contract method 0xd18af54d.
//
// Solidity: function createProxyWithCallback(address _mastercopy, bytes initializer, uint256 saltNonce, address callback) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxyWithCallback(_mastercopy common.Address, initializer []byte, saltNonce *big.Int, callback common.Address) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithCallback(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce, callback)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactor) CreateProxyWithNonce(opts *bind.TransactOpts, _mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.contract.Transact(opts, "createProxyWithNonce", _mastercopy, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactorySession) CreateProxyWithNonce(_mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithNonce(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce)
}

// CreateProxyWithNonce is a paid mutator transaction binding the contract method 0x1688f0b9.
//
// Solidity: function createProxyWithNonce(address _mastercopy, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryTransactorSession) CreateProxyWithNonce(_mastercopy common.Address, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _GnosisSafeProxyFactory.Contract.CreateProxyWithNonce(&_GnosisSafeProxyFactory.TransactOpts, _mastercopy, initializer, saltNonce)
}

// GnosisSafeProxyFactoryProxyCreationIterator is returned from FilterProxyCreation and is used to iterate over the raw logs and unpacked data for ProxyCreation events raised by the GnosisSafeProxyFactory contract.
type GnosisSafeProxyFactoryProxyCreationIterator struct {
	Event *GnosisSafeProxyFactoryProxyCreation // Event containing the contract specifics and raw log

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
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GnosisSafeProxyFactoryProxyCreation)
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
		it.Event = new(GnosisSafeProxyFactoryProxyCreation)
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
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GnosisSafeProxyFactoryProxyCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GnosisSafeProxyFactoryProxyCreation represents a ProxyCreation event raised by the GnosisSafeProxyFactory contract.
type GnosisSafeProxyFactoryProxyCreation struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreation is a free log retrieval operation binding the contract event 0xa38789425dbeee0239e16ff2d2567e31720127fbc6430758c1a4efc6aef29f80.
//
// Solidity: event ProxyCreation(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) FilterProxyCreation(opts *bind.FilterOpts) (*GnosisSafeProxyFactoryProxyCreationIterator, error) {

	logs, sub, err := _GnosisSafeProxyFactory.contract.FilterLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return &GnosisSafeProxyFactoryProxyCreationIterator{contract: _GnosisSafeProxyFactory.contract, event: "ProxyCreation", logs: logs, sub: sub}, nil
}

// WatchProxyCreation is a free log subscription operation binding the contract event 0xa38789425dbeee0239e16ff2d2567e31720127fbc6430758c1a4efc6aef29f80.
//
// Solidity: event ProxyCreation(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) WatchProxyCreation(opts *bind.WatchOpts, sink chan<- *GnosisSafeProxyFactoryProxyCreation) (event.Subscription, error) {

	logs, sub, err := _GnosisSafeProxyFactory.contract.WatchLogs(opts, "ProxyCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GnosisSafeProxyFactoryProxyCreation)
				if err := _GnosisSafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
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

// ParseProxyCreation is a log parse operation binding the contract event 0xa38789425dbeee0239e16ff2d2567e31720127fbc6430758c1a4efc6aef29f80.
//
// Solidity: event ProxyCreation(address proxy)
func (_GnosisSafeProxyFactory *GnosisSafeProxyFactoryFilterer) ParseProxyCreation(log types.Log) (*GnosisSafeProxyFactoryProxyCreation, error) {
	event := new(GnosisSafeProxyFactoryProxyCreation)
	if err := _GnosisSafeProxyFactory.contract.UnpackLog(event, "ProxyCreation", log); err != nil {
		return nil, err
	}
	return event, nil
}
