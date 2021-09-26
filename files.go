package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func getAllFiles(dirPath string, extension string) []string {
	var dirFiles []string

	err := filepath.WalkDir(dirPath, func(path string, entry fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if !entry.IsDir() && filepath.Ext(path) == extension {
			dirFiles = append(dirFiles, path)
		}
		return nil
	})

	if err != nil {
		return nil
	}

	return dirFiles
}

func fileSubstrCount(filePath string, search string) int {
	b, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Print(err)
	}
	content := string(b)

	return strings.Count(strings.ToLower(content), strings.ToLower(search))
}
