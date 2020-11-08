#! /usr/bin/env node

let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));

/* async thenable function to get all wallets from the server
   function listWallets(web3).then(console.log)
*/
const listWallets = async (web3) => {
    try {
        const result = await sendRPC_personal_listWallets(web3);
        return result;
    } catch (e) {
        return e;
    }
};

/* Make the JSON RPC call 'personal_listWallets' */
function sendRPC_personal_listWallets(web3) {
    return new Promise((resolve, reject) => {
        const id = new Date().getTime();
        web3.currentProvider.send(
            {
                method: "personal_listWallets",
                params: [],
                jsonrpc: "2.0",
                id: id,
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

let accounts;
let wallets;
web3.eth.personal.getAccounts().then((a => { accounts = a; console.log(accounts); }));
console.log("Called getAccounts");
listWallets(web3).then((w) => { wallets = w; console.log(wallets); }).catch(console.error);
console.log("Called listWallets");
