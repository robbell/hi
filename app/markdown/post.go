package markdown

import (
	"html/template"
	"strings"
	"time"

	"github.com/gernest/front"
	"github.com/robbell/hi/app/io"
	"gopkg.in/russross/blackfriday.v2"
)

// Post represents a blog post
type Post struct {
	Title       string
	Body        template.HTML
	PublishedOn time.Time
	Permalink   string
}

// ToHTMLPost converts a Markdown post to HTML
func ToHTMLPost(content string, sourcePath string) (Post, error) {
	fmHandler := front.NewMatter()
	fmHandler.Handle("---", front.YAMLHandler)
	frontMatter, body, err := fmHandler.Parse(strings.NewReader(content))
	if err != nil {
		return Post{}, err
	}

	publishedOn, err := time.Parse("2006-01-02", frontMatter["publishedOn"].(string))
	if err != nil {
		return Post{}, err
	}

	return Post{
			Title:       frontMatter["title"].(string),
			Body:        template.HTML(blackfriday.Run([]byte(body))),
			PublishedOn: publishedOn,
			Permalink:   io.ReplaceExtension(sourcePath, ""),
		},
		nil
}
