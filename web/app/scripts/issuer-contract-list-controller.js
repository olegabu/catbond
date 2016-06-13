/**
 * @class IssuerContractListController
 * @classdesc
 * @ngInject
 */
function IssuerContractListController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getIssuerContracts();
    ctl.list.forEach(function (list) {
      var bond = PeerService.getBond(list.bondId);
      list.state = bond[0].state;
    });
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

}

angular.module('issuerContractListController', [])
.controller('IssuerContractListController', IssuerContractListController);
