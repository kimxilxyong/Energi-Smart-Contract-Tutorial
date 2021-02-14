<!-- SPDX-License-Identifier: MIT -->
<!-- Copyright (c) 2020 Kim Il Yong -->
<!-- Version 1.0.0 -->

<Page name="info" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
    <Navbar backLink="Back">
        <NavLeft>
            <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
            <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
        </NavLeft>
        <NavTitle>Information</NavTitle>
        <NavRight>
            <Link panelOpen="right" class="f7-icons">square_righthalf_fill</Link>
        </NavRight>
    </Navbar>

    <Card outline>
        <CardHeader>About this faucet</CardHeader>
        <CardContent>
            This is the web frontend to the Energi3 <Link href="https://explorer.test3.energi.network" external target="_blank">Testnet</Link> faucet.
            The faucet is implemented as a SmartContract written in solidity and deployed on the Testnet.
            You have to interact with the <Link href="https://explorer.test3.energi.network/address/{$faucetAddress}/transactions" external target="_blank">faucet</Link> to get some NRG.
            Each interaction costs a very small fee, called <Link href="https://www.google.com/search?q=what+is+gas+on+a+blockchain" external target="_blank">gas</Link>. Gas is needed to keep the blockchain going.
            <p>Don't worry, gas will be provided</p>
        </CardContent>
    </Card>
    <br>
    <Card outline>
        <CardHeader>What is needed ?</CardHeader>
        <CardContent>
            <p>You need an Energi account inside a wallet to store NRGs into it. <i style="color:coral;padding-top:2px;">ATTENTION:</i> Make sure you are on the <Link href="https://explorer.test3.energi.network" external target="_blank">TESTNET</Link> (setting inside the wallet at the upper right).</p>
            <p>
                <Link href="https://wallet.test3.energi.network/generate/keystore" external target="_blank">Create a keystore Wallet</Link>
            </p>
            <p>There are several options to create a wallet, the best is a <Link href="https://medium.com/@julien.maffre/what-is-an-ethereum-keystore-file-86c8c5917b97" external target="_blank">password encrypted keystore file</Link>.
                Store this keystore file in a <Link href="https://www.google.com/search?q=save+place+to+store+the+ethereum+keystore+file" external target="_blank">safe place</Link>, for example on <Link href="https://protonmail.com/blog/usb-encryption/" external target="_blank">encrypted usb drives</Link>.
               Use linux or mac, don't use windows, if you are handling heaps of money, use your 🧠 😁</p>
            <p>From the wallet you get a public address (eg <i style="color:var(--energi-color-yellow);font-family:'Ubuntu Mono';">"0x771dDB07222A1f9442C91CF04f64F3164771BB62"</i>), this address will receive and hold your NRGs. Its like your bank account number.</p>
        </CardContent>
        <CardFooter>
            <Link href="https://www.google.com/search?q=account+wallet+keystore+privatekey" external target="_blank">More Information</Link>

            <Button fill raised href="/web3provider/" view=".view-main" panelClose>
              <i class="f7-icons button-icon">arrowtriangle_right</i>
              Next
            </Button>
        </CardFooter>
    </Card>
</Page>

  <script>
    import { CardFooter, Button, NavRight,  NavLeft, NavTitle, Link, Page, Navbar, Card, CardHeader, CardContent } from 'framework7-svelte';
    import { faucetAddress, visitedPages } from '../js/stores.js';

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
