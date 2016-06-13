/**
 * @class InvestorContractListController
 * @classdesc
 * @ngInject
 */
function InvestorContractListController($scope, $log, $interval, $uibModal, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getInvestorContracts();
    ctl.list.forEach(function (list) {
      var trade = PeerService.getTrade(list.id);
      list.price = trade[0].price;
    });
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

  ctl.open = function(trade) {
    var modalInstance = $uibModal.open({
      templateUrl: 'sell-contract-modal.html',
      controller: 'SellModalController as ctl',
      resolve: {
        trade: function() {
          return trade;
        }
      }
    });

    modalInstance.result.then(function(o) {
      PeerService.sellContract(o);
    });
  };

  ctl.create = function() {

  };

  ctl.cancel = function() {
    modalInstance.dismiss('cancel');
  };
}

function SellModalController($uibModalInstance, cfg, trade) {

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
