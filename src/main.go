package main

import (
  "fmt"
  "strconv"
  "net/http"
  "log"
  "github.com/googollee/go-socket.io"
)

var messages []string = make([]string, 0)

func main() {

  StartWebserver(3000)
}

//Initilizes and starts the webserver on specified port.
func StartWebserver(port int) {
  server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("chat")
        so.Emit("init messages", getJsonForMessages(messages))
        so.On("chat message", func(msg string) {
            log.Println("emit:", msg)
            messages = append(messages, msg)
            so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

  http.Handle("/socket.io/", server)
  //Sets up the default root to serve static files.
  http.Handle("/", http.FileServer(http.Dir("./static")))

  //Starts the webserver listening on port 3000.
  fmt.Println("Serving on port " + strconv.Itoa(port))
  http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

func getJsonForMessages(slice []string) string {
  var result string = "{\"messages\":["
  var items string = ""
  for _, str := range slice {
    items += "\"" + str + "\","
  }
  result += items
  //Removes the extra comma if there are items in the list.
  if items != "" {
    result = result[0:len([]rune(result)) - 1]
  }

  result += "]}"
  return result
}
