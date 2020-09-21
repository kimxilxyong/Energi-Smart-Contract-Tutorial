web3.currentProvider.send({
  method: "personal_listWallets",
  params: ['0x771ddb07222a1f9442c91cf04f64f3164771bb62', 'latest'],
  jsonrpc: "2.0",
  id: new Date().getTime()
}, function (error, result) {console.log(error), console.log(result)})


web3.currentProvider.send({
  method: "personal_listWallets",
  params: [],
  jsonrpc: "2.0",
  id: new Date().getTime()
}, function (error, result) {console.log(error), console.log(result)})

web3.eth.call({
  method: "personal_listWallets",
  params: [],
  jsonrpc: "2.0",
  id: new Date().getTime()}, "latest", function (error, result) {if (error) {console.log(error)} else { console.log(result) }}) 
