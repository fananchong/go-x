"use strict"

var express = require('express');

module.exports = Web

function Web() {}

Web.start = function() {
  var www = express();
  www.use('/app', express.static('app'));
  www.use('/', express.static('.'));
  www.get('/', function(req, res) {
    res.sendFile(__dirname + "/index.html");
  })
  var server = www.listen(12345, function() {
    var port = server.address().port
    console.log("listen 0.0.0.0:%s", port)
  })
}
