// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mining

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

// MiningABI is the input ABI used to generate the binding from.
const MiningABI = "[{\"inputs\":[{\"internalType\":\"contractIYam\",\"name\":\"_yam\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LpOfPid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYam\",\"name\":\"_yam\",\"type\":\"address\"}],\"name\":\"changeYam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastRewardBlock\",\"type\":\"uint256\"}],\"name\":\"getYamBlockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pending\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accYamPerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPerBlock\",\"type\":\"uint256\"}],\"name\":\"setYamPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAddrsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userIdCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userIdMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yam\",\"outputs\":[{\"internalType\":\"contractIYam\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"yamPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MiningBin is the compiled bytecode used for deploying new contracts.
var MiningBin = "0x{\n\t\"generatedSources\": [\n\t\t{\n\t\t\t\"ast\": {\n\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\"src\": \"0:339:8\",\n\t\t\t\t\"statements\": [\n\t\t\t\t\t{\n\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\"src\": \"6:3:8\",\n\t\t\t\t\t\t\"statements\": []\n\t\t\t\t\t},\n\t\t\t\t\t{\n\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\"src\": \"108:229:8\",\n\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"154:26:8\",\n\t\t\t\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"expression\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"163:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"171:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"revert\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulExpressionStatement\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"156:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"condition\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"dataEnd\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"129:7:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"138:9:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"sub\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"125:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"125:23:8\"\n\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"150:2:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"32\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"slt\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"121:3:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"121:32:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIf\",\n\t\t\t\t\t\t\t\t\t\"src\": \"118:2:8\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulVariableDeclaration\",\n\t\t\t\t\t\t\t\t\t\"src\": \"189:29:8\",\n\t\t\t\t\t\t\t\t\t\"value\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"208:9:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"mload\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"202:5:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"202:16:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"variables\": [\n\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"193:5:8\",\n\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"body\": {\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulBlock\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"281:26:8\",\n\t\t\t\t\t\t\t\t\t\t\"statements\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"expression\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"290:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"298:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"revert\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:6:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulExpressionStatement\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"283:22:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"condition\": {\n\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"240:5:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"251:5:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"arguments\": [\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"266:3:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"160\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"271:1:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"1\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"shl\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"262:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"262:11:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"kind\": \"number\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulLiteral\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"275:1:8\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"type\": \"\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"value\": \"1\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"sub\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"258:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"258:19:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"and\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"247:3:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"247:31:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"name\": \"eq\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"237:2:8\"\n\t\t\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\t\t\"src\": \"237:42:8\"\n\t\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t\t],\n\t\t\t\t\t\t\t\t\t\t\"functionName\": {\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"iszero\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"230:6:8\"\n\t\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulFunctionCall\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"230:50:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIf\",\n\t\t\t\t\t\t\t\t\t\"src\": \"227:2:8\"\n\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulAssignment\",\n\t\t\t\t\t\t\t\t\t\"src\": \"316:15:8\",\n\t\t\t\t\t\t\t\t\t\"value\": {\n\t\t\t\t\t\t\t\t\t\t\"name\": \"value\",\n\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\"src\": \"326:5:8\"\n\t\t\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t\t\t\"variableNames\": [\n\t\t\t\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\t\t\t\"nodeType\": \"YulIdentifier\",\n\t\t\t\t\t\t\t\t\t\t\t\"src\": \"316:6:8\"\n\t\t\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t\t\t]\n\t\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\t]\n\t\t\t\t\t\t},\n\t\t\t\t\t\t\"name\": \"abi_decode_tuple_t_contract$_IYam_$1057_fromMemory\",\n\t\t\t\t\t\t\"nodeType\": \"YulFunctionDefinition\",\n\t\t\t\t\t\t\"parameters\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"headStart\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"74:9:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t},\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"dataEnd\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"85:7:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t],\n\t\t\t\t\t\t\"returnVariables\": [\n\t\t\t\t\t\t\t{\n\t\t\t\t\t\t\t\t\"name\": \"value0\",\n\t\t\t\t\t\t\t\t\"nodeType\": \"YulTypedName\",\n\t\t\t\t\t\t\t\t\"src\": \"97:6:8\",\n\t\t\t\t\t\t\t\t\"type\": \"\"\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t],\n\t\t\t\t\t\t\"src\": \"14:323:8\"\n\t\t\t\t\t}\n\t\t\t\t]\n\t\t\t},\n\t\t\t\"contents\": \"{\\n    { }\\n    function abi_decode_tuple_t_contract$_IYam_$1057_fromMemory(headStart, dataEnd) -> value0\\n    {\\n        if slt(sub(dataEnd, headStart), 32) { revert(value0, value0) }\\n        let value := mload(headStart)\\n        if iszero(eq(value, and(value, sub(shl(160, 1), 1)))) { revert(value0, value0) }\\n        value0 := value\\n    }\\n}\",\n\t\t\t\"id\": 8,\n\t\t\t\"language\": \"Yul\",\n\t\t\t\"name\": \"#utility.yul\"\n\t\t}\n\t],\n\t\"linkReferences\": {},\n\t\"object\": \"60806040526729a2241af62c00006002556006805460ff19169055600060075534801561002b57600080fd5b5060405162001b7938038062001b7983398101604081905261004c916100cd565b60006100566100c9565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600180546001600160a01b0319166001600160a01b0392909216919091179055436008556100fb565b3390565b6000602082840312156100de578081fd5b81516001600160a01b03811681146100f4578182fd5b9392505050565b611a6e806200010b6000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80635c975abb116100f9578063b0c7044b11610097578063e2bbb15811610071578063e2bbb1581461035e578063e490240014610371578063e4c75c2714610384578063f2fde38b14610397576101c4565b8063b0c7044b1461033b578063ca6633751461034e578063d431b1ac14610356576101c4565b80637b79f239116100d35780637b79f239146103025780637db083dc1461030a5780638da5cb5b1461031257806393f1a40b1461031a576101c4565b80635c975abb146102dd578063630b5ba1146102f2578063715018a6146102fa576101c4565b80632f0efcfc116101665780634498e8f1116101405780634498e8f11461029c57806348cd4cb1146102af57806351eb05a6146102b75780635312ea8e146102ca576101c4565b80632f0efcfc1461026357806330689c7114610276578063441a3e7014610289576101c4565b80631526fe27116101a25780631526fe271461021157806317caf6f1146102355780631ab06ee51461023d5780632b8bbbe814610250576101c4565b8063071c0332146101c9578063081e3eda146101e75780630abe9f69146101fc575b600080fd5b6101d16103aa565b6040516101de91906116a5565b60405180910390f35b6101ef6103b9565b6040516101de9190611941565b61020f61020a3660046115f7565b6103bf565b005b61022461021f3660046115f7565b610414565b6040516101de959493929190611701565b6101ef61045f565b61020f61024b366004611668565b610465565b61020f61025e366004611656565b61053e565b6101ef6102713660046115f7565b610723565b61020f6102843660046115bb565b610743565b61020f610297366004611668565b6107a4565b6101ef6102aa3660046115bb565b6107d6565b6101ef6107e8565b61020f6102c53660046115f7565b6107ee565b61020f6102d83660046115f7565b61095a565b6102e5610987565b6040516101de91906116f6565b61020f610990565b61020f6109b7565b6101ef610a40565b6101ef610a46565b6101d1610a4c565b61032d610328366004611627565b610a5b565b6040516101de92919061194a565b6101ef6103493660046115bb565b610a7f565b6101ef610a91565b61020f610a97565b61020f61036c366004611668565b610aea565b6101d161037f3660046115f7565b610b99565b6101ef610392366004611627565b610bc3565b61020f6103a53660046115bb565b610d87565b6001546001600160a01b031681565b60035490565b6103c7610e47565b6001600160a01b03166103d8610a4c565b6001600160a01b0316146104075760405162461bcd60e51b81526004016103fe90611854565b60405180910390fd5b61040f610990565b600255565b6003818154811061042457600080fd5b6000918252602090912060059091020180546001820154600283015460038401546004909401546001600160a01b0390931694509092909185565b60075481565b61046d610e47565b6001600160a01b031661047e610a4c565b6001600160a01b0316146104a45760405162461bcd60e51b81526004016103fe90611854565b6104ac610990565b6104fd816104f7600385815481106104d457634e487b7160e01b600052603260045260246000fd5b906000526020600020906005020160010154600754610e4b90919063ffffffff16565b90610e5e565b600781905550806003838154811061052557634e487b7160e01b600052603260045260246000fd5b9060005260206000209060050201600101819055505050565b610546610e47565b6001600160a01b0316610557610a4c565b6001600160a01b03161461057d5760405162461bcd60e51b81526004016103fe90611854565b6001600160a01b0381166105a35760405162461bcd60e51b81526004016103fe9061190a565b6105ab610990565b600060085443116105be576008546105c0565b435b6007549091506105d09084610e5e565b6007556040805160a0810182526001600160a01b038481168252602082018681529282018481526000606084018181526080850182815260038054600180820183559190945295517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b600590940293840180546001600160a01b031916919096161790945594517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c82015590517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d82015592517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85e840155517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85f909201919091556106f86103b9565b61070291906119af565b6001600160a01b039092166000908152600560205260409020919091555050565b60025460009061073d906107374385610e4b565b90610e6a565b92915050565b61074b610e47565b6001600160a01b031661075c610a4c565b6001600160a01b0316146107825760405162461bcd60e51b81526004016103fe90611854565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b60065460ff16156107c75760405162461bcd60e51b81526004016103fe90611762565b6107d2828233610e76565b5050565b600a6020526000908152604090205481565b60085481565b60006003828154811061081157634e487b7160e01b600052603260045260246000fd5b90600052602060002090600502019050806002015443116108325750610957565b600481015480610849575043600290910155610957565b60006108588360020154610723565b90506000811161086a57505050610957565b600061088f600754610889866001015485610e6a90919063ffffffff16565b90610fdf565b6001546040516340c10f1960e01b81529192506000916001600160a01b03909116906340c10f19906108c790309086906004016116dd565b602060405180830381600087803b1580156108e157600080fd5b505af11580156108f5573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061091991906115d7565b9050801561094857610942610937856108898564e8d4a51000610e6a565b600387015490610e5e565b60038601555b43856002018190555050505050505b50565b60065460ff161561097d5760405162461bcd60e51b81526004016103fe90611762565b6109578133610feb565b60065460ff1681565b60035460005b818110156107d2576109a7816107ee565b6109b0816119f2565b9050610996565b6109bf610e47565b6001600160a01b03166109d0610a4c565b6001600160a01b0316146109f65760405162461bcd60e51b81526004016103fe90611854565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b60025481565b600b5490565b6000546001600160a01b031690565b60046020908152600092835260408084209091529082529020805460019091015482565b60056020526000908152604090205481565b60095481565b610a9f610e47565b6001600160a01b0316610ab0610a4c565b6001600160a01b031614610ad65760405162461bcd60e51b81526004016103fe90611854565b6006805460ff19811660ff90911615179055565b60065460ff1615610b0d5760405162461bcd60e51b81526004016103fe90611762565b336000908152600a6020526040902054610b8e5760098054906000610b31836119f2565b9091555050600954336000818152600a6020526040812092909255600b805460018101825592527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db990910180546001600160a01b03191690911790555b6107d28282336110bc565b600b8181548110610ba957600080fd5b6000918252602090912001546001600160a01b0316905081565b60008060038481548110610be757634e487b7160e01b600052603260045260246000fd5b60009182526020808320878452600480835260408086206001600160a01b03808b16885294528086206003600590960290930194850154855491516370a0823160e01b815295975092959294929316916370a0823191610c49913091016116a5565b60206040518083038186803b158015610c6157600080fd5b505afa158015610c75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c99919061160f565b835490915015610d7a578360020154431115610d3b576000610cbe8560020154610723565b90506000610cdf600754610889886001015485610e6a90919063ffffffff16565b9050610cfe610cf7846108898464e8d4a51000610e6a565b8590610e5e565b9350610d2e8560010154610d2864e8d4a51000610889888a60000154610e6a90919063ffffffff16565b90610e4b565b965050505050505061073d565b8360020154431415610d7a57610d6f8360010154610d2864e8d4a51000610889868860000154610e6a90919063ffffffff16565b94505050505061073d565b5060009695505050505050565b610d8f610e47565b6001600160a01b0316610da0610a4c565b6001600160a01b031614610dc65760405162461bcd60e51b81526004016103fe90611854565b6001600160a01b038116610dec5760405162461bcd60e51b81526004016103fe90611799565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b3390565b6000610e5782846119af565b9392505050565b6000610e578284611958565b6000610e578284611990565b600060038481548110610e9957634e487b7160e01b600052603260045260246000fd5b600091825260208083208784526004825260408085206001600160a01b03881686529092529220805460059092029092019250841115610eeb5760405162461bcd60e51b81526004016103fe906117df565b610ef4856107ee565b6000610f228260010154610d2864e8d4a5100061088987600301548760000154610e6a90919063ffffffff16565b90508015610f3457610f3484826111ff565b8415610f72578154610f469086610e4b565b82556004830154610f579086610e4b565b60048401558254610f72906001600160a01b0316858761139f565b60038301548254610f8d9164e8d4a510009161088991610e6a565b826001018190555085846001600160a01b03167ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b56887604051610fcf9190611941565b60405180910390a3505050505050565b6000610e578284611970565b60006003838154811061100e57634e487b7160e01b600052603260045260246000fd5b600091825260208083208684526004825260408085206001600160a01b0380891687529352842080548582556001820195909555600590930201805490945091929161105c9116858361139f565b600483015461106b9082610e4b565b836004018190555084846001600160a01b03167fbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595836040516110ad9190611941565b60405180910390a35050505050565b6000600384815481106110df57634e487b7160e01b600052603260045260246000fd5b600091825260208083208784526004825260408085206001600160a01b0388168652909252922060059091029091019150611119856107ee565b80541561116257600061114e8260010154610d2864e8d4a5100061088987600301548760000154610e6a90919063ffffffff16565b905080156111605761116084826111ff565b505b83156111a257815461117f906001600160a01b03168430876113f5565b805461118b9085610e5e565b8155600482015461119c9085610e5e565b60048301555b600382015481546111bd9164e8d4a510009161088991610e6a565b816001018190555084836001600160a01b03167f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15866040516110ad9190611941565b6001546040516370a0823160e01b81526000916001600160a01b0316906370a08231906112309030906004016116a5565b60206040518083038186803b15801561124857600080fd5b505afa15801561125c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611280919061160f565b9050808211156113145760015460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb906112bc90869085906004016116dd565b602060405180830381600087803b1580156112d657600080fd5b505af11580156112ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061130e91906115d7565b5061139a565b60015460405163a9059cbb60e01b81526001600160a01b039091169063a9059cbb9061134690869086906004016116dd565b602060405180830381600087803b15801561136057600080fd5b505af1158015611374573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061139891906115d7565b505b505050565b61139a8363a9059cbb60e01b84846040516024016113be9291906116dd565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611416565b611398846323b872dd60e01b8585856040516024016113be939291906116b9565b600061146b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166114a59092919063ffffffff16565b80519091501561139a578080602001905181019061148991906115d7565b61139a5760405162461bcd60e51b81526004016103fe906118c0565b60606114b484846000856114bc565b949350505050565b6060824710156114de5760405162461bcd60e51b81526004016103fe9061180e565b6114e78561157c565b6115035760405162461bcd60e51b81526004016103fe90611889565b600080866001600160a01b0316858760405161151f9190611689565b60006040518083038185875af1925050503d806000811461155c576040519150601f19603f3d011682016040523d82523d6000602084013e611561565b606091505b5091509150611571828286611582565b979650505050505050565b3b151590565b60608315611591575081610e57565b8251156115a15782518084602001fd5b8160405162461bcd60e51b81526004016103fe919061172f565b6000602082840312156115cc578081fd5b8135610e5781611a23565b6000602082840312156115e8578081fd5b81518015158114610e57578182fd5b600060208284031215611608578081fd5b5035919050565b600060208284031215611620578081fd5b5051919050565b60008060408385031215611639578081fd5b82359150602083013561164b81611a23565b809150509250929050565b60008060408385031215611639578182fd5b6000806040838503121561167a578182fd5b50508035926020909101359150565b6000825161169b8184602087016119c6565b9190910192915050565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b901515815260200190565b6001600160a01b03959095168552602085019390935260408401919091526060830152608082015260a00190565b600060208252825180602084015261174e8160408501602087016119c6565b601f01601f19169190910160400192915050565b60208082526019908201527f4d696e696e6720686173206265656e2073757370656e64656400000000000000604082015260600190565b60208082526026908201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160408201526564647265737360d01b606082015260800190565b6020808252601590820152741dda5d1a191c985dd6585b4e881b9bdd0819dbdbd9605a1b604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b6020808252601c908201527f5f6c70546f6b656e20697320746865207a65726f206164647265737300000000604082015260600190565b90815260200190565b918252602082015260400190565b6000821982111561196b5761196b611a0d565b500190565b60008261198b57634e487b7160e01b81526012600452602481fd5b500490565b60008160001904831182151516156119aa576119aa611a0d565b500290565b6000828210156119c1576119c1611a0d565b500390565b60005b838110156119e15781810151838201526020016119c9565b838111156113985750506000910152565b6000600019821415611a0657611a06611a0d565b5060010190565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b038116811461095757600080fdfea264697066735822122045644365a536b9cffce2667384b5367ff5956b69463fbe6d6f22c7f5c4aa7c9a64736f6c63430008010033\",\n\t\"opcodes\": \"PUSH1 0x80 PUSH1 0x40 MSTORE PUSH8 0x29A2241AF62C0000 PUSH1 0x2 SSTORE PUSH1 0x6 DUP1 SLOAD PUSH1 0xFF NOT AND SWAP1 SSTORE PUSH1 0x0 PUSH1 0x7 SSTORE CALLVALUE DUP1 ISZERO PUSH2 0x2B JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x40 MLOAD PUSH3 0x1B79 CODESIZE SUB DUP1 PUSH3 0x1B79 DUP4 CODECOPY DUP2 ADD PUSH1 0x40 DUP2 SWAP1 MSTORE PUSH2 0x4C SWAP2 PUSH2 0xCD JUMP JUMPDEST PUSH1 0x0 PUSH2 0x56 PUSH2 0xC9 JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP4 AND SWAP1 DUP2 OR DUP3 SSTORE PUSH1 0x40 MLOAD SWAP3 SWAP4 POP SWAP2 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP1 DUP3 SWAP1 LOG3 POP PUSH1 0x1 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE NUMBER PUSH1 0x8 SSTORE PUSH2 0xFB JUMP JUMPDEST CALLER SWAP1 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0xDE JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND DUP2 EQ PUSH2 0xF4 JUMPI DUP2 DUP3 REVERT JUMPDEST SWAP4 SWAP3 POP POP POP JUMP JUMPDEST PUSH2 0x1A6E DUP1 PUSH3 0x10B PUSH1 0x0 CODECOPY PUSH1 0x0 RETURN INVALID PUSH1 0x80 PUSH1 0x40 MSTORE CALLVALUE DUP1 ISZERO PUSH2 0x10 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP PUSH1 0x4 CALLDATASIZE LT PUSH2 0x1C4 JUMPI PUSH1 0x0 CALLDATALOAD PUSH1 0xE0 SHR DUP1 PUSH4 0x5C975ABB GT PUSH2 0xF9 JUMPI DUP1 PUSH4 0xB0C7044B GT PUSH2 0x97 JUMPI DUP1 PUSH4 0xE2BBB158 GT PUSH2 0x71 JUMPI DUP1 PUSH4 0xE2BBB158 EQ PUSH2 0x35E JUMPI DUP1 PUSH4 0xE4902400 EQ PUSH2 0x371 JUMPI DUP1 PUSH4 0xE4C75C27 EQ PUSH2 0x384 JUMPI DUP1 PUSH4 0xF2FDE38B EQ PUSH2 0x397 JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0xB0C7044B EQ PUSH2 0x33B JUMPI DUP1 PUSH4 0xCA663375 EQ PUSH2 0x34E JUMPI DUP1 PUSH4 0xD431B1AC EQ PUSH2 0x356 JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x7B79F239 GT PUSH2 0xD3 JUMPI DUP1 PUSH4 0x7B79F239 EQ PUSH2 0x302 JUMPI DUP1 PUSH4 0x7DB083DC EQ PUSH2 0x30A JUMPI DUP1 PUSH4 0x8DA5CB5B EQ PUSH2 0x312 JUMPI DUP1 PUSH4 0x93F1A40B EQ PUSH2 0x31A JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x5C975ABB EQ PUSH2 0x2DD JUMPI DUP1 PUSH4 0x630B5BA1 EQ PUSH2 0x2F2 JUMPI DUP1 PUSH4 0x715018A6 EQ PUSH2 0x2FA JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x2F0EFCFC GT PUSH2 0x166 JUMPI DUP1 PUSH4 0x4498E8F1 GT PUSH2 0x140 JUMPI DUP1 PUSH4 0x4498E8F1 EQ PUSH2 0x29C JUMPI DUP1 PUSH4 0x48CD4CB1 EQ PUSH2 0x2AF JUMPI DUP1 PUSH4 0x51EB05A6 EQ PUSH2 0x2B7 JUMPI DUP1 PUSH4 0x5312EA8E EQ PUSH2 0x2CA JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x2F0EFCFC EQ PUSH2 0x263 JUMPI DUP1 PUSH4 0x30689C71 EQ PUSH2 0x276 JUMPI DUP1 PUSH4 0x441A3E70 EQ PUSH2 0x289 JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x1526FE27 GT PUSH2 0x1A2 JUMPI DUP1 PUSH4 0x1526FE27 EQ PUSH2 0x211 JUMPI DUP1 PUSH4 0x17CAF6F1 EQ PUSH2 0x235 JUMPI DUP1 PUSH4 0x1AB06EE5 EQ PUSH2 0x23D JUMPI DUP1 PUSH4 0x2B8BBBE8 EQ PUSH2 0x250 JUMPI PUSH2 0x1C4 JUMP JUMPDEST DUP1 PUSH4 0x71C0332 EQ PUSH2 0x1C9 JUMPI DUP1 PUSH4 0x81E3EDA EQ PUSH2 0x1E7 JUMPI DUP1 PUSH4 0xABE9F69 EQ PUSH2 0x1FC JUMPI JUMPDEST PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH2 0x1D1 PUSH2 0x3AA JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0x1DE SWAP2 SWAP1 PUSH2 0x16A5 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 RETURN JUMPDEST PUSH2 0x1EF PUSH2 0x3B9 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0x1DE SWAP2 SWAP1 PUSH2 0x1941 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x20A CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0x3BF JUMP JUMPDEST STOP JUMPDEST PUSH2 0x224 PUSH2 0x21F CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0x414 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0x1DE SWAP6 SWAP5 SWAP4 SWAP3 SWAP2 SWAP1 PUSH2 0x1701 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x45F JUMP JUMPDEST PUSH2 0x20F PUSH2 0x24B CALLDATASIZE PUSH1 0x4 PUSH2 0x1668 JUMP JUMPDEST PUSH2 0x465 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x25E CALLDATASIZE PUSH1 0x4 PUSH2 0x1656 JUMP JUMPDEST PUSH2 0x53E JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x271 CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0x723 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x284 CALLDATASIZE PUSH1 0x4 PUSH2 0x15BB JUMP JUMPDEST PUSH2 0x743 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x297 CALLDATASIZE PUSH1 0x4 PUSH2 0x1668 JUMP JUMPDEST PUSH2 0x7A4 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x2AA CALLDATASIZE PUSH1 0x4 PUSH2 0x15BB JUMP JUMPDEST PUSH2 0x7D6 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x7E8 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x2C5 CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0x7EE JUMP JUMPDEST PUSH2 0x20F PUSH2 0x2D8 CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0x95A JUMP JUMPDEST PUSH2 0x2E5 PUSH2 0x987 JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0x1DE SWAP2 SWAP1 PUSH2 0x16F6 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x990 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x9B7 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0xA40 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0xA46 JUMP JUMPDEST PUSH2 0x1D1 PUSH2 0xA4C JUMP JUMPDEST PUSH2 0x32D PUSH2 0x328 CALLDATASIZE PUSH1 0x4 PUSH2 0x1627 JUMP JUMPDEST PUSH2 0xA5B JUMP JUMPDEST PUSH1 0x40 MLOAD PUSH2 0x1DE SWAP3 SWAP2 SWAP1 PUSH2 0x194A JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x349 CALLDATASIZE PUSH1 0x4 PUSH2 0x15BB JUMP JUMPDEST PUSH2 0xA7F JUMP JUMPDEST PUSH2 0x1EF PUSH2 0xA91 JUMP JUMPDEST PUSH2 0x20F PUSH2 0xA97 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x36C CALLDATASIZE PUSH1 0x4 PUSH2 0x1668 JUMP JUMPDEST PUSH2 0xAEA JUMP JUMPDEST PUSH2 0x1D1 PUSH2 0x37F CALLDATASIZE PUSH1 0x4 PUSH2 0x15F7 JUMP JUMPDEST PUSH2 0xB99 JUMP JUMPDEST PUSH2 0x1EF PUSH2 0x392 CALLDATASIZE PUSH1 0x4 PUSH2 0x1627 JUMP JUMPDEST PUSH2 0xBC3 JUMP JUMPDEST PUSH2 0x20F PUSH2 0x3A5 CALLDATASIZE PUSH1 0x4 PUSH2 0x15BB JUMP JUMPDEST PUSH2 0xD87 JUMP JUMPDEST PUSH1 0x1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP2 JUMP JUMPDEST PUSH1 0x3 SLOAD SWAP1 JUMP JUMPDEST PUSH2 0x3C7 PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x3D8 PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x407 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 REVERT JUMPDEST PUSH2 0x40F PUSH2 0x990 JUMP JUMPDEST PUSH1 0x2 SSTORE JUMP JUMPDEST PUSH1 0x3 DUP2 DUP2 SLOAD DUP2 LT PUSH2 0x424 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 SWAP1 SWAP2 KECCAK256 PUSH1 0x5 SWAP1 SWAP2 MUL ADD DUP1 SLOAD PUSH1 0x1 DUP3 ADD SLOAD PUSH1 0x2 DUP4 ADD SLOAD PUSH1 0x3 DUP5 ADD SLOAD PUSH1 0x4 SWAP1 SWAP5 ADD SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP4 AND SWAP5 POP SWAP1 SWAP3 SWAP1 SWAP2 DUP6 JUMP JUMPDEST PUSH1 0x7 SLOAD DUP2 JUMP JUMPDEST PUSH2 0x46D PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x47E PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x4A4 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH2 0x4AC PUSH2 0x990 JUMP JUMPDEST PUSH2 0x4FD DUP2 PUSH2 0x4F7 PUSH1 0x3 DUP6 DUP2 SLOAD DUP2 LT PUSH2 0x4D4 JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x5 MUL ADD PUSH1 0x1 ADD SLOAD PUSH1 0x7 SLOAD PUSH2 0xE4B SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 PUSH2 0xE5E JUMP JUMPDEST PUSH1 0x7 DUP2 SWAP1 SSTORE POP DUP1 PUSH1 0x3 DUP4 DUP2 SLOAD DUP2 LT PUSH2 0x525 JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x5 MUL ADD PUSH1 0x1 ADD DUP2 SWAP1 SSTORE POP POP POP JUMP JUMPDEST PUSH2 0x546 PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x557 PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x57D JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND PUSH2 0x5A3 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x190A JUMP JUMPDEST PUSH2 0x5AB PUSH2 0x990 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x8 SLOAD NUMBER GT PUSH2 0x5BE JUMPI PUSH1 0x8 SLOAD PUSH2 0x5C0 JUMP JUMPDEST NUMBER JUMPDEST PUSH1 0x7 SLOAD SWAP1 SWAP2 POP PUSH2 0x5D0 SWAP1 DUP5 PUSH2 0xE5E JUMP JUMPDEST PUSH1 0x7 SSTORE PUSH1 0x40 DUP1 MLOAD PUSH1 0xA0 DUP2 ADD DUP3 MSTORE PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP5 DUP2 AND DUP3 MSTORE PUSH1 0x20 DUP3 ADD DUP7 DUP2 MSTORE SWAP3 DUP3 ADD DUP5 DUP2 MSTORE PUSH1 0x0 PUSH1 0x60 DUP5 ADD DUP2 DUP2 MSTORE PUSH1 0x80 DUP6 ADD DUP3 DUP2 MSTORE PUSH1 0x3 DUP1 SLOAD PUSH1 0x1 DUP1 DUP3 ADD DUP4 SSTORE SWAP2 SWAP1 SWAP5 MSTORE SWAP6 MLOAD PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85B PUSH1 0x5 SWAP1 SWAP5 MUL SWAP4 DUP5 ADD DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND SWAP2 SWAP1 SWAP7 AND OR SWAP1 SWAP5 SSTORE SWAP5 MLOAD PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85C DUP3 ADD SSTORE SWAP1 MLOAD PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85D DUP3 ADD SSTORE SWAP3 MLOAD PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85E DUP5 ADD SSTORE MLOAD PUSH32 0xC2575A0E9E593C00F959F8C92F12DB2869C3395A3B0502D05E2516446F71F85F SWAP1 SWAP3 ADD SWAP2 SWAP1 SWAP2 SSTORE PUSH2 0x6F8 PUSH2 0x3B9 JUMP JUMPDEST PUSH2 0x702 SWAP2 SWAP1 PUSH2 0x19AF JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP3 AND PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x5 PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SWAP2 SWAP1 SWAP2 SSTORE POP POP JUMP JUMPDEST PUSH1 0x2 SLOAD PUSH1 0x0 SWAP1 PUSH2 0x73D SWAP1 PUSH2 0x737 NUMBER DUP6 PUSH2 0xE4B JUMP JUMPDEST SWAP1 PUSH2 0xE6A JUMP JUMPDEST SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH2 0x74B PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x75C PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x782 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x1 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x6 SLOAD PUSH1 0xFF AND ISZERO PUSH2 0x7C7 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1762 JUMP JUMPDEST PUSH2 0x7D2 DUP3 DUP3 CALLER PUSH2 0xE76 JUMP JUMPDEST POP POP JUMP JUMPDEST PUSH1 0xA PUSH1 0x20 MSTORE PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x8 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x3 DUP3 DUP2 SLOAD DUP2 LT PUSH2 0x811 JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST SWAP1 PUSH1 0x0 MSTORE PUSH1 0x20 PUSH1 0x0 KECCAK256 SWAP1 PUSH1 0x5 MUL ADD SWAP1 POP DUP1 PUSH1 0x2 ADD SLOAD NUMBER GT PUSH2 0x832 JUMPI POP PUSH2 0x957 JUMP JUMPDEST PUSH1 0x4 DUP2 ADD SLOAD DUP1 PUSH2 0x849 JUMPI POP NUMBER PUSH1 0x2 SWAP1 SWAP2 ADD SSTORE PUSH2 0x957 JUMP JUMPDEST PUSH1 0x0 PUSH2 0x858 DUP4 PUSH1 0x2 ADD SLOAD PUSH2 0x723 JUMP JUMPDEST SWAP1 POP PUSH1 0x0 DUP2 GT PUSH2 0x86A JUMPI POP POP POP PUSH2 0x957 JUMP JUMPDEST PUSH1 0x0 PUSH2 0x88F PUSH1 0x7 SLOAD PUSH2 0x889 DUP7 PUSH1 0x1 ADD SLOAD DUP6 PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 PUSH2 0xFDF JUMP JUMPDEST PUSH1 0x1 SLOAD PUSH1 0x40 MLOAD PUSH4 0x40C10F19 PUSH1 0xE0 SHL DUP2 MSTORE SWAP2 SWAP3 POP PUSH1 0x0 SWAP2 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 PUSH4 0x40C10F19 SWAP1 PUSH2 0x8C7 SWAP1 ADDRESS SWAP1 DUP7 SWAP1 PUSH1 0x4 ADD PUSH2 0x16DD JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x8E1 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0x8F5 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND DUP3 ADD DUP1 PUSH1 0x40 MSTORE POP DUP2 ADD SWAP1 PUSH2 0x919 SWAP2 SWAP1 PUSH2 0x15D7 JUMP JUMPDEST SWAP1 POP DUP1 ISZERO PUSH2 0x948 JUMPI PUSH2 0x942 PUSH2 0x937 DUP6 PUSH2 0x889 DUP6 PUSH5 0xE8D4A51000 PUSH2 0xE6A JUMP JUMPDEST PUSH1 0x3 DUP8 ADD SLOAD SWAP1 PUSH2 0xE5E JUMP JUMPDEST PUSH1 0x3 DUP7 ADD SSTORE JUMPDEST NUMBER DUP6 PUSH1 0x2 ADD DUP2 SWAP1 SSTORE POP POP POP POP POP POP JUMPDEST POP JUMP JUMPDEST PUSH1 0x6 SLOAD PUSH1 0xFF AND ISZERO PUSH2 0x97D JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1762 JUMP JUMPDEST PUSH2 0x957 DUP2 CALLER PUSH2 0xFEB JUMP JUMPDEST PUSH1 0x6 SLOAD PUSH1 0xFF AND DUP2 JUMP JUMPDEST PUSH1 0x3 SLOAD PUSH1 0x0 JUMPDEST DUP2 DUP2 LT ISZERO PUSH2 0x7D2 JUMPI PUSH2 0x9A7 DUP2 PUSH2 0x7EE JUMP JUMPDEST PUSH2 0x9B0 DUP2 PUSH2 0x19F2 JUMP JUMPDEST SWAP1 POP PUSH2 0x996 JUMP JUMPDEST PUSH2 0x9BF PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x9D0 PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0x9F6 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x40 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP1 DUP4 SWAP1 LOG3 PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x2 SLOAD DUP2 JUMP JUMPDEST PUSH1 0xB SLOAD SWAP1 JUMP JUMPDEST PUSH1 0x0 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND SWAP1 JUMP JUMPDEST PUSH1 0x4 PUSH1 0x20 SWAP1 DUP2 MSTORE PUSH1 0x0 SWAP3 DUP4 MSTORE PUSH1 0x40 DUP1 DUP5 KECCAK256 SWAP1 SWAP2 MSTORE SWAP1 DUP3 MSTORE SWAP1 KECCAK256 DUP1 SLOAD PUSH1 0x1 SWAP1 SWAP2 ADD SLOAD DUP3 JUMP JUMPDEST PUSH1 0x5 PUSH1 0x20 MSTORE PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SLOAD DUP2 JUMP JUMPDEST PUSH1 0x9 SLOAD DUP2 JUMP JUMPDEST PUSH2 0xA9F PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0xAB0 PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0xAD6 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x6 DUP1 SLOAD PUSH1 0xFF NOT DUP2 AND PUSH1 0xFF SWAP1 SWAP2 AND ISZERO OR SWAP1 SSTORE JUMP JUMPDEST PUSH1 0x6 SLOAD PUSH1 0xFF AND ISZERO PUSH2 0xB0D JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1762 JUMP JUMPDEST CALLER PUSH1 0x0 SWAP1 DUP2 MSTORE PUSH1 0xA PUSH1 0x20 MSTORE PUSH1 0x40 SWAP1 KECCAK256 SLOAD PUSH2 0xB8E JUMPI PUSH1 0x9 DUP1 SLOAD SWAP1 PUSH1 0x0 PUSH2 0xB31 DUP4 PUSH2 0x19F2 JUMP JUMPDEST SWAP1 SWAP2 SSTORE POP POP PUSH1 0x9 SLOAD CALLER PUSH1 0x0 DUP2 DUP2 MSTORE PUSH1 0xA PUSH1 0x20 MSTORE PUSH1 0x40 DUP2 KECCAK256 SWAP3 SWAP1 SWAP3 SSTORE PUSH1 0xB DUP1 SLOAD PUSH1 0x1 DUP2 ADD DUP3 SSTORE SWAP3 MSTORE PUSH32 0x175B7A638427703F0DBE7BB9BBF987A2551717B34E79F33B5B1008D1FA01DB9 SWAP1 SWAP2 ADD DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND SWAP1 SWAP2 OR SWAP1 SSTORE JUMPDEST PUSH2 0x7D2 DUP3 DUP3 CALLER PUSH2 0x10BC JUMP JUMPDEST PUSH1 0xB DUP2 DUP2 SLOAD DUP2 LT PUSH2 0xBA9 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 SWAP1 SWAP2 KECCAK256 ADD SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND SWAP1 POP DUP2 JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x3 DUP5 DUP2 SLOAD DUP2 LT PUSH2 0xBE7 JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 DUP1 DUP4 KECCAK256 DUP8 DUP5 MSTORE PUSH1 0x4 DUP1 DUP4 MSTORE PUSH1 0x40 DUP1 DUP7 KECCAK256 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP12 AND DUP9 MSTORE SWAP5 MSTORE DUP1 DUP7 KECCAK256 PUSH1 0x3 PUSH1 0x5 SWAP1 SWAP7 MUL SWAP1 SWAP4 ADD SWAP5 DUP6 ADD SLOAD DUP6 SLOAD SWAP2 MLOAD PUSH4 0x70A08231 PUSH1 0xE0 SHL DUP2 MSTORE SWAP6 SWAP8 POP SWAP3 SWAP6 SWAP3 SWAP5 SWAP3 SWAP4 AND SWAP2 PUSH4 0x70A08231 SWAP2 PUSH2 0xC49 SWAP2 ADDRESS SWAP2 ADD PUSH2 0x16A5 JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP7 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0xC61 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS STATICCALL ISZERO DUP1 ISZERO PUSH2 0xC75 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND DUP3 ADD DUP1 PUSH1 0x40 MSTORE POP DUP2 ADD SWAP1 PUSH2 0xC99 SWAP2 SWAP1 PUSH2 0x160F JUMP JUMPDEST DUP4 SLOAD SWAP1 SWAP2 POP ISZERO PUSH2 0xD7A JUMPI DUP4 PUSH1 0x2 ADD SLOAD NUMBER GT ISZERO PUSH2 0xD3B JUMPI PUSH1 0x0 PUSH2 0xCBE DUP6 PUSH1 0x2 ADD SLOAD PUSH2 0x723 JUMP JUMPDEST SWAP1 POP PUSH1 0x0 PUSH2 0xCDF PUSH1 0x7 SLOAD PUSH2 0x889 DUP9 PUSH1 0x1 ADD SLOAD DUP6 PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 POP PUSH2 0xCFE PUSH2 0xCF7 DUP5 PUSH2 0x889 DUP5 PUSH5 0xE8D4A51000 PUSH2 0xE6A JUMP JUMPDEST DUP6 SWAP1 PUSH2 0xE5E JUMP JUMPDEST SWAP4 POP PUSH2 0xD2E DUP6 PUSH1 0x1 ADD SLOAD PUSH2 0xD28 PUSH5 0xE8D4A51000 PUSH2 0x889 DUP9 DUP11 PUSH1 0x0 ADD SLOAD PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 PUSH2 0xE4B JUMP JUMPDEST SWAP7 POP POP POP POP POP POP POP PUSH2 0x73D JUMP JUMPDEST DUP4 PUSH1 0x2 ADD SLOAD NUMBER EQ ISZERO PUSH2 0xD7A JUMPI PUSH2 0xD6F DUP4 PUSH1 0x1 ADD SLOAD PUSH2 0xD28 PUSH5 0xE8D4A51000 PUSH2 0x889 DUP7 DUP9 PUSH1 0x0 ADD SLOAD PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP5 POP POP POP POP POP PUSH2 0x73D JUMP JUMPDEST POP PUSH1 0x0 SWAP7 SWAP6 POP POP POP POP POP POP JUMP JUMPDEST PUSH2 0xD8F PUSH2 0xE47 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0xDA0 PUSH2 0xA4C JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND EQ PUSH2 0xDC6 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1854 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND PUSH2 0xDEC JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1799 JUMP JUMPDEST PUSH1 0x0 DUP1 SLOAD PUSH1 0x40 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP6 AND SWAP4 SWAP3 AND SWAP2 PUSH32 0x8BE0079C531659141344CD1FD0A4F28419497F9722A3DAAFE3B4186F6B6457E0 SWAP2 LOG3 PUSH1 0x0 DUP1 SLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB NOT AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND SWAP2 SWAP1 SWAP2 OR SWAP1 SSTORE JUMP JUMPDEST CALLER SWAP1 JUMP JUMPDEST PUSH1 0x0 PUSH2 0xE57 DUP3 DUP5 PUSH2 0x19AF JUMP JUMPDEST SWAP4 SWAP3 POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0xE57 DUP3 DUP5 PUSH2 0x1958 JUMP JUMPDEST PUSH1 0x0 PUSH2 0xE57 DUP3 DUP5 PUSH2 0x1990 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x3 DUP5 DUP2 SLOAD DUP2 LT PUSH2 0xE99 JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 DUP1 DUP4 KECCAK256 DUP8 DUP5 MSTORE PUSH1 0x4 DUP3 MSTORE PUSH1 0x40 DUP1 DUP6 KECCAK256 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP9 AND DUP7 MSTORE SWAP1 SWAP3 MSTORE SWAP3 KECCAK256 DUP1 SLOAD PUSH1 0x5 SWAP1 SWAP3 MUL SWAP1 SWAP3 ADD SWAP3 POP DUP5 GT ISZERO PUSH2 0xEEB JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x17DF JUMP JUMPDEST PUSH2 0xEF4 DUP6 PUSH2 0x7EE JUMP JUMPDEST PUSH1 0x0 PUSH2 0xF22 DUP3 PUSH1 0x1 ADD SLOAD PUSH2 0xD28 PUSH5 0xE8D4A51000 PUSH2 0x889 DUP8 PUSH1 0x3 ADD SLOAD DUP8 PUSH1 0x0 ADD SLOAD PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 POP DUP1 ISZERO PUSH2 0xF34 JUMPI PUSH2 0xF34 DUP5 DUP3 PUSH2 0x11FF JUMP JUMPDEST DUP5 ISZERO PUSH2 0xF72 JUMPI DUP2 SLOAD PUSH2 0xF46 SWAP1 DUP7 PUSH2 0xE4B JUMP JUMPDEST DUP3 SSTORE PUSH1 0x4 DUP4 ADD SLOAD PUSH2 0xF57 SWAP1 DUP7 PUSH2 0xE4B JUMP JUMPDEST PUSH1 0x4 DUP5 ADD SSTORE DUP3 SLOAD PUSH2 0xF72 SWAP1 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP6 DUP8 PUSH2 0x139F JUMP JUMPDEST PUSH1 0x3 DUP4 ADD SLOAD DUP3 SLOAD PUSH2 0xF8D SWAP2 PUSH5 0xE8D4A51000 SWAP2 PUSH2 0x889 SWAP2 PUSH2 0xE6A JUMP JUMPDEST DUP3 PUSH1 0x1 ADD DUP2 SWAP1 SSTORE POP DUP6 DUP5 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH32 0xF279E6A1F5E320CCA91135676D9CB6E44CA8A08C0B88342BCDB1144F6511B568 DUP8 PUSH1 0x40 MLOAD PUSH2 0xFCF SWAP2 SWAP1 PUSH2 0x1941 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG3 POP POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH2 0xE57 DUP3 DUP5 PUSH2 0x1970 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x3 DUP4 DUP2 SLOAD DUP2 LT PUSH2 0x100E JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 DUP1 DUP4 KECCAK256 DUP7 DUP5 MSTORE PUSH1 0x4 DUP3 MSTORE PUSH1 0x40 DUP1 DUP6 KECCAK256 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP1 DUP10 AND DUP8 MSTORE SWAP4 MSTORE DUP5 KECCAK256 DUP1 SLOAD DUP6 DUP3 SSTORE PUSH1 0x1 DUP3 ADD SWAP6 SWAP1 SWAP6 SSTORE PUSH1 0x5 SWAP1 SWAP4 MUL ADD DUP1 SLOAD SWAP1 SWAP5 POP SWAP2 SWAP3 SWAP2 PUSH2 0x105C SWAP2 AND DUP6 DUP4 PUSH2 0x139F JUMP JUMPDEST PUSH1 0x4 DUP4 ADD SLOAD PUSH2 0x106B SWAP1 DUP3 PUSH2 0xE4B JUMP JUMPDEST DUP4 PUSH1 0x4 ADD DUP2 SWAP1 SSTORE POP DUP5 DUP5 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH32 0xBB757047C2B5F3974FE26B7C10F732E7BCE710B0952A71082702781E62AE0595 DUP4 PUSH1 0x40 MLOAD PUSH2 0x10AD SWAP2 SWAP1 PUSH2 0x1941 JUMP JUMPDEST PUSH1 0x40 MLOAD DUP1 SWAP2 SUB SWAP1 LOG3 POP POP POP POP POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0x3 DUP5 DUP2 SLOAD DUP2 LT PUSH2 0x10DF JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x32 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x0 SWAP2 DUP3 MSTORE PUSH1 0x20 DUP1 DUP4 KECCAK256 DUP8 DUP5 MSTORE PUSH1 0x4 DUP3 MSTORE PUSH1 0x40 DUP1 DUP6 KECCAK256 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP9 AND DUP7 MSTORE SWAP1 SWAP3 MSTORE SWAP3 KECCAK256 PUSH1 0x5 SWAP1 SWAP2 MUL SWAP1 SWAP2 ADD SWAP2 POP PUSH2 0x1119 DUP6 PUSH2 0x7EE JUMP JUMPDEST DUP1 SLOAD ISZERO PUSH2 0x1162 JUMPI PUSH1 0x0 PUSH2 0x114E DUP3 PUSH1 0x1 ADD SLOAD PUSH2 0xD28 PUSH5 0xE8D4A51000 PUSH2 0x889 DUP8 PUSH1 0x3 ADD SLOAD DUP8 PUSH1 0x0 ADD SLOAD PUSH2 0xE6A SWAP1 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST SWAP1 POP DUP1 ISZERO PUSH2 0x1160 JUMPI PUSH2 0x1160 DUP5 DUP3 PUSH2 0x11FF JUMP JUMPDEST POP JUMPDEST DUP4 ISZERO PUSH2 0x11A2 JUMPI DUP2 SLOAD PUSH2 0x117F SWAP1 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP5 ADDRESS DUP8 PUSH2 0x13F5 JUMP JUMPDEST DUP1 SLOAD PUSH2 0x118B SWAP1 DUP6 PUSH2 0xE5E JUMP JUMPDEST DUP2 SSTORE PUSH1 0x4 DUP3 ADD SLOAD PUSH2 0x119C SWAP1 DUP6 PUSH2 0xE5E JUMP JUMPDEST PUSH1 0x4 DUP4 ADD SSTORE JUMPDEST PUSH1 0x3 DUP3 ADD SLOAD DUP2 SLOAD PUSH2 0x11BD SWAP2 PUSH5 0xE8D4A51000 SWAP2 PUSH2 0x889 SWAP2 PUSH2 0xE6A JUMP JUMPDEST DUP2 PUSH1 0x1 ADD DUP2 SWAP1 SSTORE POP DUP5 DUP4 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH32 0x90890809C654F11D6E72A28FA60149770A0D11EC6C92319D6CEB2BB0A4EA1A15 DUP7 PUSH1 0x40 MLOAD PUSH2 0x10AD SWAP2 SWAP1 PUSH2 0x1941 JUMP JUMPDEST PUSH1 0x1 SLOAD PUSH1 0x40 MLOAD PUSH4 0x70A08231 PUSH1 0xE0 SHL DUP2 MSTORE PUSH1 0x0 SWAP2 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND SWAP1 PUSH4 0x70A08231 SWAP1 PUSH2 0x1230 SWAP1 ADDRESS SWAP1 PUSH1 0x4 ADD PUSH2 0x16A5 JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP7 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x1248 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS STATICCALL ISZERO DUP1 ISZERO PUSH2 0x125C JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND DUP3 ADD DUP1 PUSH1 0x40 MSTORE POP DUP2 ADD SWAP1 PUSH2 0x1280 SWAP2 SWAP1 PUSH2 0x160F JUMP JUMPDEST SWAP1 POP DUP1 DUP3 GT ISZERO PUSH2 0x1314 JUMPI PUSH1 0x1 SLOAD PUSH1 0x40 MLOAD PUSH4 0xA9059CBB PUSH1 0xE0 SHL DUP2 MSTORE PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 PUSH4 0xA9059CBB SWAP1 PUSH2 0x12BC SWAP1 DUP7 SWAP1 DUP6 SWAP1 PUSH1 0x4 ADD PUSH2 0x16DD JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x12D6 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0x12EA JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND DUP3 ADD DUP1 PUSH1 0x40 MSTORE POP DUP2 ADD SWAP1 PUSH2 0x130E SWAP2 SWAP1 PUSH2 0x15D7 JUMP JUMPDEST POP PUSH2 0x139A JUMP JUMPDEST PUSH1 0x1 SLOAD PUSH1 0x40 MLOAD PUSH4 0xA9059CBB PUSH1 0xE0 SHL DUP2 MSTORE PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP1 SWAP2 AND SWAP1 PUSH4 0xA9059CBB SWAP1 PUSH2 0x1346 SWAP1 DUP7 SWAP1 DUP7 SWAP1 PUSH1 0x4 ADD PUSH2 0x16DD JUMP JUMPDEST PUSH1 0x20 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 PUSH1 0x0 DUP8 DUP1 EXTCODESIZE ISZERO DUP1 ISZERO PUSH2 0x1360 JUMPI PUSH1 0x0 DUP1 REVERT JUMPDEST POP GAS CALL ISZERO DUP1 ISZERO PUSH2 0x1374 JUMPI RETURNDATASIZE PUSH1 0x0 DUP1 RETURNDATACOPY RETURNDATASIZE PUSH1 0x0 REVERT JUMPDEST POP POP POP POP PUSH1 0x40 MLOAD RETURNDATASIZE PUSH1 0x1F NOT PUSH1 0x1F DUP3 ADD AND DUP3 ADD DUP1 PUSH1 0x40 MSTORE POP DUP2 ADD SWAP1 PUSH2 0x1398 SWAP2 SWAP1 PUSH2 0x15D7 JUMP JUMPDEST POP JUMPDEST POP POP POP JUMP JUMPDEST PUSH2 0x139A DUP4 PUSH4 0xA9059CBB PUSH1 0xE0 SHL DUP5 DUP5 PUSH1 0x40 MLOAD PUSH1 0x24 ADD PUSH2 0x13BE SWAP3 SWAP2 SWAP1 PUSH2 0x16DD JUMP JUMPDEST PUSH1 0x40 DUP1 MLOAD PUSH1 0x1F NOT DUP2 DUP5 SUB ADD DUP2 MSTORE SWAP2 SWAP1 MSTORE PUSH1 0x20 DUP2 ADD DUP1 MLOAD PUSH1 0x1 PUSH1 0x1 PUSH1 0xE0 SHL SUB AND PUSH1 0x1 PUSH1 0x1 PUSH1 0xE0 SHL SUB NOT SWAP1 SWAP4 AND SWAP3 SWAP1 SWAP3 OR SWAP1 SWAP2 MSTORE PUSH2 0x1416 JUMP JUMPDEST PUSH2 0x1398 DUP5 PUSH4 0x23B872DD PUSH1 0xE0 SHL DUP6 DUP6 DUP6 PUSH1 0x40 MLOAD PUSH1 0x24 ADD PUSH2 0x13BE SWAP4 SWAP3 SWAP2 SWAP1 PUSH2 0x16B9 JUMP JUMPDEST PUSH1 0x0 PUSH2 0x146B DUP3 PUSH1 0x40 MLOAD DUP1 PUSH1 0x40 ADD PUSH1 0x40 MSTORE DUP1 PUSH1 0x20 DUP2 MSTORE PUSH1 0x20 ADD PUSH32 0x5361666545524332303A206C6F772D6C6576656C2063616C6C206661696C6564 DUP2 MSTORE POP DUP6 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND PUSH2 0x14A5 SWAP1 SWAP3 SWAP2 SWAP1 PUSH4 0xFFFFFFFF AND JUMP JUMPDEST DUP1 MLOAD SWAP1 SWAP2 POP ISZERO PUSH2 0x139A JUMPI DUP1 DUP1 PUSH1 0x20 ADD SWAP1 MLOAD DUP2 ADD SWAP1 PUSH2 0x1489 SWAP2 SWAP1 PUSH2 0x15D7 JUMP JUMPDEST PUSH2 0x139A JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x18C0 JUMP JUMPDEST PUSH1 0x60 PUSH2 0x14B4 DUP5 DUP5 PUSH1 0x0 DUP6 PUSH2 0x14BC JUMP JUMPDEST SWAP5 SWAP4 POP POP POP POP JUMP JUMPDEST PUSH1 0x60 DUP3 SELFBALANCE LT ISZERO PUSH2 0x14DE JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x180E JUMP JUMPDEST PUSH2 0x14E7 DUP6 PUSH2 0x157C JUMP JUMPDEST PUSH2 0x1503 JUMPI PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP1 PUSH2 0x1889 JUMP JUMPDEST PUSH1 0x0 DUP1 DUP7 PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB AND DUP6 DUP8 PUSH1 0x40 MLOAD PUSH2 0x151F SWAP2 SWAP1 PUSH2 0x1689 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x40 MLOAD DUP1 DUP4 SUB DUP2 DUP6 DUP8 GAS CALL SWAP3 POP POP POP RETURNDATASIZE DUP1 PUSH1 0x0 DUP2 EQ PUSH2 0x155C JUMPI PUSH1 0x40 MLOAD SWAP2 POP PUSH1 0x1F NOT PUSH1 0x3F RETURNDATASIZE ADD AND DUP3 ADD PUSH1 0x40 MSTORE RETURNDATASIZE DUP3 MSTORE RETURNDATASIZE PUSH1 0x0 PUSH1 0x20 DUP5 ADD RETURNDATACOPY PUSH2 0x1561 JUMP JUMPDEST PUSH1 0x60 SWAP2 POP JUMPDEST POP SWAP2 POP SWAP2 POP PUSH2 0x1571 DUP3 DUP3 DUP7 PUSH2 0x1582 JUMP JUMPDEST SWAP8 SWAP7 POP POP POP POP POP POP POP JUMP JUMPDEST EXTCODESIZE ISZERO ISZERO SWAP1 JUMP JUMPDEST PUSH1 0x60 DUP4 ISZERO PUSH2 0x1591 JUMPI POP DUP2 PUSH2 0xE57 JUMP JUMPDEST DUP3 MLOAD ISZERO PUSH2 0x15A1 JUMPI DUP3 MLOAD DUP1 DUP5 PUSH1 0x20 ADD REVERT JUMPDEST DUP2 PUSH1 0x40 MLOAD PUSH3 0x461BCD PUSH1 0xE5 SHL DUP2 MSTORE PUSH1 0x4 ADD PUSH2 0x3FE SWAP2 SWAP1 PUSH2 0x172F JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x15CC JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 CALLDATALOAD PUSH2 0xE57 DUP2 PUSH2 0x1A23 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x15E8 JUMPI DUP1 DUP2 REVERT JUMPDEST DUP2 MLOAD DUP1 ISZERO ISZERO DUP2 EQ PUSH2 0xE57 JUMPI DUP2 DUP3 REVERT JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x1608 JUMPI DUP1 DUP2 REVERT JUMPDEST POP CALLDATALOAD SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 DUP5 SUB SLT ISZERO PUSH2 0x1620 JUMPI DUP1 DUP2 REVERT JUMPDEST POP MLOAD SWAP2 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x40 DUP4 DUP6 SUB SLT ISZERO PUSH2 0x1639 JUMPI DUP1 DUP2 REVERT JUMPDEST DUP3 CALLDATALOAD SWAP2 POP PUSH1 0x20 DUP4 ADD CALLDATALOAD PUSH2 0x164B DUP2 PUSH2 0x1A23 JUMP JUMPDEST DUP1 SWAP2 POP POP SWAP3 POP SWAP3 SWAP1 POP JUMP JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x40 DUP4 DUP6 SUB SLT ISZERO PUSH2 0x1639 JUMPI DUP2 DUP3 REVERT JUMPDEST PUSH1 0x0 DUP1 PUSH1 0x40 DUP4 DUP6 SUB SLT ISZERO PUSH2 0x167A JUMPI DUP2 DUP3 REVERT JUMPDEST POP POP DUP1 CALLDATALOAD SWAP3 PUSH1 0x20 SWAP1 SWAP2 ADD CALLDATALOAD SWAP2 POP JUMP JUMPDEST PUSH1 0x0 DUP3 MLOAD PUSH2 0x169B DUP2 DUP5 PUSH1 0x20 DUP8 ADD PUSH2 0x19C6 JUMP JUMPDEST SWAP2 SWAP1 SWAP2 ADD SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP2 SWAP1 SWAP2 AND DUP2 MSTORE PUSH1 0x20 ADD SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP4 DUP5 AND DUP2 MSTORE SWAP2 SWAP1 SWAP3 AND PUSH1 0x20 DUP3 ADD MSTORE PUSH1 0x40 DUP2 ADD SWAP2 SWAP1 SWAP2 MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP3 SWAP1 SWAP3 AND DUP3 MSTORE PUSH1 0x20 DUP3 ADD MSTORE PUSH1 0x40 ADD SWAP1 JUMP JUMPDEST SWAP1 ISZERO ISZERO DUP2 MSTORE PUSH1 0x20 ADD SWAP1 JUMP JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB SWAP6 SWAP1 SWAP6 AND DUP6 MSTORE PUSH1 0x20 DUP6 ADD SWAP4 SWAP1 SWAP4 MSTORE PUSH1 0x40 DUP5 ADD SWAP2 SWAP1 SWAP2 MSTORE PUSH1 0x60 DUP4 ADD MSTORE PUSH1 0x80 DUP3 ADD MSTORE PUSH1 0xA0 ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 PUSH1 0x20 DUP3 MSTORE DUP3 MLOAD DUP1 PUSH1 0x20 DUP5 ADD MSTORE PUSH2 0x174E DUP2 PUSH1 0x40 DUP6 ADD PUSH1 0x20 DUP8 ADD PUSH2 0x19C6 JUMP JUMPDEST PUSH1 0x1F ADD PUSH1 0x1F NOT AND SWAP2 SWAP1 SWAP2 ADD PUSH1 0x40 ADD SWAP3 SWAP2 POP POP JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x19 SWAP1 DUP3 ADD MSTORE PUSH32 0x4D696E696E6720686173206265656E2073757370656E64656400000000000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x26 SWAP1 DUP3 ADD MSTORE PUSH32 0x4F776E61626C653A206E6577206F776E657220697320746865207A65726F2061 PUSH1 0x40 DUP3 ADD MSTORE PUSH6 0x646472657373 PUSH1 0xD0 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x15 SWAP1 DUP3 ADD MSTORE PUSH21 0x1DDA5D1A191C985DD6585B4E881B9BDD0819DBDBD9 PUSH1 0x5A SHL PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x26 SWAP1 DUP3 ADD MSTORE PUSH32 0x416464726573733A20696E73756666696369656E742062616C616E636520666F PUSH1 0x40 DUP3 ADD MSTORE PUSH6 0x1C8818D85B1B PUSH1 0xD2 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE DUP2 DUP2 ADD MSTORE PUSH32 0x4F776E61626C653A2063616C6C6572206973206E6F7420746865206F776E6572 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x1D SWAP1 DUP3 ADD MSTORE PUSH32 0x416464726573733A2063616C6C20746F206E6F6E2D636F6E7472616374000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x2A SWAP1 DUP3 ADD MSTORE PUSH32 0x5361666545524332303A204552433230206F7065726174696F6E20646964206E PUSH1 0x40 DUP3 ADD MSTORE PUSH10 0x1BDD081CDD58D8D95959 PUSH1 0xB2 SHL PUSH1 0x60 DUP3 ADD MSTORE PUSH1 0x80 ADD SWAP1 JUMP JUMPDEST PUSH1 0x20 DUP1 DUP3 MSTORE PUSH1 0x1C SWAP1 DUP3 ADD MSTORE PUSH32 0x5F6C70546F6B656E20697320746865207A65726F206164647265737300000000 PUSH1 0x40 DUP3 ADD MSTORE PUSH1 0x60 ADD SWAP1 JUMP JUMPDEST SWAP1 DUP2 MSTORE PUSH1 0x20 ADD SWAP1 JUMP JUMPDEST SWAP2 DUP3 MSTORE PUSH1 0x20 DUP3 ADD MSTORE PUSH1 0x40 ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP3 NOT DUP3 GT ISZERO PUSH2 0x196B JUMPI PUSH2 0x196B PUSH2 0x1A0D JUMP JUMPDEST POP ADD SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP3 PUSH2 0x198B JUMPI PUSH4 0x4E487B71 PUSH1 0xE0 SHL DUP2 MSTORE PUSH1 0x12 PUSH1 0x4 MSTORE PUSH1 0x24 DUP2 REVERT JUMPDEST POP DIV SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP2 PUSH1 0x0 NOT DIV DUP4 GT DUP3 ISZERO ISZERO AND ISZERO PUSH2 0x19AA JUMPI PUSH2 0x19AA PUSH2 0x1A0D JUMP JUMPDEST POP MUL SWAP1 JUMP JUMPDEST PUSH1 0x0 DUP3 DUP3 LT ISZERO PUSH2 0x19C1 JUMPI PUSH2 0x19C1 PUSH2 0x1A0D JUMP JUMPDEST POP SUB SWAP1 JUMP JUMPDEST PUSH1 0x0 JUMPDEST DUP4 DUP2 LT ISZERO PUSH2 0x19E1 JUMPI DUP2 DUP2 ADD MLOAD DUP4 DUP3 ADD MSTORE PUSH1 0x20 ADD PUSH2 0x19C9 JUMP JUMPDEST DUP4 DUP2 GT ISZERO PUSH2 0x1398 JUMPI POP POP PUSH1 0x0 SWAP2 ADD MSTORE JUMP JUMPDEST PUSH1 0x0 PUSH1 0x0 NOT DUP3 EQ ISZERO PUSH2 0x1A06 JUMPI PUSH2 0x1A06 PUSH2 0x1A0D JUMP JUMPDEST POP PUSH1 0x1 ADD SWAP1 JUMP JUMPDEST PUSH4 0x4E487B71 PUSH1 0xE0 SHL PUSH1 0x0 MSTORE PUSH1 0x11 PUSH1 0x4 MSTORE PUSH1 0x24 PUSH1 0x0 REVERT JUMPDEST PUSH1 0x1 PUSH1 0x1 PUSH1 0xA0 SHL SUB DUP2 AND DUP2 EQ PUSH2 0x957 JUMPI PUSH1 0x0 DUP1 REVERT INVALID LOG2 PUSH5 0x6970667358 0x22 SLT KECCAK256 GASLIMIT PUSH5 0x4365A536B9 0xCF 0xFC 0xE2 PUSH7 0x7384B5367FF595 PUSH12 0x69463FBE6D6F22C7F5C4AA7C SWAP11 PUSH5 0x736F6C6343 STOP ADDMOD ADD STOP CALLER \",\n\t\"sourceMap\": \"322:9483:7:-:0;;;1196:8;1167:37;;1497:26;;;-1:-1:-1;;1497:26:7;;;1518:5;1618:34;;2283:88;;;;;;;;;;;;;;;;;;;;;;;;;;;;:::i;:::-;867:17:0;887:12;:10;:12::i;:::-;909:6;:18;;-1:-1:-1;;;;;;909:18:0;-1:-1:-1;;;;;909:18:0;;;;;;;942:43;;909:18;;-1:-1:-1;909:18:0;942:43;;909:6;;942:43;-1:-1:-1;2317:3:7;:10;;-1:-1:-1;;;;;;2317:10:7;-1:-1:-1;;;;;2317:10:7;;;;;;;;;;2351:12;2338:10;:25;322:9483;;586:96:4;665:10;586:96;:::o;14:323:8:-;;150:2;138:9;129:7;125:23;121:32;118:2;;;171:6;163;156:22;118:2;202:16;;-1:-1:-1;;;;;247:31:8;;237:42;;227:2;;298:6;290;283:22;227:2;326:5;108:229;-1:-1:-1;;;108:229:8:o;:::-;322:9483:7;;;;;;\"\n}"

// DeployMining deploys a new Ethereum contract, binding an instance of Mining to it.
func DeployMining(auth *bind.TransactOpts, backend bind.ContractBackend, _yam common.Address) (common.Address, *types.Transaction, *Mining, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiningBin), backend, _yam)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mining{MiningCaller: MiningCaller{contract: contract}, MiningTransactor: MiningTransactor{contract: contract}, MiningFilterer: MiningFilterer{contract: contract}}, nil
}

// Mining is an auto generated Go binding around an Ethereum contract.
type Mining struct {
	MiningCaller     // Read-only binding to the contract
	MiningTransactor // Write-only binding to the contract
	MiningFilterer   // Log filterer for contract events
}

// MiningCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiningCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiningTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiningFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiningSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiningSession struct {
	Contract     *Mining           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiningCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiningCallerSession struct {
	Contract *MiningCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MiningTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiningTransactorSession struct {
	Contract     *MiningTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiningRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiningRaw struct {
	Contract *Mining // Generic contract binding to access the raw methods on
}

// MiningCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiningCallerRaw struct {
	Contract *MiningCaller // Generic read-only contract binding to access the raw methods on
}

// MiningTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiningTransactorRaw struct {
	Contract *MiningTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMining creates a new instance of Mining, bound to a specific deployed contract.
func NewMining(address common.Address, backend bind.ContractBackend) (*Mining, error) {
	contract, err := bindMining(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mining{MiningCaller: MiningCaller{contract: contract}, MiningTransactor: MiningTransactor{contract: contract}, MiningFilterer: MiningFilterer{contract: contract}}, nil
}

// NewMiningCaller creates a new read-only instance of Mining, bound to a specific deployed contract.
func NewMiningCaller(address common.Address, caller bind.ContractCaller) (*MiningCaller, error) {
	contract, err := bindMining(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiningCaller{contract: contract}, nil
}

// NewMiningTransactor creates a new write-only instance of Mining, bound to a specific deployed contract.
func NewMiningTransactor(address common.Address, transactor bind.ContractTransactor) (*MiningTransactor, error) {
	contract, err := bindMining(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiningTransactor{contract: contract}, nil
}

// NewMiningFilterer creates a new log filterer instance of Mining, bound to a specific deployed contract.
func NewMiningFilterer(address common.Address, filterer bind.ContractFilterer) (*MiningFilterer, error) {
	contract, err := bindMining(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiningFilterer{contract: contract}, nil
}

// bindMining binds a generic wrapper to an already deployed contract.
func bindMining(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiningABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mining *MiningRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mining.Contract.MiningCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mining *MiningRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mining.Contract.MiningTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mining *MiningRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mining.Contract.MiningTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mining *MiningCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mining.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mining *MiningTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mining.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mining *MiningTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mining.Contract.contract.Transact(opts, method, params...)
}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_Mining *MiningCaller) LpOfPid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "LpOfPid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_Mining *MiningSession) LpOfPid(arg0 common.Address) (*big.Int, error) {
	return _Mining.Contract.LpOfPid(&_Mining.CallOpts, arg0)
}

// LpOfPid is a free data retrieval call binding the contract method 0xb0c7044b.
//
// Solidity: function LpOfPid(address ) view returns(uint256)
func (_Mining *MiningCallerSession) LpOfPid(arg0 common.Address) (*big.Int, error) {
	return _Mining.Contract.LpOfPid(&_Mining.CallOpts, arg0)
}

// GetYamBlockReward is a free data retrieval call binding the contract method 0x2f0efcfc.
//
// Solidity: function getYamBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_Mining *MiningCaller) GetYamBlockReward(opts *bind.CallOpts, _lastRewardBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "getYamBlockReward", _lastRewardBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYamBlockReward is a free data retrieval call binding the contract method 0x2f0efcfc.
//
// Solidity: function getYamBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_Mining *MiningSession) GetYamBlockReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _Mining.Contract.GetYamBlockReward(&_Mining.CallOpts, _lastRewardBlock)
}

// GetYamBlockReward is a free data retrieval call binding the contract method 0x2f0efcfc.
//
// Solidity: function getYamBlockReward(uint256 _lastRewardBlock) view returns(uint256)
func (_Mining *MiningCallerSession) GetYamBlockReward(_lastRewardBlock *big.Int) (*big.Int, error) {
	return _Mining.Contract.GetYamBlockReward(&_Mining.CallOpts, _lastRewardBlock)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mining *MiningCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mining *MiningSession) Owner() (common.Address, error) {
	return _Mining.Contract.Owner(&_Mining.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mining *MiningCallerSession) Owner() (common.Address, error) {
	return _Mining.Contract.Owner(&_Mining.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mining *MiningCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mining *MiningSession) Paused() (bool, error) {
	return _Mining.Contract.Paused(&_Mining.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Mining *MiningCallerSession) Paused() (bool, error) {
	return _Mining.Contract.Paused(&_Mining.CallOpts)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_Mining *MiningCaller) Pending(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "pending", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_Mining *MiningSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Mining.Contract.Pending(&_Mining.CallOpts, _pid, _user)
}

// Pending is a free data retrieval call binding the contract method 0xe4c75c27.
//
// Solidity: function pending(uint256 _pid, address _user) view returns(uint256)
func (_Mining *MiningCallerSession) Pending(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _Mining.Contract.Pending(&_Mining.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accYamPerShare, uint256 totalAmount)
func (_Mining *MiningCaller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccYamPerShare  *big.Int
	TotalAmount     *big.Int
}, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccYamPerShare  *big.Int
		TotalAmount     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccYamPerShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accYamPerShare, uint256 totalAmount)
func (_Mining *MiningSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccYamPerShare  *big.Int
	TotalAmount     *big.Int
}, error) {
	return _Mining.Contract.PoolInfo(&_Mining.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accYamPerShare, uint256 totalAmount)
func (_Mining *MiningCallerSession) PoolInfo(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccYamPerShare  *big.Int
	TotalAmount     *big.Int
}, error) {
	return _Mining.Contract.PoolInfo(&_Mining.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Mining *MiningCaller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Mining *MiningSession) PoolLength() (*big.Int, error) {
	return _Mining.Contract.PoolLength(&_Mining.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256)
func (_Mining *MiningCallerSession) PoolLength() (*big.Int, error) {
	return _Mining.Contract.PoolLength(&_Mining.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Mining *MiningCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Mining *MiningSession) StartBlock() (*big.Int, error) {
	return _Mining.Contract.StartBlock(&_Mining.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Mining *MiningCallerSession) StartBlock() (*big.Int, error) {
	return _Mining.Contract.StartBlock(&_Mining.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Mining *MiningCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Mining *MiningSession) TotalAllocPoint() (*big.Int, error) {
	return _Mining.Contract.TotalAllocPoint(&_Mining.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Mining *MiningCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Mining.Contract.TotalAllocPoint(&_Mining.CallOpts)
}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mining *MiningCaller) UserAddrs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "userAddrs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mining *MiningSession) UserAddrs(arg0 *big.Int) (common.Address, error) {
	return _Mining.Contract.UserAddrs(&_Mining.CallOpts, arg0)
}

// UserAddrs is a free data retrieval call binding the contract method 0xe4902400.
//
// Solidity: function userAddrs(uint256 ) view returns(address)
func (_Mining *MiningCallerSession) UserAddrs(arg0 *big.Int) (common.Address, error) {
	return _Mining.Contract.UserAddrs(&_Mining.CallOpts, arg0)
}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mining *MiningCaller) UserAddrsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "userAddrsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mining *MiningSession) UserAddrsLength() (*big.Int, error) {
	return _Mining.Contract.UserAddrsLength(&_Mining.CallOpts)
}

// UserAddrsLength is a free data retrieval call binding the contract method 0x7db083dc.
//
// Solidity: function userAddrsLength() view returns(uint256)
func (_Mining *MiningCallerSession) UserAddrsLength() (*big.Int, error) {
	return _Mining.Contract.UserAddrsLength(&_Mining.CallOpts)
}

// UserIdCount is a free data retrieval call binding the contract method 0xca663375.
//
// Solidity: function userIdCount() view returns(uint256)
func (_Mining *MiningCaller) UserIdCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "userIdCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserIdCount is a free data retrieval call binding the contract method 0xca663375.
//
// Solidity: function userIdCount() view returns(uint256)
func (_Mining *MiningSession) UserIdCount() (*big.Int, error) {
	return _Mining.Contract.UserIdCount(&_Mining.CallOpts)
}

// UserIdCount is a free data retrieval call binding the contract method 0xca663375.
//
// Solidity: function userIdCount() view returns(uint256)
func (_Mining *MiningCallerSession) UserIdCount() (*big.Int, error) {
	return _Mining.Contract.UserIdCount(&_Mining.CallOpts)
}

// UserIdMap is a free data retrieval call binding the contract method 0x4498e8f1.
//
// Solidity: function userIdMap(address ) view returns(uint256)
func (_Mining *MiningCaller) UserIdMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "userIdMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserIdMap is a free data retrieval call binding the contract method 0x4498e8f1.
//
// Solidity: function userIdMap(address ) view returns(uint256)
func (_Mining *MiningSession) UserIdMap(arg0 common.Address) (*big.Int, error) {
	return _Mining.Contract.UserIdMap(&_Mining.CallOpts, arg0)
}

// UserIdMap is a free data retrieval call binding the contract method 0x4498e8f1.
//
// Solidity: function userIdMap(address ) view returns(uint256)
func (_Mining *MiningCallerSession) UserIdMap(arg0 common.Address) (*big.Int, error) {
	return _Mining.Contract.UserIdMap(&_Mining.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Mining *MiningCaller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Mining *MiningSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Mining.Contract.UserInfo(&_Mining.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_Mining *MiningCallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _Mining.Contract.UserInfo(&_Mining.CallOpts, arg0, arg1)
}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mining *MiningCaller) Yam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "yam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mining *MiningSession) Yam() (common.Address, error) {
	return _Mining.Contract.Yam(&_Mining.CallOpts)
}

// Yam is a free data retrieval call binding the contract method 0x071c0332.
//
// Solidity: function yam() view returns(address)
func (_Mining *MiningCallerSession) Yam() (common.Address, error) {
	return _Mining.Contract.Yam(&_Mining.CallOpts)
}

// YamPerBlock is a free data retrieval call binding the contract method 0x7b79f239.
//
// Solidity: function yamPerBlock() view returns(uint256)
func (_Mining *MiningCaller) YamPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mining.contract.Call(opts, &out, "yamPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YamPerBlock is a free data retrieval call binding the contract method 0x7b79f239.
//
// Solidity: function yamPerBlock() view returns(uint256)
func (_Mining *MiningSession) YamPerBlock() (*big.Int, error) {
	return _Mining.Contract.YamPerBlock(&_Mining.CallOpts)
}

// YamPerBlock is a free data retrieval call binding the contract method 0x7b79f239.
//
// Solidity: function yamPerBlock() view returns(uint256)
func (_Mining *MiningCallerSession) YamPerBlock() (*big.Int, error) {
	return _Mining.Contract.YamPerBlock(&_Mining.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_Mining *MiningTransactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "add", _allocPoint, _lpToken)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_Mining *MiningSession) Add(_allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _Mining.Contract.Add(&_Mining.TransactOpts, _allocPoint, _lpToken)
}

// Add is a paid mutator transaction binding the contract method 0x2b8bbbe8.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken) returns()
func (_Mining *MiningTransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address) (*types.Transaction, error) {
	return _Mining.Contract.Add(&_Mining.TransactOpts, _allocPoint, _lpToken)
}

// ChangeYam is a paid mutator transaction binding the contract method 0x30689c71.
//
// Solidity: function changeYam(address _yam) returns()
func (_Mining *MiningTransactor) ChangeYam(opts *bind.TransactOpts, _yam common.Address) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "changeYam", _yam)
}

// ChangeYam is a paid mutator transaction binding the contract method 0x30689c71.
//
// Solidity: function changeYam(address _yam) returns()
func (_Mining *MiningSession) ChangeYam(_yam common.Address) (*types.Transaction, error) {
	return _Mining.Contract.ChangeYam(&_Mining.TransactOpts, _yam)
}

// ChangeYam is a paid mutator transaction binding the contract method 0x30689c71.
//
// Solidity: function changeYam(address _yam) returns()
func (_Mining *MiningTransactorSession) ChangeYam(_yam common.Address) (*types.Transaction, error) {
	return _Mining.Contract.ChangeYam(&_Mining.TransactOpts, _yam)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Deposit(&_Mining.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Deposit(&_Mining.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Mining *MiningTransactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Mining *MiningSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.EmergencyWithdraw(&_Mining.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_Mining *MiningTransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.EmergencyWithdraw(&_Mining.TransactOpts, _pid)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Mining *MiningTransactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Mining *MiningSession) MassUpdatePools() (*types.Transaction, error) {
	return _Mining.Contract.MassUpdatePools(&_Mining.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_Mining *MiningTransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _Mining.Contract.MassUpdatePools(&_Mining.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mining *MiningTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mining *MiningSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mining.Contract.RenounceOwnership(&_Mining.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mining *MiningTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mining.Contract.RenounceOwnership(&_Mining.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_Mining *MiningTransactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "set", _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_Mining *MiningSession) Set(_pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Set(&_Mining.TransactOpts, _pid, _allocPoint)
}

// Set is a paid mutator transaction binding the contract method 0x1ab06ee5.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint) returns()
func (_Mining *MiningTransactorSession) Set(_pid *big.Int, _allocPoint *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Set(&_Mining.TransactOpts, _pid, _allocPoint)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Mining *MiningTransactor) SetPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "setPause")
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Mining *MiningSession) SetPause() (*types.Transaction, error) {
	return _Mining.Contract.SetPause(&_Mining.TransactOpts)
}

// SetPause is a paid mutator transaction binding the contract method 0xd431b1ac.
//
// Solidity: function setPause() returns()
func (_Mining *MiningTransactorSession) SetPause() (*types.Transaction, error) {
	return _Mining.Contract.SetPause(&_Mining.TransactOpts)
}

// SetYamPerBlock is a paid mutator transaction binding the contract method 0x0abe9f69.
//
// Solidity: function setYamPerBlock(uint256 _newPerBlock) returns()
func (_Mining *MiningTransactor) SetYamPerBlock(opts *bind.TransactOpts, _newPerBlock *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "setYamPerBlock", _newPerBlock)
}

// SetYamPerBlock is a paid mutator transaction binding the contract method 0x0abe9f69.
//
// Solidity: function setYamPerBlock(uint256 _newPerBlock) returns()
func (_Mining *MiningSession) SetYamPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.SetYamPerBlock(&_Mining.TransactOpts, _newPerBlock)
}

// SetYamPerBlock is a paid mutator transaction binding the contract method 0x0abe9f69.
//
// Solidity: function setYamPerBlock(uint256 _newPerBlock) returns()
func (_Mining *MiningTransactorSession) SetYamPerBlock(_newPerBlock *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.SetYamPerBlock(&_Mining.TransactOpts, _newPerBlock)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mining *MiningTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mining *MiningSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mining.Contract.TransferOwnership(&_Mining.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mining *MiningTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mining.Contract.TransferOwnership(&_Mining.TransactOpts, newOwner)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Mining *MiningTransactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Mining *MiningSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.UpdatePool(&_Mining.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns()
func (_Mining *MiningTransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.UpdatePool(&_Mining.TransactOpts, _pid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Withdraw(&_Mining.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Mining *MiningTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Mining.Contract.Withdraw(&_Mining.TransactOpts, _pid, _amount)
}

// MiningDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Mining contract.
type MiningDepositIterator struct {
	Event *MiningDeposit // Event containing the contract specifics and raw log

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
func (it *MiningDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningDeposit)
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
		it.Event = new(MiningDeposit)
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
func (it *MiningDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningDeposit represents a Deposit event raised by the Mining contract.
type MiningDeposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MiningDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MiningDepositIterator{contract: _Mining.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MiningDeposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningDeposit)
				if err := _Mining.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) ParseDeposit(log types.Log) (*MiningDeposit, error) {
	event := new(MiningDeposit)
	if err := _Mining.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiningEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Mining contract.
type MiningEmergencyWithdrawIterator struct {
	Event *MiningEmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MiningEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningEmergencyWithdraw)
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
		it.Event = new(MiningEmergencyWithdraw)
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
func (it *MiningEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningEmergencyWithdraw represents a EmergencyWithdraw event raised by the Mining contract.
type MiningEmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MiningEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MiningEmergencyWithdrawIterator{contract: _Mining.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MiningEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningEmergencyWithdraw)
				if err := _Mining.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) ParseEmergencyWithdraw(log types.Log) (*MiningEmergencyWithdraw, error) {
	event := new(MiningEmergencyWithdraw)
	if err := _Mining.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiningOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mining contract.
type MiningOwnershipTransferredIterator struct {
	Event *MiningOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MiningOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningOwnershipTransferred)
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
		it.Event = new(MiningOwnershipTransferred)
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
func (it *MiningOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningOwnershipTransferred represents a OwnershipTransferred event raised by the Mining contract.
type MiningOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mining *MiningFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MiningOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mining.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MiningOwnershipTransferredIterator{contract: _Mining.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mining *MiningFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MiningOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mining.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningOwnershipTransferred)
				if err := _Mining.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mining *MiningFilterer) ParseOwnershipTransferred(log types.Log) (*MiningOwnershipTransferred, error) {
	event := new(MiningOwnershipTransferred)
	if err := _Mining.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MiningWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Mining contract.
type MiningWithdrawIterator struct {
	Event *MiningWithdraw // Event containing the contract specifics and raw log

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
func (it *MiningWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MiningWithdraw)
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
		it.Event = new(MiningWithdraw)
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
func (it *MiningWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MiningWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MiningWithdraw represents a Withdraw event raised by the Mining contract.
type MiningWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MiningWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MiningWithdrawIterator{contract: _Mining.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MiningWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Mining.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MiningWithdraw)
				if err := _Mining.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_Mining *MiningFilterer) ParseWithdraw(log types.Log) (*MiningWithdraw, error) {
	event := new(MiningWithdraw)
	if err := _Mining.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
