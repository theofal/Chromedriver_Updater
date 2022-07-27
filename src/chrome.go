package src

import (
	"os"
	"strings"
)

type Chrome struct {
	path string
}

func (chrome *Chrome) verifyChromeExists() (bool, string) {
	if strings.Contains(osInfo.GOOS, "darwin") {
		chrome.path = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
		_, err := os.Stat(chrome.path)
		if os.IsNotExist(err) {
			return false, ""
		}
		return true, chrome.path
	}
	if strings.Contains(osInfo.GOOS, "linux") {
		chrome.path = ""
		return false, "Linux not implemented yet."
	}
	if strings.Contains(osInfo.GOOS, "win") {
		chrome.path = ""
		return false, "Windows not implemented yet."
	}
	logger.Errorf("%v not supported yet.", osInfo.GOOS)
	os.Exit(0)
	return false, ""
}

/*
Verify Chrome exists
Get Chrome version
*/
