function stakeAll() {
    var totalBal = 0;
    for (var acctNum in eth.accounts) {
        var acct = eth.accounts[acctNum];
        console.log(acct);
        personal.unlockAccount(acct,null,0,true);
    }
    miner.stakingStatus();
    return true;
};

