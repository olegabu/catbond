/**
 * @class DemoController
 * @classdesc
 * @ngInject
 */
function DemoController($log, $state, 
    cfg, TimeService, UserService, PeerService) {

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
  
  ctl.user = UserService.getUser();

  ctl.users = UserService.getUsers();
  
  ctl.setUser = function() {
    UserService.setUser(ctl.user);

    if(ctl.user.role === 'issuer') {
      $state.go('demo.issuerContractList');
    }
    else if(ctl.user.role === 'investor') {
      $state.go('demo.investorContractList');
    }
  };

}

angular.module('demoController', [])
.controller('DemoController', DemoController);