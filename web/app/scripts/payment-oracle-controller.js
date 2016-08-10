/**
 * @class MarketController
 * @classdesc
 * @ngInject
 */
function PaymentOracleController($scope, $log, $interval, PeerService) {

  var ctl = this;
  ctl.transferData = {}
  ctl.transferData.date = new Date();

  ctl.transfer = function(transfer) {
    PeerService.transfer(transfer);
    ctl.msg = "Transfer has been done!";
    ctl.transferData = {};
    ctl.transferData.date = new Date();
    return ctl.transferData;
  };

}

angular.module('paymentOracleController', [])
.controller('PaymentOracleController', PaymentOracleController);
