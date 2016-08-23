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

type response struct {
	State 	string `json:"state"`
	Msg 	string `json:"message"`
}


func (t *BondChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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

func (t *BondChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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
			return nil, errors.New("Incorrect arguments. Expecting tradeId, ownerId.")
		}

		tradeId, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return nil, errors.New("Incorrect tradeId. Uint64 expected.")
		}
		
		return t.buy(stub, tradeId, args[1])

	} else if function == "confirm" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect arguments. Expecting contractId")
		}

		return t.confirm(stub, args[0])

	} else if function == "sell" {
		if len(args) != 2 {
			return nil, errors.New("Incorrect arguments. Expecting contractId, price.")
		}

		price, err := strconv.ParseUint(args[1], 10, 64)
		if err != nil {
			return nil, errors.New("Incorrect price. Uint64 expected.")
		}

		return t.sell(stub, args[0], price)

	} else if function == "couponsPaid" {
		if len(args) != 2 {
			return nil, errors.New("Incorrect arguments. Expecting issuerId, bondId.")
		}
		// TODO: add role check
		return t.couponsPaid(stub, args[0], args[1])

	} else {
		log.Errorf("function: %s, args: %s", function, args)
		return nil, errors.New("Received unknown function invocation")
	}
}



// Query callback representing the query of a chaincode
func (t *BondChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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

	} else if function == "getContracts" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect arguments. Expecting id.")
		}

		role, err := t.getCallerRole(stub)
		if err != nil {
			return nil, err
		}
		company := args[0]

		if role == "issuer" {
			contracts, err := t.getIssuerContracts(stub, company)
			if err != nil {
				return nil, err
			}
			return json.Marshal(contracts)

		} else if role == "investor" {
			contracts, err := t.getOwnerContracts(stub, company)
			if err != nil {
				return nil, err
			}
			return json.Marshal(contracts)

		} else if role == "auditor" {
			contracts, err := t.getAllContracts(stub)
			if err != nil {
				return nil, err
			}
			return json.Marshal(contracts)
		} else {
			return nil, errors.New("Incorrect caller role. Expecting investor, issuer or auditor.")
		}
	} else if function == "getTrades" {
		if len(args) != 0 {
			return nil, errors.New("Incorrect arguments. Expecting no arguments.")
		}

		role, err := t.getCallerRole(stub)
		if err != nil {
			return nil, err
		}

		if role == "auditor" {
			trades, err := t.getAllTrades(stub)
			if err != nil {
				return nil, err
			}

			return json.Marshal(trades)
		} else {
			trades, err := t.getTradesByType(stub, "offer")
			if err != nil {
				return nil, err
			}

			return json.Marshal(trades)
		}

	} else if function == "verifyBuyRequest" {
		if len(args) != 2 {
			return nil, errors.New("Incorrect arguments. Expecting 2 arguments: contractId and price")
		}
		contractId := args[0]
		price, err := strconv.ParseUint(args[1], 10, 64)
		var msg response
		if err != nil {
			msg.State = "ERROR"
			msg.Msg = "cannot incorrect price"
			log.Debugf("incorrect price: %+v", err)
		}else {
			log.Debugf("Verify trade")
			msg = t.verifyTradeForContract(stub, contractId, price)
		}
		log.Debugf("Response: %+v", msg)
		return json.Marshal(msg)
	} else {
		log.Errorf("function: %s, args: %s", function, args)
		return nil, errors.New("Received unknown function invocation")
	}
}

func (t *BondChaincode) incrementAndGetCounter(stub shim.ChaincodeStubInterface, counterName string) (result uint64, err error) {
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

func (t *BondChaincode) getCallerCompany(stub shim.ChaincodeStubInterface) (string, error) {
	callerCompany, err := stub.ReadCertAttribute("company")
	if err != nil {
		log.Error("Failed fetching caller's company. Error: " + err.Error())
		return "", err
	}
	log.Debugf("Caller company is: %s", callerCompany)
	return string(callerCompany), nil
}

func (t *BondChaincode) getCallerRole(stub shim.ChaincodeStubInterface) (string, error) {
	callerRole, err := stub.ReadCertAttribute("role")
	if err != nil {
		log.Error("Failed fetching caller role. Error: " + err.Error())
		return "", err
	}
	log.Debugf("Caller role is: %s", callerRole)
	return string(callerRole), nil
}

func main() {
	err := shim.Start(new(BondChaincode))
	if err != nil {
		log.Critical("Error starting InsuranceFrontingChaincode: %s", err)
	}
}