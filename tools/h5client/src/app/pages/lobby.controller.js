(function () {
    'use strict';

    module.exports = PageLobby;

    function PageLobby() { }

    PageLobby.onController = function ($scope, $http, user) {
        $scope.txtaccount = user.account;
        $scope.txtname = "";
        $scope.txttips = "正在获取角色信息...";
        $scope.click = function () {
            onClick();
        };
        function onClick() {
        }
    };

    PageLobby.onShow = function () {
        if (PageLobby.scope.enable) {
            PageLobby.scope.txtaccount = PageLobby.user.account;
        }
    };

})();