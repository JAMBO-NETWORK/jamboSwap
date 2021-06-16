// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mortgage

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

// MortgageABI is the input ABI used to generate the binding from.
const MortgageABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mortgage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYam\",\"name\":\"_yam\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newAmount\",\"type\":\"uint256\"}],\"name\":\"setAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYam\",\"name\":\"_yam\",\"type\":\"address\"}],\"name\":\"setYam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAddrsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yam\",\"outputs\":[{\"internalType\":\"contractIYam\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MortgageBin is the compiled bytecode used for deploying new contracts.
var MortgageBin = "0x{\n\t\"generatedSources\": [\n\t\t{\n\t\t\t\"ast\": {\n\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\"src\": \"0:339:8\",\n\t\t\t\t\"statements\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\"src\": \"6:3:8\",\n\t\t\t\t\t\t\"statements\": []\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\"src\": \"108:229:8\",\n\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"154:26:8\",\n\t\t\t\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"expression\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"163:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"171:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"revert\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulExpressionStatement\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"condition\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"dataEnd\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"129:7:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"138:9:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"sub\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"125:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"125:23:8\"\n\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"150:2:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"32\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"slt\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"121:3:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"121:32:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIf\",\n\t\t\t\t\t\t\t\t\t\"src\": \"118:2:8\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulVariableDeclaration\",\n\t\t\t\t\t\t\t\t\t\"src\": \"189:29:8\",\n\t\t\t\t\t\t\t\t\t\"value\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"208:9:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"mload\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"202:5:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"202:16:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"variables\": [\n\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"193:5:8\",\n\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"281:26:8\",\n\t\t\t\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"expression\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"290:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"298:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"revert\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulExpressionStatement\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"condition\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"240:5:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"251:5:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"266:3:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"160\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"271:1:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"1\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"shl\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"262:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"262:11:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"275:1:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"1\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"sub\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"258:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"258:19:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"and\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"247:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"247:31:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"eq\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"237:2:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"237:42:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"iszero\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"230:6:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"230:50:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIf\",\n\t\t\t\t\t\t\t\t\t\"src\": \"227:2:8\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulAssignment\",\n\t\t\t\t\t\t\t\t\t\"src\": \"316:15:8\",\n\t\t\t\t\t\t\t\t\t\"value\": {\n\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"326:5:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"variableNames\": [\n\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"316:6:8\"\n\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t\"name\": \"abi_decode_tuple_t_contract$_IYam_$1057_fromMemory\",\n\t\t\t\t\t\t\"nodeType\": \"YulFunctionDefinition\",\n\t\t\t\t\t\t\"parameters\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"74:9:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"dataEnd\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"85:7:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t],\n\t\t\t\t\t\t\"returnVariables\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"97:6:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t],\n\t\t\t\t\t\t\"src\": \"14:323:8\"\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t},\n\t\t\t\"contents\": \"{\\n    { }\\n    function abi_decode_tuple_t_contract$_IYam_$1057_fromMemory(headStart, dataEnd) -> value0\\n    {\\n        if slt(sub(dataEnd, headStart), 32) { revert(value0, value0) }\\n        let value := mload(headStart)\\n        if iszero(eq(value, and(value, sub(shl(160, 1), 1)))) { revert(value0, value0) }\\n        value0 := value\\n    }\\n}\",\n\t\t\t\"id\": 8,\n\t\t\t\"language\": \"Yul\",\n\t\t\t\"name\": \"#utility.yul\"\n\t\t}\n\t],\n\t\"linkReferences\": {},\n\t\"object\": \"6080604052600a60015534801561001557600080fd5b50604051610c63380380610c63833981016040819052610034916100b1565b600061003e6100ad565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600280546001600160a01b0319166001600160a01b03929092169190911790556100df565b3390565b6000602082840312156100c2578081fd5b81516001600160a01b03811681146100d8578182fd5b9392505050565b610b75806100ee6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b1461013d5780639de7e36314610145578063a7f4377914610158578063aa8c217c14610160578063e490240014610168578063f2fde38b1461017b576100b4565b8063071c0332146100b95780631be3175e146100d7578063271f88b4146100ec5780632bcc23f8146100ff578063715018a6146101205780637db083dc14610128575b600080fd5b6100c161018e565b6040516100ce9190610871565b60405180910390f35b6100ea6100e536600461083d565b61019d565b005b6100ea6100fa36600461083d565b6102ed565b61011261010d366004610801565b610331565b6040516100ce9291906108a9565b6100ea610356565b6101306103df565b6040516100ce9190610aa5565b6100c16103e5565b6100ea610153366004610801565b6103f4565b6100ea610455565b6101306104c7565b6100c161017636600461083d565b6104cd565b6100ea610189366004610801565b6104f7565b6002546001600160a01b031681565b6001546101b290670de0b6b3a76400006105b7565b8110156101da5760405162461bcd60e51b81526004016101d190610981565b60405180910390fd5b60006101e46105ca565b6002549091506101ff906001600160a01b03168230856105ce565b6001600160a01b03818116600090815260046020526040902054166102a6576040805180820182526001600160a01b038084168083526020808401878152600083815260049092529481209351845493166001600160a01b031993841617845593516001938401556003805493840181559093527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101805490911690911790556102e9565b6001600160a01b0381166000908152600460205260409020600101546102cc908361062c565b6001600160a01b0382166000908152600460205260409020600101555b5050565b6102f56105ca565b6001600160a01b03166103066103e5565b6001600160a01b03161461032c5760405162461bcd60e51b81526004016101d1906109ef565b600155565b600460205260009081526040902080546001909101546001600160a01b039091169082565b61035e6105ca565b6001600160a01b031661036f6103e5565b6001600160a01b0316146103955760405162461bcd60e51b81526004016101d1906109ef565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60035490565b6000546001600160a01b031690565b6103fc6105ca565b6001600160a01b031661040d6103e5565b6001600160a01b0316146104335760405162461bcd60e51b81526004016101d1906109ef565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b600061045f6105ca565b6001600160a01b0381166000908152600460205260409020600101549091508061049b5760405162461bcd60e51b81526004016101d1906109b8565b6001600160a01b038083166000908152600460205260408120600101556002546102e991168383610638565b60015481565b600381815481106104dd57600080fd5b6000918252602090912001546001600160a01b0316905081565b6104ff6105ca565b6001600160a01b03166105106103e5565b6001600160a01b0316146105365760405162461bcd60e51b81526004016101d1906109ef565b6001600160a01b03811661055c5760405162461bcd60e51b81526004016101d1906108f5565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b60006105c38284610ac6565b9392505050565b3390565b610626846323b872dd60e01b8585856040516024016105ef93929190610885565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261065c565b50505050565b60006105c38284610aae565b6106578363a9059cbb60e01b84846040516024016105ef9291906108a9565b505050565b60006106b1826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166106eb9092919063ffffffff16565b80519091501561065757808060200190518101906106cf919061081d565b6106575760405162461bcd60e51b81526004016101d190610a5b565b60606106fa8484600085610702565b949350505050565b6060824710156107245760405162461bcd60e51b81526004016101d19061093b565b61072d856107c2565b6107495760405162461bcd60e51b81526004016101d190610a24565b600080866001600160a01b031685876040516107659190610855565b60006040518083038185875af1925050503d80600081146107a2576040519150601f19603f3d011682016040523d82523d6000602084013e6107a7565b606091505b50915091506107b78282866107c8565b979650505050505050565b3b151590565b606083156107d75750816105c3565b8251156107e75782518084602001fd5b8160405162461bcd60e51b81526004016101d191906108c2565b600060208284031215610812578081fd5b81356105c381610b27565b60006020828403121561082e578081fd5b815180151581146105c3578182fd5b60006020828403121561084e578081fd5b5035919050565b60008251610867818460208701610ae5565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b60006020825282518060208401526108e1816040850160208701610ae5565b601f01601f19169190910160400192915050565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b6020808252601a908201527f4e6f7420656e6f756768206d6f72746761676520616d6f756e74000000000000604082015260600190565b6020808252601c908201527f72656d6f76653a20496e73756666696369656e742062616c616e636500000000604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b90815260200190565b60008219821115610ac157610ac1610b11565b500190565b6000816000190483118215151615610ae057610ae0610b11565b500290565b60005b83811015610b00578181015183820152602001610ae8565b838111156106265750506000910152565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b0381168114610b3c57600080fd5b5056fea26469706673582212201bd9a9ff51568a35b2ecb6bdceabb6c3fac715d710b1a49a45d2668635cc4b0d64736f6c63430008010033\",\n\t\"opcodes\": \"PUSH1 0x80 PUSH1 0x40 MSTORE PUSH1 0xA PUSH1 0x1 SSTORE CALLVALUE DUP1 ISZERO PUSH2 0x15 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 MLOAD PUSH2 0xC63 CODESIZE SUB DUP1 PUSH2 0xC63 DUP4 CODECOPY DUP2 ADD PUSH1 0x40 DUP2 SWAP1 MSTORE PUSH2 0x34 SWAP2 PUSH2 0xB1 JUMP JUMPDEST PUSH1 0x0 PUSH2 0x3E PUSH2 0xAD JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP4 AND SWAP1 DUP2 OR DUP3 SSTORE PUSH1 0x40 MLOAD SWAP3 SWAP4 POP SWAP2 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP1 DUP3 SWAP1 LOG3 POP PUSH1 0x2 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE PUSH2 0xDF JUMP JUMPDEST CALLER SWAP1 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0xC2 JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND DUP2 EQ PUSH2 0xD8 JUMPI DUP2 DUP3 REVERT JUMPDEST SWAP4 SWAP3 POP POP POP JUMP JUMPDEST PUSH2 0xB75 DUP1 PUSH2 0xEE PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN INVALID PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x4 CALLDATASIZE LT PUSH2 0xB4 JUMPI PUSH1 0x0 CALLDATALOAD PUSH1 0xE0 SHR DUP1 PUSH4 0x8DA5CB5B GT PUSH2 0x71 JUMPI DUP1 PUSH4 0x8DA5CB5B EQ PUSH2 0x13D JUMPI DUP1 PUSH4 0x9DE7E363 EQ PUSH2 0x145 JUMPI DUP1 PUSH4 0xA7F43779 EQ PUSH2 0x158 JUMPI DUP1 PUSH4 0xAA8C217C EQ PUSH2 0x160 JUMPI DUP1 PUSH4 0xE4902400 EQ PUSH2 0x168 JUMPI DUP1 PUSH4 0xF2FDE38B EQ PUSH2 0x17B JUMPI PUSH2 0xB4 JUMP JUMPDEST DUP1 PUSH4 0x71C0332 EQ PUSH2 0xB9 JUMPI DUP1 PUSH4 0x1BE3175E EQ PUSH2 0xD7 JUMPI DUP1 PUSH4 0x271F88B4 EQ PUSH2 0xEC JUMPI DUP1 PUSH4 0x2BCC23F8 EQ PUSH2 0xFF JUMPI DUP1 PUSH4 0x715018A6 EQ PUSH2 0x120 JUMPI DUP1 PUSH4 0x7DB083DC EQ PUSH2 0x128 JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0xC1 PUSH2 0x18E JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0xCE SWAP2 SWAP1 PUSH2 0x871 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST PUSH2 0xEA PUSH2 0xE5 CALLDATASIZE PUSH1 0x4 PUSH2 0x83D JUMP JUMPDEST PUSH2 0x19D JUMP JUMPDEST STOP JUMPDEST PUSH2 0xEA PUSH2 0xFA CALLDATASIZE PUSH1 0x4 PUSH2 0x83D JUMP JUMPDEST PUSH2 0x2ED JUMP JUMPDEST PUSH2 0x112 PUSH2 0x10D CALLDATASIZE PUSH1 0x4 PUSH2 0x801 JUMP JUMPDEST PUSH2 0x331 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0xCE SWAP3 SWAP2 SWAP1 PUSH2 0x8A9 JUMP JUMPDEST PUSH2 0xEA PUSH2 0x356 JUMP JUMPDEST PUSH2 0x130 PUSH2 0x3DF JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0xCE SWAP2 SWAP1 PUSH2 0xAA5 JUMP JUMPDEST PUSH2 0xC1 PUSH2 0x3E5 JUMP JUMPDEST PUSH2 0xEA PUSH2 0x153 CALLDATASIZE PUSH1 0x4 PUSH2 0x801 JUMP JUMPDEST PUSH2 0x3F4 JUMP JUMPDEST PUSH2 0xEA PUSH2 0x455 JUMP JUMPDEST PUSH2 0x130 PUSH2 0x4C7 JUMP JUMPDEST PUSH2 0xC1 PUSH2 0x176 CALLDATASIZE PUSH1 0x4 PUSH2 0x83D JUMP JUMPDEST PUSH2 0x4CD JUMP JUMPDEST PUSH2 0xEA PUSH2 0x189 CALLDATASIZE PUSH1 0x4 PUSH2 0x801 JUMP JUMPDEST PUSH2 0x4F7 JUMP JUMPDEST PUSH1 0x2 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP2 JUMP JUMPDEST PUSH1 0x1 SLOAD PUSH2 0x1B2 SWAP1 PUSH8 0xDE0B6B3A7640000 PUSH2 0x5B7 JUMP JUMPDEST DUP2 LT ISZERO PUSH2 0x1DA JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x981 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST PUSH1 0x0 PUSH2 0x1E4 PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x2 SLOAD SWAP1 SWAP2 POP PUSH2 0x1FF SWAP1 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP3 ADDRESS DUP6 PUSH2 0x5CE JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 DUP2 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SLOAD AND PUSH2 0x2A6 JUMPI PUSH1 0x40 DUP1 MLOAD DUP1 DUP3 ADD DUP3 MSTORE PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP5 AND DUP1 DUP4 MSTORE PUSH1 0x20 DUP1 DUP5 ADD DUP8 DUP2 MSTORE PUSH1 0x0 DUP4 DUP2 MSTORE PUSH1 0x4 SWAP1 SWAP3 MSTORE SWAP5 DUP2 KECCAK256 SWAP4 MLOAD DUP5 SLOAD SWAP4 AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT SWAP4 DUP5 AND OR DUP5 SSTORE SWAP4 MLOAD PUSH1 0x1 SWAP4 DUP5 ADD SSTORE PUSH1 0x3 DUP1 SLOAD SWAP4 DUP5 ADD DUP2 SSTORE SWAP1 SWAP4 MSTORE PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85B SWAP1 SWAP2 ADD DUP1 SLOAD SWAP1 SWAP2 AND SWAP1 SWAP2 OR SWAP1 SSTORE PUSH2 0x2E9 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 PUSH1 0x1 ADD SLOAD PUSH2 0x2CC SWAP1 DUP4 PUSH2 0x62C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP3 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 PUSH1 0x1 ADD SSTORE JUMPDEST POP POP JUMP JUMPDEST PUSH2 0x2F5 PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x306 PUSH2 0x3E5 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x32C JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x9EF JUMP JUMPDEST PUSH1 0x1 SSTORE JUMP JUMPDEST PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x40 SWAP1 KECCAK256 DUP1 SLOAD PUSH1 0x1 SWAP1 SWAP2 ADD SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 DUP3 JUMP JUMPDEST PUSH2 0x35E PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x36F PUSH2 0x3E5 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x395 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x9EF JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x40 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP1 DUP4 SWAP1 LOG3 PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x3 SLOAD SWAP1 JUMP JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND SWAP1 JUMP JUMPDEST PUSH2 0x3FC PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x40D PUSH2 0x3E5 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x433 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x9EF JUMP JUMPDEST PUSH1 0x2 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x0 PUSH2 0x45F PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 PUSH1 0x1 ADD SLOAD SWAP1 SWAP2 POP DUP1 PUSH2 0x49B JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x9B8 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP4 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x4 PUSH1 0x20 MSTORE PUSH1 0x40 DUP2 KECCAK256 PUSH1 0x1 ADD SSTORE PUSH1 0x2 SLOAD PUSH2 0x2E9 SWAP2 AND DUP4 DUP4 PUSH2 0x638 JUMP JUMPDEST PUSH1 0x1 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x3 DUP2 DUP2 SLOAD DUP2 LT PUSH2 0x4DD JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 SWAP1 SWAP2 KECCAK256 ADD SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND SWAP1 POP DUP2 JUMP JUMPDEST PUSH2 0x4FF PUSH2 0x5CA JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x510 PUSH2 0x3E5 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x536 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x9EF JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND PUSH2 0x55C JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x8F5 JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x40 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP6 AND SWAP4 SWAP3 AND SWAP2 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP2 LOG3 PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x0 PUSH2 0x5C3 DUP3 DUP5 PUSH2 0xAC6 JUMP JUMPDEST SWAP4 SWAP3 POP POP POP JUMP JUMPDEST CALLER SWAP1 JUMP JUMPDEST PUSH2 0x626 DUP5 PUSH4 0x23B872DD PUSH1 0xE0 SHL DUP6 DUP6 DUP6 PUSH1 0x40 MLOAD PUSH1 0x24 ADD PUSH2 0x5EF SWAP4 SWAP3 SWAP2 SWAP1 PUSH2 0x885 JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD PUSH1 0x1F NOT DUP2 DUP5 SUB ADD DUP2 MSTORE SWAP2 SWAP1 MSTORE PUSH1 0x20 DUP2 ADD DUP1 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xE0 SHL SUB AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xE0 SHL SUB NOT SWAP1 SWAP4 AND SWAP3 SWAP1 SWAP3 OR SWAP1 SWAP2 MSTORE PUSH2 0x65C JUMP JUMPDEST POP POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x5C3 DUP3 DUP5 PUSH2 0xAAE JUMP JUMPDEST PUSH2 0x657 DUP4 PUSH4 0xA9059CBB PUSH1 0xE0 SHL DUP5 DUP5 PUSH1 0x40 MLOAD PUSH1 0x24 ADD PUSH2 0x5EF SWAP3 SWAP2 SWAP1 PUSH2 0x8A9 JUMP JUMPDEST POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0x6B1 DUP3 PUSH1 0x40 MLOAD DUP1 PUSH1 0x40 ADD PUSH1 0x40 MSTORE DUP1 PUSH1 0x20 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x5361666545524332303A206C6F772D6C6576656C2063616C6C206661696C6564 DUP2 MSTORE POP DUP6 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x6EB SWAP1 SWAP3 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST DUP1 MLOAD SWAP1 SWAP2 POP ISZERO PUSH2 0x657 JUMPI DUP1 DUP1 PUSH1 0x20 ADD SWAP1 MLOAD DUP2 ADD SWAP1 PUSH2 0x6CF SWAP2 SWAP1 PUSH2 0x81D JUMP JUMPDEST PUSH2 0x657 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0xA5B JUMP JUMPDEST PUSH1 0x60 PUSH2 0x6FA DUP5 DUP5 PUSH1 0x0 DUP6 PUSH2 0x702 JUMP JUMPDEST SWAP5 SWAP4 POP POP POP POP JUMP JUMPDEST PUSH1 0x60 DUP3 SELFBALANCE LT ISZERO PUSH2 0x724 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0x93B JUMP JUMPDEST PUSH2 0x72D DUP6 PUSH2 0x7C2 JUMP JUMPDEST PUSH2 0x749 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP1 PUSH2 0xA24 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP7 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP6 DUP8 PUSH1 0x40 MLOAD PUSH2 0x765 SWAP2 SWAP1 PUSH2 0x855 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP8 GAS CALL SWAP3 POP POP POP RETURNDATASIZE DUP1 PUSH1 0x0 DUP2 EQ PUSH2 0x7A2 JUMPI PUSH1 0x40 MLOAD SWAP2 POP PUSH1 0x1F NOT PUSH1 0x3F RETURNDATASIZE ADD AND DUP3 ADD PUSH1 0x40 MSTORE RETURNDATASIZE DUP3 MSTORE RETURNDATASIZE PUSH1 0x0 PUSH1 0x20 DUP5 ADD RETURNDATACOPY PUSH2 0x7A7 JUMP JUMPDEST PUSH1 0x60 SWAP2 POP JUMPDEST POP SWAP2 POP SWAP2 POP PUSH2 0x7B7 DUP3 DUP3 DUP7 PUSH2 0x7C8 JUMP JUMPDEST SWAP8 SWAP7 POP POP POP POP POP POP POP JUMP JUMPDEST EXTCODESIZE ISZERO ISZERO SWAP1 JUMP JUMPDEST PUSH1 0x60 DUP4 ISZERO PUSH2 0x7D7 JUMPI POP DUP2 PUSH2 0x5C3 JUMP JUMPDEST DUP3 MLOAD ISZERO PUSH2 0x7E7 JUMPI DUP3 MLOAD DUP1 DUP5 PUSH1 0x20 ADD REVERT JUMPDEST DUP2 PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x1D1 SWAP2 SWAP1 PUSH2 0x8C2 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x812 JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 CALLDATALOAD PUSH2 0x5C3 DUP2 PUSH2 0xB27 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x82E JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 MLOAD DUP1 ISZERO ISZERO DUP2 EQ PUSH2 0x5C3 JUMPI DUP2 DUP3 REVERT JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x84E JUMPI DUP1 DUP2 REVERT JUMPDEST POP CALLDATALOAD SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP3 MLOAD PUSH2 0x867 DUP2 DUP5 PUSH1 0x20 DUP8 ADD PUSH2 0xAE5 JUMP JUMPDEST SWAP2 SWAP1 SWAP2 ADD SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP2 SWAP1 SWAP2 AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP4 DUP5 AND DUP2 MSTORE SWAP2 SWAP1 SWAP3 AND PUSH1 0x20 DUP3 ADD MSTORE PUSH1 0x40 DUP2 ADD SWAP2 SWAP1 SWAP2 MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND DUP3 MSTORE PUSH1 0x20 DUP3 ADD MSTORE PUSH1 0x40 ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 MSTORE DUP3 MLOAD DUP1 PUSH1 0x20 DUP5 ADD MSTORE PUSH2 0x8E1 DUP2 PUSH1 0x40 DUP6 ADD PUSH1 0x20 DUP8 ADD PUSH2 0xAE5 JUMP JUMPDEST PUSH1 0x1F ADD PUSH1 0x1F NOT AND SWAP2 SWAP1 SWAP2 ADD PUSH1 0x40 ADD SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x26 SWAP1 DUP3 ADD MSTORE PUSH32 0x4F776E61626C653A206E6577206F776E657220697320746865207A65726F2061 PUSH1 0x40 DUP3 ADD MSTORE PUSH6 0x646472657373 PUSH1 0xD0 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x26 SWAP1 DUP3 ADD MSTORE PUSH32 0x416464726573733A20696E73756666696369656E742062616C616E636520666F PUSH1 0x40 DUP3 ADD MSTORE PUSH6 0x1C8818D85B1B PUSH1 0xD2 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x1A SWAP1 DUP3 ADD MSTORE PUSH32 0x4E6F7420656E6F756768206D6F72746761676520616D6F756E74000000000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x1C SWAP1 DUP3 ADD MSTORE PUSH32 0x72656D6F76653A20496E73756666696369656E742062616C616E636500000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE DUP2 DUP2 ADD MSTORE PUSH32 0x4F776E61626C653A2063616C6C6572206973206E6F7420746865206F776E6572 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x1D SWAP1 DUP3 ADD MSTORE PUSH32 0x416464726573733A2063616C6C20746F206E6F6E2D636F6E7472616374000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x2A SWAP1 DUP3 ADD MSTORE PUSH32 0x5361666545524332303A204552433230206F7065726174696F6E20646964206E PUSH1 0x40 DUP3 ADD MSTORE PUSH10 0x1BDD081CDD58D8D95959 PUSH1 0xB2 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST SWAP1 DUP2 MSTORE PUSH1 0x20 ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP3 NOT DUP3 GT ISZERO PUSH2 0xAC1 JUMPI PUSH2 0xAC1 PUSH2 0xB11 JUMP JUMPDEST POP ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP2 PUSH1 0x0 NOT DIV DUP4 GT DUP3 ISZERO ISZERO AND ISZERO PUSH2 0xAE0 JUMPI PUSH2 0xAE0 PUSH2 0xB11 JUMP JUMPDEST POP MUL SWAP1 JUMP JUMPDEST PUSH1 0x0 JUMPDEST DUP4 DUP2 LT ISZERO PUSH2 0xB00 JUMPI DUP2 DUP2 ADD MLOAD DUP4 DUP3 ADD MSTORE PUSH1 0x20 ADD PUSH2 0xAE8 JUMP JUMPDEST DUP4 DUP2 GT ISZERO PUSH2 0x626 JUMPI POP POP PUSH1 0x0 SWAP2 ADD MSTORE JUMP JUMPDEST PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x11 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND DUP2 EQ PUSH2 0xB3C JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP JUMP INVALID LOG2 PUSH5 0x6970667358 0x22 SLT KECCAK256 SHL 0xD9 0xA9 SELFDESTRUCT MLOAD JUMP DUP11 CALLDATALOAD 0xB2 0xEC 0xB6 0xBD 0xCE 0xAB 0xB6 0xC3 STATICCALL 0xC7 ISZERO 0xD7 LT 0xB1 LOG4 SWAP11 GASLIMIT 0xD2 PUSH7 0x8635CC4B0D6473 PUSH16 0x6C634300080100330000000000000000 \",\n\t\"sourceMap\": \"264:1596:7:-:0;;;394:2;370:26;;627:52;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;867:17:0;887:12;:10;:12::i;:::-;909:6;:18;;-1:-1:-1;;;;;;909:18:0;-1:-1:-1;;;;;909:18:0;;;;;;;942:43;;909:18;;-1:-1:-1;909:18:0;942:43;;909:6;;942:43;-1:-1:-1;661:3:7;:10;;-1:-1:-1;;;;;;661:10:7;-1:-1:-1;;;;;661:10:7;;;;;;;;;;264:1596;;586:96:4;665:10;586:96;:::o;14:323:8:-;;150:2;138:9;129:7;125:23;121:32;118:2;;;171:6;163;156:22;118:2;202:16;;-1:-1:-1;;;;;247:31:8;;237:42;;227:2;;298:6;290;283:22;227:2;326:5;108:229;-1:-1:-1;;;108:229:8:o;:::-;264:1596:7;;;;;;\"\n}"

// DeployMortgage deploys a new Ethereum contract, binding an instance of Mortgage to it.
func DeployMortgage(auth *bind.TransactOpts, backend bind.ContractBackend, _yam common.Address) (common.Address, *types.Transaction, *Mortgage, error) {
	parsed, err := abi.JSON(strings.NewReader(MortgageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MortgageBin), backend, _yam)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mortgage{MortgageCaller: MortgageCaller{contract: contract}, MortgageTransactor: MortgageTransactor{contract: contract}, MortgageFilterer: MortgageFilterer{contract: contract}}, nil
}

// Mortgage is an auto generated Go binding around an Ethereum contract.
type Mortgage struct {
	MortgageCaller     // Read-only binding to the contract
	MortgageTransactor // Write-only binding to the contract
	MortgageFilterer   // Log filterer for contract events
}

// MortgageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MortgageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortgageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MortgageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortgageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MortgageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MortgageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MortgageSession struct {
	Contract     *Mortgage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MortgageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MortgageCallerSession struct {
	Contract *MortgageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MortgageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MortgageTransactorSession struct {
	Contract     *MortgageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MortgageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MortgageRaw struct {
	Contract *Mortgage // Generic contract binding to access the raw methods on
}

// MortgageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MortgageCallerRaw struct {
	Contract *MortgageCaller // Generic read-only contract binding to access the raw methods on
}

// MortgageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MortgageTransactorRaw struct {
	Contract *MortgageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMortgage creates a new instance of Mortgage, bound to a specific deployed contract.
func NewMortgage(address common.Address, backend bind.ContractBackend) (*Mortgage, error) {
	contract, err := bindMortgage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mortgage{MortgageCaller: MortgageCaller{contract: contract}, MortgageTransactor: MortgageTransactor{contract: contract}, MortgageFilterer: MortgageFilterer{contract: contract}}, nil
}

// NewMortgageCaller creates a new read-only instance of Mortgage, bound to a specific deployed contract.
func NewMortgageCaller(address common.Address, caller bind.ContractCaller) (*MortgageCaller, error) {
	contract, err := bindMortgage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MortgageCaller{contract: contract}, nil
}

// NewMortgageTransactor creates a new write-only instance of Mortgage, bound to a specific deployed contract.
func NewMortgageTransactor(address common.Address, transactor bind.ContractTransactor) (*MortgageTransactor, error) {
	contract, err := bindMortgage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MortgageTransactor{contract: contract}, nil
}

// NewMortgageFilterer creates a new log filterer instance of Mortgage, bound to a specific deployed contract.
func NewMortgageFilterer(address common.Address, filterer bind.ContractFilterer) (*MortgageFilterer, error) {
	contract, err := bindMortgage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MortgageFilterer{contract: contract}, nil
}

// bindMortgage binds a generic wrapper to an already deployed contract.
func bindMortgage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MortgageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortgage *MortgageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mortgage.Contract.MortgageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortgage *MortgageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortgage.Contract.MortgageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortgage *MortgageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortgage.Contract.MortgageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mortgage *MortgageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mortgage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mortgage *MortgageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortgage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mortgage *MortgageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mortgage.Contract.contract.Transact(opts, method, params...)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Mortgage *MortgageCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Mortgage *MortgageSession) Amount() (*big.Int, error) {
	return _Mortgage.Contract.Amount(&_Mortgage.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Mortgage *MortgageCallerSession) Amount() (*big.Int, error) {
	return _Mortgage.Contract.Amount(&_Mortgage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mortgage *MortgageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mortgage *MortgageSession) Owner() (common.Address, error) {
	return _Mortgage.Contract.Owner(&_Mortgage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mortgage *MortgageCallerSession) Owner() (common.Address, error) {
	return _Mortgage.Contract.Owner(&_Mortgage.CallOpts)
}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mortgage *MortgageCaller) UserAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "userAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mortgage *MortgageSession) UserAddrs(arg0 *big.Int) (common.Address, error) {
	return _Mortgage.Contract.UserAddrs(&_Mortgage.CallOpts, arg0)
}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mortgage *MortgageCallerSession) UserAddrs(arg0 *big.Int) (common.Address, error) {
	return _Mortgage.Contract.UserAddrs(&_Mortgage.CallOpts, arg0)
}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mortgage *MortgageCaller) UserAddrsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "userAddrsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mortgage *MortgageSession) UserAddrsLength() (*big.Int, error) {
	return _Mortgage.Contract.UserAddrsLength(&_Mortgage.CallOpts)
}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mortgage *MortgageCallerSession) UserAddrsLength() (*big.Int, error) {
	return _Mortgage.Contract.UserAddrsLength(&_Mortgage.CallOpts)
}

// UserMap is a free data retrieval call binding the contract method 0x2bcc23f8.
//
// Solidity: function userMap(address ) view returns(address user, uint256 amount)
func (_Mortgage *MortgageCaller) UserMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "userMap", arg0)

	outstruct := new(struct {
		User   common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserMap is a free data retrieval call binding the contract method 0x2bcc23f8.
//
// Solidity: function userMap(address ) view returns(address user, uint256 amount)
func (_Mortgage *MortgageSession) UserMap(arg0 common.Address) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	return _Mortgage.Contract.UserMap(&_Mortgage.CallOpts, arg0)
}

// UserMap is a free data retrieval call binding the contract method 0x2bcc23f8.
//
// Solidity: function userMap(address ) view returns(address user, uint256 amount)
func (_Mortgage *MortgageCallerSession) UserMap(arg0 common.Address) (struct {
	User   common.Address
	Amount *big.Int
}, error) {
	return _Mortgage.Contract.UserMap(&_Mortgage.CallOpts, arg0)
}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mortgage *MortgageCaller) Yam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mortgage.contract.Call(opts, &out, "yam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mortgage *MortgageSession) Yam() (common.Address, error) {
	return _Mortgage.Contract.Yam(&_Mortgage.CallOpts)
}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mortgage *MortgageCallerSession) Yam() (common.Address, error) {
	return _Mortgage.Contract.Yam(&_Mortgage.CallOpts)
}

// Mortgage is a paid mutator transaction binding the contract method 0x1be3175e.
//
// Solidity: function mortgage(uint256 _amount) returns()
func (_Mortgage *MortgageTransactor) Mortgage(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "mortgage", _amount)
}

// Mortgage is a paid mutator transaction binding the contract method 0x1be3175e.
//
// Solidity: function mortgage(uint256 _amount) returns()
func (_Mortgage *MortgageSession) Mortgage(_amount *big.Int) (*types.Transaction, error) {
	return _Mortgage.Contract.Mortgage(&_Mortgage.TransactOpts, _amount)
}

// Mortgage is a paid mutator transaction binding the contract method 0x1be3175e.
//
// Solidity: function mortgage(uint256 _amount) returns()
func (_Mortgage *MortgageTransactorSession) Mortgage(_amount *big.Int) (*types.Transaction, error) {
	return _Mortgage.Contract.Mortgage(&_Mortgage.TransactOpts, _amount)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Mortgage *MortgageTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Mortgage *MortgageSession) Remove() (*types.Transaction, error) {
	return _Mortgage.Contract.Remove(&_Mortgage.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Mortgage *MortgageTransactorSession) Remove() (*types.Transaction, error) {
	return _Mortgage.Contract.Remove(&_Mortgage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mortgage *MortgageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mortgage *MortgageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mortgage.Contract.RenounceOwnership(&_Mortgage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mortgage *MortgageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mortgage.Contract.RenounceOwnership(&_Mortgage.TransactOpts)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _newAmount) returns()
func (_Mortgage *MortgageTransactor) SetAmount(opts *bind.TransactOpts, _newAmount *big.Int) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "setAmount", _newAmount)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _newAmount) returns()
func (_Mortgage *MortgageSession) SetAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _Mortgage.Contract.SetAmount(&_Mortgage.TransactOpts, _newAmount)
}

// SetAmount is a paid mutator transaction binding the contract method 0x271f88b4.
//
// Solidity: function setAmount(uint256 _newAmount) returns()
func (_Mortgage *MortgageTransactorSession) SetAmount(_newAmount *big.Int) (*types.Transaction, error) {
	return _Mortgage.Contract.SetAmount(&_Mortgage.TransactOpts, _newAmount)
}

// SetYam is a paid mutator transaction binding the contract method 0x9de7e363.
//
// Solidity: function setYam(address _yam) returns()
func (_Mortgage *MortgageTransactor) SetYam(opts *bind.TransactOpts, _yam common.Address) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "setYam", _yam)
}

// SetYam is a paid mutator transaction binding the contract method 0x9de7e363.
//
// Solidity: function setYam(address _yam) returns()
func (_Mortgage *MortgageSession) SetYam(_yam common.Address) (*types.Transaction, error) {
	return _Mortgage.Contract.SetYam(&_Mortgage.TransactOpts, _yam)
}

// SetYam is a paid mutator transaction binding the contract method 0x9de7e363.
//
// Solidity: function setYam(address _yam) returns()
func (_Mortgage *MortgageTransactorSession) SetYam(_yam common.Address) (*types.Transaction, error) {
	return _Mortgage.Contract.SetYam(&_Mortgage.TransactOpts, _yam)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mortgage *MortgageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mortgage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mortgage *MortgageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mortgage.Contract.TransferOwnership(&_Mortgage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mortgage *MortgageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mortgage.Contract.TransferOwnership(&_Mortgage.TransactOpts, newOwner)
}

// MortgageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mortgage contract.
type MortgageOwnershipTransferredIterator struct {
	Event *MortgageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MortgageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MortgageOwnershipTransferred)
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
		it.Event = new(MortgageOwnershipTransferred)
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
func (it *MortgageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MortgageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MortgageOwnershipTransferred represents a OwnershipTransferred event raised by the Mortgage contract.
type MortgageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mortgage *MortgageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MortgageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mortgage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MortgageOwnershipTransferredIterator{contract: _Mortgage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mortgage *MortgageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MortgageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mortgage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MortgageOwnershipTransferred)
				if err := _Mortgage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mortgage *MortgageFilterer) ParseOwnershipTransferred(log types.Log) (*MortgageOwnershipTransferred, error) {
	event := new(MortgageOwnershipTransferred)
	if err := _Mortgage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
