package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gernest/front"
)

type markdownPost struct {
	Title       string
	Body        string
	PublishedOn time.Time
	Permalink   string
}

func (m markdownPost) bodyAsHTML() string {
	return "<p>" + m.Body + "</p>"
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
			Body:        body,
			PublishedOn: publishedOn,
			Permalink:   frontMatter["permalink"].(string),
		},
		true
}
