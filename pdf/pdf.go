package pdf

import (
	"photoReporter/image"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

// CreatePDFReport creates a pdf file with signs and photos.
func CreatePDFReport(pathToDir, pathToSave string) error {
	pdfInstance := gofpdf.New("P", "mm", "A4", "")
	pdfInstance.AddFont("Montserrat-ExtraLight", "", "./fonts/Montserrat-ExtraLight.json")
	fontSize := 13.0
	pdfInstance.SetFont("Montserrat-ExtraLight", "", fontSize)
	ht := pdfInstance.PointConvert(fontSize)
	tr := pdfInstance.UnicodeTranslatorFromDescriptor("./fonts/cp1251") // "" defaults to "cp1252"

	write := func(pathToImage, imageName string, isNewPage bool) {
		if isNewPage {
			pdfInstance.MultiCell(190, ht, tr(imageName), "", "C", false)
			pdfInstance.Image(pathToImage, 25, 40, 0, 100, false, "", 0, "")
		} else {
			pdfInstance.MoveTo(0, 160)
			pdfInstance.MultiCell(190, ht, tr(imageName), "", "C", false)
			pdfInstance.Image(pathToImage, 25, 180, 0, 100, false, "", 0, "")
		}
	}

	photos := image.GetPhotos(pathToDir)
	for index, photo := range photos {
		if index%2 == 0 {
			pdfInstance.AddPage()
			write(photo.Path, photo.Name, true)
		} else {
			write(photo.Path, photo.Name, false)
		}
	}

	lastSlashIndex := strings.LastIndexAny(pathToDir, "/") + 1
	resultFileName := pathToDir[lastSlashIndex:] + ".pdf"
	err := pdfInstance.OutputFileAndClose(pathToSave + "/" + resultFileName)
	if err != nil {
		return err
	}

	return nil
}
