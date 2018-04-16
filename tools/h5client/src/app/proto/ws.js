(function () {
    'use strict';

    var Util = require('../util/util.js');
    var NetMsgHead = require('./netmsg_head.js');

    module.exports = WS;

    function WS(name, onConnect) {
        this.name = name;
        this.onConnect = onConnect;
        this.ws = null;                     // WebSocket对象
        this.recvbuf = null;                // 粘包处理用
        this.cmds = {};
    }

    var proto = WS.prototype;

    proto.Connect = function (ip, port) {
        var self = this;
        self.ws = Util.initWebSocket(ip, port,
            self.onopen.bind(self),
            self.onmessage.bind(self),
            self.onclose.bind(self)
        );
    };

    proto.Send = function(cmd, msg){
        var buf = Buffer.from(msg.serializeBinary());
        var data = Buffer.concat([new NetMsgHead(buf.length, cmd).encode(), buf]);
        this.ws.send(data);
    };

    proto.onopen = function () {
        this.ws.binaryType = 'arraybuffer';
        this.onConnect();
    };

    proto.onmessage = function (data) {
        var buf = Buffer.from(data);
        if (!!this.recvbuf) {
            buf = Buffer.concat([this.recvbuf, buf]);
            this.recvbuf = null;
        }
        var pos = 0;
        while (pos < buf.length) {
            if (buf.length - pos < NetMsgHead.len) {
                break;
            }
            var curBuf = buf.slice(pos);
            var head = new NetMsgHead(0, 0);
            head.decode(curBuf);
            if (buf.length - pos < head.msgSize()) {
                break;
            }
            this.onmessage_switch(head, curBuf);
            pos += head.msgSize();
        }
        if (pos < buf.length) {
            this.recvbuf = buf.slice(pos);
        }
    };

    proto.onclose = function () {
        console.log('[' + this.name + '] onclose');
    };

    proto.onmessage_switch = function (head, buf) {
        if (head.cmd in this.cmds) {
            this.cmds[head.cmd](buf);
        } else {
            console.log('[' + this.name + '] recv message, cmd:', head.cmd, ', size:', head.size, ', flag:', head.flag);
        }
    };

})();