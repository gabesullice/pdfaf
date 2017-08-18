package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/gabesullice/pdfaf/lib/core"
)

type Options struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Parse arguments to configure request
	options := parseOptions()

	// Initiate the request and get an io.Reader for the PDF
	r := toPDFReader(ctx, options)

	// Copy the PDF to a file
	copyToFile(r, "./output.pdf")
}

func copyToFile(r io.Reader, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = io.Copy(file, r); err != nil {
		log.Fatalln(err)
	}
}

func parseOptions() (options Options) {
	if err := json.Unmarshal([]byte(os.Args[1]), &options); err != nil {
		log.Fatalln(err)
	}
	return options
}

func toPDFReader(ctx context.Context, options Options) io.Reader {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
		core.ChromePath("chromium"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	return r
}
