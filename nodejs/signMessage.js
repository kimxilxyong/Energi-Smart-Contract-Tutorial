#! /usr/bin/env node

const { argv } = require('process');
let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

const msg = argv[2];
const adr = argv[3];

web3.eth.sign(msg, adr).then((s) => { console.log("Signature is: " + s); }).catch((e) => console.log(e.message));
