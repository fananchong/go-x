"use strict"

var Page = require('./page.js')

module.exports = PageLobby

function PageLobby() {}

PageLobby.onController = function($scope, $http, user, pageEvent) {
  $scope.click = function() {
    Page.showPage(pageEvent, 'room');
  };
}