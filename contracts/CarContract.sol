pragma solidity ^0.4.16;

import './Pausable.sol';
import './ReentrancyGuard.sol';

contract NotaryPlatformToken{
    function isTokenContract() returns (bool);
    function transferFrom(address _from, address _to, uint256 _value) returns (bool);
}

contract CarContracts is Ownable, ReentrancyGuard{

    struct Deal{
        address seller;
        address buyer;
        bytes32 hash;
    }

    bytes[] private index;

    mapping (bytes => Deal) private carDeals;
    
    // Only for testnet
    mapping (address => bytes) public seller;
    mapping (address => bytes) public buyer;
    
    uint256 private contractFee;
    uint256 private feeSum;
    
    address private notaryApp;
    NotaryPlatformToken private notaryToken;
    
    event LogDeal(address indexed sellerAddress, address indexed buyerAddress, bytes dealID, uint256 dealfee, bytes32 dealHash);
    event AppAccountUpdated(address indexed oldAddress,address indexed newAddress);
    event TokenAddressUpdated(address indexed oldAddress,address indexed newAddress);
    event ContractFeeUpdated(uint256 indexed oldValue,uint256 indexed newValue);
    event Error();
    
    function CarContracts(address _owner,address _notaryApp, address _notaryToken) Ownable(_owner){
        notaryApp = _notaryApp;
        notaryToken = NotaryPlatformToken(_notaryToken);
        require(notaryToken.isTokenContract());
        contractFee = 10;
        
    }
    
    function ntryApp() external constant returns (address){ return notaryApp;}
    
    function ntryToken() external constant returns (address){ return notaryToken;}
    
    function fee() external constant returns (uint256){ return contractFee;}
    
    function totalEarned() external constant returns (uint256){
        return feeSum;
    }
    
    function totalContracts() external constant returns (uint256){
        return index.length;
    }
    
    function updateFee(uint256 _value) external onlyOwner returns (bool){
        require(_value > 0);
        ContractFeeUpdated(contractFee,_value);
        contractFee = _value;
        return true;
    }
    
    function updateAppAccount(address _address) external onlyOwner returns (bool){
        require(_address != 0x00);
        AppAccountUpdated(notaryApp,_address);
        notaryApp = _address;
        return true;
    }
    
    function updateTokenAddress(address _address) external onlyOwner returns (bool){
        require(_address != 0x00);
        TokenAddressUpdated(notaryToken,_address);
        notaryToken  = NotaryPlatformToken(_address);
        if(!notaryToken.isTokenContract()){
            revert();
        } 
        return true;
    }
    
    /**
     * @param cid unique contractID calculated at server
     * @param _seller Car seller public address
     * @param _buyer  Car buyer public address
     * @param _hash Hash of contract attributes
     * @notice Arguments sequence for hash function
     * bytes cid, address _seller, address _buyer,
     * uint256 year,bytes make, bytes model, bytes vin,
     * bytes carType, bytes color, bytes engine_no,
     * uint8 mileage, uint256 totalPrice,
     * uint256 downPayment, uint256 remainingPayment,
     * uint256 remainingPaymentDate
     */
    function carDeal(bytes cid, address _seller, address _buyer,bytes32 _hash) nonReentrant() onlyApp() returns (bool){
        require(carDeals[cid].seller == 0x00);
    
        if(!notaryToken.transferFrom(_buyer, owner,contractFee)){
            Error();
            return false;
        }
        
        feeSum += contractFee;
        
        
        carDeals[cid].seller = _seller;
        carDeals[cid].buyer = _buyer;
        carDeals[cid].hash = _hash;
        
        index.push(cid);
        seller[_seller] = cid;
        seller[_buyer] = cid;
        
        LogDeal(_seller,_buyer,cid,contractFee,_hash);
        
        return true;
    }
    
    function getDeal(bytes cid) external constant returns(address,address,bytes32){
        require(carDeals[cid].seller != 0x00);
        return(carDeals[cid].seller,carDeals[cid].buyer,carDeals[cid].hash);
    }
    
    function getID(uint256 _index) external constant returns(bytes){
        return(index[_index]);
    }
    
    modifier onlyApp(){
        require(msg.sender == notaryApp);
        _;
    }
}