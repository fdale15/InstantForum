package user

/* User package was created to model a User of the forum */

import (
  "encoding/json"
)

type User struct {
  Username string
  Password string
  SocketID string
}

//Returns a user.
//Parameter takes a user object represented as a JSON object.
func GetUserForJSON(jsondata string) *User {
  u := new(User)
  json.Unmarshal([]byte(jsondata), u)
  return u
}

//Returns a JSON string representing the User object.
func (u *User) ToJSONString() string {
  jsondata, _ := json.Marshal(u)
  return string(jsondata)
}

//Returns a user based on the username provided out of the slice of users provided.
func GetUserForUsername(username string, usernames []*User) *User {
  var user *User = nil
  for _, u := range usernames {
    if u.Username == username {
      user = u
    }
  }
  return user
}

//Returns a user based on the socketid provided out of the slice of users provided.
func GetUserForSocketID(socketID string, users []*User) *User {
  var u *User = nil
  for _, user := range users {
    if user.SocketID == socketID {
      u = user
    }
  }
  return u
}
