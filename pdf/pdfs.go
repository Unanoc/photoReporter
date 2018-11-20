package pdf

import (
	"fmt"
	"io/ioutil"
	"log"
)

// CreatePDFReports creates a pdf files with signs and photos for all folders.
func CreatePDFReports(pathToDir, pathToSave string) error {
	dirs := GetFolders(pathToDir)
	fmt.Println(dirs)

	for _, dir := range dirs {
		fullPathToPhotoDir := pathToDir + "/" + dir
		err := CreatePDFReport(fullPathToPhotoDir, pathToSave)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetFolders returns a slice of folders.
func GetFolders(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	dirs := make([]string, 0)
	for _, f := range files {
		if f.IsDir() && f.Name() != "$RECYCLE.BIN" {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}
