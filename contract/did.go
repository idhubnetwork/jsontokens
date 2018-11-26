// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package did

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

// DidABI is the input ABI used to generate the binding from.
const DidABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"validDelegate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"setAttribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"}],\"name\":\"identityOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"revokeDelegateSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegateSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"name\":\"delegate\",\"type\":\"address\"},{\"name\":\"validity\",\"type\":\"uint256\"}],\"name\":\"addDelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"sigV\",\"type\":\"uint8\"},{\"name\":\"sigR\",\"type\":\"bytes32\"},{\"name\":\"sigS\",\"type\":\"bytes32\"},{\"name\":\"name\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"revokeAttributeSigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"identity\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"changed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDOwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"delegateType\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDDelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"validTo\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"previousChange\",\"type\":\"uint256\"}],\"name\":\"DIDAttributeChanged\",\"type\":\"event\"}]"

// Did is an auto generated Go binding around an Ethereum contract.
type Did struct {
	DidCaller     // Read-only binding to the contract
	DidTransactor // Write-only binding to the contract
	DidFilterer   // Log filterer for contract events
}

// DidCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidSession struct {
	Contract     *Did              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidCallerSession struct {
	Contract *DidCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidTransactorSession struct {
	Contract     *DidTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidRaw struct {
	Contract *Did // Generic contract binding to access the raw methods on
}

// DidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidCallerRaw struct {
	Contract *DidCaller // Generic read-only contract binding to access the raw methods on
}

// DidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidTransactorRaw struct {
	Contract *DidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDid creates a new instance of Did, bound to a specific deployed contract.
func NewDid(address common.Address, backend bind.ContractBackend) (*Did, error) {
	contract, err := bindDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Did{DidCaller: DidCaller{contract: contract}, DidTransactor: DidTransactor{contract: contract}, DidFilterer: DidFilterer{contract: contract}}, nil
}

// NewDidCaller creates a new read-only instance of Did, bound to a specific deployed contract.
func NewDidCaller(address common.Address, caller bind.ContractCaller) (*DidCaller, error) {
	contract, err := bindDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidCaller{contract: contract}, nil
}

// NewDidTransactor creates a new write-only instance of Did, bound to a specific deployed contract.
func NewDidTransactor(address common.Address, transactor bind.ContractTransactor) (*DidTransactor, error) {
	contract, err := bindDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidTransactor{contract: contract}, nil
}

// NewDidFilterer creates a new log filterer instance of Did, bound to a specific deployed contract.
func NewDidFilterer(address common.Address, filterer bind.ContractFilterer) (*DidFilterer, error) {
	contract, err := bindDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidFilterer{contract: contract}, nil
}

// bindDid binds a generic wrapper to an already deployed contract.
func bindDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.DidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.DidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Did *DidCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Did.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Did *DidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Did.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Did *DidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Did.Contract.contract.Transact(opts, method, params...)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed( address) constant returns(uint256)
func (_Did *DidCaller) Changed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "changed", arg0)
	return *ret0, err
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed( address) constant returns(uint256)
func (_Did *DidSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Changed(&_Did.CallOpts, arg0)
}

// Changed is a free data retrieval call binding the contract method 0xf96d0f9f.
//
// Solidity: function changed( address) constant returns(uint256)
func (_Did *DidCallerSession) Changed(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Changed(&_Did.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates( address,  bytes32,  address) constant returns(uint256)
func (_Did *DidCaller) Delegates(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "delegates", arg0, arg1, arg2)
	return *ret0, err
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates( address,  bytes32,  address) constant returns(uint256)
func (_Did *DidSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Did.Contract.Delegates(&_Did.CallOpts, arg0, arg1, arg2)
}

// Delegates is a free data retrieval call binding the contract method 0x0d44625b.
//
// Solidity: function delegates( address,  bytes32,  address) constant returns(uint256)
func (_Did *DidCallerSession) Delegates(arg0 common.Address, arg1 [32]byte, arg2 common.Address) (*big.Int, error) {
	return _Did.Contract.Delegates(&_Did.CallOpts, arg0, arg1, arg2)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCaller) IdentityOwner(opts *bind.CallOpts, identity common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "identityOwner", identity)
	return *ret0, err
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// IdentityOwner is a free data retrieval call binding the contract method 0x8733d4e8.
//
// Solidity: function identityOwner(identity address) constant returns(address)
func (_Did *DidCallerSession) IdentityOwner(identity common.Address) (common.Address, error) {
	return _Did.Contract.IdentityOwner(&_Did.CallOpts, identity)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCaller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "nonce", arg0)
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce( address) constant returns(uint256)
func (_Did *DidCallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _Did.Contract.Nonce(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCaller) Owners(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x022914a7.
//
// Solidity: function owners( address) constant returns(address)
func (_Did *DidCallerSession) Owners(arg0 common.Address) (common.Address, error) {
	return _Did.Contract.Owners(&_Did.CallOpts, arg0)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(identity address, delegateType bytes32, delegate address) constant returns(bool)
func (_Did *DidCaller) ValidDelegate(opts *bind.CallOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Did.contract.Call(opts, out, "validDelegate", identity, delegateType, delegate)
	return *ret0, err
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(identity address, delegateType bytes32, delegate address) constant returns(bool)
func (_Did *DidSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _Did.Contract.ValidDelegate(&_Did.CallOpts, identity, delegateType, delegate)
}

// ValidDelegate is a free data retrieval call binding the contract method 0x622b2a3c.
//
// Solidity: function validDelegate(identity address, delegateType bytes32, delegate address) constant returns(bool)
func (_Did *DidCallerSession) ValidDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (bool, error) {
	return _Did.Contract.ValidDelegate(&_Did.CallOpts, identity, delegateType, delegate)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(identity address, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidTransactor) AddDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addDelegate", identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(identity address, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddDelegate(&_Did.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegate is a paid mutator transaction binding the contract method 0xa7068d66.
//
// Solidity: function addDelegate(identity address, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidTransactorSession) AddDelegate(identity common.Address, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddDelegate(&_Did.TransactOpts, identity, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidTransactor) AddDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "addDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddDelegateSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// AddDelegateSigned is a paid mutator transaction binding the contract method 0x9c2c1b2b.
//
// Solidity: function addDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address, validity uint256) returns()
func (_Did *DidTransactorSession) AddDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.AddDelegateSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate, validity)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwner(opts *bind.TransactOpts, identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwner", identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(identity address, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwner(identity common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwner(&_Did.TransactOpts, identity, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactor) ChangeOwnerSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "changeOwnerSigned", identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// ChangeOwnerSigned is a paid mutator transaction binding the contract method 0x240cf1fa.
//
// Solidity: function changeOwnerSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, newOwner address) returns()
func (_Did *DidTransactorSession) ChangeOwnerSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, newOwner common.Address) (*types.Transaction, error) {
	return _Did.Contract.ChangeOwnerSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, newOwner)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidTransactor) RevokeAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttribute", identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttribute is a paid mutator transaction binding the contract method 0x00c023da.
//
// Solidity: function revokeAttribute(identity address, name bytes32, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttribute(identity common.Address, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttribute(&_Did.TransactOpts, identity, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidTransactor) RevokeAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeAttributeSigned", identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeAttributeSigned is a paid mutator transaction binding the contract method 0xe476af5c.
//
// Solidity: function revokeAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes) returns()
func (_Did *DidTransactorSession) RevokeAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte) (*types.Transaction, error) {
	return _Did.Contract.RevokeAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(identity address, delegateType bytes32, delegate address) returns()
func (_Did *DidTransactor) RevokeDelegate(opts *bind.TransactOpts, identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeDelegate", identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(identity address, delegateType bytes32, delegate address) returns()
func (_Did *DidSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.Contract.RevokeDelegate(&_Did.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegate is a paid mutator transaction binding the contract method 0x80b29f7c.
//
// Solidity: function revokeDelegate(identity address, delegateType bytes32, delegate address) returns()
func (_Did *DidTransactorSession) RevokeDelegate(identity common.Address, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.Contract.RevokeDelegate(&_Did.TransactOpts, identity, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address) returns()
func (_Did *DidTransactor) RevokeDelegateSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "revokeDelegateSigned", identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address) returns()
func (_Did *DidSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.Contract.RevokeDelegateSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// RevokeDelegateSigned is a paid mutator transaction binding the contract method 0x93072684.
//
// Solidity: function revokeDelegateSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, delegateType bytes32, delegate address) returns()
func (_Did *DidTransactorSession) RevokeDelegateSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, delegateType [32]byte, delegate common.Address) (*types.Transaction, error) {
	return _Did.Contract.RevokeDelegateSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, delegateType, delegate)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttribute(opts *bind.TransactOpts, identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttribute", identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttribute is a paid mutator transaction binding the contract method 0x7ad4b0a4.
//
// Solidity: function setAttribute(identity address, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttribute(identity common.Address, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttribute(&_Did.TransactOpts, identity, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactor) SetAttributeSigned(opts *bind.TransactOpts, identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.contract.Transact(opts, "setAttributeSigned", identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// SetAttributeSigned is a paid mutator transaction binding the contract method 0x123b5e98.
//
// Solidity: function setAttributeSigned(identity address, sigV uint8, sigR bytes32, sigS bytes32, name bytes32, value bytes, validity uint256) returns()
func (_Did *DidTransactorSession) SetAttributeSigned(identity common.Address, sigV uint8, sigR [32]byte, sigS [32]byte, name [32]byte, value []byte, validity *big.Int) (*types.Transaction, error) {
	return _Did.Contract.SetAttributeSigned(&_Did.TransactOpts, identity, sigV, sigR, sigS, name, value, validity)
}

// DidDIDAttributeChangedIterator is returned from FilterDIDAttributeChanged and is used to iterate over the raw logs and unpacked data for DIDAttributeChanged events raised by the Did contract.
type DidDIDAttributeChangedIterator struct {
	Event *DidDIDAttributeChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDAttributeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDAttributeChanged)
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
		it.Event = new(DidDIDAttributeChanged)
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
func (it *DidDIDAttributeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDAttributeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDAttributeChanged represents a DIDAttributeChanged event raised by the Did contract.
type DidDIDAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDAttributeChanged is a free log retrieval operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes32, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDAttributeChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDAttributeChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDAttributeChangedIterator{contract: _Did.contract, event: "DIDAttributeChanged", logs: logs, sub: sub}, nil
}

// WatchDIDAttributeChanged is a free log subscription operation binding the contract event 0x18ab6b2ae3d64306c00ce663125f2bd680e441a098de1635bd7ad8b0d44965e4.
//
// Solidity: e DIDAttributeChanged(identity indexed address, name bytes32, value bytes, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDAttributeChanged(opts *bind.WatchOpts, sink chan<- *DidDIDAttributeChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDAttributeChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDAttributeChanged)
				if err := _Did.contract.UnpackLog(event, "DIDAttributeChanged", log); err != nil {
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

// DidDIDDelegateChangedIterator is returned from FilterDIDDelegateChanged and is used to iterate over the raw logs and unpacked data for DIDDelegateChanged events raised by the Did contract.
type DidDIDDelegateChangedIterator struct {
	Event *DidDIDDelegateChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDDelegateChanged)
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
		it.Event = new(DidDIDDelegateChanged)
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
func (it *DidDIDDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDDelegateChanged represents a DIDDelegateChanged event raised by the Did contract.
type DidDIDDelegateChanged struct {
	Identity       common.Address
	DelegateType   [32]byte
	Delegate       common.Address
	ValidTo        *big.Int
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDDelegateChanged is a free log retrieval operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: e DIDDelegateChanged(identity indexed address, delegateType bytes32, delegate address, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) FilterDIDDelegateChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDDelegateChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDDelegateChangedIterator{contract: _Did.contract, event: "DIDDelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDIDDelegateChanged is a free log subscription operation binding the contract event 0x5a5084339536bcab65f20799fcc58724588145ca054bd2be626174b27ba156f7.
//
// Solidity: e DIDDelegateChanged(identity indexed address, delegateType bytes32, delegate address, validTo uint256, previousChange uint256)
func (_Did *DidFilterer) WatchDIDDelegateChanged(opts *bind.WatchOpts, sink chan<- *DidDIDDelegateChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDDelegateChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDDelegateChanged)
				if err := _Did.contract.UnpackLog(event, "DIDDelegateChanged", log); err != nil {
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

// DidDIDOwnerChangedIterator is returned from FilterDIDOwnerChanged and is used to iterate over the raw logs and unpacked data for DIDOwnerChanged events raised by the Did contract.
type DidDIDOwnerChangedIterator struct {
	Event *DidDIDOwnerChanged // Event containing the contract specifics and raw log

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
func (it *DidDIDOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidDIDOwnerChanged)
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
		it.Event = new(DidDIDOwnerChanged)
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
func (it *DidDIDOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidDIDOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidDIDOwnerChanged represents a DIDOwnerChanged event raised by the Did contract.
type DidDIDOwnerChanged struct {
	Identity       common.Address
	Owner          common.Address
	PreviousChange *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDIDOwnerChanged is a free log retrieval operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: e DIDOwnerChanged(identity indexed address, owner address, previousChange uint256)
func (_Did *DidFilterer) FilterDIDOwnerChanged(opts *bind.FilterOpts, identity []common.Address) (*DidDIDOwnerChangedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.FilterLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return &DidDIDOwnerChangedIterator{contract: _Did.contract, event: "DIDOwnerChanged", logs: logs, sub: sub}, nil
}

// WatchDIDOwnerChanged is a free log subscription operation binding the contract event 0x38a5a6e68f30ed1ab45860a4afb34bcb2fc00f22ca462d249b8a8d40cda6f7a3.
//
// Solidity: e DIDOwnerChanged(identity indexed address, owner address, previousChange uint256)
func (_Did *DidFilterer) WatchDIDOwnerChanged(opts *bind.WatchOpts, sink chan<- *DidDIDOwnerChanged, identity []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _Did.contract.WatchLogs(opts, "DIDOwnerChanged", identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidDIDOwnerChanged)
				if err := _Did.contract.UnpackLog(event, "DIDOwnerChanged", log); err != nil {
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
