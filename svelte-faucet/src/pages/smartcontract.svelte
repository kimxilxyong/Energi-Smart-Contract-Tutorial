<Page name="smartcontract" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
    <Navbar backLink="Back">
      <NavLeft>
            <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
            <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
      </NavLeft>
      <NavTitle>Smart Contract</NavTitle>
   </Navbar>

   <Card outline>
    <CardHeader>About Smart Contracts</CardHeader>
    <CardContent>
      What Is a Smart Contract?

      A smart contract is a self-executing contract with the terms of the agreement between buyer and seller being directly written into lines of code. The code and the agreements contained therein exist across a distributed, decentralized blockchain network. The code controls the execution, and transactions are trackable and irreversible.
    </CardContent>
    <CardFooter><Link href="https://www.google.com/search?q=about+smart+contracts" external target="_blank">More Information</Link></CardFooter>
  </Card>
  <br>

</Page>

<script>
    import { NavLeft, NavTitle, Link, Page, Navbar, Card, CardHeader, CardContent, CardFooter } from 'framework7-svelte';
    import { global } from '../js/stores.js';

    function onShow() {
      console.log("OnShow smartcontract");
      global.update(() => {
        global.smartcontract = true;
        global.web3 = false;
        global.info = false;
        return global;
      });
    }
    function onHide() {
        global.update(() => {
            global.smartcontract = false;
            return global;
        });
    }
</script>
