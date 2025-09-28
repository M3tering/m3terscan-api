// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package m3tering

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
	_ = abi.ConvertType
)

// M3teringMetaData contains all meta data concerning the M3tering contract.
var M3teringMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"anchorBlock\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainLength\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"accountBlob\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"nonceBlob\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"NewState\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SP1_PROGRAM_VKEY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"bytes6\",\"name\":\"\",\"type\":\"bytes6\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"anchorBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"accountBlob\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"nonceBlob\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"commitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"io\",\"type\":\"uint256\"}],\"name\":\"latestStateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"bytes6\",\"name\":\"\",\"type\":\"bytes6\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"programVKey_\",\"type\":\"bytes32\"}],\"name\":\"setProgramVKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"bytes6\",\"name\":\"\",\"type\":\"bytes6\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"at\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"stateAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// M3teringABI is the input ABI used to generate the binding from.
// Deprecated: Use M3teringMetaData.ABI instead.
var M3teringABI = M3teringMetaData.ABI

// M3tering is an auto generated Go binding around an Ethereum contract.
type M3tering struct {
	M3teringCaller     // Read-only binding to the contract
	M3teringTransactor // Write-only binding to the contract
	M3teringFilterer   // Log filterer for contract events
}

// M3teringCaller is an auto generated read-only Go binding around an Ethereum contract.
type M3teringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// M3teringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type M3teringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// M3teringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type M3teringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// M3teringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type M3teringSession struct {
	Contract     *M3tering         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// M3teringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type M3teringCallerSession struct {
	Contract *M3teringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// M3teringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type M3teringTransactorSession struct {
	Contract     *M3teringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// M3teringRaw is an auto generated low-level Go binding around an Ethereum contract.
type M3teringRaw struct {
	Contract *M3tering // Generic contract binding to access the raw methods on
}

// M3teringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type M3teringCallerRaw struct {
	Contract *M3teringCaller // Generic read-only contract binding to access the raw methods on
}

// M3teringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type M3teringTransactorRaw struct {
	Contract *M3teringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewM3tering creates a new instance of M3tering, bound to a specific deployed contract.
func NewM3tering(address common.Address, backend bind.ContractBackend) (*M3tering, error) {
	contract, err := bindM3tering(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &M3tering{M3teringCaller: M3teringCaller{contract: contract}, M3teringTransactor: M3teringTransactor{contract: contract}, M3teringFilterer: M3teringFilterer{contract: contract}}, nil
}

// NewM3teringCaller creates a new read-only instance of M3tering, bound to a specific deployed contract.
func NewM3teringCaller(address common.Address, caller bind.ContractCaller) (*M3teringCaller, error) {
	contract, err := bindM3tering(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &M3teringCaller{contract: contract}, nil
}

// NewM3teringTransactor creates a new write-only instance of M3tering, bound to a specific deployed contract.
func NewM3teringTransactor(address common.Address, transactor bind.ContractTransactor) (*M3teringTransactor, error) {
	contract, err := bindM3tering(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &M3teringTransactor{contract: contract}, nil
}

// NewM3teringFilterer creates a new log filterer instance of M3tering, bound to a specific deployed contract.
func NewM3teringFilterer(address common.Address, filterer bind.ContractFilterer) (*M3teringFilterer, error) {
	contract, err := bindM3tering(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &M3teringFilterer{contract: contract}, nil
}

// bindM3tering binds a generic wrapper to an already deployed contract.
func bindM3tering(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := M3teringMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_M3tering *M3teringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _M3tering.Contract.M3teringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_M3tering *M3teringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _M3tering.Contract.M3teringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_M3tering *M3teringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _M3tering.Contract.M3teringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_M3tering *M3teringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _M3tering.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_M3tering *M3teringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _M3tering.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_M3tering *M3teringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _M3tering.Contract.contract.Transact(opts, method, params...)
}

// SP1PROGRAMVKEY is a free data retrieval call binding the contract method 0xf1825792.
//
// Solidity: function SP1_PROGRAM_VKEY() view returns(bytes32)
func (_M3tering *M3teringCaller) SP1PROGRAMVKEY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "SP1_PROGRAM_VKEY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SP1PROGRAMVKEY is a free data retrieval call binding the contract method 0xf1825792.
//
// Solidity: function SP1_PROGRAM_VKEY() view returns(bytes32)
func (_M3tering *M3teringSession) SP1PROGRAMVKEY() ([32]byte, error) {
	return _M3tering.Contract.SP1PROGRAMVKEY(&_M3tering.CallOpts)
}

// SP1PROGRAMVKEY is a free data retrieval call binding the contract method 0xf1825792.
//
// Solidity: function SP1_PROGRAM_VKEY() view returns(bytes32)
func (_M3tering *M3teringCallerSession) SP1PROGRAMVKEY() ([32]byte, error) {
	return _M3tering.Contract.SP1PROGRAMVKEY(&_M3tering.CallOpts)
}

// Account is a free data retrieval call binding the contract method 0x2dd7c658.
//
// Solidity: function account(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCaller) Account(opts *bind.CallOpts, tokenId *big.Int) ([6]byte, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "account", tokenId)

	if err != nil {
		return *new([6]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([6]byte)).(*[6]byte)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x2dd7c658.
//
// Solidity: function account(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringSession) Account(tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.Account(&_M3tering.CallOpts, tokenId)
}

// Account is a free data retrieval call binding the contract method 0x2dd7c658.
//
// Solidity: function account(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCallerSession) Account(tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.Account(&_M3tering.CallOpts, tokenId)
}

// AnchorBlock is a free data retrieval call binding the contract method 0x26d95d03.
//
// Solidity: function anchorBlock() view returns(bytes32)
func (_M3tering *M3teringCaller) AnchorBlock(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "anchorBlock")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AnchorBlock is a free data retrieval call binding the contract method 0x26d95d03.
//
// Solidity: function anchorBlock() view returns(bytes32)
func (_M3tering *M3teringSession) AnchorBlock() ([32]byte, error) {
	return _M3tering.Contract.AnchorBlock(&_M3tering.CallOpts)
}

// AnchorBlock is a free data retrieval call binding the contract method 0x26d95d03.
//
// Solidity: function anchorBlock() view returns(bytes32)
func (_M3tering *M3teringCallerSession) AnchorBlock() ([32]byte, error) {
	return _M3tering.Contract.AnchorBlock(&_M3tering.CallOpts)
}

// ChainLength is a free data retrieval call binding the contract method 0x219ee89c.
//
// Solidity: function chainLength() view returns(uint256)
func (_M3tering *M3teringCaller) ChainLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "chainLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainLength is a free data retrieval call binding the contract method 0x219ee89c.
//
// Solidity: function chainLength() view returns(uint256)
func (_M3tering *M3teringSession) ChainLength() (*big.Int, error) {
	return _M3tering.Contract.ChainLength(&_M3tering.CallOpts)
}

// ChainLength is a free data retrieval call binding the contract method 0x219ee89c.
//
// Solidity: function chainLength() view returns(uint256)
func (_M3tering *M3teringCallerSession) ChainLength() (*big.Int, error) {
	return _M3tering.Contract.ChainLength(&_M3tering.CallOpts)
}

// LatestStateAddress is a free data retrieval call binding the contract method 0x7e033b88.
//
// Solidity: function latestStateAddress(uint256 io) view returns(address)
func (_M3tering *M3teringCaller) LatestStateAddress(opts *bind.CallOpts, io *big.Int) (common.Address, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "latestStateAddress", io)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LatestStateAddress is a free data retrieval call binding the contract method 0x7e033b88.
//
// Solidity: function latestStateAddress(uint256 io) view returns(address)
func (_M3tering *M3teringSession) LatestStateAddress(io *big.Int) (common.Address, error) {
	return _M3tering.Contract.LatestStateAddress(&_M3tering.CallOpts, io)
}

// LatestStateAddress is a free data retrieval call binding the contract method 0x7e033b88.
//
// Solidity: function latestStateAddress(uint256 io) view returns(address)
func (_M3tering *M3teringCallerSession) LatestStateAddress(io *big.Int) (common.Address, error) {
	return _M3tering.Contract.LatestStateAddress(&_M3tering.CallOpts, io)
}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCaller) Nonce(opts *bind.CallOpts, tokenId *big.Int) ([6]byte, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "nonce", tokenId)

	if err != nil {
		return *new([6]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([6]byte)).(*[6]byte)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringSession) Nonce(tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.Nonce(&_M3tering.CallOpts, tokenId)
}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCallerSession) Nonce(tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.Nonce(&_M3tering.CallOpts, tokenId)
}

// State is a free data retrieval call binding the contract method 0x29c50ca7.
//
// Solidity: function state(uint256 at, bytes4 selector, uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCaller) State(opts *bind.CallOpts, at *big.Int, selector [4]byte, tokenId *big.Int) ([6]byte, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "state", at, selector, tokenId)

	if err != nil {
		return *new([6]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([6]byte)).(*[6]byte)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x29c50ca7.
//
// Solidity: function state(uint256 at, bytes4 selector, uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringSession) State(at *big.Int, selector [4]byte, tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.State(&_M3tering.CallOpts, at, selector, tokenId)
}

// State is a free data retrieval call binding the contract method 0x29c50ca7.
//
// Solidity: function state(uint256 at, bytes4 selector, uint256 tokenId) view returns(bytes6)
func (_M3tering *M3teringCallerSession) State(at *big.Int, selector [4]byte, tokenId *big.Int) ([6]byte, error) {
	return _M3tering.Contract.State(&_M3tering.CallOpts, at, selector, tokenId)
}

// StateAddress is a free data retrieval call binding the contract method 0x09a73d27.
//
// Solidity: function stateAddress(uint256 at, bytes4 selector) view returns(address)
func (_M3tering *M3teringCaller) StateAddress(opts *bind.CallOpts, at *big.Int, selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _M3tering.contract.Call(opts, &out, "stateAddress", at, selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateAddress is a free data retrieval call binding the contract method 0x09a73d27.
//
// Solidity: function stateAddress(uint256 at, bytes4 selector) view returns(address)
func (_M3tering *M3teringSession) StateAddress(at *big.Int, selector [4]byte) (common.Address, error) {
	return _M3tering.Contract.StateAddress(&_M3tering.CallOpts, at, selector)
}

// StateAddress is a free data retrieval call binding the contract method 0x09a73d27.
//
// Solidity: function stateAddress(uint256 at, bytes4 selector) view returns(address)
func (_M3tering *M3teringCallerSession) StateAddress(at *big.Int, selector [4]byte) (common.Address, error) {
	return _M3tering.Contract.StateAddress(&_M3tering.CallOpts, at, selector)
}

// CommitState is a paid mutator transaction binding the contract method 0x12d8317d.
//
// Solidity: function commitState(bytes accountBlob, bytes nonceBlob, bytes proof) returns()
func (_M3tering *M3teringTransactor) CommitState(opts *bind.TransactOpts, accountBlob []byte, nonceBlob []byte, proof []byte) (*types.Transaction, error) {
	return _M3tering.contract.Transact(opts, "commitState", accountBlob, nonceBlob, proof)
}

// CommitState is a paid mutator transaction binding the contract method 0x12d8317d.
//
// Solidity: function commitState(bytes accountBlob, bytes nonceBlob, bytes proof) returns()
func (_M3tering *M3teringSession) CommitState(accountBlob []byte, nonceBlob []byte, proof []byte) (*types.Transaction, error) {
	return _M3tering.Contract.CommitState(&_M3tering.TransactOpts, accountBlob, nonceBlob, proof)
}

// CommitState is a paid mutator transaction binding the contract method 0x12d8317d.
//
// Solidity: function commitState(bytes accountBlob, bytes nonceBlob, bytes proof) returns()
func (_M3tering *M3teringTransactorSession) CommitState(accountBlob []byte, nonceBlob []byte, proof []byte) (*types.Transaction, error) {
	return _M3tering.Contract.CommitState(&_M3tering.TransactOpts, accountBlob, nonceBlob, proof)
}

// SetProgramVKey is a paid mutator transaction binding the contract method 0x61ed2713.
//
// Solidity: function setProgramVKey(bytes32 programVKey_) returns()
func (_M3tering *M3teringTransactor) SetProgramVKey(opts *bind.TransactOpts, programVKey_ [32]byte) (*types.Transaction, error) {
	return _M3tering.contract.Transact(opts, "setProgramVKey", programVKey_)
}

// SetProgramVKey is a paid mutator transaction binding the contract method 0x61ed2713.
//
// Solidity: function setProgramVKey(bytes32 programVKey_) returns()
func (_M3tering *M3teringSession) SetProgramVKey(programVKey_ [32]byte) (*types.Transaction, error) {
	return _M3tering.Contract.SetProgramVKey(&_M3tering.TransactOpts, programVKey_)
}

// SetProgramVKey is a paid mutator transaction binding the contract method 0x61ed2713.
//
// Solidity: function setProgramVKey(bytes32 programVKey_) returns()
func (_M3tering *M3teringTransactorSession) SetProgramVKey(programVKey_ [32]byte) (*types.Transaction, error) {
	return _M3tering.Contract.SetProgramVKey(&_M3tering.TransactOpts, programVKey_)
}

// M3teringNewStateIterator is returned from FilterNewState and is used to iterate over the raw logs and unpacked data for NewState events raised by the M3tering contract.
type M3teringNewStateIterator struct {
	Event *M3teringNewState // Event containing the contract specifics and raw log

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
func (it *M3teringNewStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(M3teringNewState)
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
		it.Event = new(M3teringNewState)
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
func (it *M3teringNewStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *M3teringNewStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// M3teringNewState represents a NewState event raised by the M3tering contract.
type M3teringNewState struct {
	From        common.Address
	AnchorBlock [32]byte
	ChainLength *big.Int
	AccountBlob []byte
	NonceBlob   []byte
	Proof       []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewState is a free log retrieval operation binding the contract event 0xbf5e49846b17f125ea39652d829be1dc16cd038b68c85c4e722f139d97064091.
//
// Solidity: event NewState(address indexed from, bytes32 indexed anchorBlock, uint256 indexed chainLength, bytes accountBlob, bytes nonceBlob, bytes proof)
func (_M3tering *M3teringFilterer) FilterNewState(opts *bind.FilterOpts, from []common.Address, anchorBlock [][32]byte, chainLength []*big.Int) (*M3teringNewStateIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var anchorBlockRule []interface{}
	for _, anchorBlockItem := range anchorBlock {
		anchorBlockRule = append(anchorBlockRule, anchorBlockItem)
	}
	var chainLengthRule []interface{}
	for _, chainLengthItem := range chainLength {
		chainLengthRule = append(chainLengthRule, chainLengthItem)
	}

	logs, sub, err := _M3tering.contract.FilterLogs(opts, "NewState", fromRule, anchorBlockRule, chainLengthRule)
	if err != nil {
		return nil, err
	}
	return &M3teringNewStateIterator{contract: _M3tering.contract, event: "NewState", logs: logs, sub: sub}, nil
}

// WatchNewState is a free log subscription operation binding the contract event 0xbf5e49846b17f125ea39652d829be1dc16cd038b68c85c4e722f139d97064091.
//
// Solidity: event NewState(address indexed from, bytes32 indexed anchorBlock, uint256 indexed chainLength, bytes accountBlob, bytes nonceBlob, bytes proof)
func (_M3tering *M3teringFilterer) WatchNewState(opts *bind.WatchOpts, sink chan<- *M3teringNewState, from []common.Address, anchorBlock [][32]byte, chainLength []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var anchorBlockRule []interface{}
	for _, anchorBlockItem := range anchorBlock {
		anchorBlockRule = append(anchorBlockRule, anchorBlockItem)
	}
	var chainLengthRule []interface{}
	for _, chainLengthItem := range chainLength {
		chainLengthRule = append(chainLengthRule, chainLengthItem)
	}

	logs, sub, err := _M3tering.contract.WatchLogs(opts, "NewState", fromRule, anchorBlockRule, chainLengthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(M3teringNewState)
				if err := _M3tering.contract.UnpackLog(event, "NewState", log); err != nil {
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

// ParseNewState is a log parse operation binding the contract event 0xbf5e49846b17f125ea39652d829be1dc16cd038b68c85c4e722f139d97064091.
//
// Solidity: event NewState(address indexed from, bytes32 indexed anchorBlock, uint256 indexed chainLength, bytes accountBlob, bytes nonceBlob, bytes proof)
func (_M3tering *M3teringFilterer) ParseNewState(log types.Log) (*M3teringNewState, error) {
	event := new(M3teringNewState)
	if err := _M3tering.contract.UnpackLog(event, "NewState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
