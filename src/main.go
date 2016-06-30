package main

import (
  "fmt"
  "strconv"
  "net/http"
  "log"
  "github.com/googollee/go-socket.io"
  . "models/post"
  . "models/comment_package"
  . "models/login_package"
  . "models/user"
)

//These slices hold the posts of the forum and users that have logged in.
var posts []*Post = make([]*Post, 0)
var users []*User = make([]*User, 0)

//The entry point of the program.
func main() {
  StartWebserver(3000)
}

//Initilizes and starts the webserver on specified port.
func StartWebserver(port int) {
  //Sets up the Socket.io server and sets http to use to handle requests on /socket.io/ path.
  server := SetupForumServer()
  http.Handle("/socket.io/", server)

  //Sets up the default index to serve static files.
  http.Handle("/", http.FileServer(http.Dir("./static")))

  //Starts the webserver listening on port 3000.
  fmt.Println("Serving on port " + strconv.Itoa(port))
  http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

//Sets up the socket.io server for handling of forum posts.
func SetupForumServer() *socketio.Server {
  //Creates a new Socket.IO server object.
  server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    //Sets up the connection event.
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("forum")
        //Sends current posts to a newly initiated connection.
        so.Emit("init posts", GetJSONForPosts(posts))

        //Sets up the socket's forum post event.
        so.On("forum post", func(msg string) {
            p := GetPostForJSON(msg)

            log.Println("emit forum post:", msg)
            //Posts are recorded for future connections.
            posts = append(posts, p)
            //Posts are broadcast to everyone listening to the forum room.
            so.BroadcastTo("forum", "forum post", p.ToJSONString())
            so.Emit("forum post", p.ToJSONString())
        })

        //Sets up the socket's comment post event.
        so.On("comment post", func(comment string) {
          log.Println("emit comment post:", comment)
          cp := GetCommentPackageForJSON(comment)
          p := GetPostForID(cp.PostID, posts)
          p.Comments = append(p.Comments, cp.Comment)
          so.BroadcastTo("forum", "comment post", comment)
        })
        //Handles the login event.
        so.On("login", func(user string) {
          sentuser := GetUserForJSON(user)
          storeduser := GetUserForUsername(sentuser.Username, users)

          lp := new(LoginPackage)
          lp.LoggedIn = false
          lp.Username = sentuser.Username

          //if the user has logged in before check for password, otherwise login and store password.
          if storeduser != nil {
            if storeduser.Password == sentuser.Password {
              storeduser.SocketID = so.Id()
              lp.LoggedIn = true
              log.Println("Log in: " + storeduser.Username)
            }
          } else {
            lp.LoggedIn = true
            sentuser.SocketID = so.Id()
            users = append(users, sentuser)
            log.Println("New user: " + sentuser.Username)
          }

          so.Emit("login", lp.ToJSONString())
        })

        so.On("delete", func(msg int) {
          postId := msg
          post := GetPostForID(postId, posts)
          user := GetUserForSocketID(so.Id(), users)

          if user != nil {
            if post.Author == user.Username {
              log.Println("emit delete: {post}")
              log.Println(post)
              posts = RemoveFromPostSlice(postId, posts)
              so.BroadcastTo("forum", "delete", msg)
              so.Emit("delete", msg)
            }
          }
        })

        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })
    return server
}
