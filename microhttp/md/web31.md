# Writing smartcontract DApps on the Energi3 blockchain

## Part X : Web3.js, Wallets, Accounts and Addresses

### 1. Address

An address is a public part, used to receive Energi's (NRG) and known by everyone. If someone wants to transfer money to you, you just have to tell them your Energi address. Its like your bank number.

``` js
> account = web3.eth.accounts.privateKeyToAccount("0xebd96b36ff9ce0c5aba9daaa037f767f66dfa4adf644291465436df482d56831", false);
{
  address: '0x03FB09251eC05ee9Ca36c98644070B89111D4b3F',
  privateKey: '0xebd96b36ff9ce0c5aba9daaa037f767f66dfa4adf644291465436df482d56831',
  signTransaction: [Function: signTransaction],
  sign: [Function: sign],
  encrypt: [Function: encrypt]
}
> account.address
'0x03FB09251eC05ee9Ca36c98644070B89111D4b3F'
>
```

An address is just a 20 byte string


### 2. Account

As you can see in the address example, an account is more than just a string. Its an object which contains besides the address also the privateKey.
The privateKey, as the name says ;) , is a 32 byte string, which gives you control over your account. With the privateKey you can send money to someone. This also means, that everybody who knows your privateKey can take all your money and you cant get it back. So never let anyone know it.
You should write down your pk or at least store it on an encrypted USB flash stick. If you loose it, your money is gone.

 * [Eth accounts](https://web3js.readthedocs.io/en/v1.2.11/web3-eth-accounts.html)

For fun you can check if there is some money inside the above account and then steal it ;)

Its generally a bad idea to use privateKeys to secure your funds. Better is to use [encrypted keystore files](https://www.google.com/search?q=encrypted+keystore+files).


``` js
> web3.eth.accounts.encrypt("0xebd96b36ff9ce0c5aba9daaa037f767f66dfa4adf644291465436df482d56831","nopassword")
{
  version: 3,
  id: '950c3b57-717d-4b3a-ba0a-b6e3d9cff9dc',
  address: '03fb09251ec05ee9ca36c98644070b89111d4b3f',
  crypto: {
    ciphertext: 'b3ccfcc77def98303aaa1d3fbee6bd78e5b7c7f72b04ce46a9a55bf53f35f762',
    cipherparams: { iv: 'f766302998a42b836215de76cc073b56' },
    cipher: 'aes-128-ctr',
    kdf: 'scrypt',
    kdfparams: {
      dklen: 32,
      salt: '5f0c0827c4d4b2b4c38a933a34a2b7512916ff77e0824535464bf9c9f5a56626',
      n: 8192,
      r: 8,
      p: 1
    },
    mac: '43943c1f0e671650091973b3309f921b3f80d527b72e65a9c808461b0b57f8a8'
  }
}
>
```

Save this json to a UTC-.. file. This is now a password encrypted keystore file which you can open in [MyEnergiWallet](https://docs.energi.software/downloads/myenergiwallet)


### 3. Wallet

A wallet is a collection of one or more accounts. Its usually a piece of software you use to manage your accounts.

* [Download My energi wallet](https://docs.energi.software/en/downloads/myenergiwallet)

