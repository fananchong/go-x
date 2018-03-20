function loadPage(app, page) {
  app.directive('runoob' + toUpper(page), function() {
    return {
      templateUrl: 'pages/' + page + '.html'
    };
  });
  window['onLoad' + toUpper(page)](app);
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

function bindEvent_showPage($scope, pageEvent, page) {
  pageEvent.on('showPage', function(event, data) {
    $scope.enable = (data == page);
  }, $scope);
}
