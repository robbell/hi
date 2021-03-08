package processors

import (
	"bytes"
	"sort"
	"strings"
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
func (i *Index) Process(content string, sourcePath string) error {

	if !strings.HasSuffix(sourcePath, ".md") {
		return nil
	}

	post, err := markdown.ToHTMLPost(content, sourcePath)
	if err != nil {
		return err
	}

	if post.IsListed {
		i.posts = append(i.posts, post)
	}
	return nil
}

// Finish creates the index page
func (i *Index) Finish() error {
	sort.Sort(ByDate(i.posts))

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

type ByDate []markdown.Post

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Less(i, j int) bool { return a[i].PublishedOn.After(a[j].PublishedOn) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
