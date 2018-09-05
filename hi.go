package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
