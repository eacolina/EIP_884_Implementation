pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/token/ERC20/BasicToken.sol";

contract StockToken is BasicToken {

    string public symbol;
    string public name;
    string public by_LawsHash;
    address public owner;
    mapping(address => bool) public authorizedShareholders;

    constructor (string _symbol,string _name, uint _supply, string hash) public  {
        symbol = _symbol;
        name = _name;
        totalSupply_ = _supply;
        by_LawsHash = hash;
        owner = msg.sender;
        addAddressToWhitelist(msg.sender);
        balances[msg.sender] = _supply;
    }

    modifier onlyIfWhitlisted (address _address) { // modifier to restrict access only to whitelisted accounts
        require(isWhitelisted(_address));
        _;
    }
    modifier onlyOwner() { // modifier to restrict access only the owner of the contract
        require(msg.sender == owner);
        _;
    }

    function isWhitelisted (address _address) public view returns (bool){ // function to check if address is whitelisted
        return authorizedShareholders[_address];
    }

    function addAddressToWhitelist (address _address) onlyOwner public returns (bool) { // add an address to the authorizedShareholders mapping
        require(!isWhitelisted(_address));
        authorizedShareholders[_address] = true;
    }
}