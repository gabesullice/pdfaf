package main

import (
	"context"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/gabesullice/pdfaf/lib/core"
)

type Options struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start a Chrome process
	go func() { runChrome(ctx) }()

	// Parse arguments to configure request
	options := parseOptions()

	// Sleep a little to ensure Chrome is started
	//time.Sleep(500 * time.Millisecond)

	// Initiate the request and get an io.Reader for the PDF
	r := toPDFReader(ctx, options)

	// Copy the PDF to a file
	copyToFile(r, "./output.pdf")
}

func copyToFile(r io.Reader, filename string) {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(file, r); err != nil {
		panic(err)
	}
}

func parseOptions() (options Options) {
	if err := json.Unmarshal([]byte(os.Args[1]), &options); err != nil {
		panic(err)
	}
	return options
}

func toPDFReader(ctx context.Context, options Options) io.Reader {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
	)

	if err != nil {
		panic(err)
	}

	return r
}

func runChrome(ctx context.Context) {
	cmd := exec.CommandContext(ctx, "chromium", "--headless", "--remote-debugging-port=9222", "--disable-gpu")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
