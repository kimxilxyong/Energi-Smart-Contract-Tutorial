/* eslint-disable ParseError */
// SPDX-License-Identifier: MIT -->
// Copyright (c) 2020 Kim Il Yong -->
// Version 1.0.0 -->


import { writable } from 'svelte/store';

// Version 106
export const faucetAddress = writable("0x1EDf7947F7b95bA658D0A74024Dd8092e4D4831c");
export const gasDonorAddress = writable("0x3b5abf9b81d5b62df65bf63e163463d0aa42e53b");

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




