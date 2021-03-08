package main

import (
	"net/http"

	"github.com/robbell/hi/hook"
	"github.com/robbell/hi/processors"
)

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) rebuild(w http.ResponseWriter, r *http.Request) {

	pushEvent, err := hook.ValidatePushEvent(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return // To do: log error or report failure
	}

	repo := newRepo(*pushEvent.Repo.Owner.Name, *pushEvent.Repo.Name)
	if err = repo.process(&processors.Post{}, &processors.Index{}, &processors.Resources{}, processors.NewTags()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
