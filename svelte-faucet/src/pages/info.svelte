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
            <p>You need an Energi account inside a wallet to store NRGs into it. <b style="color:red;">ATTENTION:</b> Make sure you are on the <b style="color:red;">TESTNET</b> (setting upper right).</p>
            <p><a href="https://wallet.energi.network/generate/keystore" external target="_blank">Create a Wallet</a></p>
            <p>There are several options to create a wallet, the best is a password encrypted keystore file.
            From the wallet you get a public address (eg "0x165gue5dfiugnlkjsvfguhiwsdeft5676wegbuhj"), this address will receive and hold the NRGs. Its like your bank account number.</p>
        </CardContent>
    </Card>
</Page>

  <script>
    import { NavLeft, NavTitle, Link, Page, Navbar, Card, CardHeader, CardContent } from 'framework7-svelte';
    import { global } from '../js/stores.js';

    function onShow() {
        global.update(() => {
            global.info = true;
            global.web3 = false;
            global.smartcontract = false;
            return global;
        });
    }
    function onHide()  {
        global.update(() => {
            global.info = false;
            return global;
        });
    }
  </script>
