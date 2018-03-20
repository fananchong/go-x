function onControllerLogin($scope, pageEvent, pageName) {
  $scope.enable = true;
  $scope.click = function() {
    showPage(pageEvent, 'lobby');
  };
}
