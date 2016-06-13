/**
 * @class MarketController
 * @classdesc
 * @ngInject
 */
function MarketController($scope, $log, $interval, $uibModal, PeerService) {

  var ctl = this;
  
  var init = function() {
    ctl.list = PeerService.getOffers();    
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  $interval(init, 1000);
  
  ctl.open = function(trade) {
    var modalInstance = $uibModal.open({
      templateUrl: 'buy-contract-modal.html',
      controller: 'BuyModalController as ctl',
      resolve: {
        trade: function() {
          return trade;
        }
      }
    });

    modalInstance.result.then(function(o) {
      PeerService.buyContract(o);
    });
  };
  
  ctl.create = function() {
    
  };
  
  ctl.cancel = function() {
    modalInstance.dismiss('cancel');
  };

}

function BuyModalController($uibModalInstance, cfg, trade) {

  var ctl = this;
  
  ctl.trade = trade;
  
  ctl.ok = function () {
    $uibModalInstance.close(ctl.trade);
  };

  ctl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
}

angular.module('marketController', [])
.controller('MarketController', MarketController)
.controller('BuyModalController', BuyModalController);