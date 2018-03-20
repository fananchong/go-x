function onControllerRoom($scope, pageEvent, pageName) {
  $scope.click = function() {
    showPage(pageEvent, 'login');
  };
}
