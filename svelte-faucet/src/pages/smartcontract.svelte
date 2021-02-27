<!-- SPDX-License-Identifier: MIT -->
<!-- Copyright (c) 2020 Kim Il Yong -->
<!-- Version 1.0.0 -->

<Page name="smartcontract" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
  <Navbar backLink="Back">
    <NavLeft>
      <!--  <Link iconIos="f7:back" iconAurora="f7:back" iconMd="material:back" backLink="Back" /> -->
      <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
    </NavLeft>
    <NavTitle>Smart Contract</NavTitle>
    <NavRight>
      <Link panelOpen="right" class="f7-icons">square_righthalf_fill</Link>
    </NavRight>
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
  <Card outline>
    <CardHeader >Energi Testnet faucet Smart-Contract details</CardHeader>
    <CardContent>
      <List>
        {#if faucetInfo.version === 0}
          <ListItem title="Loading from {$web3URL} ...">
            <div slot="media" class="{spin_css}">🌼</div>
          </ListItem>
        {:else}
          <ListItem title="Version: { faucetInfo.version}"></ListItem>
          <ListItem class="colorturkis" title="Balance: { ethers.utils.formatEther(faucetInfo.balance) } NRG"></ListItem>
                  <!-- OK-Square U+F634  clipboard U+F429 -->
          <ListItem title="Contract Address: {faucetInfo.contractAddress}" after="&nbsp;">
            <div slot="after-title" class="colortext" style="margin-left: 20px;">
              {#if faucetInfo.contractAddress !== ethers.constants.AddressZero}
                <span class="bround bgreen fss3 {isClipboardDoneClass}">
                Copied!
                </span>
                <i on:click={ () => {
                                      isClipboardDoneClass = "animate__animated animate__fadeOut animate__delay-2s";
                                      copyToClipboard(faucetInfo.contractAddress);
                                      setTimeout(() => {
                                                          if (isClipboardDoneClass !== "hidden") {
                                                            isClipboardDoneClass = "hidden";
                                                          }
                                                        }, 3300);
                                    }
                            } title="Copy to clipboard" class="nf-icons nf-hc-1x colorgrey">&#xF429;
                </i>
              {/if}
            </div>
        </ListItem>
        <ListItem footer="Last recipient" title="{faucetInfo.lastRecipient.name}" header="Got {faucetInfo.lastRecipient.amount} NRG { timeDifference(new Date(), faucetInfo.lastRecipient.date)}">
          <i slot="media">{faucetInfo.lastRecipient.country} </i>
        </ListItem>
      {/if}
      </List>
    </CardContent>
    <CardFooter>
      Node Address: {$web3URL}
    </CardFooter>
  </Card>

  <Card outline>
    <CardHeader>Giver or Taker</CardHeader>
      <CardContent>
        <List>
          {#if (recipient.balance === ethers.constants.Zero && donor.balance === ethers.constants.Zero)}
            <ListItem header="Which one are you?" title="The all-knowing, all-seeing Trash Heap has not spoken yet" link="https://www.youtube.com/watch?v=cZKc4D86ctE" target="_blank" external>
              <div slot=media>👳‍♀</div>
            </ListItem>
          {:else}
            <ListItem header="Which one are you?" title="The all-knowing, all-seeing Trash Heap said: You are a {userIsDonor ? 'Donor':'Recipient'}" link="https://www.youtube.com/watch?v=cZKc4D86ctE" target="_blank" external>
              <div slot=media>👳‍♀</div>
            </ListItem>
          {/if}
          <Segmented raised>
            <Button tabLink="#tabRecipient" class='{userIsDonor ? "":"tab-link-active"}' on:click="{(e) => {userSelfSelected = true;}}">Recipient</Button>
            <Button tabLink="#tabDonor" class='{userIsDonor ? "tab-link-active":""}' on:click="{(e) => {userSelfSelected = true;}}">Donor</Button>
          </Segmented>
        </List>
      </CardContent>
    <CardFooter> </CardFooter>
  </Card>

  <Tabs>
    <!-- class="page-content" -->
    <Tab id="tabRecipient" class='{userIsDonor ? "":"tab-active"}'>
      <Card outline>
        <CardHeader>Recipient</CardHeader>
        <CardContent>
          <List>

            {#if !recipient.address || recipient.address === ethers.constants.AddressZero || manualRecipientAddressEntry === true}
              <ListItem link="/web3provider/" view=".view-main" panelClose header="No or invalid recipient address" title="Please go back to provider/account selection (Web3Provider, step 2) to select a valid address" footer="Click here to go to step 2">
                <div slot="media">🥵</div>
              </ListItem>

              <List accordionList accordionOpposite>
                <ListItem accordionItem title="I know what i'm doing" id="AccordiumItemRecipient">
                  <AccordionItem>
                    <AccordionContent>

                      <!-- name="Nr of NRGs" inputReadonly style="margin-right: 10px;" -->
                      <ListInput value="{recipient.name}" onInput="{(e) => recipientNameChanged(e)}" onBlur="{(e) => recipientNameChanged(e)}" onInputClear="{()=>{recipient.name='';}}" label="Please tell us your nick" info="Your name, can by any string, min 3 characters" errorMessage="Not empty please!"
                        outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                        inputId="inputIDRecipientNameWW">
                        <div slot="media">{recipient.country}</div>
                        <span slot="content-end">
                          <small>NRG amount</small>
                          <Stepper value={recipient.amount} style="background-color: aquamarine;" class="font07"
                                   onStepperChange={(v) => {recipient.amount = v}} max="9999"
                                   small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                        </span>
                      </ListInput>
                      <ListInput label="Energi Account address" type="text" placeholder="Enter account address" info="42 char hexnumber starting with 0x" errorMessage="Only hexnumbers starting with 0x please!"
                          clearButton outline required floatingLabel
                          value="{recipient.address}"
                          onInputClear="{() => { manualRecipientAddressEntry = false; recipient.address = ''; }}"
                          onInput ="{(e) => { manualRecipientAddressEntry = true; recipientAddressChanged(e); }}"
                          onBlur="{(e) => { manualRecipientAddressEntry = true; recipientAddressChanged(e); }}"
                          validate pattern="^0x[a-fA-F0-9]*"
                          onValidate="{(isValid) => { if (!isValid) {recipient.isAddressValid = false}; isValid = false; return false;}}"
                          inputId="inputIDmanualRecipientAccount" class="mono-small" spellcheck="false" autocomplete="on">
                        <div slot="media">📫</div>
                      </ListInput>

                      <ListItem class="seperator"></ListItem>

                      <ListItem header="I want to execute the requesting TX with the Webwallet" title="Request {recipient.amount} NRGs to {recipient.address}">
                        <img slot="media" src="NRG16x16.png" alt="NRG" width="16" height="16">
                        <span slot="after" style="display:flex;flex-direction:column;top:-0.5em;position:relative;">
                          <small>Confirmations</small>
                          <Stepper name="Confirmations" value={confirmations} style="background-color: aquamarine;"
                                    onStepperChange={(v) => {confirmations = v}} max="100"
                                    small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                        </span>
                        <div slot="after-end">
                          <Button
                            fill
                            raised
                            class={requestButtonEnabled ? "" : "disablebar"} style="margin-left:0.8em;margin-top:0.35em;bgcolor:{requestButtonEnabled ? 'var(--energi-color-green)' : 'var(--energi-color-grey)'}"
                            on:click="{(e) => {if (requestButtonEnabled) {requestDonationWithWebWallet();}}}"
                          >
                            <i class="f7-icons button-icon">arrow_down_circle</i>
                            Request
                          </Button>
                        </div>
                      </ListItem>

                      <ListInput value="{recipient.txhash}" onInput="{(e) => recipientTxChanged(e)}" onBlur="{(e) => recipientTxChanged(e)}" onInputClear="{()=>{recipient.txhash='';}}"
                                  label="Transaction hash" info="The transaction you want to check" errorMessage="Not empty please!" outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                                  inputId="inputIDcheckTXRecipient">
                        <div slot="media">🈺</div>
                        <div slot="content-end">
                          <Button
                            fill
                            raised
                            class={recipientTxCheckButtonEnabled ? "" : "disablebar"}
                            style="margin-right:0.8em;margin-top:0.4em;"
                            on:click="{(e) => {if (recipientTxCheckButtonEnabled) {checkTransaction(recipient.txhash);}}}"
                          >
                            <i class="f7-icons button-icon">camera</i>
                            Check
                          </Button>
                        </div>
                      </ListInput>

                    </AccordionContent>
                  </AccordionItem>
                </ListItem>
              </List>

            {:else}
              <ListItem header="Reciving address" title="{recipient.address}" after="Balance {ethers.utils.formatEther(recipient.balance)} NRG">
                <div slot="media">👼</div>
              </ListItem>

              <!-- name="Nr of NRGs" inputReadonly style="margin-right: 10px;" -->
              <ListInput value="{recipient.name}" onInput="{(e) => recipientNameChanged(e)}" onBlur="{(e) => recipientNameChanged(e)}" onInputClear="{()=>{recipient.name='';}}" label="Your Name" info="Your name, can by any string, min 3 characters" errorMessage="Not empty please!"
                outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                inputId="inputIDRecipientName">
                <div slot="media">{recipient.country}</div>
                <span slot="content-end">
                  <small>NRG amount</small>
                    <Stepper value={recipient.amount} style="background-color: aquamarine;" class="font07"
                            onStepperChange={(v) => {recipient.amount = v}} max="9999"
                            small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                </span>
              </ListInput>

              <ListItem header="Start transaction" title="Request {recipient.amount} NRGs to {recipient.address}">
                <img slot="media" src="NRG16x16.png" alt="NRG" width="16" height="16">
                <div slot="after">
                  <Button
                    fill
                    raised
                    class={requestButtonEnabled ? "" : "disablebar"} style="margin-left:0.8em;margin-top:0.35em;bgcolor:{requestButtonEnabled ? 'var(--energi-color-green)' : 'var(--energi-color-grey)'}"
                    on:click="{(e) => {if (requestButtonEnabled) {requestDonation();}}}"
                  >
                    <i class="f7-icons button-icon">arrow_down_circle</i>
                    Request
                  </Button>
                </div>
              </ListItem>
            {/if}
            {#if transactionReciepts.length > 0}
              <ListItem  header="{spin_css_tx ? ( transactionReciepts[0].blockNumber !== -1 ? 'Waiting for confirmations of transaction ...' : 'Waiting for transaction to be mined ...' ) : 'Transaction finished'}" title="TX: {transactionReciepts[0].transactionHash}"
                          footer="{spin_css_tx ? `In progress, please wait for ${confirmations} confirmations`:''}" after="&nbsp;">
                <div slot="media" class="{spin_css_tx}">⏳</div>
                <div slot="after-title" class="colortext" style="margin-left: 20px;">
                  <span class="bround bgreen fss3 {isClipboardDoneClass}">Copied!</span>
                  <i on:click={ () => {
                                        isClipboardDoneClass = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(transactionReciepts[0].transactionHash);
                                        setTimeout(() => {
                                                            if (isClipboardDoneClass !== "hidden") {
                                                              isClipboardDoneClass = "hidden";
                                                            }
                                                          }, 3300);
                                      }
                              } title="Copy to clipboard" class="nf-icons nf-hc-1x colorgrey">&#xF429;
                  </i>
                </div>
              </ListItem>
              {#if transactionReciepts[0].blockNumber > 0}
                <ListItem  header="Transaction was mined in block {transactionReciepts[0].blockNumber}" title="Time: {transactionReciepts[0].time}, Value: {transactionReciepts[0].value} NRG" footer="Gas used: {transactionReciepts[0].gasUsed}">
                  <div slot="media">⚒️</div>
                </ListItem>
              {/if}
              {#each transactionReciepts as tr , i}
                {#if i > 0 && tr.confirmations > 0 && tr.blockNumber > 0}
                  <!-- transactionHash -->
                  <ListItem  header="{tr.confirmations} Confirmation(s) for block {tr.blockNumber}" title="Time: {tr.time}, Value: {tr.value} NRG" footer="Gas used: {tr.gasUsed}">
                    <div slot="media">✅</div>
                  </ListItem>
                {/if}
              {/each}
            {/if}
          </List>

        </CardContent>
        <CardFooter>AccountStore: {$ethAccount.accountStore} at {$web3URL} <Link on:click={(e) => {requestGas(e);}}>Request gas</Link></CardFooter>
      </Card>
    </Tab>
    <!-- class="page-content" -->
    <Tab id="tabDonor" class='{userIsDonor ? "tab-active":""}'>
      <Card outline>
        <CardHeader>Donor</CardHeader>
        <CardContent>
          <List>

            {#if !donor.address || donor.address === ethers.constants.AddressZero}
              <ListItem link="/web3provider/" view=".view-main" panelClose header="No or invalid donor address" title="Please go back to provider/account selection (Web3Provider, step 2) to select a valid address" footer="Click here to go to step 2">
                <div slot="media">🥵</div>
              </ListItem>

              <List accordionList accordionOpposite>
                <ListItem accordionItem title="I know what i'm doing" id="AccordiumItemDonor">
                  <AccordionItem>
                    <AccordionContent>

                      <!-- name="Nr of NRGs" inputReadonly style="margin-right: 10px;" -->
                      <ListInput value="{donor.name}" onInput="{(e) => donorNameChanged(e)}" onBlur="{(e) => donorNameChanged(e)}" onInputClear="{()=>{donor.name='';}}" label="Please tell us your nick" info="Your name, can by any string, min 3 characters" errorMessage="Not empty please!"
                        outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                        inputId="inputIDDonorNameWW">
                        <div slot="media">{donor.country}</div>
                        <span slot="content-end">
                          <small>NRG amount</small>
                            <Stepper value={donor.amount} style="background-color: aquamarine;" class="font07"
                                    onStepperChange={(v) => {donor.amount = v}} max="9999"
                                    small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                      </span>
                      </ListInput>

                      <ListItem class="seperator"></ListItem>

                      <ListItem header="I want to execute the donation TX with the Webwallet" title="Send a donation of {donor.amount} NRGs to {faucetInfo.contractAddress}">
                        <img slot="media" src="NRG16x16.png" alt="NRG" width="16" height="16">
                        <span slot="after" style="display:flex;flex-direction:column;top:-0.5em;position:relative;">
                          <small>Confirmations</small>
                          <Stepper name="Confirmations" value={confirmations} style="background-color: aquamarine;"
                                    onStepperChange={(v) => {confirmations = v}} max="100"
                                    small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                        </span>
                        <div slot="after-end">
                          <Button
                            fill
                            raised
                            class={sendButtonEnabled ? "" : "disablebar"}
                            style="margin-left:0.8em;margin-top:0.35em;bgcolor:{sendButtonEnabled ? 'var(--energi-color-green)' : 'var(--energi-color-grey)'}"
                            on:click="{(e) => { if (sendButtonEnabled) {sendDonationWithWebWallet();} }}"
                          >
                            <i class="f7-icons button-icon">arrow_up_circle</i>
                            Send
                          </Button>
                        </div>
                      </ListItem>

                      <ListInput value="{donor.txhash}" onInput="{(e) => donorTxChanged(e)}" onBlur="{(e) => donorTxChanged(e)}" onInputClear="{()=>{donor.txhash='';}}"
                                  label="Transaction hash" info="The transaction you want to check" errorMessage="Not empty please!" outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                                  inputId="inputIDcheckTXDonor">
                        <div slot="media">🈺</div>
                        <div slot="content-end">
                          <Button
                            fill
                            raised
                            class={donorTxCheckButtonEnabled ? "" : "disablebar"}
                            style="margin-right:0.8em;margin-top:0.4em;"
                            on:click="{(e) => {if (donorTxCheckButtonEnabled) { checkTransaction(donor.txhash); }}}"
                          >
                            <i class="f7-icons button-icon">camera</i>
                            Check
                          </Button>
                        </div>
                      </ListInput>

                    </AccordionContent>
                  </AccordionItem>
                </ListItem>
              </List>

            {:else}

              <ListItem header="Donating address" title="{donor.address}" after="Balance {ethers.utils.formatEther(donor.balance)} NRG">
                <div slot="media">🎅🏻</div>
              </ListItem>

              <!-- name="Nr of NRGs" inputReadonly style="margin-right: 10px;" -->
              <ListInput value="{donor.name}" onInput="{(e) => donorNameChanged(e)}" onBlur="{(e) => donorNameChanged(e)}" onInputClear="{()=>{donor.name='';}}" label="Your Name" info="Your name, can by any string, min 3 characters" errorMessage="Not empty please!"
                outline floatingLabel type="text" clearButton spellcheck="false" autocomplete="off"
                inputId="inputIDDonorName">
                <div slot="media">{donor.country}</div>
                <span slot="content-end">
                  <small>NRG amount</small>
                  <Stepper value={donor.amount} style="background-color: aquamarine;" class="font07"
                            onStepperChange={(v) => {donor.amount = v}} max="9999"
                            small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                </span>
              </ListInput>

              <ListItem header="Start transaction" title="Send a donation of {donor.amount} NRGs to {faucetInfo.contractAddress}">
                <img slot="media" src="NRG16x16.png" alt="NRG" width="16" height="16">
                <span slot="after" style="display:flex;flex-direction:column;top:-0.5em;position:relative;">
                  <small>Confirmations</small>
                  <Stepper name="Confirmations" value={confirmations} style="background-color: aquamarine;"
                            onStepperChange={(v) => {confirmations = v}} max="100"
                            small wraps autorepeat autorepeatDynamic autocomplete="off"/>
                </span>
                <div slot="after-end">
                  <Button
                    fill
                    raised
                    class={sendButtonEnabled ? "" : "disablebar"}
                    style="margin-left:0.8em;margin-top:0.35em;bgcolor:{sendButtonEnabled ? 'var(--energi-color-green)' : 'var(--energi-color-grey)'}"
                    on:click="{(e) => { if (sendButtonEnabled) {sendDonation();} }}"
                  >
                    <i class="f7-icons button-icon">arrow_up_circle</i>
                    Send
                  </Button>
                </div>
              </ListItem>
            {/if}

            {#if transactionReciepts.length > 0}
              <ListItem  header="{spin_css_tx ? ( transactionReciepts[0].blockNumber !== -1 ? 'Waiting for confirmations of transaction ...' : 'Waiting for transaction to be mined ...' ) : 'Transaction finished'}" title="TX: {transactionReciepts[0].transactionHash}"
                         footer="{spin_css_tx ? `In progress, please wait for ${confirmations} confirmations`:''}" after="&nbsp;">
                <div slot="media" class="{spin_css_tx}">⏳</div>
                <div slot="after-title" class="colortext" style="margin-left: 20px;">
                  <span class="bround bgreen fss3 {isClipboardDoneClass}">Copied!</span>
                  <i on:click={ () => {
                                        isClipboardDoneClass = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(transactionReciepts[0].transactionHash);
                                        setTimeout(() => {
                                                            if (isClipboardDoneClass !== "hidden") {
                                                              isClipboardDoneClass = "hidden";
                                                            }
                                                          }, 3300);
                                      }
                              } title="Copy to clipboard" class="nf-icons nf-hc-1x colorgrey">&#xF429;
                  </i>
                </div>
              </ListItem>
              {#if transactionReciepts[0].blockNumber > 0}
                <ListItem  header="Transaction was mined in block {transactionReciepts[0].blockNumber}" title="Time: {transactionReciepts[0].time}, Value: {transactionReciepts[0].value} NRG" footer="Gas used: {transactionReciepts[0].gasUsed}">
                  <div slot="media">⚒️</div>
                </ListItem>
              {/if}
              {#each transactionReciepts as tr , i}
                {#if i > 0 && tr.confirmations > 0 && tr.blockNumber > 0}
                  <!-- transactionHash -->
                  <ListItem  header="{tr.confirmations} Confirmation(s) for block {tr.blockNumber}" title="Time: {tr.time}, Value: {tr.value} NRG" footer="Gas used: {tr.gasUsed}">
                    <div slot="media">✅</div>
                  </ListItem>
                {/if}
              {/each}
            {/if}
          </List>

        </CardContent>
        <CardFooter><Link href="https://www.google.com/search?q=about+smart+contracts" external target="_blank">More Information</Link></CardFooter>
      </Card>
    </Tab>
  </Tabs>

  <Card outline>
    <CardHeader>Log</CardHeader>
    <CardContent>
      <List>
        <ListInput
          outline
          disabled
          floatingLabel
          label="Log"
          type="textarea"
          clearButton
          resizable
          value="{logText}"
          onInputClear="{() => { logText = ''; }}"
          id="yellow-mono-small"
          autocomplete="off"
          spellcheck="false">
            <i slot="media" style="color: white;" class="nf-icons nf-hc-fw">&#xf02d;</i>
      </ListInput>
      <ListItem class="seperator"></ListItem>
      </List>
    </CardContent>
    <CardFooter><Link href="https://www.google.com/search?q=about+smart+contracts" external target="_blank">More Information</Link></CardFooter>
  </Card>

</Page>

<script>
import { f7, AccordionContent, AccordionItem, Block, NavRight, NavLeft, NavTitle, Link, Page, Navbar, Subnavbar, Button, Segmented, Card, CardHeader, CardContent, CardFooter, List, ListItem, ListInput, Tabs, Tab, Stepper } from 'framework7-svelte';
import { visitedPages } from '../js/stores.js';
import { ethers } from "ethers";
import { tick, getContext, onMount, onDestroy } from "svelte";
import {
  web3URL,
  faucetAddress,
  ethAccount,
  gasDonorAddress
} from "../js/stores.js";
import { errorToReason, errorToCode, getCountryFlag, timeDifference, copyToClipboard, getTime, scrollTo, sleep } from "../js/utils.js";
import { buildCallrequestDonationArguments, buildCallDonationArguments, callRequest, callDonation, fetchFaucetDetailsAsync, fetchFaucetDetails } from "../js/web3utils.js";

import {abi106} from '../js/abi.js';
const abi = abi106;

const dividerInNRG = "500";

let faucetInfo = {
  version: 0,
  balance: ethers.constants.Zero,
  calculatedBalance: ethers.constants.Zero,
  donorsCount: 0,
  recipientsCount: 0,
  owner: ethers.constants.AddressZero,
  contractAddress: $faucetAddress,
  date: new Date(),
  error: 'none yet',
  lastRecipient: {name:"No one", address: ethers.constants.AddressZero, count: 0, amount: ethers.constants.Zero, country: "🗺", date: new Date(),},
};

let recipient = {
  name: "",
  nameOverwritten: false,
  country: "🗺",
  balance: ethers.constants.Zero,
  amount: 100,
  address: ethers.constants.AddressZero,
  isAddressValid: false,
  signer: null,
};

let donor = {
  name: "unknown benefactor",
  nameOverwritten: false,
  country: "🗺",
  balance: ethers.constants.Zero,
  amount: 1000,
  address: ethers.constants.AddressZero,
  isAddressValid: false,
  signer: null,
};

let result = {
    status: false,
    transactionHash: ethers.constants.HashZero,
    blockNumber: 0,
    blockHash: ethers.constants.HashZero,
    gasUsed: 0,
    confirmations: 0,
    logsBloom: ethers.constants.HashZero,
    sdnTX: ethers.constants.HashZero,
    error: null,
    errorReason: "",
    errorCode: "",
    errorStep: "",
  };

let transactionReciepts = [];
let confirmations = 3;
let userIsDonor = false;
let userSelfSelected = false;
let manualRecipientAddressEntry = false;

let logText = "";
let spin_css = "nf-hc-spin";
let spin_css_tx;
let isClipboardDoneClass = "hidden";
let requestButtonEnabled = false;
let sendButtonEnabled = false;
let donorTxCheckButtonEnabled = false;
let recipientTxCheckButtonEnabled = false;


// Events ------------------------------
function recipientNameChanged(e) {
  recipient.name = e.target.value;
  recipient.nameOverwritten = true;
  if (recipient.isAddressValid && recipient.name.length > 2) {
    requestButtonEnabled = true;
  } else {
    requestButtonEnabled = false;
  }
}
function recipientAddressChanged(e) {
  recipient.address = e.target.value;
  if (ethers.utils.isAddress(recipient.address)) {
    recipient.address = ethers.utils.getAddress(recipient.address);
    recipient.isAddressValid = true;
  } else {
    recipient.isAddressValid = false;
  }
  if (recipient.isAddressValid && recipient.name.length > 2) {
    requestButtonEnabled = true;
  } else {
    requestButtonEnabled = false;
  }
}
function recipientTxChanged(e) {
  recipient.txhash = e.target.value;
  if (recipient.txhash.length === 66) {
    recipientTxCheckButtonEnabled = true;
  } else {
    recipientTxCheckButtonEnabled = false;
  }
}

function donorNameChanged(e) {
  donor.name = e.target.value;
  donor.nameOverwritten = true;
  if (donor.name.length > 2) {
    sendButtonEnabled = true;
  } else {
    sendButtonEnabled = false;
  }
}
function donorTxChanged(e) {
  donor.txhash = e.target.value;
  if (donor.txhash.length === 66) {
    donorTxCheckButtonEnabled = true;
  } else {
    donorTxCheckButtonEnabled = false;
  }
}
// END Events --------------------------

async function requestGas(e) {
  try {
    console.log("Request gas event:", e);
    e.stopPropagation();

    if (recipient.address !== ethers.constants.AddressZero) {
      let response = await fetch('/ajax/requestGas?toAdr=' + recipient.address);
      if (response.ok) {
        const json = await response.json();
        console.log("Gas Response:", json);

        if (json.error && json.error.code && json.error.message) {
          if (json.error.code === -32000 && json.error.message === "authentication needed: password or unlock") {
            createErrorPopup("NO_GAS_DONOR", "Account for gas donation is locked. Tell the admin to unlock it!", true);
          } else if (json.error.code === 1) {
            createErrorPopup("GAS_DONATION", "Server said: " + json.error.message, false);
          } else {
            createErrorPopup("GAS_SERVER_ERROR", json.error.code + ", " + json.error.message, true);
          }
        }


      } else {
        console.log("Request gas failed:", response);
      }
    } else {
      createErrorPopup("INVALID_PARAMETER", "Recipient address not set", true);
    }
  } catch (e) {
    console.log("ERROR:", e);
    return Promise.reject(e);
  }
}

/**
* requestDonation() calls the contract function requestDonation(address payable _to, uint _amount, string memory _name, bytes8 _country) public
*
* @param none
* @return nothing
*/
async function requestDonation() {

  const startSeconds = new Date();
  try {

    logText = logText + "\n" + getTime() + ": Requesting " + recipient.amount + " NRG from faucet for " + recipient.name + " in " + recipient.country + " to address '" + recipient.address + "'";
    transactionReciepts = [];
    spin_css_tx = "nf-hc-spin";

    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    const signer = provider.getSigner(recipient.address);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

    // The Contract object
    const faucet = new ethers.Contract(faucetInfo.contractAddress, abi, signer);

    let lastConfirmation = -1;

    let result = await callRequest(faucet, recipient.name, recipient.country, recipient.amount, confirmations, (receipt) => {
      console.log("Reciept Callback", receipt);

      if (receipt.status && receipt.blockNumber) {

        if (lastConfirmation < receipt.confirmations) {
          logText = logText + "\n" + getTime() + ":**************************************";
          logText = logText + "\n" + getTime() + ": Confirmation: '" + receipt.confirmations + "'";
          logText = logText + "\n" + getTime() + ": Status: '" + receipt.status + "'";
          logText = logText + "\n" + getTime() + ": Time: '" + receipt.time + "'";
          logText = logText + "\n" + getTime() + ": TX Hash: '" + receipt.transactionHash + "'";
          logText = logText + "\n" + getTime() + ": Block: '" + receipt.blockNumber + "'";
          logText = logText + "\n" + getTime() + ": Gas used: '" + receipt.gasUsed + "'";
          logText = logText + "\n" + getTime() + ":**************************************";

          receipt.balance = ethers.utils.formatEther(recipient.balance);

          transactionReciepts.push(receipt);
          transactionReciepts = transactionReciepts;
        }

        lastConfirmation = receipt.confirmations;

      } else if (receipt.hash && receipt.confirmations === 0) {
        receipt.transactionHash = receipt.hash;
        transactionReciepts.push(receipt);
        transactionReciepts = transactionReciepts;
      }
    });

    /*
    let result = {
      status: false,
      transactionHash: ethers.constants.HashZero,
      blockNumber: 0,
      blockHash: ethers.constants.HashZero,
      gasUsed: 0,
      confirmations: 0,
      logsBloom: ethers.constants.HashZero,
      sdnTX: ethers.constants.HashZero,
      error: null,
      errorReason: "",
      errorCode: "",
    }; */

    logText = logText + "\n" + getTime() + ": Reuesting donation finished in " + (new Date() - startSeconds)/1000 + " seconds";
    logText = logText + "\n" + getTime() + ": Status: '" + result.status + "'";
    logText = logText + "\n" + getTime() + ": Time: '" + result.time + "'";
    logText = logText + "\n" + getTime() + ": TX Hash: '" + result.transactionHash + "'";
    logText = logText + "\n" + getTime() + ": Block: '" + result.blockNumber + "'";
    logText = logText + "\n" + getTime() + ": Confirmations: '" + result.confirmations + "'";
    logText = logText + "\n" + getTime() + ": Gas used: '" + result.gasUsed + "'";
  } catch (e) {
    // TODO REMOVE
    console.log("catch in requestDonation", e);
    let reason = errorToReason(e);
    let code = errorToCode(e);
    logText = logText + "\n" + getTime() + ": Code: '" + code + "'";
    logText = logText + "\n" + getTime() + ": Reason: '" + reason + "'";
    logText = logText + "\n" + getTime() + ": Requesting donation FAILED in " + (new Date() - startSeconds)/1000 + " seconds";
    // CALL_EXCEPTION is an contract error
    createErrorPopup(code, reason, (code === "INVALID_SIGNER" || code === "CALL_EXCEPTION" || code === "INSUFFICIENT_FUNDS" ? false : true));
  }
  spin_css_tx = "";
}

/**
* requestDonationWithWebWallet() builds the url to call the contract function requestDonation(address payable _to, uint _amount, string memory _name, bytes8 _country) public with the webwallet
*
* @param none
* @return nothing
*/
async function requestDonationWithWebWallet() {
  const startSeconds = new Date();
  try {

    logText = logText + "\n" + getTime() + ": Requesting with Webwallet " + recipient.amount + " NRG to '" + recipient.address + "'' from '" + recipient.name + "' in country '" + recipient.country + "'";
    transactionReciepts = [];
    spin_css_tx = "nf-hc-spin";

    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    //const signer = provider.getSigner(recipient.address);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);
    // Use a read-only signer as we call only view/pure functions here
    const signer = new ethers.VoidSigner(recipient.address, provider);

    // The Contract object
    const faucet = new ethers.Contract(faucetInfo.contractAddress, abi, signer);

    let lastConfirmation = -1;

    let result = await buildCallrequestDonationArguments(faucet, recipient.address, recipient.name, recipient.country, recipient.amount);

    if (result.status === 1) {
      window.open(result.url,'_blank');
    }

    /*
    let result = {
      status: 0,
      url: "",
      error: {code: "JAVASCRIPT_911", reason: "callStatic returned null", message: "🥶 CALL 911 or shout U RTardApe",}
    }; */

    logText = logText + "\n" + getTime() + ": Building unsigned donation transaction finished in " + (new Date() - startSeconds)/1000 + " seconds";

  } catch (e) {
    // TODO REMOVE
    console.log("catch in sendDonationWithWebWallet", e);
    let reason = errorToReason(e);
    let code = errorToCode(e);
    logText = logText + "\n" + getTime() + ": Code: '" + code + "'";
    logText = logText + "\n" + getTime() + ": Reason: '" + reason + "'";
    logText = logText + "\n" + getTime() + ": Building unsigned request transaction FAILED in " + (new Date() - startSeconds)/1000 + " seconds";
    // CALL_EXCEPTION is an contract error, INVALID_SIGNER means account is not unlocked
    createErrorPopup(code, reason, (code === "INVALID_SIGNER" || code === "CALL_EXCEPTION" ? false : true));
  }
  spin_css_tx = "";
}


/**
* checkTransaction(txhash) waits and shows confirmation for transaction hash
*
* @param txhash the hash to check
* @return nothing
*/
async function checkTransaction(txhash) {
  const startSeconds = new Date();
  try {
    logText = logText + "\n" + getTime() + ": Waiting for transaction '" + txhash;
    transactionReciepts = [];
    spin_css_tx = "nf-hc-spin";

    // Inform GUI that we are waiting for a transaction, shows the spinning hourglass immediat, first block has txhash only set, blocknumber is not known yet
    /* <ListItem  header="{spin_css_tx ? 'Waiting for transaction to be mined ...' : 'Transaction finished'}" title="TX: {transactionReciepts[0].transactionHash}"
                  footer="{spin_css_tx ? `In progress, please wait for ${confirmations} confirmations`:''}" after="&nbsp;"> */
        // Add to display queue
    let txDisplayItem = {transactionHash: txhash, confirmations: confirmations, blockNumber: -1, time: getTime(), balance: 0, gasUsed: 'N/A',};
    transactionReciepts.push(txDisplayItem);
    transactionReciepts = transactionReciepts;


    const provider = new ethers.providers.JsonRpcProvider($web3URL);
    const txResponse = await provider.getTransaction(txhash).catch((e) => { e.step = "getTransaction"; throw (e); });
    console.log("txResponse", txResponse);

    let lastConfirmation = -1;

    // Callback -------------------------------------------------------------------------------------
    const receiptHandler = (receipt) => {
      console.log("Transaction Receipt", receipt);
      if (receipt.status && receipt.blockNumber) {

        if (lastConfirmation < receipt.confirmations) {
          logText = logText + "\n" + getTime() + ":**************************************";
          logText = logText + "\n" + getTime() + ": Confirmation: '" + receipt.confirmations + "'";
          logText = logText + "\n" + getTime() + ": Status: '" + receipt.status + "'";
          logText = logText + "\n" + getTime() + ": TX Hash: '" + receipt.transactionHash + "'";
          logText = logText + "\n" + getTime() + ": Block: '" + receipt.blockNumber + "'";
          logText = logText + "\n" + getTime() + ": Gas used: '" + receipt.gasUsed + "'";
          logText = logText + "\n" + getTime() + ":**************************************";

          if (!receipt.time) {
            receipt.time = getTime();
          }
          if (!receipt.value) {
            receipt.value = ethers.utils.formatEther(txResponse.value);
          }

          if (receipt.confirmations === 1) {
            transactionReciepts.shift();                // remove first element
                                                        // essentially updating the first block
          }

          transactionReciepts.push(receipt);
          transactionReciepts = transactionReciepts;
        }

        lastConfirmation = receipt.confirmations;

      } else if (receipt.confirmations === 0) {

        if (receipt.blockNumber === null) {
          receipt.blockNumber = -1;
        }

        transactionReciepts.push(receipt);
        transactionReciepts = transactionReciepts;
      }
    };
    // END Callback ---------------------------------------------------------------------------------

    // Did we get information about the transaction? If external node this can be null!
    if (txResponse) {

      // TODO REMOVE
      console.log("txResponse.blockNumber:", txResponse.blockNumber);
      console.log("txResponse.timestamp:", txResponse.timestamp);

      // Already mined ?
      if (txResponse.blockNumber) {

        if (!txResponse.timestamp) {
          console.log("Getting block:", txResponse.blockNumber);
          const block = await provider.getBlock(txResponse.blockNumber);
          txResponse.timestamp = block.timestamp;
          console.log("txResponse.timestamp:", txResponse.timestamp);
        }

        logText = logText + "\n" + getTime() + ":**************************************";
        logText = logText + "\n" + getTime() + ": Transaction mined in block: '" + txResponse.blockNumber + "'";
        logText = logText + "\n" + getTime() + ":**************************************";

/*        transactionHash
          <ListItem  header="{tr.confirmations} Confirmation(s) for block {tr.blockNumber}" title="Timestamp: {tr.time}, Balance: {tr.balance}" footer="Gas used: {tr.gasUsed}">
          </ListItem>
 */
        // Add to display queue
        let txDisplayItem = {transactionHash: txhash, confirmations: txResponse.confirmations, blockNumber: txResponse.blockNumber, time: getTime(txResponse.timestamp*1000), value: ethers.utils.formatEther(txResponse.value), gasUsed: 'N/A',};
        transactionReciepts.shift();                // remove first element
        transactionReciepts.push(txDisplayItem);    // essentially updating the first block
        transactionReciepts = transactionReciepts;

        // TODO REMOVE
        console.log("txDisplayItem added:", txDisplayItem);

        lastConfirmation = txResponse.confirmations;
      }

      // Wait for more confirmations?
      if (confirmations > lastConfirmation) {
        console.log("Waiting for:", txResponse.hash);
        provider.on(txResponse.hash, receiptHandler);

        // Wait for number of confirmations
        //txResponse.confirmations = 0;
        const txReceipt = await txResponse.wait(confirmations).catch((e) => { e.step = "txResponse.wait"; throw (e); });

        provider.removeListener(txResponse.hash, receiptHandler);
      }

      logText = logText + "\n" + getTime() + ": Transaction finished in " + (new Date() - startSeconds)/1000 + " seconds";
      /* logText = logText + "\n" + getTime() + ": Status: '" + txResponse.status + "'";
      logText = logText + "\n" + getTime() + ": TX Hash: '" + txResponse.transactionHash + "'"; */
      logText = logText + "\n" + getTime() + ": Block: '" + txResponse.blockNumber + "'";
      /* logText = logText + "\n" + getTime() + ": Confirmations: '" + txResponse.confirmations + "'";
      logText = logText + "\n" + getTime() + ": Gas used: '" + txResponse.gasUsed + "'"; */
    } else {
      createErrorPopup("INVALID_TRANSACTION", "TX '" + txhash + "' seems to be invalid, please try again in a minute", false);
    }

  } catch (e) {
    // TODO REMOVE
    console.log("catch in checkDonationTransaction", e);
    let reason = errorToReason(e);
    let code = errorToCode(e);
    logText = logText + "\n" + getTime() + ": Code: '" + code + "'";
    logText = logText + "\n" + getTime() + ": Reason: '" + reason + "'";
    // CALL_EXCEPTION is an contract error
    createErrorPopup(code, reason, (code === "INVALID_SIGNER" || code === "CALL_EXCEPTION" ? false : true));
  }
  spin_css_tx = "";
}

/**
* sendDonation() calls the contract function Donation(string memory _name, bytes8 _country) public payable
*
* @param none
* @return nothing
*/
async function sendDonation() {
  const startSeconds = new Date();
  try {

    logText = logText + "\n" + getTime() + ": Sending " + donor.amount + " NRG to faucet from '" + donor.name + "' in '" + donor.country + "' at address '" + donor.address + "'";
    transactionReciepts = [];
    spin_css_tx = "nf-hc-spin";

    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    const signer = provider.getSigner(donor.address);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

    // The Contract object
    const faucet = new ethers.Contract(faucetInfo.contractAddress, abi, signer);

    let lastConfirmation = -1;

    let result = await callDonation(faucet, donor.name, donor.country, donor.amount, confirmations, (receipt) => {
      console.log("Reciept Callback", receipt);

      if (receipt.status && receipt.blockNumber) {

        if (lastConfirmation < receipt.confirmations) {
          logText = logText + "\n" + getTime() + ":**************************************";
          logText = logText + "\n" + getTime() + ": Confirmation: '" + receipt.confirmations + "'";
          logText = logText + "\n" + getTime() + ": Time: '" + receipt.time + "'";
          logText = logText + "\n" + getTime() + ": Status: '" + receipt.status + "'";
          logText = logText + "\n" + getTime() + ": TX Hash: '" + receipt.transactionHash + "'";
          logText = logText + "\n" + getTime() + ": Block: '" + receipt.blockNumber + "'";
          logText = logText + "\n" + getTime() + ": Gas used: '" + receipt.gasUsed + "'";
          logText = logText + "\n" + getTime() + ":**************************************";

          receipt.balance = ethers.utils.formatEther(donor.balance);

          transactionReciepts.push(receipt);
          transactionReciepts = transactionReciepts;
        }

        lastConfirmation = receipt.confirmations;

      } else if (receipt.hash && receipt.confirmations === 0) {
        receipt.transactionHash = receipt.hash;
        transactionReciepts.push(receipt);
        transactionReciepts = transactionReciepts;
      }
    });

    /*
    let result = {
      status: false,
      transactionHash: ethers.constants.HashZero,
      blockNumber: 0,
      blockHash: ethers.constants.HashZero,
      gasUsed: 0,
      confirmations: 0,
      logsBloom: ethers.constants.HashZero,
      sdnTX: ethers.constants.HashZero,
      error: null,
      errorReason: "",
      errorCode: "",
    }; */

    logText = logText + "\n" + getTime() + ": Sending donation finished in " + (new Date() - startSeconds)/1000 + " seconds";
    logText = logText + "\n" + getTime() + ": Status: '" + result.status + "'";
    logText = logText + "\n" + getTime() + ": TX Hash: '" + result.transactionHash + "'";
    logText = logText + "\n" + getTime() + ": Block: '" + result.blockNumber + "'";
    logText = logText + "\n" + getTime() + ": Confirmations: '" + result.confirmations + "'";
    logText = logText + "\n" + getTime() + ": Gas used: '" + result.gasUsed + "'";
  } catch (e) {
    // TODO REMOVE
    console.log("catch in sendDonation", e);
    let reason = errorToReason(e);
    let code = errorToCode(e);
    logText = logText + "\n" + getTime() + ": Code: '" + code + "'";
    logText = logText + "\n" + getTime() + ": Reason: '" + reason + "'";
    logText = logText + "\n" + getTime() + ": Sending donation FAILED in " + (new Date() - startSeconds)/1000 + " seconds";
    // CALL_EXCEPTION is an contract error
    createErrorPopup(code, reason, (code === "INVALID_SIGNER" || code === "CALL_EXCEPTION" ? false : true));
  }
  spin_css_tx = "";
}

/**
* sendDonationWithWebWallet() builds the url to call the contract function Donation(string memory _name, bytes8 _country) public payable from the webwallet
*
* @param none
* @return nothing
*/
async function sendDonationWithWebWallet() {
  const startSeconds = new Date();
  try {

    logText = logText + "\n" + getTime() + ": Sending with Webwallet " + donor.amount + " NRG to '" + faucetInfo.contractAddress + "'' from '" + donor.name + "' in country '" + donor.country + "'";
    transactionReciepts = [];
    spin_css_tx = "nf-hc-spin";

    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    //const signer = provider.getSigner(donor.address);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);
    // Use a read-only signer as we call only view/pure functions
    //const signer = new ethers.VoidSigner(fromAddress, provider);

    // No signer needed here

    // The Contract object
    const faucet = new ethers.Contract(faucetInfo.contractAddress, abi, provider);

    let lastConfirmation = -1;

    let result = await buildCallDonationArguments(faucet, donor.name, donor.country, donor.amount);

    if (result.status === 1) {
      window.open(result.url,'_blank');
    }

    /*
    let result = {
      status: 0,
      url: "",
      error: {code: "JAVASCRIPT_911", reason: "callStatic returned null", message: "🥶 CALL 911 or shout U RTardApe",}
    }; */

    logText = logText + "\n" + getTime() + ": Building unsigned donation transaction finished in " + (new Date() - startSeconds)/1000 + " seconds";

  } catch (e) {
    // TODO REMOVE
    console.log("catch in sendDonationWithWebWallet", e);
    let reason = errorToReason(e);
    let code = errorToCode(e);
    logText = logText + "\n" + getTime() + ": Code: '" + code + "'";
    logText = logText + "\n" + getTime() + ": Reason: '" + reason + "'";
    logText = logText + "\n" + getTime() + ": Building unsigned donation transaction FAILED in " + (new Date() - startSeconds)/1000 + " seconds";
    // CALL_EXCEPTION is an contract error, INVALID_SIGNER means account is not unlocked
    createErrorPopup(code, reason, (code === "INVALID_SIGNER" || code === "CALL_EXCEPTION" ? false : true));
  }
  spin_css_tx = "";
}

/**
* getRecipientBalance() gets the balance and name for recipient.address
*
* @param none
* @return nothing
*/
const getRecipientBalance = async () => {
  try {
    recipient.isAddressValid = false;
    if (recipient.address && recipient.address !== ethers.constants.AddressZero && ethers.utils.isAddress(recipient.address)) {
      recipient.isAddressValid = true;
      const provider = new ethers.providers.JsonRpcProvider($web3URL);
      // The Contract object
      const faucet = new ethers.Contract($faucetAddress, abi, provider);

      let b = provider.getBalance(recipient.address);

      if (!recipient.nameOverwritten) {
        let n = faucet.getRecipientName(recipient.address);
        recipient.name = await n;
      }

      if (recipient.name.length > 2) {
        requestButtonEnabled = true;
      }

      recipient.balance = await b;

      if (!userSelfSelected && recipient.balance && recipient.balance.gt( ethers.utils.parseEther(dividerInNRG)) ) {
        userIsDonor = true;
      }
    }
    return recipient.balance;

  } catch (e) {
    console.log("getRecipientBalance", e);
    console.log("e.code", e.code);
    console.log("e.message", e.message);
    console.log("e.reason", e.reason);

    if (!e.step) {
      e.step = "getBalance";
    }
    throw(e);
  }
}

/**
* getDonorBalance() gets the balance and name for donor.address
*
* @param none
* @return nothing
*/
const getDonorBalance = async () => {
  try {
    donor.isAddressValid = false;
    if (donor.address && donor.address !== ethers.constants.AddressZero && ethers.utils.isAddress(donor.address)) {
      donor.isAddressValid = true;
      const provider = new ethers.providers.JsonRpcProvider($web3URL);
      // The Contract object
      const faucet = new ethers.Contract($faucetAddress, abi, provider);

      let b = provider.getBalance(donor.address);
      if (!donor.nameOverwritten) {
        let n = faucet.getDonorName(donor.address);
        donor.name = await n.catch((e) => { e.step = "getDonorName"; throw (e); });
      }

      donor.balance = await b.catch((e) => { e.step = "getBalance"; throw (e); });
      sendButtonEnabled = true;

      if (!userSelfSelected && donor.balance && donor.balance.lt( ethers.utils.parseEther(dividerInNRG)) ) {
        userIsDonor = false;
      }
    }
    return donor.balance;

  } catch (e) {
    console.log("getDonorBalance", e);
    console.log("e.code", e.code);
    console.log("e.message", e.message);
    console.log("e.reason", e.reason);

    if (!e.step) {
      e.step = "getBalance";
    }
    throw(e);
  }
}

/**
* getFaucetDetails() gets info about the contract
*
* @param faucetAddress address of the contract
* @return is a structure named faucetDetails
          let faucetDetails = {
            version: 0,
            balance: ethers.constants.Zero,
            calculatedBalance: ethers.constants.Zero,
            donorsCount: 0,
            recipientsCount: 0,
            owner: ethers.constants.AddressZero,
            contractAddress: ethers.constants.AddressZero,
            date: new Date(),
            error: 'none yet',
            lastRecipient: {name:"No one", address: ethers.constants.AddressZero, count: 0, amount: ethers.constants.Zero, country: "🗺", date: new Date(),},
          };
*
*/
const getFaucetDetails = async (faucetAddress) => {

  let fetchStep = "Setup Connection";
  let result = faucetInfo;
  try {

      const provider = new ethers.providers.JsonRpcProvider($web3URL);

      // The provider also allows signing transactions to
      // send ether and pay to change state within the blockchain.
      // For this, we need the account signer...
      // const signer = provider.getSigner(gasDonorAddress);
      // const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);
      // Use a read-only signer for this function
      //const signer = new ethers.VoidSigner(gasDonorAddress, provider);

      // No signer needed here

      // The Contract object
      const faucet = new ethers.Contract(faucetAddress, abi, provider);


      fetchStep = "Fetch Details";
      console.log("getFaucetDetails From FaucetAdr", faucetAddress);
      //let info = await fetchFaucetDetails($web3URL, faucetAddress, abi);
      let info = await fetchFaucetDetailsAsync(faucet);
      console.log("getFaucetDetails result info", info);

      fetchStep = "GetLastRecipient";
      // Get last Recipient Payment
      // function getRecipientByIndex(uint _i) public view returns (address _recipient, uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country)
      let index = info.recipientsCount-1;
      let r = await faucet.getRecipientByIndex(index).catch((e) => { console.log("getRecipientByIndex", index, "failed:", e); e.step = "faucet.getRecipientByIndex(" + index + ")"; throw (e);});
      if (r) {
        info.lastRecipient.name = r._name
        info.lastRecipient.address = r._recipient;
        info.lastRecipient.count = r._numPayments;
        info.lastRecipient.amount = ethers.utils.formatEther( r._totalBalance );

        console.log("getRecipientByIndex", index, "info", info);

        if (r._country === "0x5465727261000000") {    // Terra
          info.lastRecipient.country = "🌎";
        } else if (r._country === "0x0000000000000000") {    // Zero-bytes
          info.lastRecipient.country = "🌍";
        } else {
          info.lastRecipient.country = ethers.utils.toUtf8String(r._country);
        }
        fetchStep = "GetPaymentDetails";
        let pd = await faucet.getPaymentDetails(info.lastRecipient.address, info.lastRecipient.count).catch((e) => { console.log("getPaymentDetails", info.lastRecipient.address, "payment", info.lastRecipient.count, "failed:", e); e.step = "faucet.getPaymentDetails(" + info.lastRecipient.address + ", " + info.lastRecipient.count + ")"; throw (e);});
        if (pd) {
          console.log("DETAILS recipient: ", pd);
          info.lastRecipient.date = pd._timestamp*1000;
        }
      }

      info.date = new Date();
      info.error = { reason: "Success", code: "OK", step: "all done", };
      result = info;
  }
  catch (e) {
    console.log("getFaucetDetails CATCH:", e);
    if (!e.step) {
      e.step = fetchStep;
    }
    throw(e);
  }
  spin_css = "";
  return result;
};

let faucetIVId;
let balanceIVId;

const keypressEvent = (e) => {
  if (e.code === "Enter") {
    if (e.target.id === "inputIDmanualRecipientAccount" || e.target.id === "inputIDRecipientNameWW") {
      //console.log("keypress", e);
      if (requestButtonEnabled) {
        requestDonationWithWebWallet();
      }
    } else if (e.target.id === "inputIDDonorNameWW") {
      if (sendButtonEnabled) {
        sendDonationWithWebWallet();
      }
    } else if (e.target.id === "inputIDDonorName") {
      if (sendButtonEnabled) {
        sendDonation();
      }
    } else if (e.target.id === "inputIDRecipientName") {
      if (requestButtonEnabled) {
        requestDonation();
      }
    } else if (e.target.id === "inputIDcheckTXRecipient") {
      if (recipientTxCheckButtonEnabled) {
        checkTransaction(recipient.txhash);
      }
    } else if (e.target.id === "inputIDcheckTXDonor") {
      if (donorTxCheckButtonEnabled) {
        checkTransaction(donor.txhash);
      }
    }
  }
}

const addKeypressEnterEvents = () => {
  let inputAddress = document.getElementById("inputIDmanualRecipientAccount");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDcheckTXRecipient");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDcheckTXDonor");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDDonorNameWW");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDRecipientNameWW");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDDonorName");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }
  inputAddress = document.getElementById("inputIDRecipientName");
  if (inputAddress) {
    inputAddress.removeEventListener("keypress", keypressEvent);
    inputAddress.addEventListener("keypress", keypressEvent);
    //console.log("Added keypress for inputIDmanualRecipientAccount");
  }

}

function onShow() {
  console.log("OnShow smartcontract");
  visitedPages.update(() => {
    visitedPages.smartcontract = true;
    visitedPages.web3 = false;
    visitedPages.info = false;
    return visitedPages;
  });

  faucetInfo.contractAddress = $faucetAddress;
  getCountryFlag().catch((e) => {createErrorPopup("An error occured", e)}).then((f) => {donor.country = f; recipient.country = f});
  recipient.address = $ethAccount.address;
  donor.address = $ethAccount.address;

  addKeypressEnterEvents();

  // Initial load after 0.1 seconds for the user to notice it starts to load
  setTimeout( () => {
    getFaucetDetails($faucetAddress).catch(
      (e) => {
        if (e.step) {
          createErrorPopup("Faucet details", "At " + e.step + ": " + e);
        } else {
          createErrorPopup("Faucet details", e);
        } }).then(
      (i) => {
        if (i) {
          faucetInfo = i;
        }
      })
  }, 100);

  // Refresh faucet details every 60 seconds
  faucetIVId = setInterval( () => {
    getFaucetDetails($faucetAddress).catch(
      (e) => {
        if (e.step) {
          createErrorPopup("Faucet details", "At " + e.step + ": " + e);
        } else {
          createErrorPopup("Faucet details", e);
        } }).then(
      (i) => {
        if (i) {
          faucetInfo = i;
        }
      })
  }, 60000);

  // Refresh recipient and donor balance every 5 seconds
  balanceIVId = setInterval( () => {

    addKeypressEnterEvents();

    getRecipientBalance().catch((e) => {createErrorPopup("Recipient balance", e);});
    getDonorBalance().catch((e) => {createErrorPopup("Donor balance", e);});

  }, 5000);

}

function onHide() {
  console.log("OnHide smartcontract");
  visitedPages.update(() => {
    visitedPages.smartcontract = false;
    return visitedPages;
  });

  if (faucetIVId) {
    clearInterval(faucetIVId);
  }
  faucetIVId = null;

  if (balanceIVId) {
    clearInterval(balanceIVId);
  }
  balanceIVId = null;
}

let popup;
function createErrorPopup(title, error, showGit) {
  try {
    // Create popup
    if (popup) {
      console.log("createErrorPopup", popup);
      popup.close();
      f7.popup.close(popup);
      f7.popup.destroy(popup);
      //popup = null;
    }
    let git = "";
    if (showGit) {
      git = '<a class="link external" href="https://github.com/kimxilxyong/Energi-Smart-Contract-Tutorial/issues" target="_blank" external="true">Please open an issue on github</a>';
    }

    popup = f7.popup.create({
      content: `
        <div class="popup" id="ErrorPopup">
          <div class="page">
            <div class="navbar">
            <div class="navbar-bg" style="background-color: red"></div>
              <div class="navbar-inner">
                <div class="title" style="font-size:1em;">${title}</div>
                <div class="right"><a href="#" class="link popup-close">X</a></div>
              </div>
            </div>
            <div class="page-content">
              <div class="block">
                <p class="mono-small">${error}<br>
                ${git}
                </p>
              </div>
            </div>
          </div>
        </div>
      `.trim(),
    });

    // Open it
    popup.open();
  } catch(e) {
    console.log(e);
  }
}
</script>
