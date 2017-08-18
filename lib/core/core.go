package core

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

type Options struct {
	URL        string
	Headers    []byte
	ChromePath string
	listener   *net.TCPListener
	bindAddr   string
}

type OptionFunc func(Options) Options

func URL(url string) OptionFunc {
	return func(o Options) Options {
		o.URL = url
		return o
	}
}

func ChromePath(path string) OptionFunc {
	return func(o Options) Options {
		o.ChromePath = path
		return o
	}
}

func Headers(headers map[string]string) OptionFunc {
	return func(o Options) Options {
		o.Headers, _ = json.Marshal(headers)
		return o
	}
}

func newTCPListener(addr string) (*net.TCPListener, error) {
	res, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return net.ListenTCP("tcp", res)
}

func ToPDF(ctx context.Context, ops ...OptionFunc) (io.Reader, error) {
	l, err := newTCPListener("localhost:0")
	if err != nil {
		return nil, err
	}

	opts := applyOptions(Options{
		ChromePath: "chromium",
		bindAddr:   "http://" + l.Addr().String(),
		listener:   l,
	}, ops...)

	cmd, err := startChrome(ctx, opts)
	defer cmd.Process.Kill()
	if err != nil {
		return nil, err
	}

	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
	devt := devtool.New(opts.bindAddr)
	pt, err := devt.Get(ctx, devtool.Page)
	if err != nil {
		pt, err = devt.Create(ctx)
		if err != nil {
			return nil, err
		}
	}

	log.Println("Connected to headleass chrome at " + opts.bindAddr)

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

func startChrome(ctx context.Context, opts Options) (*exec.Cmd, error) {
	file, err := opts.listener.File()
	if err != nil {
		return nil, err
	}

	args := []string{
		"--headless",
		"--remote-debugging-socket-fd=3",
		"--disable-gpu",
		"--no-sandbox",
		"--homedir=/tmp",
		"--single-process",
		"--data-path=/tmp/data-path",
		"--disk-cache-dir=/tmp/cache-dir",
	}

	cmd := exec.CommandContext(ctx, opts.ChromePath, args...)
	cmd.ExtraFiles = []*os.File{file}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err = cmd.Start()
	return cmd, err
}
