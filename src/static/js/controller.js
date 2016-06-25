angular.module('InstantForum', ['btford.socket-io'])
  .factory('ifsock', function (socketFactory) {
    var socket = socketFactory();

    return socket;
  })
  .controller('PostController', function ($scope, ifsock) {
      $scope.text = "this is a test";

      ifsock.on('init messages', function (msgs) {
        console.log(msgs);
        $scope.messages = JSON.parse(msgs).messages;
      });

      ifsock.on('chat message', function (msg) {
        console.log(msg);
        $scope.messages.push(msg);
      });

      $scope.send = function () {
        console.log('test');
        ifsock.emit('chat message', $scope.text);
        $scope.messages.push($scope.text)
      }
  });
