/**
 * @class IssuerContractListController
 * @classdesc
 * @ngInject
 */
function IssuerContractListController($scope, $log, $interval, PeerService, $rootScope) {

  var ctl = this;
  
  var init = function() {
//    ctl.list = PeerService.getIssuerContracts();
        PeerService.getIssuerContracts().then(function(list) {
          ctl.list = list;
        });
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  if($rootScope._timer){
    $interval.cancel($rootScope._timer);
  }
  $rootScope._timer = $interval(init, 2000);

}

angular.module('issuerContractListController', [])
.controller('IssuerContractListController', IssuerContractListController);