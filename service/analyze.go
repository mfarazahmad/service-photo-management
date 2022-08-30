package service

import (
	"log"
	"os"
	m "service-photo-management/model"
	"service-photo-management/service/analyze"
)

var (
	ROOT_DIR string = "C://Users/faraz/Pictures"

	MASTER_LIST []m.ALBUM
)

func AnalyzeFolder() {
	os.Chdir(ROOT_DIR)
	directory, _ := os.ReadDir(ROOT_DIR)
	log.Print(directory)

	media := []m.FILE{}
	media = analyze.TraverseDirectory(directory, ROOT_DIR, media)
	log.Print(media)

	// TODO: Add Logic for Figuring out What Layer?
	// ie: Category -> Album -> Pictures
}
