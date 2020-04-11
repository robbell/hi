// package main

// import (
// 	"io/ioutil"
// 	"path"
// )

// type markdownRepository interface {
// 	getAll() []markdownPost
// 	getByTitle(string) markdownPost
// }

// type fileSystemMarkdownRepository struct {
// 	sourceDirectory string
// }

// func newFileSystemMarkdownRepository() markdownRepository {
// 	return fileSystemMarkdownRepository{
// 		sourceDirectory: "./static",
// 	}
// }

// func (s fileSystemMarkdownRepository) getAll() []markdownPost {
// 	return getFilesInDirectory(s.sourceDirectory)
// }

// func getFilesInDirectory(directory string) []markdownPost {
// 	files, err := ioutil.ReadDir(directory)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var posts []markdownPost

// 	for _, directoryItem := range files {

// 		if directoryItem.IsDir() {
// 			posts = append(posts, getFilesInDirectory(path.Join(directory, directoryItem.Name()))...)
// 		} else if post, ok := newMarkdownPost(path.Join(directory, directoryItem.Name())); ok {
// 			posts = append(posts, post)
// 		}
// 	}

// 	return posts
// }

// func (s fileSystemMarkdownRepository) getByTitle(title string) markdownPost {
// 	post, _ := newMarkdownPost(path.Join(s.sourceDirectory, title+".md"))
// 	return post
// }
