package processors

import "github.com/robbell/hi/markdown"

// Processor interface for the site rebuild pipeline
type Processor interface {
	Process(markdown.Post) error
	Finish() error
}
