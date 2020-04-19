package processors

import (
	"bytes"
	"text/template"

	"github.com/robbell/hi/app/io"
	"github.com/robbell/hi/app/markdown"
)

// SinglePost processor generates HTML pages for single posts
type SinglePost struct{}

// Process generates an HTML page for a single post
func (p SinglePost) Process(post markdown.Post) error {
	var postBuffer bytes.Buffer

	tmpl := template.Must(template.ParseFiles("./templates/single-post.html"))
	if err := tmpl.Execute(&postBuffer, post); err != nil {
		return err
	}

	writePath := "./static/" + post.Permalink + ".html"
	if err := io.WriteToDisk(writePath, postBuffer.String()); err != nil {
		return err
	}

	return nil
}

// Finish is a NOP for a single post
func (p SinglePost) Finish() error {
	return nil
}
