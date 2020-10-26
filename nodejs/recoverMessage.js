#! /usr/bin/env node

const { argv } = require('process');
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

const msg = argv[2];
const signature = argv[3];

//web3.eth.accounts.recover(msg, signature).then((a) => { console.log("Signator address: " + a); }).catch((e) => console.log(e.message));

let signator;

console.log("Start web3.eth.accounts.recover(msg, signature, true)");
signator = web3.eth.accounts.recover(msg, signature, true);
console.log("Signator address is: ** " + signator + " **");
console.log("Start web3.eth.accounts.recover(msg, signature, false)");
signator = web3.eth.accounts.recover(msg, signature, false);
console.log("Signator address is: ** " + signator + " **");
console.log("Start web3.eth.accounts.recover(msg, signature)");
signator = web3.eth.accounts.recover(msg, signature);
console.log("Signator address is: ** " + signator + " **");
console.log("End web3.eth.accounts.recover");

console.log("Start web3.eth.personal.ecRecover");
web3.eth.personal.ecRecover(msg, signature).then((a) => { console.log("Signator address is: ** " + a + " **"); }).catch((e) => console.log(e.message));

