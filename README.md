# Hi 👋

A developer-friendly static site generator, written in Go, backed by GitHub.

![Sample Component diagram](docs/hi-component.svg)

## To Do

### Features

- [x] Serve static files
- [x] Add GitHub hook to notify of changes
- [x] Define post Markdown format
- [x] Static page generation
  - [x] Templates
  - [x] Single pages
  - [x] List pages (homepage, tags)
- [ ] Helm charts
  - [x] Persistent storage for generated site
  - [ ] Secret management
- [ ] Tests! (Probably to come later as I'm still learning Go)
- [ ] Sorting by date on tag lists
- [ ] Templates stored in source repo

### Development practices

- [ ] Should I copy the application to the root of my Docker image?
- [ ] Optimise Dockerfile - seems to be rebuilding when later layers are changed
- [ ] Understand Go testing
- [x] Understand the details of Go Modules

## Future
- [ ] Micro browser support (https://24ways.org/2019/microbrowsers-are-everywhere/)
- [ ] Medium publishing
- [ ] Pipeline for site generation
  - [x] HTML page building
  - [ ] AMP page building
  - [ ] List building
    - [x] Index
    - [ ] By tag
    - [ ] By type
  - [ ] RSS building
  - [ ] Related pages
- [x] File extension removal

## Notes

To publish a new version of the app, build and push to Docker Hub using the below, incrementing the image tag:

```
docker buildx build --platform linux/amd64 -t robbell/hi:latest -t robbell/hi:1.1.x --push .
```

After publishing, update the image tag reference for the Container App to match.
