/* eslint-disable ParseError */
// SPDX-License-Identifier: MIT -->
// Copyright (c) 2020 Kim Il Yong -->
// Version 1.0.0 -->


import { writable } from 'svelte/store';

// Version 106
export const faucetAddress = writable("0x57B98F76bB39546F97BccD1EF0A38b2d9E074494");
export const gasDonorAddress = writable("0x09ae1a5ddfd481cfd3cc4390b0e08a0832709a06");

export const web3URL = writable("https://nodeapi.test3.energi.network");

const visitedSvelte = {
  info: false,
  wallet: false,
  web3: false,
  smartcontract: false,
};
const ethAccountObject = {
  eth: {},
  address: "",
  encrypt: {},
  privateKey: "",
  sign: {},
  signTransaction: {},
  isManagedOnNode: false,
  isLocked: true,
  isManagedInBrowserWallet: false,
  isManagedByWebWallet: false,
  accountStore: "none",
};

export const visitedPages = writable(visitedSvelte);
export const isUnlocked = writable(false);
export const ethAccount = writable(ethAccountObject);




