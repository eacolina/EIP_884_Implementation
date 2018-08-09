var TokenStock = artifacts.require('./StockToken.sol');
var Whitlistable = artifacts.require('./Whitelistable.sol')

module.exports = function(deployer) {
    deployer.deploy(Whitlistable).then((A) => {
        deployer.deploy(TokenStock,'ROKK','Rokk3rCrowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca',A.address);
    })
    
}