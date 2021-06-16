// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "./IYam.sol";

contract Mining is Ownable {
    using SafeMath for uint256;
    using SafeERC20 for IERC20;

    // Info of each user.
    struct UserInfo {
        uint256 amount;     // How many LP tokens the user has provided.
        uint256 rewardDebt; // Reward debt.
    }

    // Info of each pool.
    struct PoolInfo {
        IERC20 lpToken;           // Address of LP token contract.
        uint256 allocPoint;       // How many allocation points assigned to this pool. YAMs to distribute per block.
        uint256 lastRewardBlock;  // Last block number that YAMs distribution occurs.
        uint256 accYamPerShare; // Accumulated Yams per share, times 1e12.
        uint256 totalAmount;    // Total amount of current pool deposit.
    }

    // The yam Token!
    IYam public yam;
    // yam tokens created per block.
    uint256 public yamPerBlock = 3 * 1e18;
    // Info of each pool.
    PoolInfo[] public poolInfo;
    // Info of each user that stakes LP tokens.
    mapping(uint256 => mapping(address => UserInfo)) public userInfo;
    // pid corresponding address
    mapping(address => uint256) public LpOfPid;
    // Control mining
    bool public paused = false;
    // Total allocation points. Must be the sum of all allocation points in all pools.
    uint256 public totalAllocPoint = 0;
    // The block number when yam mining starts.
    uint256 public startBlock;

    // ------------------------用于统计用户---------------------------------
    uint256 public userIdCount;
    mapping(address => uint256) public userIdMap;
    address[] public userAddrs;
    // ------------------------用于统计用户---------------------------------

    event Deposit(address indexed user, uint256 indexed pid, uint256 amount);
    event Withdraw(address indexed user, uint256 indexed pid, uint256 amount);
    event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount);

    constructor(IYam _yam) {
        yam = _yam;
        startBlock = block.number;
    }

    // Set the number of yam produced by each block
    function setYamPerBlock(uint256 _newPerBlock) public onlyOwner {
        massUpdatePools();
        yamPerBlock = _newPerBlock;
    }

    function changeYam(IYam _yam) public onlyOwner {
        yam = _yam;
    }

    function poolLength() public view returns (uint256) {
        return poolInfo.length;
    }

    function userAddrsLength() public view returns (uint256) {
        return userAddrs.length;
    }

    function setPause() public onlyOwner {
        paused = !paused;
    }

    // Add a new lp to the pool. Can only be called by the owner.
    // XXX DO NOT add the same LP token more than once. Rewards will be messed up if you do.
    function add(uint256 _allocPoint, IERC20 _lpToken) public onlyOwner {
        require(address(_lpToken) != address(0), "_lpToken is the zero address");
        massUpdatePools();
        uint256 lastRewardBlock = block.number > startBlock ? block.number : startBlock;
        totalAllocPoint = totalAllocPoint.add(_allocPoint);
        poolInfo.push(PoolInfo({
        lpToken : _lpToken,
        allocPoint : _allocPoint,
        lastRewardBlock : lastRewardBlock,
        accYamPerShare : 0,
        totalAmount : 0
        }));
        LpOfPid[address(_lpToken)] = poolLength() - 1;
    }

    // Update the given pool's yam allocation point. Can only be called by the owner.
    function set(uint256 _pid, uint256 _allocPoint) public onlyOwner {
        massUpdatePools();
        totalAllocPoint = totalAllocPoint.sub(poolInfo[_pid].allocPoint).add(_allocPoint);
        poolInfo[_pid].allocPoint = _allocPoint;
    }

    function getYamBlockReward(uint256 _lastRewardBlock) public view returns (uint256) {
        return (block.number.sub(_lastRewardBlock)).mul(yamPerBlock);
    }

    // Update reward variables for all pools. Be careful of gas spending!
    function massUpdatePools() public {
        uint256 length = poolInfo.length;
        for (uint256 pid = 0; pid < length; ++pid) {
            updatePool(pid);
        }
    }

    // Update reward variables of the given pool to be up-to-date.
    function updatePool(uint256 _pid) public {
        PoolInfo storage pool = poolInfo[_pid];
        if (block.number <= pool.lastRewardBlock) {
            return;
        }
        uint256 lpSupply = pool.lpToken.balanceOf(address(this));
        if (lpSupply == 0) {
            pool.lastRewardBlock = block.number;
            return;
        }
        uint256 blockReward = getYamBlockReward(pool.lastRewardBlock);
        if (blockReward <= 0) {
            return;
        }
        uint256 yamReward = blockReward.mul(pool.allocPoint).div(totalAllocPoint);
        bool minRet = yam.mint(address(this), yamReward);
        if (minRet) {
            pool.accYamPerShare = pool.accYamPerShare.add(yamReward.mul(1e12).div(lpSupply));
        }
        pool.lastRewardBlock = block.number;
    }

    // View function to see pending yams on frontend.
    function pending(uint256 _pid, address _user) external view returns (uint256) {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][_user];
        uint256 accYamPerShare = pool.accYamPerShare;
        uint256 lpSupply = pool.lpToken.balanceOf(address(this));
        if (user.amount > 0) {
            if (block.number > pool.lastRewardBlock) {
                uint256 blockReward = getYamBlockReward(pool.lastRewardBlock);
                uint256 yamReward = blockReward.mul(pool.allocPoint).div(totalAllocPoint);
                accYamPerShare = accYamPerShare.add(yamReward.mul(1e12).div(lpSupply));
                return user.amount.mul(accYamPerShare).div(1e12).sub(user.rewardDebt);
            }
            if (block.number == pool.lastRewardBlock) {
                return user.amount.mul(accYamPerShare).div(1e12).sub(user.rewardDebt);
            }
        }
        return 0;
    }

    // Deposit LP tokens to HecoPool for yam allocation.
    function deposit(uint256 _pid, uint256 _amount) public notPause {
        // 统计用户
        if (userIdMap[msg.sender] == 0) {
            userIdCount++;
            userIdMap[msg.sender] = userIdCount;
            userAddrs.push(msg.sender);
        }

        depositYam(_pid, _amount, msg.sender);
    }

    function depositYam(uint256 _pid, uint256 _amount, address _user) private {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][_user];
        updatePool(_pid);
        if (user.amount > 0) {
            uint256 pendingAmount = user.amount.mul(pool.accYamPerShare).div(1e12).sub(user.rewardDebt);
            if (pendingAmount > 0) {
                safeYamTransfer(_user, pendingAmount);
            }
        }
        if (_amount > 0) {
            pool.lpToken.safeTransferFrom(_user, address(this), _amount);
            user.amount = user.amount.add(_amount);
            pool.totalAmount = pool.totalAmount.add(_amount);
        }
        user.rewardDebt = user.amount.mul(pool.accYamPerShare).div(1e12);
        emit Deposit(_user, _pid, _amount);
    }

    // Withdraw LP tokens from HecoPool.
    function withdraw(uint256 _pid, uint256 _amount) public notPause {
        withdrawYam(_pid, _amount, msg.sender);
    }

    function withdrawYam(uint256 _pid, uint256 _amount, address _user) private {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][_user];
        require(user.amount >= _amount, "withdrawYam: not good");
        updatePool(_pid);
        uint256 pendingAmount = user.amount.mul(pool.accYamPerShare).div(1e12).sub(user.rewardDebt);
        if (pendingAmount > 0) {
            safeYamTransfer(_user, pendingAmount);
        }
        if (_amount > 0) {
            user.amount = user.amount.sub(_amount);
            pool.totalAmount = pool.totalAmount.sub(_amount);
            pool.lpToken.safeTransfer(_user, _amount);
        }
        user.rewardDebt = user.amount.mul(pool.accYamPerShare).div(1e12);
        emit Withdraw(_user, _pid, _amount);
    }

    // Withdraw without caring about rewards. EMERGENCY ONLY.
    function emergencyWithdraw(uint256 _pid) public notPause {
        emergencyWithdrawYam(_pid, msg.sender);
    }

    function emergencyWithdrawYam(uint256 _pid, address _user) private {
        PoolInfo storage pool = poolInfo[_pid];
        UserInfo storage user = userInfo[_pid][_user];
        uint256 amount = user.amount;
        user.amount = 0;
        user.rewardDebt = 0;
        pool.lpToken.safeTransfer(_user, amount);
        pool.totalAmount = pool.totalAmount.sub(amount);
        emit EmergencyWithdraw(_user, _pid, amount);
    }

    // Safe yam transfer function, just in case if rounding error causes pool to not have enough yams.
    function safeYamTransfer(address _to, uint256 _amount) internal {
        uint256 yamBal = yam.balanceOf(address(this));
        if (_amount > yamBal) {
            yam.transfer(_to, yamBal);
        } else {
            yam.transfer(_to, _amount);
        }
    }

    modifier notPause() {
        require(paused == false, "Mining has been suspended");
        _;
    }
}