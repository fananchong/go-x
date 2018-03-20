function initRoom(app) {
  app.controller("room", roomCtrl);

  roomCtrl.$inject = [
    '$scope'
  ];

  function roomCtrl($scope) {
    scopes.room = $scope;
    $scope.enable = false;
  }
}
