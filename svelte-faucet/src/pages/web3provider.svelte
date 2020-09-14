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
        A JSON-RPC Provider, also called a <b style="color:green;">Web3</b> Provider, is an Energi Blockchain node, which offers an api that one can use to
        interact with the blockchain, like sending or receiving money or calling functions in a SmartContract.<br>
        In this faucet we will be using the Web3.js library to interact with a local node or a remote node. A local node can hold <strong>accounts</strong>, if keystore is available, a remote node does, most probally, not.<br>
        <strong>If you use a remote node, you have to provide an account yourself, ether as a <b style="color:green;">keystore file</b> or as a private key</strong><br>
        There is a third possibility: <strong>MetaMask</strong>. This is a browser extension for Ethereum, but you can use it for Energi, too.
        Although MetaMask is not a recommended way for Energi, experienced users may have fun playing with it on the testnet.
      </CardContent>
      <CardFooter><Link href="https://www.google.com/search?q=web3+create+provider" external target="_blank">More Information</Link></CardFooter>
    </Card>
    <br>
    <Card outline>
      <CardHeader>Connect to an Energi Testnet RPC Provider
      </CardHeader>
      <CardContent>

      <Row noGap>
        <Col width="20">
          <List noHairlinesMd>
            <ListItem
              radio
              radioIcon="start"
              title="Local"
              value="http://localhost:49796"
              name="radio-url"
              onChange={ e => radioProviderChanged(e)}
            ></ListItem>
            <ListItem
              radio
              radioIcon="start"
              title="Remote"
              value="https://nodeapi.test3.energi.network"
              name="radio-url"
              onChange={ e => radioProviderChanged(e)}
            ></ListItem>
          </List>
        </Col>
        <Col width="60">
          <List noHairlinesMd>
            <ListInput
              outline
              label="Web3 Provider"
              floatingLabel
              type="text"
              clearButton
              value={web3URL}
              onInputClear={ () => {web3URL = "";}}
              onBlur={ e => urlChanged(e)}
              style="vertical-align: middle;margin-top:3%;"
            >
            <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
            </ListInput>
          </List>
        </Col>
        <Col width="20">
          <Button fill raised on:click={onConnectUrl} style="vertical-align: middle;margin-top:21%;bgcolor:var(--energi-color-green)">Connect</Button>
        </Col>
      </Row>
      <Row noGap>
        <Col width="5">
          <img class:pngColorGreen={connectedWeb3} src="./images/connectedRemote.png" title="Connection state" alt="Connected" height="30px" width="30px" style="left:12px;top:15px;position:relative;" />
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
              id="yellow-mono-small"
            >
            </ListInput>
          </List>
        </Col>
      </Row>
      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a>
      <p>{#if connectedWeb3}connected{:else}selected{/if} URL: {web3URL}</p>
      </CardFooter>
    </Card>


    <Card outline>
      <CardHeader>Choose an account to use</CardHeader>

      <CardContent>

        <List accordionList accordionOpposite>

          <ListItem accordionItem title="Keystore File" onClick={ () => {if (providerOpen) {providerOpen = false;}}} >
            <AccordionItem opened={keystoreOpen}>
            <AccordionContent>

                <Row noGap>
                  <Col width="30">
                    <input type="file" id="fileinput" class="hidden" bind:files={files} bind:this={browseInput} onChange={ e => console.log(e)}>
                    <Button onClick={() => browseInput.click()} fill raised style="align: middle;bgcolor:var(--energi-color-green)">Keystore File</Button>
                  </Col>
                  <Col width="50">
                    <ListInput
                      outline
                      label="Keystore File Password"
                      floatingLabel
                      type="text"
                      clearButton
                      value={password}
                      onInputClear={ () => {password = "";}}
                      onBlur={ e => passwordChanged(e)}
                    >
                    <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
                    </ListInput>

                  </Col>
                  <Col width="20">
                    <Button onClick={() => unlockKeystore()} fill raised style="vertical-align: middle;margin-top:5%;bgcolor:var(--energi-color-green)">Unlock</Button>
                  </Col>
                </Row>

            </AccordionContent>
          </AccordionItem>
          </ListItem>
          <ListItem accordionItem title="Private Key" onClick={ () => {if (providerOpen) {providerOpen = false;}}}>
            <AccordionItem opened={privateKeyOpen}>
            <AccordionContent>
              <ListInput
              outline
              label="Energi Privatekey"
              floatingLabel
              type="text"
              clearButton
              value={privateKey}
              onInputClear={ () => {privateKey = "";}}
              onBlur={ e => privateKeyChanged(e)}
            >
              <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
            </ListInput>
            </AccordionContent>
          </AccordionItem>
          </ListItem>
          <ListItem accordionItem title="Accounts active on the provider" onAccordionOpened={ () => {scrollTo("scrolltarget1");}}  onClick={ () => {if (providerOpen) {providerOpen = false;}}}>
            <AccordionItem opened={providerOpen}>
            <AccordionContent>

                <Row noGap>
                  <Col width="100">
                    <List noHairlinesMd>
                      {#if accounts[0]}
                        {#each accounts as { a }, i}
                          <ListItem
                            radio
                            radioIcon="start"
                            title={accounts[i]}
                            value={accounts[i]}
                            name="radio-account"
                            onChange={ e => radioAccountsChanged(e)}
                            id="yellow-mono-small"
                          ></ListItem>
                        {/each}
                      {/if}

                      <ListInput
                        outline
                        label="Energi Account"
                        floatingLabel
                        type="text"
                        clearButton
                        value={$selectedAccount}
                        onInputClear={ () => {$selectedAccount = "";}}
                        onBlur={ e => accountSelectedChanged(e)}
                      >
                        <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
                      </ListInput>

                  </List>
                </Col>
              </Row>

            </AccordionContent>
          </AccordionItem>
          </ListItem>
        </List>


      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a>
      <p>Account: {$selectedAccount}</p>
      </CardFooter>
    </Card>
    <p id="scrolltarget1">&nbsp;</p>



<!--
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
-->
</Page>

  <script>
    import { f7, AccordionItem, AccordionContent, NavLeft, NavTitle, Link, Row, Col, Card, CardHeader, CardContent, CardFooter, Page, Navbar, ListItem, List, ListInput, Button } from 'framework7-svelte';
    import { visitedPages, selectedAccount, getTime, scrollTo } from '../js/stores.js';
    //let Web3 = require('web3');
    //import { ethereum, chainName, isListening, nativeCurrency, selectedAccount, whenReady } from 'svelte-web3';
    import Web3 from 'web3';
    import { getContext } from 'svelte';

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

    let browseInput; // the input to open the file browse dialog
    let accordionItem2;
    let keystoreJson;
    let privateKey;
    let files;
    let password;
    $: {
        if (files && files[0]) {
            let keyfile = files[0];
            let reader = new FileReader();
            reader.onload = function(evt) {
              keystoreJson = evt.target.result;
            };
            reader.readAsText(keyfile);
        }
    }

    //const web3 = new Web3('http://localhost:8545');  //, null, options);

    const web3 = getContext('web3');
    console.log("Get web3 from context: ", web3);


    let web3URL='';
    let connectedWeb3 = false;
    let logText='';
    let accounts = [];
    let keystoreOpen = false;
    let privateKeyOpen = false;
    let providerOpen = false;


    //const enable = () => ethereum.setBrowserProvider()
    //const enable = () => ethereum.setProvider('http://localhost:49796');
    //$: balance = whenReady($selectedAccount, a => $web3.eth.getBalance(a));
    //$: balance = whenReady(true, c => { if (c) return 0; else return 99;});
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
    function radioProviderChanged(e) {
      console.log("radioProviderChanged(e): ", e.target.value );
      web3URL = e.target.value;
    }
    function radioAccountsChanged(e) {
      console.log("radioAccountsChanged(e): ", e.target.value );
      $selectedAccount = e.target.value;
    }
    function urlChanged(e) {
      console.log("urlChanged(e): " + e.target.value );
      web3URL = e.target.value;
    }
    function accountSelectedChanged(e) {
      console.log("accountSelected(e): ", e.target.value );
      $selectedAccount = e.target.value;
    }
    function passwordChanged(e) {
      console.log("passwordChanged(e): ", e.target.value );
      password = e.target.value;
    }
    function privateKeyChanged(e) {
      console.log("privateKeyChanged(e): ", e.target.value );
      privateKey = e.target.value;
    }

    function unlockKeystore() {
      // open the wallet in web3
      console.log("unlockKeystore: " + keystoreJson );
      console.log(accordionItem2);
      console.log(accordionItem2.opened);
      accordionItem2.opened = true;
      console.log(accordionItem2.opened);
      accordionItem2 = accordionItem2;
    }


    function onConnectUrl() {
      let web3Error = '';
      let callsFinished = 0;
      let complete = 2;

      logText = getTime() + ": Try connect to "  + web3URL;
      console.log("onConnect to ", web3URL);

      f7.dialog.preloader('Connecting to ' + web3URL);

      let p = web3.setProvider(new Web3.providers.HttpProvider(web3URL));
      console.log('SetProvider: ', p);


        web3.eth.getProtocolVersion()
         // .once('transactionHash', function(hash){ console.log(hash); })
         // .once('receipt', function(receipt){ console.log(receipt); })
          //.on('confirmation', function(confNumber, receipt){ console.log(confNumber, receipt); })
          //.on('error', function(error){ console.log(error); })
          .then((protocolVersion) => { callsFinished++;
                                       console.log(`Protocol Version: ${protocolVersion}`);
                                       logText = logText + '\n' + getTime() + ": Protocol Version " + protocolVersion;
                                       if (callsFinished > complete) {connectedWeb3 = true; f7.dialog.close();}

                                      }
          )
          .catch(function(error) {callsFinished++; if (callsFinished > complete) { f7.dialog.close(); }
                                  web3Error = error; console.log('catch:', error); logText = logText + `\n${getTime()}: ${error}`;}
          );

        web3.eth.getGasPrice()
          .then((gasPrice) => { callsFinished++;
                                console.log(`Gas Price: ${gasPrice}`);
                                logText = logText + `\n${getTime()}: Gas Price: ${gasPrice} wei, ${web3.utils.fromWei(gasPrice)} NRG`;
                                if (callsFinished > complete) {connectedWeb3 = true; f7.dialog.close();}

                              }
          )
          .catch((error) => {callsFinished++; if (callsFinished > complete) f7.dialog.close();
                            web3Error = error; console.log(error); logText = logText + `\n${getTime()}: ${error}`;
                            }
          );

        web3.eth.personal.getAccounts()
          .then((a) => { callsFinished++;
                        accounts = a;
                        console.log(`accounts: ${accounts}`);
                        logText = logText + `\n${getTime()}: Accounts: ${accounts}`;
                        if (callsFinished > complete) {connectedWeb3 = true; f7.dialog.close();}
                        if (accounts[0]) { providerOpen = true; keystoreOpen = false; }
                        //make the parent to scroll into view, smoothly!
                        scrollTo("scrolltarget1");
                       }

          )
          .catch((error) => { callsFinished++; if (callsFinished > 0) f7.dialog.close();
                             web3Error = error; console.log(error); logText = logText + `\n${getTime()}: ${error}`;
                             if (!accounts[0]) { providerOpen = false; keystoreOpen = true; }
                             //make the parent to scroll into view, smoothly!
                             scrollTo("scrolltarget1");
                            }
          );



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
      visitedPages.update(() => {
        visitedPages.web3 = true;
        visitedPages.smartcontract = false;
        visitedPages.info = false;
            return visitedPages;
        });
    }
    function onHide()  {
      visitedPages.update(() => {
        visitedPages.web3 = false;
            return visitedPages;
        });
    }
  </script>
