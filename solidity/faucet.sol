pragma solidity >=0.6.0;

/**
 * SPDX-License-Identifier: MIT
 * Copyright (c) 2020 Kim Il Yong 
 */
 
import "./safemath.sol"; 
import "./signedsafemath.sol"; 

/**
 * @title OwnedContract
 * @dev Set & change contract owner
 */
contract OwnedContract {
    
    uint8 constant public version = 101;

    // The creator of this contract
    address private owner;

    // event for EVM logging
    event OwnerSet(address indexed oldOwner, address indexed newOwner);
    
    // modifier to check if caller is owner
    modifier isOwner() {
        // If the first argument of 'require' evaluates to 'false', execution terminates and all
        // changes to the state and to Ether balances are reverted.
        // This used to consume all gas in old EVM versions, but not anymore.
        // It is often a good idea to use 'require' to check if functions are called correctly.
        // As a second argument, you can also provide an explanation about what went wrong.
        require(msg.sender == owner, "Caller is not owner");
        _;
    }
    
    /**
     * @dev Set contract deployer as owner
     */
    constructor() public {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }

    /**
     * @dev Change owner
     * @param newOwner address of new owner
     */
    function changeOwner(address newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @dev Return owner address 
     * @return address of owner
     */
    function getOwner() external view returns (address) {
        return owner;
    }
    
}

/**
 * @title Faucet
 * @dev Donate money to Developers in Need
 */
contract Faucet is OwnedContract {
    
    using SafeMath for uint;
    using SignedSafeMath for int;
    
    event ReceiveDonation(address indexed donor, uint amount, uint donationnumber, uint timestamp);
    event WithdrawMoney(address indexed to, uint indexed amount);
    
    struct Payment {
        uint amount;
        bool withdraw;
        uint timestamp;
    }
    
    struct Balance {
        int totalBalance;
        uint32 numPayments;
        mapping(uint => Payment) payments;
    }

    uint private totalBalance;
    mapping(address => Balance) private Donations;    
    mapping(address => Balance) private Payments;
    
        /**
     * @dev Set contract deployer as owner
     */
    constructor() public {
        totalBalance = 0;
        //super;
    }
    
    
    /**
     * @dev Delete the smart contract
     * @param _to address to send remaining money in the contract to
     */
        /**
     * @dev Delete the smart contract
     * @param _to address to send remaining money in the contract to
     */
    function destroySmartContract(address payable _to) public isOwner {
        require(_to != address(0), "send to address is null");
        selfdestruct(_to);
    }
    
    receive() external payable {
        Donation();
    }
    
     /**
     * @dev called when contract receives funding
     */
    function Donation() public payable {
        require(msg.sender != address(0),  "address is null");
        require((Donations[msg.sender].numPayments + 1 > Donations[msg.sender].numPayments), "uint32 overflow");
        
        uint32 numPayments;
        numPayments = Donations[msg.sender].numPayments + 1;
        
        Donations[msg.sender].totalBalance = Donations[msg.sender].totalBalance.add(int(msg.value));
        Donations[msg.sender].numPayments = numPayments;
        Donations[msg.sender].payments[numPayments].amount = msg.value;
        Donations[msg.sender].payments[numPayments].withdraw = false;
        Donations[msg.sender].payments[numPayments].timestamp = now;
        totalBalance = totalBalance.add(msg.value);
        
        emit ReceiveDonation(msg.sender, msg.value, numPayments, Donations[msg.sender].payments[numPayments].timestamp);
    }
    
    /**
     * @dev Return balance of contract
     * @return balance
     */
    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
    
    /**
     * @dev Return total calculated balance
     * @return uint balance
     */
    function getCalculatedBalance() public view returns (uint) {
        return totalBalance;
    }
    
    /**
     * @dev Return Donations from Address
     * @param  _from Address
     * @return Balance
     */
    function getDonationBalance(address _from) public view returns (int) {
        return Donations[_from].totalBalance;
    }
    
    /**
     * @dev Return Donations count from Address
     * @param  _from Address
     * @return number of donations
     */
    function getDonationCounts(address _from) public view returns (uint32) {
        return Donations[_from].numPayments;
    }

    /**
     * @dev Return Donation details
     * @param  _from Address
     * @param  _nr index
     * @return _amount
     * @return _withdraw
     * @return _timestamp
     */
    function getDonationDetails(address _from, uint32 _nr) public view returns (uint _amount, bool _withdraw, uint _timestamp) {
        _amount = Donations[_from].payments[_nr].amount;
        _withdraw = Donations[_from].payments[_nr].withdraw;
        _timestamp = Donations[_from].payments[_nr].timestamp; 
    }
    

    /**
     * @dev Withdraw money from contract
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function withdrawDonation(address payable _to, uint _amount) public isOwner {
        require(_to != address(0));
        require(int(_amount) <= Donations[_to].totalBalance, "withdraw amount too high");
        require(_amount <= address(this).balance, "not enough contract funds");
        require(Donations[_to].numPayments + 1 > Donations[_to].numPayments, "uint32 overflow");
        uint32 numPayments;
        numPayments = Donations[_to].numPayments + 1;
        
        Donations[_to].totalBalance = Donations[_to].totalBalance.sub(int(_amount));
        Donations[_to].numPayments = numPayments;
        Donations[_to].payments[numPayments].amount = _amount;
        Donations[_to].payments[numPayments].withdraw = true;
        Donations[_to].payments[numPayments].timestamp = now;
        
        totalBalance = totalBalance.sub(_amount);
        
        _to.transfer(_amount);
        emit WithdrawMoney(_to, _amount);
    }
    
}




