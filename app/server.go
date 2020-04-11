package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	banner = `
  _     _        __
 | |__ (_)  (( ,\ \ \  ))
 | '_ \| |    ,\ \ \ \  ,
 | | | | | _  \ ' ' ' \/ )
 |_| |_|_|(_)  \        /
--------------------------
%s
`
)

type server struct{}

func (s server) Start() {
	router := mux.NewRouter()
	handler := newHandler()

	router.StrictSlash(true)
	router.HandleFunc("/hooks/rebuild", handler.rebuild)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./images/")))

	fmt.Printf(banner, "Hi server is listening on http://localhost:80")

	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
