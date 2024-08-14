// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// WeweMetaData contains all meta data concerning the Wewe contract.
var WeweMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ticker_\",\"type\":\"string\"}],\"name\":\"setNameAndTicker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWhitelistContract\",\"type\":\"address\"}],\"name\":\"setWhitelistContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WeweABI is the input ABI used to generate the binding from.
// Deprecated: Use WeweMetaData.ABI instead.
var WeweABI = WeweMetaData.ABI

// Wewe is an auto generated Go binding around an Ethereum contract.
type Wewe struct {
	WeweCaller     // Read-only binding to the contract
	WeweTransactor // Write-only binding to the contract
	WeweFilterer   // Log filterer for contract events
}

// WeweCaller is an auto generated read-only Go binding around an Ethereum contract.
type WeweCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeweTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WeweTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeweFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WeweFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeweSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WeweSession struct {
	Contract     *Wewe             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WeweCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WeweCallerSession struct {
	Contract *WeweCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WeweTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WeweTransactorSession struct {
	Contract     *WeweTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WeweRaw is an auto generated low-level Go binding around an Ethereum contract.
type WeweRaw struct {
	Contract *Wewe // Generic contract binding to access the raw methods on
}

// WeweCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WeweCallerRaw struct {
	Contract *WeweCaller // Generic read-only contract binding to access the raw methods on
}

// WeweTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WeweTransactorRaw struct {
	Contract *WeweTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWewe creates a new instance of Wewe, bound to a specific deployed contract.
func NewWewe(address common.Address, backend bind.ContractBackend) (*Wewe, error) {
	contract, err := bindWewe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wewe{WeweCaller: WeweCaller{contract: contract}, WeweTransactor: WeweTransactor{contract: contract}, WeweFilterer: WeweFilterer{contract: contract}}, nil
}

// NewWeweCaller creates a new read-only instance of Wewe, bound to a specific deployed contract.
func NewWeweCaller(address common.Address, caller bind.ContractCaller) (*WeweCaller, error) {
	contract, err := bindWewe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WeweCaller{contract: contract}, nil
}

// NewWeweTransactor creates a new write-only instance of Wewe, bound to a specific deployed contract.
func NewWeweTransactor(address common.Address, transactor bind.ContractTransactor) (*WeweTransactor, error) {
	contract, err := bindWewe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WeweTransactor{contract: contract}, nil
}

// NewWeweFilterer creates a new log filterer instance of Wewe, bound to a specific deployed contract.
func NewWeweFilterer(address common.Address, filterer bind.ContractFilterer) (*WeweFilterer, error) {
	contract, err := bindWewe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WeweFilterer{contract: contract}, nil
}

// bindWewe binds a generic wrapper to an already deployed contract.
func bindWewe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WeweABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wewe *WeweRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wewe.Contract.WeweCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wewe *WeweRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wewe.Contract.WeweTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wewe *WeweRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wewe.Contract.WeweTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wewe *WeweCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wewe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wewe *WeweTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wewe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wewe *WeweTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wewe.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Wewe *WeweCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Wewe *WeweSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Wewe.Contract.Allowance(&_Wewe.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Wewe *WeweCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Wewe.Contract.Allowance(&_Wewe.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Wewe *WeweCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Wewe *WeweSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Wewe.Contract.BalanceOf(&_Wewe.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Wewe *WeweCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Wewe.Contract.BalanceOf(&_Wewe.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wewe *WeweCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wewe *WeweSession) Decimals() (uint8, error) {
	return _Wewe.Contract.Decimals(&_Wewe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wewe *WeweCallerSession) Decimals() (uint8, error) {
	return _Wewe.Contract.Decimals(&_Wewe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wewe *WeweCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wewe *WeweSession) Name() (string, error) {
	return _Wewe.Contract.Name(&_Wewe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wewe *WeweCallerSession) Name() (string, error) {
	return _Wewe.Contract.Name(&_Wewe.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wewe *WeweCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wewe *WeweSession) Owner() (common.Address, error) {
	return _Wewe.Contract.Owner(&_Wewe.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wewe *WeweCallerSession) Owner() (common.Address, error) {
	return _Wewe.Contract.Owner(&_Wewe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wewe *WeweCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wewe *WeweSession) Symbol() (string, error) {
	return _Wewe.Contract.Symbol(&_Wewe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wewe *WeweCallerSession) Symbol() (string, error) {
	return _Wewe.Contract.Symbol(&_Wewe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wewe *WeweCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wewe *WeweSession) TotalSupply() (*big.Int, error) {
	return _Wewe.Contract.TotalSupply(&_Wewe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wewe *WeweCallerSession) TotalSupply() (*big.Int, error) {
	return _Wewe.Contract.TotalSupply(&_Wewe.CallOpts)
}

// WhitelistContract is a free data retrieval call binding the contract method 0x84900b04.
//
// Solidity: function whitelistContract() view returns(address)
func (_Wewe *WeweCaller) WhitelistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wewe.contract.Call(opts, &out, "whitelistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistContract is a free data retrieval call binding the contract method 0x84900b04.
//
// Solidity: function whitelistContract() view returns(address)
func (_Wewe *WeweSession) WhitelistContract() (common.Address, error) {
	return _Wewe.Contract.WhitelistContract(&_Wewe.CallOpts)
}

// WhitelistContract is a free data retrieval call binding the contract method 0x84900b04.
//
// Solidity: function whitelistContract() view returns(address)
func (_Wewe *WeweCallerSession) WhitelistContract() (common.Address, error) {
	return _Wewe.Contract.WhitelistContract(&_Wewe.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Wewe *WeweTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Wewe *WeweSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Approve(&_Wewe.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Wewe *WeweTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Approve(&_Wewe.TransactOpts, spender, amount)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes extraData) returns(bool)
func (_Wewe *WeweTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "approveAndCall", spender, amount, extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes extraData) returns(bool)
func (_Wewe *WeweSession) ApproveAndCall(spender common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Wewe.Contract.ApproveAndCall(&_Wewe.TransactOpts, spender, amount, extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 amount, bytes extraData) returns(bool)
func (_Wewe *WeweTransactorSession) ApproveAndCall(spender common.Address, amount *big.Int, extraData []byte) (*types.Transaction, error) {
	return _Wewe.Contract.ApproveAndCall(&_Wewe.TransactOpts, spender, amount, extraData)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Wewe *WeweTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Wewe *WeweSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Burn(&_Wewe.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Wewe *WeweTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Burn(&_Wewe.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Wewe *WeweTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Wewe *WeweSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.BurnFrom(&_Wewe.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Wewe *WeweTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.BurnFrom(&_Wewe.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Wewe *WeweTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Wewe *WeweSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.DecreaseAllowance(&_Wewe.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Wewe *WeweTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.DecreaseAllowance(&_Wewe.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Wewe *WeweTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Wewe *WeweSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.IncreaseAllowance(&_Wewe.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Wewe *WeweTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.IncreaseAllowance(&_Wewe.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Wewe *WeweTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Wewe *WeweSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Mint(&_Wewe.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns()
func (_Wewe *WeweTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Mint(&_Wewe.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wewe *WeweTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wewe *WeweSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wewe.Contract.RenounceOwnership(&_Wewe.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Wewe *WeweTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Wewe.Contract.RenounceOwnership(&_Wewe.TransactOpts)
}

// SetNameAndTicker is a paid mutator transaction binding the contract method 0xc388c167.
//
// Solidity: function setNameAndTicker(string name_, string ticker_) returns()
func (_Wewe *WeweTransactor) SetNameAndTicker(opts *bind.TransactOpts, name_ string, ticker_ string) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "setNameAndTicker", name_, ticker_)
}

// SetNameAndTicker is a paid mutator transaction binding the contract method 0xc388c167.
//
// Solidity: function setNameAndTicker(string name_, string ticker_) returns()
func (_Wewe *WeweSession) SetNameAndTicker(name_ string, ticker_ string) (*types.Transaction, error) {
	return _Wewe.Contract.SetNameAndTicker(&_Wewe.TransactOpts, name_, ticker_)
}

// SetNameAndTicker is a paid mutator transaction binding the contract method 0xc388c167.
//
// Solidity: function setNameAndTicker(string name_, string ticker_) returns()
func (_Wewe *WeweTransactorSession) SetNameAndTicker(name_ string, ticker_ string) (*types.Transaction, error) {
	return _Wewe.Contract.SetNameAndTicker(&_Wewe.TransactOpts, name_, ticker_)
}

// SetWhitelistContract is a paid mutator transaction binding the contract method 0x12f26140.
//
// Solidity: function setWhitelistContract(address newWhitelistContract) returns()
func (_Wewe *WeweTransactor) SetWhitelistContract(opts *bind.TransactOpts, newWhitelistContract common.Address) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "setWhitelistContract", newWhitelistContract)
}

// SetWhitelistContract is a paid mutator transaction binding the contract method 0x12f26140.
//
// Solidity: function setWhitelistContract(address newWhitelistContract) returns()
func (_Wewe *WeweSession) SetWhitelistContract(newWhitelistContract common.Address) (*types.Transaction, error) {
	return _Wewe.Contract.SetWhitelistContract(&_Wewe.TransactOpts, newWhitelistContract)
}

// SetWhitelistContract is a paid mutator transaction binding the contract method 0x12f26140.
//
// Solidity: function setWhitelistContract(address newWhitelistContract) returns()
func (_Wewe *WeweTransactorSession) SetWhitelistContract(newWhitelistContract common.Address) (*types.Transaction, error) {
	return _Wewe.Contract.SetWhitelistContract(&_Wewe.TransactOpts, newWhitelistContract)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Wewe *WeweTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Wewe *WeweSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Transfer(&_Wewe.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Wewe *WeweTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.Transfer(&_Wewe.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Wewe *WeweTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Wewe *WeweSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.TransferFrom(&_Wewe.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Wewe *WeweTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Wewe.Contract.TransferFrom(&_Wewe.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wewe *WeweTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Wewe.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wewe *WeweSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wewe.Contract.TransferOwnership(&_Wewe.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Wewe *WeweTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Wewe.Contract.TransferOwnership(&_Wewe.TransactOpts, newOwner)
}

// WeweApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wewe contract.
type WeweApprovalIterator struct {
	Event *WeweApproval // Event containing the contract specifics and raw log

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
func (it *WeweApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeweApproval)
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
		it.Event = new(WeweApproval)
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
func (it *WeweApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeweApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeweApproval represents a Approval event raised by the Wewe contract.
type WeweApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Wewe *WeweFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WeweApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wewe.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WeweApprovalIterator{contract: _Wewe.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Wewe *WeweFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WeweApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wewe.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeweApproval)
				if err := _Wewe.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Wewe *WeweFilterer) ParseApproval(log types.Log) (*WeweApproval, error) {
	event := new(WeweApproval)
	if err := _Wewe.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeweOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Wewe contract.
type WeweOwnershipTransferredIterator struct {
	Event *WeweOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WeweOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeweOwnershipTransferred)
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
		it.Event = new(WeweOwnershipTransferred)
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
func (it *WeweOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeweOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeweOwnershipTransferred represents a OwnershipTransferred event raised by the Wewe contract.
type WeweOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wewe *WeweFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WeweOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wewe.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WeweOwnershipTransferredIterator{contract: _Wewe.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wewe *WeweFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WeweOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Wewe.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeweOwnershipTransferred)
				if err := _Wewe.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Wewe *WeweFilterer) ParseOwnershipTransferred(log types.Log) (*WeweOwnershipTransferred, error) {
	event := new(WeweOwnershipTransferred)
	if err := _Wewe.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeweTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wewe contract.
type WeweTransferIterator struct {
	Event *WeweTransfer // Event containing the contract specifics and raw log

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
func (it *WeweTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeweTransfer)
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
		it.Event = new(WeweTransfer)
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
func (it *WeweTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeweTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeweTransfer represents a Transfer event raised by the Wewe contract.
type WeweTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Wewe *WeweFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WeweTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wewe.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WeweTransferIterator{contract: _Wewe.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Wewe *WeweFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WeweTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wewe.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeweTransfer)
				if err := _Wewe.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Wewe *WeweFilterer) ParseTransfer(log types.Log) (*WeweTransfer, error) {
	event := new(WeweTransfer)
	if err := _Wewe.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
