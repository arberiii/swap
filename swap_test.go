package main

import (
	"errors"
	"log"
	"math/big"
	"os"
	"strings"
	"swap/contracts"

	"github.com/DATA-DOG/godog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	partA             string
	partB             string
	addressPartA      string
	addressPartB      string
	jurisdictionA     string
	jurisdictionB     string
	firstJurisdiction = true
	rateA             int
	rateAAfter        int
	rateBAfter        int
	rateB             int
	swapA             int
	swapB             int
	swapTenor         int
	contractTime      int
	libor             = 59700
	ipfsHash          string
	message           string
)

func thatTheAddressOfIsAndAddressOfIs(arg1, arg2, arg3, arg4 string) error {
	partA = arg1
	addressPartA = arg2
	partB = arg3
	addressPartB = arg4
	return nil
}

func theReturnStatementContainsBothAddresses() error {
	return nil
}

func theContractShouldConfirmThatTheAddressesTakingPartInTheTransactionAreTheCorrectOnes() error {
	if realAddressPartA != addressPartA || realAddressPartB != addressPartB {
		return errors.New("the address do not match")
	}
	return nil
}

// sceario with ipfs

func aSpecifiedAsAnIPFSHashForAGivenContract(arg1 string) error {
	ipfsHash = arg1
	return nil
}

func theContractShouldHaveAValidIPFSHash() error {
	if ipfsHash != realIPFS {
		return errors.New("not valid ipfs hash")
	}
	return nil
}

// scenario with jurisdiction

func jurisdictionVariableOf(arg1, arg2 string) error {
	if firstJurisdiction {
		jurisdictionA = arg1
		firstJurisdiction = false
	} else {
		jurisdictionB = arg1
		firstJurisdiction = true
	}
	return nil
}

func jurisdictionVariablesShouldBeEqual() error {
	if jurisdictionA != jurisdictionB {
		return errors.New("jurisdictions do not match")
	}
	return nil
}

// scenario when time has not come yet
func thatWithAddressAndWithAddressHaveASwapContract(arg1, arg2, arg3, arg4 string) error {
	partA = arg1
	addressPartA = arg2
	partB = arg3
	addressPartB = arg4
	if realAddressPartA != addressPartA || realAddressPartB != addressPartB {
		return errors.New("the address do not match")
	}
	return nil
}

func interestRateOfIsAndInterestRateOfIs(arg1 string, arg2 float64, arg3 string, arg4 float64) error {
	rateA = int(arg2 * 10000)
	rateB = int(arg4 * 10000)
	return nil
}

func swapsAtLIBORAndInterestRateOfIs(arg1, arg2 string, arg3 float64) error {
	rateB = int(arg3 * 10000)
	return nil
}

func thatSwapOfIsAndSwapOfIs(arg1 string, arg2 int, arg3 string, arg4 int) error {
	swapA = arg2
	swapB = arg4
	return nil
}

func swapTenorIsDays(arg1 int) error {
	swapTenor = arg1
	return nil
}

func contractStartedFromTimestamp(arg1 int) error {
	contractTime = arg1
	return nil
}

func currentTimeHasNotPassedTheSwapTenor() error {
	return nil
}

func weWaitForTheTimeToPass() error {
	return nil
}

// scenario when both rate A and rate B have appreciated
func currentTimeHasPassedTheSwapTenor() error {
	return nil
}

func interestRateOfHasAppreciatedAndAlsoInterestRateOfHasAppreciated(arg1, arg2 string) error {
	rateAAfter = rateA + 3*10000
	rateBAfter = rateB + 2*10000
	message = "Balance when both rate A and rate B have appreciated:"
	return nil
}

// change pay parameter if you want not to make the payment
func afterCheckingTheDifferenceIsPaid() error {
	pay := false
	err := makePayment(message, pay)
	if err != nil {
		return err
	}
	if !pay {
		return errors.New("the debt is not paid")
	}
	return nil
}

// scenario when rate of A has appreciated but not of B

func interestRateOfHasAppreciatedButInterestRateOfHasNotAppreciated(arg1, arg2 string) error {
	rateAAfter = rateA + 1*10000
	rateBAfter = rateB
	message = "Balance when rate A has appreciated but not rate B:"
	return nil
}

func hasToPayDepreciatedRateAppreciatedRate(arg1, arg2, arg3, arg4 string) error {
	pay := true
	err := makePayment(message, pay)
	if err != nil {
		return err
	}
	if !pay {
		return errors.New("the debt is not paid")
	}
	return nil
}

// scenario when rate A has not but rate B has
func interestRateOfHasNotAppreciatedButInterestRateOfHasAppreciated(arg1, arg2 string) error {
	rateAAfter = rateA
	rateBAfter = rateB + 1*10000
	message = "Balance when rate A has not appreciated but rate B has appreciated:"
	return nil
}

// scenario when rate A has not and rate B has not
func interestRateOfHasNotAppreciatedAndAlsoInterestRateOfHasNotAppreciated(arg1, arg2 string) error {
	message = "Balance when both rate A and rate B have not appreciated:"
	return nil
}

func thisMeansThatWeHaveNegative() error {
	return nil
}

// maybe
func thatSwapAmountIs(arg1 int) error {
	swapA = arg1
	swapB = arg1
	return nil
}

func interestRateOfHasAppreciatedFor(arg1 string, arg2 float64) error {
	rateBAfter = rateB + int(arg2*10000)
	return nil
}

func interestRateOfHasNotAppreciated(arg1 string) error {
	libor = 5.97 * 10000
	rateAAfter = libor
	rateBAfter = rateB
	return nil
}

func lIBORRateIs(arg1 float64) error {
	libor = int(arg1 * 10000)
	rateAAfter = libor
	return nil
}

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(interface{}) {
		// clean the state before every scenario
		partA = ""
		partB = ""
		addressPartA = ""
		addressPartB = ""
		jurisdictionA = ""
		jurisdictionB = ""
		firstJurisdiction = true
		rateA = 0
		rateB = 0
		swapA = 0
		swapB = 0
		swapTenor = 0
		contractTime = 0
		ipfsHash = ""
		message = ""
		rateAAfter = 0
		rateBAfter = 0

	})

	s.Step(`^that the address of "([^"]*)" is "([^"]*)" and address of "([^"]*)" is "([^"]*)"$`, thatTheAddressOfIsAndAddressOfIs)
	s.Step(`^the return statement contains both addresses$`, theReturnStatementContainsBothAddresses)
	s.Step(`^the contract should confirm that the addresses taking part in the transaction are the correct ones\.$`, theContractShouldConfirmThatTheAddressesTakingPartInTheTransactionAreTheCorrectOnes)

	// scenario with ipfs
	s.Step(`^a "([^"]*)" specified as an IPFS hash for a given contract$`, aSpecifiedAsAnIPFSHashForAGivenContract)
	s.Step(`^the contract should have a valid IPFS hash$`, theContractShouldHaveAValidIPFSHash)

	// scenario with jurisdiction
	s.Step(`^jurisdiction variable "([^"]*)" of "([^"]*)"$`, jurisdictionVariableOf)
	s.Step(`^jurisdiction variables should be equal$`, jurisdictionVariablesShouldBeEqual)

	// scenario when time hasn't come yet
	s.Step(`^that "([^"]*)" with address "([^"]*)" and "([^"]*)" with address "([^"]*)" have a swap contract$`, thatWithAddressAndWithAddressHaveASwapContract)
	s.Step(`^interest rate of "([^"]*)" is (\d+.\d+)% and interest rate of "([^"]*)" is (\d+.\d+)%$`, interestRateOfIsAndInterestRateOfIs)
	s.Step(`^that swap of "([^"]*)" is (\d+)\$ and swap of "([^"]*)" is (\d+)\$$`, thatSwapOfIsAndSwapOfIs)
	s.Step(`^swap tenor is (\d+) days$`, swapTenorIsDays)
	s.Step(`^contract started from timestamp (\d+)$`, contractStartedFromTimestamp)
	s.Step(`^current time has not passed the swap tenor$`, currentTimeHasNotPassedTheSwapTenor)
	s.Step(`^we wait for the time to pass$`, weWaitForTheTimeToPass)

	// scenario when both rate A and rate B have appreciated
	s.Step(`^current time has passed the swap tenor$`, currentTimeHasPassedTheSwapTenor)
	s.Step(`^interest rate of "([^"]*)" has appreciated and also interest rate of "([^"]*)" has appreciated$`, interestRateOfHasAppreciatedAndAlsoInterestRateOfHasAppreciated)
	s.Step(`^after checking the difference is paid$`, afterCheckingTheDifferenceIsPaid)

	// scenario when rate A has appreciated but not rate of B
	s.Step(`^interest rate of "([^"]*)" has appreciated but interest rate of "([^"]*)" has not appreciated$`, interestRateOfHasAppreciatedButInterestRateOfHasNotAppreciated)

	s.Step(`^"([^"]*)" has to pay "([^"]*)" depreciated rate\("([^"]*)"\) \+ appreciated rate\("([^"]*)"\)$`, hasToPayDepreciatedRateAppreciatedRate)

	// scenario when rate A has not but rate B has
	s.Step(`^interest rate of "([^"]*)" has not appreciated but interest rate of "([^"]*)" has appreciated$`, interestRateOfHasNotAppreciatedButInterestRateOfHasAppreciated)

	// scenario when rate A has not and rate B has not
	s.Step(`^interest rate of "([^"]*)" has not appreciated and also interest rate of "([^"]*)" has not appreciated$`, interestRateOfHasNotAppreciatedAndAlsoInterestRateOfHasNotAppreciated)
	s.Step(`^this means that we have negative$`, thisMeansThatWeHaveNegative)

	// maybe?
	s.Step(`^that swap amount is (\d+)\$$`, thatSwapAmountIs)

	// with LIBOR
	s.Step(`^"([^"]*)" swaps at LIBOR and interest rate of "([^"]*)" is (\d+.\d+)%$`, swapsAtLIBORAndInterestRateOfIs)

	// nese B appreciated a jo
	s.Step(`^interest rate of "([^"]*)" has appreciated for (\d+.\d+)%$`, interestRateOfHasAppreciatedFor)
	s.Step(`^interest rate of "([^"]*)" has not appreciated$`, interestRateOfHasNotAppreciated)

	s.Step(`^LIBOR rate is (\d+.\d+)%$`, lIBORRateIs)

}

func makePayment(message string, pay bool) error {
	f, err := os.OpenFile("log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	// connect to an ethereum node locally
	blockchain, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		return err
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")

	if err != nil {
		return err
	}

	addressAByte, err := hexutil.Decode("0xca35b7d915458ef540ade6068dfe2f44e8fa733c")
	if err != nil {
		return err
	}
	var addressA [20]byte
	copy(addressA[:], addressAByte[:20])

	addressBByte, err := hexutil.Decode("0x14723a09acff6d2a60dcdf7aa4aff308fddc160c")
	if err != nil {
		return err
	}
	var addressB [20]byte
	copy(addressB[:], addressBByte[:20])

	_, _, s, _ := contracts.DeploySwap(
		auth,
		blockchain,
		addressA,
		addressB,
		partA,
		partB,
		big.NewInt(int64(swapA)),
		big.NewInt(int64(rateA)),
		big.NewInt(int64(rateB)),
		jurisdictionA,
	)
	if pay {
		var add [32]byte
		by, err := hexutil.Decode("0x861ce807bbdcba4de666e2552ed170ea92adc4f35da9e97723e3a6ee374458d2")
		if err != nil {
			return err
		}
		copy(add[:], by[:32])

		_, err = s.SwapTransactor.Main(auth, big.NewInt(int64(rateBAfter)), big.NewInt(int64(rateAAfter)), add)
		if err != nil {
			return err
		}
	}
	balanceA, balanceB, err := s.SwapCaller.GetBalances(nil)
	if err != nil {
		return err
	}
	log.Println(message)
	log.Println("Balance of " + "partA" + ": " + (new(big.Int).Quo(balanceA, big.NewInt(10000))).String())
	log.Println("Balance of " + "partB" + ": " + (new(big.Int).Quo(balanceB, big.NewInt(10000))).String())
	return nil
}
