function startApp() {
  var app = angular.module("app", ['templates']);
  closePreload(app);
  initPage(app, 'stage');
}
