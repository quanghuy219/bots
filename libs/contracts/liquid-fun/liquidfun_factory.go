// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquid_fun_contracts

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

// LiquidFunFactoryMetaData contains all meta data concerning the LiquidFunFactory contract.
var LiquidFunFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"BlueChipMemeLaunched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityManager\",\"type\":\"address\"}],\"name\":\"LiquidityManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"memeRegistry\",\"type\":\"address\"}],\"name\":\"MemeRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pumpVault\",\"type\":\"address\"}],\"name\":\"PumpVaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"swapRouter\",\"type\":\"address\"}],\"name\":\"SwapRouterSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LIQUIDITY_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MEME_REGISTRY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUMP_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STABLE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWAP_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidityManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"memeRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"launchMeme\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityManager\",\"type\":\"address\"}],\"name\":\"setLiquidityManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_memeRegistry\",\"type\":\"address\"}],\"name\":\"setMemeRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pumpVault\",\"type\":\"address\"}],\"name\":\"setPumpVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_swapRouter\",\"type\":\"address\"}],\"name\":\"setSwapRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LiquidFunFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidFunFactoryMetaData.ABI instead.
var LiquidFunFactoryABI = LiquidFunFactoryMetaData.ABI

// LiquidFunFactory is an auto generated Go binding around an Ethereum contract.
type LiquidFunFactory struct {
	LiquidFunFactoryCaller     // Read-only binding to the contract
	LiquidFunFactoryTransactor // Write-only binding to the contract
	LiquidFunFactoryFilterer   // Log filterer for contract events
}

// LiquidFunFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidFunFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidFunFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidFunFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidFunFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidFunFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidFunFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidFunFactorySession struct {
	Contract     *LiquidFunFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidFunFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidFunFactoryCallerSession struct {
	Contract *LiquidFunFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LiquidFunFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidFunFactoryTransactorSession struct {
	Contract     *LiquidFunFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LiquidFunFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidFunFactoryRaw struct {
	Contract *LiquidFunFactory // Generic contract binding to access the raw methods on
}

// LiquidFunFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidFunFactoryCallerRaw struct {
	Contract *LiquidFunFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidFunFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidFunFactoryTransactorRaw struct {
	Contract *LiquidFunFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidFunFactory creates a new instance of LiquidFunFactory, bound to a specific deployed contract.
func NewLiquidFunFactory(address common.Address, backend bind.ContractBackend) (*LiquidFunFactory, error) {
	contract, err := bindLiquidFunFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactory{LiquidFunFactoryCaller: LiquidFunFactoryCaller{contract: contract}, LiquidFunFactoryTransactor: LiquidFunFactoryTransactor{contract: contract}, LiquidFunFactoryFilterer: LiquidFunFactoryFilterer{contract: contract}}, nil
}

// NewLiquidFunFactoryCaller creates a new read-only instance of LiquidFunFactory, bound to a specific deployed contract.
func NewLiquidFunFactoryCaller(address common.Address, caller bind.ContractCaller) (*LiquidFunFactoryCaller, error) {
	contract, err := bindLiquidFunFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryCaller{contract: contract}, nil
}

// NewLiquidFunFactoryTransactor creates a new write-only instance of LiquidFunFactory, bound to a specific deployed contract.
func NewLiquidFunFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidFunFactoryTransactor, error) {
	contract, err := bindLiquidFunFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryTransactor{contract: contract}, nil
}

// NewLiquidFunFactoryFilterer creates a new log filterer instance of LiquidFunFactory, bound to a specific deployed contract.
func NewLiquidFunFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidFunFactoryFilterer, error) {
	contract, err := bindLiquidFunFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryFilterer{contract: contract}, nil
}

// bindLiquidFunFactory binds a generic wrapper to an already deployed contract.
func bindLiquidFunFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidFunFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidFunFactory *LiquidFunFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidFunFactory.Contract.LiquidFunFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidFunFactory *LiquidFunFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.LiquidFunFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidFunFactory *LiquidFunFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.LiquidFunFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidFunFactory *LiquidFunFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidFunFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidFunFactory *LiquidFunFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidFunFactory *LiquidFunFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.contract.Transact(opts, method, params...)
}

// LIQUIDITYMANAGER is a free data retrieval call binding the contract method 0x328ad467.
//
// Solidity: function LIQUIDITY_MANAGER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) LIQUIDITYMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "LIQUIDITY_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIQUIDITYMANAGER is a free data retrieval call binding the contract method 0x328ad467.
//
// Solidity: function LIQUIDITY_MANAGER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) LIQUIDITYMANAGER() (common.Address, error) {
	return _LiquidFunFactory.Contract.LIQUIDITYMANAGER(&_LiquidFunFactory.CallOpts)
}

// LIQUIDITYMANAGER is a free data retrieval call binding the contract method 0x328ad467.
//
// Solidity: function LIQUIDITY_MANAGER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) LIQUIDITYMANAGER() (common.Address, error) {
	return _LiquidFunFactory.Contract.LIQUIDITYMANAGER(&_LiquidFunFactory.CallOpts)
}

// MEMEREGISTRY is a free data retrieval call binding the contract method 0x4e14e8db.
//
// Solidity: function MEME_REGISTRY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) MEMEREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "MEME_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MEMEREGISTRY is a free data retrieval call binding the contract method 0x4e14e8db.
//
// Solidity: function MEME_REGISTRY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) MEMEREGISTRY() (common.Address, error) {
	return _LiquidFunFactory.Contract.MEMEREGISTRY(&_LiquidFunFactory.CallOpts)
}

// MEMEREGISTRY is a free data retrieval call binding the contract method 0x4e14e8db.
//
// Solidity: function MEME_REGISTRY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) MEMEREGISTRY() (common.Address, error) {
	return _LiquidFunFactory.Contract.MEMEREGISTRY(&_LiquidFunFactory.CallOpts)
}

// PUMPVAULT is a free data retrieval call binding the contract method 0x3c880b10.
//
// Solidity: function PUMP_VAULT() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) PUMPVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "PUMP_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUMPVAULT is a free data retrieval call binding the contract method 0x3c880b10.
//
// Solidity: function PUMP_VAULT() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) PUMPVAULT() (common.Address, error) {
	return _LiquidFunFactory.Contract.PUMPVAULT(&_LiquidFunFactory.CallOpts)
}

// PUMPVAULT is a free data retrieval call binding the contract method 0x3c880b10.
//
// Solidity: function PUMP_VAULT() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) PUMPVAULT() (common.Address, error) {
	return _LiquidFunFactory.Contract.PUMPVAULT(&_LiquidFunFactory.CallOpts)
}

// STABLETOKEN is a free data retrieval call binding the contract method 0x7754f887.
//
// Solidity: function STABLE_TOKEN() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) STABLETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "STABLE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STABLETOKEN is a free data retrieval call binding the contract method 0x7754f887.
//
// Solidity: function STABLE_TOKEN() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) STABLETOKEN() (common.Address, error) {
	return _LiquidFunFactory.Contract.STABLETOKEN(&_LiquidFunFactory.CallOpts)
}

// STABLETOKEN is a free data retrieval call binding the contract method 0x7754f887.
//
// Solidity: function STABLE_TOKEN() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) STABLETOKEN() (common.Address, error) {
	return _LiquidFunFactory.Contract.STABLETOKEN(&_LiquidFunFactory.CallOpts)
}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) SWAPROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "SWAP_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) SWAPROUTER() (common.Address, error) {
	return _LiquidFunFactory.Contract.SWAPROUTER(&_LiquidFunFactory.CallOpts)
}

// SWAPROUTER is a free data retrieval call binding the contract method 0xc6005893.
//
// Solidity: function SWAP_ROUTER() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) SWAPROUTER() (common.Address, error) {
	return _LiquidFunFactory.Contract.SWAPROUTER(&_LiquidFunFactory.CallOpts)
}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) UNISWAPFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "UNISWAP_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) UNISWAPFACTORY() (common.Address, error) {
	return _LiquidFunFactory.Contract.UNISWAPFACTORY(&_LiquidFunFactory.CallOpts)
}

// UNISWAPFACTORY is a free data retrieval call binding the contract method 0xc74c0fac.
//
// Solidity: function UNISWAP_FACTORY() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) UNISWAPFACTORY() (common.Address, error) {
	return _LiquidFunFactory.Contract.UNISWAPFACTORY(&_LiquidFunFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidFunFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidFunFactory *LiquidFunFactorySession) Owner() (common.Address, error) {
	return _LiquidFunFactory.Contract.Owner(&_LiquidFunFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LiquidFunFactory *LiquidFunFactoryCallerSession) Owner() (common.Address, error) {
	return _LiquidFunFactory.Contract.Owner(&_LiquidFunFactory.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _owner, address stableToken, address swapRouter, address liquidityManager, address vault, address uniswapFactory, address memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, stableToken common.Address, swapRouter common.Address, liquidityManager common.Address, vault common.Address, uniswapFactory common.Address, memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "initialize", _owner, stableToken, swapRouter, liquidityManager, vault, uniswapFactory, memeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _owner, address stableToken, address swapRouter, address liquidityManager, address vault, address uniswapFactory, address memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) Initialize(_owner common.Address, stableToken common.Address, swapRouter common.Address, liquidityManager common.Address, vault common.Address, uniswapFactory common.Address, memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.Initialize(&_LiquidFunFactory.TransactOpts, _owner, stableToken, swapRouter, liquidityManager, vault, uniswapFactory, memeRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _owner, address stableToken, address swapRouter, address liquidityManager, address vault, address uniswapFactory, address memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) Initialize(_owner common.Address, stableToken common.Address, swapRouter common.Address, liquidityManager common.Address, vault common.Address, uniswapFactory common.Address, memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.Initialize(&_LiquidFunFactory.TransactOpts, _owner, stableToken, swapRouter, liquidityManager, vault, uniswapFactory, memeRegistry)
}

// LaunchMeme is a paid mutator transaction binding the contract method 0xd81a9e17.
//
// Solidity: function launchMeme(string name, string symbol) returns(address token)
func (_LiquidFunFactory *LiquidFunFactoryTransactor) LaunchMeme(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "launchMeme", name, symbol)
}

// LaunchMeme is a paid mutator transaction binding the contract method 0xd81a9e17.
//
// Solidity: function launchMeme(string name, string symbol) returns(address token)
func (_LiquidFunFactory *LiquidFunFactorySession) LaunchMeme(name string, symbol string) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.LaunchMeme(&_LiquidFunFactory.TransactOpts, name, symbol)
}

// LaunchMeme is a paid mutator transaction binding the contract method 0xd81a9e17.
//
// Solidity: function launchMeme(string name, string symbol) returns(address token)
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) LaunchMeme(name string, symbol string) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.LaunchMeme(&_LiquidFunFactory.TransactOpts, name, symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidFunFactory *LiquidFunFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.RenounceOwnership(&_LiquidFunFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.RenounceOwnership(&_LiquidFunFactory.TransactOpts)
}

// SetLiquidityManager is a paid mutator transaction binding the contract method 0x3c1624d4.
//
// Solidity: function setLiquidityManager(address _liquidityManager) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) SetLiquidityManager(opts *bind.TransactOpts, _liquidityManager common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "setLiquidityManager", _liquidityManager)
}

// SetLiquidityManager is a paid mutator transaction binding the contract method 0x3c1624d4.
//
// Solidity: function setLiquidityManager(address _liquidityManager) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) SetLiquidityManager(_liquidityManager common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetLiquidityManager(&_LiquidFunFactory.TransactOpts, _liquidityManager)
}

// SetLiquidityManager is a paid mutator transaction binding the contract method 0x3c1624d4.
//
// Solidity: function setLiquidityManager(address _liquidityManager) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) SetLiquidityManager(_liquidityManager common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetLiquidityManager(&_LiquidFunFactory.TransactOpts, _liquidityManager)
}

// SetMemeRegistry is a paid mutator transaction binding the contract method 0x66a14956.
//
// Solidity: function setMemeRegistry(address _memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) SetMemeRegistry(opts *bind.TransactOpts, _memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "setMemeRegistry", _memeRegistry)
}

// SetMemeRegistry is a paid mutator transaction binding the contract method 0x66a14956.
//
// Solidity: function setMemeRegistry(address _memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) SetMemeRegistry(_memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetMemeRegistry(&_LiquidFunFactory.TransactOpts, _memeRegistry)
}

// SetMemeRegistry is a paid mutator transaction binding the contract method 0x66a14956.
//
// Solidity: function setMemeRegistry(address _memeRegistry) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) SetMemeRegistry(_memeRegistry common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetMemeRegistry(&_LiquidFunFactory.TransactOpts, _memeRegistry)
}

// SetPumpVault is a paid mutator transaction binding the contract method 0xde9dfbba.
//
// Solidity: function setPumpVault(address _pumpVault) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) SetPumpVault(opts *bind.TransactOpts, _pumpVault common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "setPumpVault", _pumpVault)
}

// SetPumpVault is a paid mutator transaction binding the contract method 0xde9dfbba.
//
// Solidity: function setPumpVault(address _pumpVault) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) SetPumpVault(_pumpVault common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetPumpVault(&_LiquidFunFactory.TransactOpts, _pumpVault)
}

// SetPumpVault is a paid mutator transaction binding the contract method 0xde9dfbba.
//
// Solidity: function setPumpVault(address _pumpVault) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) SetPumpVault(_pumpVault common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetPumpVault(&_LiquidFunFactory.TransactOpts, _pumpVault)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) SetSwapRouter(opts *bind.TransactOpts, _swapRouter common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "setSwapRouter", _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetSwapRouter(&_LiquidFunFactory.TransactOpts, _swapRouter)
}

// SetSwapRouter is a paid mutator transaction binding the contract method 0x41273657.
//
// Solidity: function setSwapRouter(address _swapRouter) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) SetSwapRouter(_swapRouter common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.SetSwapRouter(&_LiquidFunFactory.TransactOpts, _swapRouter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidFunFactory *LiquidFunFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.TransferOwnership(&_LiquidFunFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LiquidFunFactory *LiquidFunFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LiquidFunFactory.Contract.TransferOwnership(&_LiquidFunFactory.TransactOpts, newOwner)
}

// LiquidFunFactoryBlueChipMemeLaunchedIterator is returned from FilterBlueChipMemeLaunched and is used to iterate over the raw logs and unpacked data for BlueChipMemeLaunched events raised by the LiquidFunFactory contract.
type LiquidFunFactoryBlueChipMemeLaunchedIterator struct {
	Event *LiquidFunFactoryBlueChipMemeLaunched // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryBlueChipMemeLaunchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryBlueChipMemeLaunched)
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
		it.Event = new(LiquidFunFactoryBlueChipMemeLaunched)
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
func (it *LiquidFunFactoryBlueChipMemeLaunchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryBlueChipMemeLaunchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryBlueChipMemeLaunched represents a BlueChipMemeLaunched event raised by the LiquidFunFactory contract.
type LiquidFunFactoryBlueChipMemeLaunched struct {
	TokenId *big.Int
	Token   common.Address
	Name    string
	Symbol  string
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlueChipMemeLaunched is a free log retrieval operation binding the contract event 0x7843a4e38906df6197ae6885d7e728368c930ba313f87b1751aa2f51181a6a8b.
//
// Solidity: event BlueChipMemeLaunched(uint256 indexed tokenId, address indexed token, string name, string symbol, address creator)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterBlueChipMemeLaunched(opts *bind.FilterOpts, tokenId []*big.Int, token []common.Address) (*LiquidFunFactoryBlueChipMemeLaunchedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "BlueChipMemeLaunched", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryBlueChipMemeLaunchedIterator{contract: _LiquidFunFactory.contract, event: "BlueChipMemeLaunched", logs: logs, sub: sub}, nil
}

// WatchBlueChipMemeLaunched is a free log subscription operation binding the contract event 0x7843a4e38906df6197ae6885d7e728368c930ba313f87b1751aa2f51181a6a8b.
//
// Solidity: event BlueChipMemeLaunched(uint256 indexed tokenId, address indexed token, string name, string symbol, address creator)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchBlueChipMemeLaunched(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryBlueChipMemeLaunched, tokenId []*big.Int, token []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "BlueChipMemeLaunched", tokenIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryBlueChipMemeLaunched)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "BlueChipMemeLaunched", log); err != nil {
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

// ParseBlueChipMemeLaunched is a log parse operation binding the contract event 0x7843a4e38906df6197ae6885d7e728368c930ba313f87b1751aa2f51181a6a8b.
//
// Solidity: event BlueChipMemeLaunched(uint256 indexed tokenId, address indexed token, string name, string symbol, address creator)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseBlueChipMemeLaunched(log types.Log) (*LiquidFunFactoryBlueChipMemeLaunched, error) {
	event := new(LiquidFunFactoryBlueChipMemeLaunched)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "BlueChipMemeLaunched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LiquidFunFactory contract.
type LiquidFunFactoryInitializedIterator struct {
	Event *LiquidFunFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryInitialized)
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
		it.Event = new(LiquidFunFactoryInitialized)
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
func (it *LiquidFunFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryInitialized represents a Initialized event raised by the LiquidFunFactory contract.
type LiquidFunFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiquidFunFactoryInitializedIterator, error) {

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryInitializedIterator{contract: _LiquidFunFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryInitialized)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseInitialized(log types.Log) (*LiquidFunFactoryInitialized, error) {
	event := new(LiquidFunFactoryInitialized)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactoryLiquidityManagerSetIterator is returned from FilterLiquidityManagerSet and is used to iterate over the raw logs and unpacked data for LiquidityManagerSet events raised by the LiquidFunFactory contract.
type LiquidFunFactoryLiquidityManagerSetIterator struct {
	Event *LiquidFunFactoryLiquidityManagerSet // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryLiquidityManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryLiquidityManagerSet)
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
		it.Event = new(LiquidFunFactoryLiquidityManagerSet)
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
func (it *LiquidFunFactoryLiquidityManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryLiquidityManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryLiquidityManagerSet represents a LiquidityManagerSet event raised by the LiquidFunFactory contract.
type LiquidFunFactoryLiquidityManagerSet struct {
	LiquidityManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityManagerSet is a free log retrieval operation binding the contract event 0xc783f916f4a84aafa6573e6ffd2cd7d99dbf370e160fbaa2d6c585ae7247737e.
//
// Solidity: event LiquidityManagerSet(address liquidityManager)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterLiquidityManagerSet(opts *bind.FilterOpts) (*LiquidFunFactoryLiquidityManagerSetIterator, error) {

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "LiquidityManagerSet")
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryLiquidityManagerSetIterator{contract: _LiquidFunFactory.contract, event: "LiquidityManagerSet", logs: logs, sub: sub}, nil
}

// WatchLiquidityManagerSet is a free log subscription operation binding the contract event 0xc783f916f4a84aafa6573e6ffd2cd7d99dbf370e160fbaa2d6c585ae7247737e.
//
// Solidity: event LiquidityManagerSet(address liquidityManager)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchLiquidityManagerSet(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryLiquidityManagerSet) (event.Subscription, error) {

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "LiquidityManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryLiquidityManagerSet)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "LiquidityManagerSet", log); err != nil {
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

// ParseLiquidityManagerSet is a log parse operation binding the contract event 0xc783f916f4a84aafa6573e6ffd2cd7d99dbf370e160fbaa2d6c585ae7247737e.
//
// Solidity: event LiquidityManagerSet(address liquidityManager)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseLiquidityManagerSet(log types.Log) (*LiquidFunFactoryLiquidityManagerSet, error) {
	event := new(LiquidFunFactoryLiquidityManagerSet)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "LiquidityManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactoryMemeRegistrySetIterator is returned from FilterMemeRegistrySet and is used to iterate over the raw logs and unpacked data for MemeRegistrySet events raised by the LiquidFunFactory contract.
type LiquidFunFactoryMemeRegistrySetIterator struct {
	Event *LiquidFunFactoryMemeRegistrySet // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryMemeRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryMemeRegistrySet)
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
		it.Event = new(LiquidFunFactoryMemeRegistrySet)
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
func (it *LiquidFunFactoryMemeRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryMemeRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryMemeRegistrySet represents a MemeRegistrySet event raised by the LiquidFunFactory contract.
type LiquidFunFactoryMemeRegistrySet struct {
	MemeRegistry common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMemeRegistrySet is a free log retrieval operation binding the contract event 0x7412c986896c51559f78d45350a43e8851694d25fe55554d4df25b25def375d5.
//
// Solidity: event MemeRegistrySet(address memeRegistry)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterMemeRegistrySet(opts *bind.FilterOpts) (*LiquidFunFactoryMemeRegistrySetIterator, error) {

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "MemeRegistrySet")
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryMemeRegistrySetIterator{contract: _LiquidFunFactory.contract, event: "MemeRegistrySet", logs: logs, sub: sub}, nil
}

// WatchMemeRegistrySet is a free log subscription operation binding the contract event 0x7412c986896c51559f78d45350a43e8851694d25fe55554d4df25b25def375d5.
//
// Solidity: event MemeRegistrySet(address memeRegistry)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchMemeRegistrySet(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryMemeRegistrySet) (event.Subscription, error) {

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "MemeRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryMemeRegistrySet)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "MemeRegistrySet", log); err != nil {
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

// ParseMemeRegistrySet is a log parse operation binding the contract event 0x7412c986896c51559f78d45350a43e8851694d25fe55554d4df25b25def375d5.
//
// Solidity: event MemeRegistrySet(address memeRegistry)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseMemeRegistrySet(log types.Log) (*LiquidFunFactoryMemeRegistrySet, error) {
	event := new(LiquidFunFactoryMemeRegistrySet)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "MemeRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LiquidFunFactory contract.
type LiquidFunFactoryOwnershipTransferredIterator struct {
	Event *LiquidFunFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryOwnershipTransferred)
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
		it.Event = new(LiquidFunFactoryOwnershipTransferred)
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
func (it *LiquidFunFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the LiquidFunFactory contract.
type LiquidFunFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidFunFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryOwnershipTransferredIterator{contract: _LiquidFunFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryOwnershipTransferred)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidFunFactoryOwnershipTransferred, error) {
	event := new(LiquidFunFactoryOwnershipTransferred)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactoryPumpVaultSetIterator is returned from FilterPumpVaultSet and is used to iterate over the raw logs and unpacked data for PumpVaultSet events raised by the LiquidFunFactory contract.
type LiquidFunFactoryPumpVaultSetIterator struct {
	Event *LiquidFunFactoryPumpVaultSet // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactoryPumpVaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactoryPumpVaultSet)
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
		it.Event = new(LiquidFunFactoryPumpVaultSet)
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
func (it *LiquidFunFactoryPumpVaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactoryPumpVaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactoryPumpVaultSet represents a PumpVaultSet event raised by the LiquidFunFactory contract.
type LiquidFunFactoryPumpVaultSet struct {
	PumpVault common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPumpVaultSet is a free log retrieval operation binding the contract event 0x4388c69db632a80459bc7f7523b9ea93e09cdce38b3b24e443bb42031be865eb.
//
// Solidity: event PumpVaultSet(address pumpVault)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterPumpVaultSet(opts *bind.FilterOpts) (*LiquidFunFactoryPumpVaultSetIterator, error) {

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "PumpVaultSet")
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactoryPumpVaultSetIterator{contract: _LiquidFunFactory.contract, event: "PumpVaultSet", logs: logs, sub: sub}, nil
}

// WatchPumpVaultSet is a free log subscription operation binding the contract event 0x4388c69db632a80459bc7f7523b9ea93e09cdce38b3b24e443bb42031be865eb.
//
// Solidity: event PumpVaultSet(address pumpVault)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchPumpVaultSet(opts *bind.WatchOpts, sink chan<- *LiquidFunFactoryPumpVaultSet) (event.Subscription, error) {

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "PumpVaultSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactoryPumpVaultSet)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "PumpVaultSet", log); err != nil {
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

// ParsePumpVaultSet is a log parse operation binding the contract event 0x4388c69db632a80459bc7f7523b9ea93e09cdce38b3b24e443bb42031be865eb.
//
// Solidity: event PumpVaultSet(address pumpVault)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParsePumpVaultSet(log types.Log) (*LiquidFunFactoryPumpVaultSet, error) {
	event := new(LiquidFunFactoryPumpVaultSet)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "PumpVaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidFunFactorySwapRouterSetIterator is returned from FilterSwapRouterSet and is used to iterate over the raw logs and unpacked data for SwapRouterSet events raised by the LiquidFunFactory contract.
type LiquidFunFactorySwapRouterSetIterator struct {
	Event *LiquidFunFactorySwapRouterSet // Event containing the contract specifics and raw log

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
func (it *LiquidFunFactorySwapRouterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidFunFactorySwapRouterSet)
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
		it.Event = new(LiquidFunFactorySwapRouterSet)
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
func (it *LiquidFunFactorySwapRouterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidFunFactorySwapRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidFunFactorySwapRouterSet represents a SwapRouterSet event raised by the LiquidFunFactory contract.
type LiquidFunFactorySwapRouterSet struct {
	SwapRouter common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwapRouterSet is a free log retrieval operation binding the contract event 0x6307ba8e3a4d6f90cda83ffa9c59c26256a075f79abca3852f6539fc6d44653f.
//
// Solidity: event SwapRouterSet(address swapRouter)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) FilterSwapRouterSet(opts *bind.FilterOpts) (*LiquidFunFactorySwapRouterSetIterator, error) {

	logs, sub, err := _LiquidFunFactory.contract.FilterLogs(opts, "SwapRouterSet")
	if err != nil {
		return nil, err
	}
	return &LiquidFunFactorySwapRouterSetIterator{contract: _LiquidFunFactory.contract, event: "SwapRouterSet", logs: logs, sub: sub}, nil
}

// WatchSwapRouterSet is a free log subscription operation binding the contract event 0x6307ba8e3a4d6f90cda83ffa9c59c26256a075f79abca3852f6539fc6d44653f.
//
// Solidity: event SwapRouterSet(address swapRouter)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) WatchSwapRouterSet(opts *bind.WatchOpts, sink chan<- *LiquidFunFactorySwapRouterSet) (event.Subscription, error) {

	logs, sub, err := _LiquidFunFactory.contract.WatchLogs(opts, "SwapRouterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidFunFactorySwapRouterSet)
				if err := _LiquidFunFactory.contract.UnpackLog(event, "SwapRouterSet", log); err != nil {
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

// ParseSwapRouterSet is a log parse operation binding the contract event 0x6307ba8e3a4d6f90cda83ffa9c59c26256a075f79abca3852f6539fc6d44653f.
//
// Solidity: event SwapRouterSet(address swapRouter)
func (_LiquidFunFactory *LiquidFunFactoryFilterer) ParseSwapRouterSet(log types.Log) (*LiquidFunFactorySwapRouterSet, error) {
	event := new(LiquidFunFactorySwapRouterSet)
	if err := _LiquidFunFactory.contract.UnpackLog(event, "SwapRouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
