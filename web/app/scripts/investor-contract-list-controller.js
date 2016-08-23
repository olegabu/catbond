/**
 * @class InvestorContractListController
 * @classdesc
 * @ngInject
 */
function InvestorContractListController($scope, $log, $interval, $uibModal, $rootScope,
    PeerService) {

  var ctl = this;
  
  var init = function() {
    PeerService.getInvestorContracts().then(function(list) {
      ctl.list = list;
    });
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  if($rootScope._timer){
    $interval.cancel($rootScope._timer);
  }
  $rootScope._timer = $interval(init, 2000);
  
  ctl.open = function(contract) {
    var modalInstance = $uibModal.open({
      templateUrl: 'sell-contract-modal.html',
      controller: 'SellModalController as ctl',
      resolve: {
        trade: function() {
          return {
            contractId: contract.id,
            price: 100,
          }
        }
      }
    });

    modalInstance.result.then(function(trade) {
      PeerService.sell(trade.contractId, trade.price);
    });
  };

}

function SellModalController($uibModalInstance, trade) {

  var ctl = this;
  
  ctl.trade = trade;
  
  ctl.ok = function () {
    $uibModalInstance.close(ctl.trade);
  };

  ctl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
}

angular.module('investorContractListController', [])
.controller('InvestorContractListController', InvestorContractListController)
.controller('SellModalController', SellModalController);