"use strict"

var Page = require('./page.js')

module.exports = PageLogin

function PageLogin() {}

PageLogin.onController = function($scope, pageEvent, pageName) {
  $scope.enable = true;
  $scope.click = function() {
    Page.showPage(pageEvent, 'lobby');
  };
}
