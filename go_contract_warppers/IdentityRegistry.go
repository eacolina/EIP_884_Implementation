// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// IdentityRegistryABI is the input ABI used to generate the binding from.
const IdentityRegistryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"addIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"updatedAddress\",\"type\":\"address\"},{\"name\":\"newHash\",\"type\":\"string\"}],\"name\":\"updateIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"identityMap\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addressAdded\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"identityHash\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"authorizedBy\",\"type\":\"address\"}],\"name\":\"IdentityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"updatedAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousHash\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"newHash\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"authorizedBy\",\"type\":\"address\"}],\"name\":\"IdentityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressAdded\",\"type\":\"address\"}],\"name\":\"AddressAddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressRemoved\",\"type\":\"address\"}],\"name\":\"AddressRemovedFromWhitelist\",\"type\":\"event\"}]"

// IdentityRegistryBin is the compiled bytecode used for deploying new contracts.
const IdentityRegistryBin = `0x608060405234801561001057600080fd5b5060018054600160a060020a03191633179055610997806100326000396000f30060806040526004361061008d5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663286dd3f581146100925780633af32abf146100b55780635f319243146100ea5780637b9417c8146101515780638da5cb5b146101725780639b19251a146101a3578063a82567d5146101c4578063dfccf23f1461022b575b600080fd5b34801561009e57600080fd5b506100b3600160a060020a03600435166102c1565b005b3480156100c157600080fd5b506100d6600160a060020a0360043516610336565b604080519115158252519081900360200190f35b3480156100f657600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100b3958335600160a060020a03169536956044949193909101919081908401838280828437509497506103549650505050505050565b34801561015d57600080fd5b506100b3600160a060020a036004351661058c565b34801561017e57600080fd5b50610187610603565b60408051600160a060020a039092168252519081900360200190f35b3480156101af57600080fd5b506100d6600160a060020a0360043516610612565b3480156101d057600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100b3958335600160a060020a03169536956044949193909101919081908401838280828437509497506106279650505050505050565b34801561023757600080fd5b5061024c600160a060020a0360043516610837565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561028657818101518382015260200161026e565b50505050905090810190601f1680156102b35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600154600160a060020a031633146102d857600080fd5b6102e181610336565b15156102ec57600080fd5b600160a060020a038116600081815260208190526040808220805460ff191690555133917f5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb91a350565b600160a060020a031660009081526020819052604090205460ff1690565b600154606090600160a060020a0316331461036e57600080fd5b600160a060020a038316600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156104095780601f106103de57610100808354040283529160200191610409565b820191906000526020600020905b8154815290600101906020018083116103ec57829003601f168201915b50505050509050805160001415156104a757604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f54686973206964656e7469747920776173207265676973746572656420616c7260448201527f6561647900000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a038316600090815260026020908152604090912083516104d0928501906108d0565b506104da8361058c565b33600160a060020a031683600160a060020a03167f86cc2097bcb5dabdee3522860f3d6b0ab9c6d24adde6941695fb4f25baf8a881846040518080602001828103825283818151815260200191508051906020019080838360005b8381101561054d578181015183820152602001610535565b50505050905090810190601f16801561057a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a3505050565b600154600160a060020a031633146105a357600080fd5b6105ac81610336565b156105b657600080fd5b600160a060020a038116600081815260208190526040808220805460ff191660011790555133917f94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa691a350565b600154600160a060020a031681565b60006020819052908152604090205460ff1681565b600154606090600160a060020a0316331461064157600080fd5b600160a060020a038316600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156106dc5780601f106106b1576101008083540402835291602001916106dc565b820191906000526020600020905b8154815290600101906020018083116106bf57829003601f168201915b5050505050905080516000141515156106f457600080fd5b600160a060020a0383166000908152600260209081526040909120835161071d928501906108d0565b5033600160a060020a031683600160a060020a03167f40df790fd345471c77ba608aeb1a950504ec123523383018d37bd5ceb497e8058385604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561079657818101518382015260200161077e565b50505050905090810190601f1680156107c35780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156107f65781810151838201526020016107de565b50505050905090810190601f1680156108235780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3505050565b600260208181526000928352604092839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156108c85780601f1061089d576101008083540402835291602001916108c8565b820191906000526020600020905b8154815290600101906020018083116108ab57829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061091157805160ff191683800117855561093e565b8280016001018555821561093e579182015b8281111561093e578251825591602001919060010190610923565b5061094a92915061094e565b5090565b61096891905b8082111561094a5760008155600101610954565b905600a165627a7a723058200a8a75530a4ace0fc6c9dd30459ca4859389c55739c541cb8fa319d75fa686da0029`

// DeployIdentityRegistry deploys a new Ethereum contract, binding an instance of IdentityRegistry to it.
func DeployIdentityRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IdentityRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// IdentityRegistry is an auto generated Go binding around an Ethereum contract.
type IdentityRegistry struct {
	IdentityRegistryCaller     // Read-only binding to the contract
	IdentityRegistryTransactor // Write-only binding to the contract
	IdentityRegistryFilterer   // Log filterer for contract events
}

// IdentityRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistrySession struct {
	Contract     *IdentityRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryCallerSession struct {
	Contract *IdentityRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IdentityRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryTransactorSession struct {
	Contract     *IdentityRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IdentityRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryRaw struct {
	Contract *IdentityRegistry // Generic contract binding to access the raw methods on
}

// IdentityRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryCallerRaw struct {
	Contract *IdentityRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryTransactorRaw struct {
	Contract *IdentityRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistry creates a new instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistry(address common.Address, backend bind.ContractBackend) (*IdentityRegistry, error) {
	contract, err := bindIdentityRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistry{IdentityRegistryCaller: IdentityRegistryCaller{contract: contract}, IdentityRegistryTransactor: IdentityRegistryTransactor{contract: contract}, IdentityRegistryFilterer: IdentityRegistryFilterer{contract: contract}}, nil
}

// NewIdentityRegistryCaller creates a new read-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryCaller, error) {
	contract, err := bindIdentityRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryCaller{contract: contract}, nil
}

// NewIdentityRegistryTransactor creates a new write-only instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryTransactor, error) {
	contract, err := bindIdentityRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryTransactor{contract: contract}, nil
}

// NewIdentityRegistryFilterer creates a new log filterer instance of IdentityRegistry, bound to a specific deployed contract.
func NewIdentityRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFilterer, error) {
	contract, err := bindIdentityRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFilterer{contract: contract}, nil
}

// bindIdentityRegistry binds a generic wrapper to an already deployed contract.
func bindIdentityRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.IdentityRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.IdentityRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistry *IdentityRegistryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IdentityRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistry *IdentityRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.contract.Transact(opts, method, params...)
}

// IdentityMap is a free data retrieval call binding the contract method 0xdfccf23f.
//
// Solidity: function identityMap( address) constant returns(string)
func (_IdentityRegistry *IdentityRegistryCaller) IdentityMap(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "identityMap", arg0)
	return *ret0, err
}

// IdentityMap is a free data retrieval call binding the contract method 0xdfccf23f.
//
// Solidity: function identityMap( address) constant returns(string)
func (_IdentityRegistry *IdentityRegistrySession) IdentityMap(arg0 common.Address) (string, error) {
	return _IdentityRegistry.Contract.IdentityMap(&_IdentityRegistry.CallOpts, arg0)
}

// IdentityMap is a free data retrieval call binding the contract method 0xdfccf23f.
//
// Solidity: function identityMap( address) constant returns(string)
func (_IdentityRegistry *IdentityRegistryCallerSession) IdentityMap(arg0 common.Address) (string, error) {
	return _IdentityRegistry.Contract.IdentityMap(&_IdentityRegistry.CallOpts, arg0)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "isWhitelisted", _address)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) IsWhitelisted(_address common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsWhitelisted(&_IdentityRegistry.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _IdentityRegistry.Contract.IsWhitelisted(&_IdentityRegistry.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistrySession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IdentityRegistry *IdentityRegistryCallerSession) Owner() (common.Address, error) {
	return _IdentityRegistry.Contract.Owner(&_IdentityRegistry.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IdentityRegistry.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistrySession) Whitelist(arg0 common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Whitelist(&_IdentityRegistry.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_IdentityRegistry *IdentityRegistryCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _IdentityRegistry.Contract.Whitelist(&_IdentityRegistry.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "addAddressToWhitelist", _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistrySession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddAddressToWhitelist(&_IdentityRegistry.TransactOpts, _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddAddressToWhitelist(&_IdentityRegistry.TransactOpts, _address)
}

// AddIdentity is a paid mutator transaction binding the contract method 0x5f319243.
//
// Solidity: function addIdentity(_address address, hash string) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) AddIdentity(opts *bind.TransactOpts, _address common.Address, hash string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "addIdentity", _address, hash)
}

// AddIdentity is a paid mutator transaction binding the contract method 0x5f319243.
//
// Solidity: function addIdentity(_address address, hash string) returns()
func (_IdentityRegistry *IdentityRegistrySession) AddIdentity(_address common.Address, hash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddIdentity(&_IdentityRegistry.TransactOpts, _address, hash)
}

// AddIdentity is a paid mutator transaction binding the contract method 0x5f319243.
//
// Solidity: function addIdentity(_address address, hash string) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) AddIdentity(_address common.Address, hash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.AddIdentity(&_IdentityRegistry.TransactOpts, _address, hash)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "removeAddressFromWhitelist", _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistrySession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RemoveAddressFromWhitelist(&_IdentityRegistry.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.RemoveAddressFromWhitelist(&_IdentityRegistry.TransactOpts, _address)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0xa82567d5.
//
// Solidity: function updateIdentity(updatedAddress address, newHash string) returns()
func (_IdentityRegistry *IdentityRegistryTransactor) UpdateIdentity(opts *bind.TransactOpts, updatedAddress common.Address, newHash string) (*types.Transaction, error) {
	return _IdentityRegistry.contract.Transact(opts, "updateIdentity", updatedAddress, newHash)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0xa82567d5.
//
// Solidity: function updateIdentity(updatedAddress address, newHash string) returns()
func (_IdentityRegistry *IdentityRegistrySession) UpdateIdentity(updatedAddress common.Address, newHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.UpdateIdentity(&_IdentityRegistry.TransactOpts, updatedAddress, newHash)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0xa82567d5.
//
// Solidity: function updateIdentity(updatedAddress address, newHash string) returns()
func (_IdentityRegistry *IdentityRegistryTransactorSession) UpdateIdentity(updatedAddress common.Address, newHash string) (*types.Transaction, error) {
	return _IdentityRegistry.Contract.UpdateIdentity(&_IdentityRegistry.TransactOpts, updatedAddress, newHash)
}

// IdentityRegistryAddressAddedToWhitelistIterator is returned from FilterAddressAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddressAddedToWhitelist events raised by the IdentityRegistry contract.
type IdentityRegistryAddressAddedToWhitelistIterator struct {
	Event *IdentityRegistryAddressAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryAddressAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryAddressAddedToWhitelist)
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
		it.Event = new(IdentityRegistryAddressAddedToWhitelist)
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
func (it *IdentityRegistryAddressAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryAddressAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryAddressAddedToWhitelist represents a AddressAddedToWhitelist event raised by the IdentityRegistry contract.
type IdentityRegistryAddressAddedToWhitelist struct {
	AuthorizedBy common.Address
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToWhitelist is a free log retrieval operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterAddressAddedToWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressAdded []common.Address) (*IdentityRegistryAddressAddedToWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryAddressAddedToWhitelistIterator{contract: _IdentityRegistry.contract, event: "AddressAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToWhitelist is a free log subscription operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchAddressAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *IdentityRegistryAddressAddedToWhitelist, AuthorizedBy []common.Address, AddressAdded []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryAddressAddedToWhitelist)
				if err := _IdentityRegistry.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
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

// IdentityRegistryAddressRemovedFromWhitelistIterator is returned from FilterAddressRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromWhitelist events raised by the IdentityRegistry contract.
type IdentityRegistryAddressRemovedFromWhitelistIterator struct {
	Event *IdentityRegistryAddressRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryAddressRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryAddressRemovedFromWhitelist)
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
		it.Event = new(IdentityRegistryAddressRemovedFromWhitelist)
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
func (it *IdentityRegistryAddressRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryAddressRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryAddressRemovedFromWhitelist represents a AddressRemovedFromWhitelist event raised by the IdentityRegistry contract.
type IdentityRegistryAddressRemovedFromWhitelist struct {
	AuthorizedBy   common.Address
	AddressRemoved common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterAddressRemovedFromWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressRemoved []common.Address) (*IdentityRegistryAddressRemovedFromWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryAddressRemovedFromWhitelistIterator{contract: _IdentityRegistry.contract, event: "AddressRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromWhitelist is a free log subscription operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchAddressRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *IdentityRegistryAddressRemovedFromWhitelist, AuthorizedBy []common.Address, AddressRemoved []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryAddressRemovedFromWhitelist)
				if err := _IdentityRegistry.contract.UnpackLog(event, "AddressRemovedFromWhitelist", log); err != nil {
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

// IdentityRegistryIdentityAddedIterator is returned from FilterIdentityAdded and is used to iterate over the raw logs and unpacked data for IdentityAdded events raised by the IdentityRegistry contract.
type IdentityRegistryIdentityAddedIterator struct {
	Event *IdentityRegistryIdentityAdded // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryIdentityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryIdentityAdded)
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
		it.Event = new(IdentityRegistryIdentityAdded)
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
func (it *IdentityRegistryIdentityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryIdentityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryIdentityAdded represents a IdentityAdded event raised by the IdentityRegistry contract.
type IdentityRegistryIdentityAdded struct {
	AddressAdded common.Address
	IdentityHash string
	AuthorizedBy common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIdentityAdded is a free log retrieval operation binding the contract event 0x86cc2097bcb5dabdee3522860f3d6b0ab9c6d24adde6941695fb4f25baf8a881.
//
// Solidity: e IdentityAdded(addressAdded indexed address, identityHash string, authorizedBy indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterIdentityAdded(opts *bind.FilterOpts, addressAdded []common.Address, authorizedBy []common.Address) (*IdentityRegistryIdentityAddedIterator, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	var authorizedByRule []interface{}
	for _, authorizedByItem := range authorizedBy {
		authorizedByRule = append(authorizedByRule, authorizedByItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "IdentityAdded", addressAddedRule, authorizedByRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryIdentityAddedIterator{contract: _IdentityRegistry.contract, event: "IdentityAdded", logs: logs, sub: sub}, nil
}

// WatchIdentityAdded is a free log subscription operation binding the contract event 0x86cc2097bcb5dabdee3522860f3d6b0ab9c6d24adde6941695fb4f25baf8a881.
//
// Solidity: e IdentityAdded(addressAdded indexed address, identityHash string, authorizedBy indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchIdentityAdded(opts *bind.WatchOpts, sink chan<- *IdentityRegistryIdentityAdded, addressAdded []common.Address, authorizedBy []common.Address) (event.Subscription, error) {

	var addressAddedRule []interface{}
	for _, addressAddedItem := range addressAdded {
		addressAddedRule = append(addressAddedRule, addressAddedItem)
	}

	var authorizedByRule []interface{}
	for _, authorizedByItem := range authorizedBy {
		authorizedByRule = append(authorizedByRule, authorizedByItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "IdentityAdded", addressAddedRule, authorizedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryIdentityAdded)
				if err := _IdentityRegistry.contract.UnpackLog(event, "IdentityAdded", log); err != nil {
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

// IdentityRegistryIdentityUpdatedIterator is returned from FilterIdentityUpdated and is used to iterate over the raw logs and unpacked data for IdentityUpdated events raised by the IdentityRegistry contract.
type IdentityRegistryIdentityUpdatedIterator struct {
	Event *IdentityRegistryIdentityUpdated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryIdentityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryIdentityUpdated)
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
		it.Event = new(IdentityRegistryIdentityUpdated)
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
func (it *IdentityRegistryIdentityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryIdentityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryIdentityUpdated represents a IdentityUpdated event raised by the IdentityRegistry contract.
type IdentityRegistryIdentityUpdated struct {
	UpdatedAddress common.Address
	PreviousHash   string
	NewHash        string
	AuthorizedBy   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterIdentityUpdated is a free log retrieval operation binding the contract event 0x40df790fd345471c77ba608aeb1a950504ec123523383018d37bd5ceb497e805.
//
// Solidity: e IdentityUpdated(updatedAddress indexed address, previousHash string, newHash string, authorizedBy indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) FilterIdentityUpdated(opts *bind.FilterOpts, updatedAddress []common.Address, authorizedBy []common.Address) (*IdentityRegistryIdentityUpdatedIterator, error) {

	var updatedAddressRule []interface{}
	for _, updatedAddressItem := range updatedAddress {
		updatedAddressRule = append(updatedAddressRule, updatedAddressItem)
	}

	var authorizedByRule []interface{}
	for _, authorizedByItem := range authorizedBy {
		authorizedByRule = append(authorizedByRule, authorizedByItem)
	}

	logs, sub, err := _IdentityRegistry.contract.FilterLogs(opts, "IdentityUpdated", updatedAddressRule, authorizedByRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryIdentityUpdatedIterator{contract: _IdentityRegistry.contract, event: "IdentityUpdated", logs: logs, sub: sub}, nil
}

// WatchIdentityUpdated is a free log subscription operation binding the contract event 0x40df790fd345471c77ba608aeb1a950504ec123523383018d37bd5ceb497e805.
//
// Solidity: e IdentityUpdated(updatedAddress indexed address, previousHash string, newHash string, authorizedBy indexed address)
func (_IdentityRegistry *IdentityRegistryFilterer) WatchIdentityUpdated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryIdentityUpdated, updatedAddress []common.Address, authorizedBy []common.Address) (event.Subscription, error) {

	var updatedAddressRule []interface{}
	for _, updatedAddressItem := range updatedAddress {
		updatedAddressRule = append(updatedAddressRule, updatedAddressItem)
	}

	var authorizedByRule []interface{}
	for _, authorizedByItem := range authorizedBy {
		authorizedByRule = append(authorizedByRule, authorizedByItem)
	}

	logs, sub, err := _IdentityRegistry.contract.WatchLogs(opts, "IdentityUpdated", updatedAddressRule, authorizedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryIdentityUpdated)
				if err := _IdentityRegistry.contract.UnpackLog(event, "IdentityUpdated", log); err != nil {
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

// WhitelistableABI is the input ABI used to generate the binding from.
const WhitelistableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressAdded\",\"type\":\"address\"}],\"name\":\"AddressAddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressRemoved\",\"type\":\"address\"}],\"name\":\"AddressRemovedFromWhitelist\",\"type\":\"event\"}]"

// WhitelistableBin is the compiled bytecode used for deploying new contracts.
const WhitelistableBin = `0x608060405234801561001057600080fd5b5060018054600160a060020a03191633179055610296806100326000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663286dd3f581146100715780633af32abf146100945780637b9417c8146100c95780638da5cb5b146100ea5780639b19251a1461011b575b600080fd5b34801561007d57600080fd5b50610092600160a060020a036004351661013c565b005b3480156100a057600080fd5b506100b5600160a060020a03600435166101b1565b604080519115158252519081900360200190f35b3480156100d557600080fd5b50610092600160a060020a03600435166101cf565b3480156100f657600080fd5b506100ff610246565b60408051600160a060020a039092168252519081900360200190f35b34801561012757600080fd5b506100b5600160a060020a0360043516610255565b600154600160a060020a0316331461015357600080fd5b61015c816101b1565b151561016757600080fd5b600160a060020a038116600081815260208190526040808220805460ff191690555133917f5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb91a350565b600160a060020a031660009081526020819052604090205460ff1690565b600154600160a060020a031633146101e657600080fd5b6101ef816101b1565b156101f957600080fd5b600160a060020a038116600081815260208190526040808220805460ff191660011790555133917f94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa691a350565b600154600160a060020a031681565b60006020819052908152604090205460ff16815600a165627a7a72305820fd8d207f77fc7f3b2daadcbb66da80780efcca87ed99617881cab26ac6e694d30029`

// DeployWhitelistable deploys a new Ethereum contract, binding an instance of Whitelistable to it.
func DeployWhitelistable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Whitelistable, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WhitelistableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Whitelistable{WhitelistableCaller: WhitelistableCaller{contract: contract}, WhitelistableTransactor: WhitelistableTransactor{contract: contract}, WhitelistableFilterer: WhitelistableFilterer{contract: contract}}, nil
}

// Whitelistable is an auto generated Go binding around an Ethereum contract.
type Whitelistable struct {
	WhitelistableCaller     // Read-only binding to the contract
	WhitelistableTransactor // Write-only binding to the contract
	WhitelistableFilterer   // Log filterer for contract events
}

// WhitelistableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistableSession struct {
	Contract     *Whitelistable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhitelistableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistableCallerSession struct {
	Contract *WhitelistableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WhitelistableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistableTransactorSession struct {
	Contract     *WhitelistableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WhitelistableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistableRaw struct {
	Contract *Whitelistable // Generic contract binding to access the raw methods on
}

// WhitelistableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistableCallerRaw struct {
	Contract *WhitelistableCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistableTransactorRaw struct {
	Contract *WhitelistableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistable creates a new instance of Whitelistable, bound to a specific deployed contract.
func NewWhitelistable(address common.Address, backend bind.ContractBackend) (*Whitelistable, error) {
	contract, err := bindWhitelistable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Whitelistable{WhitelistableCaller: WhitelistableCaller{contract: contract}, WhitelistableTransactor: WhitelistableTransactor{contract: contract}, WhitelistableFilterer: WhitelistableFilterer{contract: contract}}, nil
}

// NewWhitelistableCaller creates a new read-only instance of Whitelistable, bound to a specific deployed contract.
func NewWhitelistableCaller(address common.Address, caller bind.ContractCaller) (*WhitelistableCaller, error) {
	contract, err := bindWhitelistable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistableCaller{contract: contract}, nil
}

// NewWhitelistableTransactor creates a new write-only instance of Whitelistable, bound to a specific deployed contract.
func NewWhitelistableTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistableTransactor, error) {
	contract, err := bindWhitelistable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistableTransactor{contract: contract}, nil
}

// NewWhitelistableFilterer creates a new log filterer instance of Whitelistable, bound to a specific deployed contract.
func NewWhitelistableFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistableFilterer, error) {
	contract, err := bindWhitelistable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistableFilterer{contract: contract}, nil
}

// bindWhitelistable binds a generic wrapper to an already deployed contract.
func bindWhitelistable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelistable *WhitelistableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelistable.Contract.WhitelistableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelistable *WhitelistableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelistable.Contract.WhitelistableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelistable *WhitelistableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelistable.Contract.WhitelistableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Whitelistable *WhitelistableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Whitelistable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Whitelistable *WhitelistableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Whitelistable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Whitelistable *WhitelistableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Whitelistable.Contract.contract.Transact(opts, method, params...)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Whitelistable *WhitelistableCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelistable.contract.Call(opts, out, "isWhitelisted", _address)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Whitelistable *WhitelistableSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Whitelistable.Contract.IsWhitelisted(&_Whitelistable.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_Whitelistable *WhitelistableCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _Whitelistable.Contract.IsWhitelisted(&_Whitelistable.CallOpts, _address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelistable *WhitelistableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Whitelistable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelistable *WhitelistableSession) Owner() (common.Address, error) {
	return _Whitelistable.Contract.Owner(&_Whitelistable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Whitelistable *WhitelistableCallerSession) Owner() (common.Address, error) {
	return _Whitelistable.Contract.Owner(&_Whitelistable.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelistable *WhitelistableCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Whitelistable.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelistable *WhitelistableSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelistable.Contract.Whitelist(&_Whitelistable.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_Whitelistable *WhitelistableCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Whitelistable.Contract.Whitelist(&_Whitelistable.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Whitelistable.contract.Transact(opts, "addAddressToWhitelist", _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Whitelistable.Contract.AddAddressToWhitelist(&_Whitelistable.TransactOpts, _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableTransactorSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Whitelistable.Contract.AddAddressToWhitelist(&_Whitelistable.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Whitelistable.contract.Transact(opts, "removeAddressFromWhitelist", _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Whitelistable.Contract.RemoveAddressFromWhitelist(&_Whitelistable.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_Whitelistable *WhitelistableTransactorSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Whitelistable.Contract.RemoveAddressFromWhitelist(&_Whitelistable.TransactOpts, _address)
}

// WhitelistableAddressAddedToWhitelistIterator is returned from FilterAddressAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddressAddedToWhitelist events raised by the Whitelistable contract.
type WhitelistableAddressAddedToWhitelistIterator struct {
	Event *WhitelistableAddressAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *WhitelistableAddressAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistableAddressAddedToWhitelist)
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
		it.Event = new(WhitelistableAddressAddedToWhitelist)
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
func (it *WhitelistableAddressAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistableAddressAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistableAddressAddedToWhitelist represents a AddressAddedToWhitelist event raised by the Whitelistable contract.
type WhitelistableAddressAddedToWhitelist struct {
	AuthorizedBy common.Address
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToWhitelist is a free log retrieval operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_Whitelistable *WhitelistableFilterer) FilterAddressAddedToWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressAdded []common.Address) (*WhitelistableAddressAddedToWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _Whitelistable.contract.FilterLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistableAddressAddedToWhitelistIterator{contract: _Whitelistable.contract, event: "AddressAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToWhitelist is a free log subscription operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_Whitelistable *WhitelistableFilterer) WatchAddressAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *WhitelistableAddressAddedToWhitelist, AuthorizedBy []common.Address, AddressAdded []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _Whitelistable.contract.WatchLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistableAddressAddedToWhitelist)
				if err := _Whitelistable.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
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

// WhitelistableAddressRemovedFromWhitelistIterator is returned from FilterAddressRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromWhitelist events raised by the Whitelistable contract.
type WhitelistableAddressRemovedFromWhitelistIterator struct {
	Event *WhitelistableAddressRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *WhitelistableAddressRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistableAddressRemovedFromWhitelist)
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
		it.Event = new(WhitelistableAddressRemovedFromWhitelist)
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
func (it *WhitelistableAddressRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistableAddressRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistableAddressRemovedFromWhitelist represents a AddressRemovedFromWhitelist event raised by the Whitelistable contract.
type WhitelistableAddressRemovedFromWhitelist struct {
	AuthorizedBy   common.Address
	AddressRemoved common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_Whitelistable *WhitelistableFilterer) FilterAddressRemovedFromWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressRemoved []common.Address) (*WhitelistableAddressRemovedFromWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _Whitelistable.contract.FilterLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistableAddressRemovedFromWhitelistIterator{contract: _Whitelistable.contract, event: "AddressRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromWhitelist is a free log subscription operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_Whitelistable *WhitelistableFilterer) WatchAddressRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *WhitelistableAddressRemovedFromWhitelist, AuthorizedBy []common.Address, AddressRemoved []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _Whitelistable.contract.WatchLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistableAddressRemovedFromWhitelist)
				if err := _Whitelistable.contract.UnpackLog(event, "AddressRemovedFromWhitelist", log); err != nil {
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
