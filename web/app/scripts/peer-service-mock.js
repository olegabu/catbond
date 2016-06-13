/**
 * @class PeerService
 * @classdesc
 * @ngInject
 */
function PeerService($log, $q, $http, cfg, UserService) {

  // jshint shadow: true
  var PeerService = this;
  var tradeId = 1100;
  
  PeerService.getIssuerContracts = function() {
    return _.filter(cfg.contracts, function(o) { 
      return o.issuerId === UserService.getUser().id;
    });
  };
  
  PeerService.getInvestorContracts = function() {
    return _.filter(cfg.contracts, function(o) { 
      return o.ownerId === UserService.getUser().id;
    });
  };
  
  PeerService.getOffers = function() {
    return _.filter(cfg.trades, function(o) { 
      return o.state === 'offer';
    });
  };
  
  PeerService.getBonds = function() {
    return _.filter(cfg.bonds, function(o) { 
      return o.issuerId === UserService.getUser().id;
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
          state: 'active'
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
  
  PeerService.buyContract = function(trade) {
    var buyerId = UserService.getUser().id;
    
    var t = _.find(cfg.trades, function(o) {
      return o.id === trade.id;
    });
    
    //TODO first put the trade in captured state 
    // later payment oracle sets it to settled
    t.state = 'settled';
    t.buyerId = buyerId;
    
    var c = _.find(cfg.contracts, function(o) {
      return o.id === trade.contractId;
    });
    
    c.ownerId = buyerId;
  };

}

angular.module('peerService', []).service('PeerService', PeerService);