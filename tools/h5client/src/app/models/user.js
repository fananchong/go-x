(function () {
    'use strict';

    require('../proto/common_pb.js');
    var Login = require('./login.js');
    var Gateway = require('./gateway.js');
    var Lobby = require('./lobby.js');

    module.exports = User;

    function User() {
        this.login = new Login(this);
        this.gateway = new Gateway(this);
        this.lobby = new Lobby(this);
        this.account = "";
        this.token = "";
        this.gatewayIP = "";
        this.gatewayPort = 0;
    }

    var proto1 = User.prototype;

    proto1.Login = function (data) {
        var bytes = Array.prototype.slice.call(Buffer.from(data), 0);
        var msg = proto.proto.MsgLoginResult.deserializeBinary(bytes);
        console.log('MsgLoginResult.Err:', msg.getErr());
        console.log('MsgLoginResult.Token:', msg.getToken());
        console.log('MsgLoginResult.Address:', msg.getAddress());

        if (msg.getErr() == 0) {
            this.token = msg.getToken();
            this.gatewayIP = msg.getAddress().split(":")[0];
            this.gatewayPort = parseInt(msg.getAddress().split(":")[1]);
            this.gateway.Login();
        } else {
            alert('帐号验证失败。错误码：', msg.getErr());
        }
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