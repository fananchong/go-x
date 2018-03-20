function initLobby(app) {
  app.controller("lobby", function($scope) {
    scopes.lobby = $scope
    $scope.enable = false;
  });
}
