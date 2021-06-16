// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "./IYam.sol";

contract Mortgage is Ownable {
    using SafeMath for uint256;
    using SafeERC20 for IYam;

    uint256 public amount = 10; // 需要抵押的数量

    IYam public yam;

    struct User {
        address user;
        uint256 amount;
    }

    address[] public userAddrs;
    mapping(address => User) public userMap;

    constructor(IYam _yam) {
        yam = _yam;
    }

    function mortgage(uint256 _amount) public {
        require(_amount >= amount.mul(1e18), "Not enough mortgage amount");
        address _sender = _msgSender();
        yam.safeTransferFrom(_sender, address(this), _amount);
        if (userMap[_sender].user == address(0)) { // new user
            userMap[_sender] =  User({
            user : _sender,
            amount : _amount
            });
            userAddrs.push(_sender);
        } else { // old user
            userMap[_sender].amount = userMap[_sender].amount.add(_amount);
        }
    }

    function remove() public {
        address _sender = _msgSender();
        uint256 _amount = userMap[_sender].amount;
        if (_amount > 0) {
            userMap[_sender].amount = 0;
            yam.safeTransfer(_sender, _amount);
        }
    }

    function userAddrsLength() public view returns (uint256) {
        return userAddrs.length;
    }

    function setYam(IYam _yam) public onlyOwner {
        yam = _yam;
    }

    function setAmount(uint256 _newAmount) public onlyOwner {
        amount = _newAmount;
    }


}