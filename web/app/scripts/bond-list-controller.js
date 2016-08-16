/**
 * @class BondListController
 * @classdesc
 * @ngInject
 */
function BondListController($scope, $log, $interval, $uibModal, PeerService) {

  var ctl = this;
  
  var init = function() {
    // ctl.list = PeerService.getBonds();
    PeerService.getBonds().then(function(list) {
      ctl.list = list;
    });
  };
  
  $scope.$on('$viewContentLoaded', init);
  
  $interval(init, 10000);
  
  ctl.open = function() {
    var modalInstance = $uibModal.open({
      templateUrl: 'create-bond-modal.html',
      controller: 'CreateBondModalController as ctl',
    });

    modalInstance.result.then(function(o) {
      PeerService.createBond(o);
    });
  };
  
  ctl.create = function() {
    
  };
  
  ctl.cancel = function() {
    modalInstance.dismiss('cancel');
  };

}

function CreateBondModalController($uibModalInstance, cfg) {

  var ctl = this;
  
  ctl.triggers = cfg.triggers;
  
  ctl.bond = {term: 24, principal: 100000, rate: 600, trigger: cfg.triggers[0]};
  
  ctl.ok = function () {
    $uibModalInstance.close(ctl.bond);
  };

  ctl.cancel = function () {
    $uibModalInstance.dismiss('cancel');
  };
}

angular.module('bondListController', [])
.controller('BondListController', BondListController)
.controller('CreateBondModalController', CreateBondModalController);