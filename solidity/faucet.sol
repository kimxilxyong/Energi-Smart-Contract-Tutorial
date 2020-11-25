pragma solidity ^0.7.5;

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

    // 104 name interface is bytes32
    // 104 published to '0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF'
    // 105 name interface is string
    // 105 published to '0xfC79349137862639A035c71C36fD4d71B2a5D668'
    // http://localhost:49796
    uint16 constant public version = 105;

    // The creator of this contract
    address payable private owner;

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
    constructor() {
        owner = msg.sender; // 'msg.sender' is sender of current call, contract deployer for a constructor
        emit OwnerSet(address(0), owner);
    }

    /**
     * @dev Change owner
     * @param newOwner address of new owner
     */
    function changeOwner(address payable newOwner) public isOwner {
        emit OwnerSet(owner, newOwner);
        owner = newOwner;
    }

    /**
     * @dev Return owner address
     * @return address of owner
     */
    function getOwner() public view returns (address payable) {
        return owner;
    }

    /**
     * @dev Return owner address
     * @return address of owner
     */
    function getVersion() public pure returns (uint16) {
        return version;
    }
}

/**
 * @title Faucet
 * @dev Donate money to Developers in Need
 */
contract Faucet is OwnedContract {

    using SafeMath for uint;
    using SafeMath for uint64;
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
        uint totalBalance;
        uint64 numPayments;
        uint64 index;
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
    constructor() {
        totalBalance = 0;
        //super;
    }

    /**
     * @dev Delete the smart contract
     */
    function destroySmartContract() public isOwner {
        //selfdestruct(super.getOwner());
        selfdestruct(super.getOwner());
    }

    receive() external payable {
        string memory name = "Anonymous Donor";
        Donation(name);
    }

    /**
     * @dev Return the address of the contract
     * @return address
     */
    function getAddress() public view returns (address) {
        return address(this);
    }

    /**
     * @dev Return balance of contract
     * @return balance
     */
    function getBalance() public view returns (uint) {
        return (payable(address(this))).balance;
    }

    /**
     * @dev Return total calculated balance
     * @return uint balance
     */
    function getCalculatedBalance() public view returns (uint) {
        return totalBalance;
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
    function getDonorByIndex(uint _i) public view returns (address _from, uint _totalBalance, uint _numPayments, string memory _name) {
        require(_i < 100, "Max Donor count is 100");
        require(_i < Donors.length, "Invalid Donor index");

        _from = Donors[_i];
        _totalBalance = Donations[_from].totalBalance;
        _numPayments = Donations[_from].numPayments;
        _name = Donations[_from].name;
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
    function getRecipientByIndex(uint _i) public view returns (address _recipient, uint _totalBalance, uint _numPayments, string memory _name) {
        require(_i < Recipients.length, "Invalid Recipient index");

        _recipient = Recipients[_i];
        _totalBalance = Payments[_recipient].totalBalance;
        _numPayments = Payments[_recipient].numPayments;
        _name = Payments[_recipient].name;
    }

    /**
     * @dev Return Donations from Address
     * @param  _from Address
     * @return Balance
     */
    function getDonorBalance(address _from) public view returns (uint) {
        return Donations[_from].totalBalance;
    }

    /**
     * @dev Return Donations to Address
     * @param  _to Address
     * @return Balance
     */
    function getRecipientBalance(address _to) public view returns (uint) {
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
    function setDonorName(string calldata _name) public {
        Donations[msg.sender].name = _name;
    }

    /**
     * @dev Sets a recipients name
     * @param  _name of recipient
     */
    function setRecipientName(string calldata  _name) public {
        Payments[msg.sender].name = _name;
    }

    /**
     * @dev Return Donor info
     * @param  _from Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getDonorByAddress(address _from) public view returns (uint _totalBalance, uint _numPayments) {
        _totalBalance = getDonorBalance(_from);
        _numPayments = getDonationCounts(_from);
    }

    /**
     * @dev Return Recipient info
     * @param  _to Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getRecipientByAddress(address _to) public view returns (uint _totalBalance, uint _numPayments) {
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
     * @dev called when contract receives funding
     */
    function Donation(string memory _name) public payable {
        require(msg.sender != address(0),  "address is null");
        require(msg.value > 0, "amount should be greater than 0");
        //require((Donations[msg.sender].numPayments + 1 > Donations[msg.sender].numPayments), "uint32 overflow");

        uint64 numPayments;
        numPayments = Donations[msg.sender].numPayments.add64(uint64(1));
        if (numPayments == 1) {
            Donations[msg.sender].index = uint64(Donors.length);
            Donations[msg.sender].name = _name;
            Donors.push(msg.sender);
        }

        Donations[msg.sender].totalBalance = Donations[msg.sender].totalBalance.add(msg.value);
        Donations[msg.sender].numPayments = numPayments;
        Donations[msg.sender].payments[numPayments].amount = msg.value;
        Donations[msg.sender].payments[numPayments].withdraw = false;
        Donations[msg.sender].payments[numPayments].timestamp = block.timestamp;

        //addDonor(msg.sender);
        totalBalance = totalBalance.add(msg.value);

        emit ReceiveDonation(msg.sender, msg.value, numPayments, block.timestamp);
    }

    /**
     * @dev Withdraw money from contract
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function withdrawDonation(address payable _to, uint _amount) public {
        require(_to != address(0));
        require(_amount <= Donations[_to].totalBalance, "withdraw amount higher than your Donations");
        require(_amount <= address(this).balance, "not enough contract funds");
        //require(Donations[_to].numPayments + 1 > Donations[_to].numPayments, "uint32 overflow");
        uint64 numPayments;
        numPayments = Donations[_to].numPayments.add64(uint64(1));

        Donations[_to].totalBalance = Donations[_to].totalBalance.sub(_amount);
        Donations[_to].numPayments = numPayments;
        Donations[_to].payments[numPayments].amount = _amount;
        Donations[_to].payments[numPayments].withdraw = true;
        Donations[_to].payments[numPayments].timestamp = block.timestamp;

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
        // allow only 1/10th of funds to pay out
        require(_amount <= address(this).balance.div(10), "amount too high, maximum is one tenth of faucet funds");

        //bool AllowTransfer = true;
        uint timeLastTransfer = 0; // time of last request
        uint64 numPayments;
        numPayments = Payments[_to].numPayments.add64(1);
        if (numPayments == 1) {

            // new Recipient
            Payments[_to].index = uint64(Recipients.length);
            Recipients.push(_to);

        } else {

            // check if the last payment is more than 12 hours ago
            for(uint i = Payments[_to].numPayments; i > 0; i--) {
                timeLastTransfer = Math.max(timeLastTransfer, Payments[_to].payments[i].timestamp);
            }
            if (timeLastTransfer > block.timestamp.sub(60*60*12)) {
                revert("Only one request per 12 hours is possible");
            }
        }

        Payments[_to].totalBalance = Payments[_to].totalBalance.add(_amount);
        Payments[_to].numPayments = numPayments;
        Payments[_to].payments[numPayments].amount = _amount;
        Payments[_to].payments[numPayments].withdraw = true;
        Payments[_to].payments[numPayments].timestamp = block.timestamp;

        totalBalance = totalBalance.sub(_amount);
        _to.transfer(_amount);
        emit GivenDonation(_to, _amount, numPayments, block.timestamp);
    }

    /**
     * @dev Request bootstrap gas from the faucet
     * @param _to address to send the gas to
     */
    function giveBootstrapGas(address payable _to, uint _amount) public isOwner {
        require(_amount <= 1 ether, "1 NRG is more than enough gas!");
        _to.transfer(_amount);
    }
}
