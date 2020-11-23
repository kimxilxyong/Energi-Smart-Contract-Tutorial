#! /usr/bin/env node

const Web3 = require('web3');
const ethers = require('ethers');
const abi = require('./faucet104.abi.json');
const { result } = require('underscore');

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

const getContractVersionSigned = async (web3, contract, fromAddress, privateKey) => {
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

const callPureContractMethod = async (contract, method) => {
    try {
        const result = await contract.methods[method]().call();
        if (result) {
            return result;
        }
    } catch (e) {
        console.log(e);
        return e;
    }
}

const sendDonation = async (contract, donorName, wei) => {
    let result;
    try {
        const options = { gas: "1000000", value: wei, from: contract.defaultAddress, };
        const ta = contract.methods.Donation(donorName).send(options);
        ta.on('transactionHash', (hash) => { console.log("TransactionHash:", hash); })
            .on('confirmation', (confirmationNumber, receipt) => { console.log("Confirmation Nr:", confirmationNumber); })
            .on('error', console.error)
            .on('receipt', (receipt) => { console.log("Receipt:", receipt); })
            .catch((e) => { console.log("Send Promise:", e); });

        result = await ta;

    } catch (e) {
        result = e;
    }
    return result;
}


const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

const contractAddr = "0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF";
//const fromAddress = "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c";
const fromAddress = "0x03fb09251ec05ee9ca36c98644070b89111d4b3f";
let privateKey;

const faucet = new web3.eth.Contract(abi, contractAddr);
faucet.defaultAddress = fromAddress;


callPureContractMethod(faucet, "getVersion()").catch((e) => { console.log("getVersion failed:", e); }).then((r) => { console.log("Version:", r); });
callPureContractMethod(faucet, "getOwner()").catch((e) => { console.log("getOwner failed:", e); }).then((r) => { console.log("Owner:", r); });
//callPureContractMethod(faucet, "getBalance()").catch((e) => { console.log("getBalance failed:", e); }).then((r) => { console.log("Balance:", r); });
callPureContractMethod(faucet, "getCalculatedBalance()").catch((e) => { console.log("getCalculatedBalance failed:", e); }).then((r) => { console.log("CalculatedBalance:", r); });
//getContractVersion(web3, faucet, fromAddress, privateKey).catch((e) => { console.log("getVersion failed:", e); }).then((r) => { console.log("Version:", r); });

console.log("Donating 1 NRG");
sendDonation(faucet, web3.utils.utf8ToHex("Kim from 🇦🇹"), web3.utils.toWei("1")).catch((e) => { console.log("sendDonation failed:", e); }).then((r) => { console.log("sendDonation:", r); });
console.log("waiting for result");
