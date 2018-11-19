package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

const outputPath = "/Users/lee/Desktop/"

func getPhotos(path string) (list []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	list = make([]string, 0)
	for _, f := range files {
		if f.Name() != ".DS_Store" {
			list = append(list, f.Name())
		}
	}
	return
}

func getNameOfDir(path string) string {
	beginIndex := strings.LastIndexAny(path, "/")
	fmt.Println(path[beginIndex:])
	return path[beginIndex:] + ".pdf"
}

func main() {
	if len(os.Args) < 2 {
		panic("path is not defined")
	}
	path := string(os.Args[1])

	pdf := gofpdf.New("P", "mm", "A4", "") // A4 210.0 x 297.0
	// See documentation for details on how to generate fonts
	pdf.AddFont("Montserrat-ExtraLight", "", "Montserrat-ExtraLight.json")
	fontSize := 13.0
	pdf.SetFont("Montserrat-ExtraLight", "", fontSize)
	ht := pdf.PointConvert(fontSize)
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251") // "" defaults to "cp1252"
	write := func(image string, isNewPage bool) {
		if isNewPage {
			pdf.MultiCell(190, ht, tr(image), "", "C", false)
			pdf.Image(path+"/"+image, 25, 40, 0, 100, false, "", 0, "")
		} else {
			pdf.MoveTo(0, 160)
			pdf.MultiCell(190, ht, tr(image), "", "C", false)
			pdf.Image(path+"/"+image, 25, 180, 0, 100, false, "", 0, "")
		}
	}

	photos := getPhotos(path + "/")
	for index, photo := range photos {
		if index%2 == 0 {
			pdf.AddPage()
			write(photo, true)
		} else {
			write(photo, false)
		}
	}
	nameOfReport := outputPath + getNameOfDir(path)
	fmt.Println(nameOfReport)
	err := pdf.OutputFileAndClose(nameOfReport)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Report was successfully created!")
	}
}
