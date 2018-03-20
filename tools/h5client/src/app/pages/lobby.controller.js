function initLobby(app) {
  app.directive("runoobLobby", function() {
    return {
      templateUrl: 'pages/lobby.html'
    };
  });
  app.controller("lobby", function($scope) {
    scopes.lobby = $scope;
    $scope.enable = false;
  });
}
