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

func main() {
	runChrome(context.TODO())

	file, err := os.Create("./output.pdf")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	headerMap := make(map[string]string)
	if err := json.Unmarshal([]byte(os.Args[2]), &headerMap); err != nil {
		panic(err)
	}
	r, err := core.RequestToPDF(
		core.URL(os.Args[1]),
		core.Headers(headerMap),
	)

	if err != nil {
		panic(err)
	}

	if _, err = io.Copy(file, r); err != nil {
		panic(err)
	}
}

func runChrome(ctx context.Context) {
	go func() {
		cmd := exec.CommandContext(ctx, "chromium", "--headless", "--remote-debugging-port=9222", "--disable-gpu")
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(500 * time.Millisecond)
}
