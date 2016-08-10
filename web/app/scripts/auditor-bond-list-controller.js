/**
 * @class IssuerContractListController
 * @classdesc
 * @ngInject
 */
function AuditorBondListController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getBonds();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

}

angular.module('auditorBondListController', [])
.controller('AuditorBondListController', AuditorBondListController);
