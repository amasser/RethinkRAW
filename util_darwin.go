package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func init() {
	// ignore Process Serial Number argument
	for i, a := range os.Args {
		if strings.HasPrefix(a, "-psn_") {
			os.Args = append(os.Args[:i], os.Args[i+1:]...)
			break
		}
	}
}

func findChrome() string {
	versions := []string{"Google Chrome", "Chromium"}

	for _, v := range versions {
		c := filepath.Join("/Applications", v+".app", "Contents/MacOS", v)
		if _, err := os.Stat(c); err == nil {
			return c
		}
	}
	return ""
}

func exitChrome(cmd *exec.Cmd) {}

func openURLCmd(url string) *exec.Cmd {
	return exec.Command("open", url)
}

func isHidden(fi os.FileInfo) bool {
	// TODO: check for UF_HIDDEN flag?
	return strings.HasPrefix(fi.Name(), ".")
}

func getANSIPath(path string) (string, error) {
	return path, nil
}

func hideConsole() {}

func bringToTop() {}

func handleConsoleCtrl(c chan<- os.Signal) {}
