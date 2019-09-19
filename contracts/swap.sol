pragma solidity ^0.4.24;

contract Swap{
    
    struct Agreement{
        uint256 swapAmount;//5,000,000
        uint256 libor;//5.56%
        uint256 FTSE;//10%
        uint256 swapTenor;//180 days or 10 seconds or whatever.
    }
    
    struct Party{
        address addr;
        string partyName;
        uint256 partyBalance;
        string jurisdiction;
        bytes32 IPFShash;
        Agreement agreement;
    }
    
    mapping(address => Party) parties;
    address address_A;
    address address_B;
    
    
    constructor(
        
        address address_A_,
        address address_B_,
        string partyAname,
        string partyBname,
        uint256 swap_amount,
        uint256 LIBOR_amount,
        uint256 FTSE_amount,
        string jurisdiction_
        
        ) public {
            
        address party_A_addr = address_A_;
        address party_B_addr = address_B_;
        string memory party_A_name = partyAname;
        string memory party_B_name = partyBname;
        
        Agreement memory agrement_A;
        agrement_A.swapAmount = swap_amount; 
        agrement_A.libor = LIBOR_amount;
        agrement_A.swapTenor = 180;
        
        Party memory A;
        A.addr = party_A_addr;
        A.partyName = party_A_name;
        A.partyBalance = 0;
        A.jurisdiction = jurisdiction_;
        A.agreement = agrement_A;
        A.IPFShash = 0x861ce807bbdcba4de666e2552ed170ea92adc4f35da9e97723e3a6ee374458d2;
        
        Agreement memory agrement_B;
        agrement_B.swapAmount = swap_amount; 
        agrement_B.FTSE = FTSE_amount; //10 * 100 == 10%
        agrement_B.swapTenor = 180;
        
        Party memory B;
        B.addr = party_B_addr;
        B.partyName = party_B_name;
        B.partyBalance = 0;
        B.jurisdiction = jurisdiction_;
        B.agreement = agrement_B;
        B.IPFShash = 0x861ce807bbdcba4de666e2552ed170ea92adc4f35da9e97723e3a6ee374458d2;
        
        parties[party_A_addr] = A;
        parties[party_B_addr] = B;
        
        address_A = party_A_addr;
        address_B = party_B_addr;
        
        
    }
    
    function main(uint256 _FTSE, uint256 _LIBOR, bytes32 _IPFShash) public returns(bool) {
        require(clearValues());
        require( keccak256(parties[address_A].jurisdiction) == keccak256(parties[address_B].jurisdiction));
        require(_IPFShash == parties[address_A].IPFShash);
        require(_IPFShash == parties[address_B].IPFShash);
        payments(_FTSE, _LIBOR);

    }

    function payments(uint256 _FTSE , uint256 _LIBOR) private  returns(uint){
        
        if(_FTSE > parties[address_B].agreement.FTSE){
            
            uint256 FTSEChange = _FTSE - parties[address_B].agreement.FTSE;
            
            uint256 swap_amount = parties[address_A].agreement.swapAmount;
            uint256 swap_tenor = parties[address_A].agreement.swapTenor;
            
            
            uint256 amountToPayB = (swap_amount * _LIBOR / 10000) * swap_tenor/360;
            uint256 amountToPayA = (swap_amount * FTSEChange / 10000);
            
            if (amountToPayA > amountToPayB){
                
                uint256 finalAmount = amountToPayA - amountToPayB;
                parties[address_A].partyBalance += finalAmount;
                return parties[address_A].partyBalance;
                
              //finalAmount i shtohet bilancit te party A
            
                
            }else if(amountToPayA < amountToPayB){
                
                finalAmount = amountToPayB - amountToPayA;
                parties[address_B].partyBalance += finalAmount;
                return parties[address_B].partyBalance;
                //finalAmount i shtohet bilancit te party B
        
            }else if(amountToPayA == amountToPayB){
                
                return 0;
                
            }
            
        }else{
            
            FTSEChange = parties[address_B].agreement.FTSE - _FTSE;
            
            swap_amount = parties[address_A].agreement.swapAmount;
            swap_tenor = parties[address_A].agreement.swapTenor;
            
            amountToPayB = ((swap_amount * _LIBOR / 10000) * swap_tenor/360) + 
                (swap_amount * FTSEChange / 10000);
                
            //in this case B does not need to pay A     
            //B bilance += amountToPayB
            parties[address_B].partyBalance += amountToPayB;
            return parties[address_B].partyBalance;
            
        }
        
    }
    
    function getBalances() public view returns(uint256,uint256){
        return (
            parties[address_A].partyBalance, 
            parties[address_B].partyBalance
            );
    }
    
    function clearValues() private returns(bool){
        parties[address_A].partyBalance = 0;
        parties[address_B].partyBalance = 0;
        return true;
    }
    
 
}