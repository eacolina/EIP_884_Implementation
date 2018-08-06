pragma solidity ^0.4.24;

// import "./ERC20NoTransfer.sol";
import "./ERC20NoTransfer.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";

contract StockToken is ERC20NoTransfer {

    string public symbol;
    string public name;
    string public byLawsHash;
    address public owner;
    uint public decimals;
    mapping(address => bool) public authorizedShareholders;

    event AddressAddedToWhitelist(address indexed AuthorizedBy, address indexed AddressAdded);

    constructor (string _symbol,string _name, uint _supply, string hash) public  {
        symbol = _symbol;
        name = _name;
        totalSupply_ = _supply;
        byLawsHash = hash;
        owner = msg.sender;
        addAddressToWhitelist(msg.sender);
        balances[msg.sender] = _supply;
    }

    modifier onlyIfWhitelisted (address _address) { // modifier to restrict access only to whitelisted accounts
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
        emit AddressAddedToWhitelist(msg.sender, _address);
    }

    function transfer(address _to, uint256 _value) onlyIfWhitelisted(_to) public returns (bool){
        return transferAfterWhitelist(_to, _value);
    }
}