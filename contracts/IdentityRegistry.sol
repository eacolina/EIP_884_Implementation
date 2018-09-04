pragma solidity ^0.4.24;

import "./Whitelistable.sol";

contract IdentityRegistry is Whitelistable {
    mapping(address => string) public identityMap; // maps an address to an idenity hash

    event IdentityAdded(address indexed addressAdded, string identityHash, address indexed authorizedBy);
    event IdentityUpdated(address indexed updatedAddress, string previousHash, string newHash, address indexed authorizedBy);

    constructor() public Whitelistable(){ // empty constructor used to call the whitelistable constructor
    }

    function addIdentity(address _address, string hash) onlyOwner() public {
        bytes memory value = bytes(identityMap[_address]);
        require(value.length == 0, "This identity was registered already");
        identityMap[_address] = hash;
        addAddressToWhitelist(_address);
        emit IdentityAdded(_address, hash, msg.sender);
    }

    function updateIdentity(address updatedAddress, string newHash) onlyOwner() public {
        bytes memory previousHash = bytes(identityMap[updatedAddress]);
        require(previousHash.length != 0);
        identityMap[updatedAddress] = newHash;
        emit IdentityUpdated(updatedAddress, string(previousHash), newHash, msg.sender);
    }
}
