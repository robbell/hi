package processors

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/robbell/hi/app/io"
	"github.com/robbell/hi/app/markdown"
)

// TagList processor generates HTML lists of posts by tag
type TagList struct {
	tags map[string][]markdown.Post
}

// NewTagList creates a new TagList
func NewTagList() *TagList {
	tl := TagList{}
	tl.tags = make(map[string][]markdown.Post)
	return &tl
}

// Process collects information about posts to list
func (l *TagList) Process(post markdown.Post) error {
	for _, tag := range post.Tags {
		l.tags[tag] = append(l.tags[tag], post)
	}

	return nil
}

// Finish creates the listing page for posts
func (l *TagList) Finish() error {
	for tag, posts := range l.tags {
		var listBuffer bytes.Buffer

		tmpl := template.Must(template.ParseFiles("./templates/all-posts.html"))
		if err := tmpl.Execute(&listBuffer, posts); err != nil {
			return err
		}

		if err := io.WriteToDisk(fmt.Sprintf("./static/tags/%v.html", tag), listBuffer.String()); err != nil {
			return err
		}
	}

	return nil
}
