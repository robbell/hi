package main

import (
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/gernest/front"
	"gopkg.in/russross/blackfriday.v2"
)

type markdownPost struct {
	Title       string
	Body        template.HTML
	PublishedOn time.Time
	Permalink   string
}

type markdownPostFactory struct {
	fmHandler *front.Matter
}

func newMarkdownPostFactory() markdownPostFactory {
	fmHandler := front.NewMatter()
	fmHandler.Handle("---", front.YAMLHandler)

	return markdownPostFactory{
		fmHandler: fmHandler,
	}
}

func (f markdownPostFactory) loadMarkdownPost(markdownPath string) (markdownPost, bool) {
	if filepath.Ext(markdownPath) != ".md" {
		return markdownPost{}, false
	}

	file, err := os.Open(markdownPath)

	if err != nil {
		return markdownPost{}, false
	}

	frontMatter, body, err := f.fmHandler.Parse(file)

	if err != nil {
		panic(err)
	}

	publishedOn, _ := time.Parse("2006-01-02", frontMatter["publishedOn"].(string))

	return markdownPost{
			Title:       frontMatter["title"].(string),
			Body:        template.HTML(blackfriday.Run([]byte(body))),
			PublishedOn: publishedOn,
			Permalink:   frontMatter["permalink"].(string),
		},
		true
}
