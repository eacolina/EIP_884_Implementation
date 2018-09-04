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
    IdentityRegistry public platformWhitelist;
    
    constructor(string _symbol, string _name, uint _supply, string hash, address _registry) public  {
        symbol = _symbol;
        name = _name;
        totalSupply_ = _supply;
        byLawsHash = hash;
        owner = msg.sender;
        balances[msg.sender] = _supply;
        platformWhitelist = IdentityRegistry(_registry);
    }

    modifier onlyIfWhitelisted(address _address) { // modifier to restrict access only to whitelisted accounts
        require(platformWhitelist.isWhitelisted(_address));
        _;
    }

    function transfer(address _to, uint256 _value) onlyIfWhitelisted(_to) public returns (bool){
        super.transfer(_to, _value);
    }

    function togglePrivateCompany() public {
        isPrivateCompany = !isPrivateCompany;
    }
    
}