package analyze

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	c "service-photo-management/constants"
	m "service-photo-management/model"
	"service-photo-management/util"
)

func fileSetter(file *os.File, dirName string) *m.FILE {
	info, _ := os.Stat(file.Name())
	contentType := strings.ToUpper(filepath.Ext(file.Name()))
	parentPath := filepath.Dir(file.Name())

	isVideo := util.Contains(contentType, c.VIDEO_FORMATS)
	length := 0.0
	if isVideo {
		length = 0.0
	}

	newFile := &m.FILE{
		DATE_MODIFIED: info.ModTime().String(),
		NAME:          file.Name(),
		PATH:          fmt.Sprintf("%s/%s", parentPath, file.Name()),
		ABLUM_NAME:    dirName,
		FILE_TYPE:     contentType,
		SIZE:          info.Size(),
		LENGTH:        length,
	}

	return newFile
}

func TraverseDirectory(directory []fs.DirEntry, dirName string, media []m.FILE) []m.FILE {
	log.Printf("Traversing Directory: %s", dirName)
	directoryDone := make(chan []m.FILE)
	fileDone := make(chan []m.FILE)

	for _, file := range directory {

		if file.IsDir() {
			log.Print("Go Routine started for Directory! ")

			go func() { directoryDone <- handleDirectory(file) }()
			media = append(media, <-directoryDone...)
		} else {
			log.Print("Go Routine started for File! ")

			go func() { fileDone <- handleFile(file, dirName) }()
			media = append(media, <-fileDone...)
		}
	}
	return media
}

func handleDirectory(file fs.DirEntry) []m.FILE {
	newDirectory := util.OpenDirectory(file.Name())
	util.GoToChild(fmt.Sprintf("./%s", file.Name()))

	media := []m.FILE{}
	media = TraverseDirectory(newDirectory, file.Name(), media)
	return media
}

func handleFile(file fs.DirEntry, dirName string) []m.FILE {
	openedFile := util.FileOpen(fmt.Sprintf("./%s", file.Name()))
	fileDetails := fileSetter(openedFile, dirName)

	media := []m.FILE{}
	media = append(media, *fileDetails)

	defer openedFile.Close()

	return media
}
