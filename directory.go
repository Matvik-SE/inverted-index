package main

import (
	"io/fs"
	"path/filepath"
)

var dirFiles []string

func getDirFiles(dirPath string) []string {
	err := filepath.WalkDir(dirPath, walkHandler)

	if err != nil {
		return nil
	}

	return dirFiles
}

func walkHandler(path string, entry fs.DirEntry, e error) error {
	if e != nil {
		return e
	}

	if !entry.IsDir() {
		dirFiles = append(dirFiles, path)
	}

	return nil
}
