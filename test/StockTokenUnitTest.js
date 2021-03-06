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
    const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000';
    
    beforeEach('Create a new contract instance', async () => {
        platformWhitelist = await IdentityRegistry.new();
        await platformWhitelist.addIdentity(OWNER,"OWNER");
        await platformWhitelist.addIdentity(accounts[2], "Account 2"); // in case a whitelisted account is needed
        instance = await StockToken.new('ROKK','Rokk3r Crowdbuild',1000000, '022841754bd3d55d221fdb46a178cee5e223937eebaccc56efc415e7e63823ca',platformWhitelist.address);
        await instance.addAddressToWhitelist(whitelistedAccount);
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

        it('should create a StockToken  and assign tokens', async() => {
            let account = accounts[0];
            const balance = await instance.balanceOf(account); // get balance
            const totalSupply = await instance.totalSupply(); // get total supply of coins
            assert.isTrue(balance.eq(totalSupply), "The owner of the contract doesn't own all tokens"); //Owner's balance should be equal to total supply
        })
        it('should add the owner address to the shareholders array at index 1 and index 0 should be 0x0', async() => {
            let account = accounts[0];
            let ownersCount = await instance.ownersCount();
            let zeroIndexAddress = await instance.tokenOwners(0);
            let onlyOwner = await instance.tokenOwners(1);
            assert.equal(zeroIndexAddress,ZERO_ADDRESS,"This address should always be the 0 address");
            assert.equal(onlyOwner,OWNER,"The address in the token owners array is not the OWNER");
            assert.isTrue(ownersCount.eq(1),"There should only be one owner at this point")
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
            const destAccount = accounts[3];
            try{
                await instance.transfer(destAccount, 100);
                assert.fail();
            } catch(err) {
                assert.isOk(/revert/.test(err.message), "There was no REVERT error thrown")
            }
        })
        it('should check if destAccount is in tokenHolders array and add it if needed and save it\'s index in the tokenHoldersIndex mapping', async() =>{
            const destAccount = whitelistedAccount; // address to which the token will be sent
            await instance.transfer(destAccount, 100);
            let tokenOwners = await instance.getTokenOwners();
            let length = await instance.ownersCount();
            assert.isTrue(length.eq(2), "The length was not updated correctly");
            assert.equal(tokenOwners[2],destAccount,"The account wasn\'t added correctly");
        })
        it('it should only add to tokenOwners if it\'s balance is not 0', async() => {
            const destAccount = whitelistedAccount; // address to which the token will be sent
            await instance.transfer(destAccount, 100);
            await instance.transfer(destAccount,100);
            let owners = await instance.ownersCount();
            assert.equal(owners,2);
        })
        it("it should remove destAccount from tokenHolders array if new balance is 0 and make it's index -1 in the mapping",async() =>{
            const destAccount = whitelistedAccount; // address to which the token will be sent
            const totalSupply = 1000000;
            let tokenOwners = await instance.getTokenOwners();
            await instance.transfer(destAccount, totalSupply);
            tokenOwners = await instance.getTokenOwners();
            let balance = await instance.balanceOf(OWNER);
            assert.equal(tokenOwners.length,2,"The account wasn\'t removed correctly")
            assert.equal(tokenOwners[1],destAccount,"Not the right account")
        })
    })

    describe('togglePrivateCompany', async () => {
        it('should change the isPrivateCompany to its opposite value', async () => {
            assert.isTrue(await instance.isPrivateCompany(), 'The company is public');
            await instance.togglePrivateCompany()
            assert.isFalse(await instance.isPrivateCompany(), "Company is still private");
        })

        it('should emit an changedCompanyStatus event', async() => {
            const tx = await instance.togglePrivateCompany();
            const event = tx.logs[0];
            assert.equal(event.event, "ChangedCompanyStatus", "changedCompanyStatus was not the sent event"); // check if the right event was fired
            assert.equal(event.args.authorizedBy, OWNER, "The function was called from owner but log is not")
            assert.isFalse(event.args.newStatus, "The account removed doesn't match") // testing for false because isPrivateCompany is true by default
        })
    })

    describe('getTokenOwners', async() =>{
        it('should return an array with all the owner addresses',async ()=>{
            let array = await instance.getTokenOwners()
            assert.equal(array.length,2,"Array should contain only one address")
            assert.equal(array[1],OWNER,"This address should be the address of the token owner")
        })
    })

})
