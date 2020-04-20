package models

import "github.com/robbell/hi/app/markdown"

// List represents a list of posts and a title
type List struct {
	Posts []markdown.Post
	Title string
}
