function initLogin(app) {
  app.controller("login", function($scope) {
    scopes.login = $scope
    $scope.enable = true;
    $scope.click = function() {
      showView("lobby");
    };
  });
}
