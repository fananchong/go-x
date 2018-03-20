var app;
var scopes;
var rootScope;

function StartApp() {
  app = angular.module("app", ['templates']);
  scopes = {};

  // 关闭加载界面
  ClosePreload();

  app.directive("runoobStage", function() {
    return {
      templateUrl: 'pages/stage.html'
    };
  });


  app.directive("runoobLogin", function() {
    return {
      templateUrl: 'pages/login.html'
    };
  });

  app.directive("runoobLobby", function() {
    return {
      templateUrl: 'pages/lobby.html'
    };
  });

  app.directive("runoobRoom", function() {
    return {
      templateUrl: 'pages/room.html'
    };
  });

  app.controller("lobby", function($scope) {
    scopes.lobby = $scope;
    $scope.enable = false;
  });

  app.controller("login", function($scope) {
    scopes.login = $scope;
    $scope.enable = true;
    $scope.click = function() {
      Show("lobby");
    };
  });

  app.controller("room", function($scope) {
    scopes.room = $scope;
    $scope.enable = false;
  });

}