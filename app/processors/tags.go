package processors

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/robbell/hi/io"
	"github.com/robbell/hi/markdown"
	"github.com/robbell/hi/models"
)

// Tags processor generates HTML lists of posts by tag
type Tags struct {
	tags map[string][]markdown.Post
}

// NewTags creates a new Tags processor
func NewTags() *Tags {
	tl := Tags{}
	tl.tags = make(map[string][]markdown.Post)
	return &tl
}

// Process collects information about posts to list by tag
func (t *Tags) Process(post markdown.Post) error {
	for _, tag := range post.Tags {
		t.tags[tag] = append(t.tags[tag], post)
	}

	return nil
}

// Finish creates the listing page for posts by tag
func (t *Tags) Finish() error {
	for tag, posts := range t.tags {
		var listBuffer bytes.Buffer

		tmpl := template.Must(template.ParseFiles("./templates/list.html", "./templates/base.html"))
		if err := tmpl.Execute(&listBuffer, models.List{Title: fmt.Sprintf("Posts tagged with \"%v\"", tag), Posts: posts}); err != nil {
			return err
		}

		if err := io.WriteToDisk(fmt.Sprintf("./static/tags/%v.html", tag), listBuffer.String()); err != nil {
			return err
		}
	}

	return nil
}
