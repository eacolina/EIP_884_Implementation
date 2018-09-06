pragma solidity ^0.4.24;

// import "./ERC20NoTransfer.sol";

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./IdentityRegistry.sol";
import "openzeppelin-solidity/contracts/token/ERC20/BasicToken.sol";


contract StockToken is BasicToken {

    string public symbol;
    string public name;
    string public byLawsHash;
    address public owner;
    uint public decimals;
    bool public isPrivateCompany = true;
    mapping(address => uint) tokenOwnersIndex;
    address[] public tokenOwners;
    IdentityRegistry public platformWhitelist;
    
    constructor(string _symbol, string _name, uint _supply, string hash, address _registry) public  {
        symbol = _symbol;
        name = _name;
        totalSupply_ = _supply;
        byLawsHash = hash;
        owner = msg.sender;
        balances[msg.sender] = _supply;
        platformWhitelist = IdentityRegistry(_registry);
        tokenOwners.push(0);
        uint index = tokenOwners.push(msg.sender);
        tokenOwnersIndex[msg.sender] = index - 1;
    }

    modifier onlyIfWhitelisted(address _address) { // modifier to restrict access only to whitelisted accounts
        require(platformWhitelist.isWhitelisted(_address));
        _;
    }

    function transfer(address _to, uint256 _value) onlyIfWhitelisted(_to) public returns (bool){
        uint index = 0;
        if(tokenOwnersIndex[_to] == 0) {
            index = tokenOwners.push(_to) - 1;
            tokenOwnersIndex[_to] = index;
        }
        super.transfer(_to, _value);
        if(balanceOf(msg.sender) == 0){
            removeTokenOwner(msg.sender);
        }
    }

    function togglePrivateCompany() public {
        isPrivateCompany = !isPrivateCompany;
    }

    // Return the number of shareholders in the company
    function ownersCount() public view returns(uint){
        return tokenOwners.length - 1;
    }
    
    // Return an array with all the token owners
    function getTokenOwners() public view returns(address[]){
        return tokenOwners;
    }

    // Removes entry from array at index and resizes the array appropriatly
    function removeFromTokenOwnersArray(uint index) internal {
        address lastElement = tokenOwners[tokenOwners.length - 1];
        tokenOwners[index] = lastElement;
        tokenOwners.length--;
    }

    // Removes a token owner from the list of shareholders
    function removeTokenOwner(address holder) internal {
        uint i = tokenOwnersIndex[holder];
        removeFromTokenOwnersArray(i);
        tokenOwnersIndex[holder] = 0;
    }
    
}