package main

import (
  "fmt"
  "strconv"
  "net/http"
  "log"
  "github.com/googollee/go-socket.io"
  . "models/post"
  . "models/comment_package"
)

var posts []*Post = make([]*Post, 0)

func main() {
  StartWebserver(3000)
}

//Initilizes and starts the webserver on specified port.
func StartWebserver(port int) {
  //Sets up the path to server socket.io requests from.
  server := SetupForumServer()
  http.Handle("/socket.io/", server)

  //Sets up the default root to serve static files.
  http.Handle("/", http.FileServer(http.Dir("./static")))

  //Starts the webserver listening on port 3000.
  fmt.Println("Serving on port " + strconv.Itoa(port))
  http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

func getJsonForPosts(slice []*Post) string {
  var result string = "{\"posts\":["
  var items string = ""
  for _, p := range slice {
    items += p.ToJSONString() + ","
  }
  result += items
  //Removes the extra comma if there are items in the list.
  if items != "" {
    result = result[0:len([]rune(result)) - 1]
  }

  result += "]}"
  return result
}

//Sets up the socket.io server for handling of forum posts.
func SetupForumServer() *socketio.Server {
  server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    //Sets up the connection event.
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("forum")
        //Sends current posts to a newly initiated connection.
        so.Emit("init posts", getJsonForPosts(posts))

        //Sets up the socket's forum post event.
        so.On("forum post", func(msg string) {
            p := GetPostForJSON(msg)
            log.Println("emit:", msg)
            //Posts are recorded for future connections.
            posts = append(posts, p)
            //Posts are broadcast to everyone listening to the forum room.
            so.BroadcastTo("forum", "forum post", msg)
        })

        //Sets up the socket's comment post event.
        so.On("comment post", func(comment string) {
          log.Println("emit:", comment)
          cp := GetCommentPackageForJSON(comment)
          p := GetPostForID(cp.PostID)
          p.Comments = append(p.Comments, cp.Comment)
          so.BroadcastTo("forum", "comment post", comment)
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

func GetPostForID(id int) *Post {
  p := new(Post)
  for _, post := range posts {
    if post.ID == id {
      p = post
    }
  }
  return p
}
