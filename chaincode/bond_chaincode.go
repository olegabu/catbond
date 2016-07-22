/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

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
		return nil, errors.New("Failed creating Contracts table.")
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

		return t.createBond(stub, newBond)

	} else if function == "createPolicy" {
		if len(args) != 3 {
			return nil, errors.New("Incorrect arguments. Expecting contract ID, coverage and premium.")
		}

		//coverage, err := strconv.ParseUint(args[1], 10, 64)
		//if err != nil {
		//	return nil, errors.New("Incorrect coverage. Uint64 expected.")
		//}
		//
		//premium, err := strconv.ParseUint(args[2], 10, 64)
		//if err != nil {
		//	return nil, errors.New("Incorrect premium. Uint64 expected.")
		//}

		//_, err = t.createPolicy(stub, args[0], coverage, premium)
		//return nil, err
		return nil, nil

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

	} else if function == "getPolicy" {
		//if len(args) != 1 {
		//	return nil, errors.New("Incorrect number of arguments. Expecting policyID.")
		//}
		//
		//policy, err := t.getPolicy(stub, args[0])
		//if err != nil {
		//	return nil, err
		//}
		//
		//return json.Marshal(policy)
		return nil, nil

	}else {
		log.Errorf("function: %s, args: %s", function, args)
		return nil, errors.New("Received unknown function invocation")
	}
}

func main() {
	err := shim.Start(new(BondChaincode))
	if err != nil {
		log.Critical("Error starting InsuranceFrontingChaincode: %s", err)
	}
}