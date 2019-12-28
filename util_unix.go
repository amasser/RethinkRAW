// +build !windows
// +build !darwin

package main

import (
	"os"
	"os/exec"
	"strings"
)

func findChrome() string {
	versions := []string{"google-chrome-stable", "google-chrome", "chromium-browser", "chromium"}

	for _, v := range versions {
		if c, err := exec.LookPath(v); err == nil {
			return c
		}
	}
	return ""
}

func exitChrome(cmd *exec.Cmd) {}

func openURLCmd(url string) *exec.Cmd {
	return exec.Command("xdg-open", url)
}

func isHidden(fi os.FileInfo) bool {
	return strings.HasPrefix(fi.Name(), ".")
}

func getANSIPath(path string) (string, error) {
	return path, nil
}

func hideConsole() {}

func bringToTop() {}

func handleConsoleCtrl(c chan<- os.Signal) {}