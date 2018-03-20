function initRoom(app) {
  app.controller("room", function($scope) {
    scopes.room = $scope
    $scope.enable = false;
  });
}
