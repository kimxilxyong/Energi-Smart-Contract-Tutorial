#! /usr/bin/env node
const { argv } = require('process');

const { ethers } = require("ethers");

//const provider = new ethers.providers.Web3Provider(window.ethereum)
//const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');

// The provider also allows signing transactions to
// send ether and pay to change state within the blockchain.
// For this, we need the account signer...
//const signer = provider.getSigner();


let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

async function getTransactionsByAccount(myaccount, startBlockNumber, endBlockNumber) {
    if (endBlockNumber == null) {
        endBlockNumber = await web3.eth.getBlockNumber();
        console.log("Using endBlockNumber: " + endBlockNumber);
    }
    if (startBlockNumber == null) {
        startBlockNumber = endBlockNumber - 300000;
        console.log("Using startBlockNumber: " + startBlockNumber);
    }

    myaccount = myaccount.toUpperCase();
    console.log("Searching for transactions to/from account \"" + myaccount + "\" within blocks " + startBlockNumber + " and " + endBlockNumber);

    for (var i = startBlockNumber; i <= endBlockNumber; i++) {
        if (i % 100000 == 0) {
            console.log("Searching block " + i);
        }
        var block = await web3.eth.getBlock(i, true);
        if (block != null && block.transactions != null) {
            block.transactions.forEach(function (e) {
                if ( (e.to && e.from) && ( myaccount == "*" || myaccount == e.from.toUpperCase() || myaccount == e.to.toUpperCase() ) ) {
                    console.log(  "   tx hash         : " + e.hash + "\n"
                                + "   nonce           : " + e.nonce + "\n"
                                + "   blockHash       : " + e.blockHash + "\n"
                                + "   blockNumber     : " + e.blockNumber + "\n"
                                + "   transactionIndex: " + e.transactionIndex + "\n"
                                + "   from            : " + e.from + "\n"
                                + "   to              : " + e.to + "\n"
                                + "   value           : " + e.value + "\n"
                                + "   time            : " + block.timestamp + " " + new Date(block.timestamp * 1000).toGMTString() + "\n"
                                + "   gasPrice        : " + e.gasPrice + "\n"
                                + "   gas             : " + e.gas + "\n"
                                + "   input           : " + e.input);
                }
            })
        }
    }
}

const address = argv[2] || "0x771dDB07222A1f9442C91CF04f64F3164771BB62";
console.log("Info for address ", address);

// Get transaction count of an account...
let sendCount;
web3.eth.getTransactionCount(address).then((c) => { sendCount = c; console.log("getTransactionCount for ", address, " is ", sendCount); });

let test;
test = getTransactionsByAccount(address);



