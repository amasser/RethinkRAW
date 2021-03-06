package chrome

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func findChrome() {
	versions := []string{`Google\Chrome`, `Chromium`}
	prefixes := []string{os.Getenv("LOCALAPPDATA"), os.Getenv("PROGRAMFILES"), os.Getenv("PROGRAMFILES(X86)")}
	suffix := `\Application\chrome.exe`

	for _, v := range versions {
		for _, p := range prefixes {
			if p != "" {
				c := filepath.Join(p, v, suffix)
				if _, err := os.Stat(c); err == nil {
					chrome = c
					return
				}
			}
		}
	}
}

func (c *Cmd) Exit() {
	for i := 0; i < 10; i++ {
		if exec.Command("taskkill", "/pid", strconv.Itoa(c.cmd.Process.Pid)).Run() != nil {
			return
		}
		time.Sleep(time.Second / 10)
	}
}
