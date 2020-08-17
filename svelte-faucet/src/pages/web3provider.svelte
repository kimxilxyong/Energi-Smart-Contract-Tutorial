<Page>
    <Navbar title="Connect to the Energi Testnet" backLink="Back" />
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
      <CardHeader>Connect to an RPC Providers</CardHeader>
      <CardContent>
        <List>
          LIST
          <ListItem>Item1
            <Block>
              <p>Popup content goes here.</p>
            </Block>
          </ListItem>
          <ListItem>Item2
            <Block>
              <p>Popup content goes here.</p>
            </Block>
          </ListItem>
        </List>
      </CardContent>
      <CardFooter><a href="https://www.google.com/search?q=web3+create+provider"  external target="_blank">More Information</a></CardFooter>
    </Card>





</Page>
  <script>
    import { Block, Card, CardHeader, CardContent, CardFooter, Page, Navbar, Link, List, ListItem } from 'framework7-svelte';
    import { onMount } from 'svelte';
    import { global } from '../js/stores.js';

    onMount(() => {
        global.update(() => {
            global.web3 = true;
            return global;
        });
    });
  </script>
