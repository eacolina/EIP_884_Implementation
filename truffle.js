var HDWalletProvider = require("truffle-hdwallet-provider");
var mnemonic = "bag various stumble orchard print plate relief gossip defense spirit during raven"
module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
  networks: {
    ropsten: {
      provider: function() {
        return new HDWalletProvider(mnemonic, "https://ropsten.infura.io/tRmXv2MPxIZMNr29IaYX ")
      },
      network_id: 3,
      gas:4209000
    },
    development: {
      host:'localhost',
      port:7545,
      network_id:'*'
    }
  }
};
