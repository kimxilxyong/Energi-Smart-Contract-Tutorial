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
        A RPC Provider, also called <b style="color:green;">a Web3</b> Provider, is a Energi Blockchain node, which offers an api that one can use to
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
      <CardHeader>Connection to the official Energi Testnet RPC Provider<br>
      <strong>https://nodeapi.test3.energi.network</strong>
      </CardHeader>
      <CardContent>
      <List noHairlinesMd>
        <ListInput
          outline
          label="URL"
          floatingLabel
          type="url"
          placeholder="https://nodeapi.test3.energi.network"
          clearButton
          bind:value={urlRemote}
        >
        <i class="icon demo-list-icon" slot="media" />
        </ListInput>
       </List>List>
      <Button class="col" fill raised onClick={onConnectRemote}color:var(--energi-color-text)>Connect</Button>

      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a></CardFooter>
    </Card>
</Page>

  <script>
    import { NavLeft, NavTitle, Link, Block, Card, CardHeader, CardContent, CardFooter, Page, Navbar, List, ListInput, Button } from 'framework7-svelte';
    import { global } from '../js/stores.js';

    let urlRemote;

    function onConnectRemote() {
      console.log("onConnectRemote to " + urlRemote);
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
