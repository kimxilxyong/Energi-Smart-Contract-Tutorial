// Utility functions to list and remove old version peers

var currentVersion = "3.0.6";
var clearWatch;

function checkPeersByVersion(latestVersion, removePeer, showAll) {
    var invalidCount = 0;
    var wasRemoved = false;
    for (var i in admin.peers) {
        if (admin.peers[i] === undefined)
           continue;
        var name = String(admin.peers[i].name);
        var version = name.substring(9,14);
        if ((version != latestVersion) || showAll) {
            var enode = admin.peers[i].enode;
            invalidCount++;
            if (version != latestVersion) {
                console.log("**** OLD Peer " + i + " name " + name + " version " + version);
            } else {
                console.log("Peer " + i + " name " + name + " version " + version); 
            }
            console.log(enode); 
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


