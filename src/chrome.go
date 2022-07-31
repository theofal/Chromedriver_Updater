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
		chromePath := getChromeVersionLinux()
		logger.Debugf("Google Chrome path set to: %s", chromePath)
		if chromePath == "" {
			return false, chromePath
		}
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

func getChromeVersionLinux() string {
	list := []string{
		"google-chrome",
		"google-chrome-stable",
		"google-chrome-beta",
		"google-chrome-dev",
		"chromium-browser",
		"chromium",
	}

	for _, appName := range list {
		out, err := exec.Command("which", appName).Output()
		logger.Debugf("Trying to find %s binary: %s", appName, string(out))
		if err == nil {
			return string(out)
		}
		continue
	}
	logger.Error("Could not find Google Chrome app.")
	return ""
}

/*
https://developer.chrome.com/docs/versionhistory/examples/
*/
