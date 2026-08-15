package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/robbell/hi/processors"
)

type localRepo struct {
	root string
}

func newLocalRepo(root string) *localRepo {
	return &localRepo{root: root}
}

func (r *localRepo) process(processors ...processors.Processor) error {
	err := filepath.Walk(r.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == ".github" {
				return filepath.SkipDir
			}

			return nil
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(r.root, path)
		if err != nil {
			return err
		}

		rel = filepath.ToSlash(rel)

		for _, processor := range processors {
			if err := processor.Process(string(content), rel); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	for _, processor := range processors {
		if err := processor.Finish(); err != nil {
			return err
		}
	}

	return nil
}
