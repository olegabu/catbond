package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"errors"
	"strconv"
)

//trades: [{
//id: 1000,
//contractId: 'issuer0.2017.6.13.600.0',
//sellerId: 'issuer0',
//price: 100,
//state: 'offer'
//},

type trade struct {
	State 		string `json:"state"`
	ContractId 	string `json:"contractId"`
	Id 		uint64 `json:"id"`
	SellerId 	string `json:"sellerId"`
	Price 		uint64 `json:"price"`
}

func (trade_ *trade) readFromRow(row shim.Row) {
	trade_.State 		= row.Columns[0].GetString_()
	trade_.ContractId 	= row.Columns[1].GetString_()
	trade_.Id 		= row.Columns[2].GetUint64()
	trade_.SellerId 	= row.Columns[3].GetString_()
	trade_.Price 		= row.Columns[4].GetUint64()
}

func (t *BondChaincode) initTrades(stub *shim.ChaincodeStub) (error) {
	// Create trades table
	err := stub.CreateTable("Trades", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "State", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ContractId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_UINT64, Key: true},
		&shim.ColumnDefinition{Name: "SellerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Price", Type: shim.ColumnDefinition_UINT64, Key: false},
	})
	if err != nil {
		log.Criticalf("Cannot initialize Trades")
		return errors.New("Failed creating Trades table.")
	}

	err = stub.PutState("TradesCounter", []byte(strconv.FormatUint(0, 10)))
	if err != nil {
		return err
	}

	return nil
}

func (t *BondChaincode) createTradeForContract(stub *shim.ChaincodeStub, contract_ contract, price uint64) ([]byte, error) {
	log.Debugf("function: %s, args: %s", "createTradeForContract", contract_.Id)
	var trade_ trade
	trade_.State = "offer"
	trade_.ContractId = contract_.Id

	counter, err := t.incrementAndGetCounter(stub, "TradesCounter")
	if err != nil {
		return nil, err
	}

	trade_.Id = counter
	trade_.SellerId = contract_.OwnerId
	trade_.Price = price

	if ok, err := stub.InsertRow("Trades", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: trade_.State}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.ContractId}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Id}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.SellerId}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Price}}},
	}); !ok {
		log.Error("Failed inserting new trade: " + err.Error())
		return nil, err
	}

	_, err = t.changeContractState(stub, contract_.IssuerId, contract_.Id, "offer")
	return nil, err
}

func (t *BondChaincode) buy(stub *shim.ChaincodeStub, contractId string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", "buy", contractId)
	// TODO get buyer id
	// var newOwnerId string = "newOwnerId"

	trade_, err := t.getOfferTradeForContract(stub, contractId)
	if err != nil {
		message := "Failed buying trade. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Delete old trade entry with "offer" state
	var columns []shim.Column
	stateOfferColumn := shim.Column{Value: &shim.Column_String_{String_: "offer"}}
	columns = append(columns, stateOfferColumn)
	ContractIdColumn := shim.Column{Value: &shim.Column_String_{String_: contractId}}
	columns = append(columns, ContractIdColumn)
	if err := stub.DeleteRow("Trades", columns); err != nil {
		message := "Failed retrieving trades. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Create new trade entry with "settled" state
	trade_.State = "settled"
	if ok, err := stub.InsertRow("Trades", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: trade_.State}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.ContractId}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Id}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.SellerId}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Price}}},
	}); !ok {
		log.Error("Failed inserting new trade: " + err.Error())
		return nil, err
	}

	// TODO add money transfer
	// TODO transfer Contract ownership
	return nil, nil
}

func (t *BondChaincode) getOfferTrades(stub *shim.ChaincodeStub) (trades []trade, err error) {
	var columns []shim.Column
	stateOffer := shim.Column{Value: &shim.Column_String_{String_: "offer"}}
	columns = append(columns, stateOffer)

	rows, err := stub.GetRows("Trades", columns)
	if err != nil {
		message := "Failed retrieving trades. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	for row := range rows {
		var result trade
		result.readFromRow(row)

		log.Debugf("getOfferTrades result includes: %+v", result)
		trades = append(trades, result)
	}

	return trades, nil
}

func (t *BondChaincode) getOfferTradeForContract(stub *shim.ChaincodeStub, contractId string) (trade, error) {
	var columns []shim.Column
	stateOfferColumn := shim.Column{Value: &shim.Column_String_{String_: "offer"}}
	columns = append(columns, stateOfferColumn)
	ContractIdColumn := shim.Column{Value: &shim.Column_String_{String_: contractId}}
	columns = append(columns, ContractIdColumn)

	row, err := stub.GetRow("Trades", columns)
	if err != nil {
		message := "Failed retrieving trade. Error: " + err.Error()
		log.Error(message)
		return trade{}, errors.New(message)
	}

	var result trade
	result.readFromRow(row)
	log.Debugf("getOfferTradeForContract result is: %+v", result)
	return result, nil
}