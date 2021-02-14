// SPDX-License-Identifier: MIT -->
// Copyright (c) 2020 Kim Il Yong -->
// Version 1.0.0 -->

import HomePage from '../pages/home.svelte';
import InfoPage from '../pages/info.svelte';
import Web3ProviderPage from '../pages/web3provider.svelte';
import SmartContractPage from '../pages/smartcontract.svelte';

import AboutPage from '../pages/about.svelte';
import NotFoundPage from '../pages/404.svelte';

// Export the route definition object
let spa_routes = {
  // Exact path
  '/': HomePage,
  '/info/': InfoPage,
  '/web3provider/': Web3ProviderPage,
  '/smartcontract/': SmartContractPage,
  '/about/': AboutPage,
  // Catch-all, must be last
  '*': NotFoundPage,
};

export default spa_routes;
