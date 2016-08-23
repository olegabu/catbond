angular.module('app', ['ui.router',
                       'ui.bootstrap',
                       'timeService',
                       'userService',
                       'peerService',
                       'demoController',
                       'bondListController',
                       'issuerContractListController',
                       'investorContractListController',
                       'marketController',
                       'offlineController',
                       'config'])

.config(function($stateProvider, $urlRouterProvider) {

  $urlRouterProvider.otherwise('/');

  $stateProvider
  .state('demo', {
    url: '/',
    templateUrl: 'partials/demo.html',
    controller: 'DemoController as ctl'
  })
  .state('demo.issuerContractList', {
    url: 'issuer-contracts',
    templateUrl: 'partials/issuer-contract-list.html',
    controller: 'IssuerContractListController as ctl'
  })
  .state('demo.investorContractList', {
    url: 'investor-contracts',
    templateUrl: 'partials/investor-contract-list.html',
    controller: 'InvestorContractListController as ctl'
  })
  .state('demo.bondList', {
    url: 'bonds',
    templateUrl: 'partials/bond-list.html',
    controller: 'BondListController as ctl'
  })
  .state('demo.market', {
    url: 'market',
    templateUrl: 'partials/market.html',
    controller: 'MarketController as ctl'
  })
  .state('demo.offline', {
    url: 'offline',
    templateUrl: 'partials/offline.html',
    controller: 'OfflineController as ctl'
  });

});
