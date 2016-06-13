/**
 * @class MarketController
 * @classdesc
 * @ngInject
 */
function catastropheOracleController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getTriggers();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

  ctl.trigger = function(catastrophe) {
    if (catastrophe)
      return PeerService.trigger(catastrophe);
  };

}

angular.module('catastropheOracleController', [])
.controller('CatastropheOracleController', catastropheOracleController);
