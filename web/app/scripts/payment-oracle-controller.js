/**
 * @class MarketController
 * @classdesc
 * @ngInject
 */
function PaymentOracleController($scope, $log, $interval, PeerService) {

  var ctl = this;

  var init = function() {
    ctl.transferData = {}
    ctl.transferData.date = new Date();
  };

  $scope.$on('$viewContentLoaded', init);

  $interval(init, 1000);

  ctl.transfer = function(transfer) {
    return PeerService.transfer(transfer);
  };

}

angular.module('paymentOracleController', [])
.controller('PaymentOracleController', PaymentOracleController);
