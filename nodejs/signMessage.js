#! /usr/bin/env node

const { argv } = require('process');
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

const msg = argv[2]; // Message
const adr = argv[3] || "0x771ddb07222a1f9442c91cf04f64f3164771bb62"; // address 0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c

web3.eth.sign(msg, adr).then((s) => { console.log("Signature is: " + s); }).catch((e) => console.log(e.message));
