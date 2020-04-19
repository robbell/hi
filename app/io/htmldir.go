package io

import (
	"net/http"
	"strings"
)

// FriendlyFileSystem provides friendly URLs for HTML file access
type FriendlyFileSystem struct {
	Fs http.FileSystem
}

// Open wrapper to http.FileSystem
func (ffs FriendlyFileSystem) Open(path string) (http.File, error) {

	if !strings.HasSuffix(path, ".html") {
		if f, err := ffs.Fs.Open(path + ".html"); err == nil {
			return f, nil
		}
	}

	f, err := ffs.Fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := strings.TrimSuffix(path, "/") + "/index.html"
		if _, err := ffs.Fs.Open(index); err != nil {
			return nil, err
		}
	}

	return f, nil
}
