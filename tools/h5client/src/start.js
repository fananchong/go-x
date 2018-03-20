function startApp() {
  var app = angular.module("app", ['templates']);
  closePreload(app);
  initPageEventGenerator(app);
  loadPage(app, 'stage');
}
