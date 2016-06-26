package login_package

import (
  "encoding/json"
)

type Login_package struct {
  Username string
  LoggedIn bool
}

func NewLoginPackage(username string, loggedin bool) *Login_package {
  lp := new(Login_package)
  lp.Username = username
  lp.LoggedIn = loggedin
  return lp
}

func GetLoginPackageForJSON(jsondata string) *Login_package {
  lp := new(Login_package)
  json.Unmarshal([]byte(jsondata), lp)
  return lp
}

func (lp *Login_package) ToJSONString() string {
  jsondata, _ := json.Marshal(lp)
  return string(jsondata)
}
