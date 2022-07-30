package src

import (
	"os"
	"os/exec"
	"strings"
)

type Chrome struct {
	version string
}

func verifyChromeExists() (bool, string) {
	logger.Info("Verifying chrome exists.")
	if osInfo.OS == "mac" {
		chromePath := "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
		_, err := os.Stat(chromePath)
		if os.IsNotExist(err) {
			return false, ""
		}
		return true, chromePath
	}
	if osInfo.OS == "linux" {
		// TODO: change to implement other names of chrome packages
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
	if osInfo.OS == "win" {
		chromePath := "Windows not implemented yet."
		return false, chromePath
	}
	logger.Fatalf("%s not supported yet.", osInfo.OS)
	return false, ""
}

func (chrome *Chrome) getChromeVersion() string {
	var chromeExists, chromePath = verifyChromeExists()
	if chromeExists {
		if osInfo.OS == "mac" || osInfo.OS == "linux" {
			logger.Info("Getting Google Chrome version")
			out, err := exec.Command(chromePath, "--version").Output()
			if err != nil {
				logger.Fatal(err)
			}
			chrome.version = strings.Split(string(out), " ")[2]
			logger.Infof("Google Chrome detected: %s, %s", chrome.version, chromePath)
			return chrome.version
		}
		if osInfo.OS == "win" {
			logger.Fatal("Windows not supported yet.")
			return ""
		}
		logger.Fatalf("%s not supported yet.", osInfo.OS)
		return ""
	}
	logger.Debugf("Google Chrome detected: %v", false)
	return ""
}

/*
https://developer.chrome.com/docs/versionhistory/examples/
*/
