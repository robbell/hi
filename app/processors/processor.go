package processors

// Processor interface for the site rebuild pipeline
type Processor interface {
	Process(string, string) error
	Finish() error
}
