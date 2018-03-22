"use strict"

var Page = require('./page.js')

module.exports = PageRoom

function PageRoom() {}

PageRoom.onController = function($scope, $http, user, pageEvent) {
  $scope.click = function() {
    Page.showPage(pageEvent, 'login');
  };
}