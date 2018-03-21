var express = require('express');
var app = express();

app.use('/app', express.static('app'));
app.use('/', express.static('.'));

app.get('/', function(req, res) {
  res.sendFile(__dirname + "/index.html");
})

var server = app.listen(12345, function() {
  var port = server.address().port
  console.log("listen 0.0.0.0:%s", port)
})
