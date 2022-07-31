package src

import (
	"os"
	"os/exec"
	"path/filepath"
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
			logger.Fatal("Could not find Google Chrome app.")
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
				logger.Fatalf("An error occured while getting the chrome version: %s", err)
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
	logger.Fatalf("Google Chrome detected: %v", false)
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
		logger.Debugf("Trying to find %s binary: %s", appName, strings.Split(string(out), "\n")[0])
		if err == nil {
			output := strings.Split(string(out), "\n")[0]
			symlinks, err := filepath.EvalSymlinks(output)
			if err != nil {
				logger.Fatalf("An error occurred while evaluation symlink: %s", err)
			}
			return symlinks
		}
		continue
	}
	logger.Fatal("Could not find Google Chrome app.")
	return ""
}
