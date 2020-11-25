#! /usr/bin/env node

const Web3 = require('web3');
const ethers = require('ethers');

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
        throw (e);
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
        return result;
    } catch (e) {
        throw (e);
        //return e;
    }
}

const sendDonation = async (contract, donorName, wei) => {
    let result;
    try {
        const options = { gas: "1000000", value: wei, from: contract.defaultAddress, };
        const ta = contract.methods.Donation(donorName).send(options);
        ta.on('transactionHash', (hash) => { console.log("TransactionHash:", hash); })
            .on('confirmation', (confirmationNumber, receipt) => { console.log("Confirmation Nr:", confirmationNumber); })
            .on('error', (e) => { lastErrorObject = e; })
            .on('receipt', (receipt) => { console.log("Receipt:", receipt); })
            .catch((e) => { return e;});  //console.log("Send Promise:", e); });

        result = await ta;

    } catch (e) {
        throw (e);
        //result = e;
    }
    return result;
}



// 104 name interface is bytes32
// 104 published to '0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF'
// 104 name interface is string
// 105 published to '0xfC79349137862639A035c71C36fD4d71B2a5D668'
//const contractAddr = "0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF";
const contractAddr = "0xfC79349137862639A035c71C36fD4d71B2a5D668";
const fromAddress = "0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30";
//const fromAddress = "0x771ddb07222a1f9442c91cf04f64f3164771bb62";
//const fromAddress = "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c";
//const fromAddress = "0x03fb09251ec05ee9ca36c98644070b89111d4b3f";

//const name = "Алекс из Новосибирска 🇷🇺 🥶";
//const name = "Anton aus Tirol 🇦🇹 🎿";
//const name = "🇹🇼台湾 is not 💩🇨🇳中国💩";
const name = "免费西藏你中国混蛋 🤬"; // free tibet
//web3.utils.utf8ToHex(name)

const txConfirmationCount = 5;

let privateKey;
let lastErrorObject;

const abi = require('./faucet105.abi.json');
const web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));
const faucet = new web3.eth.Contract(abi, contractAddr);

faucet.defaultAddress = fromAddress;
faucet.transactionConfirmationBlocks = txConfirmationCount;

callPureContractMethod(faucet, "getVersion()").catch((e) => { console.log("getVersion failed:", e); }).then((r) => { console.log("Version:", r); });
callPureContractMethod(faucet, "getOwner()").catch((e) => { console.log("getOwner failed:", e); }).then((r) => { console.log("Owner:", r); });
//callPureContractMethod(faucet, "getBalance()").catch((e) => { console.log("getBalance failed:", e); }).then((r) => { console.log("Balance:", r); });
callPureContractMethod(faucet, "getCalculatedBalance()").catch((e) => { console.log("getCalculatedBalance failed:", e); }).then((r) => { console.log("CalculatedBalance:", r); });
//getContractVersion(web3, faucet, fromAddress, privateKey).catch((e) => { console.log("getVersion failed:", e); }).then((r) => { console.log("Version:", r); });

console.log("Donating 1 NRG");
sendDonation(faucet, name, web3.utils.toWei("1"))
    .catch((e) => {
        if (e.toString().includes("authentication needed")) {
            console.log("***********************************************************************************************");
            if (e.toString().includes("only staking")) {
                console.log("Your account is in staking only mode, you need to unlock it in the nodes console.");
            } else {
                console.log("Your account is locked, you need to unlock it in the nodes console.");
            }
            console.log("personal.unlockAccount('" + fromAddress + "',null,secondsToUnlock,false)");
            console.log("***********************************************************************************************");
        }
        console.log("**************************");
        console.log("**** Donation failed with:", e);
        console.log("**************************");
    })
    .then((r) => {
        if (r) {
            if (r.status) {
                console.log("Donation succesfull 🤑");
                console.log("Transaction Hash", r.transactionHash);
                console.log("Gas used", r.gasUsed);
                console.log("Status", r.status);
                console.log("Block Number", r.blockNumber);
                console.log("****************************************");
                console.log("🧐 Waiting for " + faucet.transactionConfirmationBlocks + " more confirmations 🏁");

            } else {
                console.log("Donation returned status false");
                console.log(r);
            }
        }
    });
console.log("waiting for result");
