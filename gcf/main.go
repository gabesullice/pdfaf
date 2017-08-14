package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gabesullice/pdfaf/lib/core"
)

type Options struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start a Chrome process
	go func() { runChrome(ctx) }()

	decoder := json.NewDecoder(base64.NewDecoder(base64.StdEncoding, r.Body))
	var options Options
	if err := decoder.Decode(&options); err != nil {
		log.Printf("Options: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Initiate the request and get an io.Reader for the PDF
	pdf, err := toPDFReader(ctx, options)
	if err != nil {
		log.Printf("To PDF: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := copyToFile(pdf, "./output.pdf"); err != nil {
		log.Printf("Copy: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Success!")

	// Copy the PDF into the response
	//if _, err = io.Copy(w, pdf); err != nil {
	//	log.Printf("Copy: %s\n", err)
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
}

func toPDFReader(ctx context.Context, options Options) (io.Reader, error) {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
	)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func copyToFile(r io.Reader, filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	if _, err = io.Copy(file, r); err != nil {
		return err
	}
	return nil
}

func runChrome(ctx context.Context) {
	cmd := exec.CommandContext(ctx, "chromium", "--headless", "--remote-debugging-port=9222", "--disable-gpu")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
