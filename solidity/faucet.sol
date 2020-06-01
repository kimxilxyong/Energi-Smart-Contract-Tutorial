pragma solidity ^0.6.0;

/**
 * SPDX-License-Identifier: MIT
 * Copyright (c) 2020 Kim Il Yong 
 */
 
import "./safemath.sol";
import "./signedsafemath.sol";
import "./math.sol";

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
    event GivenDonation(address indexed recipient, uint amount, uint donationnumber, uint timestamp);
    event WithdrawMoney(address indexed to, uint indexed amount);
    
    struct Payment {
        uint amount;
        bool withdraw;
        uint timestamp;
    }
    
    struct Balance {
        int totalBalance;
        uint numPayments;
        string name;
        mapping(uint => Payment) payments;
    }

    uint private totalBalance;
    
    address[] private Donors;
    address[] private Recipients;
    
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
     * @dev Store the address of a donor
     * @param _d address of donor
     * @return index of address
     */
    function addDonor(address _d) internal returns (uint) {
        uint i;
        for(i = 0; i < Donors.length; i++) {
            if (Donors[i] == _d) {
                return i;
            }
        }
        Donors.push(_d);
        return Donors.length - 1;
    }
    
    /**
     * @dev Get the number of donors
     * @return amount of donors
     */
    function getDonorCount() public view returns (uint) {
        return Donors.length;
    }
    
    /**
     * @dev Get the address of a donor
     * @return address of donor
     */
    function getDonorAddress(uint _i) public view returns (address) {
        require(_i < Donors.length, "index _i out of bounds");
        return Donors[_i];
    }

    /**
     * @dev Get the address and name of a donor
     * @param _i index of the  address of a donor
     * @return _from address of donor
     * @return _totalBalance
     * @return _numPayments
     * @return _name name of donor
     */
    function getDonorByIndex(uint _i) public view returns (address _from, int _totalBalance, uint _numPayments, string memory _name) {
        require(_i < Donors.length, "Invalid Donor index");
        
        _from = Donors[_i];
        _totalBalance = Donations[_from].totalBalance;
        _numPayments = Donations[_from].numPayments;
        _name = Donations[_from].name;
    }

    /**
     * @dev Store the address of a recipient
     * @param _r address of recipient
     * @return index of address
     */
    function addRecipient(address _r) internal returns (uint) {
        uint i;
        for(i = 0; i < Recipients.length; i++) {
            if (Recipients[i] == _r) {
                return i;
            }
        }
        Recipients.push(_r);
        return Recipients.length - 1;
    }
    
    /**
     * @dev Get the number of donors
     * @return amount of donors
     */
    function getRecipientsCount() public view returns (uint) {
        return Recipients.length;
    }
    
    /**
     * @dev Get the address of a recipient
     * @return address of recipient
     */
    function getRecipientAddress(uint _i) public view returns (address) {
        require(_i < Recipients.length, "index _i out of bounds");
        return Recipients[_i];
    }

    /**
     * @dev Get the address and name of a recipient
     * @param _i index of the address of a recipient
     * @return _recipient address of recipient
     * @return _totalBalance
     * @return _numPayments
     * @return _name name of recipient
     */
    function getRecipientByIndex(uint _i) public view returns (address _recipient, int _totalBalance, uint _numPayments, string memory _name) {
        require(_i < Recipients.length, "Invalid Recipient index");
        
        _recipient = Recipients[_i];
        _totalBalance = Payments[_recipient].totalBalance;
        _numPayments = Payments[_recipient].numPayments;
        _name = Payments[_recipient].name;
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
    function getDonorBalance(address _from) public view returns (int) {
        return Donations[_from].totalBalance;
    }

    /**
     * @dev Return Donations to Address
     * @param  _to Address
     * @return Balance
     */
    function getRecipientBalance(address _to) public view returns (int) {
        return Payments[_to].totalBalance;
    }
    
    /**
     * @dev Return Donations count from Address
     * @param  _from Address
     * @return number of donations
     */
    function getDonationCounts(address _from) public view returns (uint) {
        return Donations[_from].numPayments;
    }

    /**
     * @dev Return Payments count to Address
     * @param  _to Address
     * @return number of payments
     */
    function getPaymentCounts(address _to) public view returns (uint) {
        return Payments[_to].numPayments;
    }

    /**
     * @dev Return Donations from Address name
     * @param  _from Address
     * @return name of donor
     */
    function getDonorName(address _from) public view returns (string memory) {
        return Donations[_from].name;
    }

    /**
     * @dev Return Donations to Address name
     * @param  _to Address
     * @return name of recipient
     */
    function getRecipientName(address _to) public view returns (string memory) {
        return Payments[_to].name;
    }

    /**
     * @dev Sets a donor name
     * @param  _name of donor
     */
    function setDonorName(string memory _name) public {
        Donations[msg.sender].name = _name;
    }

    /**
     * @dev Sets a recipients name
     * @param  _name of recipient
     */
    function setRecipientName(string memory _name) public {
        Payments[msg.sender].name = _name;
    }

    /**
     * @dev Return Donor info
     * @param  _from Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getDonorByAddress(address _from) public view returns (int _totalBalance, uint _numPayments) {
        _totalBalance = getDonorBalance(_from);
        _numPayments = getDonationCounts(_from);
    }
    
    /**
     * @dev Return Recipient info
     * @param  _to Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getRecipientByAddress(address _to) public view returns (int _totalBalance, uint _numPayments) {
        _totalBalance = getRecipientBalance(_to);
        _numPayments = getPaymentCounts(_to);
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
     * @dev Return Payments details
     * @param  _to Address
     * @param  _nr index
     * @return _amount
     * @return _withdraw
     * @return _timestamp
     */
    function getPaymentDetails(address _to, uint32 _nr) public view returns (uint _amount, bool _withdraw, uint _timestamp) {
        _amount = Payments[_to].payments[_nr].amount;
        _withdraw = Payments[_to].payments[_nr].withdraw;
        _timestamp = Payments[_to].payments[_nr].timestamp; 
    }
    
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
        require(msg.value > 0, "amount should be greater than 0");
        //require((Donations[msg.sender].numPayments + 1 > Donations[msg.sender].numPayments), "uint32 overflow");
        
        uint numPayments;
        numPayments = Donations[msg.sender].numPayments.add(1);
        
        Donations[msg.sender].totalBalance = Donations[msg.sender].totalBalance.add(int(msg.value));
        Donations[msg.sender].numPayments = numPayments;
        Donations[msg.sender].payments[numPayments].amount = msg.value;
        Donations[msg.sender].payments[numPayments].withdraw = false;
        Donations[msg.sender].payments[numPayments].timestamp = now;
        
        addDonor(msg.sender);
        totalBalance = totalBalance.add(msg.value);
        
        emit ReceiveDonation(msg.sender, msg.value, numPayments, now);
    }
        
    /**
     * @dev Withdraw money from contract
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function withdrawDonation(address payable _to, uint _amount) public isOwner {
        require(_to != address(0));
        require(int(_amount) <= Donations[_to].totalBalance, "withdraw amount higher than your Donations");
        require(_amount <= address(this).balance, "not enough contract funds");
        //require(Donations[_to].numPayments + 1 > Donations[_to].numPayments, "uint32 overflow");
        uint numPayments;
        numPayments = Donations[_to].numPayments.add(1);
        
        Donations[_to].totalBalance = Donations[_to].totalBalance.sub(int(_amount));
        Donations[_to].numPayments = numPayments;
        Donations[_to].payments[numPayments].amount = _amount;
        Donations[_to].payments[numPayments].withdraw = true;
        Donations[_to].payments[numPayments].timestamp = now;
        
        totalBalance = totalBalance.sub(_amount);
        
        _to.transfer(_amount);
        emit WithdrawMoney(_to, _amount);
    }
    
    /**
     * @dev Request a donation from the faucet
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function requestDonation(address payable _to, uint _amount) public {
        
        require(_to != address(0), "null address supplied");
        require(_amount <= address(this).balance, "not enough contract funds");

        //bool AllowTransfer = true;
        uint timeLastTransfer = 0; // time of last request
        
        // check if the last payment is more than 24 hours ago
        for(uint i = Payments[_to].numPayments; i > 0; i--) {
            timeLastTransfer = Math.max(timeLastTransfer, Payments[_to].payments[i].timestamp);
        }
        if (timeLastTransfer > 0) {
            if (timeLastTransfer > now.sub(60*60*24)) {
                revert("Only one request per 24 hours is allowed");
            }
        }
        
        // allow only 1/10th of funds to pay out
        require(_amount <= address(this).balance.div(10), "amount too high, maximum is one tenth of faucet funds");
                
        uint numPayments;
        numPayments = Payments[_to].numPayments.add(1);
        
        Payments[_to].totalBalance = Payments[_to].totalBalance.add(int(_amount));
        Payments[_to].numPayments = numPayments;
        Payments[_to].payments[numPayments].amount = _amount;
        Payments[_to].payments[numPayments].withdraw = true;
        Payments[_to].payments[numPayments].timestamp = now;
        
        addRecipient(_to);
        totalBalance = totalBalance.sub(_amount);
        _to.transfer(_amount);
        emit GivenDonation(_to, _amount, numPayments, now);
    }

    /**
     * @dev Request bootstrap gas from the faucet
     * @param _to address to send the gas to
     */
    function requestGas(address payable _to) public isOwner {
        _to.transfer(300000);
    }

}
