/**
 * @class PeerService
 * @classdesc
 * @ngInject
 */
function PeerService($log, $q, $http, cfg, UserService) {

  // jshint shadow: true
  var PeerService = this;
  var tradeId = 1100;

  PeerService.getContracts = function() {
    var role = UserService.getUser().role;
    return _.filter(cfg.contracts, function(o) {
      var userId = UserService.getUser().id;
      if (role === 'issuer')
        return o.issuerId === userId;
      if (role === 'investor')
        return o.ownerId === userId;
      if (role === 'auditor')
        return o;
    });
  };

  PeerService.getTrades = function() {
    var role = UserService.getUser().role;
    return _.filter(cfg.trades, function(o) {
      if (role === 'investor')
        return o.sellerId === UserService.getUser().id;
      if (role === 'auditor')
        return o;
    });
  };

  PeerService.getOffers = function() {
    return _.filter(cfg.trades, function(o) {
      return o.state === 'offer';
    });
  };

  PeerService.getTransfers = function() {
    return _.filter(cfg.trades, function(o) {
      return o.state === 'captured';
    });
  };

  PeerService.getTriggers = function() {
    return cfg.triggers;
  };

  PeerService.getTrade = function(contractId) {
    return _.filter(cfg.trades, function(o) {
      return o.contractId === contractId;
    });
  };

  PeerService.getBonds = function() {
    var role = UserService.getUser().role;
    return _.filter(cfg.bonds, function(o) {
      if (role === 'issuer')
        return o.issuerId === UserService.getUser().id;
      if (role === 'auditor')
        return o;
    });
  };

  PeerService.getBond = function(bondId) {
    return _.filter(cfg.bonds, function(o) {
      return o.id === bondId;
    });
  };

  var getMaturityDate = function(term) {
    var now = new Date();
    return new Date(now.getFullYear(), now.getMonth() + term, now.getDate());
  };

  var getMaturityDateString = function(term) {
    var m = getMaturityDate(term);
    return m.getFullYear() + '.' + (m.getMonth() + 1) + '.' + m.getDate();
  };

  PeerService.createBond = function(bond) {
    bond.issuerId = UserService.getUser().id;
    bond.maturityDate = getMaturityDateString(bond.term);
    bond.state = 'active';

    bond.id = bond.issuerId + '.' + bond.maturityDate + '.' + bond.rate;

    // fail if bond with this id already exists
    var exist = _.find(cfg.bonds, function(o) {
      return o.id === bond.id;
    });

    if(exist) {
      $log.error('bond already exists', exist);
      return;
    }

    cfg.bonds.push(bond);

    // create contracts and offer them for sale: create trades in offer state

    var numContracts = bond.principal / 100000;
    var i;

    for(i=0; i < numContracts; i++) {
      var contract = {
          id: bond.id + '.' + i,
          bondId: bond.id,
          issuerId: bond.issuerId,
          ownerId: bond.issuerId,
          couponsPaid: 0,
          state: 'offer'
      };

      var trade = {
          id: tradeId++,
          contractId: contract.id,
          sellerId: bond.issuerId,
          price: 100,
          state: 'offer'
      }

      cfg.contracts.push(contract);
      cfg.trades.push(trade);
    }
  };

  PeerService.transfer = function(transferData) {
    var splits = transferData.note.split('.');
    var action = splits[0];
    if (action === 'trade') {
      var tradeId = parseFloat(splits[1]);
      var t = _.find(cfg.trades, function(o) {
        return o.id === tradeId;
      });
      t.state = 'settled';
    }
    if (action === 'coupon' || action === 'premium') {
      splits.splice(0,1);
      var contractId = splits.join('.');
      var contract = _.find(cfg.contracts, function(o) {
        console.log('o.id', o.id);
        console.log('contractId', contractId);
        console.log('o.id typeof', typeof o.id);
        console.log('contractId typeof', typeof contractId);
        return o.id === contractId;
      });
      if(action === 'coupon') {
        contract.couponsPaid += 1;
      }
      if(action === 'premium') {
        contract.state === 'cancel';
      }
    }
    return transferData;
  };

  PeerService.buy = function(tradeId) {
    var buyerId = UserService.getUser().id;

    var trade = _.find(cfg.trades, function(o) {
      return o.id === tradeId;
    });

    //TODO first put the trade in captured state
    // later payment oracle sets it to settled
    trade.state = 'settled';
    trade.buyerId = buyerId;

    var contract = _.find(cfg.contracts, function(o) {
      return o.id === trade.contractId;
    });

    contract.state = 'active';
    contract.ownerId = buyerId;
  };

  PeerService.sell = function(contractId, price) {
    var sellerId = UserService.getUser().id;

    var contract = _.find(cfg.contracts, function(o) {
      return o.id === contractId;
    });

    // set contract's state so it cannot be sold twice
    contract.state = 'offer';

    var trade = {
        id: tradeId++,
        contractId: contract.id,
        sellerId: sellerId,
        price: price,
        state: 'offer'
    }

    cfg.trades.push(trade);
  };

  PeerService.trigger = function(catastrophe) {
    var bonds = _.filter(cfg.bonds, function(o) {
      return o.trigger === catastrophe;
    });

    bonds.forEach(function(bond) {
      bond.state = 'triggered';
    });
  };
}

angular.module('peerService', []).service('PeerService', PeerService);
