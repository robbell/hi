package processors

import (
	"strings"

	"github.com/robbell/hi/io"
)

// Resources processor merges resource files
type Resources struct{}

// Process writes resource files unchanged
func (r Resources) Process(content string, sourcePath string) error {

	if strings.HasSuffix(sourcePath, ".md") {
		return nil
	}

	writePath := "./static/" + sourcePath
	if err := io.WriteToDisk(writePath, content); err != nil {
		return err
	}

	return nil
}

// Finish is a NOP for resources
func (r Resources) Finish() error {
	return nil
}
