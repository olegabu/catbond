/**
 * @class InvestorTradeListController
 * @classdesc
 * @ngInject
 */
function InvestorTradeListController($scope, $log, $interval, $uibModal, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getTrades();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);
}


angular.module('investorTradeListController', [])
.controller('InvestorTradeListController', InvestorTradeListController);
