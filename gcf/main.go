package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

	decoder := json.NewDecoder(r.Body)
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

	encoder := base64.NewEncoder(base64.StdEncoding, w)
	defer encoder.Close()
	// Copy the PDF into the response
	if _, err = io.Copy(encoder, pdf); err != nil {
		log.Printf("Copy: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func toPDFReader(ctx context.Context, options Options) (io.Reader, error) {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
		core.ChromePath("./chrome-linux/chrome"),
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
