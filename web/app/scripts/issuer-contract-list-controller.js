/**
 * @class IssuerContractListController
 * @classdesc
 * @ngInject
 */
function IssuerContractListController($scope, $log, $interval, PeerService) {

  var ctl = this;
  
  var init = function() {
    ctl.list = PeerService.getIssuerContracts();    
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  $interval(init, 1000);

}

angular.module('issuerContractListController', [])
.controller('IssuerContractListController', IssuerContractListController);