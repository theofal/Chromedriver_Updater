package src

import (
	"os"
	"os/exec"
	"strings"
)

type Chrome struct {
}

func (chrome *Chrome) verifyChromeExists() (bool, string) {
	if strings.ToLower(osInfo.GOOS) == "darwin" {
		chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
		_, err := os.Stat(chromePath)
		if os.IsNotExist(err) {
			return false, ""
		}
		return true, chromePath
	}
	if strings.ToLower(osInfo.GOOS) == "linux" {
		// TODO: change to implement other names of chrome
		out, err := exec.Command("which", "google-chrome").Output()
		if err != nil {
			logger.Fatal(err)
		}
		_, err = os.Stat(string(out))
		if os.IsNotExist(err) {
			return false, ""
		}
		chromePath := string(out)
		return true, chromePath
	}
	if strings.Contains(strings.ToLower(osInfo.GOOS), "win") {
		chromePath := "Windows not implemented yet."
		return false, chromePath
	}
	logger.Fatalf("%v not supported yet.", osInfo.GOOS)
	return false, ""
}

func (chrome *Chrome) getChromeVersion() string {
	chromeExists, chromePath := chrome.verifyChromeExists()
	if chromeExists {
		if strings.ToLower(osInfo.GOOS) == "darwin" || strings.ToLower(osInfo.GOOS) == "linux" {
			out, err := exec.Command(chromePath, "-v").Output()
			if err != nil {
				logger.Fatal(err)
			}
			return string(out)
		}
		if strings.Contains(strings.ToLower(osInfo.GOOS), "win") {
			return "Windows not implemented yet."
		}
		logger.Fatalf("%v not supported yet.", osInfo.GOOS)
		return ""
	}
	return ""
}

/*
Verify Chrome exists
Get Chrome version
https://developer.chrome.com/docs/versionhistory/examples/
*/
