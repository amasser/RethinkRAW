package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"rethinkraw/internal/chrome"
	"rethinkraw/internal/config"
	"rethinkraw/internal/util"
	"rethinkraw/osutil"
)

var shutdown = make(chan os.Signal, 1)

func init() {
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	util.HideConsole()
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if err := config.SetupPaths(); err != nil {
		return err
	}

	url := url.URL{
		Scheme: "http",
		Host:   "[::1]:39639",
	}

	if len(os.Args) > 1 {
		if fi, err := os.Stat(os.Args[1]); err != nil {
			return err
		} else if abs, err := filepath.Abs(os.Args[1]); err != nil {
			return err
		} else if len(os.Args) > 2 {
			url.Path = "/batch/" + batches.New(os.Args[1:])
		} else if fi.IsDir() {
			url.Path = "/gallery/" + toURLPath(abs)
		} else {
			url.Path = "/photo/" + toURLPath(abs)
		}
	}

	if err := testDNGConverter(); err != nil {
		url.Path = "/dngconv.html"
	}

	var server bool
	if ln, err := net.Listen("tcp", url.Host); err == nil {
		server = true
		http := setupHTTP()
		exif, err := setupExifTool()
		if err != nil {
			return err
		}
		defer func() {
			http.Shutdown(context.Background())
			exif.Shutdown()
			os.RemoveAll(config.TempDir)
		}()
		go http.Serve(ln)
	}

	if chrome.Installed() {
		data := filepath.Join(config.DataDir, "chrome")
		cache := filepath.Join(config.TempDir, "chrome")
		cmd := chrome.New(url.String(), data, cache)

		if err := cmd.Start(); err != nil {
			return err
		}
		go func() {
			<-shutdown
			cmd.Exit()
		}()
		if err := cmd.Wait(); err != nil {
			return err
		}
	} else {
		if err := osutil.ShellOpen(url.String()); err != nil {
			return err
		}
		if server {
			<-shutdown
		}
	}

	return nil
}
