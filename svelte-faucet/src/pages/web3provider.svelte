<Page name="web3" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
  <Navbar backLink="Back">
    <NavLeft>
       <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
      <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
    </NavLeft>
    <NavTitle>Connect to the Energi Testnet</NavTitle>
  </Navbar>

    <Card outline>
      <CardHeader>About RPC Providers</CardHeader>
      <CardContent>
        A JSON-RPC Provider, also called <b style="color:green;">a Web3</b> Provider, is a Energi Blockchain node, which offers an api that one can use to
        interact with the blockchain, like sending or receiving money or calling functions in a SmartContract.<br>
        In this faucet we will be using the Web3.js library to interact with a local (IPC) node or a remote (HTTP) node.<br>
        There is a third possibility: <strong>MetaMask</strong>. This is a browser extension for Ethereum, but you can use it for Energi, too.
        Although MetaMask is not a recommended way, experienced users may have fun playing with it on the testnet.
      </CardContent>
      <CardFooter><Link href="https://www.google.com/search?q=web3+create+provider" external target="_blank">More Information</Link></CardFooter>
    </Card>
    <br>
    <br>
    <Card outline>
      <CardHeader>Connection to the official Energi Testnet RPC Provider:&nbsp;https://nodeapi.test3.energi.network
      </CardHeader>
      <CardContent>

      <Row>
        <Col width="20">
          <List noHairlinesMd>
            <ListItem
              radio
              radioIcon="start"
              title="Local"
              value="http://localhost:49796"
              name="radio-url"
              onChange={ e => radioChanged(e)}
            ></ListItem>
            <ListItem
              radio
              radioIcon="start"
              title="Remote"
              value="https://nodeapi.test3.energi.network"
              name="radio-url"
              onChange={ e => radioChanged(e)}
            ></ListItem>
          </List>
        </Col>
        <Col width="65">
          <List noHairlinesMd>
            <ListInput
              outline
              label="Web3 Provider"
              floatingLabel
              type="text"
              clearButton
              value={urlRemote}
              onInputClear={ () => {urlRemote = "";}}
            >
            <!-- <i class="icon demo-list-icon" slot="media" /> -->
            </ListInput>
          </List>
        </Col>
        <Col width="15">
          <Button fill raised on:click={onConnectRemote} style="margin-top:16px;bgcolor:var(--energi-color-green)">Connect</Button>
        </Col>
      </Row>
      <Row noGap>
        <Col width="5">
          <img class:pngColorGreen={connectedRemote} src="./images/connectedRemote.png" alt="Connected Remote" height="30px" width="30px" style="left:12px;top:15px;position:relative;" />
        </Col>
        <Col width="95">
          <List noHairlinesMd>
            <ListInput
              outline
              floatingLabel
              label="Log"
              type="textarea"
              clearButton
              resizable
              value={logText}
              onInputClear={ () => {logText = "";} }
              id="logtextarea"
            >
            <!-- <i class="icon demo-list-icon" slot="media" /> -->
            </ListInput>
          </List>
        </Col>
      </Row>
      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a></CardFooter>
    </Card>





    <h1>Hello {name}!</h1>

    <p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>

    <p>Visit the <a href="https://web3js.readthedocs.io/en/">Web3.js documentation</a> to learn how to use Web3.js library.</p>

    <p>This svelte example app to the window provider (eg Metamask) : {$selectedAccount} </p>

    {#if $isListening}
    <p>
      Selected account: {$selectedAccount || 'not defined'}
    </p>
    {#if $chainName}
    <p>
      Balance on {$chainName}:
      {#await balance}<span>waiting...</span>{:then value}<span>{value}</span>{/await} {$nativeCurrency.symbol}
    </p>

    {/if}
    {/if}










</Page>

  <script>
    import { f7, NavLeft, NavTitle, Link, Row, Col, Card, CardHeader, CardContent, CardFooter, Page, Navbar, ListItem, List, ListInput, Button } from 'framework7-svelte';
    import { global, getTime } from '../js/stores.js';
    //let Web3 = require('web3');
    import { ethereum, chainName, isListening, nativeCurrency, selectedAccount, whenReady } from 'svelte-web3';
    import Web3 from 'web3';

const options = {
    defaultAccount: '0x0',
    defaultBlock: 'latest',
    defaultGas: 1,
    defaultGasPrice: 0,
    transactionBlockTimeout: 50,
    transactionConfirmationBlocks: 24,
    transactionPollingTimeout: 480
    //transactionSigner: new CustomTransactionSigner()
    };

    const web3 = new Web3('http://localhost:8545');  //, null, options);

    let name='FAUCET';
    let urlRemote='';
    let connectedRemote = false;
    let logText='';


    //const enable = () => ethereum.setBrowserProvider()
    //const enable = () => ethereum.setProvider('http://localhost:49796');
    //$: balance = whenReady($selectedAccount, a => $web3.eth.getBalance(a));
    $: balance = whenReady(true, c => { if (c) return 0; else return 99;});
    /*
    const unsubscribe = selectedAccount.subscribe(selectedAccount => {
      console.log('subscribe Account:', selectedAccount);
    });
    const unsubscribe1 = isListening.subscribe(isListening => {
      console.log('subscribe isListening:', isListening);
      if (isListening) logText = logText + '\nWeb3.js is listening';
    });
    const unsubscribe2 = web3s.subscribe(web3s => {
      console.log('subscribe web3:', web3s);
    });
    const unsubscribe3 = chainName.subscribe(chainName => {
      console.log('subscribe chainName:', chainName);
    });
    */
    function radioChanged(e) {
      console.log("radioChanged(e): " + e.target.value );
      urlRemote = e.target.value;
    }



    function onConnectRemote() {
      let web3Error = '';
      let callsFinished = 0;
      let complete = 1;

      connectedRemote = false;
      logText = "Try connect to "  + urlRemote;
      console.log("onConnectRemote to ", urlRemote);

      //f7.dialog.preloader('Connecting to ' + urlRemote);

      let p = web3.setProvider(new Web3.providers.HttpProvider(urlRemote));
      console.log('SetProvider: ', p);


        web3.eth.getProtocolVersion()
         // .once('transactionHash', function(hash){ console.log(hash); })
         // .once('receipt', function(receipt){ console.log(receipt); })
          //.on('confirmation', function(confNumber, receipt){ console.log(confNumber, receipt); })
          //.on('error', function(error){ console.log(error); })
          .then((protocolVersion) => { console.log(`Protocol Version: ${protocolVersion}`);
                                       logText = logText + '\n' + getTime() + ": Protocol Version " + protocolVersion;
                                       if (callsFinished > complete) {connectedRemote = true; f7.dialog.close();}
                                       callsFinished++;
                                      }
          )
          .catch(function(error) {if (callsFinished > complete) { f7.dialog.close(); }
                                  callsFinished++; web3Error = error; console.log('catch:', error); logText = logText + "\n" + getTime()  + error;
                                 }
          );

        web3.eth.getGasPrice(console.log)
          .then((gasPrice) => { console.log(`Gas Price: ${gasPrice}`);
                                logText = logText + `\n${getTime()}: Gas Price: ${gasPrice}`;
                                if (callsFinished > complete) {connectedRemote = true; f7.dialog.close();}
                                callsFinished++;
                              }
          )
          .catch((error) => {if (callsFinished > complete) f7.dialog.close();
                           callsFinished++; web3Error = error; console.log(error); logText = logText + `\n${getTime()}: ${error}`;
                          });

        web3.eth.personal.getAccounts(console.log)
          .then((accounts) => { console.log(`accounts: ${accounts}`);
                                logText = logText + `\n${getTime()}: Accounts: ${accounts}`;
                                if (callsFinished > complete) {connectedRemote = true; f7.dialog.close();}
                                callsFinished++;
                              })
          .catch((error) => {if (callsFinished > 0) f7.dialog.close();
                           callsFinished++; web3Error = error; console.log(error); logText = logText + `\n${getTime()}: ${error}`;
                          });



      //console.log('after setProvider Web3: ', web3);
      //console.log('setProvider: ', p);



      /*
      ethereum.setProvider(urlRemote)
      .then(result => {
        console.log("Connected: " + result);
        connectedRemote = true;
        logText = 'Connection to ' + urlRemote + ' OK!';
        f7.dialog.close();
      })
      .catch(error => {
        console.log("Connection failed: " + error);
        connectedRemote = false;
        logText = 'Connection to ' + urlRemote + ' FAILED!\n' + error;
        f7.dialog.close();
      });
      */
      // Connect to JSON RPC provider
      /*
      let provider = ethereum.setProvider(urlRemote);
      provider.then(result => {
        console.log("Connected: " + result);
        connectedRemote = true;
        logText = 'Connection to ' + urlRemote + ' OK!';
        f7.dialog.close();
      });
      provider.catch(error => {
        console.log("Connection failed: " + error);
        connectedRemote = false;
        logText = 'Connection to ' + urlRemote + ' FAILED!\n' + error;
        f7.dialog.close();
      });
      */
    }

    function onShow() {
        global.update(() => {
            global.web3 = true;
            global.smartcontract = false;
            global.info = false;
            return global;
        });
    }
    function onHide()  {
        global.update(() => {
            global.web3 = false;
            return global;
        });
    }
  </script>
