#! /usr/bin/env node

const execute = async () => {

  // 104 name interface is bytes32
  // 104 published to '0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF'
  //const contractAddr = "0x4b7c29e9c7e5132B140884A9dAc98427dcF97AbF";
  //const abi = require('./faucet104.abi.json');

  // 105 name interface is string
  // 105 published to '0xfC79349137862639A035c71C36fD4d71B2a5D668'
  //const contractAddr = "0xfC79349137862639A035c71C36fD4d71B2a5D668";
  //const abi = require('./faucet105.abi.json');

  // 106 add country
  // 106 published to '0x1EDf7947F7b95bA658D0A74024Dd8092e4D4831c'
  const contractAddr = "0x1EDf7947F7b95bA658D0A74024Dd8092e4D4831c";
  const abi = require('./faucet106.abi.json');
  const ethers = require('ethers');

  const fromAddress = "0x3b5abf9b81d5b62df65bf63e163463d0aa42e53b";
  const toAddress = "0xdc67d898cc5fdd982308d1cb0c2301913f8f438c";

  // amount in NRG
  const amountNRG = "1.2"; // NRG

  const name = "Elon から Pretoria 🚗🚀🥠";
  const country = "🇿🇦";



  // value part of transaction
  const wei = ethers.utils.parseEther(amountNRG);

  const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');

  // The provider also allows signing transactions to
  // send ether and pay to change state within the blockchain.
  // For this, we need the account signer...
  //const signer = provider.getSigner(fromAddress);
  //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

  // !!! No signer, we want the raw transaction to pass to webwallet, which will do the signing

  // The Contract object
  const faucet = new ethers.Contract(contractAddr, abi, provider);

  try {

    let b = await faucet.getBalance().catch((e) => { throw(e); }); //.then((r) => { console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r))); });
    console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(b)));

    b = await faucet.getCalculatedBalance().catch((e) => { throw(e); }); //.then((r) => { console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r))); });
    console.log("Faucet Calculated Balance:", parseFloat(ethers.utils.formatEther(b)));

    // debugging: trigger error
    //toAddress = "";                                           // reason: network does not support ENS, code: UNSUPPORTED_OPERATION
    //toAddress = "not an address";                             // reason: network does not support ENS, code: UNSUPPORTED_OPERATION
    //toAddress = "0x";                                         // reason: invalid address, code: INVALID_ARGUMENT
    //toAddress = "0x0";                                        // reason: invalid address, code: INVALID_ARGUMENT
    //toAddress = "0xdc67d898cc5fdd982308d1cb0c2301913f8f438C"; // reason: bad address checksum, code: INVALID_ARGUMENT
    //toAddress = "0x0000000000000000000000000000000000000000"; // reason: null address supplied, code: CALL_EXCEPTION -- sc was called 🥳
    //toAddress = "0x0000000000000000000000000000000000000001"; // reason: amount too high, maximum is one tenth of faucet funds, code: CALL_EXCEPTION


    // Ask for donation, but do not really execute
    // function requestDonation(address payable _to, uint _amount, string memory _name, bytes8 _country) public {
    let r = await faucet.callStatic.requestDonation(toAddress, wei, name, ethers.utils.toUtf8Bytes(country)).catch((e) => { throw(e); });
    if (r) {
      console.log("requestDonation would execute, getting tx data ...");
      let unsignedTx = await faucet.populateTransaction.requestDonation(toAddress, wei, name, ethers.utils.toUtf8Bytes(country)).catch((e) => { throw(e); });
      if (unsignedTx) {
        console.log(unsignedTx);
        console.log("Building webwallet URL for requesting", amountNRG, "NRG to account", toAddress);

        let url = "https://wallet.test3.energi.network/account/send/?gasLimit=3000000&value=0";
        url += "&to=" + contractAddr;
        url += "&data=" + unsignedTx.data;
        url += "&warnings=0&ts=" + Math.floor(new Date() / 1000);

        console.log(url);

      /* https://wallet.energi.network/account/send/?
      gasLimit=3000000
      &value=0
      &to=0x0000000000000000000000000000000000000309
      &data=0x6112fe2e00000000000000000000000000000000000000000000003635c9adc5dea00000
      &warnings=0
      &ts=1608741991354 */

      }
    } else {
      console.log("requestDonation would NOT execute, ABORT");
    }

    // Send donation, but do not really execute
    // function Donation(string memory _name, bytes8 _country) public payable {
    let d = await faucet.callStatic.Donation(name, ethers.utils.toUtf8Bytes(country)).catch((e) => { throw(e); });
    if (d) {
      console.log("Donation would execute, getting tx data ...");
      let unsignedTx = await faucet.populateTransaction.Donation(name, ethers.utils.toUtf8Bytes(country)).catch((e) => { throw(e); });
      if (unsignedTx) {
        console.log(unsignedTx);
        console.log("Building webwallet URL for donating", amountNRG, "NRG to faucet at", contractAddr, "from", name, "in", country);

        let url = "https://wallet.test3.energi.network/account/send/?gasLimit=3000000&value=" + wei;
        url += "&to=" + contractAddr;
        url += "&data=" + unsignedTx.data;
        url += "&warnings=0&ts=" + Math.floor(new Date() / 1000);

        console.log(url);
      }
    } else {
      console.log("Donation would NOT execute, ABORT");
    }

  } catch (e) {
    if (e && e.reason) {
      console.log("Error reason:", e.reason);
      if (e.transaction) {
        console.log("Data:", e.transaction.data);
        console.log("To:", e.transaction.to);
      }
    }
    console.log(e);
  }

};


execute();
