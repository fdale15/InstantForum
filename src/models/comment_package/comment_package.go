package comment_package

/* The comment_package package was created to allow a
comment to be sent with its corresponding post's id. */

import (
  "encoding/json"
  . "models/post"
)

type CommentPackage struct {
  PostID int
  Comment Post
}

//Returns a CommentPackage.
//The parameter should be a CommentPackage represented in JSON.
func GetCommentPackageForJSON(jsondata string) *CommentPackage {
  cp := new(CommentPackage)
  json.Unmarshal([]byte(jsondata), cp)
  return cp
}
