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

  const fromAddress = "0xdc67d898cc5fdd982308d1cb0c2301913f8f438c";
  let toAddress = "0xdc67d898cc5fdd982308d1cb0c2301913f8f438c";

  // how many confirmations do we want to wait for
  const waitForConfirmations = 3;

  // how much do we want to receive
  const spendeFromFaucetInETH = "100"; // NRG

  const ethers = require('ethers');
  const wei = ethers.utils.parseEther(spendeFromFaucetInETH);


  // If you don't specify a //url//, Ethers connects to the default
  // (i.e. ``http:/\/localhost:8545``)
  const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');

  // The provider also allows signing transactions to
  // send ether and pay to change state within the blockchain.
  // For this, we need the account signer...
  const signer = provider.getSigner(fromAddress);
  //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

  // !!! No signer, we want the raw transaction to pass to webwallet, which will do the signing

  // The Contract object
  const faucet = new ethers.Contract(contractAddr, abi, signer);

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
    let r = await faucet.callStatic.requestDonation(toAddress, wei, "Armer Mikkel", ethers.utils.toUtf8Bytes("🇸🇪")).catch((e) => { throw(e); });
    if (r) {
      console.log("requestDonation would execute, getting tx data ...");
      let unsignedTx = await faucet.populateTransaction.requestDonation(toAddress, wei, "Armer Mikkel", ethers.utils.toUtf8Bytes("🇸🇪")).catch((e) => { throw(e); });
      if (unsignedTx) {
        console.log(unsignedTx);
        console.log("Building webwallet URL for requesting", spendeFromFaucetInETH, "NRG to account", toAddress);

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
};

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
};

execute();
