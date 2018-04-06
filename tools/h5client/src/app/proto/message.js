(function () {
    'use strict';

    var md5 = require("blueimp-md5");
    require('sprintf-js');

    module.exports = Message;

    function Message() { }

    var SIGN1 = "5UY6$f$h";
    var SIGN2 = "3wokZB%q";
    var SIGN3 = "%2Fi9TRf";

    Message.msgurl = function (addr, port, version, cmd, msg) {
        var t = String(new Date().getTime());
        var c = String(cmd);
        var s = md5(SIGN1 + c + SIGN2 + t + SIGN3 + version);
        var d = msg.serializeBinary();
        console.log("d1=", d);
        var url = sprintf("http://%s:%s/msg", addr, port);
        var params = { 'c': c, 't': t, 'd': d, 's': s };
        return [url, params];
    };

    Message.posturl = function ($http, url, data, cb_success, cb_fail) {
        $http({
            url: url,
            method: 'POST',
            data: data,
            async: false,
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                "Access-Control-Allow-Origin": "*",
                'Access-Control-Allow-Methods': 'POST',
                'Access-Control-Allow-Headers': 'Accept,X-Custom-Header,X-Requested-With,Content-Type, Origin'
            },
            transformRequest: function (obj) {
                var str = [];
                for (var s in obj) {
                    str.push(encodeURIComponent(s) + "=" + encodeURIComponent(obj[s]));
                    if (s == 'd') {
                        console.log("d2=", encodeURIComponent(obj[s]));
                    }
                }
                return str.join("&");
            }
        }).then(cb_success, cb_fail);
    };

    Message.geturl = function ($http, url, params, cb_success, cb_fail) {
        $http({
            url: url,
            method: 'GET',
            params: params,
            async: false,
            transformRequest: function (obj) {
                var str = [];
                for (var s in obj) {
                    str.push(encodeURIComponent(s) + "=" + encodeURIComponent(obj[s]));
                }
                return str.join("&");
            }
        }).then(cb_success, cb_fail);
    };
})();