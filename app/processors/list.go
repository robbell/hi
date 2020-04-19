package processors

import (
	"bytes"
	"text/template"

	"github.com/robbell/hi/app/io"
	"github.com/robbell/hi/app/markdown"
)

// List processor generates HTML lists of posts
type List struct {
	posts []markdown.Post
}

// Process collects information about posts to list
func (l *List) Process(post markdown.Post) error {
	l.posts = append(l.posts, post)
	return nil
}

// Finish creates the listing page for posts
func (l *List) Finish() error {
	var listBuffer bytes.Buffer

	tmpl := template.Must(template.ParseFiles("./templates/all-posts.html"))
	if err := tmpl.Execute(&listBuffer, l.posts); err != nil {
		return err
	}

	if err := io.WriteToDisk("./static/index.html", listBuffer.String()); err != nil {
		return err
	}

	return nil
}
