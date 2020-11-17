package processors

import (
	"bytes"
	"text/template"

	"github.com/robbell/hi/io"
	"github.com/robbell/hi/markdown"
	"github.com/robbell/hi/models"
)

// Index processor generates the HTML index page
type Index struct {
	posts []markdown.Post
}

// Process collects information about posts to display on index
func (i *Index) Process(post markdown.Post) error {
	i.posts = append(i.posts, post)
	return nil
}

// Finish creates the index page
func (i *Index) Finish() error {
	var listBuffer bytes.Buffer

	tmpl := template.Must(template.ParseFiles("./templates/list.html", "./templates/base.html"))
	if err := tmpl.Execute(&listBuffer, models.List{Title: "", Posts: i.posts}); err != nil {
		return err
	}

	if err := io.WriteToDisk("./static/index.html", listBuffer.String()); err != nil {
		return err
	}

	return nil
}
