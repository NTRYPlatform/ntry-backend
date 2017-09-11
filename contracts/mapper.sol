pragma solidity ^0.4.13;

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

    function mapAddress(bytes16 secondary)
        secondaryAddressMustBeUnique(secondary) {
        // If primary address is already in use, throw error
        if (primaryToSecondary[msg.sender].inUse) revert();
        // If there is no mapping, this does nothing
        secondaryInUse[primaryToSecondary[msg.sender].uid] = false;

        primaryToSecondary[msg.sender] = secondaryAddress(secondary, true);
        secondaryInUse[secondary] = true;

        AddressMapped(msg.sender, secondary);
    }
}