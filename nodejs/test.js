#! /usr/bin/env node

const { argv } = require('process');
const { ethers } = require("ethers");
const { SSL_OP_EPHEMERAL_RSA } = require('constants');

async function getBalanceHistory(provider, address, blocks) {

    let endBlockNumber;
    try {
        endBlockNumber = await provider.getBlockNumber();
    } catch (e) {
        console.log("Failed to get current blocknumber, exiting");
        console.log(e);
        return (-1);
    }

    let startBlockNumber = endBlockNumber - blocks;
    console.log("Using endBlockNumber: " + endBlockNumber);
    console.log("Using startBlockNumber: " + startBlockNumber);

    const arrayBalances = [];
    // fetch all block balances async
    for (let i = startBlockNumber; i <= endBlockNumber; i++) {
        try {
            arrayBalances.push({
                balancePromise: provider.getBalance(address, i),
                block: i,
            });
            if (i % 20 === 0) {
                await sleep(200);
            }
        } catch (error) {
            console.log("GetBalance push Promise error:", e, " for BlockIndex ", i);
        }
    }
    console.log("DEBUG - 1");

    return (arrayBalances);
}

function sleep(milliseconds) {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
}

async function TestSleep() {
    let i = 40;
    console.log("Sleep");
    if (i % 20 === 0) {
        await sleep(5000);
    }
    console.log("Slept 5 seconds");
    return (1);
}

//const provider = new ethers.providers.Web3Provider(window.ethereum)
const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
const address = argv[2] || "0x771dDB07222A1f9442C91CF04f64F3164771BB62";
const blocks = 3;  //300000;

console.log("Starting collection for address", address, "from the last", blocks, "blocks");

let arrayCoinHistory;

getBalanceHistory(provider, address, blocks).catch(() => { }).then(console.log);

console.log("arrayCoinHistory");
console.log(arrayCoinHistory);

return (0);


