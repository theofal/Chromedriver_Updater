package src

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
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
	logger.Fatalf("%v not supported yet.", osInfo.OS)
	return false
}

func (chromedriver *Chromedriver) getChromedriverVersion() string {
	if chromedriver.verifyChromedriverExists() {
		if osInfo.OS == "mac" || osInfo.OS == "linux" {
			logger.Info("Getting Google Chrome version")
			out, err := exec.Command(chromedriver.path, "--version").Output()
			if err != nil {
				logger.Fatal(err)
			}
			chromedriver.version = string(out)
			return chromedriver.version
		}
		if osInfo.OS == "win" {
			logger.Fatal("Windows not implemented yet.")
			return chromedriver.version
		}
		logger.Fatalf("%v not supported yet.", osInfo.OS)
		return ""
	}
	return ""
}

func (chromedriver *Chromedriver) removeOldChromedriver() *Chromedriver {
	if chromedriver.getChromedriverVersion() != "" {
		err := os.Remove(chromedriver.path)
		if err != nil {
			logger.Fatal(err)
		}
		return chromedriver
	}
	return chromedriver
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
		"https://chromedriver.storage.googleapis.com/index.html?path=%v/chromedriver_%v%v.zip", version, osInfo.OS, osInfo.ARCH,
	)
	response, err := http.Get(downloadPath)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(response.Body)

	if response.StatusCode != 200 {
		logger.Fatal("Received non 200 response code", err)
	}
	//Create a empty file
	file, err := os.Create("~/Downloads/chromedriver.zip")
	if err != nil {
		logger.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Fatal(err)
		}
	}(file)

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.Fatal(err)
	}
	chromedriver.exists = true
	return chromedriver
}

func (chromedriver *Chromedriver) unzipChromedriver() *Chromedriver {
	chromedriver.removeOldChromedriver()
	// TODO
	// si chromedriver.path diff de nil, on le met là
	// sinon, on le met à l'endroit par défaut /usr/local/bin/chromedriver
	// supprimer le zip file
	return chromedriver
}

/*
[x] Verify Chromedriver exists
[x] Get Chromedriver version
[] Download Chromedriver
[x] Remove old Chromedriver
[] Unzip Chromedriver
*/
