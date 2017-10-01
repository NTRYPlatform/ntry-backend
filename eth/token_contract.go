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

// TokenContractFunctionsABI is the input ABI used to generate the binding from.
const TokenContractFunctionsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"tweakUpgrading\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"setMigrationAgent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"migrationAgent\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalMigrated\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"tweakState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradingEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RealeaseTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Migrate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"Upgrading\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenContractFunctions is an auto generated Go binding around an Ethereum contract.
type TokenContractFunctions struct {
	TokenContractFunctionsCaller     // Read-only binding to the contract
	TokenContractFunctionsTransactor // Write-only binding to the contract
}

// TokenContractFunctionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenContractFunctionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenContractFunctionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenContractFunctionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenContractFunctionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenContractFunctionsSession struct {
	Contract     *TokenContractFunctions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenContractFunctionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenContractFunctionsCallerSession struct {
	Contract *TokenContractFunctionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TokenContractFunctionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenContractFunctionsTransactorSession struct {
	Contract     *TokenContractFunctionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TokenContractFunctionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenContractFunctionsRaw struct {
	Contract *TokenContractFunctions // Generic contract binding to access the raw methods on
}

// TokenContractFunctionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenContractFunctionsCallerRaw struct {
	Contract *TokenContractFunctionsCaller // Generic read-only contract binding to access the raw methods on
}

// TokenContractFunctionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenContractFunctionsTransactorRaw struct {
	Contract *TokenContractFunctionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenContractFunctions creates a new instance of TokenContractFunctions, bound to a specific deployed contract.
func NewTokenContractFunctions(address common.Address, backend bind.ContractBackend) (*TokenContractFunctions, error) {
	contract, err := bindTokenContractFunctions(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenContractFunctions{TokenContractFunctionsCaller: TokenContractFunctionsCaller{contract: contract}, TokenContractFunctionsTransactor: TokenContractFunctionsTransactor{contract: contract}}, nil
}

// NewTokenContractFunctionsCaller creates a new read-only instance of TokenContractFunctions, bound to a specific deployed contract.
func NewTokenContractFunctionsCaller(address common.Address, caller bind.ContractCaller) (*TokenContractFunctionsCaller, error) {
	contract, err := bindTokenContractFunctions(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenContractFunctionsCaller{contract: contract}, nil
}

// NewTokenContractFunctionsTransactor creates a new write-only instance of TokenContractFunctions, bound to a specific deployed contract.
func NewTokenContractFunctionsTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenContractFunctionsTransactor, error) {
	contract, err := bindTokenContractFunctions(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenContractFunctionsTransactor{contract: contract}, nil
}

// bindTokenContractFunctions binds a generic wrapper to an already deployed contract.
func bindTokenContractFunctions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenContractFunctionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenContractFunctions *TokenContractFunctionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenContractFunctions.Contract.TokenContractFunctionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenContractFunctions *TokenContractFunctionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TokenContractFunctionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenContractFunctions *TokenContractFunctionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TokenContractFunctionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenContractFunctions *TokenContractFunctionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenContractFunctions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenContractFunctions *TokenContractFunctionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenContractFunctions *TokenContractFunctionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x9e87fac8.
//
// Solidity: function Paused() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "Paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x9e87fac8.
//
// Solidity: function Paused() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) Paused() (bool, error) {
	return _TokenContractFunctions.Contract.Paused(&_TokenContractFunctions.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x9e87fac8.
//
// Solidity: function Paused() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) Paused() (bool, error) {
	return _TokenContractFunctions.Contract.Paused(&_TokenContractFunctions.CallOpts)
}

// RealeaseTime is a free data retrieval call binding the contract method 0xf784dd61.
//
// Solidity: function RealeaseTime() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsCaller) RealeaseTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "RealeaseTime")
	return *ret0, err
}

// RealeaseTime is a free data retrieval call binding the contract method 0xf784dd61.
//
// Solidity: function RealeaseTime() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsSession) RealeaseTime() (*big.Int, error) {
	return _TokenContractFunctions.Contract.RealeaseTime(&_TokenContractFunctions.CallOpts)
}

// RealeaseTime is a free data retrieval call binding the contract method 0xf784dd61.
//
// Solidity: function RealeaseTime() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) RealeaseTime() (*big.Int, error) {
	return _TokenContractFunctions.Contract.RealeaseTime(&_TokenContractFunctions.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TokenContractFunctions *TokenContractFunctionsCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TokenContractFunctions *TokenContractFunctionsSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TokenContractFunctions.Contract.Allowance(&_TokenContractFunctions.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TokenContractFunctions.Contract.Allowance(&_TokenContractFunctions.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenContractFunctions *TokenContractFunctionsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenContractFunctions *TokenContractFunctionsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TokenContractFunctions.Contract.BalanceOf(&_TokenContractFunctions.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TokenContractFunctions.Contract.BalanceOf(&_TokenContractFunctions.CallOpts, _owner)
}

// IsTokenContract is a free data retrieval call binding the contract method 0x50372085.
//
// Solidity: function isTokenContract() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCaller) IsTokenContract(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "isTokenContract")
	return *ret0, err
}

// IsTokenContract is a free data retrieval call binding the contract method 0x50372085.
//
// Solidity: function isTokenContract() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) IsTokenContract() (bool, error) {
	return _TokenContractFunctions.Contract.IsTokenContract(&_TokenContractFunctions.CallOpts)
}

// IsTokenContract is a free data retrieval call binding the contract method 0x50372085.
//
// Solidity: function isTokenContract() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) IsTokenContract() (bool, error) {
	return _TokenContractFunctions.Contract.IsTokenContract(&_TokenContractFunctions.CallOpts)
}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsCaller) MigrationAgent(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "migrationAgent")
	return *ret0, err
}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsSession) MigrationAgent() (common.Address, error) {
	return _TokenContractFunctions.Contract.MigrationAgent(&_TokenContractFunctions.CallOpts)
}

// MigrationAgent is a free data retrieval call binding the contract method 0x8328dbcd.
//
// Solidity: function migrationAgent() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) MigrationAgent() (common.Address, error) {
	return _TokenContractFunctions.Contract.MigrationAgent(&_TokenContractFunctions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsSession) Owner() (common.Address, error) {
	return _TokenContractFunctions.Contract.Owner(&_TokenContractFunctions.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) Owner() (common.Address, error) {
	return _TokenContractFunctions.Contract.Owner(&_TokenContractFunctions.CallOpts)
}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsCaller) TotalMigrated(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "totalMigrated")
	return *ret0, err
}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsSession) TotalMigrated() (*big.Int, error) {
	return _TokenContractFunctions.Contract.TotalMigrated(&_TokenContractFunctions.CallOpts)
}

// TotalMigrated is a free data retrieval call binding the contract method 0x95a0f5eb.
//
// Solidity: function totalMigrated() constant returns(uint256)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) TotalMigrated() (*big.Int, error) {
	return _TokenContractFunctions.Contract.TotalMigrated(&_TokenContractFunctions.CallOpts)
}

// UpgradingEnabled is a free data retrieval call binding the contract method 0xc35506c0.
//
// Solidity: function upgradingEnabled() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCaller) UpgradingEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenContractFunctions.contract.Call(opts, out, "upgradingEnabled")
	return *ret0, err
}

// UpgradingEnabled is a free data retrieval call binding the contract method 0xc35506c0.
//
// Solidity: function upgradingEnabled() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) UpgradingEnabled() (bool, error) {
	return _TokenContractFunctions.Contract.UpgradingEnabled(&_TokenContractFunctions.CallOpts)
}

// UpgradingEnabled is a free data retrieval call binding the contract method 0xc35506c0.
//
// Solidity: function upgradingEnabled() constant returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsCallerSession) UpgradingEnabled() (bool, error) {
	return _TokenContractFunctions.Contract.UpgradingEnabled(&_TokenContractFunctions.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Approve(&_TokenContractFunctions.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Approve(&_TokenContractFunctions.TransactOpts, _spender, _value)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) Claim() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Claim(&_TokenContractFunctions.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) Claim() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Claim(&_TokenContractFunctions.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.DecreaseApproval(&_TokenContractFunctions.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.DecreaseApproval(&_TokenContractFunctions.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.IncreaseApproval(&_TokenContractFunctions.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.IncreaseApproval(&_TokenContractFunctions.TransactOpts, _spender, _addedValue)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(_value uint256) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) Migrate(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "migrate", _value)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(_value uint256) returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) Migrate(_value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Migrate(&_TokenContractFunctions.TransactOpts, _value)
}

// Migrate is a paid mutator transaction binding the contract method 0x454b0608.
//
// Solidity: function migrate(_value uint256) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) Migrate(_value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Migrate(&_TokenContractFunctions.TransactOpts, _value)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(_agent address) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) SetMigrationAgent(opts *bind.TransactOpts, _agent common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "setMigrationAgent", _agent)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(_agent address) returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) SetMigrationAgent(_agent common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.SetMigrationAgent(&_TokenContractFunctions.TransactOpts, _agent)
}

// SetMigrationAgent is a paid mutator transaction binding the contract method 0x75e2ff65.
//
// Solidity: function setMigrationAgent(_agent address) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) SetMigrationAgent(_agent common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.SetMigrationAgent(&_TokenContractFunctions.TransactOpts, _agent)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Transfer(&_TokenContractFunctions.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.Transfer(&_TokenContractFunctions.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TransferFrom(&_TokenContractFunctions.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TransferFrom(&_TokenContractFunctions.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TransferOwnership(&_TokenContractFunctions.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TransferOwnership(&_TokenContractFunctions.TransactOpts, newOwner)
}

// TweakState is a paid mutator transaction binding the contract method 0xb84dda70.
//
// Solidity: function tweakState() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) TweakState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "tweakState")
}

// TweakState is a paid mutator transaction binding the contract method 0xb84dda70.
//
// Solidity: function tweakState() returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) TweakState() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TweakState(&_TokenContractFunctions.TransactOpts)
}

// TweakState is a paid mutator transaction binding the contract method 0xb84dda70.
//
// Solidity: function tweakState() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) TweakState() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TweakState(&_TokenContractFunctions.TransactOpts)
}

// TweakUpgrading is a paid mutator transaction binding the contract method 0x5a6f7aef.
//
// Solidity: function tweakUpgrading() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactor) TweakUpgrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenContractFunctions.contract.Transact(opts, "tweakUpgrading")
}

// TweakUpgrading is a paid mutator transaction binding the contract method 0x5a6f7aef.
//
// Solidity: function tweakUpgrading() returns()
func (_TokenContractFunctions *TokenContractFunctionsSession) TweakUpgrading() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TweakUpgrading(&_TokenContractFunctions.TransactOpts)
}

// TweakUpgrading is a paid mutator transaction binding the contract method 0x5a6f7aef.
//
// Solidity: function tweakUpgrading() returns()
func (_TokenContractFunctions *TokenContractFunctionsTransactorSession) TweakUpgrading() (*types.Transaction, error) {
	return _TokenContractFunctions.Contract.TweakUpgrading(&_TokenContractFunctions.TransactOpts)
}
