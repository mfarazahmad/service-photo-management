package util

import (
	"io/fs"
	"log"
	"os"
)

func FileOpen(file string) *os.File {
	log.Printf("Opening File %s", file)
	openedFile, _ := os.Open(file)
	return openedFile
}

func FileReader(file string) {
	log.Printf("Reading File %s", file)
	document, _ := os.ReadFile(file)
	log.Print(document)
}

func OpenDirectory(currentPath string) []fs.DirEntry {
	directories, _ := os.ReadDir(currentPath)
	return directories
}

func GoToParent(path string) {
	log.Printf("Going to Parent: %s", path)
	os.Chdir(path)
}

func GoToChild(path string) {
	log.Printf("Going to SubFolder: %s", path)
	os.Chdir(path)
}

func IsDirectoryComplete() bool {
	return true
}

func WrapError(err error) {
	log.Print(err)
}

func Contains(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
