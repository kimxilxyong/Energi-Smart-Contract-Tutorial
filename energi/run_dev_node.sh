#!/bin/bash
# SPDX-License-Identifier: MIT
# Copyright (c) 2020 Kim Il Yong
##############################################################################################################
# Description: This script is to start energi core in testmode as a provider for the faucet deployed on a local running microhttp
# 🥶 ATTENTION: don't run --rpccorsdomain "*" on a server with real money! 🥵
#
# Download ethers.js
# npm install ethers
#
# In the nodes console:
#
#  const ethers = require('ethers');
#  const provider = new ethers.providers.JsonRpcProvider('http://localhost:49796');
#
#  console.log('Provider: ', provider);
#
##############################################################################################################

./energi3 --testnet --verbosity 3 --gcmode=archive --syncmode=full --preload utils.js,peers.js,send.js,stake.js --rpccorsdomain "*" --rpc --rpcaddr "127.0.0.1" --rpcvhosts "localhost" --rpcapi admin,web3,eth,debug,personal,net,energi console

