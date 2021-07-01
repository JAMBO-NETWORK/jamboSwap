const jamAddr = "0xf30925b40C9AEAba7bEbe89549336E49F3A4d360"
const mortgageAddr = "0x8dCd7D2De573b8020d295A948b09712DAcC69a89"
const miningAddr = "0x4de2A4cA11d4b523cd96a5d6C2516111474671b7"

const MAX_INT = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
const base18 = 1e18;
const mortgageValue = '10'; // 质押jam数量
// const url = "http://192.168.1.132:8062";
const url = "http://47.74.86.87:8062";
const chainType = "BSC";
// const chainId = 97 // BSC测试网
// const testText = "Please choose BSC test network"
const chainId = 56 // BSC主网
const testText = "Please choose BSC main network"
var curChainId;

// 链接钱包的地址
var defaultAccount = "";
var mortgageInstance;
var jamInstance;
var miningInstance;

