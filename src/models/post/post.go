package post

import (
  "encoding/json"
)

//Used to keep unique ID
var count int = 0

type Post struct {
  ID int
  Content string
  Author string
  Comments []Post
}



//Returns a JSON string representation of the Post object.
func (p Post) ToJSONString() string {
  /*result := "{"

  result += "ID:" + strconv.Itoa(p.ID) + ","
  result += "Content:\"" + p.Content + "\","
  result += "Author:\"" + p.Author + "\","

  comments := "["
  for _, post := range p.Comments {
    comments += post.ToJSONString() + ","
  }
  if comments != "[" {
    comments += comments[0:len([]rune(comments))-1]
    comments += "]"
    result += comments
  }

  result += "}"
  return result*/

  jsondata, _ := json.Marshal(p)

  return string(jsondata)
}

//Creates a new post object.
func NewPost(content string, author string, comments []Post) *Post {
  count += 1
  post := new(Post)
  post.ID = count
  post.Content = content
  post.Author = author
  post.Comments = make([]Post, 0)
  copy(post.Comments, comments)

  return post
}

//Returns a Post object from JSON representation
func GetPostForJSON(jsondata string) *Post {
  p := new(Post)
  json.Unmarshal([]byte(jsondata), p)
  if p.ID == 0 {
    count += 1
    p.ID = count
  }
  return p
}
