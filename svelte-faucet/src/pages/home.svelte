<Page name="home">
  <!-- Top Navbar -->
  <Navbar sliding={true}>
    <NavLeft>
      <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
    </NavLeft>
    <NavTitle sliding>Testnet Faucet</NavTitle>
    <NavRight>

    </NavRight>
</Navbar>
  <!-- Toolbar -->
  <Toolbar bottom>
    <Link>Right Link</Link>
    <Link iconIos="f7:menu" iconAurora="f7:menu" iconMd="material:menu" panelOpen="left" />
  </Toolbar>
  <!-- Page content -->

  <Card outline>
    <CardHeader style="color:var(--energi-color-green);">Energi Testnet faucet Smart-Contract details</CardHeader>
    <CardContent>
      <List>
        <ListItem style="color:var(--energi-color-turkis);" title="Version: {r.version}"></ListItem>
        <ListItem style="color:var(--energi-color-green);" title="Balance: { ethers.utils.formatEther(r.balance) } NRG"></ListItem>
        <ListItem style="color:var(--energi-color-green);" title="Calculated Balance: { ethers.utils.formatEther(r.calculatedBalance) } NRG"></ListItem>
        <!-- after={donorCountAfter} man &#xf263;  people &#xf007;-->
        {#if r.donorsCount > 0}
          <ListItem onClick="{(e) => {onDonorClicked(e);}}" title="Donors Count: {r.donorsCount}"  badge={r.donorsCount} footer="Click to load donors">
            <i slot="media" class="nf-icons nf-hc-fw {nf_hc_spin} colororange">&#xf263;</i>
            <span slot="after">
                  <Stepper name="Nr of TXs" value={numberOfDonorsToFetch}
                          onStepperChange={(v) => {numberOfDonorsToFetch = v}}
                          small wraps autorepeat autorepeatDynamic inputReadonly/>
            </span>
          </ListItem>
        {:else}
          <ListItem title="Donors Count: {r.donorsCount}"  badge={r.donorsCount} footer="No donors found 🙄">
            <i slot="media" class="nf-icons nf-hc-fw colorgrey">&#xf263;</i>
          </ListItem>
        {/if}
        <li>
          <ul>
            {#if donorsArray.length > 0}
              <ListItem title="Donors list: {donorsArray.length}"></ListItem>
              {#each donorsArray as donor, i}
                <ListItem title="{donor.name}" after="Donated: { ethers.utils.formatEther(donor.balance) } NRG &nbsp;&nbsp;" badge="{donor.index}">
                  <i slot="media" class="nf-icons nf-hc-fw colororange">&#xf007;</i>
                  <i slot="after-end">&nbsp;&nbsp;</i>
                </ListItem>
              {/each}
            {/if}
          </ul>
        </li>
        <ListItem title="Recepients Count: {r.recepientsCount}">
          <i slot="media" class="nf-icons nf-hc-fw {nf_hc_spinR} colororpink">&#xf263;</i>
        </ListItem>
        <!-- OK-Square U+F634  clipboard U+F429 -->
        <ListItem style="color:var(--energi-color-yellow);" title="Owner: {r.owner}" after="&nbsp;">
            <div slot="after-title" class="colortext" style="margin-left: 20px;">
              {#if r.owner !== ethers.constants.AddressZero}
                <span class="bround bgreen fss3 {isClipboardDoneClass1}">
                Copied!
                </span>
                <i on:click={ () => {
                                        isClipboardDoneClass1 = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(r.owner);
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

        <ListItem style="color:yellow;"title="Contract Address: {r.contractAddress}" after="&nbsp;">
            <div slot="after-title" class="colortext" style="margin-left: 20px;">
              {#if r.contractAddress !== ethers.constants.AddressZero}
                <span class="bround bgreen fss3 {isClipboardDoneClass2}">
                Copied!
                </span>
                <i on:click={ () => {
                                        isClipboardDoneClass2 = "animate__animated animate__fadeOut animate__delay-2s";
                                        copyToClipboard(r.contractAddress);
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
        <ListItem title="Result: {r.error}" class="{ r.error === "OK" ? 'colorgreen' : r.error.toString().includes("Error") ? 'colorred' : ''}"></ListItem>
        <ListItem title="Activity: {currentActivity}"></ListItem>
      </List>
    </CardContent>
  </Card>

  <Card outline>
    <CardHeader style="color:var(--energi-color-green);">Logs, stats and more</CardHeader>
    <CardContent>
        <List noHairlinesMd>
            <ListItem style="color:var(--energi-color-turkis);" title="JSON-RPC provider node: {$web3URL}"></ListItem>
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
                spellcheck="false">
                    <!-- <i slot="media" class="nf-icons nf-hc-fw animate__animated animate__pulse animate__infinite" style="color:{connectedColor};">&#xf817;</i> -->
                    <!-- <i slot="media" class="nf-icons nf-hc-fw {animatePulseConnect}" style="color:{connectedColor};">&#xf817;</i> -->
                    <i slot="media" class="nf-icons nf-hc-fw" style="">&#xf817;</i>
            </ListInput>
        </List>

    </CardContent>
  </Card>


  <Block strong>
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
  </List>
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
  import ethers from "ethers";
  import { tick, getContext } from "svelte";
  import {
    web3URL,
    faucetAddress,
    gasDonorAddress
  } from "../js/stores.js";
  //import {  } from "../js/web3utils.js";
  import { timeDifference, copyToClipboard, getTime, scrollTo } from "../js/utils.js";
  import {abi106} from '../js/abi.js';
  const abi = abi106;

  let animateSwingClass = "animate__animated animate__swing animate__infinite colorturkis";
  let isClipboardDoneClass1 = "hidden";
  let isClipboardDoneClass2 = "hidden";

  let nf_hc_spin = "";
  let nf_hc_spinR = "";
  let donorsColor;

  let activityBaseStatus = "Loading finished";
  let currentActivity = "Details not loaded yet";
  let lastUpdated = new Date();

  let donorCountBadge = "?";
  let donorCountAfter = "Loading..."
  let numberOfDonorsToFetch = 5;

  let donorsArray = [];

  let logText = "";

  const result = {
      version: 0,
      balance: ethers.constants.Zero,
      calculatedBalance: ethers.constants.Zero,
      donorsCount: 0,
      recepientsCount: 0,
      owner: ethers.constants.AddressZero,
      contractAddress: ethers.constants.AddressZero,
      error: 'none yet',
    };
  let r = result;

  function sleep(milliseconds) {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
  }

  const onDonorClicked = async (e) => {
    console.log("onDonorClicked", e);
/*     if (nf_hc_spin) {
      nf_hc_spin = "";
    } else {
      nf_hc_spin = "nf-hc-spin";
    } */

    try {
      donorsArray = [];
      nf_hc_spin = "nf-hc-spin";

      await sleep(1000);

      let donorCount = await getDonorList($faucetAddress, $gasDonorAddress, numberOfDonorsToFetch);
      console.log("Result donorCount:", donorCount);
      if (donorCount > 0) {
        r.error = "none";
        activityBaseStatus = "Loading finished: " + donorCount + " Donors";
        currentActivity = activityBaseStatus;
      } else {
        r.error = "Found not one donor!?! 🤮 This should not happen";
      }

    } catch (e) {
      r.error = e;
    }
    nf_hc_spin = "";
    r = r;
  }

  /* Returns number of donors loaded, fills array donorsArray */
  const getDonorList = async (faucetAddress, gasDonorAddress, count) => {
    try {

      //donorsArray.length = 0;

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

      const donorCount = await faucet.getDonorCount().catch((e) => { console.log("getDonorCount failed:", e); });
      if (donorCount && donorCount > 0) {
        console.log("donorCount:", donorCount);
        console.log("count:", count);

        let newDonor = {};

        /* Loop over the first count donors */
        for (let i = 0; i < donorCount && i < count; i++) {
          newDonor = await getDonor(faucet, i);
          if (newDonor) {

            console.log("1ALength", donorsArray.length);

            console.log("Push:", newDonor);
            donorsArray.push(newDonor);
            // trigger react
            donorsArray = donorsArray;

            console.log("2ALength", donorsArray.length);

            //await tick();
            await sleep(2000);
          }
        }
      }
      console.log("Donors Array:", donorsArray);
      return donorsArray.length;

    } catch (e) {
      throw (e);
    }
  }

  /* Fetch one donor by index 0 - ? */
  const getDonor = async (faucet, index) => {
    try {
      console.log("Start fetching Donor: ", index);
      let fetchedDonor = {
        name: "",
        index: -1,
        count: 0,
        balance: ethers.constants.Zero,
        address: ethers.constants.AddressZero,
        error: "",
      };

      console.log("start getDonorByIndex", index);
      // solidity: function getDonorByIndex(uint _i) public view returns (address _from, uint _totalBalance, uint _numPayments, string memory _name)
      let d = await faucet.getDonorByIndex(index).catch((e) => { console.log("getDonorByIndex", index, "failed:", e); });
      if (d) {
        console.log("Name", d._name);
        fetchedDonor.name = d._name;
        fetchedDonor.address = d._from;
        fetchedDonor.index = index;
        fetchedDonor.count = d._numPayments;
        fetchedDonor.balance = d._totalBalance;
      } else {
        console.log("NOT D, DAMMIT");;
      }

      return fetchedDonor;

    } catch (error) {
      console.log("CATCH getDonorList:", e);
      throw(e);
    }

  }

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

        console.log("Start fetching");
        fetchStep = "getVersion";
        let r = await faucet.getVersion(); /* .catch((e) => {    console.log("getVersion failed:", e);
                                                            logText = logText + "\n" + getTime() + ": getVersion failed: " + e;
                                                       }); */
        if (r) {
          console.log("Faucet Version:", r);
          result.version = r;
        }

        fetchStep = "getOwner";
        r = await faucet.getOwner(); /* .catch((e) => {  logText = logText + "\n" + getTime() + ": getOwner failed: " + e;
                                                    console.log("getOwner failed:", e); }); */
        if (r) {
          console.log("Faucet Owner:", r);
          result.owner = r;
        }

        fetchStep = "getBalance";
        r = await faucet.getBalance(); /* .catch((e) => {console.log("getBalance failed:", e);
                                                    logText = logText + "\n" + getTime() + ": getBalance failed: " + e;
                                                    }); */
        if (r) {
          console.log("Faucet Balance:", parseFloat(ethers.utils.formatEther(r)));
          result.balance = r;
        }

        fetchStep = "getCalculatedBalance";
        r = await faucet.getCalculatedBalance();/* .catch((e) => { console.log("getCalculatedBalance failed:", e); }); */
        if (r) {
          console.log("Faucet Calculated Balance:", parseFloat(ethers.utils.formatEther(r)));
          result.calculatedBalance = r;
        }

        fetchStep = "getDonorCount";
        r = await faucet.getDonorCount(); /* .catch((e) => { console.log("getDonorCount failed:", e); }); */
        if (r) {
          console.log("Donor Count:", r);
          result.donorsCount = r;
        }

        fetchStep = "getRecipientsCount";
        r = await faucet.getRecipientsCount(); /* .catch((e) => { console.log("getRecipientsCount failed:", e); }); */
        if (r) {
          console.log("Recipients Count:", r);
          result.recepientsCount = r;
        }

        fetchStep = "getAddress";
        r = await faucet.getAddress(); /* .catch((e) => { console.log("getAddress failed:", e); }); */
        if (r) {
          console.log("Address:", r);
          result.contractAddress = r;
        }

        result.error = "OK";
        return result;
    }
    catch (e) {
      //console.log("CATCH:", e);
      throw(fetchStep + " failed, " + e);
      //return e;
    }
  }

  const loadFaucetStatistics = () => {
    currentActivity = "Details loading ...";
    const startSeconds = new Date();

    if (logText.length > 1000) {
        logText = getTime() + ": Starting to fetch data from " + $web3URL;
    } else {
        logText = logText + "\n" + getTime() + ": Starting to fetch data from " + $web3URL;
    }

    getFaucetDetails($faucetAddress, $gasDonorAddress)
    .catch((e) => {
                    console.log(".catch:", e);
                    r.error = e;
                    r = r;
                    activityBaseStatus = "Loading failed";
                    currentActivity = activityBaseStatus;
                    lastUpdated = new Date();
                    logText = logText + "\n" + getTime() + ": " + e;
                    logText = logText + "\n" + getTime() + ": Fechting faucet data failed in " + (new Date() - startSeconds)/1000 + " seconds"
                  })
    .then((res) => {
                    console.log(".then:", res);
                    if (res) {
                      r = res;
                      activityBaseStatus = "Loading finished";
                      currentActivity = activityBaseStatus;
                      lastUpdated = new Date();
                      logText = logText + "\n" + getTime() + ": Fechting faucet data finished in " + (new Date() - startSeconds)/1000 + " seconds"
                    }
                  });

  };


  // Initial load after 5 seconds for the user to notice it starts to load
  setTimeout(() => {
      loadFaucetStatistics();
  }, 5000);

  // Refresh faucet details every 60 seconds
  setInterval(() => {
      loadFaucetStatistics();
  }, 60000);

  // twitter style time ago
  setInterval(() => {
    if (currentActivity.startsWith(activityBaseStatus)) {
      currentActivity = activityBaseStatus + " " + timeDifference(new Date(), lastUpdated);
    }
  }, 1050);

</script>
