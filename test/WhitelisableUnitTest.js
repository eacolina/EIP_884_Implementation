const Whitelistable = artifacts.require('./Whitelistable.sol');

contract('Whitelistable', async (accounts) =>{
    let instance;
    const OWNER = accounts[0];

    beforeEach('Create a new contract instance', async () => {
        instance = await Whitelistable.new();
    })

    describe('addAddressToWhitelist', () => {
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
            await instance.addAddressToWhitelist(account);
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

    describe('removeAddressFromWhitelist', async () => {
        it('should remove the given address from the whitelist mapping', async () => {
            let account = accounts[1]
            await instance.addAddressToWhitelist(account)
            const isWhitelistedBefore = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isTrue(isWhitelistedBefore, "The address was not whitelisted"); // This address should be whitelisted already
            await instance.removeAddressFromWhitelist(account);
            const isWhitelistedAfter = await instance.isWhitelisted(account); // Check if address is in whitelist
            assert.isFalse(isWhitelistedAfter, "The address was not whitelisted"); // This address should be whitelisted already
        })
        it('should emit AddressRemovedFromWhitelist with correct parameters', async() => {
            let account = accounts[1];
            await instance.addAddressToWhitelist(account)
            const tx = await instance.removeAddressFromWhitelist(account)
            const event = tx.logs[0];
            assert.equal(event.event, "AddressRemovedFromWhitelist", "AddressRemovedToWhitelist was not the sent event"); // check if the right event was fired
            assert.equal(event.args.AuthorizedBy, OWNER, "The function was called from owner but log is not")
            assert.equal(event.args.AddressRemoved, account, "The account removed doesn't match")
        })
        it("should throw when non-owner tries to remove from whitelist", async () => {
            let account = accounts[1]; // non owner account
            await instance.addAddressToWhitelist(account)
            // Try-catch to check if error thrown
            try {
                await instance.removeAddressFromWhitelist(account,{from:account}) // try to whitelist account from non owner
                assert.fail(); // if above line doesn't throw, throw a fail
            } catch (err) {
                assert.ok(/revert/.test(err.message),"There was no REVERT error") // check what king of error was thrown
            }
        })
    })
})