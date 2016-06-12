/**
 * @class DemoController
 * @classdesc
 * @ngInject
 */
function DemoController($log, $q, $interval, $timeout, cfg, PeerService) {

  var ctl = this;

  ctl.now = new Date();

  var clockStep = 10;

  var addTime = function() {
    ctl.now.setMinutes(ctl.now.getMinutes() + clockStep);
  };

  // call init

  ctl.tick = function() {

    var i = 0, minutes = ctl.now.getMinutes();

    if(minutes === 0) {
      // call top of hour function
    }

    addTime();
  };

  var stop;

  ctl.clock = function() {
    if(angular.isDefined(stop)) {
      $interval.cancel(stop);
      stop = undefined;
    }
    else {
      stop = $interval(ctl.tick, 1000);
    }
  };

}

angular.module('demoController', [])
.controller('DemoController', DemoController);