function initLogin(app) {
  app.directive("runoobLogin", function() {
    return {
      templateUrl: 'pages/login.html'
    };
  });
  app.controller("login", function($scope) {
    scopes.login = $scope;
    $scope.enable = true;
    $scope.click = function() {
      showView("lobby");
    };
  });
}
