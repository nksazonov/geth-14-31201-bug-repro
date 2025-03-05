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

// DummyPreShanghaiMetaData contains all meta data concerning the DummyPreShanghai contract.
var DummyPreShanghaiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"}],\"name\":\"setX\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600a600081905550610133806100286000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c80630c55699c1460375780634018d9aa146051575b600080fd5b603d6069565b604051604891906090565b60405180910390f35b606760048036038101906063919060d5565b606f565b005b60005481565b8060008190555050565b6000819050919050565b608a816079565b82525050565b600060208201905060a360008301846083565b92915050565b600080fd5b60b5816079565b811460bf57600080fd5b50565b60008135905060cf8160ae565b92915050565b60006020828403121560e85760e760a9565b5b600060f48482850160c2565b9150509291505056fea2646970667358221220d8a6d25dc0089902e25f0a18a39513a876ccd33b0373fbca08ef602cbf6478a464736f6c63430008130033",
}

// DummyPreShanghaiABI is the input ABI used to generate the binding from.
// Deprecated: Use DummyPreShanghaiMetaData.ABI instead.
var DummyPreShanghaiABI = DummyPreShanghaiMetaData.ABI

// DummyPreShanghaiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DummyPreShanghaiMetaData.Bin instead.
var DummyPreShanghaiBin = DummyPreShanghaiMetaData.Bin

// DeployDummyPreShanghai deploys a new Ethereum contract, binding an instance of DummyPreShanghai to it.
func DeployDummyPreShanghai(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DummyPreShanghai, error) {
	parsed, err := DummyPreShanghaiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DummyPreShanghaiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DummyPreShanghai{DummyPreShanghaiCaller: DummyPreShanghaiCaller{contract: contract}, DummyPreShanghaiTransactor: DummyPreShanghaiTransactor{contract: contract}, DummyPreShanghaiFilterer: DummyPreShanghaiFilterer{contract: contract}}, nil
}

// DummyPreShanghai is an auto generated Go binding around an Ethereum contract.
type DummyPreShanghai struct {
	DummyPreShanghaiCaller     // Read-only binding to the contract
	DummyPreShanghaiTransactor // Write-only binding to the contract
	DummyPreShanghaiFilterer   // Log filterer for contract events
}

// DummyPreShanghaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DummyPreShanghaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPreShanghaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DummyPreShanghaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPreShanghaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DummyPreShanghaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DummyPreShanghaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DummyPreShanghaiSession struct {
	Contract     *DummyPreShanghai // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DummyPreShanghaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DummyPreShanghaiCallerSession struct {
	Contract *DummyPreShanghaiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DummyPreShanghaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DummyPreShanghaiTransactorSession struct {
	Contract     *DummyPreShanghaiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DummyPreShanghaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DummyPreShanghaiRaw struct {
	Contract *DummyPreShanghai // Generic contract binding to access the raw methods on
}

// DummyPreShanghaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DummyPreShanghaiCallerRaw struct {
	Contract *DummyPreShanghaiCaller // Generic read-only contract binding to access the raw methods on
}

// DummyPreShanghaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DummyPreShanghaiTransactorRaw struct {
	Contract *DummyPreShanghaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDummyPreShanghai creates a new instance of DummyPreShanghai, bound to a specific deployed contract.
func NewDummyPreShanghai(address common.Address, backend bind.ContractBackend) (*DummyPreShanghai, error) {
	contract, err := bindDummyPreShanghai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DummyPreShanghai{DummyPreShanghaiCaller: DummyPreShanghaiCaller{contract: contract}, DummyPreShanghaiTransactor: DummyPreShanghaiTransactor{contract: contract}, DummyPreShanghaiFilterer: DummyPreShanghaiFilterer{contract: contract}}, nil
}

// NewDummyPreShanghaiCaller creates a new read-only instance of DummyPreShanghai, bound to a specific deployed contract.
func NewDummyPreShanghaiCaller(address common.Address, caller bind.ContractCaller) (*DummyPreShanghaiCaller, error) {
	contract, err := bindDummyPreShanghai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DummyPreShanghaiCaller{contract: contract}, nil
}

// NewDummyPreShanghaiTransactor creates a new write-only instance of DummyPreShanghai, bound to a specific deployed contract.
func NewDummyPreShanghaiTransactor(address common.Address, transactor bind.ContractTransactor) (*DummyPreShanghaiTransactor, error) {
	contract, err := bindDummyPreShanghai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DummyPreShanghaiTransactor{contract: contract}, nil
}

// NewDummyPreShanghaiFilterer creates a new log filterer instance of DummyPreShanghai, bound to a specific deployed contract.
func NewDummyPreShanghaiFilterer(address common.Address, filterer bind.ContractFilterer) (*DummyPreShanghaiFilterer, error) {
	contract, err := bindDummyPreShanghai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DummyPreShanghaiFilterer{contract: contract}, nil
}

// bindDummyPreShanghai binds a generic wrapper to an already deployed contract.
func bindDummyPreShanghai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DummyPreShanghaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyPreShanghai *DummyPreShanghaiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyPreShanghai.Contract.DummyPreShanghaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyPreShanghai *DummyPreShanghaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.DummyPreShanghaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyPreShanghai *DummyPreShanghaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.DummyPreShanghaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DummyPreShanghai *DummyPreShanghaiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DummyPreShanghai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DummyPreShanghai *DummyPreShanghaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DummyPreShanghai *DummyPreShanghaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.contract.Transact(opts, method, params...)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPreShanghai *DummyPreShanghaiCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DummyPreShanghai.contract.Call(opts, &out, "x")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPreShanghai *DummyPreShanghaiSession) X() (*big.Int, error) {
	return _DummyPreShanghai.Contract.X(&_DummyPreShanghai.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_DummyPreShanghai *DummyPreShanghaiCallerSession) X() (*big.Int, error) {
	return _DummyPreShanghai.Contract.X(&_DummyPreShanghai.CallOpts)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPreShanghai *DummyPreShanghaiTransactor) SetX(opts *bind.TransactOpts, _x *big.Int) (*types.Transaction, error) {
	return _DummyPreShanghai.contract.Transact(opts, "setX", _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPreShanghai *DummyPreShanghaiSession) SetX(_x *big.Int) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.SetX(&_DummyPreShanghai.TransactOpts, _x)
}

// SetX is a paid mutator transaction binding the contract method 0x4018d9aa.
//
// Solidity: function setX(uint256 _x) returns()
func (_DummyPreShanghai *DummyPreShanghaiTransactorSession) SetX(_x *big.Int) (*types.Transaction, error) {
	return _DummyPreShanghai.Contract.SetX(&_DummyPreShanghai.TransactOpts, _x)
}
