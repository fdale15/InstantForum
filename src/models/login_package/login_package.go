package login_package

/* The login_package package was created to
allow the server to send the username logged in as
and a value indicating whether the login attempt was
successful or not.
(This is not really a great way of handling a login.)*/

import (
  "encoding/json"
)

type LoginPackage struct {
  Username string
  LoggedIn bool
}

//Returns a LoginPackage object.
//Parameter takes a LoginPackage object represented as a JSON object.
func GetLoginPackageForJSON(jsondata string) *LoginPackage {
  lp := new(LoginPackage)
  json.Unmarshal([]byte(jsondata), lp)
  return lp
}

//Member function of LoginPackage object.
//It returns the JSON string representation of the login object.
func (lp *LoginPackage) ToJSONString() string {
  jsondata, _ := json.Marshal(lp)
  return string(jsondata)
}
