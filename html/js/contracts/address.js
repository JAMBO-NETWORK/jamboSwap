const jamAddr = "0x490DA0C1c6387Cc60EE1734C17d246466e55C889"
const mortgageAddr = "0x31DfAe660bDa041efed02f2ed7170E29EF373f94"
const miningAddr = "0x11409781648E436A3d294e0728a87d8bcA29dA0F"

const MAX_INT = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
const base18 = 1e18;
const mortgageValue = '10'; // 质押jam数量
const url = "http://192.168.1.132:8062";
// const url = "http://119.8.103.132:8062";
const chainType = "BSC";

// 链接钱包的地址
var defaultAccount = "";
var mortgageInstance;
var jamInstance;
var miningInstance;

