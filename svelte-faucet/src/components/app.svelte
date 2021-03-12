
<svelte:head>
<!--   <link href="https://fonts.googleapis.com/css?family=Open Sans" rel="stylesheet">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/material-design-iconic-font/2.2.0/css/material-design-iconic-font.min.css">
 -->
  <meta charset="UTF-8">
  <meta name="description" content="Smart Contract tutorials">
  <meta name="keywords" content="svelte, solidity, smart contract, energi3, web3, ethers.js, ethers.io, crypto wallets, energi, JavaScript">
  <meta name="author" content="Kim Il">
  <meta name="generator" content="svelte">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</svelte:head>

<App params={ f7params } themeDark>

  <!-- Left panel with cover effect when hidden -->
  <Panel left reveal themeDark visibleBreakpoint={960}>
    <View>
      <Page>
        <Navbar>
            <Link href="https://www.energi.world/" external target="_blank">
              <img src="./images/energi_logo.svg" alt="Energi Logo" height="60%" width="60%"/>
            </Link>
        </Navbar>

        <!-- Navigation links -->
        <Card outline>
          <CardHeader style="{infoStyle}">Start</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/info/" view=".view-main" panelClose title="About" style="{infoStyleLink}"/>
              <ListItem link="https://wallet.test3.energi.network/generate/keystore" external target="_blank" panelClose title="Create a Wallet"/>
            </List>
          </CardContent>
        </Card>
        <br>
        <Card outline>
          <CardHeader style="{web3Style}">Connect</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/web3provider/" view=".view-main" panelClose title="Web3Provider"  style="{web3StyleLink}"/>
            </List>
          </CardContent>
        </Card>
        <br>
        <Card outline>
          <CardHeader style="{scStyle}">Interact</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/smartcontract/" view=".view-main" panelClose title="SmartContract" style="{scStyleLink}"/>
            </List>
          </CardContent>
        </Card>
        <br>
        <br>
        <Card outline>
          <CardHeader>Navigation</CardHeader>
          <CardContent>
            <List>
              <ListItem link="#" view=".view-main" back panelClose title="Back in history"/>
              <ListItem link="/" view=".view-main" panelClose title="Home" style="{homeStyleLink}"/>
            </List>
          </CardContent>
        </Card>

      </Page>
    </View>
  </Panel>

  <!-- Your main view, should have "view-main" class -->
<!--   <View main class="safe-areas" url="/" pushState="true"/> -->

  <!-- Right panel resizeable with reveal effect when hidden -->
  <Panel right reveal themeDark visibleBreakpoint={1360} resizable>
    <View>
      <Page>
        <div class="navbar theme-dark">
          <div class="navbar-bg">
           <!-- <img src="./images/EnergiLogo-Light.png" alt="Energi Logo" height="60%" width="60%" style="left:11px;top:12px;position:absolute;"/> -->
          </div>
        </div>

        <Card outline>
          <CardHeader>Browser</CardHeader>
          <CardContent>
            <List>
              {#if browserDetails.date === 0}
                <ListItem header="No data yet" footer="Waiting for identification"/>
              {:else}
                {#if browserDetails.browser.error}
                  <ListItem title="Unknown browser"/>
                {:else}
                  <ListItem title="{browserDetails.browser.name}" header="Name"/>
                  <ListItem title="{browserDetails.browser.version}" header="Version"/>
                {/if}
                {#if browserDetails.wallet.exists}
                  <ListItem title="Crypto wallet detected" footer="ChainID: {browserDetails.wallet.chain}"/>
                  <ListItem header="Selected Address" footer="{ window.ethereum ? window.ethereum.selectedAddress : 'None'}"/>
                  <ListItem header="Active Address" footer="{ browserDetails.wallet.activeAccount ? browserDetails.wallet.activeAccount : 'None'}"/>
                  {#if browserDetails.wallet.enabled}
                    <ListItem title="Crypto wallet enabled" footer="{browserDetails.wallet.chain}"/>
                    {#if browserDetails.wallet.accounts.length === 0}
                      <ListItem title="No accounts found"/>
                    {:else}
                      {#each browserDetails.wallet.accounts as acc, i}
                        <ListItem title="{acc}" header="Account {i+1}"/>
                      {/each}
                    {/if}
                  {:else}
                    <ListItem title="Crypto wallet disabled"/>
                    {#if browserDetails.wallet.error}
                      <ListItem title="{browserDetails.wallet.error.code}" footer="{browserDetails.wallet.error.reason}"/>
                      <ListItem onClick="{(e) => { e.stopPropagation(); openCryptoWallet(e) }}" title="{browserDetails.todo.text}}" link/>
                    {/if}
                  {/if}
                {:else}
                  <ListItem title="No crypto wallet detected"/>
                {/if}
                <ListItem title="{browserDetails.time}" header="Last Update"/>
              {/if}
            </List>
          </CardContent>
        </Card>
      </Page>
    </View>
  </Panel>

   <!-- Your main view, should have "view-main" class -->
  <View main class="safe-areas" url="/" pushState="true"/>


  <!-- Popup -->
  <Popup id="my-popup">
    <View>
      <Page>
        <Navbar title="Popup">
          <NavRight>
            <Link popupClose>Close</Link>
          </NavRight>
        </Navbar>
        <Block>
          <p>Popup content goes here.</p>
        </Block>
      </Page>
    </View>
  </Popup>

  <LoginScreen id="my-login-screen">
    <View>
      <Page loginScreen>
        <LoginScreenTitle>Login</LoginScreenTitle>
        <List form>
          <ListInput
            type="text"
            name="username"
            placeholder="Your username"
            autocomplete="off"
            value={username}
            onInput={(e) => username = e.target.value}
          />
          <ListInput
            type="password"
            name="password"
            placeholder="Your password"
            autocomplete="current-password"
            value={password}
            onInput={(e) => password = e.target.value}
          />
        </List>
        <List>
          <ListButton title="Sign In" onClick={() => alertLoginData()} />
        </List>
        <BlockFooter>
          Some text about login information.<br />Click "Sign In" to close Login Screen
        </BlockFooter>
      </Page>
    </View>
  </LoginScreen>
</App>
<script>
  import { onMount, onDestroy, setContext } from 'svelte';
  import {
    f7,
    f7ready,
    App,
    Appbar,
    Panel,
    Card,
    CardHeader,
    CardContent,
    Views,
    View,
    Popup,
    Page,
    Navbar,
    Toolbar,
    NavRight,
    Link,
    Icon,
    Block,
    BlockTitle,
    LoginScreen,
    LoginScreenTitle,
    List,
    ListItem,
    ListInput,
    ListButton,
    BlockFooter
  } from 'framework7-svelte';
  import Web3 from 'web3';

  // Framework7 routes
  import routes from '../js/routes';

  import { visitedPages } from '../js/stores.js';
  import { initBrowserDetails, getBrowserDetails, openCryptoWallet } from '../js/utils.js';

  // store the web3 object globally into a svelte context
  const web3 = new Web3('http://localhost:49796');
  setContext('web3', web3);


  let infoStyle = `color:var(--energi-color-text)`;
  let web3Style = infoStyle;
  let scStyle = infoStyle;


  let infoStyleLink = `color:var(--energi-color-text)`;
  let web3StyleLink = infoStyleLink;
  let scStyleLink = infoStyleLink;
  let homeStyleLink = infoStyleLink;

  const unsubscribe = visitedPages.subscribe(visitedPages => {

    infoStyleLink = `color:var(--energi-color-text)`;
    web3StyleLink = infoStyleLink;
    scStyleLink = infoStyleLink;

    if (visitedPages.info) {
      infoStyle = `color:var(--energi-color-green);`;
      infoStyleLink = `color:var(--energi-color-green-secondary);`;
    }
    if (visitedPages.web3) {
      web3Style = `color:var(--energi-color-green);`;
      web3StyleLink = `color:var(--energi-color-green-secondary);`;
    }
    if (visitedPages.smartcontract) {
      scStyle = `color:var(--energi-color-green);`;
      scStyleLink = `color:var(--energi-color-green-secondary);`;
    }
  });

  // Framework7 Parameters
  let f7params = {
    name: 'Faucet', // App name
    theme: 'auto', // Automatic theme detection

    // App routes
    routes: routes,
    // Register service worker
    serviceWorker: {
        path: '/service-worker.js',
    },
    clicks: {
        externalLinks: '.external',
    },
    touch: {
        tapHold: true,
    },
    panel: {
        swipe: true,
        visibleBreakpoint: 1024,
    },
    // Extended by Dialog component:
    dialog: {
        title: 'Faucet',
    },
    // Extended by Statusbar component:
    statusbar: {
        iosOverlaysWebview: true,
    },
  };
  // Login screen demo data
  let username = '';
  let password = '';

  let browserDetails = initBrowserDetails();

  /**
   * Listening for MetaMask address changes.
   * @param  {Function} callback Resolve when address is changed
  */
  if (window.ethereum) {
    window.ethereum.on('accountsChanged', (accounts) => {
      // Time to reload your interface with accounts[0]
      console.log("accountsChanged", accounts);
      browserDetails.wallet.activeAccount = accounts[0];
    });
  }

  const updateBrowserDetails = () => {
    // TODO REMOVE
    //console.log("getBrowserDetails...");
    getBrowserDetails().catch((e) => {
      console.log("ERROR getBrowserDetails", e);
      browserDetails.browser.error = e;
      browserDetails.wallet.error = e;
    }).then((bd) => {
      if (bd) {
        browserDetails = bd;
      } else {
        console.log("ERROR getBrowserDetails undefined");
      }
    });
  }

  function alertLoginData() {
    f7.dialog.alert('Username: ' + username + '<br>Password: ' + password, () => {
      f7.loginScreen.close();
    });
  };

  let clearIv;

  onMount(() => {
    f7ready(() => {
      // Refresh browser details every 60 seconds
      clearIv = setInterval(() => {
        updateBrowserDetails();
      }, 60000);
      // Call F7 APIs here
      updateBrowserDetails();
    });
  });
  onDestroy(() => {
    clearInterval(clearIv);
    unsubscribe();
	});
</script>
