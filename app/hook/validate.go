package hook

import (
	"net/http"

	"github.com/google/go-github/github"
)

// ValidatePushEvent returns a valid PushEvent or nil if invalid
func ValidatePushEvent(r *http.Request) (e *github.PushEvent, err error) {

	payload, err := github.ValidatePayload(r, []byte("[Replaced]")) // To do: replace with env variable
	if err != nil {
		return nil, err
	}

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		return nil, err
	}

	return event.(*github.PushEvent), nil
}
