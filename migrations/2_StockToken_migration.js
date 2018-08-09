var TokenStock = artifacts.require('./StockToken');
var IdentityRegistry = artifacts.require('./IdentityRegistry')

module.exports = function(deployer) {
    deployer.deploy(IdentityRegistry).then((A) => {
        deployer.deploy(TokenStock,'ROKK','Rokk3rCrowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca',A.address);
    })
}