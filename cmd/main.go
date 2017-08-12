package main

import (
	"github.com/gabesullice/pdfaf/lib/core"
	"os"
)

func main() {
	core.RequestToPDF(core.URL(os.Args[1]))
}
