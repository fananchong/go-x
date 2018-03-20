var app;
var scopes;

function LoadStage() {
  app = angular.module("app", []);
  scopes = {};
}

function Show(view) {
  for (var key in scopes) {
    scopes[key].enable = false;
  }
  scopes[view].enable = true;
}


function ClosePreload() {
  app.run(['$window', '$animate', '$location', '$document', '$timeout',
    function Execute($window,
      $animate,
      $location,
      $document,
      $timeout,
      settingsModel,
      projectModel) {

      function closePreload() {
        $timeout(function() {
          var element = angular.element(document.getElementById('page-preload'));
          $animate.addClass(element, 'preload-fade')
            .then(function() {
              element.remove();
            });
        }, 500);
      }

      closePreload();
    }
  ]);
}