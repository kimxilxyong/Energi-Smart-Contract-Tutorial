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

      <Row noGap>
        <Col width="80">
          <List noHairlinesMd>
            <ListInput
              outline
              label="URL of JSON-RPC"
              floatingLabel
              type="url"
              placeholder="https://nodeapi.test3.energi.network"
              autocomplete="https://nodeapi.test3.energi.network"
              clearButton
              resizable
              onInput={ e => onInput(e) }
            >
            <!-- <i class="icon demo-list-icon" slot="media" /> -->
            </ListInput>
          </List>
        </Col>
        <Col width="20">
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
              label="Log"
              type="textarea"
              clearButton
              resizable

            >
            <!-- <i class="icon demo-list-icon" slot="media" /> -->
            </ListInput>
          </List>
        </Col>
      </Row>
      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a></CardFooter>
    </Card>
</Page>

  <script>
    import { f7, NavLeft, NavTitle, Link, Row, Col, Card, CardHeader, CardContent, CardFooter, Page, Navbar,Input, List, ListInput, Button } from 'framework7-svelte';
    import { global } from '../js/stores.js';
    //let Web3 = require('web3');

    let urlRemote;
    let connectedRemote = false;

    function onInput(e) {
      console.log("onInput: " + e.target.value );
      urlRemote = e.target.value;
    }
    function onConnectRemote() {
      console.log("onConnectRemote to " + urlRemote);

      f7.dialog.preloader('Connecting ' + urlRemote);
      setTimeout(() => {
        connectedRemote = ! connectedRemote;
        f7.dialog.close();
      }, 2000);
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
