package src

import (
	"fmt"
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
	if osInfo.OS == "win" {
		chromePath := "Windows not implemented yet."
		return false, chromePath
	}
	logger.Fatalf("%v not supported yet.", osInfo.OS)
	return false, ""
}

func (chrome *Chrome) getChromeVersion() string {
	chromeExists, chromePath := verifyChromeExists()
	logger.Infof("chrome exists: %v", chromeExists)
	if chromeExists {
		fmt.Println(osInfo)
		if osInfo.OS == "mac" || osInfo.OS == "linux" {
			logger.Info("Getting Google Chrome version")
			out, err := exec.Command(chromePath, "--version").Output()
			if err != nil {
				logger.Fatal(err)
			}
			fmt.Println(string(out))
			chrome.version = strings.Split(string(out), " ")[2]
			fmt.Println(chrome.version)
			return chrome.version
		}
		if osInfo.OS == "win" {
			logger.Fatal("Windows not implemented yet.")
			return ""
		}
		logger.Fatalf("%v not supported yet.", osInfo.OS)
		return ""
	}
	return ""
}

/*
https://developer.chrome.com/docs/versionhistory/examples/
*/
