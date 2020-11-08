<!-- SPDX-License-Identifier: MIT -->
<!--  Copyright (c) 2020 Kim Il Yong -->

<Page name="web3" onPageBeforeIn="{onShow}" onPageBeforeOut="{onHide}">
  <Navbar backLink="Back">
    <NavLeft>
      <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
      <Link
        iconIos="f7:menu"
        iconAurora="f7:menu"
        iconMd="material:menu"
        panelOpen="left"
      />
    </NavLeft>
    <NavTitle>Connect to the Energi Testnet</NavTitle>
  </Navbar>

  <Card outline>
    <CardHeader>About RPC Providers</CardHeader>
    <CardContent>
      A JSON-RPC Provider, also called a
      <b style="color:green;">Web3</b>
      Provider, is an Energi Blockchain node, which offers an api that one can
      use to interact with the blockchain, like sending or receiving money or
      calling functions in a SmartContract.<br />
      In this faucet we will be using the Web3.js library to interact with a
      local node or a remote node. A local node can hold
      <strong>accounts</strong>, if keystore is available, a remote node does,
      most probally, not.<br />
      <strong>If you use a remote node, you have to provide an account yourself,
        ether as a
        <b style="color:green;">keystore file</b>
        or as a private key</strong><br />
      There is a third possibility:
      <strong>MetaMask</strong>. This is a browser extension for Ethereum, but
      you can use it for Energi, too. Although MetaMask is not a recommended way
      for Energi, experienced users may have fun playing with it on the testnet.
    </CardContent>
    <CardFooter>
      <Link
        href="https://www.google.com/search?q=web3+create+provider"
        external
        target="_blank"
      >
        More Information
      </Link>
    </CardFooter>
  </Card>
  <br />
  <Card outline>
    <CardHeader>Connect to an Energi Testnet RPC Provider</CardHeader>
    <CardContent>
      <Row noGap class="flex-row">
        <Col width="20" class="flex-col fixed">
          <List noHairlinesMd>
            <ListItem
              radio
              radioIcon="start"
              title="Local"
              value="http://localhost:49796"
              name="radio-url"
              id="small"
              onChange="{(e) => radioProviderChanged(e)}"
            />
            <ListItem
              radio
              radioIcon="start"
              title="Remote"
              value="https://nodeapi.test3.energi.network"
              name="radio-url"
              id="small"
              onChange="{(e) => radioProviderChanged(e)}"
            />
          </List>
        </Col>
      <Col width="60" class="flex-col dynamic top5">
          <List noHairlinesMd>
            <ListInput
              outline
              label="Web3 Provider"
              floatingLabel
              type="text"
              clearButton
              value="{web3URL}"
              onInputClear="{() => { web3URL = ''; }}"
              onBlur="{(e) => urlChanged(e)}"
              style="vertical-align:top;"
              spellcheck="false"
              id="mono-small"
            >
              <!-- <i class="icon demo-list-icon" slot="media" />   margin-top:24px -->
            </ListInput>
          </List>
        </Col>
        <Col width="20" class="flex-col fixed top2">
          <Button
            fill
            raised
            on:click="{ConnectToRPCProviderUrl}"
          >
            <i class="f7-icons button-icon">link</i>
            Connect
          </Button>
        </Col>
      </Row>
      <Row noGap class="flex-row">
        <Col width="100" class="flex-col dynamic">
        <List noHairlinesMd>

          <ListInput
              outline
              disabled
              floatingLabel
              label="Log"
              type="textarea"
              clearButton
              resizable
              value="{logConnect}"
              onInputClear="{() => {
                logConnect = '';
              }}"
              id="yellow-mono-small"
              spellcheck="false">
                    <!-- <i slot="media" class="nf-icons nf-hc-fw animate__animated animate__pulse animate__infinite" style="color:{connectedColor};">&#xf817;</i> -->
                    <i slot="media" class="nf-icons nf-hc-fw {connectedAnimate}" style="color:{connectedColor};">&#xf817;</i>

            </ListInput>
          </List>
    </Col>
      </Row>
    </CardContent>
    <CardFooter>
      <a
        href="https://www.google.com/search?q=web3+create+provider"
        external
        target="_blank"
      >More Information</a>
      <p>
        {#if connectedWeb3}connected{:else}selected{/if}
        URL:
        {web3URL}
      </p>
    </CardFooter>
  </Card>

  <Card outline>

    <CardHeader>Choose an account to use</CardHeader>

    <CardContent>
      <List accordionList accordionOpposite>

        <ListItem accordionItem title="From Keystore File" id="AccordiumItem1">

          <AccordionItem>
            <AccordionContent>
            <Block>

                <Row noGap class="flex-row">
                  <Col class="flex-col fixed top1">       <!-- align-items:stretch;resize:none;flex: 0 0 0%;"> -->
                    <input
                      type="file"
                      id="fileinput"
                      class="hidden"
                      bind:files
                      bind:this="{browseInput}"
                      onChange="{(e) => console.log(e)}"
                    />
                    <Button
                      onClick="{() => browseInput.click()}"
                      fill
                      raised
                      style=""
                    >
                      <i class="f7-icons button-icon">doc_text_search</i>
                      Choose File
                    </Button>
                  </Col>
                  <Col width="70" class="flex-col dynamic top4">     <!-- "align-items:stretch;resize:horizontal;flex: 1 1 100%;"> -->
                    <ListInputPassword
                      label="Keystore Password"
                      id="mono-small"
                      inputId="PasswordField"
                      style="position:relative;top:-0.5rem;"
                      clearButton="{true}"
                      bind:password
                    />
                  </Col>
                  <Col width="15" class="flex-col fixed top1">       <!-- "align-items:stretch;resize:none;flex: 0 0 0%;"> -->
                    <Button
                      onClick="{() => unlockKeystore()}"
                      fill
                      raised
                    >
                      <i class="f7-icons button-icon">lock_open</i>
                      Unlock
                    </Button>
                  </Col>
                </Row>

              </Block>
            </AccordionContent>
          </AccordionItem>
        </ListItem>

        <ListItem accordionItem title="From Private Key" id="AccordiumItem2">
          <AccordionItem>
            <AccordionContent>
              <Block>
                <Row noGap class="flex-row">
                  <Col width="90" class="flex-col dynamic top62">
                    <ListInput
                      outline
                      label="Energi Privatekey"
                      floatingLabel
                      type="text"
                      placeholder="Enter private key"
                      info="Hexnumber starting with 0x"
                      errorMessage="Only hexnumbers starting with 0x please!"
                      required
                      validate
                      pattern="^0x[a-fA-F0-9]*"
                      clearButton
                      value="{privateKey}"
                      onInputClear="{() => {
                        privateKey = '';
                      }}"
                      onBlur="{(e) => onPrivateKeyChanged(e)}"
                      onValidate="{(isValid) => checkPrivateKeyValid(isValid)}"
                      inputId="inputIDPrivateKey"
                      spellcheck="false"
                    >
                    </ListInput>
                  </Col>
                  <Col class="flex-col fixed top9">       <!-- align-items:stretch;resize:none;flex: 0 0 0%;"> -->
                    <Button fill raised on:click="{ScanPrivateKey}">
                      <i class="f7-icons button-icon">barcode_viewfinder</i>
                    Scan
                    </Button>
                  </Col>
                </Row>
            </Block>
            </AccordionContent>
          </AccordionItem>
        </ListItem>

        <ListItem accordionItem title="Accounts active on the provider" id="AccordiumItem3" onAccordionOpened="{() => { scrollTo('scrolltarget1'); }}">
          <AccordionItem>
            <AccordionContent>
              <Row noGap>
                <Col width="100">
                  <List noHairlinesMd>
                    {#if accounts[0]}
                      {#each accounts as { a }, i}
                        {#await getAccountStatus(accounts[i])}
                          <ListItem
                            radio
                            radioIcon="start"
                            title="{accounts[i]} waiting..."
                            value="{accounts[i]}"
                            name="radio-account"
                            onChange="{(e) => radioAccountsChanged(e,i)}"
                            id="yellow-mono-small"
                          />
                        {:then ad}
                          <ListItem
                            radio
                            radioIcon="start"
                            title="{ad.address}, {ad.status}, Balance: {Web3.utils.fromWei(ad.balance, 'ether')} NRG"
                            value="{accounts[i]}"
                            name="radio-account"
                            onChange="{(e) => radioAccountsChanged(e,i)}"
                            id="yellow-mono-small"
                          />
                        {/await}
                      {/each}
                    {:else}
                      <ListItem
                        radio
                        radioIcon="start"
                        title="No account found"
                        value=""
                        name="radio-account"
                        onChange="{(e) => radioAccountsChanged(e)}"
                        id="grey-mono-small"
                      />
                    {/if}
                  </List>
                </Col>
              </Row>
            </AccordionContent>
          </AccordionItem>
        </ListItem>
      </List>
    </CardContent>
    <CardFooter>
      <a href="https://www.google.com/search?q=web3+create+provider" external target="_blank">More Information</a>
      <p>Address: {selectedAddress}</p>
    </CardFooter>
  </Card>

  <i id="scrolltarget1">&nbsp;</i>

  <Card outline>

    <CardHeader>Verify account address</CardHeader>

    <CardContent>
      <Block>
        <Row noGap class="flex-row">
          <Col width="90" class="flex-col dynamic">
            <List>
            <ListInput
              outline
              label="Energi Address"
              floatingLabel
              type="text"
              clearButton
              info="Hexnumber starting with 0x"
              errorMessage="Only hexnumbers starting with 0x please!"
              validate
              pattern="^0x[a-fA-F0-9]*"
              id="mono-small"
              value="{selectedAddress}"
              onInputClear="{() => {
                selectedAddress = '';
              }}"
              onBlur="{(e) => addressSelectedChanged(e)}"
              inputId="inputIDSelectedAddress"
              spellcheck="false"
            >

            </ListInput>
            </List>
          </Col>
          <Col class="flex-col fixed top3">
          <Button
              fill
              raised
              on:click="{TestAccount}"
            >
              <i class="f7-icons size-20" style="margin-right: 5px;margin-top:5px">search</i>
              Verify Address
            </Button>
          </Col>
        </Row>
        <Row noGap class="flex-row">
          <Col class="flex-col fixed">
            <Block style="top:1.6em;">
              <i class="f7-icons" style="align:center; font-size:24px; color: {creditcardColor};">creditcard</i>
            </Block>
          </Col>
          <Col width="100" class="flex-col dynamic">
            <List noHairlinesMd>
              <ListInput
                outline
                disabled
                floatingLabel
                label="Log"
                type="textarea"
                clearButton
                resizable
                value="{logAccount}"
                onInputClear="{() => {
                  logAccount = '';
                }}"
                id="yellow-mono-small"
                spellcheck="false"
              />
            </List>
          </Col>
        </Row>
      </Block>

      <Block>

        <List>
          <ListItem></ListItem>
          <ListItem header="Balance" title={balanceTitle}>
            <i slot="media" class="f7-icons size-20" style="color:{plusminusColor};">plus_slash_minus</i>
          </ListItem>
          <ListItem header="Sign a message" title={signedTitle}>
            <i slot="media" class="f7-icons size-20" style="color:{pawColor};">paw</i>
          </ListItem>
          {#if isManagedOnNode}
            <ListItem header="Choosen Account Address" title="{$ethAccount.address}" footer="{ choosenAddressIsValid ? 'Is validated' : 'Is not validated'}" after="">
              <i slot="media" class="f7-icons size-20" style="color:{thumbsupColor};">hand_thumbsup</i>
              <i on:click={CopyToClipboard($ethAccount.address, false)} slot="after-end" class="f7-icons size-20" style="margin-right:5px;color:{rocketColor}">rocket</i>
            </ListItem>
            <ListItem header="Private Key" title="Privatekey is managed by node {web3URL}" footer="Use personal.UnlockAccount on the nodes console" after="" link="">
              <i slot="media" class="f7-icons size-20" style="color:{privatekeyColor};">lock_slash</i>
            <!--  <i slot="after-start" class="f7-icons size-20">calendar_badge_plus</i> -->
              <i on:click={CopyToClipboard($ethAccount.privateKey, true)}  slot="after-start" class="f7-icons size-20" style="margin-right:5px">lock_open</i>
            </ListItem>
          {:else}
            <ListItem header="Choosen Account Address" title="{$ethAccount.address}" footer="{ choosenAddressIsValid ? 'Is validated' : 'Is not validated'}">
              <i slot="media" class="f7-icons size-20" style="color:{thumbsupColor};">hand_thumbsup</i>
            </ListItem>
            <ListItem header="Private Key" title="{$ethAccount.privateKey}" footer="IMPORTANT, don't loose it !!! Copy it to the clipboard and store it on an encrypted usb key" after="Copy">
              <i slot="media" class="f7-icons size-20" style="color:{privatekeyColor};">umbrella envelope_open</i>
            <!--  <i slot="after-start" class="f7-icons size-20">calendar_badge_plus</i> -->
              <i on:click={CopyToClipboard($ethAccount.privateKey, true)}  slot="after-start" class="f7-icons size-20" style="margin-right:5px">tray_arrow_down</i>
            </ListItem>
          {/if}
        </List>
      </Block>
    </CardContent>
    <CardFooter>
      <a
        href="https://www.google.com/search?q=ethereum+address+account+privatekey"
        external
        target="_blank"
      >More Information</a>
      <p>Balance: {selectedAddressBalance}</p>
    <div style="display:flex;flex-direction: row;justify-content: flex-end;align-items:center;">
        <Button disabled={!nextButtonEnabled} fill raised href="/smartcontract/" view=".view-main" panelClose class={nextButtonEnabled ? "" : "disablebar"} style="display:flex;align:center;align-items:center;padding-left:15%;bgcolor:{nextButtonEnabled ? 'var(--energi-color-green)' : 'var(--energi-color-grey)'}">Next
          <i class="f7-icons size-20" style="margin-left:0.2rem;margin-top:0.3rem">arrowtriangle_right</i>
        </Button>
      </div>
    </CardFooter>
  </Card>
  <i id="scrolltarget2">&nbsp;</i>

</Page>

<script>
  import {
    f7,
    AccordionItem,
    AccordionContent,
    NavLeft,
    NavTitle,
    Link,
    Row,
    Col,
    Card,
    CardHeader,
    CardContent,
    CardFooter,
    Page,
    Navbar,
    ListItem,
    List,
    Block,
    BlockHeader,
    ListInput,
    Button,
    Preloader,
  } from "framework7-svelte";
  import Web3 from "web3";
  import { tick, getContext } from "svelte";
  import {
    visitedPages,
    isUnlocked,
    ethAccount,
    getTime,
    scrollTo,
  } from "../js/stores.js";
  import { listWallets } from "../js/web3utils.js";
  import ListInputPassword, { setFocus } from "../components/input-password.svelte";


  const options = {
    defaultAccount: "0x0",
    defaultBlock: "latest",
    defaultGas: 1,
    defaultGasPrice: 0,
    transactionBlockTimeout: 50,
    transactionConfirmationBlocks: 24,
    transactionPollingTimeout: 480,
    //transactionSigner: new CustomTransactionSigner()
  };

  let browseInput; // the input to open the file browse dialog
  let keystoreJson;
  let files;
  let password;
  let web3Error = "";
  let accountDetails = [];
  let web3URL = "";
  let connectedWeb3 = false;  // true if connection to remote or local node is established
  let isLocalNode = false;

  let logConnect = "";
  let logAccount = "";
  let accounts = [];

  let selectedAddress;
  let privateKey;
  let gotBalance;
  let isSigned;
  let isVerified;

  let connectedColor = "var(--energi-color-red)";
  let connectedAnimate = "animate__animated animate__pulse animate__infinite";
  let creditcardColor = "var(--energi-color-grey)";
  let pawColor = "var(--energi-color-grey)";
  let plusminusColor = "var(--energi-color-grey)";
  let thumbsupColor = "var(--energi-color-grey)";
  let privatekeyColor = "var(--energi-color-grey)";
  let rocketColor = "var(--energi-color-grey)";
  let selectedAddressBalance = 0;
  let balanceTitle = "N/A";
  let signedTitle = "N/A";
  let choosenAddressIsValid = false;
  let nextButtonEnabled = false;
  let isManagedOnNode = false;


  setInterval(() => {

    // check connection
    if (connectedWeb3) {
      connectedAnimate = "";
    }
  }, 1000);

  $: {
    if (files && files[0]) {
      let keyfile = files[0];
      let reader = new FileReader();
      reader.onload = function (evt) {
        keystoreJson = evt.target.result;
        document.getElementById("PasswordField").focus();
      };
      reader.onerror = function (e) {
        console.log(e);
        createErrorPopup("Error " + e.target.error.name, e.target.error.message);
        keystoreJson = "";
      };
      reader.readAsText(keyfile);
    }
  }

  const whenReady = (...args) => {
    const fn = args.pop();
    for (const arg of args) {
      if (!arg) return Promise.reject(new Error("not valid"));
    }
    /// check fn is a fn
    return fn(...args);
  };

  function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
  }

  //const web3 = new Web3('http://localhost:8545');  //, null, options);

  const web3 = getContext("web3");
  console.log("Get web3 from context: ", web3);


  //const enable = () => ethereum.setBrowserProvider()
  //const enable = () => ethereum.setProvider('http://localhost:49796');
  //$: balance = whenReady(accounts[i], a => $web3.eth.getBalance(a));
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
    console.log("radioProviderChanged(e): ", e.target.value);
    web3URL = e.target.value;
  }
  function radioAccountsChanged(e, i) {
    console.log("radioAccountsChanged(e): ", e.target.value);
    selectedAddress = e.target.value;
    isManagedOnNode = true;

    console.log("accountDetails[i]: ", accountDetails[i]);

    $ethAccount.address = selectedAddress;
    $ethAccount.isManagedOnNode = isManagedOnNode;
    $ethAccount.isUnlocked = false;
    if (accountDetails[i].status) {
      if (accountDetails[i].status !== "Locked") {
        $ethAccount.isUnlocked = true;
      }
    }
  }
  function urlChanged(e) {
    console.log("urlChanged(e): " + e.target.value);
    web3URL = e.target.value;
  }
  function addressSelectedChanged(e) {
    console.log("accountSelected(e): ", e.target.value);
    selectedAddress = e.target.value;
  }
  function passwordChanged(e) {
    console.log("passwordChanged(e): ", e.target.value);
    password = e.target.value;
  }

  function checkPrivateKeyValid(isValid) {
    console.log("Valid?", isValid);
    return false;
  }
  function onPrivateKeyChanged(c) {
    console.log("Key: ", c);
    privateKey = c.target.value;
  }

  function toggleAccordionItem(index, open) {
    let ai = document.getElementById("AccordiumItem" + index);
    if (ai) {
      //console.log(ai.classList);

      let link = ai.querySelector("a.item-link");
      if (link) {
        if (ai.classList.contains("accordion-item-opened")) {
          if (!open) {
            //ai.classList.remove("accordion-item-opened");
            link.click();
          }
        } else {
          if (open) {
            //ai.classList.add("accordion-item-opened");
            link.click();
          }
        }
      }
    }
  }

  async function ConnectToRPCProviderUrl() {
    let callsFinished = 0;
    let complete = 2;

    // Reset fetched account infos
    accounts = [];
    accountDetails = [];
    connectedWeb3 = false;
    connectedColor = "var(--energi-color-red)";
    connectedAnimate = "animate__animated animate__pulse animate__infinite";

    // Reset connection details
    isLocalNode = false;
    if (web3URL.includes("localhost") || web3URL.includes("127.0.0.1")) {
      isLocalNode = true;
    } else {
      isManagedOnNode = false;
    }

    // Reset verification details
    gotBalance = false;
    isSigned = false;
    isVerified = false;

    logConnect = getTime() + ": Try connect to " + web3URL;

    f7.dialog.preloader("Connecting to " + web3URL);

    let p = web3.setProvider(new Web3.providers.HttpProvider(web3URL));
    console.log("SetProvider: ", p);

    try {
      web3.eth.getProtocolVersion()
        // .once('transactionHash', function(hash){ console.log(hash); })
        // .once('receipt', function(receipt){ console.log(receipt); })
        //.on('confirmation', function(confNumber, receipt){ console.log(confNumber, receipt); })
        //.on('error', function(error){ console.log(error); })
        .then((protocolVersion) => {
          callsFinished++;
          logConnect = logConnect + "\n" + getTime() + ": Protocol Version " + protocolVersion;
          if (callsFinished > complete) {
            connectedWeb3 = true;
            connectedColor = "var(--energi-color-green)";
            f7.dialog.close();
          }
        })
        .catch(function (error) {
          callsFinished++;
          if (callsFinished > complete) {
            f7.dialog.close();
          }
          web3Error = error;
          console.log("catch:", error);
          logConnect = logConnect + `\n${getTime()}: ${error}`;
        });
    }
    catch (error) {
      callsFinished++;
      if (callsFinished > complete) {
        f7.dialog.close();
      }
      web3Error = error;
      console.log("catch:", error);
      logConnect = logConnect + `\n${getTime()}: ${error}`;
    }

    try {
    web3.eth.getGasPrice().then((gasPrice) => {
        callsFinished++;
        logConnect = logConnect + `\n${getTime()}: Gas Price: ${gasPrice} wei, ${web3.utils.fromWei( gasPrice )} NRG`;
        if (callsFinished > complete) {
          connectedWeb3 = true;
          connectedColor = "var(--energi-color-green)";
          f7.dialog.close();
        }
      })
      .catch((error) => {
        callsFinished++;
        if (callsFinished > complete) f7.dialog.close();
        web3Error = error;
        console.log(error);
        logConnect = logConnect + `\n${getTime()}: ${error}`;
      });
    }
    catch (error) {
        callsFinished++;
        if (callsFinished > complete) {
          f7.dialog.close();
        }
        web3Error = error;
        console.log("catch:", error);
        logConnect = logConnect + `\n${getTime()}: ${error}`;
    }

    try {
    web3.eth.personal.getAccounts()
      .then((a) => {
        callsFinished++;
        accounts = a;
        logConnect = logConnect + `\n${getTime()}: Accounts: ${accounts}`;
        if (callsFinished > complete) {
          connectedWeb3 = true;
          connectedColor = "var(--energi-color-green)";
          f7.dialog.close();
        }
        if (accounts[0]) {
          toggleAccordionItem(3, true);
        }
        //make the parent to scroll into view, smoothly!
        scrollTo("scrolltarget1");

        // fetch account status
        listWallets(web3)
          .then( async (r) => {
            console.log("listWallets:", r.result.length);
            for (let i = 0; i < r.result.length; i++) {
              accountDetails.push({
                status: r.result[i].status,
                address: r.result[i].accounts[0].address,
                balance: "-1",
              });
            };
            connectedColor = "var(--energi-color-green)";

            await tick();
            scrollTo("scrolltarget1");
          });
          /*
          .catch((e) => {
            logConnect += `\n${getTime()}: listWallets: ${e}`;
            logConnect += `\n${getTime()}: NOTE: Make sure personal is loaded in your node`;
            logConnect += `\n${getTime()}: {@html "Example: ./energi --rpcapi admin,web3,eth,debug,<strong>personal</strong>,net,energi"}`;

            accountDetails.push({
              status: "-1",
              address: e,
              balance: "-1",
            });
            if (web3URL.includes("localhost")) {
              console.log("web3URL.contains(localhost)");
              console.log(web3URL);
              connectedColor = "var(--energi-color-red)";
            } else {
              console.log("NOT web3URL.contains(localhost)");
              connectedColor = "var(--energi-color-green)";
            }
          }); */
      })
      .catch((error) => {
        callsFinished++;
        if (callsFinished > complete) {
          connectedWeb3 = true;
          f7.dialog.close();
        }
        web3Error = error;
        console.log(error);
        logConnect += `\n${getTime()}: ${error}`;
        accountDetails.push({
          status: "-1",
          address: error,
          balance: "-1",
        });
        if (isLocalNode) {
          logConnect += `\n${getTime()}: NOTE: Make sure personal is loaded in your node`;
          logConnect += `\n${getTime()}: "Example: ./energi --rpcapi admin,web3,eth,debug,personal,net,energi"`;
          connectedColor = "var(--energi-color-red)";
        } else {
          connectedColor = "var(--energi-color-green)";
        }
        if (!accounts[0]) {
          toggleAccordionItem(1, true);
        }
        //make the parent to scroll into view, smoothly!
        scrollTo("scrolltarget1");
      });
    }
    catch (error) {
      callsFinished++;
      if (callsFinished > complete) {
        f7.dialog.close();
      }
      web3Error = error;
      console.log("catch:", error);
      logConnect = logConnect + `\n${getTime()}: ${error}`;
    }
  }


  function ScanPrivateKey() {
    ResetValidEthAccount();
    f7.dialog.preloader(
      "web3.eth.accounts.privateKeyToAccount (" + privateKey
    ) + ")";

    console.log("ConvertPrivateKey: ", privateKey);

    try {
      const account = web3.eth.accounts.privateKeyToAccount(privateKey, true);
      web3.eth.accounts.wallet.add(account);
      web3.eth.defaultAccount = account.address;
      selectedAddress = web3.eth.defaultAccount;
      console.log("Address: ${account.address}");
      console.log("Account: ", account);
      logAccount = logAccount + "\n" + getTime() + ": Address " + account.address;

      // store account object for later
      $ethAccount = account;

      f7.dialog.close();
      creditcardColor = "var(--energi-color-yellow)";
    } catch (e) {
      f7.dialog.close();
      web3Error = e;
      console.log("catch:", e);
      logAccount = logAccount + `\n${getTime()}: ${web3Error}`;
      createErrorPopup("Error converting Private Key", web3Error);
      creditcardColor = "var(--energi-color-red)";
    }
  }

  async function unlockKeystore() {
    // open the wallet in web3
    f7.dialog.preloader("Unlocking Keystore");
    console.log("unlockKeystore: " + keystoreJson);

    setTimeout(() => {
      try {
        const decryptedAccount = web3.eth.accounts.decrypt(keystoreJson, password);

        if (decryptedAccount) {
          selectedAddress = decryptedAccount.address;
          privateKey = decryptedAccount.privateKey;

          web3.eth.accounts.wallet.add(decryptedAccount);
          web3.eth.defaultAccount = decryptedAccount.address;

          // store account object for later
          $ethAccount = decryptedAccount;
          $isUnlocked = true;

          logAccount += "\n" + getTime() + ": decryptedAccount **********************";
          logAccount += "\n" + JSON.stringify(decryptedAccount, undefined, 4);
          logAccount += "\n" + getTime() + ": ***************************************";
          logAccount = logAccount.trim();

          console.log(decryptedAccount);

          //make the parent to scroll into view, smoothly!
          scrollTo("scrolltarget2");
          f7.dialog.close();
        }
        } catch(e) {
            f7.dialog.close();
            createErrorPopup("Unlocking failed", e);
            creditcardColor = "var(--energi-color-red)";
        }
      }, 300);

    await tick();
  }

 function ResetValidEthAccount() {
    creditcardColor = "var(--energi-color-grey)";
    pawColor = "var(--energi-color-grey)";
    plusminusColor = "var(--energi-color-grey)";
    rocketColor = "var(--energi-color-grey)";
    thumbsupColor = $ethAccount.address ? "var(--energi-color-red)" : "var(--energi-color-grey)";
    privatekeyColor = $ethAccount.privateKey ? "var(--energi-color-red)" : "var(--energi-color-grey)";
    balanceTitle = "N/A";
    signedTitle = "N/A";
    selectedAddressBalance = 0;
    choosenAddressIsValid = false;
    nextButtonEnabled = false;
 }

function setValidatedColor() {
  creditcardColor = "var(--energi-color-red)";
  thumbsupColor = "var(--energi-color-red)";
  rocketColor = "var(--energi-color-red)";
  privatekeyColor = "var(--energi-color-red)";

  if (gotBalance) {
    plusminusColor = "var(--energi-color-green)";
  } else {
    plusminusColor = "var(--energi-color-red)";
  }

  if (isSigned) {
    pawColor = "var(--energi-color-green)";
  } else {
    pawColor = "var(--energi-color-red)";
  }
  if (isVerified) {
    creditcardColor = "var(--energi-color-green)";
    thumbsupColor = "var(--energi-color-green)";
    rocketColor = "var(--energi-color-green)";
    if (isManagedOnNode) {
      privatekeyColor = "var(--energi-color-green)";
    } else {
      privatekeyColor = "var(--energi-color-yellow)";
    }
  }
  console.log("setValidatedColor", gotBalance, isSigned, isVerified);
}

 function TestAccount() {
    f7.dialog.preloader("Testing Account");

    if (selectedAddress !== $ethAccount.address) {
      $ethAccount.address = selectedAddress;
    }

    ResetValidEthAccount();

    getBalance($ethAccount.address);
  }

  function getBalance(address) {
    if (address) {
      let checksumAddress;
      try {
        checksumAddress = Web3.utils.toChecksumAddress(address);
      } catch (error) {
        f7.dialog.close();
        setValidatedColor();
        createErrorPopup("Error occured", error);
        return null;
      }

      if (checksumAddress !== address) {
        f7.dialog.close();
        setValidatedColor();
        logAccount = logAccount + "\n" + getTime() + ": Invalid address checksum";
        logAccount = logAccount.trim();
        createErrorPopup("Error occured", "Invalid address checksum");
      } else {
              try {
                web3.eth.getBalance(address).then((b) => {
                  selectedAddressBalance = web3.utils.fromWei( b );
                  balanceTitle = selectedAddressBalance  + " NRG";
                  //plusminusColor = "var(--energi-color-green)";

                  gotBalance = true;
                  setValidatedColor();

                  logAccount = logAccount + "\n" + getTime() + ": Balance for '" + address + "' is '" + selectedAddressBalance + "'";
                  logAccount = logAccount.trim();
                  console.log("Start sign");
                  const message = "Energi Faucet Test message to be signed";
                  const signatureObject = SignMessage(message);
                  console.log("After sign", signatureObject);
                  if (signatureObject) {
                    console.log("Start recover");
                    isSigned = true;
                    const signator = web3.eth.accounts.recover(signatureObject);
                    console.log("After recover");
                    if (signator === $ethAccount.address) {
                      signedTitle = 'Signator "' + signator + '" is VALID';
                      isVerified = true;
                      choosenAddressIsValid = true;
                    } else {
                      signedTitle = 'Signator "' + signator + '" is INVALID';
                      isVerified = false;
                    }
                  } else {
                    isSigned = false;
                    createErrorPopup("Error occured", "web3.eth.accounts.sign(message, privateKey) failed");
                  }
                  setValidatedColor();
                  f7.dialog.close();
                }).catch((e) => {
                  f7.dialog.close();
                  setValidatedColor();
                  createErrorPopup("Error occured", e);
                  console.log(e);
                });
              } catch (error) {
                f7.dialog.close();
                setValidatedColor();
                createErrorPopup("Error occured", error);
              }
            }
    } else {
        setValidatedColor();
        logAccount = logAccount + "\n" + getTime() + ": Address is not set";
        logAccount = logAccount.trim();
        f7.dialog.close();
    }
  }

  function SignMessage(msg) {
    try {
      const signature = web3.eth.accounts.sign(msg, $ethAccount.privateKey);
      logAccount += "\n" + getTime() + ": Signature: ****************************";
      logAccount += "\n" + JSON.stringify(signature, undefined, 4);
      logAccount += "\n" + getTime() + ": ***************************************";
      logAccount = logAccount.trim();
      return signature;
    } catch(e) {
      creditcardColor = "var(--energi-color-red)";
      logAccount = logAccount + "\n" + getTime() + ": Signing failed: " + e;
      logAccount = logAccount.trim();
      signedTitle = "Signing failed: " + e;
      createErrorPopup("Signing failed", e);
    }
    return null;
  }


  function getAccountStatus(a) {
    return new Promise((resolve, reject) => {
      // check every second if the accountslist has been filled already
      let iv = setInterval(getBalanceForAll, 1000);

      function getBalanceForAll() {
        if (accountDetails.length > 0) {
          for (let i = 0; i < accountDetails.length; i++) {
            if (Web3.utils.toChecksumAddress(accountDetails[i].address) === a) {
              try {
                web3.eth.getBalance(a).then((b) => {
                  clearInterval(iv);
                  resolve({
                    status: accountDetails[i].status,
                    address: a,
                    balance: b,
                  });
              });
              } catch (error) {
                reject(error);
              }
            }
          }
        } else {
          console.log("accountDetails.length = 0, wait for next intervall !!!");
        }
      }
    });
  }


  function sleepFor(sleepDuration) {
    let now = new Date().getTime();
    while (new Date().getTime() < now + sleepDuration) { tick(); }
  }

  function onShow() {
    visitedPages.update(() => {
      visitedPages.web3 = true;
      visitedPages.smartcontract = false;
      visitedPages.info = false;
      return visitedPages;
    });
  }
  function onHide() {
    visitedPages.update(() => {
      visitedPages.web3 = false;
      return visitedPages;
    });
  }

  function CopyToClipboard(text, isPrivateKey) {
    if (choosenAddressIsValid || !isPrivateKey) {
      navigator.clipboard.writeText(text).then(
        function() {
          /* clipboard successfully set */
          console.log("CopyToClipboard PK: ", isPrivateKey);
          if (isPrivateKey) {
            privatekeyColor = "var(--energi-color-green)";
          } else {
            rocketColor = "var(--energi-color-green)";
          }
          choosenAddressIsValid = true;
          nextButtonEnabled = true;

        }, function() {
          /* clipboard write failed */
          createErrorPopup("Copy to Clipboard", "Failed to copy '" + text + "'");
        }
      );
    } else {
      createErrorPopup("Copy Privatekey", "Please verify your address before you continue.");
    }
  }

  let popup;
  function createErrorPopup(title, error) {
    try {
      // Create popup
      if (popup) {
        f7.popup.close(popup);
        f7.popup.destroy(popup);
        popup = null;
      }

      popup = f7.popup.create({
        content: `
          <div class="popup" id="ErrorPopup">
            <div class="page">
              <div class="navbar">
              <div class="navbar-bg" style="background-color: red"></div>
                <div class="navbar-inner">
                  <div class="title">${title}</div>
                  <div class="right"><a href="#" class="link popup-close">Close</a></div>
                </div>
              </div>
              <div class="page-content">
                <div class="block">
                  <p class="mono-small">${error}</p>
                </div>
              </div>
            </div>
          </div>
        `.trim(),
      });

      // Open it
      popup.open();
    } catch(e) {
      console.log(e);
    }

  }
</script>
