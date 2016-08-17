package main

import (
	"errors"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"

	"encoding/json"
	"strconv"
)

var log = logging.MustGetLogger("bond-traiding")

// SimpleChaincode example simple Chaincode implementation
type BondChaincode struct {
}

func (t *BondChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", function, args)

	// Create bonds table
	err := t.initBonds(stub)
	if err != nil {
		log.Criticalf("function: %s, args: %s", function, args)
		return nil, errors.New("Failed creating Bond table.")
	}
	// Create contracts table
	err = t.initContracts(stub)
	if err != nil {
		log.Criticalf("function: %s, args: %s", function, args)
		return nil, errors.New("Failed creating Contracts table.")
	}
	// Create trades table
	err = t.initTrades(stub)
	if err != nil {
		log.Criticalf("function: %s, args: %s", function, args)
		return nil, errors.New("Failed creating Trades table.")
	}

	return nil, nil
}

func (t *BondChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", function, args)

	// Handle different functions
	if function == "createBond" {
		if len(args) != 5 {
			return nil, errors.New("Incorrect arguments. Expecting issuerId, maturityDate, principal, rate and term.")
		}

		var newBond bond

		newBond.IssuerId = args[0]
		newBond.MaturityDate = args[1]

		principal, err := strconv.ParseUint(args[2], 10, 64)
		if err != nil {
			return nil, errors.New("Incorrect principa. Uint64 expected.")
		}
		newBond.Principal = principal

		rate, err := strconv.ParseUint(args[3], 10, 64)
		if err != nil {
			return nil, errors.New("Incorrect rate. Uint64 expected.")
		}
		newBond.Rate = rate

		term, err := strconv.ParseUint(args[4], 10, 64)
		if err != nil {
			return nil, errors.New("Incorrect term. Uint64 expected.")
		}
		newBond.Term = term

		// TODO check with Oleg is state should be hardcoded on a contract level
		newBond.State = "active"

		newBond.Id = newBond.IssuerId + "." + newBond.MaturityDate + "." + strconv.FormatUint(newBond.Rate, 10)

		if msg, err := t.createBond(stub, newBond); err != nil {
			return msg, err
		}
		return t.createContractsForBond(stub, newBond, 5)

	} else if function == "buy" {
		if len(args) != 2 {
			return nil, errors.New("Incorrect arguments. Expecting contractId, ownerId.")
		}
		
		return t.buy(stub, args[0], args[1])

	}else {
		log.Errorf("function: %s, args: %s", function, args)
		return nil, errors.New("Received unknown function invocation")
	}
}



// Query callback representing the query of a chaincode
func (t *BondChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", function, args)
	// Handle different functions
	if function == "getBonds" {
		var issuerId string
		if len(args) == 1 {
			issuerId = args[0]
		}

		bonds, err := t.getBonds(stub, issuerId)
		if err != nil {
			return nil, err
		}

		return json.Marshal(bonds)

	} else if function == "getIssuerContracts" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect arguments. Expecting issuerId.")
		}

		issuerId := args[0]
		contracts, err := t.getIssuerContracts(stub, issuerId)
		if err != nil {
			return nil, err
		}

		return json.Marshal(contracts)

	} else if function == "getOwnerContracts" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect arguments. Expecting ownerId.")
		}

		ownerId := args[0]
		contracts, err := t.getOwnerContracts(stub, ownerId)
		if err != nil {
			return nil, err
		}

		return json.Marshal(contracts)

	} else if function == "getOfferTrades" {
		if len(args) != 0 {
			return nil, errors.New("Incorrect arguments. Expecting no arguments.")
		}

		trades, err := t.getOfferTrades(stub)
		if err != nil {
			return nil, err
		}

		return json.Marshal(trades)

	} else {
		log.Errorf("function: %s, args: %s", function, args)
		return nil, errors.New("Received unknown function invocation")
	}
}

func (t *BondChaincode) incrementAndGetCounter(stub *shim.ChaincodeStub, counterName string) (result uint64, err error) {
	if contractIDBytes, err := stub.GetState(counterName); err != nil {
		log.Errorf("Failed retrieving %s.", counterName)
		return result, err
	} else {
		result, _ = strconv.ParseUint(string(contractIDBytes), 10, 64)
	}
	result++
	if err = stub.PutState(counterName, []byte(strconv.FormatUint(result, 10))); err != nil {
		log.Errorf("Failed saving %s!", counterName)
		return result, err
	}
	return result, err
}

func main() {
	err := shim.Start(new(BondChaincode))
	if err != nil {
		log.Critical("Error starting InsuranceFrontingChaincode: %s", err)
	}
}