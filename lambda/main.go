package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/gabesullice/pdfaf/lib/core"
)

type Options struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

func main() {}

func Handler(evt json.RawMessage, _ *runtime.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var options Options
	if err := json.Unmarshal(evt, options); err != nil {
		return err.Error(), err
	}

	// Initiate the request and get an io.Reader for the PDF
	pdf, err := toPDFReader(ctx, options)
	if err != nil {
		return err.Error(), err
	}

	src, err := ioutil.ReadAll(pdf)
	if err != nil {
		return err.Error(), err
	}

	var enc []byte
	base64.StdEncoding.Encode(enc, src)
	return enc, nil
}

func toPDFReader(ctx context.Context, options Options) (io.Reader, error) {
	r, err := core.ToPDF(
		ctx,
		core.URL(options.URL),
		core.Headers(options.Headers),
		core.ChromePath("./chrome/chrome-linux/chrome"),
	)

	if err != nil {
		return nil, err
	}

	return r, nil
}
