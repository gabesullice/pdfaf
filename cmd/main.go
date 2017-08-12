package main

import (
	"github.com/gabesullice/pdfaf/lib/core"
	"os"
)

func main() {
	file, err := os.Create("./output.pdf")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	r, err := core.RequestToPDF(core.URL(os.Args[1]))
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(file, r); err != nil {
		panic(err)
	}
}
