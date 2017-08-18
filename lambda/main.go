package main

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/gabesullice/pdfaf/lib/core"
)

type Options struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func Handle(evt json.RawMessage, _ *runtime.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var options Options
	if err := json.Unmarshal(evt, &options); err != nil {
		log.Println(err.Error())
		return err.Error(), err
	}

	log.Printf("Generating PDF for URL: %s\n", options.URL)

	// Initiate the request and get an io.Reader for the PDF
	pdf, err := toPDFReader(ctx, options)
	if err != nil {
		log.Println(err.Error())
		log.Println(options)
		return err.Error(), err
	}

	src, err := ioutil.ReadAll(pdf)
	if err != nil {
		log.Println(err.Error())
		return err.Error(), err
	}

	response := struct{ PDF []byte }{src}
	return response, nil
}

func toPDFReader(ctx context.Context, options Options) (io.Reader, error) {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
		core.ChromePath("./headless-chrome/headless_shell"),
	)

	if err != nil {
		return nil, err
	}

	return r, nil
}
