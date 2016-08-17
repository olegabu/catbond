#!/usr/bin/env bash

cd /opt/gopath/src/github.com/olegabu/catbond/support/

echo "Stopping peer and membersrvc..."
ps -ef | grep membersrvc | grep -v grep | awk '{print $2}' | xargs kill &> /tmp/kill
ps -ef | grep peer | grep -v grep | awk '{print $2}' | xargs kill &> /tmp/kill
echo "   - Stopped"

echo "Removing previous configuration and state"
#reset configuration
rm -rf /var/hyperledger/production
echo "   - Removed"

echo "Starting membersrvc."
#run service
nohup membersrvc  &> /tmp/membersrvc.log &

echo "Try to build chain code."
go build ../chaincode/

if [[ $? != 0 ]];
then
    echo "There are some mistakes in chaincode"
    exit 1;
fi

echo "   - Chain code is correct"
rm -f chaincode


echo "Starting peer..."
nohup peer node start &> /tmp/peer.log &

sleep 1
echo "   - peer started"

echo "Login all users to the Fabric"
peer network login auditor0 -p yg5DVhm0er1z
peer network login investor1 -p b7pmSxzKNFiw
peer network login investor0 -p YsWZD4qQmYxo
peer network login issuer1 -p W8G0usrU7jRk
peer network login issuer0 -p H80SiB5ODKKQ

echo "Deploying chaincode..."

OUTPUT="$(curl -XPOST -d @./commands/deployLocal.json http://localhost:5000/chaincode)"
#"$(peer chaincode deploy -p github.com/olegabu/catbond/chaincode -c '{"Function":"init", "Args": []}' -u auditor0)"
echo "   - Deployed: ${OUTPUT}"
echo "${OUTPUT}" > ./commands/HASH.txt

#sed "s/chaincodeID:.*/chaincodeID: '${OUTPUT}',/g" ../web/app/scripts/config.js > tmp.js
#mv tmp.js ../web/app/scripts/config.js

