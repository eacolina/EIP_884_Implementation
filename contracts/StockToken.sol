pragma solidity ^0.4.24;

// import "./ERC20NoTransfer.sol";
import "./ERC20NoTransfer.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./Whitelistable.sol";
contract StockToken is ERC20NoTransfer {

    string public symbol;
    string public name;
    string public byLawsHash;
    address public owner;
    uint public decimals;
    bool public isPrivateCompany = true;
    Whitelistable public platformWhitelist;
    
    constructor(string _symbol, string _name, uint _supply, string hash, address _whitelist) public  {
        symbol = _symbol;
        name = _name;
        totalSupply_ = _supply;
        byLawsHash = hash;
        owner = msg.sender;
        balances[msg.sender] = _supply;
        platformWhitelist = Whitelistable(_whitelist);
    }

    modifier onlyIfWhitelisted(address _address) { // modifier to restrict access only to whitelisted accounts
        require(platformWhitelist.isWhitelisted(_address));
        _;
    }

    modifier onlyOwner() { // modifier to restrict access only the owner of the contract
        require(msg.sender == owner);
        _;
    }

    function transfer(address _to, uint256 _value) onlyIfWhitelisted(_to) public returns (bool){
        return transferAfterWhitelist(_to, _value);
    }

    function togglePrivateCompany() public {
        isPrivateCompany = !isPrivateCompany;
    }
}