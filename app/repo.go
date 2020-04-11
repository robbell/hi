package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
)

type repo struct {
	Owner string
	Name  string
}

func newRepo(owner string, name string) *repo {
	return &repo{Owner: owner, Name: name}
}

func (s *repo) download(e *github.PushEvent) {
	context := context.Background()
	client := github.NewClient(nil)
	s.downloadDir(context, client, "/")
}

func (s *repo) downloadDir(ctx context.Context, client *github.Client, path string) error {
	_, directoryContents, r, err := client.Repositories.GetContents(ctx, s.Owner, s.Name, path, nil)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return fmt.Errorf("%d status code returned", r.StatusCode)
	}

	for _, directoryContent := range directoryContents {
		if *directoryContent.Type == "dir" {
			s.downloadDir(ctx, client, *directoryContent.Path)
		} else {
			s.downloadFile(ctx, client, *directoryContent.Path)
		}
	}

	return nil
}

func (s *repo) downloadFile(ctx context.Context, client *github.Client, path string) error {
	content, _, r, err := client.Repositories.GetContents(ctx, s.Owner, s.Name, path, nil)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return fmt.Errorf("%d status code returned", r.StatusCode)
	}

	output, _ := content.GetContent()
	path = "./static/" + path

	folderPath := filepath.Dir(path)
	os.MkdirAll(folderPath, os.ModePerm)

	localFile, _ := os.Create(path)
	io.WriteString(localFile, output)
	localFile.Sync()

	return nil
}
