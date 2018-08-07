var Whitelistable = artifacts.require('./Whitelistable.sol')

module.exports = function(deployer){
    deployer.deploy(Whitelistable);
}