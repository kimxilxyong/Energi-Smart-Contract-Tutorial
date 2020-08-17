import { writable } from 'svelte/store';

let visitedSvelte = {
    info: false,
    wallet: false,
    web3: false,
    smartcontract: false,
  };

export const global = writable(visitedSvelte);
