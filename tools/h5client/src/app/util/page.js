function loadPage(app, page) {
  app.directive('runoob' + toUpper(page), function() {
    return {
      templateUrl: 'pages/' + page + '.html'
    };
  });

  var onLoad = window['onLoad' + toUpper(page)];
  if (onLoad != null) {
    onLoad(app);
  }

  function ctrl($scope, pageEvent) {
    pageEvent.on('showPage', function(event, data) {
      $scope.enable = (data == page);
    }, $scope);
    $scope.enable = false;
    onController($scope, pageEvent, page);
  }
  var onController = window['onController' + toUpper(page)];
  if (onController != null) {
    app.controller(page, ctrl);
    ctrl.$inject = [
      '$scope',
      'pageEvent'
    ];
  }
}

function initPageEventGenerator(app) {
  app.factory('pageEvent', pageEvent);

  pageEvent.$inject = [
    '$rootScope'
  ];

  function pageEvent($rootScope) {
    var msgBus = {};
    msgBus.emit = function(msg, data) {
      data = data || {};
      $rootScope.$emit(msg, data);
    };
    msgBus.on = function(msg, func, scope) {
      var unbind = $rootScope.$on(msg, func);
      if (scope) {
        scope.$on('$destroy', unbind);
      }
    };
    return msgBus;
  }
}

// event: showPage
function showPage(pageEvent, page) {
  pageEvent.emit('showPage', page);
}
