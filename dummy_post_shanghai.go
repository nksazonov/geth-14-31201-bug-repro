// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
)

// DummyPostShanghaiMetaData contains all meta data concerning the DummyPostShanghai contract.
var DummyPostShanghaiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"setX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50600a5f81905550610126806100245f395ff3fe6080604052348015600e575f80fd5b50600436106030575f3560e01c80630c55699c1460345780634018d9aa14604e575b5f80fd5b603a6066565b60405160459190608a565b60405180910390f35b606460048036038101906060919060ca565b606b565b005b5f5481565b805f8190555050565b5f819050919050565b6084816074565b82525050565b5f602082019050609b5f830184607d565b92915050565b5f80fd5b60ac816074565b811460b5575f80fd5b50565b5f8135905060c48160a5565b92915050565b5f6020828403121560dc5760db60a1565b5b5f60e78482850160b8565b9150509291505056fea2646970667358221220c0e8bca7db5a3be0846a385cee21f50754c246e58903c30c06044d6aae37919864736f6c63430008140033",
}

// DummyPostShanghaiABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyPostShanghaiMetaData.ABI instead.
var DummyPostShanghaiABI = DummyPostShanghaiMetaData.ABI

// DummyPostShanghaiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DummyPostShanghaiMetaData.Bin instead.
var DummyPostShanghaiBin = DummyPostShanghaiMetaData.Bin

// DeployDummyPostShanghai deploys a new Ethereum contract, binding an instance of DummyPostShanghai to it.
func DeployDummyPostShanghai(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DummyPostShanghai, error) {
	parsed, err := DummyPostShanghaiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DummyPostShanghaiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyPostShanghai{DummyPostShanghaiCaller: DummyPostShanghaiCaller{contract: contract}, DummyPostShanghaiTransactor: DummyPostShanghaiTransactor{contract: contract}, DummyPostShanghaiFilterer: DummyPostShanghaiFilterer{contract: contract}}, nil
}

// DummyPostShanghai is an auto generated Go binding around an Ethereum contract.
type DummyPostShanghai struct {
	DummyPostShanghaiCaller     // Read-only binding to the contract
	DummyPostShanghaiTransactor // Write-only binding to the contract
	DummyPostShanghaiFilterer   // Log filterer for contract events
}

// DummyPostShanghaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyPostShanghaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPostShanghaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyPostShanghaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPostShanghaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyPostShanghaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPostShanghaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyPostShanghaiSession struct {
	Contract     *DummyPostShanghai // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DummyPostShanghaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyPostShanghaiCallerSession struct {
	Contract *DummyPostShanghaiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DummyPostShanghaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyPostShanghaiTransactorSession struct {
	Contract     *DummyPostShanghaiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DummyPostShanghaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyPostShanghaiRaw struct {
	Contract *DummyPostShanghai // Generic contract binding to access the raw methods on
}

// DummyPostShanghaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyPostShanghaiCallerRaw struct {
	Contract *DummyPostShanghaiCaller // Generic read-only contract binding to access the raw methods on
}

// DummyPostShanghaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyPostShanghaiTransactorRaw struct {
	Contract *DummyPostShanghaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyPostShanghai creates a new instance of DummyPostShanghai, bound to a specific deployed contract.
func NewDummyPostShanghai(address common.Address, backend bind.ContractBackend) (*DummyPostShanghai, error) {
	contract, err := bindDummyPostShanghai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyPostShanghai{DummyPostShanghaiCaller: DummyPostShanghaiCaller{contract: contract}, DummyPostShanghaiTransactor: DummyPostShanghaiTransactor{contract: contract}, DummyPostShanghaiFilterer: DummyPostShanghaiFilterer{contract: contract}}, nil
}

// NewDummyPostShanghaiCaller creates a new read-only instance of DummyPostShanghai, bound to a specific deployed contract.
func NewDummyPostShanghaiCaller(address common.Address, caller bind.ContractCaller) (*DummyPostShanghaiCaller, error) {
	contract, err := bindDummyPostShanghai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyPostShanghaiCaller{contract: contract}, nil
}

// NewDummyPostShanghaiTransactor creates a new write-only instance of DummyPostShanghai, bound to a specific deployed contract.
func NewDummyPostShanghaiTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyPostShanghaiTransactor, error) {
	contract, err := bindDummyPostShanghai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyPostShanghaiTransactor{contract: contract}, nil
}

// NewDummyPostShanghaiFilterer creates a new log filterer instance of DummyPostShanghai, bound to a specific deployed contract.
func NewDummyPostShanghaiFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyPostShanghaiFilterer, error) {
	contract, err := bindDummyPostShanghai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyPostShanghaiFilterer{contract: contract}, nil
}

// bindDummyPostShanghai binds a generic wrapper to an already deployed contract.
func bindDummyPostShanghai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyPostShanghaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyPostShanghai *DummyPostShanghaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyPostShanghai.Contract.DummyPostShanghaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyPostShanghai *DummyPostShanghaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.DummyPostShanghaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyPostShanghai *DummyPostShanghaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.DummyPostShanghaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyPostShanghai *DummyPostShanghaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyPostShanghai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyPostShanghai *DummyPostShanghaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyPostShanghai *DummyPostShanghaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPostShanghai *DummyPostShanghaiCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyPostShanghai.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPostShanghai *DummyPostShanghaiSession) X() (*big.Int, error) {
	return _DummyPostShanghai.Contract.X(&_DummyPostShanghai.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPostShanghai *DummyPostShanghaiCallerSession) X() (*big.Int, error) {
	return _DummyPostShanghai.Contract.X(&_DummyPostShanghai.CallOpts)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPostShanghai *DummyPostShanghaiTransactor) SetX(opts *bind.TransactOpts, _x *big.Int) (*types.Transaction, error) {
	return _DummyPostShanghai.contract.Transact(opts, "setX", _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPostShanghai *DummyPostShanghaiSession) SetX(_x *big.Int) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.SetX(&_DummyPostShanghai.TransactOpts, _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPostShanghai *DummyPostShanghaiTransactorSession) SetX(_x *big.Int) (*types.Transaction, error) {
	return _DummyPostShanghai.Contract.SetX(&_DummyPostShanghai.TransactOpts, _x)
}
