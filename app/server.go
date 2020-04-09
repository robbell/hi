package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{}

func (s server) Start() {
	router := mux.NewRouter()
	handler := newHandler()

	router.HandleFunc("/hooks/rebuild", handler.rebuild)

	http.FileServer(http.Dir("./static"))

	fmt.Printf("👋 Hi server is listening on http://localhost:80")

	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
