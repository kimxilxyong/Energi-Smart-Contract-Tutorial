#!/bin/bash
# SPDX-License-Identifier: MIT
# Copyright (c) 2020 Kim Il Yong 

 
#####################################################################
# Description: This script is to start energi core in testmode as a web3 provider for remix desktop IDE
#
# Download Remix
# https://github.com/ethereum/remix-desktop/releases
#
# Attach Remix to the local energi3 testnode
# DEPLOY & RUN TRANSACTIONS -> ENVIRONMENT -> Web3 Provider -> http://localhost:39796 -> OK
#
#####################################################################


./energi3 --testnet --preload utils.js,peers.js attach


