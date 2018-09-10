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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405234801561001057600080fd5b5061027a806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461005b57806370a0823114610082578063a9059cbb146100b0575b600080fd5b34801561006757600080fd5b506100706100f5565b60408051918252519081900360200190f35b34801561008e57600080fd5b5061007073ffffffffffffffffffffffffffffffffffffffff600435166100fb565b3480156100bc57600080fd5b506100e173ffffffffffffffffffffffffffffffffffffffff60043516602435610123565b604080519115158252519081900360200190f35b60015490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b3360009081526020819052604081205482111561013f57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8316151561016157600080fd5b33600090815260208190526040902054610181908363ffffffff61022916565b336000908152602081905260408082209290925573ffffffffffffffffffffffffffffffffffffffff8516815220546101c0908363ffffffff61023b16565b73ffffffffffffffffffffffffffffffffffffffff8416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b60008282111561023557fe5b50900390565b8181018281101561024857fe5b929150505600a165627a7a72305820f46a24d5bec5b00b1da132a7021f96b6d861e2a1f5d01e8350826924d90fef470029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582046eafaacb3acedcf2325b80f0ed11eaa43e9a4c3c03c2635cf650089415114860029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StockTokenABI is the input ABI used to generate the binding from.
const StockTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"isPrivateCompany\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWhitelisted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"platformWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"byLawsHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"togglePrivateCompany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownersCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokenOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_supply\",\"type\":\"uint256\"},{\"name\":\"hash\",\"type\":\"string\"},{\"name\":\"_registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"authorizedBy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newStatus\",\"type\":\"bool\"}],\"name\":\"ChangedCompanyStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressAdded\",\"type\":\"address\"}],\"name\":\"AddressAddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"AuthorizedBy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"AddressRemoved\",\"type\":\"address\"}],\"name\":\"AddressRemovedFromWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StockTokenBin is the compiled bytecode used for deploying new contracts.
const StockTokenBin = `0x60806040526009805460ff1916600117905534801561001d57600080fd5b50604051610db7380380610db78339810160409081528151602080840151928401516060850151608086015160038054600160a060020a03191633179055938601805190969586019592949190920192909160009161008191600491890190610151565b508451610095906005906020880190610151565b50600184905582516100ae906006906020860190610151565b50506007805433600160a060020a0319918216811790925560008281526020818152604080832097909755600c80548416600160a060020a039690961695909517909455600b8054600181810183557f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db9918201805486169055825490810190925581018054909316909317909155600a90925292902091909155506101ec915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061019257805160ff19168380011785556101bf565b828001600101855582156101bf579182015b828111156101bf5782518255916020019190600101906101a4565b506101cb9291506101cf565b5090565b6101e991905b808211156101cb57600081556001016101d5565b90565b610bbc806101fb6000396000f3006080604052600436106100fb5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630401a24a811461010057806306fdde031461012957806318160ddd146101b3578063286dd3f5146101da578063313ce567146101fd5780633af32abf1461021257806370a08231146102335780637b9417c8146102545780637e3e9ede1461027557806380d01bcd146102a65780638da5cb5b146102bb57806395d89b41146102d05780639b19251a146102e5578063a9059cbb14610306578063ac6af1c51461032a578063b94885461461033f578063d1b02ba714610354578063f8a14f46146103b9575b600080fd5b34801561010c57600080fd5b506101156103d1565b604080519115158252519081900360200190f35b34801561013557600080fd5b5061013e6103da565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610178578181015183820152602001610160565b50505050905090810190601f1680156101a55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101bf57600080fd5b506101c8610468565b60408051918252519081900360200190f35b3480156101e657600080fd5b506101fb600160a060020a036004351661046f565b005b34801561020957600080fd5b506101c86104e4565b34801561021e57600080fd5b50610115600160a060020a03600435166104ea565b34801561023f57600080fd5b506101c8600160a060020a0360043516610508565b34801561026057600080fd5b506101fb600160a060020a0360043516610523565b34801561028157600080fd5b5061028a61059a565b60408051600160a060020a039092168252519081900360200190f35b3480156102b257600080fd5b5061013e6105a9565b3480156102c757600080fd5b5061028a610604565b3480156102dc57600080fd5b5061013e610613565b3480156102f157600080fd5b50610115600160a060020a036004351661066e565b34801561031257600080fd5b50610115600160a060020a0360043516602435610683565b34801561033657600080fd5b506101fb610886565b34801561034b57600080fd5b506101c86108f1565b34801561036057600080fd5b506103696108fb565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103a557818101518382015260200161038d565b505050509050019250505060405180910390f35b3480156103c557600080fd5b5061028a60043561095d565b60095460ff1681565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104605780601f1061043557610100808354040283529160200191610460565b820191906000526020600020905b81548152906001019060200180831161044357829003601f168201915b505050505081565b6001545b90565b600754600160a060020a0316331461048657600080fd5b61048f816104ea565b151561049a57600080fd5b600160a060020a038116600081815260026020526040808220805460ff191690555133917f5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb91a350565b60085481565b600160a060020a031660009081526002602052604090205460ff1690565b600160a060020a031660009081526020819052604090205490565b600754600160a060020a0316331461053a57600080fd5b610543816104ea565b1561054d57600080fd5b600160a060020a038116600081815260026020526040808220805460ff191660011790555133917f94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa691a350565b600c54600160a060020a031681565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104605780601f1061043557610100808354040283529160200191610460565b600754600160a060020a031681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104605780601f1061043557610100808354040283529160200191610460565b60026020526000908152604090205460ff1681565b600c54604080517f3af32abf000000000000000000000000000000000000000000000000000000008152600160a060020a038086166004830152915160009384938793911691633af32abf9160248082019260209290919082900301818887803b1580156106f057600080fd5b505af1158015610704573d6000803e3d6000fd5b505050506040513d602081101561071a57600080fd5b5051151561072757600080fd5b60095460ff16156107ce5761073b816104ea565b15156107ce57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f41646472657373206e6f7420696e2070726976617465207368617265686f6c6460448201527f6572732077686974656c69737400000000000000000000000000000000000000606482015290519081900360840190fd5b600160a060020a0385166000908152600a6020526040812054909250151561085b57600b8054600181019091557f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db98101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091556000908152600a6020526040902081905591505b6108658585610985565b5061086f33610508565b151561087e5761087e33610a64565b505092915050565b600754600160a060020a0316331461089d57600080fd5b6009805460ff19811660ff91821615179182905560408051338152929091161515602083015280517fb52734fc46bdfb46ff91a021440a182abee932adcac1be40cde3e0838d851b549281900390910190a1565b600b546000190190565b6060600b80548060200260200160405190810160405280929190818152602001828054801561095357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610935575b5050505050905090565b600b80548290811061096b57fe5b600091825260209091200154600160a060020a0316905081565b336000908152602081905260408120548211156109a157600080fd5b600160a060020a03831615156109b657600080fd5b336000908152602081905260409020546109d6908363ffffffff610aa116565b3360009081526020819052604080822092909255600160a060020a03851681522054610a08908363ffffffff610ab316565b600160a060020a038416600081815260208181526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b600160a060020a0381166000908152600a6020526040902054610a8681610ac6565b50600160a060020a03166000908152600a6020526040812055565b600082821115610aad57fe5b50900390565b81810182811015610ac057fe5b92915050565b600b8054600091906000198101908110610adc57fe5b600091825260209091200154600b8054600160a060020a039092169250829184908110610b0557fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055600b805490610b4e906000198301610b53565b505050565b815481835581811115610b4e57600083815260209020610b4e91810190830161046c91905b80821115610b8c5760008155600101610b78565b50905600a165627a7a723058204e110903d337eb4a2032e3275a4fd02335339f6f8850db6264d5da89646db1f50029`

// DeployStockToken deploys a new Ethereum contract, binding an instance of StockToken to it.
func DeployStockToken(auth *bind.TransactOpts, backend bind.ContractBackend, _symbol string, _name string, _supply *big.Int, hash string, _registry common.Address) (common.Address, *types.Transaction, *StockToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StockTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StockTokenBin), backend, _symbol, _name, _supply, hash, _registry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StockToken{StockTokenCaller: StockTokenCaller{contract: contract}, StockTokenTransactor: StockTokenTransactor{contract: contract}, StockTokenFilterer: StockTokenFilterer{contract: contract}}, nil
}

// StockToken is an auto generated Go binding around an Ethereum contract.
type StockToken struct {
	StockTokenCaller     // Read-only binding to the contract
	StockTokenTransactor // Write-only binding to the contract
	StockTokenFilterer   // Log filterer for contract events
}

// StockTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StockTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StockTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StockTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StockTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StockTokenSession struct {
	Contract     *StockToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StockTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StockTokenCallerSession struct {
	Contract *StockTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StockTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StockTokenTransactorSession struct {
	Contract     *StockTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StockTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StockTokenRaw struct {
	Contract *StockToken // Generic contract binding to access the raw methods on
}

// StockTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StockTokenCallerRaw struct {
	Contract *StockTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StockTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StockTokenTransactorRaw struct {
	Contract *StockTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStockToken creates a new instance of StockToken, bound to a specific deployed contract.
func NewStockToken(address common.Address, backend bind.ContractBackend) (*StockToken, error) {
	contract, err := bindStockToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StockToken{StockTokenCaller: StockTokenCaller{contract: contract}, StockTokenTransactor: StockTokenTransactor{contract: contract}, StockTokenFilterer: StockTokenFilterer{contract: contract}}, nil
}

// NewStockTokenCaller creates a new read-only instance of StockToken, bound to a specific deployed contract.
func NewStockTokenCaller(address common.Address, caller bind.ContractCaller) (*StockTokenCaller, error) {
	contract, err := bindStockToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StockTokenCaller{contract: contract}, nil
}

// NewStockTokenTransactor creates a new write-only instance of StockToken, bound to a specific deployed contract.
func NewStockTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StockTokenTransactor, error) {
	contract, err := bindStockToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StockTokenTransactor{contract: contract}, nil
}

// NewStockTokenFilterer creates a new log filterer instance of StockToken, bound to a specific deployed contract.
func NewStockTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StockTokenFilterer, error) {
	contract, err := bindStockToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StockTokenFilterer{contract: contract}, nil
}

// bindStockToken binds a generic wrapper to an already deployed contract.
func bindStockToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StockTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StockToken *StockTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StockToken.Contract.StockTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StockToken *StockTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StockToken.Contract.StockTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StockToken *StockTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StockToken.Contract.StockTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StockToken *StockTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StockToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StockToken *StockTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StockToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StockToken *StockTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StockToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StockToken *StockTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StockToken *StockTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StockToken.Contract.BalanceOf(&_StockToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_StockToken *StockTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StockToken.Contract.BalanceOf(&_StockToken.CallOpts, _owner)
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_StockToken *StockTokenCaller) ByLawsHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "byLawsHash")
	return *ret0, err
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_StockToken *StockTokenSession) ByLawsHash() (string, error) {
	return _StockToken.Contract.ByLawsHash(&_StockToken.CallOpts)
}

// ByLawsHash is a free data retrieval call binding the contract method 0x80d01bcd.
//
// Solidity: function byLawsHash() constant returns(string)
func (_StockToken *StockTokenCallerSession) ByLawsHash() (string, error) {
	return _StockToken.Contract.ByLawsHash(&_StockToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StockToken *StockTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StockToken *StockTokenSession) Decimals() (*big.Int, error) {
	return _StockToken.Contract.Decimals(&_StockToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StockToken *StockTokenCallerSession) Decimals() (*big.Int, error) {
	return _StockToken.Contract.Decimals(&_StockToken.CallOpts)
}

// GetTokenOwners is a free data retrieval call binding the contract method 0xd1b02ba7.
//
// Solidity: function getTokenOwners() constant returns(address[])
func (_StockToken *StockTokenCaller) GetTokenOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "getTokenOwners")
	return *ret0, err
}

// GetTokenOwners is a free data retrieval call binding the contract method 0xd1b02ba7.
//
// Solidity: function getTokenOwners() constant returns(address[])
func (_StockToken *StockTokenSession) GetTokenOwners() ([]common.Address, error) {
	return _StockToken.Contract.GetTokenOwners(&_StockToken.CallOpts)
}

// GetTokenOwners is a free data retrieval call binding the contract method 0xd1b02ba7.
//
// Solidity: function getTokenOwners() constant returns(address[])
func (_StockToken *StockTokenCallerSession) GetTokenOwners() ([]common.Address, error) {
	return _StockToken.Contract.GetTokenOwners(&_StockToken.CallOpts)
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_StockToken *StockTokenCaller) IsPrivateCompany(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "isPrivateCompany")
	return *ret0, err
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_StockToken *StockTokenSession) IsPrivateCompany() (bool, error) {
	return _StockToken.Contract.IsPrivateCompany(&_StockToken.CallOpts)
}

// IsPrivateCompany is a free data retrieval call binding the contract method 0x0401a24a.
//
// Solidity: function isPrivateCompany() constant returns(bool)
func (_StockToken *StockTokenCallerSession) IsPrivateCompany() (bool, error) {
	return _StockToken.Contract.IsPrivateCompany(&_StockToken.CallOpts)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_StockToken *StockTokenCaller) IsWhitelisted(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "isWhitelisted", _address)
	return *ret0, err
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_StockToken *StockTokenSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _StockToken.Contract.IsWhitelisted(&_StockToken.CallOpts, _address)
}

// IsWhitelisted is a free data retrieval call binding the contract method 0x3af32abf.
//
// Solidity: function isWhitelisted(_address address) constant returns(bool)
func (_StockToken *StockTokenCallerSession) IsWhitelisted(_address common.Address) (bool, error) {
	return _StockToken.Contract.IsWhitelisted(&_StockToken.CallOpts, _address)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StockToken *StockTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StockToken *StockTokenSession) Name() (string, error) {
	return _StockToken.Contract.Name(&_StockToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StockToken *StockTokenCallerSession) Name() (string, error) {
	return _StockToken.Contract.Name(&_StockToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StockToken *StockTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StockToken *StockTokenSession) Owner() (common.Address, error) {
	return _StockToken.Contract.Owner(&_StockToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StockToken *StockTokenCallerSession) Owner() (common.Address, error) {
	return _StockToken.Contract.Owner(&_StockToken.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint256)
func (_StockToken *StockTokenCaller) OwnersCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "ownersCount")
	return *ret0, err
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint256)
func (_StockToken *StockTokenSession) OwnersCount() (*big.Int, error) {
	return _StockToken.Contract.OwnersCount(&_StockToken.CallOpts)
}

// OwnersCount is a free data retrieval call binding the contract method 0xb9488546.
//
// Solidity: function ownersCount() constant returns(uint256)
func (_StockToken *StockTokenCallerSession) OwnersCount() (*big.Int, error) {
	return _StockToken.Contract.OwnersCount(&_StockToken.CallOpts)
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_StockToken *StockTokenCaller) PlatformWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "platformWhitelist")
	return *ret0, err
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_StockToken *StockTokenSession) PlatformWhitelist() (common.Address, error) {
	return _StockToken.Contract.PlatformWhitelist(&_StockToken.CallOpts)
}

// PlatformWhitelist is a free data retrieval call binding the contract method 0x7e3e9ede.
//
// Solidity: function platformWhitelist() constant returns(address)
func (_StockToken *StockTokenCallerSession) PlatformWhitelist() (common.Address, error) {
	return _StockToken.Contract.PlatformWhitelist(&_StockToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StockToken *StockTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StockToken *StockTokenSession) Symbol() (string, error) {
	return _StockToken.Contract.Symbol(&_StockToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StockToken *StockTokenCallerSession) Symbol() (string, error) {
	return _StockToken.Contract.Symbol(&_StockToken.CallOpts)
}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners( uint256) constant returns(address)
func (_StockToken *StockTokenCaller) TokenOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "tokenOwners", arg0)
	return *ret0, err
}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners( uint256) constant returns(address)
func (_StockToken *StockTokenSession) TokenOwners(arg0 *big.Int) (common.Address, error) {
	return _StockToken.Contract.TokenOwners(&_StockToken.CallOpts, arg0)
}

// TokenOwners is a free data retrieval call binding the contract method 0xf8a14f46.
//
// Solidity: function tokenOwners( uint256) constant returns(address)
func (_StockToken *StockTokenCallerSession) TokenOwners(arg0 *big.Int) (common.Address, error) {
	return _StockToken.Contract.TokenOwners(&_StockToken.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StockToken *StockTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StockToken *StockTokenSession) TotalSupply() (*big.Int, error) {
	return _StockToken.Contract.TotalSupply(&_StockToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StockToken *StockTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StockToken.Contract.TotalSupply(&_StockToken.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_StockToken *StockTokenCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StockToken.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_StockToken *StockTokenSession) Whitelist(arg0 common.Address) (bool, error) {
	return _StockToken.Contract.Whitelist(&_StockToken.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_StockToken *StockTokenCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _StockToken.Contract.Whitelist(&_StockToken.CallOpts, arg0)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_StockToken *StockTokenTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StockToken.contract.Transact(opts, "addAddressToWhitelist", _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_StockToken *StockTokenSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _StockToken.Contract.AddAddressToWhitelist(&_StockToken.TransactOpts, _address)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(_address address) returns()
func (_StockToken *StockTokenTransactorSession) AddAddressToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _StockToken.Contract.AddAddressToWhitelist(&_StockToken.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_StockToken *StockTokenTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StockToken.contract.Transact(opts, "removeAddressFromWhitelist", _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_StockToken *StockTokenSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _StockToken.Contract.RemoveAddressFromWhitelist(&_StockToken.TransactOpts, _address)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(_address address) returns()
func (_StockToken *StockTokenTransactorSession) RemoveAddressFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _StockToken.Contract.RemoveAddressFromWhitelist(&_StockToken.TransactOpts, _address)
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_StockToken *StockTokenTransactor) TogglePrivateCompany(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StockToken.contract.Transact(opts, "togglePrivateCompany")
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_StockToken *StockTokenSession) TogglePrivateCompany() (*types.Transaction, error) {
	return _StockToken.Contract.TogglePrivateCompany(&_StockToken.TransactOpts)
}

// TogglePrivateCompany is a paid mutator transaction binding the contract method 0xac6af1c5.
//
// Solidity: function togglePrivateCompany() returns()
func (_StockToken *StockTokenTransactorSession) TogglePrivateCompany() (*types.Transaction, error) {
	return _StockToken.Contract.TogglePrivateCompany(&_StockToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StockToken *StockTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StockToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StockToken *StockTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StockToken.Contract.Transfer(&_StockToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StockToken *StockTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StockToken.Contract.Transfer(&_StockToken.TransactOpts, _to, _value)
}

// StockTokenAddressAddedToWhitelistIterator is returned from FilterAddressAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddressAddedToWhitelist events raised by the StockToken contract.
type StockTokenAddressAddedToWhitelistIterator struct {
	Event *StockTokenAddressAddedToWhitelist // Event containing the contract specifics and raw log

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
func (it *StockTokenAddressAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StockTokenAddressAddedToWhitelist)
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
		it.Event = new(StockTokenAddressAddedToWhitelist)
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
func (it *StockTokenAddressAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StockTokenAddressAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StockTokenAddressAddedToWhitelist represents a AddressAddedToWhitelist event raised by the StockToken contract.
type StockTokenAddressAddedToWhitelist struct {
	AuthorizedBy common.Address
	AddressAdded common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToWhitelist is a free log retrieval operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_StockToken *StockTokenFilterer) FilterAddressAddedToWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressAdded []common.Address) (*StockTokenAddressAddedToWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _StockToken.contract.FilterLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return &StockTokenAddressAddedToWhitelistIterator{contract: _StockToken.contract, event: "AddressAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToWhitelist is a free log subscription operation binding the contract event 0x94d300f0803cc0af9991569051574a3dafd674f3b8d03f81c5ac017d5293afa6.
//
// Solidity: e AddressAddedToWhitelist(AuthorizedBy indexed address, AddressAdded indexed address)
func (_StockToken *StockTokenFilterer) WatchAddressAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *StockTokenAddressAddedToWhitelist, AuthorizedBy []common.Address, AddressAdded []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressAddedRule []interface{}
	for _, AddressAddedItem := range AddressAdded {
		AddressAddedRule = append(AddressAddedRule, AddressAddedItem)
	}

	logs, sub, err := _StockToken.contract.WatchLogs(opts, "AddressAddedToWhitelist", AuthorizedByRule, AddressAddedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StockTokenAddressAddedToWhitelist)
				if err := _StockToken.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
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

// StockTokenAddressRemovedFromWhitelistIterator is returned from FilterAddressRemovedFromWhitelist and is used to iterate over the raw logs and unpacked data for AddressRemovedFromWhitelist events raised by the StockToken contract.
type StockTokenAddressRemovedFromWhitelistIterator struct {
	Event *StockTokenAddressRemovedFromWhitelist // Event containing the contract specifics and raw log

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
func (it *StockTokenAddressRemovedFromWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StockTokenAddressRemovedFromWhitelist)
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
		it.Event = new(StockTokenAddressRemovedFromWhitelist)
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
func (it *StockTokenAddressRemovedFromWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StockTokenAddressRemovedFromWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StockTokenAddressRemovedFromWhitelist represents a AddressRemovedFromWhitelist event raised by the StockToken contract.
type StockTokenAddressRemovedFromWhitelist struct {
	AuthorizedBy   common.Address
	AddressRemoved common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAddressRemovedFromWhitelist is a free log retrieval operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_StockToken *StockTokenFilterer) FilterAddressRemovedFromWhitelist(opts *bind.FilterOpts, AuthorizedBy []common.Address, AddressRemoved []common.Address) (*StockTokenAddressRemovedFromWhitelistIterator, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _StockToken.contract.FilterLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return &StockTokenAddressRemovedFromWhitelistIterator{contract: _StockToken.contract, event: "AddressRemovedFromWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressRemovedFromWhitelist is a free log subscription operation binding the contract event 0x5a76ef3d09a09d8a99c3d9ddddfb7770d906de670d79b5119a8c2cba903814bb.
//
// Solidity: e AddressRemovedFromWhitelist(AuthorizedBy indexed address, AddressRemoved indexed address)
func (_StockToken *StockTokenFilterer) WatchAddressRemovedFromWhitelist(opts *bind.WatchOpts, sink chan<- *StockTokenAddressRemovedFromWhitelist, AuthorizedBy []common.Address, AddressRemoved []common.Address) (event.Subscription, error) {

	var AuthorizedByRule []interface{}
	for _, AuthorizedByItem := range AuthorizedBy {
		AuthorizedByRule = append(AuthorizedByRule, AuthorizedByItem)
	}
	var AddressRemovedRule []interface{}
	for _, AddressRemovedItem := range AddressRemoved {
		AddressRemovedRule = append(AddressRemovedRule, AddressRemovedItem)
	}

	logs, sub, err := _StockToken.contract.WatchLogs(opts, "AddressRemovedFromWhitelist", AuthorizedByRule, AddressRemovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StockTokenAddressRemovedFromWhitelist)
				if err := _StockToken.contract.UnpackLog(event, "AddressRemovedFromWhitelist", log); err != nil {
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

// StockTokenChangedCompanyStatusIterator is returned from FilterChangedCompanyStatus and is used to iterate over the raw logs and unpacked data for ChangedCompanyStatus events raised by the StockToken contract.
type StockTokenChangedCompanyStatusIterator struct {
	Event *StockTokenChangedCompanyStatus // Event containing the contract specifics and raw log

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
func (it *StockTokenChangedCompanyStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StockTokenChangedCompanyStatus)
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
		it.Event = new(StockTokenChangedCompanyStatus)
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
func (it *StockTokenChangedCompanyStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StockTokenChangedCompanyStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StockTokenChangedCompanyStatus represents a ChangedCompanyStatus event raised by the StockToken contract.
type StockTokenChangedCompanyStatus struct {
	AuthorizedBy common.Address
	NewStatus    bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterChangedCompanyStatus is a free log retrieval operation binding the contract event 0xb52734fc46bdfb46ff91a021440a182abee932adcac1be40cde3e0838d851b54.
//
// Solidity: e ChangedCompanyStatus(authorizedBy address, newStatus bool)
func (_StockToken *StockTokenFilterer) FilterChangedCompanyStatus(opts *bind.FilterOpts) (*StockTokenChangedCompanyStatusIterator, error) {

	logs, sub, err := _StockToken.contract.FilterLogs(opts, "ChangedCompanyStatus")
	if err != nil {
		return nil, err
	}
	return &StockTokenChangedCompanyStatusIterator{contract: _StockToken.contract, event: "ChangedCompanyStatus", logs: logs, sub: sub}, nil
}

// WatchChangedCompanyStatus is a free log subscription operation binding the contract event 0xb52734fc46bdfb46ff91a021440a182abee932adcac1be40cde3e0838d851b54.
//
// Solidity: e ChangedCompanyStatus(authorizedBy address, newStatus bool)
func (_StockToken *StockTokenFilterer) WatchChangedCompanyStatus(opts *bind.WatchOpts, sink chan<- *StockTokenChangedCompanyStatus) (event.Subscription, error) {

	logs, sub, err := _StockToken.contract.WatchLogs(opts, "ChangedCompanyStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StockTokenChangedCompanyStatus)
				if err := _StockToken.contract.UnpackLog(event, "ChangedCompanyStatus", log); err != nil {
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

// StockTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StockToken contract.
type StockTokenTransferIterator struct {
	Event *StockTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StockTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StockTokenTransfer)
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
		it.Event = new(StockTokenTransfer)
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
func (it *StockTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StockTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StockTokenTransfer represents a Transfer event raised by the StockToken contract.
type StockTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StockToken *StockTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StockTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StockToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StockTokenTransferIterator{contract: _StockToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StockToken *StockTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StockTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StockToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StockTokenTransfer)
				if err := _StockToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
