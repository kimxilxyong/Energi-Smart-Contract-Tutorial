# Writing smartcontract DApps on the Energi3 blockchain

## Part 1 : Setting up the development Environment

### 1. A web3 RPC provider node is needed locally
   
 * [Download the Energi3 Core node](https://docs.energi.software/en/downloads/core-node)
   
 * Install or Unpack into a new Directory which is called **EnergiHOME** from now on

 * Go into **EnergiHOME/bin**

 * Create a startup script or run directly: This will need several hours to download the testnet blockchain
``` bash
  ./energi3 --testnet --verbosity 3 --rpccorsdomain "https://remix.ethereum.org" --vmdebug --rpc --rpcport 39796 --rpcaddr "127.0.0.1" --rpcvhosts "localhost" --rpcapi admin,web3,eth,debug,personal,net,energi 
```
### 2. Create account(s) using the Energi Wallet
* [Download the Energi3 local App Wallet](https://docs.energi.software/en/downloads/myenergiwallet)
* Change Network to **Testnet!**
* Generate three Keystore files (3 Accounts)

![Wallet](https://i.ibb.co/yh2hRzq/Create-Wallet.png)

* Copy the keystore files to ```~/.energicore3/testnet/keystore/```

![Keystore](https://i.ibb.co/Hxp6qWw/Keystore-Files.png)


### 3. Learn how to inspect your accounts and unlock them 

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
___

---

***


## Typographic replacements

Enable typographer option to see result.

(c) (C) (r) (R) (tm) (TM) (p) (P) +-
© © ® ® ™ ™ § § ±

test.. test... test..... test?..... test!....

!!!!!! ???? ,,  -- ---

"Smartypants, double quotes" and 'single quotes'


## Emphasis

**This is bold text**

__This is bold text__

*This is italic text*

_This is italic text_

~~Strikethrough~~


## Blockquotes


> Blockquotes can also be nested...
>> ...by using additional greater-than signs right next to each other...
> > > ...or with spaces between arrows.


## Lists

Unordered

+ Create a list by starting a line with `+`, `-`, or `*`
+ Sub-lists are made by indenting 2 spaces:
  - Marker character change forces new list start:
    * Ac tristique libero volutpat at
    + Facilisis in pretium nisl aliquet
    - Nulla volutpat aliquam velit
+ Very easy!

Ordered

1. Lorem ipsum dolor sit amet
2. Consectetur adipiscing elit
3. Integer molestie lorem at massa


1. You can use sequential numbers...
1. ...or keep all the numbers as `1.`

Start numbering with offset:

57. foo
1. bar


## Code

Inline `code`

Indented code

    // Some comments
    line 1 of code
    line 2 of code
    line 3 of code


Block code "fences"

```
Sample text here...
```

Syntax highlighting

``` js
var foo = function (bar) {
  return bar++;
};

console.log(foo(5));
```

## Tables

| Option | Description                                                               |
| ------ | ------------------------------------------------------------------------- |
| data   | path to data files to supply the data that will be passed into templates. |
| engine | engine to be used for processing templates. Handlebars is the default.    |
| ext    | extension to be used for dest files.                                      |

Right aligned columns

| Option |                                                               Description |
| -----: | ------------------------------------------------------------------------: |
|   data | path to data files to supply the data that will be passed into templates. |
| engine |    engine to be used for processing templates. Handlebars is the default. |
|    ext |                                      extension to be used for dest files. |


## Links

[link text](http://dev.nodeca.com)

[link with title](http://nodeca.github.io/pica/demo/ "title text!")

Autoconverted link https://github.com/nodeca/pica (enable linkify to see)


## Images

![Wallet](https://i.ibb.co/yh2hRzq/Create-Wallet.png)
![Stormtroopocat](https://octodex.github.com/images/stormtroopocat.jpg "The Stormtroopocat")

Like links, Images also have a footnote style syntax

![Alt text][id]

With a reference later in the document defining the URL location:

[id]: https://octodex.github.com/images/dojocat.jpg  "The Dojocat"


## Plugins

The killer feature of `markdown-it` is very effective support of
[syntax plugins](https://www.npmjs.org/browse/keyword/markdown-it-plugin).


### [Emojies](https://github.com/markdown-it/markdown-it-emoji)

> Classic markup: :wink: :crush: :cry: :tear: :laughing: :yum:
>
> Shortcuts (emoticons): :-) :-( 8-) ;)

see [how to change output](https://github.com/markdown-it/markdown-it-emoji#change-output) with twemoji.


### [Subscript](https://github.com/markdown-it/markdown-it-sub) / [Superscript](https://github.com/markdown-it/markdown-it-sup)

- 19^th^
- H~2~O


### [\<ins>](https://github.com/markdown-it/markdown-it-ins)

++Inserted text++


### [\<mark>](https://github.com/markdown-it/markdown-it-mark)

==Marked text==


### [Footnotes](https://github.com/markdown-it/markdown-it-footnote)

Footnote 1 link[^first].

Footnote 2 link[^second].

Inline footnote^[Text of inline footnote] definition.

Duplicated footnote reference[^second].

[^first]: Footnote **can have markup**

    and multiple paragraphs.

[^second]: Footnote text.


### [Definition lists](https://github.com/markdown-it/markdown-it-deflist)

Term 1

:   Definition 1
with lazy continuation.

Term 2 with *inline markup*

:   Definition 2

        { some code, part of Definition 2 }

    Third paragraph of definition 2.

_Compact style:_

Term 1
  ~ Definition 1
 
Term 2
  ~ Definition 2a
  ~ Definition 2b
  
  ```
  
  ```


### [Abbreviations](https://github.com/markdown-it/markdown-it-abbr)

This is HTML abbreviation example.

It converts "HTML", but keep intact partial entries like "xxxHTMLyyy" and so on.

*[HTML]: Hyper Text Markup Language

### [Custom containers](https://github.com/markdown-it/markdown-it-container)

::: warning
*here be dragons*
:::