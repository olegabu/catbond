package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"strings"
)

//issuerId: 'issuer0',
//id: 'issuer0.2017.6.13.600',
//principal: 500000,
//term: 12,
//maturityDate: '2017.6.13',
//rate: 600,
//trigger: 'hurricane 2 FL',
//state: 'offer'


type bond struct {
	IssuerId       string `json:"issuerId"`
	Id             string `json:"id"`
	Principal      uint64 `json:"principal"`
	Term           uint64 `json:"term"`
	MaturityDate   string `json:"maturityDate"`
	Rate           uint64 `json:"rate"`
	Trigger        string `json:"trigger"`
	State          string `json:"state"`
}


func (t *BondChaincode) initBonds(stub *shim.ChaincodeStub) (error) {
	// Create bonds table
	err := stub.CreateTable("Bonds", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "IssuerId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "Principal", Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: "Term", Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: "MaturityDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Rate", Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: "Trigger", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "State", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		log.Criticalf("Cannot initialize Bonds")
		return errors.New("Failed creating Bonds table.")
	}

	return nil
}

func (t *BondChaincode) getBonds(stub *shim.ChaincodeStub, issuerID string) ([]bond, error) {
	var columns []shim.Column
	if issuerID != "" {
		columnIssuerIDs := shim.Column{Value: &shim.Column_String_{String_: issuerID}}
		columns = append(columns, columnIssuerIDs)
	}

	rows, err := stub.GetRows("Bonds", columns)
	if err != nil {
		message := "Failed retrieving bonds. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	var bonds []bond

	for row := range rows {
		result := bond{
			IssuerId:       row.Columns[0].GetString_(),
			Id:             row.Columns[1].GetString_(),
			Principal:      row.Columns[2].GetUint64(),
			Term:           row.Columns[3].GetUint64(),
			MaturityDate:   row.Columns[4].GetString_(),
			Rate:           row.Columns[5].GetUint64(),
			Trigger:        row.Columns[6].GetString_(),
			State:          row.Columns[7].GetString_()}

		log.Debugf("getBonds result includes: %+v", result)
		bonds = append(bonds, result)
	}

	return bonds, nil
}

func (t *BondChaincode) createBond(stub *shim.ChaincodeStub, bond_ bond) ([]byte, error) {
	//TODO Verify if bond with such id is created already

	if ok, err := stub.InsertRow("Bonds", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: bond_.IssuerId}},
			&shim.Column{Value: &shim.Column_String_{String_: bond_.Id}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: bond_.Principal}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: bond_.Term}},
			&shim.Column{Value: &shim.Column_String_{String_: bond_.MaturityDate}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: bond_.Rate}},
			&shim.Column{Value: &shim.Column_String_{String_: bond_.Trigger}},
			&shim.Column{Value: &shim.Column_String_{String_: bond_.State}}},
	}); !ok {
		log.Error("Failed inserting new bond: " + err.Error())
		return nil, err
	}

	return nil, nil
}

func (t *BondChaincode) couponsPaid(stub *shim.ChaincodeStub, issuerId string, bondId string) ([]byte, error) {
	log.Debugf("couponsPaid called with issuerId:%s, bondId:%s", issuerId, bondId)

	// Get all contracts issued by issuerId
	contracts, err := t.getIssuerContracts(stub, issuerId)
	if err != nil {
		log.Error("couponsPaid failed on retrieving contracts: " + err.Error())
		return nil, err
	}

	// Iterate over the contracts and increment those that match bondId
	matchCounter := 0
	for _, contract_ := range contracts {
		// "issuer0.2017.6.13.600" expected after trimming a suffix from "issuer0.2017.6.13.600.42"
		if bondId == contract_.Id[:strings.LastIndex(contract_.Id, ".")] {
			matchCounter++
			if _, err := t.payContractCoupon(stub, contract_); err != nil {
				log.Errorf("couponsPaid failed on paying coupon for %s: %s", contract_.Id, err.Error())
				return nil, err
			}
		}
	}
	log.Debugf("couponsPaid: %d out of %d issued by %s matched %s and were paid",
		   matchCounter, len(contracts), issuerId, bondId)

	return nil, nil
}
