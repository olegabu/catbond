/**
 * @class MarketController
 * @classdesc
 * @ngInject
 */
function PaymentOracleController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.list = PeerService.getTransfers();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

  ctl.transfer = function(o) {
    return PeerService.transfer(o);
  };

}

angular.module('paymentOracleController', [])
.controller('PaymentOracleController', PaymentOracleController);
