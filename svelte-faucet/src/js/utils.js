// SPDX-License-Identifier: MIT -->
// Copyright (c) 2020 Kim Il Yong -->
// Version 1.0.0 -->


/* Get the country flag
   1be9a6884abd4c3ea143b59ca317c6b2
   645728ce29a740769339b7493f06fa17
   https://ipgeolocation.abstractapi.com/v1/?api_key=1be9a6884abd4c3ea143b59ca317c6b2&ip_address=127.5.1.2
   https://ipgeolocation.abstractapi.com/v1/?api_key=1be9a6884abd4c3ea143b59ca317c6b2&ip_address=2600:8802:901:1500:24f3:91fa:64cf:bec2
*/
import detectEthereumProvider from '@metamask/detect-provider';
import Bowser from "bowser";


export function errorToCode(e) {
  let code = "UNKNOWN_ERROR";

  try {
    if (e.hasOwnProperty("code")) {
      code = e.code;
    } else if (e.hasOwnProperty("errorCode")) {
      code = e.errorCode;
    } else if (e.hasOwnProperty("error")) {
      if (e.error.hasOwnProperty("code")) {
        code = e.error.code;
      } else if (e.error.hasOwnProperty("errorCode")) {
        code = e.error.errorCode;
      } else if (e.error.name) {
        code = e.error.name;
      }
    } else if (e.name) {
      code = e.name;
    }
  } catch (error) {
    code = "JAVASCRIPT_HASOWNPROPERTY_THROWED";
  }

  return code;
}

export function errorToReason(e) {
  let reason = "Unknown error occured";

  try {
    if (e.hasOwnProperty("reason")) {
      reason = e.reason;
    } else if (e.hasOwnProperty("message")) {
      reason = e.message;
    } else if (e.hasOwnProperty("errorReason")) {
      reason = e.errorReason;
    } else if (e.hasOwnProperty("error")) {
      if (e.error.hasOwnProperty("message")) {
        reason = e.error.message;
      } else if (e.error.hasOwnProperty("reason")) {
        reason = e.error.reason;
      } else if (e.error.hasOwnProperty("errorReason")) {
        reason = e.error.errorReason;
      } else {
        reason = JSON.stringify(e.error, null, 2);
      }
    } else {
      reason = JSON.stringify(e, null, 2);
    }

    if (e.hasOwnProperty("step")) {
      reason += "\n" + e.step;
    } else if (e.hasOwnProperty("errorStep")) {
      reason += "\n" + e.errorStep;
    }

  } catch (error) {
    if (error.message) {
      reason = error.message;
    } else {
      reason = error;
    }
  }

  return reason;
}


export const openCryptoWallet = async () => {
  try {
    if (window.ethereum) {
      /*********************************************/
      /* Access the user's accounts (per EIP-1102) */
      /*********************************************/

      // You should only attempt to request the user's accounts in response to user
      // interaction, such as a button click.
      // Otherwise, you popup-spam the user like it's 1999.
      // If you fail to retrieve the user's account(s), you should encourage the user
      // to initiate the attempt.
      let accounts = await window.ethereum.request({ method: 'eth_requestAccounts', }).catch((err) => {
        if (err.code === 4001) {
          // EIP-1193 userRejectedRequest error
          // If this happens, the user rejected the connection request.
          console.log('Please connect to MetaMask.');
          throw { reason: "User rejected get accounts request", code: "USER_REJECTED_REQUEST", step: "eth_requestAccounts", };
        } else {
          err.step = "eth_requestAccounts";
          console.error(err);
          throw (err);
        }
      });

      if (accounts) {
        console.log(accounts);
        return true;
      }
      //return { reason: "User rejected get accounts request", code: "USER_REJECTED_REQUEST", step: "eth_requestAccounts", };
      return { reason: "No account selected in browser wallet", code: "NO_ACCOUNT_FOUND", step: "eth_requestAccounts", };
    }
  } catch (error) {
    return error;
  }
};

export const initBrowserDetails = () => {
  let result = {
    browser: { name: "", version: "0.0.0", details: null, icon: "./images/unknown32x32.png", error: null, },
    wallet: { exists: false, enabled: false, chain: "-1", version: "0.0.0", connected: false, unlocked: false, activeAccount: null, accounts: [], provider: null, error: null, callbackAccountsChanged: null, },
    todo: { text: "", url: "", },
    date: 0,
    time: "",
  };
  return result;
};

export const detectBrowser = (result) => {

  if (!result) {
    result = initBrowserDetails();
    result.date = new Date();
    result.time = getTime();
  }

  try {
    // Detect the users browser
    // TODO REMOVE
    //console.log("window.navigator", window.navigator);
    //console.log("window", window);
    //console.log("window.Notification", window.Notification);

    let browser = Bowser.getParser(window.navigator.userAgent);
    // TODO REMOVE
    //console.log("The current browser is", browser);

    result.browser.details = browser;
    result.browser.name = browser.getBrowser().name;
    result.browser.version = browser.getBrowser().version;
    result.browser.error = null;

    if (window.navigator.brave) {
      result.browser.name = "Brave";
    }

    if (result.browser.name === "Brave") {
      result.browser.icon = "./images/Brave32x32.png";
    } else if (result.browser.name === "Safari") {
      result.browser.icon = "./images/Safari32x32.png";
    } else if (result.browser.name === "Microsoft Edge") {
      result.browser.icon = "./images/MicrosoftEdge32x32.png";
    } else if (result.browser.name === "Chrome") {
      result.browser.icon = "./images/Chrome32x32.png";
    } else if (result.browser.name === "Firefox") {
      result.browser.icon = "./images/Firefox32x32.png";
    } else if (result.browser.name === "Opera") {
      result.browser.icon = "./images/Opera32x32.png";
    }


  } catch (error) {
    error.step = "Get Browser ID";
    result.browser.name = "Unknown";
    result.browser.error = error;
  }
  return result;
};

export const getBrowserDetails = async (e) => {

  let result = initBrowserDetails();

  result.date = new Date();
  result.time = getTime();

  if (!window) {
    console.log("window not defined?");
    result.browser.name = "Unknown";
    result.browser.details = null;
    result.browser.error = { reason: "window not defined", code: "INVALID_RUNTIME", step: "check js runtime", };
    return result;
  }

  result = detectBrowser(result);
  try {
    result = await detectWallet(result);
  } catch (error) {
    result.wallet.error = error;
  }

  return result;
};

export const detectWallet = async (result) => {
  try {
    if (!result) {
      result = initBrowserDetails();
    }
    if (window.ethereum) {

      result.wallet.exists = true;

      /*****************************************/
      /* Detect the MetaMask Ethereum provider */
      /*****************************************/
      let browserProvider = await detectEthereumProvider();
      // TODO REMOVE
      //console.log("browserProvider:", browserProvider);
      //console.log("window.ethereum:", window.ethereum);

      if (browserProvider !== window.ethereum) {
        console.error('ERROR Do you have multiple wallets installed?');
        result.wallet.provider = browserProvider;
        throw ({ reason: "Do you have multiple wallets installed?", code: "INVALID_ETHEREUM_PROVIDER", step: "detectEthereumProvider()", });
      }

      /**********************************************************/
      /* Handle chain (network) and chainChanged (per EIP-1193) */
      /**********************************************************/

      result.wallet.chain = browserProvider.chainId;

      /***********************************************************/
      /* Handle user accounts and accountsChanged (per EIP-1193) */
      /***********************************************************/

      result.wallet.accounts = await browserProvider.request({ method: 'eth_accounts', }).catch((e) => { e.step = "request eth_accounts"; throw (e); });

      result.wallet.provider = browserProvider;

      if (result.wallet.accounts && result.wallet.accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        console.log('Please connect to MetaMask.');
        result.wallet.error = { reason: "No account selected in browser wallet", code: "NO_ACCOUNT_FOUND", step: "get accounts from wallet", };
        result.todo.text = "Unlock account(s)";
        result.todo.url = "brave://wallet/";
        result.wallet.enabled = false;
      } else {
        result.wallet.enabled = true;
        console.log("result.wallet.accounts", result.wallet.accounts);
      }
      result.wallet.connected = browserProvider._state.isConnected;
      result.wallet.unlocked = browserProvider._state.isUnlocked;
      result.wallet.version = browserProvider.networkVersion;

    } else {
      // TODO REMOVE
      //console.log("window.ethereum not defined!");
      result.wallet.error = { reason: "Browser injected crypto wallet not enabled", code: "NO_ETHEREUM_PROVIDER", step: "check window.ethereum", };
      /* brave://settings/extensions
      Ethereum provider for using Dapps => Crypto Wallets */
    }

  } catch (error) {
    //console.log("detectWallet catch:", error);
    if (error.message) {
      error.reason = error.message;
    }
    result.wallet.error = error;
  }
  return result;
};

export const requestWalletAccounts = async (result) => {
  try {
    if (window.ethereum) {

      result.wallet.exists = true;

      /*****************************************/
      /* Detect the MetaMask Ethereum provider */
      /*****************************************/
      let browserProvider = await detectEthereumProvider();
      // TODO REMOVE
      //console.log("browserProvider:", browserProvider);
      //console.log("window.ethereum:", window.ethereum);

      if (browserProvider !== window.ethereum) {
        console.error('ERROR Do you have multiple wallets installed?');
        result.wallet.provider = browserProvider;
        throw ({ reason: "Do you have multiple wallets installed?", code: "INVALID_ETHEREUM_PROVIDER", step: "detectEthereumProvider()", });
      }

      /**********************************************************/
      /* Handle chain (network) and chainChanged (per EIP-1193) */
      /**********************************************************/

      result.wallet.chain = browserProvider.chainId;

      /***********************************************************/
      /* Handle user accounts and accountsChanged (per EIP-1193) */
      /***********************************************************/

      result.wallet.accounts = await browserProvider.request({ method: 'eth_accounts', }).catch((e) => { e.step = "request eth_accounts"; throw (e); });

      result.wallet.provider = browserProvider;

      if (result.wallet.accounts && result.wallet.accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        console.log('Please connect to MetaMask.');
        result.wallet.error = { reason: "No account selected in browser wallet", code: "NO_ACCOUNT_FOUND", step: "get accounts from wallet", };
        result.todo.text = "Unlock account(s)";
        result.todo.url = "brave://wallet/";
        result.wallet.enabled = false;
      } else {
        result.wallet.enabled = true;
        console.log("result.wallet.accounts", result.wallet.accounts);
      }
      result.wallet.connected = browserProvider._state.isConnected;
      result.wallet.unlocked = browserProvider._state.isUnlocked;
      result.wallet.version = browserProvider.networkVersion;

      // add the event handler for accounts changed
      if (result.wallet.callbackAccountsChanged) {
        browserProvider.on('accountsChanged', result.wallet.callbackAccountsChanged);
      }

      /*********************************************/
      /* Access the user's accounts (per EIP-1102) */
      /*********************************************/

      // You should only attempt to request the user's accounts in response to user
      // interaction, such as a button click.
      // Otherwise, you popup-spam the user like it's 1999.
      // If you fail to retrieve the user's account(s), you should encourage the user
      // to initiate the attempt.
      if (result.wallet.enabled === false) {
        result.wallet.accounts = await browserProvider.request({ method: 'eth_requestAccounts', }).catch((err) => {
          if (err.code === 4001) {
            // EIP-1193 userRejectedRequest error
            // If this happens, the user rejected the connection request.
            console.log('Please connect to MetaMask.');
            result.wallet.error = { reason: "User rejected the connection request", code: "REQUEST_REJECTED", step: "request accounts from wallet", };
            result.todo.text = "Unlock account(s)";
            result.todo.url = "https://www.google.com/search?q=how+to+unlock+metamask";
          } else if (err.code === -32002) {
            console.log('Request pending', err);
            result.wallet.error = { reason: "Permission request already pending, please complete the Crypto Wallet setup", code: "REQUEST_PENDING", step: "request accounts from wallet", };
            result.todo.text = "Finish crypto wallet setup";
            result.todo.url = "https://www.google.com/search?q=how+to+unlock+metamask";

          } else {
            result.wallet.error = { reason: err.message, code: err.code, step: "request accounts from wallet", };
            console.error(err);
            result.todo.text = "Unknown error while getting accounts";
            result.todo.url = "https://www.google.com/search?q=how+to+unlock+metamask";
          }
        });
      }

    } else {
      console.log("window.ethereum not defined!");
      result.wallet.error = { reason: "Browser injected crypto wallet not enabled", code: "NO_ETHEREUM_PROVIDER", step: "check window.ethereum", };
      /* brave://settings/extensions
      Ethereum provider for using Dapps => Crypto Wallets */
    }

  } catch (error) {
    //console.log("requestWalletAccounts catch:", error);
    if (error.message) {
      error.reason = error.message;
    }
    result.wallet.error = error;
  }
  return result;
};


export const getCountryFlag = async () => {
  try {
    let r = await fetch("https://ipgeolocation.abstractapi.com/v1/?api_key=645728ce29a740769339b7493f06fa17", {
      method: 'GET',
      credentials: 'omit',
      cache: 'no-cache',
      headers: {
        'Content-Type': 'application/json',
        // 'Content-Type': 'application/x-www-form-urlencoded',
      },
      redirect: 'follow',
      referrerPolicy: 'no-referrer',
    });
    // TODO REMOVE debug
    //console.log("Flag raw:", r);

    let flag = "🇦🇶";
    if (r && r.status === 200) {
      try {
        r = await r.json();
        // TODO REMOVE
        //console.log("Flag json:", r);
      } catch (e) {
        return "🏴󠁵󠁳󠁴󠁸󠁿";
      }
      if (r.error) {
        if (r.error.details) {
          if (r.error.details.ip_address) {
            flag = "🚣";
          } else {
            flag = "🏴‍☠️";
          }
        } else {
          flag = "🎌";
        }
      } else {
        try {
          flag = r.flag.emoji;
          if (r.security && r.security.is_vpn && r.security.is_vpn === true) {
            flag = "🏴‍☠️" + flag;
          }
          // TODO REMOVE debug
          //console.log("Flag JSON", r.json());
        } catch (e) {
          if (!flag) {
            flag = "🇦🇶";
          }
        }
      }
    } else {
      flag = "🇺🇳";
    }
    return flag;

  } catch (e) {
    //throw (e);
    //return "🎌";
    return "🏴󠁵󠁳󠁴󠁸󠁿";
  }
};

export const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(
    () => {
        console.log("Copied", text, "to clipboard");
    },
    (e) => {
        console.log("Copy to clipboard failed", e);
    },
  );
};

const stringInject = (str = '', obj = {}) => {
  let newStr = str;
  Object.keys(obj).forEach((key) => {
    let placeHolder = `#${key}#`;
    if (newStr.includes(placeHolder)) {
      newStr = newStr.replace(placeHolder, obj[key] || " ");
    }
  });
  return newStr;
};

export const getTime = (now) => {

  let currentTime = new Date();
  if (now) {
    currentTime = new Date(now);
  }

  let hours = currentTime.getHours();
  let minutes = currentTime.getMinutes();
  let seconds = currentTime.getSeconds();

  if (minutes < 10) {
    minutes = "0" + minutes;
  }
  if (seconds < 10) {
    seconds = "0" + seconds;
  }
  return hours + ":" + minutes + ":" + seconds;
};

export const timeDifference = (current, previous) => {
  const msPerSecond = 1000;
  const msPerMinute = 60 * msPerSecond;
  const msPerHour = msPerMinute * 60;
  const msPerDay = msPerHour * 24;
  const msPerWeek = msPerDay * 7;
  const msPerMonth = msPerDay * 30;
  const msPerYear = msPerDay * 365;
  const msPerCentury = msPerYear * 100;
  const second = "just now";
  const minute = "#count# second#one# ago";
  const hour = "#count# minute#one# ago";
  const day = "#count# hour#one# ago";
  const week = "#count# day#one# ago";
  const month = "#count# week#one# ago";
  const year = "#count# month#one# ago";
  const century = "#count# year#one# ago";

  const units = [{ ms: msPerSecond, u: second, },
                { ms: msPerMinute, u: minute, },
                { ms: msPerHour, u: hour, },
                { ms: msPerDay, u: day, },
                { ms: msPerWeek, u: week, },
                { ms: msPerMonth, u: month, },
                { ms: msPerYear, u: year, },
                { ms: msPerCentury, u: century, },
               ];

  const elapsed = current - previous;
  // TODO REMOVE
  //console.log("timeDifference: current", current, "previous", previous, "elapsed", elapsed);

  let result = "";
  let ago = 101;

  for (let i = 0; i < units.length; i++) {
    let isOne = "s";
    if (elapsed < units[i].ms) {
      if (i > 0) {
        ago = Math.round(elapsed / units[i - 1].ms);
      }
      if (ago === 1) {
        isOne = "";
      }
      result = stringInject(units[i].u, { count: ago, one: isOne, });
      // TODO REMOVE
      //console.log("AGO Result:", ago, result);
      break;
    }
  }
  if (result === "") {
    result = "More than 100 years ago";
  }
  return result;

  // TODO REMOVE
/*   if (elapsed < msPerMinute) {
    return Math.round(elapsed / 1000) + ' seconds ago';
  } else if (elapsed < msPerHour) {
    return Math.round(elapsed / msPerMinute) + ' minutes ago';
  } else if (elapsed < msPerDay) {
    return Math.round(elapsed / msPerHour) + ' hours ago';
  } else if (elapsed < msPerMonth) {
    return 'approximately ' + Math.round(elapsed / msPerDay) + ' days ago';
  } else if (elapsed < msPerYear) {
    return 'approximately ' + Math.round(elapsed / msPerMonth) + ' months ago';
  } else {
    return 'approximately ' + Math.round(elapsed / msPerYear) + ' years ago';
  } */
};

export const scrollTo = (id) => {
  //get the element to be scrolled
  let sct = document.getElementById(id);
  //make the parent to scroll into view, smoothly!
  sct.scrollIntoView({ block: "end", behavior: "smooth", });
  console.log("Scroll ", id, sct);
};

export function sleep(milliseconds) {
  return new Promise(resolve => setTimeout(resolve, milliseconds));
}
