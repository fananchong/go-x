(function () {
    'use strict';

    var Util = require('../util/util.js');

    module.exports = Page;

    function Page() { }

    Page.loadPage = function (app, page) {

        app.directive('runoob' + Util.toUpper(page), function () {
            return {
                templateUrl: 'app/pages/' + page + '.html'
            };
        });

        var pageX = require('./' + page + '.controller.js');

        var onLoad = pageX.onLoad;
        if (onLoad != null) {
            onLoad(app);
        }

        function ctrl($scope, $http, user, pageEvent) {
            pageEvent.on('showPage', function (event, data) {
                $scope.enable = (data == page);
            }, $scope);
            $scope.enable = false;
            onController($scope, $http, user, pageEvent);
        }
        var onController = pageX.onController;
        if (onController != null) {
            app.controller(page, ctrl);
            ctrl.$inject = [
                '$scope',
                '$http',
                'user',
                'pageEvent'
            ];
        }
    };

    Page.initPageEventGenerator = function (app) {
        app.factory('pageEvent', obj);

        obj.$inject = [
            '$rootScope'
        ];

        function obj($rootScope) {
            var msgBus = {};
            msgBus.emit = function (msg, data) {
                data = data || {};
                $rootScope.$emit(msg, data);
            };
            msgBus.on = function (msg, func, scope) {
                var unbind = $rootScope.$on(msg, func);
                if (scope) {
                    scope.$on('$destroy', unbind);
                }
            };
            return msgBus;
        }
    };

    Page.showPage = function (pageEvent, page) {
        pageEvent.emit('showPage', page);
    };

})();