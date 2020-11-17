package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v31/github"
	"github.com/robbell/hi/markdown"
	"github.com/robbell/hi/processors"
	"golang.org/x/oauth2"
)

type repo struct {
	Owner string
	Name  string
}

func newRepo(owner string, name string) *repo {
	return &repo{Owner: owner, Name: name}
}

func (s *repo) process(processors ...processors.Processor) error {
	context := context.Background()

	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "[Replaced]"}, // To do: replace with env variable
	)
	tokenClient := oauth2.NewClient(context, tokenSource)
	client := github.NewClient(tokenClient)

	err := s.processDir(context, client, "/", processors...)

	for _, processor := range processors {
		processor.Finish() // To do: error aggregation for processors
	}

	return err
}

func (s *repo) processDir(ctx context.Context, client *github.Client, path string, processors ...processors.Processor) error {
	_, directoryContents, r, err := client.Repositories.GetContents(ctx, s.Owner, s.Name, path, nil)
	if err != nil {
		return err
	} else if r.StatusCode != 200 {
		return fmt.Errorf("%d status code returned", r.StatusCode)
	}

	for _, directoryContent := range directoryContents {
		if *directoryContent.Type == "dir" {
			s.processDir(ctx, client, *directoryContent.Path, processors...)
		} else {
			s.processFile(ctx, client, *directoryContent.Path, processors...)
		}
	}

	return nil
}

func (s *repo) processFile(ctx context.Context, client *github.Client, sourcePath string, processors ...processors.Processor) error {
	metaContent, _, r, err := client.Repositories.GetContents(ctx, s.Owner, s.Name, sourcePath, nil)
	if err != nil {
		return err
	} else if r.StatusCode != 200 {
		return fmt.Errorf("%d status code returned", r.StatusCode)
	}

	content, err := metaContent.GetContent()
	if err != nil {
		return err
	}

	post, err := markdown.ToHTMLPost(content, sourcePath)
	if err != nil {
		return err
	}

	for _, processor := range processors {
		processor.Process(post) // To do: error aggregation for processors
	}

	return nil
}
