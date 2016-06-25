package comment_package

import (
  "encoding/json"
  . "models/post"
)

type CommentPackage struct {
  PostID int
  Comment Post
}

func GetCommentPackageForJSON(jsondata string) *CommentPackage {
  cp := new(CommentPackage)
  json.Unmarshal([]byte(jsondata), cp)
  return cp
}
