// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Erc20ABI is the input ABI used to generate the binding from.
const Erc20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Erc20Bin is the compiled bytecode used for deploying new contracts.
var Erc20Bin = "0x// Code generated - DO NOT EDIT.\n// This file is a generated binding and any manual changes will be lost.\n\npackage erc20\n\nimport (\n        \"math/big\"\n        \"strings\"\n\n        ethereum \"github.com/ethereum/go-ethereum\"\n        \"github.com/ethereum/go-ethereum/accounts/abi\"\n        \"github.com/ethereum/go-ethereum/accounts/abi/bind\"\n        \"github.com/ethereum/go-ethereum/common\"\n        \"github.com/ethereum/go-ethereum/core/types\"\n        \"github.com/ethereum/go-ethereum/event\"\n)\n\n// Reference imports to suppress errors if they are not otherwise used.\nvar (\n        _ = big.NewInt\n        _ = strings.NewReader\n        _ = ethereum.NotFound\n        _ = bind.Bind\n        _ = common.Big1\n        _ = types.BloomLookup\n        _ = event.NewSubscription\n)\n\n\n\n\n\n        // Erc20ABI is the input ABI used to generate the binding from.\n        const Erc20ABI = \"[{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"Approval\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"from\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"Transfer\\\",\\\"type\\\":\\\"event\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"allowance\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"approve\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"balanceOf\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"decimals\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"totalSupply\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"recipient\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"transfer\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"recipient\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"transferFrom\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"}]\"\n\n\n\n\n                // Erc20Bin is the compiled bytecode used for deploying new contracts.\n                var Erc20Bin = \"0x[\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"Approval\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"from\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"Transfer\",\n                \"type\": \"event\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"allowance\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"approve\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"account\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"balanceOf\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"decimals\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"totalSupply\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"recipient\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"transfer\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"sender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"recipient\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"transferFrom\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        }\n]\"\n\n                // DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.\n                func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Erc20, error) {\n                  parsed, err := abi.JSON(strings.NewReader(Erc20ABI))\n                  if err != nil {\n                    return common.Address{}, nil, nil, err\n                  }\n\n                  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Erc20Bin), backend )\n                  if err != nil {\n                    return common.Address{}, nil, nil, err\n                  }\n                  return address, tx, &Erc20{ Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract} }, nil\n                }\n\n\n        // Erc20 is an auto generated Go binding around an Ethereum contract.\n        type Erc20 struct {\n          Erc20Caller     // Read-only binding to the contract\n          Erc20Transactor // Write-only binding to the contract\n          Erc20Filterer   // Log filterer for contract events\n        }\n\n        // Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.\n        type Erc20Caller struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.\n        type Erc20Transactor struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.\n        type Erc20Filterer struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // Erc20Session is an auto generated Go binding around an Ethereum contract,\n        // with pre-set call and transact options.\n        type Erc20Session struct {\n          Contract     *Erc20        // Generic contract binding to set the session for\n          CallOpts     bind.CallOpts     // Call options to use throughout this session\n          TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session\n        }\n\n        // Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,\n        // with pre-set call options.\n        type Erc20CallerSession struct {\n          Contract *Erc20Caller // Generic contract caller binding to set the session for\n          CallOpts bind.CallOpts    // Call options to use throughout this session\n        }\n\n        // Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,\n        // with pre-set transact options.\n        type Erc20TransactorSession struct {\n          Contract     *Erc20Transactor // Generic contract transactor binding to set the session for\n          TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session\n        }\n\n        // Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.\n        type Erc20Raw struct {\n          Contract *Erc20 // Generic contract binding to access the raw methods on\n        }\n\n        // Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.\n        type Erc20CallerRaw struct {\n                Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on\n        }\n\n        // Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.\n        type Erc20TransactorRaw struct {\n                Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on\n        }\n\n        // NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.\n        func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {\n          contract, err := bindErc20(address, backend, backend, backend)\n          if err != nil {\n            return nil, err\n          }\n          return &Erc20{ Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract} }, nil\n        }\n\n        // NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.\n        func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {\n          contract, err := bindErc20(address, caller, nil, nil)\n          if err != nil {\n            return nil, err\n          }\n          return &Erc20Caller{contract: contract}, nil\n        }\n\n        // NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.\n        func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {\n          contract, err := bindErc20(address, nil, transactor, nil)\n          if err != nil {\n            return nil, err\n          }\n          return &Erc20Transactor{contract: contract}, nil\n        }\n\n        // NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.\n        func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {\n          contract, err := bindErc20(address, nil, nil, filterer)\n          if err != nil {\n            return nil, err\n          }\n          return &Erc20Filterer{contract: contract}, nil\n        }\n\n        // bindErc20 binds a generic wrapper to an already deployed contract.\n        func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {\n          parsed, err := abi.JSON(strings.NewReader(Erc20ABI))\n          if err != nil {\n            return nil, err\n          }\n          return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil\n        }\n\n        // Call invokes the (constant) contract method with params as input values and\n        // sets the output to result. The result type might be a single field for simple\n        // returns, a slice of interfaces for anonymous returns and a struct for named\n        // returns.\n        func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {\n                return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)\n        }\n\n        // Transfer initiates a plain transaction to move funds to the contract, calling\n        // its default method if one is available.\n        func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {\n                return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)\n        }\n\n        // Transact invokes the (paid) contract method with params as input values.\n        func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {\n                return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)\n        }\n\n        // Call invokes the (constant) contract method with params as input values and\n        // sets the output to result. The result type might be a single field for simple\n        // returns, a slice of interfaces for anonymous returns and a struct for named\n        // returns.\n        func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {\n                return _Erc20.Contract.contract.Call(opts, result, method, params...)\n        }\n\n        // Transfer initiates a plain transaction to move funds to the contract, calling\n        // its default method if one is available.\n        func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {\n                return _Erc20.Contract.contract.Transfer(opts)\n        }\n\n        // Transact invokes the (paid) contract method with params as input values.\n        func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {\n                return _Erc20.Contract.contract.Transact(opts, method, params...)\n        }\n\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts , owner common.Address , spender common.Address ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Erc20.contract.Call(opts, &out, \"allowance\" , owner, spender)\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Erc20 *Erc20Session) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {\n                  return _Erc20.Contract.Allowance(&_Erc20.CallOpts , owner, spender)\n                }\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Erc20 *Erc20CallerSession) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {\n                  return _Erc20.Contract.Allowance(&_Erc20.CallOpts , owner, spender)\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address account) view returns(uint256)\n                func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts , account common.Address ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Erc20.contract.Call(opts, &out, \"balanceOf\" , account)\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address account) view returns(uint256)\n                func (_Erc20 *Erc20Session) BalanceOf( account common.Address ) ( *big.Int,  error) {\n                  return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts , account)\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address account) view returns(uint256)\n                func (_Erc20 *Erc20CallerSession) BalanceOf( account common.Address ) ( *big.Int,  error) {\n                  return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts , account)\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() view returns(uint256)\n                func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Erc20.contract.Call(opts, &out, \"decimals\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() view returns(uint256)\n                func (_Erc20 *Erc20Session) Decimals() ( *big.Int,  error) {\n                  return _Erc20.Contract.Decimals(&_Erc20.CallOpts )\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() view returns(uint256)\n                func (_Erc20 *Erc20CallerSession) Decimals() ( *big.Int,  error) {\n                  return _Erc20.Contract.Decimals(&_Erc20.CallOpts )\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Erc20.contract.Call(opts, &out, \"totalSupply\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Erc20 *Erc20Session) TotalSupply() ( *big.Int,  error) {\n                  return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts )\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Erc20 *Erc20CallerSession) TotalSupply() ( *big.Int,  error) {\n                  return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts )\n                }\n\n\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts , spender common.Address , amount *big.Int ) (*types.Transaction, error) {\n                        return _Erc20.contract.Transact(opts, \"approve\" , spender, amount)\n                }\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Session) Approve( spender common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.Approve(&_Erc20.TransactOpts , spender, amount)\n                }\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20TransactorSession) Approve( spender common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.Approve(&_Erc20.TransactOpts , spender, amount)\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                        return _Erc20.contract.Transact(opts, \"transfer\" , recipient, amount)\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Session) Transfer( recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.Transfer(&_Erc20.TransactOpts , recipient, amount)\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20TransactorSession) Transfer( recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.Transfer(&_Erc20.TransactOpts , recipient, amount)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts , sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                        return _Erc20.contract.Transact(opts, \"transferFrom\" , sender, recipient, amount)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20Session) TransferFrom( sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts , sender, recipient, amount)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)\n                func (_Erc20 *Erc20TransactorSession) TransferFrom( sender common.Address , recipient common.Address , amount *big.Int ) (*types.Transaction, error) {\n                  return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts , sender, recipient, amount)\n                }\n\n\n\n\n\n\n\n                // Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.\n                type Erc20ApprovalIterator struct {\n                        Event *Erc20Approval // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *Erc20ApprovalIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(Erc20Approval)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(Erc20Approval)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *Erc20ApprovalIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *Erc20ApprovalIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // Erc20Approval represents a Approval event raised by the Erc20 contract.\n                type Erc20Approval struct {\n                        Owner common.Address;\n                        Spender common.Address;\n                        Value *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {\n\n                        var ownerRule []interface{}\n                        for _, ownerItem := range owner {\n                                ownerRule = append(ownerRule, ownerItem)\n                        }\n                        var spenderRule []interface{}\n                        for _, spenderItem := range spender {\n                                spenderRule = append(spenderRule, spenderItem)\n                        }\n\n\n                        logs, sub, err := _Erc20.contract.FilterLogs(opts, \"Approval\", ownerRule, spenderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &Erc20ApprovalIterator{contract: _Erc20.contract, event: \"Approval\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {\n\n                        var ownerRule []interface{}\n                        for _, ownerItem := range owner {\n                                ownerRule = append(ownerRule, ownerItem)\n                        }\n                        var spenderRule []interface{}\n                        for _, spenderItem := range spender {\n                                spenderRule = append(spenderRule, spenderItem)\n                        }\n\n\n                        logs, sub, err := _Erc20.contract.WatchLogs(opts, \"Approval\", ownerRule, spenderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(Erc20Approval)\n                                                if err := _Erc20.contract.UnpackLog(event, \"Approval\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {\n                        event := new(Erc20Approval)\n                        if err := _Erc20.contract.UnpackLog(event, \"Approval\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.\n                type Erc20TransferIterator struct {\n                        Event *Erc20Transfer // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *Erc20TransferIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(Erc20Transfer)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(Erc20Transfer)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *Erc20TransferIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *Erc20TransferIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // Erc20Transfer represents a Transfer event raised by the Erc20 contract.\n                type Erc20Transfer struct {\n                        From common.Address;\n                        To common.Address;\n                        Value *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {\n\n                        var fromRule []interface{}\n                        for _, fromItem := range from {\n                                fromRule = append(fromRule, fromItem)\n                        }\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n\n                        logs, sub, err := _Erc20.contract.FilterLogs(opts, \"Transfer\", fromRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &Erc20TransferIterator{contract: _Erc20.contract, event: \"Transfer\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {\n\n                        var fromRule []interface{}\n                        for _, fromItem := range from {\n                                fromRule = append(fromRule, fromItem)\n                        }\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n\n                        logs, sub, err := _Erc20.contract.WatchLogs(opts, \"Transfer\", fromRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(Erc20Transfer)\n                                                if err := _Erc20.contract.UnpackLog(event, \"Transfer\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {\n                        event := new(Erc20Transfer)\n                        if err := _Erc20.contract.UnpackLog(event, \"Transfer\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }"

// DeployErc20 deploys a new Ethereum contract, binding an instance of Erc20 to it.
func DeployErc20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Erc20, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Erc20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20 *Erc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Erc20 *Erc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20Session) Decimals() (*big.Int, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Erc20 *Erc20CallerSession) Decimals() (*big.Int, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, sender, recipient, amount)
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
