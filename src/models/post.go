package post


type Post struct {
  Content string
  Author string
  Comments []Post
}

function NewPost(content string, author string, comments []Post) Post {
  post := new(Post)
  post.Content = content
  post.Author = Author
  post.Comments = make([]Post)
  copy(post.Comments, comments)

  return post
}
