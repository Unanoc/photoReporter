package main

import (
	"fmt"
	"log"
	"os"
	"photoReporter/pdf"
	"photoReporter/valid"
)

var (
	instructionText = fmt.Sprintf("-all [path to dir with photo dirs] [outdir]\n-u [path to dir with photo] [outdir]\n")
)

func main() {
	if len(os.Args) != 4 {
		fmt.Printf(instructionText)
		return
	}

	var err error

	pathToDir, err := valid.GetValidPath(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	pathToSave, err := valid.GetValidPath(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "-all":
		err = pdf.CreatePDFReports(pathToDir, pathToSave)
	case "-u":
		err = pdf.CreatePDFReport(pathToDir, pathToSave)
	default:
		fmt.Printf(instructionText)
		return
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
}
