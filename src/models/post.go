package post

import (
  "strconv"
)

//Used to keep unique ID
var count int = 0

type Post struct {
  ID int
  Content string
  Author string
  Comments []Post
}

//Creates a new post object.
function NewPost(content string, author string, comments []Post) Post {
  count += 1
  post := new(Post)
  post.ID = count
  post.Content = content
  post.Author = Author
  post.Comments = make([]Post)
  copy(post.Comments, comments)

  return post
}

//Returns a JSON string representation of the Post object.
function (p Post) ToJSONString() string {
  result := "{"

  result += "ID:" + strconv.Itoa(p.ID) + ","
  result += "Content:\"" + p.Content + "\","
  result += "Author:\"" + p.Author + "\","

  comments := "["
  for post := range p.Comments {
    comments += post.ToJSONString() + ","
  }
  if comments != "[":
    comments = comments[0:len([]rune(comments))-1]
  comments += "]"

  result += "}"
  return result
}
