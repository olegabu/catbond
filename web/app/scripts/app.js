angular.module('app', ['ui.router',
                       'ui.bootstrap',
                       'timeService',
                       'userService',
                       'peerService',
                       'demoController',
                       'oraclesController',
                       'bondListController',
                       'issuerContractListController',
                       'investorContractListController',
                       'investorTradeListController',
                       'auditorContractListController',
                       'auditorBondListController',
                       'auditorTradeListController',
                       'paymentOracleController',
                       'catastropheOracleController',
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
  .state('demo.investorTradeList', {
    url: 'investor-trades',
    templateUrl: 'partials/investor-trade-list.html',
    controller: 'InvestorTradeListController as ctl'
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
  .state('oracles', {
    url: '/oracles/',
    templateUrl: 'partials/oracles.html',
    controller: 'OraclesController as ctl'
  })
  .state('oracles.transfers', {
    url: 'payment-oracle',
    templateUrl: 'partials/transfer-form.html',
    controller: 'PaymentOracleController as ctl'
  })
  .state('oracles.triggers', {
    url: 'catastrophe-oracle',
    templateUrl: 'partials/trigger-list.html',
    controller: 'CatastropheOracleController as ctl'
  });

});
