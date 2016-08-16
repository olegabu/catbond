angular.module('config', [])
.constant('cfg', 
    {
  endpoint: 'http://localhost:5000/chaincode',
  secureContext: 'user_type1_deadbeef',
  chaincodeID: 'c31825d9780ff58e1d6bc1661d6cd07cd7a50d978168c6e4557b33f0ba4bda3105e5ee93169846d247b08bcf661bd81aaaa3c2876a2b020ae266655d868cb5d9',
  users: [{id: 'issuer0', role: 'issuer'},
          {id: 'issuer1', role: 'issuer'},
          {id: 'investor0', role: 'investor'},
          {id: 'investor1', role: 'investor'},
          {id: 'auditor0', role: 'auditor'}],
  triggers: ['hurricane 2 FL', 'earthquake 5 CA'],
  bonds: [{
            id: 'issuer0.2017.6.13.600',
            issuerId: 'issuer0',
            principal: 500000,
            term: 12,
            maturityDate: '2017.6.13',
            rate: 600,
            trigger: 'hurricane 2 FL',
            state: 'offer'
          }],
  contracts: [{
            id: 'issuer0.2017.6.13.600.0',
            issuerId: 'issuer0',
            bondId: 'issuer0.2017.6.13.600',
            ownerId: 'issuer0',
            couponsPaid: 0,
            state: 'offer'
          },
          {
            id: 'issuer0.2017.6.13.600.1',
            issuerId: 'issuer0',
            bondId: 'issuer0.2017.6.13.600',
            ownerId: 'issuer0',
            couponsPaid: 0,
            state: 'offer'
          },
          {
            id: 'issuer0.2017.6.13.600.2',
            issuerId: 'issuer0',
            bondId: 'issuer0.2017.6.13.600',
            ownerId: 'issuer0',
            couponsPaid: 0,
            state: 'offer'
          },
          {
            id: 'issuer0.2017.6.13.600.3',
            issuerId: 'issuer0',
            bondId: 'issuer0.2017.6.13.600',
            ownerId: 'issuer0',
            couponsPaid: 0,
            state: 'offer'
          },
          {
            id: 'issuer0.2017.6.13.600.4',
            issuerId: 'issuer0',
            bondId: 'issuer0.2017.6.13.600',
            ownerId: 'issuer0',
            couponsPaid: 0,
            state: 'offer'
          }],
    trades: [{
            id: 1000,
            contractId: 'issuer0.2017.6.13.600.0',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1001,
            contractId: 'issuer0.2017.6.13.600.1',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1002,
            contractId: 'issuer0.2017.6.13.600.2',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1003,
            contractId: 'issuer0.2017.6.13.600.3',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          },
          {
            id: 1004,
            contractId: 'issuer0.2017.6.13.600.4',
            sellerId: 'issuer0',
            price: 100,
            state: 'offer'
          }]
    }
);
