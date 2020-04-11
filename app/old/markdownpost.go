// package main

// import (
// 	"html/template"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"github.com/gernest/front"
// 	"gopkg.in/russross/blackfriday.v2"
// )

// type markdownPost struct {
// 	Title       string
// 	Body        template.HTML
// 	PublishedOn time.Time
// 	Permalink   string
// }

// func newMarkdownPost(path string) (markdownPost, bool) {
// 	fmHandler := front.NewMatter()
// 	fmHandler.Handle("---", front.YAMLHandler)

// 	if filepath.Ext(path) != ".md" {
// 		return markdownPost{}, false
// 	}

// 	file, err := os.Open(path)

// 	if err != nil {
// 		return markdownPost{}, false
// 	}

// 	frontMatter, body, err := fmHandler.Parse(file)

// 	if err != nil {
// 		panic(err)
// 	}

// 	publishedOn, _ := time.Parse("2006-01-02", frontMatter["publishedOn"].(string))

// 	return markdownPost{
// 			Title:       frontMatter["title"].(string),
// 			Body:        template.HTML(blackfriday.Run([]byte(body))),
// 			PublishedOn: publishedOn,
// 			Permalink:   frontMatter["permalink"].(string),
// 		},
// 		true
// }
