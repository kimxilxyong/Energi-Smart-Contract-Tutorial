<Page>
    <Navbar title="Smart Contract" backLink="Back" />
    <BlockTitle>About Smart Contracts</BlockTitle>
    <Block strong>
      <p>This is the web frontend to the Energi3 </p><p style="color:green;">Testnet</p><p> faucet.</p>
      <p>The faucet is implemented as a SmartContract written in solidity and deployed on the Testnet.</p>
      <p>You have to interact with the </p><p style="color:green;">faucet</p><p> to get some NRG.</p>
      <p>Each interaction costs a very small fee, called </p><p style="color:green;">gas</p><p>. Gas is needed to keep the blockchain going.</p>
    </Block>
    <BlockTitle>Interacting with a Smart Contract</BlockTitle>
    <Block strong>
        <p>You need an Energi account inside a wallet to store NRGs into it. ATTENTION: Make sure you are on the TESTNET (setting upper right).
            <a href="https://wallet.energi.network/generate/keystore">Create a Wallet</a></p>
        <p>There are several options to create a wallet, the best is a password encrypted keystore file.</p>
        <p>From the wallet you get a public address (eg "0x165gue5dfiugnlkjsvfguhiwsdeft5676wegbuhj"), this address will receive and hold the NRGs. Its like your bank account number</p>
      </Block>
      <h1>global.smartcontract = {global.smartcontract}</h1>
    </Page>
  <script>
    import { Page, Navbar, Block, BlockTitle } from 'framework7-svelte';
    import { onMount } from 'svelte';
    import { global } from '../js/stores.js';

    onMount(() => {
        global.update(() => {
            global.smartcontract = true;
            return global;
        });
    });
  </script>
