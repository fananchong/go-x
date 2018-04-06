(function () {
    'use strict';

    var Login = require('./login.js');

    module.exports = User;

    function User() {
        this.login = new Login(this);
        this.uid = 0;
        this.token = "";
        this.gatewayIP = "";
        this.gatewayPort = 0;
    }

    var proto = User.prototype;

    proto.Login = function (data) {
        console.log("user data = ", JSON.stringify(data));
        // this.uid = data.UID;
        // this.token = data.Token;
        // this.gatewayIP = data.LobbyAddr.split(":")[0];
        // this.gatewayPort = parseInt(data.LobbyAddr.split(":")[1]);

        // 登录Gateway
        this.gateway();
    };

    proto.gateway = function () {
        //alert('aaaa');
    };

    var u = new User();

    User.initUser = function (app) {
        app.factory('user', obj);
        obj.$inject = [
            '$rootScope'
        ];

        function obj($rootScope) {
            return u;
        }
    };

})();