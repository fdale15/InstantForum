package user

import (
  "encoding/json"
)

type User struct {
  Username string
  Password string
}

func NewUser(username, password string) *User {
  u := new(User)
  u.Username = username
  u.Password = password

  return u
}

func GetUserForJSON(jsondata string) *User {
  u := new(User)
  json.Unmarshal([]byte(jsondata), u)
  return u
}

func (u *User) ToJSONString() string {
  jsondata, _ := json.Marshal(u)
  return string(jsondata)
}

func GetUserForUsername(username string, usernames []*User) *User {
  var user *User = nil
  for _, u := range usernames {
    if u.Username == username {
      user = u
    }
  }
  return user
}
