function initStage(app) {
  app.directive("runoobStage", function() {
    return {
      templateUrl: 'pages/stage.html'
    };
  });
  initLogin(app);
  initLobby(app);
  initRoom(app);
}

function showView(view) {
  for (var key in scopes) {
    scopes[key].enable = false;
  }
  scopes[view].enable = true;
}
