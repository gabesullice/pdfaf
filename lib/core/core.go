package core

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

type Options struct {
	URL     string
	Headers []byte
}

type OptionFunc func(Options) Options

func URL(url string) OptionFunc {
	return func(o Options) Options {
		o.URL = url
		return o
	}
}

func Headers(headers map[string]string) OptionFunc {
	return func(o Options) Options {
		o.Headers, _ = json.Marshal(headers)
		return o
	}
}

func RequestToPDF(ops ...OptionFunc) (io.Reader, error) {
	opts := applyOptions(Options{}, ops...)

	ctx := context.TODO()

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New("http://127.0.0.1:9222")
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Initiate a new RPC connection to the Chrome Debugging Protocol target.
	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
	if err != nil {
		return nil, err
	}
	defer conn.Close() // Leaving connections open will leak memory.

	c := cdp.NewClient(conn)

	// Open a DOMContentEventFired client to buffer this event.
	loaded, err := c.Page.LoadEventFired(ctx)
	if err != nil {
		return nil, err
	}
	defer loaded.Close()

	// Enable events on the Page domain, it's often preferrable to create
	// event clients before enabling events so that we don't miss any.
	if err = c.Page.Enable(ctx); err != nil {
		return nil, err
	}

	if opts.Headers != nil {
		// Enable events on the Network domain.
		if err = c.Network.Enable(ctx, network.NewEnableArgs()); err != nil {
			return nil, err
		}

		if err := c.Network.SetExtraHTTPHeaders(ctx, network.NewSetExtraHTTPHeadersArgs(opts.Headers)); err != nil {
			return nil, err
		}
	}

	navArgs := page.NewNavigateArgs(opts.URL)
	if _, err = c.Page.Navigate(ctx, navArgs); err != nil {
		return nil, err
	}

	// Wait until we have a DOMContentEventFired event.
	if _, err = loaded.Recv(); err != nil {
		return nil, err
	}

	reply, err := c.Page.PrintToPDF(ctx, &page.PrintToPDFArgs{})
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(reply.Data)

	return r, nil
}

func applyOptions(o Options, ops ...OptionFunc) Options {
	for _, op := range ops {
		o = op(o)
	}
	return o
}
