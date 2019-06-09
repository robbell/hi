package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	pc := postController{}

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		pc.listAllPosts(writer)
	})

	router.HandleFunc("/{year}/{month}/{title}", func(writer http.ResponseWriter, request *http.Request) {
		pc.showPost(mux.Vars(request)["title"], writer)
	})

	fmt.Printf("👋 Hi is listening on http://localhost:80")

	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}
