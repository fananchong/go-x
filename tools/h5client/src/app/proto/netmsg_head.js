(function () {
    'use strict';

    var uint8 = require('./datatypes_uint8.js');
    var uint16 = require('./datatypes_uint16.js');
    var uint24 = require('./datatypes_uint24.js');

    module.exports = NetMsgHead;

    function NetMsgHead(size, cmd) {
        this.size = size + 2;
        this.flag = 0;
        this.cmd = cmd;
    }

    NetMsgHead.len = 6;

    var proto = NetMsgHead.prototype;

    proto.encode = function () {
        var buf1 = uint24.packerl(this.size);
        var buf2 = uint8.packerl(this.flag);
        var buf3 = uint16.packerl(this.cmd);
        return Buffer.concat([buf1, buf2, buf3]);
    };

    proto.decode = function (buf) {
        var pos = 0;
        this.size = uint24.unpackerl(buf.slice(pos));
        pos += uint24.size();
        this.flag = uint8.unpackerl(buf.slice(pos));
        pos += uint8.size();
        this.cmd = uint16.unpackerl(buf.slice(pos));
        pos += uint16.size();
    };

})();