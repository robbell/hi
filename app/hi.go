package main

import (
	"flag"
	"log"
	"os"

	"github.com/robbell/hi/processors"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "build" {
		buildFlags := flag.NewFlagSet("build", flag.ExitOnError)
		input := buildFlags.String("input", "", "content directory")
		buildFlags.Parse(os.Args[2:])

		if *input == "" {
			log.Fatal("missing --input")
		}

		if err := os.RemoveAll("./static"); err != nil {
			log.Fatal(err)
		}

		if err := os.MkdirAll("./static", os.ModePerm); err != nil {
			log.Fatal(err)
		}

		repo := newLocalRepo(*input)
		if err := repo.process(&processors.Post{}, &processors.Index{}, &processors.Resources{}, processors.NewTags()); err != nil {
			log.Fatal(err)
		}

		return
	}

	s := server{}
	s.Start()
}
