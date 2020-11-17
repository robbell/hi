package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robbell/hi/io"
)

const (
	port   = 8080
	banner = `
  _     _        __
 | |__ (_)  (( ,\ \ \  ))
 | '_ \| |    ,\ \ \ \  ,
 | | | | | _  \ ' ' ' \/ )
 |_| |_|_|(_)  \        /
--------------------------
%s%v
`
)

type server struct{}

func (s server) Start() {
	router := mux.NewRouter()
	handler := newHandler()

	router.StrictSlash(true)
	router.HandleFunc("/hooks/rebuild", handler.rebuild)
	router.PathPrefix("/").Handler(http.FileServer(io.FriendlyFileSystem{Fs: http.Dir("./static/")}))

	fmt.Printf(banner, "Hi server is listening on http://localhost:", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), router); err != nil {
		panic(err)
	}
}
