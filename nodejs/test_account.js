

let Web3 = require('web3');
let web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:49796'));
let web3 = new Web3('http://localhost:49796');

web3.eth.personal.getAccounts().then(console.log);

web3.eth.sign("Test Message", "0x771dDB07222A1f9442C91CF04f64F3164771BB62").then(console.log).catch((e) => console.log(e.message));

web3.eth.personal.ecRecover("Test Message", "0x566289ee5c5d8c3608f81dc77f6afd7a1f557a0494e01a875582619d9ffa86f7402a64bda4118d9b9b8674c63669adc74d08a5020cec5e7e8a9f0f14ce43ebe41b").then(console.log).catch(console.log)
web3.eth.accounts.recover("Test Message", "0x566289ee5c5d8c3608f81dc77f6afd7a1f557a0494e01a875582619d9ffa86f7402a64bda4118d9b9b8674c63669adc74d08a5020cec5e7e8a9f0f14ce43ebe41b", false);



web3.eth.sendTransaction({ from: "0x771ddb07222a1f9442c91cf04f64f3164771bb62", to: "0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30", value: web3.utils.toWei("1", "ether") });
