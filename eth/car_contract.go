// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package eth

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CarContractFunctionsABI is the input ABI used to generate the binding from.
const CarContractFunctionsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalEarned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateAppAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ntryApp\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"cid\",\"type\":\"uint256\"}],\"name\":\"getDeal\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ntryToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cid\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"},{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"carDeal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sellerAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"buyerAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dealID\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dealfee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dealHash\",\"type\":\"bytes32\"}],\"name\":\"LogDeal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AppAccountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"TokenAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldValue\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"ContractFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"ReentrancyAttemp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CarContractFunctions is an auto generated Go binding around an Ethereum contract.
type CarContractFunctions struct {
	CarContractFunctionsCaller     // Read-only binding to the contract
	CarContractFunctionsTransactor // Write-only binding to the contract
}

// CarContractFunctionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarContractFunctionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarContractFunctionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarContractFunctionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarContractFunctionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarContractFunctionsSession struct {
	Contract     *CarContractFunctions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CarContractFunctionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarContractFunctionsCallerSession struct {
	Contract *CarContractFunctionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// CarContractFunctionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarContractFunctionsTransactorSession struct {
	Contract     *CarContractFunctionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// CarContractFunctionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarContractFunctionsRaw struct {
	Contract *CarContractFunctions // Generic contract binding to access the raw methods on
}

// CarContractFunctionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarContractFunctionsCallerRaw struct {
	Contract *CarContractFunctionsCaller // Generic read-only contract binding to access the raw methods on
}

// CarContractFunctionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarContractFunctionsTransactorRaw struct {
	Contract *CarContractFunctionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCarContractFunctions creates a new instance of CarContractFunctions, bound to a specific deployed contract.
func NewCarContractFunctions(address common.Address, backend bind.ContractBackend) (*CarContractFunctions, error) {
	contract, err := bindCarContractFunctions(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CarContractFunctions{CarContractFunctionsCaller: CarContractFunctionsCaller{contract: contract}, CarContractFunctionsTransactor: CarContractFunctionsTransactor{contract: contract}}, nil
}

// NewCarContractFunctionsCaller creates a new read-only instance of CarContractFunctions, bound to a specific deployed contract.
func NewCarContractFunctionsCaller(address common.Address, caller bind.ContractCaller) (*CarContractFunctionsCaller, error) {
	contract, err := bindCarContractFunctions(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &CarContractFunctionsCaller{contract: contract}, nil
}

// NewCarContractFunctionsTransactor creates a new write-only instance of CarContractFunctions, bound to a specific deployed contract.
func NewCarContractFunctionsTransactor(address common.Address, transactor bind.ContractTransactor) (*CarContractFunctionsTransactor, error) {
	contract, err := bindCarContractFunctions(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &CarContractFunctionsTransactor{contract: contract}, nil
}

// bindCarContractFunctions binds a generic wrapper to an already deployed contract.
func bindCarContractFunctions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CarContractFunctionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarContractFunctions *CarContractFunctionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CarContractFunctions.Contract.CarContractFunctionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarContractFunctions *CarContractFunctionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.CarContractFunctionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarContractFunctions *CarContractFunctionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.CarContractFunctionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarContractFunctions *CarContractFunctionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CarContractFunctions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarContractFunctions *CarContractFunctionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarContractFunctions *CarContractFunctionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.contract.Transact(opts, method, params...)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "fee")
	return *ret0, err
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsSession) Fee() (*big.Int, error) {
	return _CarContractFunctions.Contract.Fee(&_CarContractFunctions.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCallerSession) Fee() (*big.Int, error) {
	return _CarContractFunctions.Contract.Fee(&_CarContractFunctions.CallOpts)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(cid uint256) constant returns(address, address, bytes32)
func (_CarContractFunctions *CarContractFunctionsCaller) GetDeal(opts *bind.CallOpts, cid *big.Int) (common.Address, common.Address, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _CarContractFunctions.contract.Call(opts, out, "getDeal", cid)
	return *ret0, *ret1, *ret2, err
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(cid uint256) constant returns(address, address, bytes32)
func (_CarContractFunctions *CarContractFunctionsSession) GetDeal(cid *big.Int) (common.Address, common.Address, [32]byte, error) {
	return _CarContractFunctions.Contract.GetDeal(&_CarContractFunctions.CallOpts, cid)
}

// GetDeal is a free data retrieval call binding the contract method 0x82fd5bac.
//
// Solidity: function getDeal(cid uint256) constant returns(address, address, bytes32)
func (_CarContractFunctions *CarContractFunctionsCallerSession) GetDeal(cid *big.Int) (common.Address, common.Address, [32]byte, error) {
	return _CarContractFunctions.Contract.GetDeal(&_CarContractFunctions.CallOpts, cid)
}

// GetID is a free data retrieval call binding the contract method 0x0d18e43c.
//
// Solidity: function getID(_index uint256) constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCaller) GetID(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "getID", _index)
	return *ret0, err
}

// GetID is a free data retrieval call binding the contract method 0x0d18e43c.
//
// Solidity: function getID(_index uint256) constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsSession) GetID(_index *big.Int) (*big.Int, error) {
	return _CarContractFunctions.Contract.GetID(&_CarContractFunctions.CallOpts, _index)
}

// GetID is a free data retrieval call binding the contract method 0x0d18e43c.
//
// Solidity: function getID(_index uint256) constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCallerSession) GetID(_index *big.Int) (*big.Int, error) {
	return _CarContractFunctions.Contract.GetID(&_CarContractFunctions.CallOpts, _index)
}

// NtryApp is a free data retrieval call binding the contract method 0x8004e5d1.
//
// Solidity: function ntryApp() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCaller) NtryApp(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "ntryApp")
	return *ret0, err
}

// NtryApp is a free data retrieval call binding the contract method 0x8004e5d1.
//
// Solidity: function ntryApp() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsSession) NtryApp() (common.Address, error) {
	return _CarContractFunctions.Contract.NtryApp(&_CarContractFunctions.CallOpts)
}

// NtryApp is a free data retrieval call binding the contract method 0x8004e5d1.
//
// Solidity: function ntryApp() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCallerSession) NtryApp() (common.Address, error) {
	return _CarContractFunctions.Contract.NtryApp(&_CarContractFunctions.CallOpts)
}

// NtryToken is a free data retrieval call binding the contract method 0xb8389713.
//
// Solidity: function ntryToken() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCaller) NtryToken(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "ntryToken")
	return *ret0, err
}

// NtryToken is a free data retrieval call binding the contract method 0xb8389713.
//
// Solidity: function ntryToken() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsSession) NtryToken() (common.Address, error) {
	return _CarContractFunctions.Contract.NtryToken(&_CarContractFunctions.CallOpts)
}

// NtryToken is a free data retrieval call binding the contract method 0xb8389713.
//
// Solidity: function ntryToken() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCallerSession) NtryToken() (common.Address, error) {
	return _CarContractFunctions.Contract.NtryToken(&_CarContractFunctions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsSession) Owner() (common.Address, error) {
	return _CarContractFunctions.Contract.Owner(&_CarContractFunctions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CarContractFunctions *CarContractFunctionsCallerSession) Owner() (common.Address, error) {
	return _CarContractFunctions.Contract.Owner(&_CarContractFunctions.CallOpts)
}

// TotalContracts is a free data retrieval call binding the contract method 0xa09037a9.
//
// Solidity: function totalContracts() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCaller) TotalContracts(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "totalContracts")
	return *ret0, err
}

// TotalContracts is a free data retrieval call binding the contract method 0xa09037a9.
//
// Solidity: function totalContracts() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsSession) TotalContracts() (*big.Int, error) {
	return _CarContractFunctions.Contract.TotalContracts(&_CarContractFunctions.CallOpts)
}

// TotalContracts is a free data retrieval call binding the contract method 0xa09037a9.
//
// Solidity: function totalContracts() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCallerSession) TotalContracts() (*big.Int, error) {
	return _CarContractFunctions.Contract.TotalContracts(&_CarContractFunctions.CallOpts)
}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCaller) TotalEarned(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CarContractFunctions.contract.Call(opts, out, "totalEarned")
	return *ret0, err
}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsSession) TotalEarned() (*big.Int, error) {
	return _CarContractFunctions.Contract.TotalEarned(&_CarContractFunctions.CallOpts)
}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() constant returns(uint256)
func (_CarContractFunctions *CarContractFunctionsCallerSession) TotalEarned() (*big.Int, error) {
	return _CarContractFunctions.Contract.TotalEarned(&_CarContractFunctions.CallOpts)
}

// CarDeal is a paid mutator transaction binding the contract method 0xd69d8724.
//
// Solidity: function carDeal(cid uint256, _seller address, _buyer address, _hash bytes32) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactor) CarDeal(opts *bind.TransactOpts, cid *big.Int, _seller common.Address, _buyer common.Address, _hash [32]byte) (*types.Transaction, error) {
	return _CarContractFunctions.contract.Transact(opts, "carDeal", cid, _seller, _buyer, _hash)
}

// CarDeal is a paid mutator transaction binding the contract method 0xd69d8724.
//
// Solidity: function carDeal(cid uint256, _seller address, _buyer address, _hash bytes32) returns(bool)
func (_CarContractFunctions *CarContractFunctionsSession) CarDeal(cid *big.Int, _seller common.Address, _buyer common.Address, _hash [32]byte) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.CarDeal(&_CarContractFunctions.TransactOpts, cid, _seller, _buyer, _hash)
}

// CarDeal is a paid mutator transaction binding the contract method 0xd69d8724.
//
// Solidity: function carDeal(cid uint256, _seller address, _buyer address, _hash bytes32) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactorSession) CarDeal(cid *big.Int, _seller common.Address, _buyer common.Address, _hash [32]byte) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.CarDeal(&_CarContractFunctions.TransactOpts, cid, _seller, _buyer, _hash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CarContractFunctions *CarContractFunctionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CarContractFunctions *CarContractFunctionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.TransferOwnership(&_CarContractFunctions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CarContractFunctions *CarContractFunctionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.TransferOwnership(&_CarContractFunctions.TransactOpts, newOwner)
}

// UpdateAppAccount is a paid mutator transaction binding the contract method 0x7544698a.
//
// Solidity: function updateAppAccount(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactor) UpdateAppAccount(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.contract.Transact(opts, "updateAppAccount", _address)
}

// UpdateAppAccount is a paid mutator transaction binding the contract method 0x7544698a.
//
// Solidity: function updateAppAccount(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsSession) UpdateAppAccount(_address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateAppAccount(&_CarContractFunctions.TransactOpts, _address)
}

// UpdateAppAccount is a paid mutator transaction binding the contract method 0x7544698a.
//
// Solidity: function updateAppAccount(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactorSession) UpdateAppAccount(_address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateAppAccount(&_CarContractFunctions.TransactOpts, _address)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_value uint256) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactor) UpdateFee(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _CarContractFunctions.contract.Transact(opts, "updateFee", _value)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_value uint256) returns(bool)
func (_CarContractFunctions *CarContractFunctionsSession) UpdateFee(_value *big.Int) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateFee(&_CarContractFunctions.TransactOpts, _value)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(_value uint256) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactorSession) UpdateFee(_value *big.Int) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateFee(&_CarContractFunctions.TransactOpts, _value)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x6691461a.
//
// Solidity: function updateTokenAddress(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactor) UpdateTokenAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.contract.Transact(opts, "updateTokenAddress", _address)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x6691461a.
//
// Solidity: function updateTokenAddress(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsSession) UpdateTokenAddress(_address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateTokenAddress(&_CarContractFunctions.TransactOpts, _address)
}

// UpdateTokenAddress is a paid mutator transaction binding the contract method 0x6691461a.
//
// Solidity: function updateTokenAddress(_address address) returns(bool)
func (_CarContractFunctions *CarContractFunctionsTransactorSession) UpdateTokenAddress(_address common.Address) (*types.Transaction, error) {
	return _CarContractFunctions.Contract.UpdateTokenAddress(&_CarContractFunctions.TransactOpts, _address)
}
