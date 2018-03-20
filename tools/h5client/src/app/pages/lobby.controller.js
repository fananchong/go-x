function initLobby(app) {
  app.controller("lobby", lobbyCtrl);

  lobbyCtrl.$inject = [
    '$scope'
  ];

  function lobbyCtrl($scope) {
    scopes.lobby = $scope;
    $scope.enable = false;
  }
}
