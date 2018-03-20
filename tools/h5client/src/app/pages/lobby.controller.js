function onControllerLobby($scope, pageEvent, pageName) {
  $scope.click = function() {
    showPage(pageEvent, 'room');
  };
}
