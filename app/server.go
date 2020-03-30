package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v29/github"
	"github.com/gorilla/mux"
)

type server struct{}

func (s server) Start() {
	router := mux.NewRouter()
	controller := newPostController()

	router.HandleFunc("/hooks/rebuild", handleWebhook)

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		controller.listAllPosts(writer)
	})

	router.HandleFunc("/{year}/{month}/{title}", func(writer http.ResponseWriter, request *http.Request) {
		controller.showPost(mux.Vars(request)["title"], writer)
	})

	fmt.Printf("👋 Hi server is listening on http://localhost:80")

	if err := http.ListenAndServe(":80", router); err != nil {
		panic(err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, _ := github.ValidatePayload(r, []byte("smelly")) // To do: replace with env variable
	event, _ := github.ParseWebHook(github.WebHookType(r), payload)
	pushEvent := event.(*github.PushEvent)

	ctx := context.Background()
	client := github.NewClient(nil)
	downloadDirectory(ctx, client, *pushEvent.Repo.Owner.Name, *pushEvent.Repo.Name, "/")
}

func downloadDirectory(ctx context.Context, client *github.Client, owner string, repo string, path string) {
	_, directoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, path, nil)

	for _, directoryContent := range directoryContents {
		if *directoryContent.Type == "dir" {
			downloadDirectory(ctx, client, owner, repo, *directoryContent.Path)
		} else {
			downloadFile(ctx, client, owner, repo, *directoryContent.Path)
		}
	}
}

func downloadFile(ctx context.Context, client *github.Client, owner string, repo string, path string) {
	content, _, _, _ := client.Repositories.GetContents(ctx, owner, repo, path, nil)
	output, _ := content.GetContent()
	path = "./static/" + path

	folderPath := filepath.Dir(path)
	os.MkdirAll(folderPath, os.ModePerm)

	localFile, _ := os.Create(path)
	io.WriteString(localFile, output)
	localFile.Sync()
}
