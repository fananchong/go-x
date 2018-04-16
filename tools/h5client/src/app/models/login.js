(function () {
    'use strict';

    require('../proto/common_pb.js');
    var Message = require('../proto/message.js');

    module.exports = Login;

    function Login(user) {
        this.user = user;
    }

    var proto1 = Login.prototype;

    proto1.Login = function ($http, account, password, ip, port) {
        var self = this;
        self.user.account = account;
        var msg = new proto.proto.MsgLogin();
        msg.setAccount(account);
        msg.setPassword(password);
        var urldata = Message.msgurl(ip, port, '0.0.1', proto.proto.MsgTypeCmd.LOGIN, msg);
        var url = urldata[0];
        var data = urldata[1];
        Message.posturl($http, url, data, function success(response) {
            console.log("login to Login success!");
            console.log('response:', response);
            self.user.Login(response.data);
        }, function fail(response) {
            console.log("login to Login fail!");
            console.log('response:', response);
            alert("login fail.\nresponse:" + JSON.stringify(response));
        });
    };

})();