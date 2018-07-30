const StockToken = artifacts.require('./StockToken.sol');

contract('StockToken', async(accounts) => {
    let instance
    const OWNER = accounts[0];
    
    beforeEach('Create a new contract instance', async () => {
        instance = await StockToken.new('ROKK','Rokk3r Crowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca');
    })

    it('should create a StockToken with correct parameters and owner', async() => {
        let name = await instance.name(); // get name of token
        let symbol = await instance.symbol(); // get symbol of token
        let supply = await instance.totalSupply(); // get supply of token 
        let hash = await instance.by_LawsHash(); // get hash of token
        let owner = await instance.owner();
        assert.equal(symbol, "ROKK","Symbol doesn't match"); // PERFROM assertions
        assert.equal(name, "Rokk3r Crowdbuild","Names doesn't match");
        assert.equal(supply, 1000000, "Supply doesn't match");
        assert.equal(hash, "022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca", "hash doesn't match");
        assert.equal(owner, accounts[0],"The owner is incorrect");
    })

    it('should allow address to be whitelisted', async () => {
        let account = accounts[1]; // account to be whitelisted
        const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
        assert.isFalse(isWhitelistedBefore, "The address was already whitelisted."); // The address should not be be whitlisted
        await instance.addAddressToWhitelist(account); // whitelist the account
        const isWhitelistedAfter = await instance.isWhitelisted(account); // Check if address is in whitelist
        assert.isTrue(isWhitelistedAfter, "Address was not whitelisted."); // Address should be whitelisted now
    })

    it('addAddressToWhitelist should throw if address is whitelisted already', async () =>{
        let account = accounts[0]; // account that should be whitlisted already
        const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
        assert.isTrue(isWhitelistedBefore, "The address was not whitlisted"); // This address should be whitelisted already
        try {
            await instance.addAddressToWhitelist(account); // Adding this address to the whitelist should throw
            assert.fail(); // if above line doesn't throw, throw a fail
        } catch (err) {
            assert.ok(/revert/.test(err.message),"There was no REVERT error") // check what king of error was thrown
        }
    })

    it("Only owner should be able to add to the whitelist, non owner should throw if attempt to add to whitelist", async () => {
        let account = accounts[1]; // non owner account
        const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
        assert.isFalse(isWhitelistedBefore, "The address was already whitelisted."); // The address should not be be whitlisted
        // Try-catch to check if error thrown
        try {
            await instance.addAddressToWhitelist(account,{from:account}) // try to whitelist account from non owner
            assert.fail(); // if above line doesn't throw, throw a fail
        } catch (err) {
            assert.ok(/revert/.test(err.message),"There was no REVERT error") // check what king of error was thrown
        }
    })

    it('should create a StockToken and create whitelisted address for owner and all tokens to it', async() => {
        let account = accounts[0];
        const isWhitelisted = await instance.isWhitelisted(account); // call the smart contract and see if address is whitelisted
        const balance = await instance.balanceOf(account); // get balance
        const totalSupply = await instance.totalSupply(); // get total supply of coins
        assert.isTrue(isWhitelisted,'The address is not whitelisted'); // address should be whitlisted as it owns the contract
        assert.isTrue(balance.eq(totalSupply), "The owner of the contract doesn't own all tokens"); //Owner's balance should be equal to total supply
    })
    
    it('should emit event AddressAddedToWhitelist with correct parameters', async() => {
        const account = accounts[1];
        const tx = await instance.addAddressToWhitelist(account,{from:OWNER}) // get TX receipt to parse logs and look for events
        const event = tx.logs[0];
        assert.equal(event.event, "AddressAddedToWhitelist", "AddressAddedToWhitelist was not the sent event"); // check if the right event was fired
        assert.equal(event.args.AuthorizedBy, OWNER, "The function was called from owner but log is not")
        assert.equal(event.args.AddressAdded, account, "The account added doesn't match")
    })
})