(function () {
  'use strict';

  var Page = require('./page.js');
  require('../proto/login_pb.js');
  var Message = require('../proto/message.js');

  module.exports = PageLogin;

  function PageLogin() { }

  PageLogin.onController = function ($scope, $http, user, pageEvent) {
    $scope.enable = true;
    $scope.txtaccount = 'test1';
    $scope.txtpassword = '123456';
    $scope.txtip = '127.0.0.1';
    $scope.txtport = 8080;
    $scope.click = function () {
      onClick($scope, $http);
    };

    function onClick($scope, $http) {
      console.log('txtaccount:', $scope.txtaccount);
      console.log('txtpassword:', $scope.txtpassword);
      console.log('txtip:', $scope.txtip);
      console.log('txtport:', $scope.txtport);

      if ($scope.txtaccount == "") {
        alert("账号名不能为空！");
        return;
      }
      if ($scope.txtpassword == "") {
        alert("密码不能为空！");
        return;
      }
      var msg = new proto.proto.MsgLogin();
      msg.setAccount($scope.txtaccount);
      msg.setPassword($scope.txtpassword);
      var urldata = Message.msgurl($scope.txtip, $scope.txtport, '0.0.1', proto.proto.MsgTypeCmd.LOGIN, msg);
      var url = urldata[0];
      var data = urldata[1];
      Message.posturl($http, url, data, function success(response) {
        console.log("login to Login success!");
        console.log('response:', response);
        user.Login(response.data);
      }, function fail(response) {
        console.log("login to Login fail!");
        console.log('response:', response);
        alert("login fail.\nresponse:" + JSON.stringify(response));
      });
    }
  };
})();