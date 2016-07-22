# Catastrophe Bond Proof Of Concept
Proof of concept demo for issuing and trading bonds on Hyperledger blockchain.
Please read the [functional spec document](docs/catbond.md).

# Changes to core.yaml 

    security.enabled = true

# Changes to membersrvc.yaml 
Add 5 more users

    auditor0: 1 yg5DVhm0er1z auditor        00011
    investor1: 1 b7pmSxzKNFiw investor        00012
    investor0: 1 YsWZD4qQmYxo investor 00013
    issuer1: 1 W8G0usrU7jRk issuer        00014
    issuer0: 1 H80SiB5ODKKQ issuer 00015


And add correct atributes to the aca.attributes section

    attribute-entry-2: auditor0;auditor;role;auditor;2015-07-13T00:00:00-03:00;;
    attribute-entry-3: investor1;investor;role;investor;2001-02-02T00:00:00-03:00;;
    attribute-entry-4: investor0;investor;role;investor;2001-02-02T00:00:00-03:00;;
    attribute-entry-5: issuer1;issuer;role;issuer;2015-01-01T00:00:00-03:00;;
    attribute-entry-6: issuer0;issuer;role;issuer;2015-01-01T00:00:00-03:00;;

run `support/deploy_chaincode.sh` to run membersrvc and one peer


# Web Application
Demo is served by an Angular single page web application. Please install and run in `web` directory.

## Install
```
npm install
bower install
```
Will download developer dependencies which may take a little while.

## Run
The web app is built and run by `gulp`:

```
gulp serve
```
