function onLoadLobby(app) {
  app.controller('lobby', lobbyCtrl);

  lobbyCtrl.$inject = [
    '$scope',
    'pageEvent'
  ];

  function lobbyCtrl($scope, pageEvent) {
    bindEvent_showPage($scope, pageEvent, 'lobby');
    $scope.enable = false;
  }
}
