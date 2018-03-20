function onLoadRoom(app) {
  app.controller('room', roomCtrl);

  roomCtrl.$inject = [
    '$scope',
    'pageEvent'
  ];

  function roomCtrl($scope, pageEvent) {
    bindEvent_showPage($scope, pageEvent, 'room');
    $scope.enable = false;
  }
}
