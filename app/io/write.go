package io

import (
	"io"
	"os"
	"path/filepath"
)

// WriteToDisk writes content to disk, creating the file and folder structure if they don't exist
func WriteToDisk(path string, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err = io.WriteString(file, content); err != nil {
		return err
	}

	if err = file.Sync(); err != nil {
		return err
	}

	return nil
}
