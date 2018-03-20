function onLoadLogin(app) {
  app.controller('login', loginCtrl);

  loginCtrl.$inject = [
    '$scope',
    'pageEvent'
  ];

  function loginCtrl($scope, pageEvent) {
    bindEvent_showPage($scope, pageEvent, 'login');
    $scope.enable = true;
    $scope.click = function() {
      showPage(pageEvent, 'lobby');
    };
  }
}
