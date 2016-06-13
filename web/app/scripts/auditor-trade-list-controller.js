/**
 * @class IssuerTradeListController
 * @classdesc
 * @ngInject
 */
function AuditorTradeListController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getAuditorTrades();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

}

angular.module('auditorTradeListController', [])
.controller('AuditorTradeListController', AuditorTradeListController);
