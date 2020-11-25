// SPDX - License - Identifier: MIT
// Copyright(c) 2020 Kim Il Yong
// Utility functions to list and remove old version peers
// start with startPeerWatch()
// stop with stopPeerWatch()
// see who is connected: listAllPeers()

var currentVersion = "3.0.7";
var clearWatch;

function checkPeersByVersion(latestVersion, removePeer, showAll) {
    var invalidCount = 0;
    var wasRemoved = false;
    for (var i = 0; i < admin.peers.length; i++) {
    //for (var i in admin.peers) {
        if (admin.peers[i] === undefined || admin.peers[i] == undefined || !admin.peers[i]) {
            // there seems to be a bug somewhere under energi 3.0.6, sometimes i get an undefined peer
            // happens only very rarely and has probably to do with removing peers in a loop over peers
            //console.log("Undefined detected: i=",i,", '",admin.peers[i],"'");
            continue;
            console.log("DEBUG after continue? 🤮 i=",i);
        }
        //console.log("After undefined: i=",i,", '",admin.peers[i],"'");

        var name = String(admin.peers[i].name);
        var version = name.substring(9,14);
        if ((version != latestVersion) || showAll) {
            var enode = admin.peers[i].enode;

            if (version != latestVersion) {
                console.log("**** OLD Peer " + i + " name " + name + " version " + version);
                console.log(enode);
                invalidCount++;
            } else {
                if (showAll == true) {
                    console.log("Peer " + i + " name " + name + " version " + version);
                }
            }
            if (removePeer == true) {
                admin.removePeer(enode);
                wasRemoved = true;
                i--;
            }
        }
    }
    console.log("Total invalid Peers: " + invalidCount);
    return wasRemoved;
};

function listAllPeers() {
    return checkPeersByVersion(currentVersion, false, true);
}

function checkAllPeers() {
    return checkPeersByVersion(currentVersion, false, false);
}

function removeAllInvalidPeers() {
    return checkPeersByVersion(currentVersion, true, false);
}

function removeInvalidPeer(version) {
   return checkPeersByVersion(version, true, false);
}

function startPeerWatch() {
    clearWatch = setInterval(removeAllInvalidPeers, 1000*60);
    return clearWatch;
}

function stopPeerWatch() {
    return clearInterval(clearWatch);
}


