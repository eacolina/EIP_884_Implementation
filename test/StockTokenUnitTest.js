/*
IMPORTANT NOTE: These tests do not cover the basic functionality of the ERC20 standard as I'm using the implementation of Open Zeppelin which has been thoroughly audited. 
*/

const StockToken = artifacts.require('./StockToken.sol');
const Whitlistable = artifacts.require('./Whitelistable.sol');
const IdentityRegistry = artifacts.require('./IdentityRegistry');

contract('StockToken', async(accounts) => {
    let instance
    let platformWhitelist
    const OWNER = accounts[0];
    const whitelistedAccount = accounts[2];
    
    beforeEach('Create a new contract instance', async () => {
        platformWhitelist = await IdentityRegistry.new();
        await platformWhitelist.addIdentity(accounts[2], "Account 2"); // in case a whitelisted account is needed
        instance = await StockToken.new('ROKK','Rokk3r Crowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca',platformWhitelist.address);
    })

    describe('constructor', () => {
        it('should create a StockToken with correct parameters and owner', async() => {
            let name = await instance.name(); // get name of token
            let symbol = await instance.symbol(); // get symbol of token
            let supply = await instance.totalSupply(); // get supply of token 
            let hash = await instance.byLawsHash(); // get hash of token
            let owner = await instance.owner();
            let decimals = await instance.decimals();
            let platformWhitelistAddress = await instance.platformWhitelist();
            assert.equal(symbol, "ROKK","Symbol doesn't match");
            assert.equal(name, "Rokk3r Crowdbuild","Names doesn't match");
            assert.equal(supply, 1000000, "Supply doesn't match");
            assert.equal(hash, "022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca", "hash doesn't match");
            assert.equal(owner, accounts[0],"The owner is incorrect");
            assert.equal(decimals.toNumber(), 0, "The decimals are not 0.");
            assert.equal(platformWhitelistAddress, platformWhitelist.address, "The whitelist pointer is incorrect");
        })

        it('should create a StockToken and whitelist the owner\'s address', async() => {
            let account = accounts[0];
            const balance = await instance.balanceOf(account); // get balance
            const totalSupply = await instance.totalSupply(); // get total supply of coins
            assert.isTrue(balance.eq(totalSupply), "The owner of the contract doesn't own all tokens"); //Owner's balance should be equal to total supply
        })
    })
    
    describe('transfer', () => {
        it('should allow for transfer of tokens to a platform whitelisted address when isPrivateCompany is false', async () => {
            const destAccount = whitelistedAccount; // address to which the token will be sent
            const initDestBalance = await instance.balanceOf(destAccount);
            const initOrigBalance = await instance.balanceOf(OWNER);
            await instance.togglePrivateCompany(); // isPrivate needs to be false for this test
            await instance.transfer(destAccount, 100); // transfer 100 tokens to dest account
            const postTransferBalance = await instance.balanceOf(destAccount);
            assert.isTrue(postTransferBalance.eq(100),"Balances dont add up")
        })
    
        it('should throw when recipient is not whitelisted in IdentityRegistry', async () => {
            const destAccount = accounts[1]; // address to which the token will be sent
            try {
                await instance.transfer(destAccount, 100);
                assert.fail()
            } catch (err) {
                assert.isOk(/revert/.test(err.message), "There was no REVERT error thrown")
            }
        })
        it('if isPrivateComapny, should throw when recipient is whitelisted in IdentityRegistry but not in the authorized mapping(token whitelist) ', async () => {
            const destAccount = accounts[2];
            try{
                await instance.transfer(destAccount, 100);
                assert.fail();
            } catch(err) {
                assert.isOk(/revert/.test(err.message), "There was no REVERT error thrown")
            }
        })
    })

    describe('togglePrivateCompany', async () => {
        it('should change the isPrivateCompany to its opposite value', async () => {
            assert.isTrue(await instance.isPrivateCompany(), 'The company is public');
            await instance.togglePrivateCompany()
            assert.isFalse(await instance.isPrivateCompany(), "Company is still private");
        })
    })
})