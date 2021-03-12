#!/bin/bash
# SPDX-License-Identifier: MIT
# Copyright (c) 2020 Kim Il Yong
#####################################################################
# Description: This script is to start energi core in testmode as a web3 provider for remix desktop IDE
# 🥶 ATTENTION: don't run --rpccorsdomain "*" on a server with real money! 🥵
#
# Download Remix
# https://github.com/ethereum/remix-desktop/releases
#
# Attach Remix to the local energi3 testnode
# DEPLOY & RUN TRANSACTIONS -> ENVIRONMENT -> Web3 Provider -> http://localhost:49796 -> OK
#
#####################################################################


./energi3 --testnet --gcmode=archive --syncmode=full --verbosity 3 --rpccorsdomain "package://a7df6d3c223593f3550b35e90d7b0b1f.mod" --vmdebug --rpc --rpcport 49796 --rpcaddr "127.0.0.1" --rpcvhosts "localhost" --rpcapi admin,web3,eth,debug,personal,net,energi console


