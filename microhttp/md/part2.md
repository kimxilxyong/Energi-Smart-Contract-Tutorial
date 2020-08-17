# Writing smartcontract DApps on the Energi3 blockchain

## Part 2 : Setting up the Go(lang) development Environment

### 1. Install the Go compiler

 * [Install the Go Language Compiler](https://golang.org/doc/install)

### 2. A good IDE is needed, use Visual Studio Code (VSCode)

 * [Install the official VS Code](https://code.visualstudio.com/docs/setup/linux)

### 3. Install the Go extensions for VS Code

Pic Extensions








### 4. Learn how to inspect your accounts and unlock them

* ```./energi3 --testnet --preload utils.js attach```

``` js
Welcome to the Energi Core JavaScript console!

instance: energi3/v3.0.6-stable/linux-amd64/go1.13.8
coinbase: 0x0000000000000000000000457068656d6572616c
at block: 145451 (Wed, 20 May 2020 15:46:44 CEST)
 datadir: /home/moe/.energicore3/testnet
 modules: admin:1.0 debug:1.0 energi:1.0 eth:1.0 masternode:1.0 miner:1.0 net:1.0 nrg:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

> personal.listWallets
[{
    accounts: [{
        address: "0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30",
        url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-12T09-34-29.733027979Z--7757a1f517d4680dba5d0ae9c984d3d394cc4a30"
    }],
    status: "Locked",
    url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-12T09-34-29.733027979Z--7757a1f517d4680dba5d0ae9c984d3d394cc4a30"
}, {
    accounts: [{
        address: "0x771ddb07222a1f9442c91cf04f64f3164771bb62",
        url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-14T13-47-18.288Z--771ddb07222a1f9442c91cf04f64f3164771bb62"
    }],
    status: "Locked",
    url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-14T13-47-18.288Z--771ddb07222a1f9442c91cf04f64f3164771bb62"
}, {
    accounts: [{
        address: "0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c",
        url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-14T13-50-27.459Z--17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c"
    }],
    status: "Locked",
    url: "keystore:///home/moe/.energicore3/testnet/keystore/UTC--2020-05-14T13-50-27.459Z--17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c"
}]
> checkAllBalances()
  eth.accounts[0]: 	0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30 	balance: 0.975950853 NRG
  eth.accounts[1]: 	0x771ddb07222a1f9442c91cf04f64f3164771bb62 	balance: 899.9460951174 NRG
  eth.accounts[2]: 	0x17fa844a3f96f2a8ae55e49c0e8ccf1bc628f90c 	balance: 97.949481646 NRG
  Total balance: 998.8715276163999 NRG
true
> personal.unlockAccount("0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30",null,0,false)
Unlock account 0x7757a1f517d4680dba5d0ae9c984d3d394cc4a30
Passphrase:
Passphrase entered: *****
true
>
```

* The accounts should now be: <```status: "Unlocked"```>, you can later use them to deploy smart contracts.

### 4. Get some money!

* Go to the faucet and beg for some tNRGs
*

### 5. Setup an easy to use solidity IDE
 * Remix is available in two versions: As a Web App or a local desktop App. I suggest using the local Version, but both do work.
 * WebApp https://remix.ethereum.org
 * LocalApp https://github.com/ethereum/remix-desktop/releases
 * Activate plugins needed for solidity

   ![Plugins](https://i.ibb.co/C7N2x76/Remix-IDEPlugins.png)
 * Optional install Remixd to load and store solidity files from the local filesystem and therefor have them managed by a git repository instead of only a simple gist: https://github.com/ethereum/remixd



