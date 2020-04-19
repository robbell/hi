package io

import "path"

// ReplaceExtension replaces one extension wth another
func ReplaceExtension(sourcePath string, newExtension string) string {
	currentExtension := path.Ext(sourcePath)
	return sourcePath[0:len(sourcePath)-len(currentExtension)] + newExtension
}
