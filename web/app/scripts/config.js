angular.module('config', [])
.constant('cfg', 
    {
  endpoint: 'https://deadbeef-6d99-41a8-988a-c31ce4ae14dc_vp1-api.blockchain.ibm.com:443/chaincode',
  secureContext: 'user_type1_deadbeef',
  chaincodeID: 'badeadbeeff1ec5ab1d077cb38907780c79260420fd94c288d6bddb710114c44969b9d560d365e38b62d379681e2acaa5b88536ebd22d6ed56c0605736349',
  users: [{id: 'issuer0', role: 'issuer'},
          {id: 'issuer1', role: 'issuer'},
          {id: 'investor0', role: 'investor'},
          {id: 'investor1', role: 'investor'},
          {id: 'auditor0', role: 'auditor'}],
  triggers: ['hurricane 2 FL', 'earthquake 5 CA'],
  bonds: [{
            id: 'issuer0.12.600',
            issuerId: 'issuer0',
            principal: 500000,
            term: 12,
            rate: 600,
            trigger: 'hurricane 2 FL',
            state: 'active'
          }],
  contracts: [{
            id: 'issuer0.12.600.0',
            issuerId: 'issuer0',
            bondId: 'issuer0.12.600',
            ownerId: 'issuer0',
            couponsPaid: 0
          },
          {
            id: 'issuer0.12.600.1',
            issuerId: 'issuer0',
            bondId: 'issuer0.12.600',
            ownerId: 'issuer0',
            couponsPaid: 0
          },
          {
            id: 'issuer0.12.600.2',
            issuerId: 'issuer0',
            bondId: 'issuer0.12.600',
            ownerId: 'issuer0',
            couponsPaid: 0
          },
          {
            id: 'issuer0.12.600.3',
            issuerId: 'issuer0',
            bondId: 'issuer0.12.600',
            ownerId: 'issuer0',
            couponsPaid: 0
          },
          {
            id: 'issuer0.12.600.4',
            issuerId: 'issuer0',
            bondId: 'issuer0.12.600',
            ownerId: 'issuer0',
            couponsPaid: 0
          }],
    trades: [{
            id: 1000,
            contractId: 'issuer0.12.600.0',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1001,
            contractId: 'issuer0.12.600.1',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1002,
            contractId: 'issuer0.12.600.2',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1003,
            contractId: 'issuer0.12.600.3',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1004,
            contractId: 'issuer0.12.600.4',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          }]
    }
);
