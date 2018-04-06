(function () {
    'use strict';

    var Page = require('./page.js');

    module.exports = PageLobby;

    function PageLobby() { }

    PageLobby.onController = function ($scope, $http, user) {
        $scope.click = function () {
            Page.showPage('room');
        };
    };
})();