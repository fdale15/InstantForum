package post

/* The post package was created the model a post on the forum */

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
  jsondata, _ := json.Marshal(p)
  return string(jsondata)
}

//Returns JSON string with posts member populated with array of JSON post objects.
func GetJSONForPosts(slice []*Post) string {
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

//Returns a post if the ID provided matches any of the posts provided.
func GetPostForID(id int, posts []*Post) *Post {
  p := new(Post)
  for _, post := range posts {
    if post.ID == id {
      p = post
    }
  }
  return p
}

//Returns the index of a post for the ID given.
func GetPostIdxForID(id int, posts []*Post) int {
  i := -1
  for idx, post := range posts {
    if post.ID == id {
      i = idx
    }
  }
  return i
}

//Removes the post with the provided id from the provided slice of posts.
func RemoveFromPostSlice(id int, posts []*Post) []*Post{
  p := make([]*Post, 0)
  for _, post := range posts {
    if post.ID != id {
      p = append(p, post)
    }
  }
  return p
}
