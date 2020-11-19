/* SPDX-License-Identifier: MIT
   Copyright (c) 2020 Kim Il Yong */

//import Web3 from 'web3';
import ethers from "ethers";

/*
function getBalanceHistory
param provider:
param address:
param changesNumber: the number of coin balance history changes */
//async function getBalanceHistory(provider, address, changesNumber) {

export const getBalanceHistory = async (provider, address, changesNumber) => {
  trieMessageCount = 0;

  if (!ethers.utils.isAddress(address)) {
    throw ("Invalid address " + address);
  }
  address = ethers.utils.getAddress(address);

  let arrayHistory = [];
  let tempArray = [{},];
  try {
    const latest = await provider.getBlockNumber();
    const balance = await provider.getBalance(address);
    // balance is zero, return
    if (balance && balance.eq(ethers.constants.Zero)) {
      return [{
        timestamp: new Date(),
        balance: 0,
        diff: 0,
        blockNr: latest,
      },];
    }
    let iteration = 0;
    let start;
    let end;
    while (changesNumber > arrayHistory.length && (start > 0 || iteration === 0)) {

      // iteration == 0: go from latest-100 to latest
      // iteration > 0: go from latest-(100*iteration+1) to latest-(100*iteration)
      start = latest - (crawlerBlocks * (iteration + 1));
      end = start + crawlerBlocks;

      if (start < 0) {
        start = 0;
      }

      tempArray = await getBalanceHistoryCrawler(provider, address, start, end);

      if (tempArray.length > 1) {
        tempArray.shift();
        arrayHistory.unshift(tempArray);
        arrayHistory = arrayHistory.flat();
      }
      iteration++;
    }
    while (arrayHistory.length > changesNumber) {
      arrayHistory.shift();
    }

    console.log("Return history: iterations", iteration, "from block", start, "to", latest);

    return arrayHistory;

  } catch (error) {
    console.error("Top level catch:", error);
    throw (error);
  }
}

/*
function getBalanceHistoryBackCrawler
param provider
param address
param blocks number of eth blocks to scan from ent-blocks to latest */
async function getBalanceHistoryBackCrawler(provider, address, blocks) {

  try {
    const end = await provider.getBlockNumber();
    const start = end - blocks;
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

  //console.log("Using startBlockNumber: ", start);
  //console.log("Using endBlockNumber: ", end);

  const arrayBalances = [];
  // fetch all block balances async
  for (let i = start; i <= end; i++) {
    try {
      arrayBalances.push({
        balancePromise: provider.getBalance(address, i).catch((e) => {
          // check for missing trie node
          if (e.toString().includes("missing trie node")) {
            if (trieMessageCount % trieMessageShowEveryXMessage === 0) {
              console.error("🐞 https://github.com/ethereum/go-ethereum/issues/17133", i);
            }
            trieMessageCount++;
            //throw ("🐞 missing trie node https://github.com/ethereum/go-ethereum/issues/17133");
          } else {
            console.error(".Catch1:", e.toString().slice(0, errorLen));
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

  const arrayChanges = [];
  let item;
  let lastBalance = ethers.constants.Zero;
  let balance = ethers.constants.Zero;

  // loop over promises and collect balance changes
  for (let i = 0; i < arrayBalances.length; i++) {
    try {

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
          lastBalance = balance;
          arrayChanges.push(
            {
              balance: balance,
              blockNr: item.blockNr,
              block: provider.getBlockWithTransactions(item.blockNr).catch((e) => {
                // check for missing trie node
                if (e.toString().includes("missing trie node")) {
                  if (trieMessageCount % trieMessageShowEveryXMessage === 0) {
                    console.error("🐞 https://github.com/ethereum/go-ethereum/issues/17133", i);
                  }
                  trieMessageCount++;
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
  // garbage collection
  arrayBalances.length = 0;

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
      let incoming;
      if (blockFull) {

        if (blockFull.number !== b.blockNr) {
          // Bailout, this can not happen (or can it 🇨🇳 🤑 🇺🇸?)
          console.log.error("CRITICAL FAILURE: 🐞 Invalid block number from getBlockWithTransactions(" + b.blockNr + "): " + blockFull.number + " !== " + b.blockNr + ", ABORTING");
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

      arrayBlocks.push({
        timestamp: timestamp,
        balance: parseFloat(ethers.utils.formatEther(b.balance)),
        diff: parseFloat(ethers.utils.formatEther(bndiff)),
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

  console.log("Collected: ", arrayBlocks.length, "Blocks in range", start, end);

  return (arrayBlocks);
}

function sleep(milliseconds) {
  return new Promise(resolve => setTimeout(resolve, milliseconds));
}


/*
// Ethereum
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545');
// Metamask
const provider = new ethers.providers.Web3Provider(window.ethereum)
*/
// my local Energi3 testnet node
/* const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
const address = argv[2] || "0x771dDB07222A1f9442C91CF04f64F3164771BB62";
const transactions = 2000;
 */
// How many blocks to scan in one iteration (bigger number means more memory)
// After an iteration the blocks are deleted and only blocks with effective balance changes are kept
// Also in one iteration each block is fetched asynchronly (at nearly the same time the fetches are fired)
const crawlerBlocks = 200;

/* if you get eth.getBalance('account', 'block-no') returns 'Error: missing trie node' #17133 */
/* add  😜--syncmode=full --gcmode=archive😝  to your node server startup */
/* https://github.com/ethereum/go-ethereum/issues/17133 */
let trieMessageCount = 0;
const trieMessageShowEveryXMessage = (crawlerBlocks * 20); // show only each X's message, because if you get one, you will get thousands


/* These consts make sure we don't overload the server (hopefully 😇)*/
const scanBlock = 100;
const scanBlockSleep = 100; // milliseconds


const startTime = Date.now();
// Get balance changes history for the last transactions count
/* getBalanceHistory(provider, address, transactions).catch((e) => { console.error(e); }).then((h) => {
  console.log('***************************************');
  console.log(h);
console.log('***************************************');
  console.log("History", h.length, "transactions");
  console.log("Elapsed:", Date.now() - startTime, "ms");
  console.log('***************************************');
}); */

/*
const blocks = 300;  //300000;
getBalanceHistoryBackCrawler(provider, address, blocks).catch((e) => { console.log(e); }).then((h) => {
    console.log('***************************************');
    console.log(h);
    console.log('***************************************');
    console.log("History", h.length, "transactions");
    console.log("Elapsed:", Date.now() - startTime, "ms");
    console.log('***************************************');
}); */

export const timeDifference = (current, previous) => {

  const msPerMinute = 60 * 1000;
  const msPerHour = msPerMinute * 60;
  const msPerDay = msPerHour * 24;
  const msPerMonth = msPerDay * 30;
  const msPerYear = msPerDay * 365;

  const elapsed = current - previous;

  if (elapsed < msPerMinute) {
    return Math.round(elapsed / 1000) + ' seconds ago';
  } else if (elapsed < msPerHour) {
    return Math.round(elapsed / msPerMinute) + ' minutes ago';
  } else if (elapsed < msPerDay) {
    return Math.round(elapsed / msPerHour) + ' hours ago';
  } else if (elapsed < msPerMonth) {
    return 'approximately ' + Math.round(elapsed / msPerDay) + ' days ago';
  } else if (elapsed < msPerYear) {
    return 'approximately ' + Math.round(elapsed / msPerMonth) + ' months ago';
  } else {
    return 'approximately ' + Math.round(elapsed / msPerYear) + ' years ago';
  }
}




export const listWallets = async (web3) => {
  try {
    const result = await sendRPC_personal_listWallets(web3);
    return result;
  } catch (e) {
    return e;
  }
};

function sendRPC_personal_listWallets(web3) {
  return new Promise((resolve, reject) => {
    web3.currentProvider.send(
      {
        method: "personal_listWallets",
        params: [],
        jsonrpc: "2.0",
        id: new Date().getTime(),
      },
      function (error, result) {
        if (error) {
          reject(error);
        } else {
          resolve(result);
        }
      }
    );
  });
}
