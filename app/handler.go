package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/robbell/hi/app/hook"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) rebuild(_ http.ResponseWriter, r *http.Request) {

	pushEvent, err := hook.ValidatePushEvent(r)
	if err != nil {
		return // To do: logging error or reporting failure
	}

	site := newSite()
	site.rebuild(pushEvent)
}



type site struct{}

func newSite() *site {
	return &site{}
}

func (s *site) rebuild(e *github.PushEvent) {

}

type source struct {
	Owner string
	Repo  string
}

func newSource(owner string, repo string) *source {
	return &source{Owner: owner, Repo: repo}
}

func (s *site) rebuild(e *github.PushEvent) {
 st
 
}

func downloadDirectory(ctx context.Context, client *github.Client, owner string, repo string, path string) {
	_, directoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, path, nil)

	for _, directoryContent := range directoryContents {
		if *directoryContent.Type == "dir" {
			downloadDirectory(ctx, client, owner, repo, *directoryContent.Path)
		} else {
			downloadFile(ctx, client, owner, repo, *directoryContent.Path)
		}
	}
}

func downloadFile(ctx context.Context, client *github.Client, owner string, repo string, path string) {
	content, _, _, _ := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	output, _ := content.GetContent()
	path = "./static/" + path

	folderPath := filepath.Dir(path)
	os.MkdirAll(folderPath, os.ModePerm)

	localFile, _ := os.Create(path)
	io.WriteString(localFile, output)
	localFile.Sync()
}