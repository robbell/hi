package main

import (
	"net/http"

	"github.com/google/go-github/github"
	"github.com/robbell/hi/app/hook"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) rebuild(_ http.ResponseWriter, r *http.Request) {

	pushEvent, err := hook.ValidatePushEvent(r)
	if err != nil {
		return // To do: log error or report failure
	}

	site := newSite()
	site.rebuild(pushEvent)
}

type site struct{}

func newSite() *site {
	return &site{}
}

func (s *site) rebuild(e *github.PushEvent) {
	r := newRepo(*e.Repo.Name)

	r.download()
}
