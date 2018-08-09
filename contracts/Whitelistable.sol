pragma solidity ^0.4.24;

contract Whitelistable {
    mapping (address => bool) public whitelist;
    address public owner;
    
    event AddressAddedToWhitelist(address indexed AuthorizedBy, address indexed AddressAdded);
    event AddressRemovedFromWhitelist(address indexed AuthorizedBy, address indexed AddressRemoved);

    modifier onlyOwner() { // modifier to restrict access only the owner of the contract
        require(msg.sender == owner);
        _;
    }

    constructor() public{
        owner = msg.sender;
    }
    
    function isWhitelisted(address _address) public view returns(bool){ // function to check if address is whitelisted
        return whitelist[_address];
    }

    function addAddressToWhitelist(address _address) onlyOwner public{ // add an address to the authorized mapping
        require(!isWhitelisted(_address));
        whitelist[_address] = true;
        emit AddressAddedToWhitelist(msg.sender, _address);
    }

    function removeAddressFromWhitelist(address _address) onlyOwner public{
        require(isWhitelisted(_address)); // check if address is whitelisted
        whitelist[_address] = false;
        emit AddressRemovedFromWhitelist(msg.sender, _address);
    }
}