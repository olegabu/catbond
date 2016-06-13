angular.module('app', ['ui.router',
                       'ui.bootstrap',
                       'timeService',
                       'userService',
                       'peerService',
                       'demoController',
                       'bondListController',
                       'issuerContractListController',
                       'investorContractListController',
                       'auditorContractListController',
                       'auditorBondListController',
                       'auditorTradeListController',
                       'marketController',
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
  .state('demo.auditorContractList', {
    url: 'auditor-contracts',
    templateUrl: 'partials/auditor-contract-list.html',
    controller: 'AuditorContractListController as ctl'
  })
  .state('demo.auditorBondList', {
    url: 'auditor-bonds',
    templateUrl: 'partials/auditor-bond-list.html',
    controller: 'AuditorBondListController as ctl'
  })
  .state('demo.auditorTradeList', {
    url: 'auditor-trades',
    templateUrl: 'partials/auditor-trades-list.html',
    controller: 'AuditorTradeListController as ctl'
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

});
