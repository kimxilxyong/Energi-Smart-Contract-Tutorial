<!-- SPDX-License-Identifier: MIT -->
<!-- Copyright (c) 2020 Kim Il Yong -->
<!-- Version 1.0.0 -->

<Page name="home" onPageBeforeIn={onShow} onPageBeforeOut={onHide}>
  <!-- Top Navbar -->
  <Navbar sliding={true}>
    <NavLeft>
      <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
    </NavLeft>
    <NavTitle sliding>Testnet Faucet</NavTitle>
    <NavRight>
      <!-- <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="right" /> -->
      <Link panelOpen="right" class="f7-icons">square_righthalf_fill</Link>
    </NavRight>
  </Navbar>
  <!-- Toolbar -->
<!--   <Toolbar bottom>
    <Link panelOpen="right" class="f7-icons">square_righthalf_fill</Link>
    <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="right" />
  </Toolbar> -->
  <!-- Page content -->

  <Card outline>
    <CardHeader style="color:var(--energi-color-green);">Energi Testnet faucet Smart-Contract details</CardHeader>
    <CardContent>
      <List>
        <ListItem style="color:var(--energi-color-turkis);" title="Version: {c.version}"></ListItem>
        <ListItem style="color:var(--energi-color-turkis);" title="Balance: { ethers.utils.formatEther(c.balance) } NRG"></ListItem>
        <ListItem style="color:var(--energi-color-turkis);" title="Calculated Balance: { ethers.utils.formatEther(c.calculatedBalance) } NRG"></ListItem>

        <!-- after={donorCountAfter} man &#xf263;  people &#xf007;-->
        {#if c.donorsCount > 0}
          <ListItem onClick="{(e) => {onDonorsClicked(e);}}" title="Donors Count: {c.donorsCount}"  badge={c.donorsCount} footer="Click to load donors">
            <i slot="media" class="nf-icons nf-hc-fw {d_nf_hc_spin} colorlime">&#xf263;</i>
            <span slot="after-start" style="margin-right: 10px;">
                  <Stepper name="Nr of TXs" value={numberOfDonorsToFetch}
                          onStepperChange={(v) => {numberOfDonorsToFetch = v}}
                          small wraps autorepeat autorepeatDynamic inputReadonly autocomplete="off"/>
            </span>
          </ListItem>
        {:else}
          <ListItem title="Donors Count: {c.donorsCount}"  badge={c.donorsCount} footer="No donors found 🙄">
            <i slot="media" class="nf-icons nf-hc-fw colorgrey">&#xf263;</i>
          </ListItem>
        {/if}
        <li>
          <ul>
            {#if donorsArray.length > 0}
              <ListItem title="Donors list: {donorsArray.length}" footer="Click to refresh" onClick="{(e) => {onDonorsClicked(e);}}"></ListItem>
              {#each donorsArray as donor, i}
                <ListItem title="{donor.name}" footer="{donor.address}" header="Donated: { ethers.utils.formatEther(donor.balance) } NRG { timeAgoDisplayItems[donor.bindIndex] ? timeAgoDisplayItems[donor.bindIndex].current : ''}" badge="{donor.count}">
                  <!-- <i slot="media" class="nf-icons nf-hc-fw colororange">&#xf007;</i> -->
                  <i slot="media" class="nf-icons nf-hc-fw nf-hc-1x">{donor.country}</i>
                  <!-- <i slot="after-end" style="margin-right:20px;">&nbsp;</i> -->
                </ListItem>
              {/each}
            {/if}
          </ul>
        </li>
        {#if donorsArray.length > 0}
          <li><ul><div class="seperator"/></ul></li>
        {/if}

        {#if c.recipientsCount > 0}
          <ListItem onClick="{(e) => {onRecipientsClicked(e);}}" title="Recipients Count: {c.recipientsCount}" badge={c.recipientsCount} footer="Click to load recipients">
            <i slot="media" class="nf-icons nf-hc-fw {r_nf_hc_spin} colorlime">&#xf263;</i>
            <span slot="after-start" style="margin-right: 10px;">
                  <Stepper name="Nr of TXs" value={numberOfRecipientsToFetch}
                          onStepperChange={(v) => {numberOfRecipientsToFetch = v}}
                          small wraps autorepeat autorepeatDynamic inputReadonly autocomplete="off"/>
            </span>
          </ListItem>
        {:else}
          <ListItem title="Recipients Count: {c.recipientsCount}"  badge={c.recipientsCount} footer="No recipients found 🙄">
            <i slot="media" class="nf-icons nf-hc-fw colorgrey">&#xf263;</i>
          </ListItem>
        {/if}
        <li>
          <ul>
            {#if recipientsArray.length > 0}
              <ListItem title="Recipients list: {recipientsArray.length}" footer="Click to refresh" onClick="{(e) => {onRecipientsClicked(e);}}"></ListItem>
              {#each recipientsArray as recipient, i}
                <ListItem title="{recipient.name}" footer="{recipient.address}" header="Recieved: { ethers.utils.formatEther(recipient.balance) } NRG { timeAgoDisplayItems[recipient.bindIndex] ? timeAgoDisplayItems[recipient.bindIndex].current : ''}" badge="{recipient.count}">
                  <!-- <i slot="media" class="nf-icons nf-hc-fw colororange">&#xf007;</i> -->
                  <i slot="media" class="nf-icons nf-hc-fw nf-hc-1x">{recipient.country}</i>
                  <!-- <i slot="after-end" style="margin-right:20px;">&nbsp;</i> -->
                </ListItem>
              {/each}
            {/if}
          </ul>
        </li>
        {#if recipientsArray.length > 0}
          <ListItem class="seperator"></ListItem>
        {/if}


        <!-- OK-Square U+F634  clipboard U+F429 -->
        <ListItem title="Contract Address: {c.contractAddress}" after="&nbsp;">
            <div slot="after-title" class="colortext" style="margin-left: 20px;">
              {#if c.contractAddress !== ethers.constants.AddressZero}
                <span class="bround bgreen fss3 {isClipboardDoneClass2}">
                Copied!
                </span>
                <i on:click={ () => {
                                        isClipboardDoneClass2 = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(c.contractAddress);
                                        setTimeout(() =>   {
                                                                if (isClipboardDoneClass2 !== "hidden") {
                                                                    isClipboardDoneClass2 = "hidden";
                                                                }
                                                            }, 3300);
                                    }
                            } title="Copy to clipboard" class="nf-icons nf-hc-1x colorgrey">&#xF429;
                </i>
              {/if}
            </div>
        </ListItem>
        <ListItem title="Owner: {c.owner}" after="&nbsp;">
            <div slot="after-title" class="colortext" style="margin-left: 20px;">
              {#if c.owner !== ethers.constants.AddressZero}
                <span class="bround bgreen fss3 {isClipboardDoneClass1}">
                Copied!
                </span>
                <i on:click={ () => {
                                        isClipboardDoneClass1 = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(c.owner);
                                        setTimeout(() =>   {
                                                                if (isClipboardDoneClass1 !== "hidden") {
                                                                    isClipboardDoneClass1 = "hidden";
                                                                }
                                                            }, 3300);
                                    }
                            } title="Copy to clipboard" class="nf-icons nf-hc-1x colorgrey">&#xF429;
                </i>
              {/if}
            </div>
        </ListItem>
        <ListItem title="Result: {status.error}" class="{status.color}"></ListItem>
        <ListItem title="Activity: { timeAgoDisplayItems[0] ? timeAgoDisplayItems[0].current : 'refresh not activated'}"></ListItem>
      </List>
    </CardContent>
  </Card>

  <Card outline>
    <CardHeader style="color:var(--energi-color-green);">Logs, stats and more</CardHeader>
    <CardContent>
        <List noHairlinesMd>
            <ListItem style="color:var(--energi-color-turkis);" title="JSON-RPC provider node: {$web3URL}"></ListItem>
            <ListItem title="Log size: {countMaxLogLength} chars maximum">
                <span slot="after">
                    <Stepper value={countMaxLogLength} max=10000 name=""
                        onStepperChange={(v) => {countMaxLogLength = v}} small wraps autorepeat autorepeatDynamic inputReadonly autocomplete="off"/>
                </span>
            </ListItem>
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
            <ListItem title="Source code repository" link="https://github.com/kimxilxyong/Energi-Smart-Contract-Tutorial" target="_blank" external>
                <i slot="media" class="f7-icons">logo_github</i>
            </ListItem>
            <ListItem title="Smartcontract solidity source code (faucet.sol)" link="https://github.com/kimxilxyong/Energi-Smart-Contract-Tutorial/blob/master/solidity/faucet.sol" target="_blank" external>
                <i slot="media" class="f7-icons colorblue">logo_github</i>
            </ListItem>
            <ListItem class="noline" title="Test your font setup" link="static/fonttest2.html" target="_blank" external>
                <img slot="media" src={browserIcon} alt="?" width="32" height="32">
            </ListItem>
            <ListItem class="seperator"></ListItem>
            <ListItem title="Powerd by">
                <div slot="inner-end" class="flex-row">
                    <!-- style='height: 2em; width: 6em;' -->
                    <div class="flex-col dynamic dinherit tooltip-init" style="min-width:32px;" data-tooltip="Svelte.io">
                      <Link href="https://svelte.dev/" external target="_blank">
                        <img class="mhauto" src="./images/svelte_logo.svg" alt="Svelte" width="32" height="32">
                      </Link>
                    </div>
                    <div class="flex-col dynamic dinherit tooltip-init" style="min-width:32px;" data-tooltip="Framework7.io">
                      <Link href="https://framework7.io/svelte/" external target="_blank" style="margin-left:15%">
                        <img class="mhauto" src="./images/framework7-24x24.png" alt="Framework7" width="24" height="24">
                      </Link>
                    </div>
                    <div class="flex-col dynamic dinherit tooltip-init" style="min-width:32px;" data-tooltip="Ethers.js">
                      <Link href="https://docs.ethers.io/" external target="_blank">
                        <img class="mhauto" src="./images/ethers_logo.svg" alt="Ethers.js" width="32" height="32">
                        <!-- style="background-color: cornflowerblue;" -->
                      </Link>
                    </div>
                    <div class="flex-col dynamic dinherit tooltip-init" style="min-width:38px;" data-tooltip="Go the Language">
                      <Link href="https://go.dev/" external target="_blank">
                        <img class="mhauto" src="./images/Go-Logo_Aqua.svg" alt="Go Lang" width="32" height="32">
                      </Link>
                    </div>
                    <div class="flex-col dynamic dinherit tooltip-init" style="min-width:68px;" data-tooltip="GIT SourceCode Management"> <!-- GIT -->
                      <Link href="https://git-scm.com/" external target="_blank">
                        <img class="mhauto" src="./images/Git-logo.svg" alt="GIT" width="64" height="32">
                      </Link>
                    </div>
                    <div class="tooltip-init flex-col dynamic dinherit" style="min-width:42px;" data-tooltip="Linux Institute (rip)"> <!-- LINUX -->
                      <Link style="align-self: baseline; flex: auto;" href="https://web.archive.org/web/20051201021143/http://www.li.org/" external target="_blank">
                        <img class="mhauto" src="./images/newtux.svg" alt="Linux" width="32" height="32">
                      </Link>
                    </div>
                    <div class="flex-col dynamic dinherit nf-icons nf-hc-fw tooltip-init" style="min-width:32px;" data-tooltip="HTML5">                <!-- HTML5 -->
                      <Link class="colororange" href="https://www.w3schools.com/html/default.asp" external target="_blank">&#xf81c;</Link>
                    </div>
                    <div class="flex-col dynamic dinherit nf-icons nf-hc-fw tooltip-init" style="min-width:32px;" data-tooltip="ECMAScript 2018">    <!-- JS -->
                      <Link class="coloryellow" href="https://www.w3schools.com/js/js_versions.asp" external target="_blank">&#xf81d;</Link>
                    </div>
                </div>
            </ListItem>
            <!-- <li><div class="seperator"/></li> -->
        </List>
    </CardContent>
    <CardFooter>
      <Link href="https://www.google.com/search?q=javascript+ethereum" external target="_blank">More Information</Link>

      <Button fill raised href="/info/" view=".view-main" panelClose>
        <i class="f7-icons button-icon">arrowtriangle_right</i>
        Start
      </Button>
    </CardFooter>
  </Card>

  <!-- chain &#xf838;  OKman &#xf263;  people &#xf007; GIT f1d2 , LogBook f02d , Linux f17c , GNU e779 , gitcat &#xf09b; html5 f81c , js f81d , linux f83c , box arrow right f705 -->


<!--   <Block strong>
    <p>This is an example of split view application layout, commonly used on tablets. The main approach of such kind of layout is that you can see different views at the same time.</p>

    <p>Each view may have different layout, different navbar type (dynamic, fixed or static) or without navbar.</p>

    <p>The fun thing is that you can easily control one view from another without any line of JavaScript just using "data-view" attribute on links.</p>
  </Block>

  <BlockTitle>Navigation</BlockTitle>
  <List>
    <ListItem link="/about/" title="About"/>
    <ListItem link="/form/" title="Form"/>
  </List>

  <BlockTitle>Modals</BlockTitle>
  <Block strong>
    <Row>
      <Col width="50">
        <Button fill raised popupOpen="#my-popup">Popup</Button>
      </Col>
      <Col width="50">
        <Button fill raised loginScreenOpen="#my-login-screen">Login Screen</Button>
      </Col>
    </Row>
  </Block>

  <BlockTitle>Panels</BlockTitle>
  <Block strong>
    <Row>
      <Col width="50">
        <Button fill raised panelOpen="left">Left Panel</Button>
      </Col>
      <Col width="50">
        <Button fill raised panelOpen="right">Right Panel</Button>
      </Col>
    </Row>
  </Block>

  <List>
    <ListItem
      title="Dynamic (Component) Route"
      link="/dynamic-route/blog/45/post/125/?foo=bar#about"
    />
    <ListItem
      title="Default Route (404)"
      link="/load-something-that-doesnt-exist/"
    />
    <ListItem
      title="Request Data & Load"
      link="/request-and-load/user/123456/"
    />
  </List> -->

</Page>

<script>
import {
    Page,
    Navbar,
    NavLeft,
    NavTitle,
    NavTitleLarge,
    NavRight,
    Link,
    Toolbar,
    Block,
    BlockTitle,
    Card,
    CardHeader,
    CardContent,
    CardFooter,
    List,
    ListItem,
    ListInput,
    Row,
    Stepper,
    Col,
    Button
} from 'framework7-svelte';
import { ethers } from "ethers";
import { tick, getContext, onMount, onDestroy } from "svelte";
import {
    web3URL,
    faucetAddress,
    gasDonorAddress
} from "../js/stores.js";
//import {  } from "../js/web3utils.js";
import { detectBrowser, timeDifference, copyToClipboard, getTime, scrollTo, sleep } from "../js/utils.js";
import {abi107} from '../js/abi.js';
const abi = abi107;

let countMaxLogLength = 3000;
let animateSwingClass = "animate__animated animate__swing animate__infinite colorturkis";
let isClipboardDoneClass1 = "hidden";
let isClipboardDoneClass2 = "hidden";

let d_nf_hc_spin = "";
let r_nf_hc_spin = "";
let donorsColor = "";

let donorCountBadge = "?";
let donorCountAfter = "Loading...";

let numberOfDonorsToFetch = 5;
let numberOfRecipientsToFetch = 5;

let donorsArray = [];
let recipientsArray = [];
let timeAgoDisplayItems = [];
let faucetIVId;

let logText = "";

let browserIcon = "./images/unknown32x32.png";
let br = detectBrowser(null);
if (br && br.browser && br.browser.icon) {
  browserIcon = br.browser.icon
}

console.log("Browser Icon", browserIcon);

const result = {
    version: 0,
    balance: ethers.constants.Zero,
    calculatedBalance: ethers.constants.Zero,
    donorsCount: 0,
    recipientsCount: 0,
    owner: ethers.constants.AddressZero,
    contractAddress: ethers.constants.AddressZero,
    error: 'none yet',
};
let c = result;

let status = {error: "none yet, wait a second 🥸", color: ""};

const loadFaucetStatistics = () => {
  /* currentActivity = "Details loading ..."; */
  timeAgoDisplayItems[0].current = "Details loading ...";
  const startSeconds = new Date();
  if (logText.length == 0 || logText.length > countMaxLogLength) {
      logText = getTime() + ": Starting to fetch data from " + $web3URL;
  } else {
      logText = logText + "\n" + getTime() + ": Starting to fetch data from " + $web3URL;
  }

  getFaucetDetails($faucetAddress, $gasDonorAddress)
    .catch(
        (e) => {
                  console.log(".catch:", e);
                  if (e && e.reason) {
                    if (e.code) {
                      if (e.step) {
                        status.error = "code: " + e.code + " reason: " + e.reason + ", at: " + e.step;;
                      } else {
                        status.error = "code: " + e.code + " reason: " + e.reason;
                      }
                    } else {
                      status.error = "reason: " + e.reason;
                    }
                    status.color = 'colorred';
                  } else {
                    status.error = e;
                    status.color = 'colorred';
                  };
                  c = c;
                  if (timeAgoDisplayItems.length > 0) {
                    timeAgoDisplayItems[0].base = "Loading failed";
                    timeAgoDisplayItems[0].current = timeAgoDisplayItems[0].base;
                    timeAgoDisplayItems[0].last = new Date();
                  }
                  logText = logText + "\n" + getTime() + ": " + e;
                  logText = logText + "\n" + getTime() + ": Fechting faucet data failed in " + (new Date() - startSeconds)/1000 + " seconds"
                }
    )
    .then(
        (res) => {
                    // TODO REMOVE
                    //console.log("getFaucetDetails.then:", res);
                    if (res) {
                      c = res;
                      status.error = 'OK - faucet details loaded';
                      status.color = 'colorgreen';
                      // TODO REMOVE
                      //console.log("getFaucetDetails.then.status:", status);

                      if (timeAgoDisplayItems.length > 0) {
                        timeAgoDisplayItems[0].base = "Loading finished";
                        timeAgoDisplayItems[0].current = timeAgoDisplayItems[0].base;
                        timeAgoDisplayItems[0].last = new Date();
                      }

                      logText = logText + "\n" + getTime() + ": Fechting faucet data finished in " + (new Date() - startSeconds)/1000 + " seconds"
                    }
                  }
    );
};

const onDonorsClicked = async (e) => {
  if (d_nf_hc_spin === "nf-hc-spin" || e.target.className.includes("stepper")) {
    e.stopPropagation();
    return;
  }
  try {
    donorsArray = [];
    d_nf_hc_spin = "nf-hc-spin";
    await sleep(100);
    let donorCount = await getDonorsList($faucetAddress, $gasDonorAddress, numberOfDonorsToFetch);
    console.log("Result donorCount:", donorCount);
    if (donorCount > 0) {
      status.error = "OK - donors loaded";
      status.color = 'colorgreen';

      /* activityBaseStatus = "Loading finished: " + donorCount + " Donors"; */
      timeAgoDisplayItems[0].base = "Loading finished: " + donorCount + " Donors";
      /* currentActivity = activityBaseStatus; */
      timeAgoDisplayItems[0].current = timeAgoDisplayItems[0].base;
    } else {
      status.error = "Found not one donor!?! 🤮 This should not happen";
      status.color = 'colorred';
    }
  } catch (e) {
    if (e && e.reason) {
      if (e.code) {
        if (e.step) {
          status.error = "code: " + e.code + " reason: " + e.reason + ", at: " + e.step;
        } else {
          status.error = "code: " + e.code + " reason: " + e.reason;
        }
      } else {
        status.error = "reason: " + e.reason;
      }
      status.color = 'colorred';
    } else {
      status.error = e;
      status.color = 'colorred';
    }
  }
  d_nf_hc_spin = "";
  status = status;
};

const onRecipientsClicked = async (e) => {
  if (r_nf_hc_spin === "nf-hc-spin" || e.target.className.includes("stepper")) {
    e.stopPropagation();
    return;
  }

  try {
    recipientsArray = [];
    r_nf_hc_spin = "nf-hc-spin";

    await sleep(100);

    // TODO REMOVE
    console.log("DEBUG: Start getRecipientsList:", $faucetAddress, $gasDonorAddress, numberOfRecipientsToFetch);
    let recipientCount = await getRecipientsList($faucetAddress, $gasDonorAddress, numberOfRecipientsToFetch);
    // TODO REMOVE
    console.log("DEBUG: Result recipientCount:", recipientCount);
    if (recipientCount > 0) {
      status.error = "OK - recipients loaded";
      status.color = 'colorgreen';
      /* activityBaseStatus = "Loading finished: " + donorCount + " Donors"; */
      timeAgoDisplayItems[0].base = "Loading finished: " + recipientCount + " Recipients";
      /* currentActivity = activityBaseStatus; */
      timeAgoDisplayItems[0].current = timeAgoDisplayItems[0].base;
    } else {
      status.error = "Found not one recipient? TODO: Hire John McAfee";
      status.color = 'colorred';
    }

  } catch (e) {
    // TODO REMOVE
    console.log("DEBUG onRecipientsClicked CATCH:", e);
    if (e && e.reason) {
      if (e.code) {
        if (e.step) {
          status.error = "code: " + e.code + " reason: " + e.reason + ", at: " + e.step;;
        } else {
          status.error = "code: " + e.code + " reason: " + e.reason;
        }
      } else {
        status.error = "reason: " + e.reason;
      }
      status.color = 'colorred';
    } else {
      status.error = e;
      status.color = 'colorred';
    }
  }
  status = status;
  r_nf_hc_spin = "";
  // TODO REMOVE
  console.log("DEBUG onRecipientsClicked FINISH:", status);
};

/* Returns number of donors loaded, fills array donorsArray */
const getDonorsList = async (faucetAddress, gasDonorAddress, count) => {
  try {
    // If you don't specify a //url//, Ethers connects to the default
    // (i.e. ``http:/\/localhost:8545``)
    //const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    //const signer = provider.getSigner(gasDonorAddress);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

    // Use a read-only signer for this function
    const signer = new ethers.VoidSigner(gasDonorAddress, provider);

    // The Contract object
    const faucet = new ethers.Contract(faucetAddress, abi, signer);

    const donorCount = await faucet.getDonorCount().catch((e) => { console.log("getDonorCount failed:", e); e.step = "faucet.getDonorCount()"; throw(e); });
    if (donorCount && donorCount > 0) {

      let newDonor = {};

      /* Loop over the first count donors */
      for (let i = 0; i < donorCount && i < count; i++) {
        newDonor = await getDonor(faucet, i);
        if (newDonor) {

          donorsArray.push(newDonor);
          // trigger react
          donorsArray = donorsArray;

          await sleep(500);
        }
      }
    }
    return donorsArray.length;

  } catch (e) {
    throw (e);
  }
};

/* Fetch one donor by index 0 - ? */
const getDonor = async (faucet, index) => {
  try {
    console.log("Start fetching Donor: ", index);
    let fetchedDonor = {
      name: "",
      country: "",
      index: -1,
      count: 0,
      balance: ethers.constants.Zero,
      address: ethers.constants.AddressZero,
      lastPaymentTime: 0,
      bindIndex: -1,
      error: "",
    };

    let startSeconds = new Date();
    if (logText.length == 0 || logText.length > countMaxLogLength) {
      logText = getTime() + ": Starting to fetch donor " + (index+1);
    } else {
      logText = logText + "\n" + getTime() + ": Starting to fetch donor " + (index+1);
    }

    // solidity: function getDonorByIndex(uint _i) public view returns (address _from, uint _totalBalance, uint _numPayments, string memory _name)
    let d = await faucet.getDonorByIndex(index).catch((e) => { console.log("getDonorByIndex", index, "failed:", e); e.step = "faucet.getDonorByIndex()"; throw (e);});
    if (d) {
      // TODO REMOVE
      console.log("DONOR: ", d);
      fetchedDonor.name = d._name;
      if (d._country === "0x5465727261000000") {    // Terra
        fetchedDonor.country = "🌎";
      } else if (d._country === "0x0000000000000000") {    // Zero-bytes
        fetchedDonor.country = "🌍";
      } else {
        fetchedDonor.country = ethers.utils.toUtf8String(d._country);
      }
      fetchedDonor.address = d._from;
      fetchedDonor.index = index;
      fetchedDonor.count = d._numPayments;
      fetchedDonor.balance = d._totalBalance;
      fetchedDonor.lastPaymentTime = 0;

      // donation details start with 1 !!
      let dd = await faucet.getDonationDetails(fetchedDonor.address, fetchedDonor.count).catch((e) => { console.log("getDonationDetails", fetchedDonor.address, "payment", fetchedDonor.count, "failed:", e); e.step = "faucet.getDonationDetails(address, count)"; throw (e);});
      if (dd) {
        console.log("DETAILS donor: ", dd);
        /* (uint _amount, bool _withdraw, uint _timestamp) */
        if (dd._withdraw === false) {
          fetchedDonor.lastPaymentTime = dd._timestamp*1000;   // block.timestamp is seconds since UNIX
          fetchedDonor.bindIndex = addTimeDisplay(" ", " ", fetchedDonor.lastPaymentTime);
        } else {
          let i = fetchedDonor.count - 1;
          while (i > 0 && dd._withdraw === true) {
            dd = await faucet.getDonationDetails(fetchedDonor.address, i).catch((e) => {console.log("getDonationDetails", fetchedDonor.address, "payment", fetchedDonor.count, "failed:", e); e.step = "faucet.getDonationDetails(address, count-1)"; throw (e);});
            if (dd) {
              fetchedDonor.lastPaymentTime = dd._timestamp*1000;
            }
            i--;
          }
          if (dd._withdraw === false) {
            fetchedDonor.lastPaymentTime = dd._timestamp*1000;   // block.timestamp is seconds since UNIX
            fetchedDonor.bindIndex = addTimeDisplay(" ", " ", fetchedDonor.lastPaymentTime);
            // TODO REMOVE
            console.log("DEBUG abnormal donor bindindex", fetchedDonor.bindIndex);
          } else {
            console.log("DEBUG abnormal donor bindindex", fetchedDonor.bindIndex);
            throw({reason:"DONOR_WITHOUT_DONATIONS", code: "BAD_DEV_OR_VM_BUG", step: "faucet.getDonationDetails() name=" + fetchedDonor.name});
          }


        }
      }
    } else {
      console.log("NOT D, DAMMIT");
      throw({reason:"NOT D, DAMMIT", code: "BAD_DEV_OR_VM_BUG", step: "faucet.getRecipientByIndex(index)"});
    }
    logText = logText + "\n" + getTime() + ": Fechting donor finished in " + (new Date() - startSeconds)/1000 + " seconds";

    // TODO REMOVE
    //console.log("DEBUG return fetchedDonor: ", fetchedDonor);
    return fetchedDonor;

  } catch (e) {
    throw(e);
  }
};

  /* Returns number of recipients loaded, fills array recipientsArray */
const getRecipientsList = async (faucetAddress, gasDonorAddress, count) => {
  try {
    // If you don't specify a //url//, Ethers connects to the default
    // (i.e. ``http:/\/localhost:8545``)
    //const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
    const provider = new ethers.providers.JsonRpcProvider($web3URL);

    // The provider also allows signing transactions to
    // send ether and pay to change state within the blockchain.
    // For this, we need the account signer...
    //const signer = provider.getSigner(gasDonorAddress);
    //const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);

    // Use a read-only signer for this function
    const signer = new ethers.VoidSigner(gasDonorAddress, provider);

    // The Contract object
    const faucet = new ethers.Contract(faucetAddress, abi, signer);

    const recipientCount = await faucet.getRecipientsCount().catch((e) => { console.log("getRecipientsCount failed:", e); e.step = "faucet.getRecipientsCount()"; throw(e); });
    if (recipientCount && recipientCount > 0) {

      let newRecipient = {};

      /* Loop over the first count recipients */
      for (let i = 0; i < recipientCount && i < count; i++) {
        newRecipient = await getRecipient(faucet, i);
        if (newRecipient) {

          recipientsArray.push(newRecipient);

          // TODO REMOVE
          console.log("DEBUG newRecipient pushed", newRecipient);

          // trigger react
          recipientsArray = recipientsArray;

          await sleep(500);
        }
      }
    }
    return recipientsArray.length;

  } catch (e) {
    throw (e);
  }
};

/* Fetch one Recipient by index 0 - ? */
const getRecipient = async (faucet, index) => {
  try {
    console.log("Start fetching Recipient: ", index);
    let fetchedRecipient = {
        name: "",
        country: "",
        index: -1,
        count: 0,
        balance: ethers.constants.Zero,
        address: ethers.constants.AddressZero,
        lastPaymentTime: 0,
        bindIndex: -1,
        error: "",
    };

    let startSeconds = new Date();
    if (logText.length == 0 || logText.length > countMaxLogLength) {
        logText = getTime() + ": Starting to fetch recipient " + (index+1);
    } else {
        logText = logText + "\n" + getTime() + ": Starting to fetch recipient " + (index+1);
    }

    // solidity: function getRecipientByIndex(uint _i) public view returns (address _recipient, uint _totalBalance, uint _numPayments, string memory _name, bytes8 _country) {
    let r = await faucet.getRecipientByIndex(index).catch((e) => { console.log("getRecipientByIndex", index, "failed:", e); e.step = "faucet.getRecipientByIndex(index)"; throw (e);});
    if (r) {
      // TODO REMOVE
      console.log("RECIPIENT: ", r);

      fetchedRecipient.name = r._name;
      if (r._country === "0x5465727261000000") {    // Terra
        fetchedRecipient.country = "🌎";
      } else if (r._country === "0x0000000000000000") {    // Zero-bytes
        fetchedRecipient.country = "🌍";
      } else {
        fetchedRecipient.country = ethers.utils.toUtf8String(r._country);
      }
      fetchedRecipient.address = r._recipient;
      fetchedRecipient.index = index;
      fetchedRecipient.count = r._numPayments;
      fetchedRecipient.balance = r._totalBalance;
      fetchedRecipient.lastPaymentTime = 0;

      // payment details start with 1 !!
      let pd = await faucet.getPaymentDetails(fetchedRecipient.address, fetchedRecipient.count).catch((e) => { console.log("getPaymentDetails", fetchedRecipient.address, "payment", fetchedRecipient.count, "failed:", e); e.step = "faucet.getPaymentDetails(address, count)"; throw (e);});
      if (pd) {
        console.log("DETAILS recipient: ", pd);
        /* (uint _amount, bool _withdraw, uint _timestamp) */
        if (pd._withdraw === true) {
          fetchedRecipient.lastPaymentTime = pd._timestamp*1000;   // block.timestamp is seconds since UNIX
          fetchedRecipient.bindIndex = addTimeDisplay(" ", " ", fetchedRecipient.lastPaymentTime);
          // TODO REMOVE
          console.log("DEBUG recipient bindindex", fetchedRecipient.bindIndex);
        } else {
          let i = fetchedRecipient.count - 1;
          while (i > 0 && pd._withdraw === false) {
            pd = await faucet.getPaymentDetails(fetchedRecipient.address, i).catch((e) => {  console.log("getPaymentDetails", fetchedRecipient.address, "payment", fetchedRecipient.count, "failed:", e); e.step = "faucet.getPaymentDetails(address, count-1)"; throw (e);});
            if (pd) {
                fetchedRecipient.lastPaymentTime = pd._timestamp*1000;
            }
            i--;
          } // while
          if (pd._withdraw === true) {
            fetchedRecipient.lastPaymentTime = pd._timestamp*1000;   // block.timestamp is seconds since UNIX
            fetchedRecipient.bindIndex = addTimeDisplay(" ", " ", fetchedRecipient.lastPaymentTime);
            // TODO REMOVE
            console.log("DEBUG abnormal recipient bindindex", fetchedRecipient.bindIndex);
          }  else {
            console.log("DEBUG abnormal recipient bindindex", fetchedRecipient.bindIndex);
            throw({reason:"RECIPIENT_WITHOUT_PAYMENTS", code: "BAD_DEV_OR_VM_BUG", step: "faucet.getPaymentDetails() name=" + fetchedRecipient.name});
          }

        } // else

      } // if
    } else {
      console.log("NOT R, DAMMIT");
      throw({reason:"NOT R, DAMMIT", code: "BAD_DEV_OR_VM_BUG", step: "faucet.getRecipientByIndex(index)"});
    };
    logText = logText + "\n" + getTime() + ": Fechting recipient finished in " + (new Date() - startSeconds)/1000 + " seconds";

    // TODO REMOVE
    //console.log("DEBUG return fetchedRecipient: ", fetchedRecipient);
    return fetchedRecipient;

  } catch (e) {
    throw(e);
  }
};

// Load details about smart contract
const getFaucetDetails = async (faucetAddress, gasDonorAddress) => {

  // console.log("getFaucetDetails From",gasDonorAddress, "FaucetAdr", faucetAddress, "abi", abi);
  let fetchStep = "Setup Connection";
  try {
      // If you don't specify a //url//, Ethers connects to the default
      // (i.e. ``http:/\/localhost:8545``)
      //const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
      const provider = new ethers.providers.JsonRpcProvider($web3URL);

      // The provider also allows signing transactions to
      // send ether and pay to change state within the blockchain.
      // For this, we need the account signer...
      // const signer = provider.getSigner(gasDonorAddress);
      // const signer = ethers.Wallet.fromEncryptedJsonSync(json, password);
      // Use a read-only signer for this function
      const signer = new ethers.VoidSigner(gasDonorAddress, provider);

      // The Contract object
      const faucet = new ethers.Contract(faucetAddress, abi, signer);

      // TODO REMOVE
      //console.log("Start fetching");
      fetchStep = "getVersion";
      let r = await faucet.getVersion(); /* .catch((e) => {    console.log("getVersion failed:", e);
                                                          logText = logText + "\n" + getTime() + ": getVersion failed: " + e;
                                                      }); */
      if (r) {
        // TODO REMOVE
        //console.log("Faucet Version:", r);
        result.version = r;
      };

      fetchStep = "getOwner";
      r = await faucet.getOwner(); /* .catch((e) => {  logText = logText + "\n" + getTime() + ": getOwner failed: " + e;
                                                  console.log("getOwner failed:", e); }); */
      if (r) {
        // TODO REMOVE
        //console.log("Faucet Owner:", r);
        result.owner = r;
      }

      fetchStep = "getBalance";
      r = await faucet.getBalance(); /* .catch((e) => {console.log("getBalance failed:", e);
                                                  logText = logText + "\n" + getTime() + ": getBalance failed: " + e;
                                                  }); */
      if (r) {
        // TODO REMOVE
        //console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r)));
        result.balance = r;
      }

      fetchStep = "getCalculatedBalance";
      r = await faucet.getCalculatedBalance();/* .catch((e) => { console.log("getCalculatedBalance failed:", e); }); */
      if (r) {
        // TODO REMOVE
        //console.log("Faucet Calculated Balance:", parseFloat(ethers.utils.formatEther(r)));
        result.calculatedBalance = r;
    }

      fetchStep = "getDonorCount";
      r = await faucet.getDonorCount(); /* .catch((e) => { console.log("getDonorCount failed:", e); }); */
      if (r) {
        // TODO REMOVE
        //console.log("Donor Count:", r);
        result.donorsCount = r;
      }

      fetchStep = "getRecipientsCount";
      r = await faucet.getRecipientsCount(); /* .catch((e) => { console.log("getRecipientsCount failed:", e); }); */
      if (r) {
        // TODO REMOVE
        //console.log("Recipients Count:", r);
        result.recipientsCount = r;
      }

      fetchStep = "getAddress";
      r = await faucet.getAddress(); /* .catch((e) => { console.log("getAddress failed:", e); }); */
      if (r) {
        // TODO REMOVE
        //console.log("Address:", r);
        result.contractAddress = r;
      }

      result.error = "OK";
      return result;
  }
  catch (e) {
    //console.log("CATCH:", e);
    e.step = fetchStep;
    throw(e);
  }
};

// add a new twitter style watcher event
function addTimeDisplay(base, current, last) {
  let newItem = {
    base: base,
    current: current,
    last: last,
    id: null,
  };
  let bindIndex = timeAgoDisplayItems.length;
  timeAgoDisplayItems.push(newItem);
  timeAgoDisplayItems[bindIndex].id = setInterval(
    () => {
          if (timeAgoDisplayItems[bindIndex]) {
            if (timeAgoDisplayItems[bindIndex].current.startsWith(timeAgoDisplayItems[bindIndex].base)) {
              timeAgoDisplayItems[bindIndex].current = timeAgoDisplayItems[bindIndex].base + " " + timeDifference(new Date(), timeAgoDisplayItems[bindIndex].last);
            }
          }
    }, 800 + Math.floor(Math.random() * 401)); // returns a random integer from 0 to 400
  return bindIndex;
}

// twitter style time ago
addTimeDisplay("Loading finished", "Details not loaded yet", new Date());

function onShow() {
  console.log("OnShow Home");

  if (timeAgoDisplayItems.length === 0) {
    // twitter style time ago
    addTimeDisplay("Loading finished", "Details not loaded yet", new Date());
  }

  // Initial load after 3 seconds for the user to notice it starts to load
  setTimeout(() => {
    loadFaucetStatistics();
  }, 3000);

  // Refresh faucet details every 60 seconds
  faucetIVId = setInterval( () => {
    loadFaucetStatistics();
  }, 600000);
}

function onHide() {
  console.log("OnHide Home");
  // remove all 'time ago' events
  timeAgoDisplayItems.forEach(element => {
    if (element.id !== undefined) {
    clearInterval(element.id);
    };
  });
  timeAgoDisplayItems = [];
  if (faucetIVId) {
    clearInterval(faucetIVId);
  }
  faucetIVId = null;
}

onMount(() => {
  console.log("OnMount Home");
});

onDestroy(() => {
  console.log("OnDestroy Home");
});
</script>
