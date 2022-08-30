package main

import (
	"log"
	"service-photo-management/service"
	"time"
)

/*
* This Service Parses through my Photography Collection, Finds the Pictures, Creates an Album Data Struct
* that contains the Filename, Type, Path, & Folder Name. Sorts them and uploads them into Google Photos.
* It also allows me to analyze the different types of files I have available recursively in a directory.
* It checks google photos to see if there are new changes to upload.
* It allows me to discover which files are raw or nef that are unedited and remaining to be completed.
 */

func main() {
	timeTracker("Folder Analyzer", service.AnalyzeFolder)
}

func timeTracker(serviceName string, fn func()) {
	startTime := time.Now()
	log.Printf("Starting %s: %s", serviceName, startTime.String())

	fn()

	endTime := time.Now()
	log.Printf("Ending %s: %s", serviceName, endTime.String())

	elapsedTime := time.Since(startTime)
	log.Printf("Elapsed Time: %d", elapsedTime)
}

/* TODO
* func viewEditList() {}
* func createAlbum() {}
* func uploadExportedFiles() {}
 */
