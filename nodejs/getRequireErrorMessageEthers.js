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

  const fromAddress = "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c";

  const ethers = require('ethers');

  // If you don't specify a //url//, Ethers connects to the default
  // (i.e. ``http:/\/localhost:8545``)
  const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');

  // The provider also allows signing transactions to
  // send ether and pay to change state within the blockchain.
  // For this, we need the account signer...
  //const signer = provider.getSigner(fromAddress);
  //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

  // Use a read-only signer as we call only view/pure functions
  const signer = new ethers.VoidSigner(fromAddress, provider);
  // The Contract object
  const faucet = new ethers.Contract(contractAddr, abi, signer);
  // invalid index, should trigger: require(_i < Donors.length, "index _i out of bounds");
  const donorIndex = 99;
  try {
    let r = await faucet.getDonorAddress(donorIndex).catch((e) => { throw (e); });
    console.log("Donor Address result:", r);
  } catch (e) {
    if (e && e.reason) {
      console.log("contract require failed with:", e.reason);
    }
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
