require('angular')
var Page = require('./app/pages/page.js')

function startApp() {
  var app = angular.module("app", []);
  closePreload(app);
  Page.initPageEventGenerator(app);
  Page.loadPage(app, 'stage');
}

startApp();
