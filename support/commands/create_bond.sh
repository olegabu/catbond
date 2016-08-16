#!/usr/bin/env bash

HASH=$(cat HASH.txt)

cd $GOPATH/src/github.com/hyperledger/fabric/devenv/
vagrant ssh -c "peer chaincode invoke -l golang -n \"${HASH}\" -c '{\"Function\": \"createBond\", \"Args\": [\"issuer0\", \"1.1.2017\", \"100000\", \"600\", \"24\"]}' -u issuer0"