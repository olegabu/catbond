/**
 * @class InvestorContractListController
 * @classdesc
 * @ngInject
 */
function InvestorContractListController($scope, $log, $interval, PeerService) {

  var ctl = this;
  
  var init = function() {
    ctl.list = PeerService.getInvestorContracts();    
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  $interval(init, 1000);

}

angular.module('investorContractListController', [])
.controller('InvestorContractListController', InvestorContractListController);