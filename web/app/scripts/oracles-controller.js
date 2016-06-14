/**
 * @class OraclesController
 * @classdesc
 * @ngInject
 */
function OraclesController(cfg, TimeService) {

  var ctl = this;

  ctl.now = function() {
    return TimeService.now;
  };

  ctl.tick = function() {
    TimeService.tick();
  };

  ctl.clock = function() {
    TimeService.clock();
  };

}

angular.module('oraclesController', [])
.controller('OraclesController', OraclesController);
