contract NotaryMapper {

   struct secondaryAddress {
     bytes16 uid;
     bool inUse;
   }

    event AddressMapped(address primary, bytes16 secondary);
    event Error(uint code, address sender);

    mapping (address => secondaryAddress) public primaryToSecondary;
    mapping (bytes16 => bool) public secondaryInUse;

    modifier secondaryAddressMustBeUnique(bytes16 secondary) {
        if(secondaryInUse[secondary]) {
            Error(1, msg.sender);
            revert();
        }
        _;
    }

    function NotaryMapper(){
        owner = msg.sender;
        tokenContract = 0x00;
    }

    function mapAddress(bytes16 secondary)
        secondaryAddressMustBeUnique(secondary) {
        require(tokenContract!=0x00);
        // If primary address is already in use, throw error
        if (primaryToSecondary[msg.sender].inUse) revert();
        // If there is no mapping, this does nothing
        secondaryInUse[primaryToSecondary[msg.sender].uid] = false;

        primaryToSecondary[msg.sender] = secondaryAddress(secondary, true);
        secondaryInUse[secondary] = true;
        
        if (!NotaryPlatformToken(tokenContract).faucet(msg.sender)){
            revert();
        }

        AddressMapped(msg.sender, secondary);
    }
    
    //*********** Only for Test App**************// 
    address public owner;
    address public tokenContract;
    function setTokenAddress(address _newAddress) external{
        require(msg.sender == owner);
        tokenContract = _newAddress;
    }
}

contract NotaryPlatformToken{
    function faucet(address _to) returns(bool);
}
