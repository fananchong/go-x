function initLogin(app) {
  app.controller("login", loginCtrl);

  loginCtrl.$inject = [
    '$scope'
  ];

  function loginCtrl($scope) {
    scopes.login = $scope
    $scope.enable = true;
    $scope.click = function() {
      showView("lobby");
    };
  }
}
