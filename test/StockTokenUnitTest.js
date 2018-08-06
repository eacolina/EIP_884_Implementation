/*
IMPORTANT NOTE: These tests do not cover the basic functionality of the ERC20 standard as I'm using the implementation of Open Zeppelin which has been thoroughly audited. 
*/

const StockToken = artifacts.require('./StockToken.sol');

contract('StockToken', async(accounts) => {
    let instance
    const OWNER = accounts[0];
    
    beforeEach('Create a new contract instance', async () => {
        instance = await StockToken.new('ROKK','Rokk3r Crowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca');
    })
    describe('constructor', () => {
        it('should create a StockToken with correct parameters and owner', async() => {
            let name = await instance.name(); // get name of token
            let symbol = await instance.symbol(); // get symbol of token
            let supply = await instance.totalSupply(); // get supply of token 
            let hash = await instance.byLawsHash(); // get hash of token
            let owner = await instance.owner();
            let decimals = await instance.decimals();
            assert.equal(symbol, "ROKK","Symbol doesn't match"); // PERFROM assertions
            assert.equal(name, "Rokk3r Crowdbuild","Names doesn't match");
            assert.equal(supply, 1000000, "Supply doesn't match");
            assert.equal(hash, "022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca", "hash doesn't match");
            assert.equal(owner, accounts[0],"The owner is incorrect");
            assert.equal(decimals.toNumber(), 0, "The decimals are not 0.");
        })

        it('should create a StockToken and create whitelisted address for owner and all tokens to it', async() => {
            let account = accounts[0];
            const isWhitelisted = await instance.isWhitelisted(account); // call the smart contract and see if address is whitelisted
            const balance = await instance.balanceOf(account); // get balance
            const totalSupply = await instance.totalSupply(); // get total supply of coins
            assert.isTrue(isWhitelisted,'The address is not whitelisted'); // address should be whitelisted as it owns the contract
            assert.isTrue(balance.eq(totalSupply), "The owner of the contract doesn't own all tokens"); //Owner's balance should be equal to total supply
        })
    })
    describe('addAddressToWhitelist', () =>{
        it('should allow address to be whitelisted by owner when address is not in whitelist', async () => {
            let account = accounts[1]; // account to be whitelisted
            const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isFalse(isWhitelistedBefore, "The address was already whitelisted."); // The address should not be be whitelisted
            await instance.addAddressToWhitelist(account); // whitelist the account
            const isWhitelistedAfter = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isTrue(isWhitelistedAfter, "Address was not whitelisted."); // Address should be whitelisted now
        })

        it('should throw if address is whitelisted already', async () =>{
            let account = accounts[0]; // account that should be whitelisted already
            const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isTrue(isWhitelistedBefore, "The address was not whitelisted"); // This address should be whitelisted already
            try {
                await instance.addAddressToWhitelist(account); // Adding this address to the whitelist should throw
                assert.fail(); // if above line doesn't throw, throw a fail
            } catch (err) {
                assert.ok(/revert/.test(err.message),"There was no REVERT error") // check what king of error was thrown
            }
        })

        it("should throw when non-owner tries to whitelist", async () => {
            let account = accounts[1]; // non owner account
            const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isFalse(isWhitelistedBefore, "The address was already whitelisted."); // The address should not be be whitelisted
            // Try-catch to check if error thrown
            try {
                await instance.addAddressToWhitelist(account,{from:account}) // try to whitelist account from non owner
                assert.fail(); // if above line doesn't throw, throw a fail
            } catch (err) {
                assert.ok(/revert/.test(err.message),"There was no REVERT error") // check what king of error was thrown
            }
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

    describe('transfer', () => {
        it('should allow for transfer of tokens to a whitelisted address', async () => {
            const destAccount = accounts[1]; // address to which the token will be sent
            const initDestBalance = await instance.balanceOf(destAccount);
            const initOrigBalance = await instance.balanceOf(OWNER);
            await instance.addAddressToWhitelist(destAccount); // whitelist the account to make it usable
            await instance.transfer(destAccount, 100); // transfer 100 tokens to dest account
        })
    
        it('should throw when recipient is not whitelisted', async () => {
            const destAccount = accounts[1]; // address to which the token will be sent
            try {
                await instance.transfer(destAccount, 100);
                assert.fail()
            } catch (err) {
                assert.isOk(/revert/.test(err.message), "There was no REVERT error thrown")
            }
        })
    })
})