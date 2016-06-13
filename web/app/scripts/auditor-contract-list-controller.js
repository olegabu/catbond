/**
 * @class IssuerContractListController
 * @classdesc
 * @ngInject
 */
function AuditorContractListController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getAuditorContracts();
    ctl.list.forEach(function (list) {
      var bond = PeerService.getBond(list.bondId);
      list.state = bond[0].state;
    });
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

}

angular.module('auditorContractListController', [])
.controller('AuditorContractListController', AuditorContractListController);
