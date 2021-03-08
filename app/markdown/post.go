package markdown

import (
	"html/template"
	"strings"
	"time"

	"github.com/gernest/front"
	"github.com/robbell/hi/io"
	"gopkg.in/russross/blackfriday.v2"
)

const summaryBreak = "[!EndSummary]"

// Post represents a blog post
type Post struct {
	Title       string
	Summary     template.HTML
	Body        template.HTML
	PublishedOn time.Time
	Permalink   string
	Tags        []string
	IsListed    bool
}

// ToHTMLPost converts a Markdown post to HTML
func ToHTMLPost(content string, sourcePath string) (Post, error) {
	fmHandler := front.NewMatter()
	fmHandler.Handle("---", front.YAMLHandler)
	frontMatter, markdown, err := fmHandler.Parse(strings.NewReader(content))

	if err != nil {
		return Post{}, err
	}

	publishedOn, err := time.Parse("2006-01-02", frontMatter["publishedOn"].(string))
	if err != nil {
		return Post{}, err
	}

	return Post{
			Title:       frontMatter["title"].(string),
			Summary:     getSummary(markdown),
			Body:        getBody(markdown),
			PublishedOn: publishedOn,
			Permalink:   getPermalink(sourcePath),
			Tags:        getTags(frontMatter),
			IsListed:    getIsListed(frontMatter),
		},
		nil
}

func getSummary(markdown string) template.HTML {

	if i := strings.Index(markdown, summaryBreak); i > 0 {
		markdown = markdown[:i]
	}

	return template.HTML(blackfriday.Run([]byte(markdown)))
}

func getBody(markdown string) template.HTML {
	sanitised := strings.Replace(markdown, summaryBreak, "", -1)
	return template.HTML(blackfriday.Run([]byte(sanitised)))
}

func getPermalink(sourcePath string) string {
	if !strings.HasPrefix(sourcePath, "/") {
		sourcePath = "/" + sourcePath
	}

	return io.ReplaceExtension(sourcePath, "")
}

func getTags(frontMatter map[string]interface{}) []string {

	if tags, found := frontMatter["tags"]; found {
		tagInterface := tags.([]interface{})
		tags := make([]string, len(tagInterface))
		for key, tag := range tagInterface {
			tags[key] = tag.(string)
		}

		return tags
	}
	return nil
}

func getIsListed(frontMatter map[string]interface{}) bool {

	if isListed, found := frontMatter["IsListed"]; found {
		return isListed.(bool)
	}
	return true
}
