Feature: Swap
 SCHEDULE to the 2002 Master Agreement

Scenario: Addresses of A and B
   Given that the address of "Party A" is "0xca35b7d915458ef540ade6068dfe2f44e8fa733c" and address of "Party B" is "0x14723a09acff6d2a60dcdf7aa4aff308fddc160c"
   And the return statement contains both addresses
   Then the contract should confirm that the addresses taking part in the transaction are the correct ones.

Scenario: IPFS UUID
   Given a "0x861ce807bbdcba4de666e2552ed170ea92adc4f35da9e97723e3a6ee374458d2" specified as an IPFS hash for a given contract
   Then the contract should have a valid IPFS hash

Scenario: comparing jurisdiction variables
   Given jurisdiction variable "USA" of "Party A"
   And jurisdiction variable "USA" of "Party B"
   Then jurisdiction variables should be equal

Scenario: Swap interest
   Given that "Party A" with address "0xca35b7d915458ef540ade6068dfe2f44e8fa733c" and "Party B" with address "0x14723a09acff6d2a60dcdf7aa4aff308fddc160c" have a swap contract
   And "Party A" swaps at LIBOR and interest rate of "Party B" is 0.03%
   And that swap amount is 5000000$
   And swap tenor is 180 days
   And contract started from timestamp 1544018183
   When current time has not passed the swap tenor
   Then we wait for the time to pass

Scenario: Swap interest
   Given that "Party A" with address "0xca35b7d915458ef540ade6068dfe2f44e8fa733c" and "Party B" with address "0x14723a09acff6d2a60dcdf7aa4aff308fddc160c" have a swap contract
   And "Party A" swaps at LIBOR and interest rate of "Party B" is 0.03%
   And that swap amount is 5000000$
   And swap tenor is 180 days
   And contract started from timestamp 1544018183
   When current time has passed the swap tenor
   But interest rate of "Party B" has appreciated for 1.0%
   And LIBOR rate is 4.0%
   Then after checking the difference is paid

Scenario: Swap interest
   Given that "Party A" with address "0xca35b7d915458ef540ade6068dfe2f44e8fa733c" and "Party B" with address "0x14723a09acff6d2a60dcdf7aa4aff308fddc160c" have a swap contract
   And "Party A" swaps at LIBOR and interest rate of "Party B" is 0.05%
   And that swap amount is 5000000$
   And swap tenor is 180 days
   And contract started from timestamp 1544018183
   When current time has passed the swap tenor
   And interest rate of "Party B" has not appreciated
   And LIBOR rate is 4.5%
   Then after checking the difference is paid
