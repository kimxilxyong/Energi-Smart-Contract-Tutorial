#! /usr/bin/env node

let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

web3.eth.personal.getAccounts().then(console.log);
