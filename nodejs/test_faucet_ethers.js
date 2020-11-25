#! /usr/bin/env node

/* const signTransaction = async (web3, contract, fromAddress, privateKey) => {
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
        }); *

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
*/

const sendDonation = async (contract, donorName, wei, waitForConfirmations) => {
    let result;
    let lastConfirms = -2;
    try {
        console.log("Calling faucet.Donation(", donorName, wei, ")");
        console.log("Transaction is pending, waiting for execution response ... ");

        const txResponse = await contract.Donation(donorName, { value: wei, });

        console.log("Transaction", txResponse.hash, "was executed, waiting for miners ... ");

        const receiptHandler = (receipt) => {
            if (receipt.confirmations === lastConfirms) {
                // I don't know why each confirmation comes twice, the two receipts are exactly the same, - ignore the second
                return;
            }
            if (lastConfirms === -2) {
                console.log("Transaction was mined in Block", receipt.blockNumber, "and used", receipt.gasUsed.toNumber(),"gas");
                console.log("Waiting for", waitForConfirmations, "confirmations");
                lastConfirms = -1;
                return;
            }
            console.log("Receipt Confirmation:", receipt.confirmations);
            console.log("****************************************");
            if ((receipt.transactionHash !== txResponse.hash) || (receipt.confirmation < lastConfirms)) {
                console.log("Chain reorg!! 🐞 🥵");
            }
            lastConfirms = receipt.confirmations;
        };
        contract.provider.on(txResponse.hash, receiptHandler);

        // Wait for number of confirmations
        txResponse.confirmations = 0;
        const txReceipt = await txResponse.wait(waitForConfirmations);

        contract.provider.removeListener(txResponse.hash, receiptHandler);

        result = txReceipt;

    } catch (e) {
        throw (e);
    }
    return result;
}


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

const execute = async () => {

    // 104 name interface is bytes32
    // 104 published to '0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF'
    // 104 name interface is string
    // 105 published to '0xfC79349137862639A035c71C36fD4d71B2a5D668'
    //const contractAddr = "0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF";
    const contractAddr = "0xfC79349137862639A035c71C36fD4d71B2a5D668";

    //const name = "Anton aus Tirol 🇦🇹 🎿";
    //const fromAddress = "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c";

    //const name = "Алекс из Новосибирска 🇷🇺 🥶";
    //const fromAddress = "0x771ddb07222a1f9442c91cf04f64f3164771bb62";

    //const name = "🇹🇼台湾 is not 💩🇨🇳中国💩";
    //const fromAddress = "0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30";

    const name = "免费西藏你中国混蛋 🤮🤢🤬"; // free tibet
    const fromAddress = "0x3b5abf9b81d5b62df65bf63e163463d0aa42e53b";


    //web3.utils.utf8ToHex(name)

    // how many confirmations do we want to wait for
    const waitForConfirmations = 3;

    // how much do we want to spend
    const spendeToFaucetInETH = "5.9"; // NRG

    let privateKey;
    let currentName;
    const ethers = require('ethers');
    const abi = require('./faucet105.abi.json');
    // If you don't specify a //url//, Ethers connects to the default
    // (i.e. ``http:/\/localhost:8545``)
    const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    const signer = provider.getSigner(fromAddress);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

    // The Contract object
    const faucet = new ethers.Contract(contractAddr, abi, signer);



    let r = await faucet.getVersion().catch((e) => { console.log("getVersion failed:", e); }); //.then((r) => { console.log("Faucet Version:", r); });
    console.log("Faucet Version:", r);

    r = await faucet.getOwner().catch((e) => { console.log("getOwner failed:", e); }); //.then((r) => { console.log("Faucet Owner:", r); });
    console.log("Faucet Owner:", r);
    //callPureContractMethod(faucet, "getOwner()").catch((e) => { console.log("getOwner failed:", e); }).then((r) => { console.log("Owner:", r); });

    r = await faucet.getBalance().catch((e) => { console.log("getBalance failed:", e); }); //.then((r) => { console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r))); });
    console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r)));
    // callPureContractMethod(faucet, "getBalance()").catch((e) => { console.log("getBalance failed:", e); }).then((r) => { console.log("Balance:", r); });

    r = await faucet.getCalculatedBalance().catch((e) => { console.log("getCalculatedBalance failed:", e); }); //.then((r) => { console.log("Faucet CalculatedBalance:", parseFloat(ethers.utils.formatEther(r))); });
    console.log("Faucet CalculatedBalance:", parseFloat(ethers.utils.formatEther(r)));
    //callPureContractMethod(faucet, "getCalculatedBalance()").catch((e) => { console.log("getCalculatedBalance failed:", e); }).then((r) => { console.log("CalculatedBalance:", r); });

    if (name) {
        faucet.getDonorName(fromAddress).catch((e) => { console.log("getDonorName failed:", e); }).then((n) => { currentName = n; });
    }

    console.log(fromAddress + "'" + name + "' is Donating",spendeToFaucetInETH,"NRG");
    sendDonation(faucet, name, ethers.utils.parseEther(spendeToFaucetInETH).toString(), waitForConfirmations)
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
            console.log("*********************");
            console.log("Donation failed with:", e);
            console.log("*********************");
        })
        .then((r) => {
            if (r) {
                if (r.status) {
                    faucet.getCalculatedBalance().catch((e) => { console.log("getCalculatedBalance failed:", e); }).then((r) => { console.log("Faucet Balance is now:", parseFloat(ethers.utils.formatEther(r))); });
                    // Update donors name
                    if (currentName && name && currentName !== name) {
                        // update donor name
                        faucet.setDonorName(name).catch((e) => { console.log("SetName failed:", e); }).then((n) => { console.log("Updated your name from '", currentName, "' to '", name + "'"); });
                    }
                    console.log("****************************************");
                    console.log("Donation to Faucet succesfull 🤑");
                    console.log("Gas used", r.gasUsed.toNumber());
                    console.log("Status", r.status);
                    console.log("Transaction Hash", r.transactionHash);
                    console.log("Block Number", r.blockNumber);
                    console.log("****************************************");
                    console.log("🏁 Finished 🏁");
                    console.log("Thank you'", name, "'for donating", parseFloat(spendeToFaucetInETH) + "NRG");

                } else {
                    console.log("Donation returned status false");
                    console.log(r);
                }
            }
        });

}

execute();
