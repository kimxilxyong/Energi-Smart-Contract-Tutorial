#! /usr/bin/env node

const { argv } = require('process');
const { ethers } = require("ethers");
const { SSL_OP_EPHEMERAL_RSA } = require('constants');
const { throws } = require('assert');

/*
function getBalanceHistory
param provider:
param address:
param changesNumber: the number of coin balance history changes */
async function getBalanceHistory(provider, address, changesNumber) {
    const crawlerBlocks = 100;

    if (!ethers.utils.isAddress(address)) {
        throw ("Invalid address " + address);
    }
    address = ethers.utils.getAddress(address);

    let arrayHistory = [];
    let tempArray = [{},];
    try {
        let latest = await provider.getBlockNumber();
        let iteration = 0;
        let start;
        let end;
        while (changesNumber > arrayHistory.length && tempArray.length > 0) {

            // iteration == 0: go from latest-100 to latest
            // iteration > 0: go from latest-(100*iteration+1) to latest-(100*iteration)
            start = latest - (crawlerBlocks * (iteration + 1));
            end = start + crawlerBlocks;

            tempArray = await getBalanceHistoryCrawler(provider, address, start, end);
            // add results on front
/*             if (  !(arrayHistory.length + tempArray.length >= changesNumber)  ) {
                tempArray.shift();
                console.log("tempArray.shift()", tempArray);
            }
 */
            if (tempArray.length > 1) {
                tempArray.shift();
                arrayHistory.unshift(tempArray);
                arrayHistory = arrayHistory.flat();
            }
        iteration++;
        }

        return arrayHistory;

    } catch (error) {
        throw (error);
    }
}

/*
function getBalanceHistoryBackCrawler
param provider
param address
param iteration how far to start back
param blocks number of eth blocks to scan from start */
async function getBalanceHistoryBackCrawler(provider, address, iteration, blocks) {

    try {
        let end = await provider.getBlockNumber();
        end = end - (iteration * blocks);
        let start = end - blocks;
        return (await getBalanceHistoryCrawler(provider, address, start, end));

    } catch (e) {
        console.log("Failed to get current blocknumber, exiting");
        console.log(e);
        throw (e);
    }
}

/*
function getBalanceHistoryCrawler
param provider
param address
param start eth start block
param end eth end block */
async function getBalanceHistoryCrawler(provider, address, start, end) {
    const errorLen = 280; // max error output length

    console.log("Using startBlockNumber: ", start);
    console.log("Using endBlockNumber: ", end);

    const arrayBalances = [];
    // fetch all block balances async
    for (let i = start; i <= end; i++) {
        try {
            arrayBalances.push({
                balancePromise: provider.getBalance(address, i).catch((e) => {
                    // check for missing trie node
                    if (e.toString().includes("missing trie node")) {
                        console.error("🐞 https://github.com/ethereum/go-ethereum/issues/17133", i);
                        //throw ("🐞 missing trie node https://github.com/ethereum/go-ethereum/issues/17133");
                        //return undefined;
                    } else {
                        console.error(".Catch:", e.toString().slice(0, errorLen));
                    }
                    return undefined;
                }),
                blockNr: i,
            });
            if (i % scanBlock === 0) {
                await sleep(scanBlockSleep);
            }
        } catch (e) {
            if (e.toString().includes("missing trie node")) {
                throw (e);
            } else {
                console.log("GetBalance push Promise error: ", e.toString().slice(0, errorLen), " for BlockIndex ", i);
            }
        }
    }
    console.log("DEBUG - 1 arrayBalances.length", arrayBalances.length);

    const arrayChanges = [];
    let item;
    let lastBalance = ethers.constants.Zero;
    let balance = ethers.constants.Zero;
    console.log("DEBUG - 1x", balance);

    // loop over promises and collect balance changes
    for (let i = 0; i < arrayBalances.length; i++) {
    //while (arrayBalancePromises.length > 0) {
        try {

            //balance = await arrayBalances.shift();
            item = arrayBalances[i];
            try {
                //console.log("item.balancePromise", item.balancePromise);
                balance = await item.balancePromise;
                //console.log("Balance:", balance)
            } catch (e) {
                console.error("await", i, "item.balancePromise:", e.toString().slice(0, errorLen));
                balance = lastBalance;
            }
            if (balance) {
                if (!balance.eq(lastBalance)) {
                    console.log("Balance change at block", item.blockNr);
                    console.log("Balance change old", lastBalance);
                    console.log("Balance change new", balance);
                    lastBalance = balance;
                    arrayChanges.push(
                        {
                            balance: balance,
                            blockNr: item.blockNr,
                            block: provider.getBlockWithTransactions(item.blockNr).catch((e) => {
                                // check for missing trie node
                                if (e.toString().includes("missing trie node")) {
                                    console.error("🐞 https://github.com/ethereum/go-ethereum/issues/17133", i);
                                    //throw ("🐞 missing trie node https://github.com/ethereum/go-ethereum/issues/17133");
                                } else {
                                    console.error(".Catch getBlock:", e.toString().slice(0, errorLen));
                                }
                                return undefined;
                            }),
                        }
                    );
                }
            }

        } catch (e) {
            if (e.toString().includes("missing trie node")) {
                throw (e);
            } else {
                console.error("getBlock push Promise error:", e.toString().slice(0, errorLen), "for BlockIndex", i);
            }
        }
    }
    console.log("DEBUG - 1xx");
    console.log("DEBUG - 2 arrayBalances.length:", arrayBalances.length);
    // garbage collection
    arrayBalances.length = 0;
    console.log("DEBUG - 2 arrayBalances.length:", arrayBalances.length);
    console.log("DEBUG - 3 arrayChanges.length:", arrayChanges.length);

    let promiseBlock = "TODO provider.getBlockWithTransactions(b.blockNr)";
    const arrayBlocks = [];
    let found = false;
    let to = ethers.constants.AddressZero;
    let from = ethers.constants.AddressZero;
    let thash = ethers.constants.HashZero;
    let timestamp = new Date(0);
    let value = ethers.constants.Zero;

    // loop over balance changes and get transaction from block

    if (arrayChanges.length > 0) {
        lastBalance = arrayChanges[0].balance;
    }
    let blockFull;
    let bndiff;
    for (let i = 0; i < arrayChanges.length; i++) {

        // get the block at which a balance change happened
        let b = arrayChanges[i];

        try {

            // wait for promise from provider.getBlockWithTransactions(b.blockNr);
            blockFull = await b.block;

            if (blockFull) {

                if (blockFull.number !== b.blockNr) {
                    // Bailout, this can not happen (or can it 🇨🇳 🤑 🇺🇸?)
                    throw ("CRITICAL FAILURE: 🐞 Invalid block number from getBlockWithTransactions(" + b.blockNr + "): " + blockFull.number + " !== " + b.blockNr + ", ABORTING");
                }
                timestamp = new Date(blockFull.timestamp * 1000);

                // how much value has moved (positive/negative), incoming: + > 0 / outgoing: - < 0
                bndiff = b.balance.sub(lastBalance);
                found = false;
                // shortify, t instead blockFull.transactions
                const t = blockFull.transactions;
                for (let i = 0; i < t.length; i++) {

                    to = ethers.utils.getAddress(t[i].to);
                    from = ethers.utils.getAddress(t[i].from);
                    thash = t[i].hash;
                    value = t[i].value;
                    incoming = false;

                    if (to === address) {
                        found = true;
                        incoming = true;
                    } else if (from === address) {
                        found = true;
                    } else if (bndiff.eq(value)) {
                        found = true;
                        if (bndiff.gt(ethers.constants.Zero)) { // greater zero, value > 0
                            incoming = true;
                        }
                    }
                    if (found) {
                        break;
                    }
                }
            } else {
                try {
                    const blockOnly = await provider.getBlock(b.blockNr);
                    timestamp = new Date(blockOnly.timestamp * 1000);
                    value = blockOnly.gasUsed;
                } catch (error) {
                    timestamp = new Date(-62167219200000);
                    value = ethers.constants.Zero;
                }
            }
            //console.log("DEBUG timestamp", timestamp);
            arrayBlocks.push({
                utc: timestamp.toUTCString(),
                balance: parseFloat( ethers.utils.formatEther(b.balance) ),
                diff: parseFloat( ethers.utils.formatEther(bndiff) ),
                blockNr: b.blockNr,
                block: {
                    timestamp: timestamp,
                    to: to,
                    from: from,
                    value: value,
                    incoming: incoming,
                    isInternal: !found,
                },
            });
            lastBalance = b.balance;

        } catch (e) {
            console.log("getBlockWithTransactions error:", e, " for BlockNumber ", b.blockNr);
        }
    }
    arrayChanges.length = 0;

    console.log("Collection finished: ", arrayBlocks.length, " Blocks");

    return (arrayBlocks);
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

/*
// Ethereum
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545');
// Metamask
const provider = new ethers.providers.Web3Provider(window.ethereum)
*/
// my local Energi3 testnet node
const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
const address = argv[2] || "0x771dDB07222A1f9442C91CF04f64F3164771BB62";
const blocks = 300;  //300000;
const transactions = 20;

/* if you get eth.getBalance('account', 'block-no') returns 'Error: missing trie node' #17133 */
/* add  😜--syncmode=full --gcmode=archive😝  to your node server startup */
/* https://github.com/ethereum/go-ethereum/issues/17133 */

/* These consts make sure we don't overload the server (hopefully 😇)*/
const scanBlock = 25;
const scanBlockSleep = 9; // milliseconds


/* crawl back number of blocks to scan for balance changes */
/* console.log("Starting collection for address", address, "from the last", blocks, "blocks"); */
/* getBalanceHistoryBackCrawler(provider, address, blocks).catch(() => { }).then((h) => { console.log("History", h); }); */

const startTime = Date.now();
// get balance changes history for the last transactions count
getBalanceHistory(provider, address, transactions).catch((e) => { console.log(e); }).then((h) => {
    console.log('***************************************');
    console.log(h);
    console.log('***************************************');
    console.log("History", h.length, "transactions");
    console.log("Elapsed:", Date.now() - startTime, "ms");
    console.log('***************************************');
});

return (0);
