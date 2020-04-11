// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// type postController struct {
// 	repository markdownRepository
// }

// func newPostController() postController {
// 	return postController{repository: newFileSystemMarkdownRepository()}
// }

// func (p postController) listAllPosts(w http.ResponseWriter) {
// 	tmpl := template.Must(template.ParseFiles("./views/all-posts.html"))
// 	tmpl.Execute(w, p.repository.getAll())
// }

// func (p postController) showPost(title string, w http.ResponseWriter) {
// 	tmpl := template.Must(template.ParseFiles("./views/single-post.html"))
// 	tmpl.Execute(w, p.repository.getByTitle(title))
// }
