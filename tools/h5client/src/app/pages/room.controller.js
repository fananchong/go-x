function initRoom(app) {
  app.directive("runoobRoom", function() {
    return {
      templateUrl: 'pages/room.html'
    };
  });
  app.controller("room", function($scope) {
    scopes.room = $scope;
    $scope.enable = false;
  });
}
