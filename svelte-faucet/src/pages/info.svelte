<Page name="info" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
    <Navbar backLink="Back">
        <NavLeft>
            <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
            <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
        </NavLeft>
        <NavTitle>Information</NavTitle>
    </Navbar>


    <Card outline>
        <CardHeader>About this faucet</CardHeader>
        <CardContent>
            This is the web frontend to the Energi3 <b style="color:green;">Testnet</b> faucet.
            The faucet is implemented as a SmartContract written in solidity and deployed on the Testnet.
            You have to interact with the <b style="color:green;">faucet</b> to get some NRG.
            Each interaction costs a very small fee, called <b style="color:green;">gas</b>. Gas is needed to keep the blockchain going.
        </CardContent>
    </Card>
    <br>
    <Card outline>
        <CardHeader>What is needed ?</CardHeader>
        <CardContent>
            <p>You need an Energi account inside a wallet to store NRGs into it. <b style="color:red;">ATTENTION:</b> Make sure you are on the <b style="color:red;">TESTNET</b> (setting inside the wallet at the upper right).</p>
            <p>
                <Link href="https://wallet.energi.network/generate/keystore" external target="_blank">Create a Wallet</Link>
            </p>
            <p>There are several options to create a wallet, the best is a password encrypted keystore file.
               From the wallet you get a public address (eg</p>
            <p style="color:var(--energi-color-yellow);font-family:'Ubuntu Mono';">"0x771dDB07222A1f9442C91CF04f64F3164771BB62"</p>
            <p> ), this address will receive and hold your NRGs. Its like your bank account number.</p>
        </CardContent>
        <CardFooter>
            <Link href="https://www.google.com/search?q=web3+create+provider" external target="_blank">More Information</Link>

            <div style="display:flex;flex-direction: row;justify-content: flex-end;align-items:center;">
              <Button fill raised href="/web3provider/" view=".view-main" panelClose style="display:flex;align:center;align-items:center;padding-left:15%;bgcolor:var(--energi-color-green)">Next
                <i class="f7-icons size-20" style="margin-left: 3%">arrowtriangle_right</i>
              </Button>
            </div>
        </CardFooter>
    </Card>
</Page>

  <script>
    import { CardFooter, Button, NavLeft, NavTitle, Link, Page, Navbar, Card, CardHeader, CardContent } from 'framework7-svelte';
    import { visitedPages } from '../js/stores.js';

    function onShow() {
        visitedPages.update(() => {
            visitedPages.info = true;
            visitedPages.web3 = false;
            visitedPages.smartcontract = false;
            return visitedPages;
        });
    }
    function onHide()  {
        visitedPages.update(() => {
            visitedPages.info = false;
            return visitedPages;
        });
    }
  </script>
