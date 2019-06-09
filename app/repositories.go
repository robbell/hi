package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gernest/front"
)

type markdownRepository interface {
	getAll() []MarkdownPost
	getByTitle(string) MarkdownPost
}

type fileSystemMarkdownRepository struct {
	sourceDirectory string
}

func (s fileSystemMarkdownRepository) getAll() []MarkdownPost {
	files, err := ioutil.ReadDir(s.sourceDirectory)

	if err != nil {
		panic(err)
	}

	fmHandler := front.NewMatter()
	fmHandler.Handle("---", front.YAMLHandler)

	var posts []MarkdownPost

	for _, file := range files {
		if post, ok := s.loadMarkdownPost(file.Name(), fmHandler); ok {
			posts = append(posts, post)
		}
	}

	return posts
}

func (s fileSystemMarkdownRepository) getByTitle(title string) MarkdownPost {
	fmHandler := front.NewMatter()
	fmHandler.Handle("---", front.YAMLHandler)
	post, _ := s.loadMarkdownPost(title+".md", fmHandler)
	return post
}

func (s fileSystemMarkdownRepository) loadMarkdownPost(markdownPath string, m *front.Matter) (MarkdownPost, bool) {
	if filepath.Ext(markdownPath) != ".md" {
		return MarkdownPost{}, false
	}

	file, err := os.Open(path.Join(s.sourceDirectory, markdownPath))

	if err != nil {
		return MarkdownPost{}, false
	}

	frontMatter, body, err := m.Parse(file)

	if err != nil {
		panic(err)
	}

	publishedOn, _ := time.Parse("2006-01-02", frontMatter["publishedOn"].(string))

	return MarkdownPost{
			Title:       frontMatter["title"].(string),
			Body:        body,
			PublishedOn: publishedOn,
			Permalink:   frontMatter["permalink"].(string),
		},
		true
}

type MarkdownPost struct {
	Title       string
	Body        string
	PublishedOn time.Time
	Permalink   string
}
