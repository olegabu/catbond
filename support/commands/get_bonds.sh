#!/usr/bin/env bash

HASH=$(cat HASH.txt)

cd $GOPATH/src/github.com/hyperledger/fabric/devenv/
vagrant ssh -c "peer chaincode query -l golang -n \"${HASH}\"  -c '{\"Function\": \"getBonds\", \"Args\": [\"issuer0\"]}' -u issuer0"