pragma solidity >=0.5.0;
pragma experimental ABIEncoderV2;

/**
 * SPDX-License-Identifier: MIT
 * Copyright (c) 2020 Kim Il Yong 
 */

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
    
     /**
     * @dev called when contract receives funding
     */
    function Donation() public payable {
        
        totalBalance = totalBalance + msg.value;
        uint32 numPayments = Donations[msg.sender].numPayments + 1;
        
        Donations[msg.sender].totalBalance += int(msg.value);
        Donations[msg.sender].numPayments = numPayments;
        Donations[msg.sender].payments[numPayments].amount = msg.value;
        Donations[msg.sender].payments[numPayments].withdraw = false;
        Donations[msg.sender].payments[numPayments].timestamp = now;
        
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
     * @dev Withdraw money from contract
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function withdrawMoney(address payable _to, uint _amount) public isOwner {
        require(_amount <= address(this).balance, "not enough funds");
        
        totalBalance = totalBalance -_amount;
        uint32 numPayments = Donations[_to].numPayments + 1;
        
        Donations[_to].totalBalance -= int(_amount);
        Donations[_to].numPayments = numPayments;
        Donations[_to].payments[numPayments].amount = _amount;
        Donations[_to].payments[numPayments].withdraw = true;
        Donations[_to].payments[numPayments].timestamp = now;
        
        _to.transfer(_amount);
        emit WithdrawMoney(_to, _amount);
    }
    
}




