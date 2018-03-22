"use strict"

var Page = require('./page.js')
require('sprintf-js')

module.exports = PageLogin

function PageLogin() {}

PageLogin.onController = function($scope, $http, user, pageEvent) {
  $scope.enable = true;
  $scope.txtaccount = 'test1';
  $scope.txtpassword = '123456';
  $scope.txtip = '127.0.0.1';
  $scope.txtport = 8080;
  $scope.click = function() {
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

    var UserLoginReq = {
      "User": $scope.txtaccount,
      "Password": $scope.txtpassword
    };
    var data = JSON.stringify(UserLoginReq);
    var url = sprintf("http://%s:%s/login", $scope.txtip, $scope.txtport);

    $http({
      url: url,
      method: 'POST',
      data: data,
      async: false,
      headers: {
        "Access-Control-Allow-Origin": "*",
        'Access-Control-Allow-Methods': 'POST',
        'Access-Control-Allow-Headers': 'Accept,X-Custom-Header,X-Requested-With,Content-Type,Origin'
      }
    }).then(function success(response) {
      console.log("login to Login success!");
      console.log('response:', response);
      user.Login(response.data)
    }, function fail(response) {
      console.log("login to Login fail!");
      console.log('response:', response);
      alert("login fail.\nresponse:" + JSON.stringify(response));
    });
  }
}