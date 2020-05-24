function checkAllBalances() {
    var totalBal = 0;
    for (var acctNum in eth.accounts) {
        var acct = eth.accounts[acctNum];
        var acctBal = web3.fromWei(eth.getBalance(acct), "ether");
        totalBal += parseFloat(acctBal);
        console.log("  eth.accounts[" + acctNum + "]: \t" + acct + " \tbalance: " + acctBal + " NRG");
    }
    console.log("  Total balance: " + totalBal + " NRG");
    return true;
};

function checkBalance(acct) {
    var acctBal = web3.fromWei(eth.getBalance(acct), "ether");
    console.log(" account:\t" + acct + " \tbalance: " + acctBal + " NRG");
    return true;
};

function mnBalances() {
    var totalBal = 0;
    for (var acctNum in eth.accounts) {
        var acct = eth.accounts[acctNum];
        var acctBal = web3.fromWei(masternode.masternodeInfo(acct).collateral, "ether");
        totalBal += parseFloat(acctBal);
        console.log("  mn.accounts[" + acctNum + "]: \t" + acct + " \tbalance: " + acctBal + " NRG");
    }
    console.log("  Total balance: " + totalBal + " NRG");
    return true;
};


// Utility functions to list and remove old version peers

var currentVersion = "3.0.6";

function checkPeersByVersion(latestVersion, removePeer) {
    var invalidCount = 0;
    for (var i in admin.peers) {
        if (admin.peers[i] === undefined)
           continue;
        var name = String(admin.peers[i].name);
        var version = name.substring(9,14);
        if (version != latestVersion) {
            var enode = admin.peers[i].enode;
            invalidCount++;
            console.log("Old Peer " + i + " name " + name + " version " + version);
            console.log(enode); 
            if (removePeer == true) {
                admin.removePeer(enode);
                i--;
            }
        }
    }
    console.log("Total invalid Peers: " + invalidCount);
    return true;
};

function checkAllPeers() {
    return checkPeersByVersion(currentVersion, false);
};

function removeAllInvalidPeers() {
    return checkPeersByVersion(currentVersion, true);
};

function removeInvalidPeer(version) {
   return checkPeersByVersion(version, true);
};


