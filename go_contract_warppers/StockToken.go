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

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isPrivateCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platformWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"byLawsHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_supply\",\"type\":\"uint256\"},{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"authorizedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"ChangedCompanyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressAdded\",\"type\":\"address\"}],\"name\":\"AddressAddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressRemoved\",\"type\":\"address\"}],\"name\":\"AddressRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"togglePrivateCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, _owner)
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_Contracts *ContractsCaller) ByLawsHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "byLawsHash")
	return *ret0, err
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_Contracts *ContractsSession) ByLawsHash() (string, error) {
	return _Contracts.Contract.ByLawsHash(&_Contracts.CallOpts)
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_Contracts *ContractsCallerSession) ByLawsHash() (string, error) {
	return _Contracts.Contract.ByLawsHash(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Contracts *ContractsSession) Decimals() (*big.Int, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Contracts *ContractsCallerSession) Decimals() (*big.Int, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_Contracts *ContractsCaller) IsPrivateCompany(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "isPrivateCompany")
	return *ret0, err
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_Contracts *ContractsSession) IsPrivateCompany() (bool, error) {
	return _Contracts.Contract.IsPrivateCompany(&_Contracts.CallOpts)
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_Contracts *ContractsCallerSession) IsPrivateCompany() (bool, error) {
	return _Contracts.Contract.IsPrivateCompany(&_Contracts.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Contracts *ContractsCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "isWhitelisted", _address)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Contracts *ContractsSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Contracts.Contract.IsWhitelisted(&_Contracts.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Contracts *ContractsCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Contracts.Contract.IsWhitelisted(&_Contracts.CallOpts, _address)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_Contracts *ContractsCaller) PlatformWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "platformWhitelist")
	return *ret0, err
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_Contracts *ContractsSession) PlatformWhitelist() (common.Address, error) {
	return _Contracts.Contract.PlatformWhitelist(&_Contracts.CallOpts)
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_Contracts *ContractsCallerSession) PlatformWhitelist() (common.Address, error) {
	return _Contracts.Contract.PlatformWhitelist(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Contracts *ContractsCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Contracts *ContractsSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Whitelist(&_Contracts.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Contracts *ContractsCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Contracts.Contract.Whitelist(&_Contracts.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Contracts *ContractsTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addAddressToWhitelist", _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Contracts *ContractsSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAddressToWhitelist(&_Contracts.TransactOpts, _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Contracts *ContractsTransactorSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAddressToWhitelist(&_Contracts.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Contracts *ContractsTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeAddressFromWhitelist", _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Contracts *ContractsSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAddressFromWhitelist(&_Contracts.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Contracts *ContractsTransactorSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAddressFromWhitelist(&_Contracts.TransactOpts, _address)
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_Contracts *ContractsTransactor) TogglePrivateCompany(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "togglePrivateCompany")
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_Contracts *ContractsSession) TogglePrivateCompany() (*types.Transaction, error) {
	return _Contracts.Contract.TogglePrivateCompany(&_Contracts.TransactOpts)
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_Contracts *ContractsTransactorSession) TogglePrivateCompany() (*types.Transaction, error) {
	return _Contracts.Contract.TogglePrivateCompany(&_Contracts.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Contracts *ContractsSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, _to, _value)
}

// ContractsAddressAddedToWhitelistIterator is returned from FilterAddressAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddressAddedToWhitelist events raised by the Contracts contract.
type ContractsAddressAddedToWhitelistIterator struct {
	Event *ContractsAddressAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractsAddressAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressAddedToWhitelist)
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
		it.Event = new(ContractsAddressAddedToWhitelist)
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
func (it *ContractsAddressAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressAddedToWhitelist represents a AddressAddedToWhitelist event raised by the Contracts contract.
type ContractsAddressAddedToWhitelist struct {
	AuthorizedBy common.Address
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToWhitelist is a free log retrieval operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_Contracts *ContractsFilterer) FilterAddressAddedToWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressAdded []common.Address) (*ContractsAddressAddedToWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressAddedToWhitelistIterator{contract: _Contracts.contract, event: "AddressAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToWhitelist is a free log subscription operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_Contracts *ContractsFilterer) WatchAddressAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *ContractsAddressAddedToWhitelist, AuthorizedBy []common.Address, AddressAdded []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressAddedToWhitelist)
				if err := _Contracts.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
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

// ContractsAddressRemovedFromWhitelistIterator is returned from FilterAddressRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromWhitelist events raised by the Contracts contract.
type ContractsAddressRemovedFromWhitelistIterator struct {
	Event *ContractsAddressRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *ContractsAddressRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAddressRemovedFromWhitelist)
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
		it.Event = new(ContractsAddressRemovedFromWhitelist)
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
func (it *ContractsAddressRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAddressRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAddressRemovedFromWhitelist represents a AddressRemovedFromWhitelist event raised by the Contracts contract.
type ContractsAddressRemovedFromWhitelist struct {
	AuthorizedBy   common.Address
	AddressRemoved common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_Contracts *ContractsFilterer) FilterAddressRemovedFromWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressRemoved []common.Address) (*ContractsAddressRemovedFromWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return &ContractsAddressRemovedFromWhitelistIterator{contract: _Contracts.contract, event: "AddressRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromWhitelist is a free log subscription operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_Contracts *ContractsFilterer) WatchAddressRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *ContractsAddressRemovedFromWhitelist, AuthorizedBy []common.Address, AddressRemoved []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAddressRemovedFromWhitelist)
				if err := _Contracts.contract.UnpackLog(event, "AddressRemovedFromWhitelist", log); err != nil {
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

// ContractsChangedCompanyStatusIterator is returned from FilterChangedCompanyStatus and is used to iterate over the raw logs and unpacked data for ChangedCompanyStatus events raised by the Contracts contract.
type ContractsChangedCompanyStatusIterator struct {
	Event *ContractsChangedCompanyStatus // Event containing the contract specifics and raw log

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
func (it *ContractsChangedCompanyStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsChangedCompanyStatus)
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
		it.Event = new(ContractsChangedCompanyStatus)
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
func (it *ContractsChangedCompanyStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsChangedCompanyStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsChangedCompanyStatus represents a ChangedCompanyStatus event raised by the Contracts contract.
type ContractsChangedCompanyStatus struct {
	AuthorizedBy common.Address
	NewStatus    bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChangedCompanyStatus is a free log retrieval operation binding the contract event 0xb52734fc46bdfb46ff91a021440a182abee932adcac1be40cde3e0838d851b54.
//
// Solidity: e ChangedCompanyStatus(authorizedBy address, newStatus bool)
func (_Contracts *ContractsFilterer) FilterChangedCompanyStatus(opts *bind.FilterOpts) (*ContractsChangedCompanyStatusIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ChangedCompanyStatus")
	if err != nil {
		return nil, err
	}
	return &ContractsChangedCompanyStatusIterator{contract: _Contracts.contract, event: "ChangedCompanyStatus", logs: logs, sub: sub}, nil
}

// WatchChangedCompanyStatus is a free log subscription operation binding the contract event 0xb52734fc46bdfb46ff91a021440a182abee932adcac1be40cde3e0838d851b54.
//
// Solidity: e ChangedCompanyStatus(authorizedBy address, newStatus bool)
func (_Contracts *ContractsFilterer) WatchChangedCompanyStatus(opts *bind.WatchOpts, sink chan<- *ContractsChangedCompanyStatus) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ChangedCompanyStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsChangedCompanyStatus)
				if err := _Contracts.contract.UnpackLog(event, "ChangedCompanyStatus", log); err != nil {
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

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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
