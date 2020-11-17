module github.com/robbell/hi

go 1.15

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1

require (
	github.com/gernest/front v0.0.0-20181129160812-ed80ca338b88
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v31 v31.0.0
	github.com/gorilla/mux v1.8.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
)
