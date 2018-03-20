var scopes = {};

function initPage(app, page) {
  app.directive('runoob' + toUpper(page), function() {
    return {
      templateUrl: 'pages/' + page + '.html'
    };
  });
  window['init' + toUpper(page)](app);
}

function showPage(page) {
  for (var key in scopes) {
    scopes[key].enable = false;
  }
  scopes[page].enable = true;
}
