package src

import (
	"Chromedriver_Updater/src/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Chromedriver struct {
	version string
	path    string
	exists  bool
}

func (chromedriver *Chromedriver) verifyChromedriverExists() bool {
	logger.Info("Verifying chromedriver exists.")
	if osInfo.OS == "mac" {
		chromedriver.path = "/usr/local/bin/chromedriver"
		_, err := os.Stat(chromedriver.path)
		chromedriver.exists = !os.IsNotExist(err)
		return chromedriver.exists
	}
	if osInfo.OS == "linux" {
		// TODO: change to implement other names of chrome
		out, err := exec.Command("which", "google-chrome").Output()
		if err != nil {
			logger.Fatal(err)
		}
		_, err = os.Stat(string(out))
		if os.IsNotExist(err) {
			return false
		}
		chromedriver.path = string(out)
		return true
	}
	if osInfo.OS == "win" {
		//chromePath := "Windows not implemented yet."
		return false
	}
	logger.Fatalf("%s not supported yet.", osInfo.OS)
	return false
}

func (chromedriver *Chromedriver) getChromedriverVersion() string {
	if chromedriver.verifyChromedriverExists() {
		if osInfo.OS == "mac" || osInfo.OS == "linux" {
			logger.Info("Getting Chromedriver version")
			out, err := exec.Command(chromedriver.path, "--version").Output()
			if err != nil {
				logger.Fatal(err)
			}

			chromedriver.version = strings.Split(string(out), " ")[1]
			logger.Infof("Chromedriver binary detected: %s, %s", chromedriver.version, chromedriver.path)
			return chromedriver.version
		}
		if osInfo.OS == "win" {
			logger.Fatal("Windows not implemented yet.")
			return chromedriver.version
		}
		logger.Fatalf("%s not supported yet.", osInfo.OS)
		return ""
	}
	logger.Infof("Chromedriver detected: %v", false)
	return ""
}

func getLatestReleaseForSpecificVersion(majorVersion string) string {
	response, err := http.Get("https://chromedriver.storage.googleapis.com/LATEST_RELEASE_" + majorVersion)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""
}

func (chromedriver *Chromedriver) downloadChromedriver(version string) *Chromedriver {
	downloadPath := fmt.Sprintf(
		"https://chromedriver.storage.googleapis.com/%s/chromedriver_%s64.zip", version, osInfo.OS, // osInfo.ARCH,
	) //TODO fix ARCH
	zipFilePath := "/tmp/chromedriver.zip"

	resp, err := http.Get(downloadPath)
	if err != nil {
		logger.Errorf("An error occurred while trying to reach website: %s", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Errorf("An error occurred while closing the file: %s", err)
		}
	}(resp.Body)

	logger.Debug("Response status: ", resp.Status)
	if resp.StatusCode != 200 {
		logger.Errorf("HTML response: %s", err)
		return chromedriver
	}

	// Create the file
	out, err := os.Create(zipFilePath)
	if err != nil {
		logger.Errorf("An error occurred while creating file: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.Errorf("An error occurred while copying file: %s", err)
	}

	chromedriver.exists = true
	chromedriver.version = version
	return chromedriver.unzipChromedriver()
}

func (chromedriver *Chromedriver) unzipChromedriver() *Chromedriver {
	chromedriver.removeFile(chromedriver.path)

	zipper := utils.NewZipper("/tmp/chromedriver.zip", strings.Replace(chromedriver.path, "chromedriver", "", 1))
	err := zipper.UnzipSource()
	if err != nil {
		logger.Error(err)
	}
	chromedriver.removeFile("/tmp/chromedriver.zip")
	// si chromedriver.path diff de nil, on le met là
	// sinon, on le met à l'endroit par défaut /usr/local/bin/chromedriver
	// supprimer le zip file
	logger.Infof("Your chromedriver has been updated. %s, %s", chromedriver.version, chromedriver.path)
	return chromedriver
}

func (chromedriver *Chromedriver) removeFile(path string) *Chromedriver {
	_, err := os.Stat(path)
	if err == nil {
		err := os.Remove(path)
		if err != nil {
			logger.Fatal(err)
		}
		return chromedriver
	}
	return chromedriver
}
