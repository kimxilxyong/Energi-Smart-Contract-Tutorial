/* eslint-disable ParseError */
/* SPDX-License-Identifier: MIT
   Copyright (c) 2020 Kim Il Yong */

import { writable } from 'svelte/store';

const visitedSvelte = {
  info: false,
  wallet: false,
  web3: false,
  smartcontract: false,
};
const ethAccountObject = {
  address: "",
  encrypt: {},
  privateKey: "",
  sign: {},
  signTransaction: {},
};

export const visitedPages = writable(visitedSvelte);
export const isUnlocked = writable(false);
export const ethAccount = writable(ethAccountObject);

export const getTime = () => {

  let currentTime = new Date();
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

export const scrollTo = (id) => {
  //get the element to be scrolled
  let sct = document.getElementById(id);
  //make the parent to scroll into view, smoothly!
  sct.scrollIntoView({ block: "end", behavior: "smooth", });
  console.log("Scroll ", id, sct);
};


