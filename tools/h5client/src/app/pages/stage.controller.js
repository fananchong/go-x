var scopes = {};

function initStage(app) {
  app.directive("runoobStage", function() {
    return {
      templateUrl: 'pages/stage.html'
    };
  });
  initView(app, 'login');
  initView(app, 'lobby');
  initView(app, 'room');
}



function initView(app, view) {
  app.directive('runoob' + toUpper(view), function() {
    return {
      templateUrl: 'pages/' + view + '.html'
    };
  });
  window['init' + toUpper(view)](app);
}

function showView(view) {
  for (var key in scopes) {
    scopes[key].enable = false;
  }
  scopes[view].enable = true;
}
