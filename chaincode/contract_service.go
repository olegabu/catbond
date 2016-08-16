package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"strconv"
)

type contract struct {
	IssuerId       string `json:"issuerId"`
	OwnerId        string `json:"ownerId"`
	Id             string `json:"id"`
	CouponsPaid    uint64 `json:"couponsPaid"`
	State          string `json:"state"`
}

func (t *BondChaincode) initContracts(stub *shim.ChaincodeStub) (error) {
	// Create contracts table
	err := stub.CreateTable("Contracts", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "IssuerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "OwnerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "CouponsPaid", Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: "State", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		log.Criticalf("Cannot initialize Contracts")
		return errors.New("Failed creating Contracts table.")
	}

	return nil
}

func (t *BondChaincode) createContractsForBond(stub *shim.ChaincodeStub, bond_ bond, numberOfContracts uint64) ([]byte, error) {
	if numberOfContracts > 128 {
		return nil, errors.New("Wrong number of contracts to create for bond.")
	}

	contract_ := contract{IssuerId: bond_.IssuerId, OwnerId: bond_.IssuerId, State: "offer"}
	for numberOfContracts > 0 {
		numberOfContracts--
		contract_.Id = bond_.Id + "." + strconv.FormatUint(numberOfContracts, 10)
		if _, err := t.createContract(stub, contract_); err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func (t *BondChaincode) createContract(stub *shim.ChaincodeStub, contract_ contract) ([]byte, error) {
	//TODO Verify if contract with such id is created already
	if ok, err := stub.InsertRow("Contracts", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: contract_.IssuerId}},
			&shim.Column{Value: &shim.Column_String_{String_: contract_.OwnerId}},
			&shim.Column{Value: &shim.Column_String_{String_: contract_.Id}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: contract_.CouponsPaid}},
			&shim.Column{Value: &shim.Column_String_{String_: contract_.State}}},
	}); !ok {
		log.Error("Failed inserting new contract: " + err.Error())
		return nil, err
	}

	return nil, nil
}

func (t *BondChaincode) getIssuerContracts(stub *shim.ChaincodeStub, issuerId string) (contracts []contract, err error) {
	var columns []shim.Column
	if issuerId != "" {
		columnIssuerIDs := shim.Column{Value: &shim.Column_String_{String_: issuerId}}
		columns = append(columns, columnIssuerIDs)
	}

	rows, err := stub.GetRows("Contracts", columns)
	if err != nil {
		message := "Failed retrieving contracts. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	for row := range rows {
		var result contract
		result.IssuerId = row.Columns[0].GetString_()
		result.OwnerId = row.Columns[1].GetString_()
		result.Id = row.Columns[2].GetString_()
		result.CouponsPaid = row.Columns[3].GetUint64()
		result.State = row.Columns[4].GetString_()

		log.Debugf("getContracts result includes: %+v", result)
	}

	return contracts, nil
}