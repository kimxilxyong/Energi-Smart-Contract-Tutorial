#!/bin/bash
# SPDX-License-Identifier: MIT
# Copyright (c) 2020 Kim Il Yong
##############################################################################################################
# Description: This script is to start energi core in testmode as a provider for local browser or node Web3.js
# 🥶 ATTENTION: don't run --rpccorsdomain "*" on a server with real money! 🥵
#
# Download Web3
# npm install web3
#
#
#  let p = web3.setProvider(new Web3.providers.HttpProvider("http://localhost:49796"));
#  console.log('Provider: ', p);
#
##############################################################################################################

./energi3 --testnet --verbosity 3 --vmdebug --gcmode=archive --syncmode=full --preload utils.js,peers.js --rpccorsdomain "*" --rpc --rpcaddr "127.0.0.1" --rpcvhosts "localhost" --rpcapi admin,web3,eth,debug,personal,net,energi console
