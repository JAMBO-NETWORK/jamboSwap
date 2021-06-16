// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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

// PairABI is the input ABI used to generate the binding from.
const PairABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PairBin is the compiled bytecode used for deploying new contracts.
var PairBin = "0x// Code generated - DO NOT EDIT.\n// This file is a generated binding and any manual changes will be lost.\n\npackage pair\n\nimport (\n        \"math/big\"\n        \"strings\"\n\n        ethereum \"github.com/ethereum/go-ethereum\"\n        \"github.com/ethereum/go-ethereum/accounts/abi\"\n        \"github.com/ethereum/go-ethereum/accounts/abi/bind\"\n        \"github.com/ethereum/go-ethereum/common\"\n        \"github.com/ethereum/go-ethereum/core/types\"\n        \"github.com/ethereum/go-ethereum/event\"\n)\n\n// Reference imports to suppress errors if they are not otherwise used.\nvar (\n        _ = big.NewInt\n        _ = strings.NewReader\n        _ = ethereum.NotFound\n        _ = bind.Bind\n        _ = common.Big1\n        _ = types.BloomLookup\n        _ = event.NewSubscription\n)\n\n\n\n\n\n        // PairABI is the input ABI used to generate the binding from.\n        const PairABI = \"[{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"Approval\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"Burn\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"Mint\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0In\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1In\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0Out\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1Out\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"Swap\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint112\\\",\\\"name\\\":\\\"reserve0\\\",\\\"type\\\":\\\"uint112\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint112\\\",\\\"name\\\":\\\"reserve1\\\",\\\"type\\\":\\\"uint112\\\"}],\\\"name\\\":\\\"Sync\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"from\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":false,\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"Transfer\\\",\\\"type\\\":\\\"event\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"DOMAIN_SEPARATOR\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"MINIMUM_LIQUIDITY\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"PERMIT_TYPEHASH\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"allowance\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"approve\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"balanceOf\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"burn\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"decimals\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint8\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint8\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"factory\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"getReserves\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint112\\\",\\\"name\\\":\\\"reserve0\\\",\\\"type\\\":\\\"uint112\\\"},{\\\"internalType\\\":\\\"uint112\\\",\\\"name\\\":\\\"reserve1\\\",\\\"type\\\":\\\"uint112\\\"},{\\\"internalType\\\":\\\"uint32\\\",\\\"name\\\":\\\"blockTimestampLast\\\",\\\"type\\\":\\\"uint32\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"initialize\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"kLast\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"mint\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"liquidity\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"name\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"string\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"nonces\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"owner\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"spender\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"deadline\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"uint8\\\",\\\"name\\\":\\\"v\\\",\\\"type\\\":\\\"uint8\\\"},{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"r\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"s\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"permit\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"price0CumulativeLast\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"price1CumulativeLast\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"skim\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount0Out\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"amount1Out\\\",\\\"type\\\":\\\"uint256\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"bytes\\\",\\\"name\\\":\\\"data\\\",\\\"type\\\":\\\"bytes\\\"}],\\\"name\\\":\\\"swap\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"symbol\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"string\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"string\\\"}],\\\"stateMutability\\\":\\\"pure\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"sync\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"token0\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"token1\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"totalSupply\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"transfer\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"from\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"to\\\",\\\"type\\\":\\\"address\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"value\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"transferFrom\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"}]\"\n\n\n\n\n                // PairBin is the compiled bytecode used for deploying new contracts.\n                var PairBin = \"0x[\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"Approval\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"sender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"Burn\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"sender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"Mint\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"sender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0In\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1In\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0Out\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1Out\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"Swap\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint112\",\n                                \"name\": \"reserve0\",\n                                \"type\": \"uint112\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint112\",\n                                \"name\": \"reserve1\",\n                                \"type\": \"uint112\"\n                        }\n                ],\n                \"name\": \"Sync\",\n                \"type\": \"event\"\n        },\n        {\n                \"anonymous\": false,\n                \"inputs\": [\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"from\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": true,\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"indexed\": false,\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"Transfer\",\n                \"type\": \"event\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"DOMAIN_SEPARATOR\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bytes32\",\n                                \"name\": \"\",\n                                \"type\": \"bytes32\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"MINIMUM_LIQUIDITY\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"pure\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"PERMIT_TYPEHASH\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bytes32\",\n                                \"name\": \"\",\n                                \"type\": \"bytes32\"\n                        }\n                ],\n                \"stateMutability\": \"pure\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"allowance\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"approve\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"balanceOf\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"burn\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"decimals\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint8\",\n                                \"name\": \"\",\n                                \"type\": \"uint8\"\n                        }\n                ],\n                \"stateMutability\": \"pure\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"factory\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"getReserves\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint112\",\n                                \"name\": \"reserve0\",\n                                \"type\": \"uint112\"\n                        },\n                        {\n                                \"internalType\": \"uint112\",\n                                \"name\": \"reserve1\",\n                                \"type\": \"uint112\"\n                        },\n                        {\n                                \"internalType\": \"uint32\",\n                                \"name\": \"blockTimestampLast\",\n                                \"type\": \"uint32\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"initialize\",\n                \"outputs\": [],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"kLast\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"mint\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"liquidity\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"name\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"string\",\n                                \"name\": \"\",\n                                \"type\": \"string\"\n                        }\n                ],\n                \"stateMutability\": \"pure\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"nonces\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"owner\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"spender\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"deadline\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"internalType\": \"uint8\",\n                                \"name\": \"v\",\n                                \"type\": \"uint8\"\n                        },\n                        {\n                                \"internalType\": \"bytes32\",\n                                \"name\": \"r\",\n                                \"type\": \"bytes32\"\n                        },\n                        {\n                                \"internalType\": \"bytes32\",\n                                \"name\": \"s\",\n                                \"type\": \"bytes32\"\n                        }\n                ],\n                \"name\": \"permit\",\n                \"outputs\": [],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"price0CumulativeLast\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"price1CumulativeLast\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"name\": \"skim\",\n                \"outputs\": [],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount0Out\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"amount1Out\",\n                                \"type\": \"uint256\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"bytes\",\n                                \"name\": \"data\",\n                                \"type\": \"bytes\"\n                        }\n                ],\n                \"name\": \"swap\",\n                \"outputs\": [],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"symbol\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"string\",\n                                \"name\": \"\",\n                                \"type\": \"string\"\n                        }\n                ],\n                \"stateMutability\": \"pure\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"sync\",\n                \"outputs\": [],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"token0\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"token1\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"\",\n                                \"type\": \"address\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [],\n                \"name\": \"totalSupply\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"stateMutability\": \"view\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"transfer\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        },\n        {\n                \"inputs\": [\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"from\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"address\",\n                                \"name\": \"to\",\n                                \"type\": \"address\"\n                        },\n                        {\n                                \"internalType\": \"uint256\",\n                                \"name\": \"value\",\n                                \"type\": \"uint256\"\n                        }\n                ],\n                \"name\": \"transferFrom\",\n                \"outputs\": [\n                        {\n                                \"internalType\": \"bool\",\n                                \"name\": \"\",\n                                \"type\": \"bool\"\n                        }\n                ],\n                \"stateMutability\": \"nonpayable\",\n                \"type\": \"function\"\n        }\n]\"\n\n                // DeployPair deploys a new Ethereum contract, binding an instance of Pair to it.\n                func DeployPair(auth *bind.TransactOpts, backend bind.ContractBackend ) (common.Address, *types.Transaction, *Pair, error) {\n                  parsed, err := abi.JSON(strings.NewReader(PairABI))\n                  if err != nil {\n                    return common.Address{}, nil, nil, err\n                  }\n\n                  address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairBin), backend )\n                  if err != nil {\n                    return common.Address{}, nil, nil, err\n                  }\n                  return address, tx, &Pair{ PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract} }, nil\n                }\n\n\n        // Pair is an auto generated Go binding around an Ethereum contract.\n        type Pair struct {\n          PairCaller     // Read-only binding to the contract\n          PairTransactor // Write-only binding to the contract\n          PairFilterer   // Log filterer for contract events\n        }\n\n        // PairCaller is an auto generated read-only Go binding around an Ethereum contract.\n        type PairCaller struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // PairTransactor is an auto generated write-only Go binding around an Ethereum contract.\n        type PairTransactor struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.\n        type PairFilterer struct {\n          contract *bind.BoundContract // Generic contract wrapper for the low level calls\n        }\n\n        // PairSession is an auto generated Go binding around an Ethereum contract,\n        // with pre-set call and transact options.\n        type PairSession struct {\n          Contract     *Pair        // Generic contract binding to set the session for\n          CallOpts     bind.CallOpts     // Call options to use throughout this session\n          TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session\n        }\n\n        // PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,\n        // with pre-set call options.\n        type PairCallerSession struct {\n          Contract *PairCaller // Generic contract caller binding to set the session for\n          CallOpts bind.CallOpts    // Call options to use throughout this session\n        }\n\n        // PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,\n        // with pre-set transact options.\n        type PairTransactorSession struct {\n          Contract     *PairTransactor // Generic contract transactor binding to set the session for\n          TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session\n        }\n\n        // PairRaw is an auto generated low-level Go binding around an Ethereum contract.\n        type PairRaw struct {\n          Contract *Pair // Generic contract binding to access the raw methods on\n        }\n\n        // PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.\n        type PairCallerRaw struct {\n                Contract *PairCaller // Generic read-only contract binding to access the raw methods on\n        }\n\n        // PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.\n        type PairTransactorRaw struct {\n                Contract *PairTransactor // Generic write-only contract binding to access the raw methods on\n        }\n\n        // NewPair creates a new instance of Pair, bound to a specific deployed contract.\n        func NewPair(address common.Address, backend bind.ContractBackend) (*Pair, error) {\n          contract, err := bindPair(address, backend, backend, backend)\n          if err != nil {\n            return nil, err\n          }\n          return &Pair{ PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract} }, nil\n        }\n\n        // NewPairCaller creates a new read-only instance of Pair, bound to a specific deployed contract.\n        func NewPairCaller(address common.Address, caller bind.ContractCaller) (*PairCaller, error) {\n          contract, err := bindPair(address, caller, nil, nil)\n          if err != nil {\n            return nil, err\n          }\n          return &PairCaller{contract: contract}, nil\n        }\n\n        // NewPairTransactor creates a new write-only instance of Pair, bound to a specific deployed contract.\n        func NewPairTransactor(address common.Address, transactor bind.ContractTransactor) (*PairTransactor, error) {\n          contract, err := bindPair(address, nil, transactor, nil)\n          if err != nil {\n            return nil, err\n          }\n          return &PairTransactor{contract: contract}, nil\n        }\n\n        // NewPairFilterer creates a new log filterer instance of Pair, bound to a specific deployed contract.\n        func NewPairFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFilterer, error) {\n          contract, err := bindPair(address, nil, nil, filterer)\n          if err != nil {\n            return nil, err\n          }\n          return &PairFilterer{contract: contract}, nil\n        }\n\n        // bindPair binds a generic wrapper to an already deployed contract.\n        func bindPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {\n          parsed, err := abi.JSON(strings.NewReader(PairABI))\n          if err != nil {\n            return nil, err\n          }\n          return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil\n        }\n\n        // Call invokes the (constant) contract method with params as input values and\n        // sets the output to result. The result type might be a single field for simple\n        // returns, a slice of interfaces for anonymous returns and a struct for named\n        // returns.\n        func (_Pair *PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {\n                return _Pair.Contract.PairCaller.contract.Call(opts, result, method, params...)\n        }\n\n        // Transfer initiates a plain transaction to move funds to the contract, calling\n        // its default method if one is available.\n        func (_Pair *PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {\n                return _Pair.Contract.PairTransactor.contract.Transfer(opts)\n        }\n\n        // Transact invokes the (paid) contract method with params as input values.\n        func (_Pair *PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {\n                return _Pair.Contract.PairTransactor.contract.Transact(opts, method, params...)\n        }\n\n        // Call invokes the (constant) contract method with params as input values and\n        // sets the output to result. The result type might be a single field for simple\n        // returns, a slice of interfaces for anonymous returns and a struct for named\n        // returns.\n        func (_Pair *PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {\n                return _Pair.Contract.contract.Call(opts, result, method, params...)\n        }\n\n        // Transfer initiates a plain transaction to move funds to the contract, calling\n        // its default method if one is available.\n        func (_Pair *PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {\n                return _Pair.Contract.contract.Transfer(opts)\n        }\n\n        // Transact invokes the (paid) contract method with params as input values.\n        func (_Pair *PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {\n                return _Pair.Contract.contract.Transact(opts, method, params...)\n        }\n\n\n                // DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.\n                //\n                // Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)\n                func (_Pair *PairCaller) DOMAINSEPARATOR(opts *bind.CallOpts ) ([32]byte, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"DOMAIN_SEPARATOR\" )\n\n                        if err != nil {\n                                return *new([32]byte),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)\n\n                        return out0,  err\n\n                }\n\n                // DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.\n                //\n                // Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)\n                func (_Pair *PairSession) DOMAINSEPARATOR() ( [32]byte,  error) {\n                  return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts )\n                }\n\n                // DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.\n                //\n                // Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)\n                func (_Pair *PairCallerSession) DOMAINSEPARATOR() ( [32]byte,  error) {\n                  return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts )\n                }\n\n                // MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.\n                //\n                // Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)\n                func (_Pair *PairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"MINIMUM_LIQUIDITY\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.\n                //\n                // Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)\n                func (_Pair *PairSession) MINIMUMLIQUIDITY() ( *big.Int,  error) {\n                  return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts )\n                }\n\n                // MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.\n                //\n                // Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)\n                func (_Pair *PairCallerSession) MINIMUMLIQUIDITY() ( *big.Int,  error) {\n                  return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts )\n                }\n\n                // PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.\n                //\n                // Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)\n                func (_Pair *PairCaller) PERMITTYPEHASH(opts *bind.CallOpts ) ([32]byte, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"PERMIT_TYPEHASH\" )\n\n                        if err != nil {\n                                return *new([32]byte),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)\n\n                        return out0,  err\n\n                }\n\n                // PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.\n                //\n                // Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)\n                func (_Pair *PairSession) PERMITTYPEHASH() ( [32]byte,  error) {\n                  return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts )\n                }\n\n                // PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.\n                //\n                // Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)\n                func (_Pair *PairCallerSession) PERMITTYPEHASH() ( [32]byte,  error) {\n                  return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts )\n                }\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Pair *PairCaller) Allowance(opts *bind.CallOpts , owner common.Address , spender common.Address ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"allowance\" , owner, spender)\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Pair *PairSession) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.Allowance(&_Pair.CallOpts , owner, spender)\n                }\n\n                // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.\n                //\n                // Solidity: function allowance(address owner, address spender) view returns(uint256)\n                func (_Pair *PairCallerSession) Allowance( owner common.Address , spender common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.Allowance(&_Pair.CallOpts , owner, spender)\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address owner) view returns(uint256)\n                func (_Pair *PairCaller) BalanceOf(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"balanceOf\" , owner)\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address owner) view returns(uint256)\n                func (_Pair *PairSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.BalanceOf(&_Pair.CallOpts , owner)\n                }\n\n                // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.\n                //\n                // Solidity: function balanceOf(address owner) view returns(uint256)\n                func (_Pair *PairCallerSession) BalanceOf( owner common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.BalanceOf(&_Pair.CallOpts , owner)\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() pure returns(uint8)\n                func (_Pair *PairCaller) Decimals(opts *bind.CallOpts ) (uint8, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"decimals\" )\n\n                        if err != nil {\n                                return *new(uint8),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)\n\n                        return out0,  err\n\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() pure returns(uint8)\n                func (_Pair *PairSession) Decimals() ( uint8,  error) {\n                  return _Pair.Contract.Decimals(&_Pair.CallOpts )\n                }\n\n                // Decimals is a free data retrieval call binding the contract method 0x313ce567.\n                //\n                // Solidity: function decimals() pure returns(uint8)\n                func (_Pair *PairCallerSession) Decimals() ( uint8,  error) {\n                  return _Pair.Contract.Decimals(&_Pair.CallOpts )\n                }\n\n                // Factory is a free data retrieval call binding the contract method 0xc45a0155.\n                //\n                // Solidity: function factory() view returns(address)\n                func (_Pair *PairCaller) Factory(opts *bind.CallOpts ) (common.Address, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"factory\" )\n\n                        if err != nil {\n                                return *new(common.Address),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)\n\n                        return out0,  err\n\n                }\n\n                // Factory is a free data retrieval call binding the contract method 0xc45a0155.\n                //\n                // Solidity: function factory() view returns(address)\n                func (_Pair *PairSession) Factory() ( common.Address,  error) {\n                  return _Pair.Contract.Factory(&_Pair.CallOpts )\n                }\n\n                // Factory is a free data retrieval call binding the contract method 0xc45a0155.\n                //\n                // Solidity: function factory() view returns(address)\n                func (_Pair *PairCallerSession) Factory() ( common.Address,  error) {\n                  return _Pair.Contract.Factory(&_Pair.CallOpts )\n                }\n\n                // GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.\n                //\n                // Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)\n                func (_Pair *PairCaller) GetReserves(opts *bind.CallOpts ) (struct{ Reserve0 *big.Int;Reserve1 *big.Int;BlockTimestampLast uint32; }, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"getReserves\" )\n\n                        outstruct := new(struct{  Reserve0 *big.Int;  Reserve1 *big.Int;  BlockTimestampLast uint32;  })\n                        if err != nil {\n                                return *outstruct, err\n                        }\n\n                        outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n                        outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)\n                        outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)\n\n                        return *outstruct, err\n\n                }\n\n                // GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.\n                //\n                // Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)\n                func (_Pair *PairSession) GetReserves() (struct{ Reserve0 *big.Int;Reserve1 *big.Int;BlockTimestampLast uint32; },  error) {\n                  return _Pair.Contract.GetReserves(&_Pair.CallOpts )\n                }\n\n                // GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.\n                //\n                // Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)\n                func (_Pair *PairCallerSession) GetReserves() (struct{ Reserve0 *big.Int;Reserve1 *big.Int;BlockTimestampLast uint32; },  error) {\n                  return _Pair.Contract.GetReserves(&_Pair.CallOpts )\n                }\n\n                // KLast is a free data retrieval call binding the contract method 0x7464fc3d.\n                //\n                // Solidity: function kLast() view returns(uint256)\n                func (_Pair *PairCaller) KLast(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"kLast\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // KLast is a free data retrieval call binding the contract method 0x7464fc3d.\n                //\n                // Solidity: function kLast() view returns(uint256)\n                func (_Pair *PairSession) KLast() ( *big.Int,  error) {\n                  return _Pair.Contract.KLast(&_Pair.CallOpts )\n                }\n\n                // KLast is a free data retrieval call binding the contract method 0x7464fc3d.\n                //\n                // Solidity: function kLast() view returns(uint256)\n                func (_Pair *PairCallerSession) KLast() ( *big.Int,  error) {\n                  return _Pair.Contract.KLast(&_Pair.CallOpts )\n                }\n\n                // Name is a free data retrieval call binding the contract method 0x06fdde03.\n                //\n                // Solidity: function name() pure returns(string)\n                func (_Pair *PairCaller) Name(opts *bind.CallOpts ) (string, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"name\" )\n\n                        if err != nil {\n                                return *new(string),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(string)).(*string)\n\n                        return out0,  err\n\n                }\n\n                // Name is a free data retrieval call binding the contract method 0x06fdde03.\n                //\n                // Solidity: function name() pure returns(string)\n                func (_Pair *PairSession) Name() ( string,  error) {\n                  return _Pair.Contract.Name(&_Pair.CallOpts )\n                }\n\n                // Name is a free data retrieval call binding the contract method 0x06fdde03.\n                //\n                // Solidity: function name() pure returns(string)\n                func (_Pair *PairCallerSession) Name() ( string,  error) {\n                  return _Pair.Contract.Name(&_Pair.CallOpts )\n                }\n\n                // Nonces is a free data retrieval call binding the contract method 0x7ecebe00.\n                //\n                // Solidity: function nonces(address owner) view returns(uint256)\n                func (_Pair *PairCaller) Nonces(opts *bind.CallOpts , owner common.Address ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"nonces\" , owner)\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Nonces is a free data retrieval call binding the contract method 0x7ecebe00.\n                //\n                // Solidity: function nonces(address owner) view returns(uint256)\n                func (_Pair *PairSession) Nonces( owner common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.Nonces(&_Pair.CallOpts , owner)\n                }\n\n                // Nonces is a free data retrieval call binding the contract method 0x7ecebe00.\n                //\n                // Solidity: function nonces(address owner) view returns(uint256)\n                func (_Pair *PairCallerSession) Nonces( owner common.Address ) ( *big.Int,  error) {\n                  return _Pair.Contract.Nonces(&_Pair.CallOpts , owner)\n                }\n\n                // Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.\n                //\n                // Solidity: function price0CumulativeLast() view returns(uint256)\n                func (_Pair *PairCaller) Price0CumulativeLast(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"price0CumulativeLast\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.\n                //\n                // Solidity: function price0CumulativeLast() view returns(uint256)\n                func (_Pair *PairSession) Price0CumulativeLast() ( *big.Int,  error) {\n                  return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts )\n                }\n\n                // Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.\n                //\n                // Solidity: function price0CumulativeLast() view returns(uint256)\n                func (_Pair *PairCallerSession) Price0CumulativeLast() ( *big.Int,  error) {\n                  return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts )\n                }\n\n                // Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.\n                //\n                // Solidity: function price1CumulativeLast() view returns(uint256)\n                func (_Pair *PairCaller) Price1CumulativeLast(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"price1CumulativeLast\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.\n                //\n                // Solidity: function price1CumulativeLast() view returns(uint256)\n                func (_Pair *PairSession) Price1CumulativeLast() ( *big.Int,  error) {\n                  return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts )\n                }\n\n                // Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.\n                //\n                // Solidity: function price1CumulativeLast() view returns(uint256)\n                func (_Pair *PairCallerSession) Price1CumulativeLast() ( *big.Int,  error) {\n                  return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts )\n                }\n\n                // Symbol is a free data retrieval call binding the contract method 0x95d89b41.\n                //\n                // Solidity: function symbol() pure returns(string)\n                func (_Pair *PairCaller) Symbol(opts *bind.CallOpts ) (string, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"symbol\" )\n\n                        if err != nil {\n                                return *new(string),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(string)).(*string)\n\n                        return out0,  err\n\n                }\n\n                // Symbol is a free data retrieval call binding the contract method 0x95d89b41.\n                //\n                // Solidity: function symbol() pure returns(string)\n                func (_Pair *PairSession) Symbol() ( string,  error) {\n                  return _Pair.Contract.Symbol(&_Pair.CallOpts )\n                }\n\n                // Symbol is a free data retrieval call binding the contract method 0x95d89b41.\n                //\n                // Solidity: function symbol() pure returns(string)\n                func (_Pair *PairCallerSession) Symbol() ( string,  error) {\n                  return _Pair.Contract.Symbol(&_Pair.CallOpts )\n                }\n\n                // Token0 is a free data retrieval call binding the contract method 0x0dfe1681.\n                //\n                // Solidity: function token0() view returns(address)\n                func (_Pair *PairCaller) Token0(opts *bind.CallOpts ) (common.Address, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"token0\" )\n\n                        if err != nil {\n                                return *new(common.Address),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)\n\n                        return out0,  err\n\n                }\n\n                // Token0 is a free data retrieval call binding the contract method 0x0dfe1681.\n                //\n                // Solidity: function token0() view returns(address)\n                func (_Pair *PairSession) Token0() ( common.Address,  error) {\n                  return _Pair.Contract.Token0(&_Pair.CallOpts )\n                }\n\n                // Token0 is a free data retrieval call binding the contract method 0x0dfe1681.\n                //\n                // Solidity: function token0() view returns(address)\n                func (_Pair *PairCallerSession) Token0() ( common.Address,  error) {\n                  return _Pair.Contract.Token0(&_Pair.CallOpts )\n                }\n\n                // Token1 is a free data retrieval call binding the contract method 0xd21220a7.\n                //\n                // Solidity: function token1() view returns(address)\n                func (_Pair *PairCaller) Token1(opts *bind.CallOpts ) (common.Address, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"token1\" )\n\n                        if err != nil {\n                                return *new(common.Address),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)\n\n                        return out0,  err\n\n                }\n\n                // Token1 is a free data retrieval call binding the contract method 0xd21220a7.\n                //\n                // Solidity: function token1() view returns(address)\n                func (_Pair *PairSession) Token1() ( common.Address,  error) {\n                  return _Pair.Contract.Token1(&_Pair.CallOpts )\n                }\n\n                // Token1 is a free data retrieval call binding the contract method 0xd21220a7.\n                //\n                // Solidity: function token1() view returns(address)\n                func (_Pair *PairCallerSession) Token1() ( common.Address,  error) {\n                  return _Pair.Contract.Token1(&_Pair.CallOpts )\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Pair *PairCaller) TotalSupply(opts *bind.CallOpts ) (*big.Int, error) {\n                        var out []interface{}\n                        err := _Pair.contract.Call(opts, &out, \"totalSupply\" )\n\n                        if err != nil {\n                                return *new(*big.Int),  err\n                        }\n\n                        out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)\n\n                        return out0,  err\n\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Pair *PairSession) TotalSupply() ( *big.Int,  error) {\n                  return _Pair.Contract.TotalSupply(&_Pair.CallOpts )\n                }\n\n                // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.\n                //\n                // Solidity: function totalSupply() view returns(uint256)\n                func (_Pair *PairCallerSession) TotalSupply() ( *big.Int,  error) {\n                  return _Pair.Contract.TotalSupply(&_Pair.CallOpts )\n                }\n\n\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 value) returns(bool)\n                func (_Pair *PairTransactor) Approve(opts *bind.TransactOpts , spender common.Address , value *big.Int ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"approve\" , spender, value)\n                }\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 value) returns(bool)\n                func (_Pair *PairSession) Approve( spender common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.Approve(&_Pair.TransactOpts , spender, value)\n                }\n\n                // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.\n                //\n                // Solidity: function approve(address spender, uint256 value) returns(bool)\n                func (_Pair *PairTransactorSession) Approve( spender common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.Approve(&_Pair.TransactOpts , spender, value)\n                }\n\n                // Burn is a paid mutator transaction binding the contract method 0x89afcb44.\n                //\n                // Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)\n                func (_Pair *PairTransactor) Burn(opts *bind.TransactOpts , to common.Address ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"burn\" , to)\n                }\n\n                // Burn is a paid mutator transaction binding the contract method 0x89afcb44.\n                //\n                // Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)\n                func (_Pair *PairSession) Burn( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Burn(&_Pair.TransactOpts , to)\n                }\n\n                // Burn is a paid mutator transaction binding the contract method 0x89afcb44.\n                //\n                // Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)\n                func (_Pair *PairTransactorSession) Burn( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Burn(&_Pair.TransactOpts , to)\n                }\n\n                // Initialize is a paid mutator transaction binding the contract method 0x485cc955.\n                //\n                // Solidity: function initialize(address , address ) returns()\n                func (_Pair *PairTransactor) Initialize(opts *bind.TransactOpts , arg0 common.Address , arg1 common.Address ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"initialize\" , arg0, arg1)\n                }\n\n                // Initialize is a paid mutator transaction binding the contract method 0x485cc955.\n                //\n                // Solidity: function initialize(address , address ) returns()\n                func (_Pair *PairSession) Initialize( arg0 common.Address , arg1 common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Initialize(&_Pair.TransactOpts , arg0, arg1)\n                }\n\n                // Initialize is a paid mutator transaction binding the contract method 0x485cc955.\n                //\n                // Solidity: function initialize(address , address ) returns()\n                func (_Pair *PairTransactorSession) Initialize( arg0 common.Address , arg1 common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Initialize(&_Pair.TransactOpts , arg0, arg1)\n                }\n\n                // Mint is a paid mutator transaction binding the contract method 0x6a627842.\n                //\n                // Solidity: function mint(address to) returns(uint256 liquidity)\n                func (_Pair *PairTransactor) Mint(opts *bind.TransactOpts , to common.Address ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"mint\" , to)\n                }\n\n                // Mint is a paid mutator transaction binding the contract method 0x6a627842.\n                //\n                // Solidity: function mint(address to) returns(uint256 liquidity)\n                func (_Pair *PairSession) Mint( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Mint(&_Pair.TransactOpts , to)\n                }\n\n                // Mint is a paid mutator transaction binding the contract method 0x6a627842.\n                //\n                // Solidity: function mint(address to) returns(uint256 liquidity)\n                func (_Pair *PairTransactorSession) Mint( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Mint(&_Pair.TransactOpts , to)\n                }\n\n                // Permit is a paid mutator transaction binding the contract method 0xd505accf.\n                //\n                // Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()\n                func (_Pair *PairTransactor) Permit(opts *bind.TransactOpts , owner common.Address , spender common.Address , value *big.Int , deadline *big.Int , v uint8 , r [32]byte , s [32]byte ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"permit\" , owner, spender, value, deadline, v, r, s)\n                }\n\n                // Permit is a paid mutator transaction binding the contract method 0xd505accf.\n                //\n                // Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()\n                func (_Pair *PairSession) Permit( owner common.Address , spender common.Address , value *big.Int , deadline *big.Int , v uint8 , r [32]byte , s [32]byte ) (*types.Transaction, error) {\n                  return _Pair.Contract.Permit(&_Pair.TransactOpts , owner, spender, value, deadline, v, r, s)\n                }\n\n                // Permit is a paid mutator transaction binding the contract method 0xd505accf.\n                //\n                // Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()\n                func (_Pair *PairTransactorSession) Permit( owner common.Address , spender common.Address , value *big.Int , deadline *big.Int , v uint8 , r [32]byte , s [32]byte ) (*types.Transaction, error) {\n                  return _Pair.Contract.Permit(&_Pair.TransactOpts , owner, spender, value, deadline, v, r, s)\n                }\n\n                // Skim is a paid mutator transaction binding the contract method 0xbc25cf77.\n                //\n                // Solidity: function skim(address to) returns()\n                func (_Pair *PairTransactor) Skim(opts *bind.TransactOpts , to common.Address ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"skim\" , to)\n                }\n\n                // Skim is a paid mutator transaction binding the contract method 0xbc25cf77.\n                //\n                // Solidity: function skim(address to) returns()\n                func (_Pair *PairSession) Skim( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Skim(&_Pair.TransactOpts , to)\n                }\n\n                // Skim is a paid mutator transaction binding the contract method 0xbc25cf77.\n                //\n                // Solidity: function skim(address to) returns()\n                func (_Pair *PairTransactorSession) Skim( to common.Address ) (*types.Transaction, error) {\n                  return _Pair.Contract.Skim(&_Pair.TransactOpts , to)\n                }\n\n                // Swap is a paid mutator transaction binding the contract method 0x022c0d9f.\n                //\n                // Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()\n                func (_Pair *PairTransactor) Swap(opts *bind.TransactOpts , amount0Out *big.Int , amount1Out *big.Int , to common.Address , data []byte ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"swap\" , amount0Out, amount1Out, to, data)\n                }\n\n                // Swap is a paid mutator transaction binding the contract method 0x022c0d9f.\n                //\n                // Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()\n                func (_Pair *PairSession) Swap( amount0Out *big.Int , amount1Out *big.Int , to common.Address , data []byte ) (*types.Transaction, error) {\n                  return _Pair.Contract.Swap(&_Pair.TransactOpts , amount0Out, amount1Out, to, data)\n                }\n\n                // Swap is a paid mutator transaction binding the contract method 0x022c0d9f.\n                //\n                // Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()\n                func (_Pair *PairTransactorSession) Swap( amount0Out *big.Int , amount1Out *big.Int , to common.Address , data []byte ) (*types.Transaction, error) {\n                  return _Pair.Contract.Swap(&_Pair.TransactOpts , amount0Out, amount1Out, to, data)\n                }\n\n                // Sync is a paid mutator transaction binding the contract method 0xfff6cae9.\n                //\n                // Solidity: function sync() returns()\n                func (_Pair *PairTransactor) Sync(opts *bind.TransactOpts ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"sync\" )\n                }\n\n                // Sync is a paid mutator transaction binding the contract method 0xfff6cae9.\n                //\n                // Solidity: function sync() returns()\n                func (_Pair *PairSession) Sync() (*types.Transaction, error) {\n                  return _Pair.Contract.Sync(&_Pair.TransactOpts )\n                }\n\n                // Sync is a paid mutator transaction binding the contract method 0xfff6cae9.\n                //\n                // Solidity: function sync() returns()\n                func (_Pair *PairTransactorSession) Sync() (*types.Transaction, error) {\n                  return _Pair.Contract.Sync(&_Pair.TransactOpts )\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address to, uint256 value) returns(bool)\n                func (_Pair *PairTransactor) Transfer(opts *bind.TransactOpts , to common.Address , value *big.Int ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"transfer\" , to, value)\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address to, uint256 value) returns(bool)\n                func (_Pair *PairSession) Transfer( to common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.Transfer(&_Pair.TransactOpts , to, value)\n                }\n\n                // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.\n                //\n                // Solidity: function transfer(address to, uint256 value) returns(bool)\n                func (_Pair *PairTransactorSession) Transfer( to common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.Transfer(&_Pair.TransactOpts , to, value)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)\n                func (_Pair *PairTransactor) TransferFrom(opts *bind.TransactOpts , from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {\n                        return _Pair.contract.Transact(opts, \"transferFrom\" , from, to, value)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)\n                func (_Pair *PairSession) TransferFrom( from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.TransferFrom(&_Pair.TransactOpts , from, to, value)\n                }\n\n                // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.\n                //\n                // Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)\n                func (_Pair *PairTransactorSession) TransferFrom( from common.Address , to common.Address , value *big.Int ) (*types.Transaction, error) {\n                  return _Pair.Contract.TransferFrom(&_Pair.TransactOpts , from, to, value)\n                }\n\n\n\n\n\n\n\n                // PairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pair contract.\n                type PairApprovalIterator struct {\n                        Event *PairApproval // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairApprovalIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairApproval)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairApproval)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairApprovalIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairApprovalIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairApproval represents a Approval event raised by the Pair contract.\n                type PairApproval struct {\n                        Owner common.Address;\n                        Spender common.Address;\n                        Value *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Pair *PairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PairApprovalIterator, error) {\n\n                        var ownerRule []interface{}\n                        for _, ownerItem := range owner {\n                                ownerRule = append(ownerRule, ownerItem)\n                        }\n                        var spenderRule []interface{}\n                        for _, spenderItem := range spender {\n                                spenderRule = append(spenderRule, spenderItem)\n                        }\n\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Approval\", ownerRule, spenderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairApprovalIterator{contract: _Pair.contract, event: \"Approval\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Pair *PairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {\n\n                        var ownerRule []interface{}\n                        for _, ownerItem := range owner {\n                                ownerRule = append(ownerRule, ownerItem)\n                        }\n                        var spenderRule []interface{}\n                        for _, spenderItem := range spender {\n                                spenderRule = append(spenderRule, spenderItem)\n                        }\n\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Approval\", ownerRule, spenderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairApproval)\n                                                if err := _Pair.contract.UnpackLog(event, \"Approval\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.\n                //\n                // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)\n                func (_Pair *PairFilterer) ParseApproval(log types.Log) (*PairApproval, error) {\n                        event := new(PairApproval)\n                        if err := _Pair.contract.UnpackLog(event, \"Approval\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // PairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pair contract.\n                type PairBurnIterator struct {\n                        Event *PairBurn // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairBurnIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairBurn)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairBurn)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairBurnIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairBurnIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairBurn represents a Burn event raised by the Pair contract.\n                type PairBurn struct {\n                        Sender common.Address;\n                        Amount0 *big.Int;\n                        Amount1 *big.Int;\n                        To common.Address;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.\n                //\n                // Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)\n                func (_Pair *PairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairBurnIterator, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Burn\", senderRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairBurnIterator{contract: _Pair.contract, event: \"Burn\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.\n                //\n                // Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)\n                func (_Pair *PairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Burn\", senderRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairBurn)\n                                                if err := _Pair.contract.UnpackLog(event, \"Burn\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.\n                //\n                // Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)\n                func (_Pair *PairFilterer) ParseBurn(log types.Log) (*PairBurn, error) {\n                        event := new(PairBurn)\n                        if err := _Pair.contract.UnpackLog(event, \"Burn\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // PairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pair contract.\n                type PairMintIterator struct {\n                        Event *PairMint // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairMintIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairMint)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairMint)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairMintIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairMintIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairMint represents a Mint event raised by the Pair contract.\n                type PairMint struct {\n                        Sender common.Address;\n                        Amount0 *big.Int;\n                        Amount1 *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.\n                //\n                // Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)\n                func (_Pair *PairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*PairMintIterator, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Mint\", senderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairMintIterator{contract: _Pair.contract, event: \"Mint\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.\n                //\n                // Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)\n                func (_Pair *PairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PairMint, sender []common.Address) (event.Subscription, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Mint\", senderRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairMint)\n                                                if err := _Pair.contract.UnpackLog(event, \"Mint\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.\n                //\n                // Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)\n                func (_Pair *PairFilterer) ParseMint(log types.Log) (*PairMint, error) {\n                        event := new(PairMint)\n                        if err := _Pair.contract.UnpackLog(event, \"Mint\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // PairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pair contract.\n                type PairSwapIterator struct {\n                        Event *PairSwap // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairSwapIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairSwap)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairSwap)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairSwapIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairSwapIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairSwap represents a Swap event raised by the Pair contract.\n                type PairSwap struct {\n                        Sender common.Address;\n                        Amount0In *big.Int;\n                        Amount1In *big.Int;\n                        Amount0Out *big.Int;\n                        Amount1Out *big.Int;\n                        To common.Address;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.\n                //\n                // Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)\n                func (_Pair *PairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairSwapIterator, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n\n\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Swap\", senderRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairSwapIterator{contract: _Pair.contract, event: \"Swap\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.\n                //\n                // Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)\n                func (_Pair *PairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {\n\n                        var senderRule []interface{}\n                        for _, senderItem := range sender {\n                                senderRule = append(senderRule, senderItem)\n                        }\n\n\n\n\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Swap\", senderRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairSwap)\n                                                if err := _Pair.contract.UnpackLog(event, \"Swap\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.\n                //\n                // Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)\n                func (_Pair *PairFilterer) ParseSwap(log types.Log) (*PairSwap, error) {\n                        event := new(PairSwap)\n                        if err := _Pair.contract.UnpackLog(event, \"Swap\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // PairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pair contract.\n                type PairSyncIterator struct {\n                        Event *PairSync // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairSyncIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairSync)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairSync)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairSyncIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairSyncIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairSync represents a Sync event raised by the Pair contract.\n                type PairSync struct {\n                        Reserve0 *big.Int;\n                        Reserve1 *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.\n                //\n                // Solidity: event Sync(uint112 reserve0, uint112 reserve1)\n                func (_Pair *PairFilterer) FilterSync(opts *bind.FilterOpts) (*PairSyncIterator, error) {\n\n\n\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Sync\")\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairSyncIterator{contract: _Pair.contract, event: \"Sync\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.\n                //\n                // Solidity: event Sync(uint112 reserve0, uint112 reserve1)\n                func (_Pair *PairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PairSync) (event.Subscription, error) {\n\n\n\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Sync\")\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairSync)\n                                                if err := _Pair.contract.UnpackLog(event, \"Sync\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.\n                //\n                // Solidity: event Sync(uint112 reserve0, uint112 reserve1)\n                func (_Pair *PairFilterer) ParseSync(log types.Log) (*PairSync, error) {\n                        event := new(PairSync)\n                        if err := _Pair.contract.UnpackLog(event, \"Sync\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }\n\n\n                // PairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pair contract.\n                type PairTransferIterator struct {\n                        Event *PairTransfer // Event containing the contract specifics and raw log\n\n                        contract *bind.BoundContract // Generic contract to use for unpacking event data\n                        event    string              // Event name to use for unpacking event data\n\n                        logs chan types.Log        // Log channel receiving the found contract events\n                        sub  ethereum.Subscription // Subscription for errors, completion and termination\n                        done bool                  // Whether the subscription completed delivering logs\n                        fail error                 // Occurred error to stop iteration\n                }\n                // Next advances the iterator to the subsequent event, returning whether there\n                // are any more events found. In case of a retrieval or parsing error, false is\n                // returned and Error() can be queried for the exact failure.\n                func (it *PairTransferIterator) Next() bool {\n                        // If the iterator failed, stop iterating\n                        if (it.fail != nil) {\n                                return false\n                        }\n                        // If the iterator completed, deliver directly whatever's available\n                        if (it.done) {\n                                select {\n                                case log := <-it.logs:\n                                        it.Event = new(PairTransfer)\n                                        if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                                it.fail = err\n                                                return false\n                                        }\n                                        it.Event.Raw = log\n                                        return true\n\n                                default:\n                                        return false\n                                }\n                        }\n                        // Iterator still in progress, wait for either a data or an error event\n                        select {\n                        case log := <-it.logs:\n                                it.Event = new(PairTransfer)\n                                if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {\n                                        it.fail = err\n                                        return false\n                                }\n                                it.Event.Raw = log\n                                return true\n\n                        case err := <-it.sub.Err():\n                                it.done = true\n                                it.fail = err\n                                return it.Next()\n                        }\n                }\n                // Error returns any retrieval or parsing error occurred during filtering.\n                func (it *PairTransferIterator) Error() error {\n                        return it.fail\n                }\n                // Close terminates the iteration process, releasing any pending underlying\n                // resources.\n                func (it *PairTransferIterator) Close() error {\n                        it.sub.Unsubscribe()\n                        return nil\n                }\n\n                // PairTransfer represents a Transfer event raised by the Pair contract.\n                type PairTransfer struct {\n                        From common.Address;\n                        To common.Address;\n                        Value *big.Int;\n                        Raw types.Log // Blockchain specific contextual infos\n                }\n\n                // FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Pair *PairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PairTransferIterator, error) {\n\n                        var fromRule []interface{}\n                        for _, fromItem := range from {\n                                fromRule = append(fromRule, fromItem)\n                        }\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n\n                        logs, sub, err := _Pair.contract.FilterLogs(opts, \"Transfer\", fromRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return &PairTransferIterator{contract: _Pair.contract, event: \"Transfer\", logs: logs, sub: sub}, nil\n                }\n\n                // WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Pair *PairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {\n\n                        var fromRule []interface{}\n                        for _, fromItem := range from {\n                                fromRule = append(fromRule, fromItem)\n                        }\n                        var toRule []interface{}\n                        for _, toItem := range to {\n                                toRule = append(toRule, toItem)\n                        }\n\n\n                        logs, sub, err := _Pair.contract.WatchLogs(opts, \"Transfer\", fromRule, toRule)\n                        if err != nil {\n                                return nil, err\n                        }\n                        return event.NewSubscription(func(quit <-chan struct{}) error {\n                                defer sub.Unsubscribe()\n                                for {\n                                        select {\n                                        case log := <-logs:\n                                                // New log arrived, parse the event and forward to the user\n                                                event := new(PairTransfer)\n                                                if err := _Pair.contract.UnpackLog(event, \"Transfer\", log); err != nil {\n                                                        return err\n                                                }\n                                                event.Raw = log\n\n                                                select {\n                                                case sink <- event:\n                                                case err := <-sub.Err():\n                                                        return err\n                                                case <-quit:\n                                                        return nil\n                                                }\n                                        case err := <-sub.Err():\n                                                return err\n                                        case <-quit:\n                                                return nil\n                                        }\n                                }\n                        }), nil\n                }\n\n                // ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.\n                //\n                // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)\n                func (_Pair *PairFilterer) ParseTransfer(log types.Log) (*PairTransfer, error) {\n                        event := new(PairTransfer)\n                        if err := _Pair.contract.UnpackLog(event, \"Transfer\", log); err != nil {\n                                return nil, err\n                        }\n                        event.Raw = log\n                        return event, nil\n                }"

// DeployPair deploys a new Ethereum contract, binding an instance of Pair to it.
func DeployPair(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pair, error) {
	parsed, err := abi.JSON(strings.NewReader(PairABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PairBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pair{PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract}}, nil
}

// Pair is an auto generated Go binding around an Ethereum contract.
type Pair struct {
	PairCaller     // Read-only binding to the contract
	PairTransactor // Write-only binding to the contract
	PairFilterer   // Log filterer for contract events
}

// PairCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairSession struct {
	Contract     *Pair             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairCallerSession struct {
	Contract *PairCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairTransactorSession struct {
	Contract     *PairTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairRaw struct {
	Contract *Pair // Generic contract binding to access the raw methods on
}

// PairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairCallerRaw struct {
	Contract *PairCaller // Generic read-only contract binding to access the raw methods on
}

// PairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairTransactorRaw struct {
	Contract *PairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPair creates a new instance of Pair, bound to a specific deployed contract.
func NewPair(address common.Address, backend bind.ContractBackend) (*Pair, error) {
	contract, err := bindPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pair{PairCaller: PairCaller{contract: contract}, PairTransactor: PairTransactor{contract: contract}, PairFilterer: PairFilterer{contract: contract}}, nil
}

// NewPairCaller creates a new read-only instance of Pair, bound to a specific deployed contract.
func NewPairCaller(address common.Address, caller bind.ContractCaller) (*PairCaller, error) {
	contract, err := bindPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairCaller{contract: contract}, nil
}

// NewPairTransactor creates a new write-only instance of Pair, bound to a specific deployed contract.
func NewPairTransactor(address common.Address, transactor bind.ContractTransactor) (*PairTransactor, error) {
	contract, err := bindPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairTransactor{contract: contract}, nil
}

// NewPairFilterer creates a new log filterer instance of Pair, bound to a specific deployed contract.
func NewPairFilterer(address common.Address, filterer bind.ContractFilterer) (*PairFilterer, error) {
	contract, err := bindPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairFilterer{contract: contract}, nil
}

// bindPair binds a generic wrapper to an already deployed contract.
func bindPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.PairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.PairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pair *PairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pair *PairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pair *PairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pair.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pair *PairCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pair.Contract.DOMAINSEPARATOR(&_Pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Pair *PairCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Pair *PairSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() pure returns(uint256)
func (_Pair *PairCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pair.Contract.MINIMUMLIQUIDITY(&_Pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Pair *PairCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Pair *PairSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() pure returns(bytes32)
func (_Pair *PairCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pair.Contract.PERMITTYPEHASH(&_Pair.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pair *PairCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pair *PairSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pair.Contract.Allowance(&_Pair.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Pair *PairCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Pair.Contract.Allowance(&_Pair.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pair *PairCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pair *PairSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Pair *PairCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Pair.Contract.BalanceOf(&_Pair.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Pair *PairCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Pair *PairSession) Decimals() (uint8, error) {
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_Pair *PairCallerSession) Decimals() (uint8, error) {
	return _Pair.Contract.Decimals(&_Pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairSession) Factory() (common.Address, error) {
	return _Pair.Contract.Factory(&_Pair.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pair *PairCallerSession) Factory() (common.Address, error) {
	return _Pair.Contract.Factory(&_Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Pair *PairCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Pair *PairSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast)
func (_Pair *PairCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pair.Contract.GetReserves(&_Pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairSession) KLast() (*big.Int, error) {
	return _Pair.Contract.KLast(&_Pair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pair *PairCallerSession) KLast() (*big.Int, error) {
	return _Pair.Contract.KLast(&_Pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pair *PairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pair *PairSession) Name() (string, error) {
	return _Pair.Contract.Name(&_Pair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_Pair *PairCallerSession) Name() (string, error) {
	return _Pair.Contract.Name(&_Pair.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pair *PairCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pair *PairSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Pair.Contract.Nonces(&_Pair.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Pair *PairCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Pair.Contract.Nonces(&_Pair.CallOpts, owner)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Price0CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Price0CumulativeLast(&_Pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairSession) Price1CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pair *PairCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Pair.Contract.Price1CumulativeLast(&_Pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Pair *PairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Pair *PairSession) Symbol() (string, error) {
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_Pair *PairCallerSession) Symbol() (string, error) {
	return _Pair.Contract.Symbol(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pair *PairCallerSession) Token0() (common.Address, error) {
	return _Pair.Contract.Token0(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pair *PairCallerSession) Token1() (common.Address, error) {
	return _Pair.Contract.Token1(&_Pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pair.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairSession) TotalSupply() (*big.Int, error) {
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pair *PairCallerSession) TotalSupply() (*big.Int, error) {
	return _Pair.Contract.TotalSupply(&_Pair.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Approve(&_Pair.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pair *PairTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Burn(&_Pair.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Pair *PairTransactor) Initialize(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "initialize", arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Pair *PairSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Initialize(&_Pair.TransactOpts, arg0, arg1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address , address ) returns()
func (_Pair *PairTransactorSession) Initialize(arg0 common.Address, arg1 common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Initialize(&_Pair.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pair *PairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Mint(&_Pair.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pair *PairTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pair.Contract.Permit(&_Pair.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pair *PairTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pair.Contract.Skim(&_Pair.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pair *PairTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pair.Contract.Swap(&_Pair.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairSession) Sync() (*types.Transaction, error) {
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pair *PairTransactorSession) Sync() (*types.Transaction, error) {
	return _Pair.Contract.Sync(&_Pair.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.Transfer(&_Pair.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pair *PairTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pair.Contract.TransferFrom(&_Pair.TransactOpts, from, to, value)
}

// PairApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pair contract.
type PairApprovalIterator struct {
	Event *PairApproval // Event containing the contract specifics and raw log

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
func (it *PairApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairApproval)
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
		it.Event = new(PairApproval)
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
func (it *PairApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairApproval represents a Approval event raised by the Pair contract.
type PairApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pair *PairFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PairApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PairApprovalIterator{contract: _Pair.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pair *PairFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PairApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairApproval)
				if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Pair *PairFilterer) ParseApproval(log types.Log) (*PairApproval, error) {
	event := new(PairApproval)
	if err := _Pair.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pair contract.
type PairBurnIterator struct {
	Event *PairBurn // Event containing the contract specifics and raw log

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
func (it *PairBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairBurn)
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
		it.Event = new(PairBurn)
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
func (it *PairBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairBurn represents a Burn event raised by the Pair contract.
type PairBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairBurnIterator{contract: _Pair.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PairBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairBurn)
				if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pair *PairFilterer) ParseBurn(log types.Log) (*PairBurn, error) {
	event := new(PairBurn)
	if err := _Pair.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pair contract.
type PairMintIterator struct {
	Event *PairMint // Event containing the contract specifics and raw log

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
func (it *PairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairMint)
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
		it.Event = new(PairMint)
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
func (it *PairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairMint represents a Mint event raised by the Pair contract.
type PairMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*PairMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &PairMintIterator{contract: _Pair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PairMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairMint)
				if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pair *PairFilterer) ParseMint(log types.Log) (*PairMint, error) {
	event := new(PairMint)
	if err := _Pair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pair contract.
type PairSwapIterator struct {
	Event *PairSwap // Event containing the contract specifics and raw log

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
func (it *PairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairSwap)
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
		it.Event = new(PairSwap)
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
func (it *PairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairSwap represents a Swap event raised by the Pair contract.
type PairSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairSwapIterator{contract: _Pair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSwap)
				if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pair *PairFilterer) ParseSwap(log types.Log) (*PairSwap, error) {
	event := new(PairSwap)
	if err := _Pair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pair contract.
type PairSyncIterator struct {
	Event *PairSync // Event containing the contract specifics and raw log

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
func (it *PairSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairSync)
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
		it.Event = new(PairSync)
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
func (it *PairSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairSync represents a Sync event raised by the Pair contract.
type PairSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) FilterSync(opts *bind.FilterOpts) (*PairSyncIterator, error) {

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &PairSyncIterator{contract: _Pair.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PairSync) (event.Subscription, error) {

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairSync)
				if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pair *PairFilterer) ParseSync(log types.Log) (*PairSync, error) {
	event := new(PairSync)
	if err := _Pair.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PairTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pair contract.
type PairTransferIterator struct {
	Event *PairTransfer // Event containing the contract specifics and raw log

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
func (it *PairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairTransfer)
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
		it.Event = new(PairTransfer)
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
func (it *PairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairTransfer represents a Transfer event raised by the Pair contract.
type PairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pair *PairFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairTransferIterator{contract: _Pair.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pair *PairFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pair.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairTransfer)
				if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Pair *PairFilterer) ParseTransfer(log types.Log) (*PairTransfer, error) {
	event := new(PairTransfer)
	if err := _Pair.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
