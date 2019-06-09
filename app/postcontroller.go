package main

import (
	"html/template"
	"net/http"
)

type postController struct{}

func (p postController) listAllPosts(w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("./views/all-posts.html"))
	repository := fileSystemMarkdownRepository{"./static"}
	tmpl.Execute(w, repository.getAll())
}

func (p postController) showPost(title string, w http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("./views/single-post.html"))
	repository := fileSystemMarkdownRepository{"./static"}
	tmpl.Execute(w, repository.getByTitle(title))
}
