let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));


async function listWallets() {
    try {
        const result = await sendRPC_personal_listWallets();
        return result;
    } catch (e) {
        return e;
    }
}

function sendRPC_personal_listWallets() {

    return new Promise((resolve, reject) => {

        web3.currentProvider.send({ method: "personal_listWallets", params: [], jsonrpc: "2.0", id: new Date().getTime() },
            function (error, result) { if (error) { reject(error); } else { resolve(result); } });
    }
    );
}

listWallets().then(console.log);



