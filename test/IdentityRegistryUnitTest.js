var IdentityRegistry = artifacts.require('./IdentityRegistry');

contract('IdentityRegistry', async(accounts) => {
    const owner = accounts[0];
    const normalAccount1 = accounts[1];
    let instance
    
    beforeEach('Create new contract instance', async () =>{
        instance = await IdentityRegistry.new();
    })

    describe('constructor', () => {
        it('should create an instance with correct owner', async () => {
            assert.equal(await instance.owner(), owner, "Owner was set incorrectly.")
        })
    })

    describe('addIdentity', () =>{
        it('should add the address to the map with it\'s corresponding hash if identity not in map and add the address to the whitelist', async () => {
            const hash = "SUPERHASH";
            await instance.addIdentity(normalAccount1,hash);
            const retrivedHash = await instance.identityMap(normalAccount1);
            const isWhitelisted = await instance.isWhitelisted(normalAccount1);
            assert.equal(retrivedHash,hash, "The hash was not added correctly.");
            assert.isTrue(isWhitelisted, "The address was not whitelisted");
        })
        it('should emit IdentityAdded event under normal conditions ', async () => {
            const hash = "SUPERHASH";
            const tx = await instance.addIdentity(normalAccount1,hash);
            const event = tx.logs[1].args; // logs[1] becuase logs[0] is the AddressAddedToWhitelist is [0]
            const authorizedBy = event.authorizedBy;
            const addressAdded = event.addressAdded;
            const identityHash = event.identityHash;
            assert.equal(authorizedBy, owner, "Authorizing address doesn't match");
            assert.equal(addressAdded, normalAccount1, "Address added is incorrect");
            assert.equal(identityHash, hash, "Identity hash doesn't match");
        })
        it('should throw if address already has an identity', async () => {
            const hash = "SUPERHASH";
            await instance.addIdentity(normalAccount1,hash);
            try {
                await instance.addIdentity(normalAccount1, hash);
                assert.fail();
            } catch (e) {
                assert.isOk(/revert/.test(e.message), "No REVERT error was thrown");

            }
        })
        it('should throw if non owner is calling',async () => {
            const hash = "SUPERHASH";
            await instance.addIdentity(normalAccount1,hash);
            try {
                await instance.addIdentity(normalAccount1, hash);
                assert.fail();
            } catch (e) {
                assert.isOk(/revert/.test(e.message), "No REVERT error was thrown");

            }
        })
    })
    
    describe('updateIdentity', () => {
        const hash = "HASH"; 
        const newHash = "SUPERHASH";
        it('should update the identity hash under normal conditions', async () => {
            await instance.addIdentity(normalAccount1, hash);
            await instance.updateIdentity(normalAccount1, newHash);
            const retrivedHash = await instance.identityMap(normalAccount1);
            assert.equal(retrivedHash, newHash);
        })
        it('should throw if the address doesn\'t have an identity yet', async () =>{
            try {
                await instance.updateIdentity(normalAccount1, hash);
                assert.fail();
            } catch (e) {
                assert.isOk(/revert/.test(e.message), "No REVERT error was thrown");

            }
        })
        it('should emit IdentityAdded event under normal conditions', async () => {
            await instance.addIdentity(normalAccount1, hash);
            const tx = await instance.updateIdentity(normalAccount1, newHash);
            const event = tx.logs[0].args;
            const addressUpdated = event.updatedAddress;
            const previousHash = event.previousHash;
            const updatedHash = event.newHash;
            const authorizedBy = event.authorizedBy;
            assert.equal(addressUpdated, normalAccount1, "The wrong account was updated");
            assert.equal(previousHash, hash, "The previous hash is incorrect");
            assert.equal(updatedHash, newHash, "The new hash is incorrect");
            assert.equal(authorizedBy, owner, "The updated wasn't made by the owner");
        })

    })

})