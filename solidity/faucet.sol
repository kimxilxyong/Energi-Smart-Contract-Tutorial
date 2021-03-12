pragma solidity >=0.7.0 <0.8.0;

/**
 * SPDX-License-Identifier: MIT
 * Copyright (c) 2020 Kim Il Yong
 */

// Connect Remix to: http://localhost:49796

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

    // 106 add country
    // 106 published to '0x1EDf7947F7b95bA658D0A74024Dd8092e4D4831c'

    // 107 one more require(_amount <= getBalance(), "not enough contract funds");
    //     set name and country at every donation & request
    //     fixed parent constructor calling
    //     no abi changes
    // 107 published to 0x57B98F76bB39546F97BccD1EF0A38b2d9E074494
    uint16 constant public version = 107;

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
        address payable oldOwner = owner;

        owner = newOwner;
        emit OwnerSet(oldOwner, newOwner);
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

    event ReceivedDonation(address indexed donor, uint amount, uint donationnumber);
    event GivenDonation(address indexed recipient, uint amount, uint donationnumber);
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
        bytes8 country;
        mapping(uint => Payment) payments;
    }

    // Allow a request only every 4 hours
    uint16 constant private graceTimeout = 14400; //60*60*4

    uint private totalBalance;

    address[] private Donors;
    address[] private Recipients;

    mapping(address => Balance) private Donations;
    mapping(address => Balance) private Payments;

    /**
     * @dev Set contract deployer as owner
     */
    constructor() OwnedContract() {
        totalBalance = 0;
    }

    /**
     * @dev Delete the smart contract
     */
    function destroySmartContract() public isOwner {
        selfdestruct(super.getOwner());
    }

    receive() external payable {
        string memory name = "Anonymous Donor";
        bytes8 country = "Terra";
        Donation(name, country);
    }

    fallback() external {
        revert("Have a look at https://github.com/kimxilxyong/Energi-Smart-Contract-Tutorial/blob/master/solidity/faucet.sol");
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
    function getDonorByIndex(uint _i) public view returns (address _from, uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country) {
        require(_i < Donors.length, "Invalid Donor index");

        _from = Donors[_i];
        _totalBalance = Donations[_from].totalBalance;
        _numPayments = Donations[_from].numPayments;
        _name = Donations[_from].name;
        _country = Donations[_from].country;
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
    function getRecipientByIndex(uint _i) public view returns (address _recipient, uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country) {
        require(_i < Recipients.length, "Invalid Recipient index");

        _recipient = Recipients[_i];
        _totalBalance = Payments[_recipient].totalBalance;
        _numPayments = Payments[_recipient].numPayments;
        _name = Payments[_recipient].name;
        _country = Payments[_recipient].country;
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
     * @dev Return Donations from Address country
     * @param  _from Address
     * @return country of donor
     */
    function getDonorCountry(address _from) public view returns (bytes8) {
        return Donations[_from].country;
    }

    /**
     * @dev Return Payments to Address name
     * @param  _to Address
     * @return name of recipient
     */
    function getRecipientName(address _to) public view returns (string memory) {
        return Payments[_to].name;
    }
    /**
     * @dev Return Payments to Address country
     * @param  _to Address
     * @return country of recipient
     */
    function getRecipientCountry(address _to) public view returns (bytes8) {
        return Payments[_to].country;
    }

    /**
     * @dev Sets a donor name
     * @param  _name of donor
     */
    function setDonorName(string calldata _name) public {
        Donations[msg.sender].name = _name;
    }
    /**
     * @dev Sets a donor country
     * @param  _country of donor
     */
    function setDonorCountry(bytes8 _country) public {
        Donations[msg.sender].country = _country;
    }

    /**
     * @dev Sets a recipients name
     * @param  _name of recipient
     */
    function setRecipientName(string calldata  _name) public {
        Payments[msg.sender].name = _name;
    }
    /**
     * @dev Sets a recipients country
     * @param  _country of recipient
     */
    function setRecipientCountry(bytes8  _country) public {
        Payments[msg.sender].country = _country;
    }

    /**
     * @dev Return Donor info
     * @param  _from Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getDonorByAddress(address _from) public view returns (uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country) {
        _totalBalance = getDonorBalance(_from);
        _numPayments = getDonationCounts(_from);
        _name = getDonorName(_from);
        _country = getDonorCountry(_from);
    }

    /**
     * @dev Return Recipient info
     * @param  _to Address
     * @return _totalBalance
     * @return _numPayments
     */
    function getRecipientByAddress(address _to) public view returns (uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country) {
        _totalBalance = getRecipientBalance(_to);
        _numPayments = getPaymentCounts(_to);
        _name = getRecipientName(_to);
        _country = getRecipientCountry(_to);
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
     * @dev called to provide funding to contract
     * @param _name Name of Recipient
     * @param _country emoji flag
     */
    function Donation(string memory _name, bytes8 _country) public payable {
        require(msg.sender != address(0),  "address is null");
        require(msg.value > 0, "amount should be greater than 0");

        uint64 numPayments;
        numPayments = Donations[msg.sender].numPayments.add64(uint64(1));
        assert(numPayments > 0);
        if (numPayments == 1) {
            Donations[msg.sender].index = uint64(Donors.length);
            Donors.push(msg.sender);
        }

        Donations[msg.sender].name = _name;
        Donations[msg.sender].country = _country;

        Donations[msg.sender].totalBalance = Donations[msg.sender].totalBalance.add(msg.value);
        Donations[msg.sender].numPayments = numPayments;
        Donations[msg.sender].payments[numPayments].amount = msg.value;
        Donations[msg.sender].payments[numPayments].withdraw = false;
        Donations[msg.sender].payments[numPayments].timestamp = block.timestamp;

        totalBalance = totalBalance.add(msg.value);

        emit ReceivedDonation(msg.sender, msg.value, numPayments);
    }

    /**
     * @dev Withdraw money from contract
     * @param _to address to send money to
     * @param _amount Amount of money to send
     */
    function withdrawDonation(address payable _to, uint _amount) public {
        require(_to != address(0), "null address supplied");
        require(_amount <= Donations[_to].totalBalance, "withdraw amount higher than your Donations");
        require(_amount <= getBalance(), "not enough contract funds");
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
     * @param _name Name of Recipient
     * @param _country emoji flag
     */
    function requestDonation(address payable _to, uint _amount, string memory _name, bytes8 _country) public {
        require(_to != address(0), "null address supplied");
        require(_amount <= getBalance(), "not enough contract funds");
        // allow only 1/10th of funds to pay out
        require(_amount <= getBalance().div(10), "amount too high, maximum is one tenth of faucet funds");

        //bool AllowTransfer = true;
        uint timeLastTransfer = 0; // time of last request
        uint64 numPayments;
        numPayments = Payments[_to].numPayments.add64(1);
        assert(numPayments > 0);
        if (numPayments == 1) {
          // new Recipient
            Payments[_to].index = uint64(Recipients.length);
            Recipients.push(_to);
        } else if (numPayments == 0) {
            revert("Internal error, integer overflow");
        } else {
            // check if the last payment is more than 12 hours ago
            timeLastTransfer = Payments[_to].payments[numPayments.sub64(1)].timestamp;

            if (timeLastTransfer > block.timestamp.sub( graceTimeout )) {     // 60*60*4
                revert("Only one request per 4 hours is possible");
            }
        }

        Payments[_to].name = _name;
        Payments[_to].country = _country;

        Payments[_to].totalBalance = Payments[_to].totalBalance.add(_amount);
        Payments[_to].numPayments = numPayments;
        Payments[_to].payments[numPayments].amount = _amount;
        Payments[_to].payments[numPayments].withdraw = true;
        Payments[_to].payments[numPayments].timestamp = block.timestamp;

        totalBalance = totalBalance.sub(_amount);
        _to.transfer(_amount);
        emit GivenDonation(_to, _amount, numPayments);
    }

    /**
     * @dev Request bootstrap gas from the faucet
     * @param _to address to send the gas to
     * @param _amount of gas to send
     */
    function giveBootstrapGas(address payable _to, uint _amount) public isOwner {
        require(_amount <= 0.1 ether, "Max 0.1 NRG is more than enough gas!");
        _to.transfer(_amount);
    }
}
