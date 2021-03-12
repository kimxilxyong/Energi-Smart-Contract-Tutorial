// SPDX - License - Identifier: MIT
// Copyright(c) 2020 Kim Il Yong
// Utility functions to send NRG from to

var i;
var time;
var txHash;

// value is in NRG as a number
function sendNrg(from, to, value) {
    time = new Date();
    //tx = eth.getTransactionReceipt(eth.sendTransaction({from:from, to: to, value: web3.toWei(value,"ether"), gas: "300000"}));
    txHash = eth.sendTransaction({from:from, to: to, value: web3.toWei(value,"ether"), gas: "300000"});
    i = setInterval(waitForConfirmation, 5000);
}

// send all funds, so "from" balance goes to exactly zero (0) wei
function sendAll(from, to) {

    var balance = eth.getBalance(from);
    var gas = new BigNumber(21000);
    var gasPrice = web3.eth.gasPrice;

    var cost = gas.mul(gasPrice);
    var sendAmount = balance.sub(cost);

    time = new Date();

    if (sendAmount < 0) {
        console.log("Insufficient funds, you need atleast '", cost, "' wei");
    } else {

        console.log("Sending wei:", sendAmount, "at GasPrice", gasPrice);

        //tx = eth.getTransactionReceipt(eth.sendTransaction({from: from, to: to, gas: gas, gasPrice: gasPrice, value: sendAmount}));
        console.log("eth.sendTransaction({ from:", from, ", to:", to, ", value:", sendAmount.toString(), ", gas: '21000', })")
        txHash = eth.sendTransaction({ from: from, to: to, value: sendAmount.toString(), gas: "21000", });
        i = setInterval(waitForConfirmation, 5000);
    }
}

function waitForConfirmation() {
    var txReceipt = eth.getTransactionReceipt(txHash);
    if (txReceipt) {
        var tx = eth.getTransaction(txHash);
        if (tx) {
            console.log("Success sending" , web3.fromWei(tx.value), "NRG, from", tx.from, "to", tx.to);
            console.log("Mined in block:", tx.blockNumber);
            console.log("Gas used:", txReceipt.cumulativeGasUsed);
            console.log("Status:", txReceipt.status);
            console.log("Transaction:", txReceipt.transactionHash);
            clearInterval(i);
        }
    } else {
        console.log("Waiting for miners ... ", ( new Date() - time )/1000, "seconds" );
    }
}
