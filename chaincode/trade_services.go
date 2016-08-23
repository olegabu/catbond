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
	Id 		uint64 `json:"id"`
	ContractId 	string `json:"contractId"`
	SellerId 	string `json:"sellerId"`
	Price 		uint64 `json:"price"`
	State 		string `json:"state"`
}

func (trade_ *trade) readFromRow(row shim.Row) {
	log.Debugf("readFromRow: %+v", row)
	trade_.Id 		= row.Columns[0].GetUint64()
	trade_.ContractId 	= row.Columns[1].GetString_()
	trade_.SellerId 	= row.Columns[2].GetString_()
	trade_.Price 		= row.Columns[3].GetUint64()
	trade_.State 		= row.Columns[4].GetString_()
}

func (trade_ *trade) toRow() (shim.Row) {
	return shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Id}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.ContractId}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.SellerId}},
			&shim.Column{Value: &shim.Column_Uint64{Uint64: trade_.Price}},
			&shim.Column{Value: &shim.Column_String_{String_: trade_.State}}},
	}
}

func (t *BondChaincode) initTrades(stub shim.ChaincodeStubInterface) (error) {
	// Create trades table
	err := stub.CreateTable("Trades", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "ID", Type: shim.ColumnDefinition_UINT64, Key: true},
		&shim.ColumnDefinition{Name: "ContractId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SellerId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Price", Type: shim.ColumnDefinition_UINT64, Key: false},
		&shim.ColumnDefinition{Name: "State", Type: shim.ColumnDefinition_STRING, Key: false},
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

func (t *BondChaincode) createTradeForContract(stub shim.ChaincodeStubInterface, contract_ contract, price uint64) ([]byte, error) {
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

	if ok, err := stub.InsertRow("Trades", trade_.toRow()); !ok {
		log.Error("Failed inserting new trade: " + err.Error())
		return nil, err
	}

	_, err = t.changeContractState(stub, contract_.IssuerId, contract_.Id, "offer")
	return nil, err
}

func (t *BondChaincode) sell(stub shim.ChaincodeStubInterface, contractId string, price uint64) ([]byte, error) {
	log.Debugf("function: %s, args: %s", "sell", contractId)

	// Get Contract
	contract_, err := t.getContractById(stub, contractId)
	if err != nil {
		message := "Failed retrieving contract. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	if _, err := t.createTradeForContract(stub, contract_, price); err != nil {
		message := "createTradeForContract failed. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	return nil, nil
}

func (t *BondChaincode) buy(stub shim.ChaincodeStubInterface, tradeId uint64, newOwnerId string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", "buy", tradeId)

	trade_, err := t.getTradeByType(stub, "offer", tradeId)
	if err != nil {
		message := "Failed buying trade. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Get Contract
	contract_, err := t.getContractById(stub, trade_.ContractId)
	if err != nil {
		message := "Failed retrieving contract. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Transfer Contract ownership
	if _, err := t.changeContractOwner(stub, contract_.IssuerId, contract_.Id, newOwnerId); err != nil {
		message := "Failed transfering contract ownership. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// TODO add money transfer

	// Create new trade entry with "settled" state
	trade_.State = "reserved"
	if ok, err := stub.ReplaceRow("Trades", trade_.toRow()); !ok {
		log.Error("Failed inserting new trade: " + err.Error())
		return nil, err
	}

	return nil, nil
}

func (t *BondChaincode) confirm(stub shim.ChaincodeStubInterface, contractId string) ([]byte, error) {
	log.Debugf("function: %s, args: %s", "buy", contractId)

	trade_, err := t.getOfferTradeForContract(stub, contractId, "reserved")
	if err != nil {
		message := "Failed confirming trade. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Get Contract
	contract_, err := t.getContractById(stub, trade_.ContractId)
	if err != nil {
		message := "Failed retrieving contract. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Transfer Contract ownership
	if _, err := t.changeContractState(stub, contract_.IssuerId, contract_.Id, "active"); err != nil {
		message := "Failed transfering contract ownership. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	// Create new trade entry with "settled" state
	trade_.State = "settled"
	if ok, err := stub.ReplaceRow("Trades", trade_.toRow()); !ok {
		log.Error("Failed inserting new trade: " + err.Error())
		return nil, err
	}

	return nil, nil
}

func (t *BondChaincode) getAllTrades(stub shim.ChaincodeStubInterface) (trades []trade, err error) {
	rows, err := stub.GetRows("Trades", []shim.Column{})
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

func (t *BondChaincode) getTradesByType(stub shim.ChaincodeStubInterface, state string) (trades []trade, err error) {
	rows, err := stub.GetRows("Trades", []shim.Column{})
	if err != nil {
		message := "Failed retrieving trades. Error: " + err.Error()
		log.Error(message)
		return nil, errors.New(message)
	}

	for row := range rows {
		var result trade
		result.readFromRow(row)
		if result.State != state {
			continue
		}
		log.Debugf("getOfferTrades result includes: %+v", result)
		trades = append(trades, result)
	}

	return trades, nil
}

func (t *BondChaincode) getTradeByType(stub shim.ChaincodeStubInterface, state string, tradeId uint64) (trade, error) {
	rows, err := stub.GetRows("Trades", []shim.Column{})
	if err != nil {
		message := "Failed retrieving trades. Error: " + err.Error()
		log.Error(message)
		return trade{}, errors.New(message)
	}

	for row := range rows {
		var result trade
		result.readFromRow(row)
		if result.State == state && result.Id == tradeId {
			log.Debugf("getOfferTradeForContract returns: %+v", result)
			return result, nil
		}
	}
	return trade{}, errors.New("No trades found for id " + strconv.FormatUint(tradeId, 10))
}

func (t *BondChaincode) getOfferTradeForContract(stub shim.ChaincodeStubInterface, contractId string, state string) (trade, error) {
	rows, err := stub.GetRows("Trades", []shim.Column{})
	if err != nil {
		message := "Failed retrieving trades. Error: " + err.Error()
		log.Error(message)
		return trade{}, errors.New(message)
	}

	for row := range rows {
		var result trade
		result.readFromRow(row)
		if result.State == state && result.ContractId == contractId {
			log.Debugf("getOfferTradeForContract returns: %+v", result)
			return result, nil
		}
	}
	return trade{}, errors.New("No trades found for contract " + contractId)
}

func (t *BondChaincode) verifyTradeForContract(stub shim.ChaincodeStubInterface, contractId string, price uint64) (response) {
	trade, err := t.getOfferTradeForContract(stub, contractId, "reserved")
	log.Debugf("getOfferTradeForContract returns: %+v", trade)
	var msg response
	if err != nil {
		msg.State = "ERROR"
		msg.Msg = "Payment is not confirmed."
		log.Debugf("contract id %s with state 'reserved' not found: %+v", contractId, err)
		return msg
	}
	if trade.Price == price{
		msg.State = "OK"
		msg.Msg = "Approved"
		log.Debugf("contract approved")
	}else{
		msg.State = "ERROR"
		msg.Msg = "Incorrect price"
		log.Debugf("contract incorrect")
	}
	return msg
}