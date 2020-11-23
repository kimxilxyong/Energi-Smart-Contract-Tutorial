#! /usr/bin/env node

const Web3 = require('web3');
const ethers = require('ethers');
const abi = require('./faucet103.abi.json');

const testPrivateKey = (pk) => {
    try {
        w = new ethers.Wallet(pk);
        console.log(w);
        return true;
    }
    catch (e) {
        console.log('Invalid private key.');
    }
    return false;
}

const signTransaction = async (web3, contract, fromAddress, privateKey) => {
    const tx = {
        // Get the nonce
        nonce: await web3.eth.getTransactionCount(fromAddress),
        // this could be provider.addresses[0] if it exists
        from: fromAddress,
        // target address, this could be a smart contract address
        to: contract._address,
        // optional if you want to specify the gas limit
        gasPrice: "20000000000",
        gas: "2100000",
        // optional if you are invoking say a payable function
        value: 0,
        // this encodes the ABI of the method and the arguements
        data: contract.methods.getVersion().encodeABI(),
    };

    let signedTransaction;
    try {
        if (privateKey && testPrivateKey(privateKey)) {
            // Private Key
            signedTransaction = await web3.eth.accounts.signTransaction(tx, privateKey);
        } else {
            // unlocked Account
            console.log('Signing with unlocked account', tx.from);
            signedTransaction = await web3.eth.signTransaction(tx, tx.from); //.catch((e) => { console.log(e); });
        }

    } catch (e) {
        console.log(e);
        signedTransaction = undefined;
    }
    return signedTransaction;
}

const sendTransaction = async (web3, signedTx) => {
    // raw transaction string may be available in .raw or
    // .rawTransaction depending on which signTransaction
    // function was called
    try {
        const sentTx = await web3.eth.sendSignedTransaction(signedTx.raw || signedTx.rawTransaction);
        if (sentTx) {
            console.log(sentTx);
        }
        return sentTx;

    } catch (e) {
        console.log(e);
        return false;
    }

    /*     sentTx.on("receipt", receipt => {
            // do something when receipt comes back
        });
        sentTx.on("error", err => {
            // do something on transaction error
        }); */

}

const getContractVersion = async (web3, contract, fromAddress, privateKey) => {
    try {
        const signedTX = await signTransaction(web3, contract, fromAddress, privateKey);
        if (signedTX) {
            const sentTX = await sendTransaction(web3, signedTX);
            console.log(sentTX);
            return 1;
        } else {
            console.log("Failed to sign transaction");
        }
    } catch (e) {
        console.log(e);
    }
    return -1;
}

const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

const contractAddr = "0x2a2215Ab0B0ccFcaFF01dD7eB6E5808BD8D50C51";
const fromAddress = "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c";
let privateKey;

const faucet = new web3.eth.Contract(abi, contractAddr);

console.log("calling getVersion(", fromAddress, privateKey);
getContractVersion(web3, faucet, fromAddress, privateKey).catch((e) => { console.log("getVersion failed:", e); }).then((r) => { console.log("Version:", r); });
console.log("waiting for result");


