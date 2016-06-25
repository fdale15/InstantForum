angular.module('InstantForum', ['btford.socket-io'])
  .factory('ifsock', function (socketFactory) {
    var socket = socketFactory();

    return socket;
  })
  .controller('PostController', function ($scope, ifsock) {
      $scope.text = "";
      $scope.author = "";
      $scope.username = "";
      $scope.comment_content = "";
      $scope.posts = [];

      ifsock.on('init posts', function (posts) {
        console.log(posts);
        $scope.posts = JSON.parse(posts).posts;
      });

      //Listens for a new thread to be posted, then appends it to the posts
      ifsock.on('forum post', function (post) {
        post = JSON.parse(post);
        post.Comments = [];
        console.log(post);
        $scope.posts.push(post);
      });

      //Listens for a comment to be posted, then appends it to the correct posts comments.
      ifsock.on('comment post', function (comment_package) {
        console.log(comment_package);
        cp = JSON.parse(comment_package);
        post_idx = GetIndexOfPostForID(cp.PostID);
        $scope.posts[post_idx].Comments.push(cp.Comment);
        console.log(post_idx);
        console.log($scope.posts[post_idx]);
      })

      //Handles sending a new post to the server.
      $scope.send = function () {
        if ($scope.author.length > 0) {
          if ($scope.text.length > 0) {
            post = {};
            post.Content = $scope.text;
            post.Author = $scope.author;
            post.Comments = [];

            console.log(post);
            ifsock.emit('forum post', JSON.stringify(post));
            $scope.posts.push(post);
            $scope.text = "";
          }
        }
        else {
          alert('You must enter a name.');
        }
      }

      //Handles sending a comment to the server.
      $scope.comment = function(postId) {
        if ($scope.author.length > 0) {
          if ($scope.comment_content.length > 0) {
            console.log("postID: " + postId);
            comment = {};
            comment.Author = $scope.author;
            comment.Content = $scope.comment_content;

            comment_package = { "PostID" : postId, "Comment" : comment };
            console.log(comment_package);
            ifsock.emit('comment post', JSON.stringify(comment_package))
            $scope.posts[GetIndexOfPostForID(postId)].Comments.push(comment);
            $scope.comment_content = "";
          }
        }
        else {
          alert('You must enter a name.');
        }
      }

      $scope.join = function () {
        $scope.author = $scope.username;
        console.log($scope.username);
      }

      var GetIndexOfPostForID = function (id) {
        console.log(id);
        idx = -1;
        for (i in $scope.posts) {
          console.log(i + " " + $scope.posts[i].ID);
          if ($scope.posts[i].ID == id) {
            idx = i;
            break;
          }
        }
        return idx;
      }
  });
