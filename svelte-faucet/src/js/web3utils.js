/* SPDX-License-Identifier: MIT
   Copyright (c) 2020 Kim Il Yong */

//import Web3 from 'web3';

export const listWallets = async (web3) => {
  try {
    const result = await sendRPC_personal_listWallets(web3);
    return result;
  } catch (e) {
    return e;
  }
};

function sendRPC_personal_listWallets(web3) {
  return new Promise((resolve, reject) => {
    web3.currentProvider.send(
      {
        method: "personal_listWallets",
        params: [],
        jsonrpc: "2.0",
        id: new Date().getTime(),
      },
      function (error, result) {
        if (error) {
          reject(error);
        } else {
          resolve(result);
        }
      }
    );
  });
}
