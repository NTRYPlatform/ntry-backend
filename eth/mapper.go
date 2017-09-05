// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package eth

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MapperABI is the input ABI used to generate the binding from.
const MapperABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"secondary\",\"type\":\"address\"}],\"name\":\"mapAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"secondaryInUse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"primaryToSecondary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"primary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"secondary\",\"type\":\"address\"}],\"name\":\"AddressMapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"code\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Error\",\"type\":\"event\"}]"

// Mapper is an auto generated Go binding around an Ethereum contract.
type Mapper struct {
	MapperCaller     // Read-only binding to the contract
	MapperTransactor // Write-only binding to the contract
}

// MapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapperSession struct {
	Contract     *Mapper           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapperCallerSession struct {
	Contract *MapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapperTransactorSession struct {
	Contract     *MapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapperRaw struct {
	Contract *Mapper // Generic contract binding to access the raw methods on
}

// MapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapperCallerRaw struct {
	Contract *MapperCaller // Generic read-only contract binding to access the raw methods on
}

// MapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapperTransactorRaw struct {
	Contract *MapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapper creates a new instance of Mapper, bound to a specific deployed contract.
func NewMapper(address common.Address, backend bind.ContractBackend) (*Mapper, error) {
	contract, err := bindMapper(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mapper{MapperCaller: MapperCaller{contract: contract}, MapperTransactor: MapperTransactor{contract: contract}}, nil
}

// NewMapperCaller creates a new read-only instance of Mapper, bound to a specific deployed contract.
func NewMapperCaller(address common.Address, caller bind.ContractCaller) (*MapperCaller, error) {
	contract, err := bindMapper(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MapperCaller{contract: contract}, nil
}

// NewMapperTransactor creates a new write-only instance of Mapper, bound to a specific deployed contract.
func NewMapperTransactor(address common.Address, transactor bind.ContractTransactor) (*MapperTransactor, error) {
	contract, err := bindMapper(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MapperTransactor{contract: contract}, nil
}

// bindMapper binds a generic wrapper to an already deployed contract.
func bindMapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mapper *MapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mapper.Contract.MapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mapper *MapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mapper.Contract.MapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mapper *MapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mapper.Contract.MapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mapper *MapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Mapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mapper *MapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mapper *MapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mapper.Contract.contract.Transact(opts, method, params...)
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Mapper *MapperCaller) PrimaryToSecondary(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Mapper.contract.Call(opts, out, "primaryToSecondary", arg0)
	return *ret0, err
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Mapper *MapperSession) PrimaryToSecondary(arg0 common.Address) (common.Address, error) {
	return _Mapper.Contract.PrimaryToSecondary(&_Mapper.CallOpts, arg0)
}

// PrimaryToSecondary is a free data retrieval call binding the contract method 0xd352221e.
//
// Solidity: function primaryToSecondary( address) constant returns(address)
func (_Mapper *MapperCallerSession) PrimaryToSecondary(arg0 common.Address) (common.Address, error) {
	return _Mapper.Contract.PrimaryToSecondary(&_Mapper.CallOpts, arg0)
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Mapper *MapperCaller) SecondaryInUse(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Mapper.contract.Call(opts, out, "secondaryInUse", arg0)
	return *ret0, err
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Mapper *MapperSession) SecondaryInUse(arg0 common.Address) (bool, error) {
	return _Mapper.Contract.SecondaryInUse(&_Mapper.CallOpts, arg0)
}

// SecondaryInUse is a free data retrieval call binding the contract method 0x6a50654a.
//
// Solidity: function secondaryInUse( address) constant returns(bool)
func (_Mapper *MapperCallerSession) SecondaryInUse(arg0 common.Address) (bool, error) {
	return _Mapper.Contract.SecondaryInUse(&_Mapper.CallOpts, arg0)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Mapper *MapperTransactor) MapAddress(opts *bind.TransactOpts, secondary common.Address) (*types.Transaction, error) {
	return _Mapper.contract.Transact(opts, "mapAddress", secondary)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Mapper *MapperSession) MapAddress(secondary common.Address) (*types.Transaction, error) {
	return _Mapper.Contract.MapAddress(&_Mapper.TransactOpts, secondary)
}

// MapAddress is a paid mutator transaction binding the contract method 0x4a270f47.
//
// Solidity: function mapAddress(secondary address) returns()
func (_Mapper *MapperTransactorSession) MapAddress(secondary common.Address) (*types.Transaction, error) {
	return _Mapper.Contract.MapAddress(&_Mapper.TransactOpts, secondary)
}
