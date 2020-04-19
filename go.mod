module github.com/robbell/hi

go 1.12

require (
	github.com/gernest/front v0.0.0-20181129160812-ed80ca338b88
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-github/v29 v29.0.3
	github.com/google/go-github/v31 v31.0.0
	github.com/gorilla/mux v1.7.4
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
