var app;

function startApp() {
  app = angular.module("app", ['templates']);
  closePreload(app);
  initStage(app);
}
