<App params={ f7params } themeDark>

  <!-- Left panel with cover effect when hidden -->
  <Panel left cover themeDark visibleBreakpoint={960}>
    <View>
      <Page>
        <Navbar title="Faucet" bgColor="green"/>

        <Card outline>
          <CardHeader style="{infoStyle}">Information</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/info/" view=".view-main" panelClose title="About"/>
              <ListItem link="https://wallet.energi.network/generate/keystore" external target="_blank" panelClose title="Create a Wallet"/>
            </List>
          </CardContent>
        </Card>
        <br>

        <Card outline>
          <CardHeader style="{web3Style}">Connect</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/web3provider/" view=".view-main" panelClose title="Web3Provider"/>
            </List>
          </CardContent>
        </Card>
        <br>

        <Card outline>
          <CardHeader style="{scStyle}">Interact</CardHeader>
          <CardContent>
            <List>
              <ListItem link="/smartcontract/" view=".view-main" panelClose title="SmartContract"/>
            </List>
          </CardContent>
        </Card>
        <br>
        <br>
        <Card outline align=botton>
          <CardHeader>Navigation</CardHeader>
          <CardContent>
            <List>
              <ListItem link="#" view=".view-main" back panelClose title="Back in history"/>
            </List>
          </CardContent>
        </Card>

      </Page>
    </View>
  </Panel>


  <!-- Right panel with reveal effect-->
  <Panel right reveal themeDark>
    <View>
      <Page>
        <Navbar title="Right Panel"/>
        <Block>Right panel content goes here</Block>
      </Page>
    </View>
  </Panel>


  <!-- Your main view, should have "view-main" class -->
  <View main class="safe-areas" url="/" />


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
            value={username}
            onInput={(e) => username = e.target.value}
          />
          <ListInput
            type="password"
            name="password"
            placeholder="Your password"
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
  import { onMount, onDestroy } from 'svelte';

  import {
    f7,
    f7ready,
    App,
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

  import routes from '../js/routes';
  import { global } from '../js/stores.js';


  let infoStyle = ``;
  //let web3Visited = false;
  let web3Style = "";
  //let scVisited = false;
  let scStyle = ``;

  const unsubscribe = global.subscribe(global => {

    if (global.info) {
      infoStyle = `color:green;`;
    }
    if (global.web3) {
      web3Style = `color:green;`;
    }
    if (global.smartcontract) {
      scStyle = `color:green;`;
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
  };
  // Login screen demo data
  let username = '';
  let password = '';


  function alertLoginData() {
    f7.dialog.alert('Username: ' + username + '<br>Password: ' + password, () => {
      f7.loginScreen.close();
    });
  }
  onMount(() => {
    f7ready(() => {

      // Call F7 APIs here
    });
  });
  onDestroy(() => {
    unsubscribe();
	});
</script>
