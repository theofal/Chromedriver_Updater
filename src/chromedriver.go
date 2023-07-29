package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/theofal/Chromedriver_Updater/src/utils"
)

type Chromedriver struct {
	majorVersion string
	version      string
	path         string
	exists       bool
	downloadPath string
}

func (chromedriver *Chromedriver) verifyChromedriverExists() bool {
	logger.Info("Verifying chromedriver exists.")
	if osInfo.OS == "mac" || osInfo.OS == "linux" {
		_, err := os.Stat(chromedriver.path)
		chromedriver.exists = !os.IsNotExist(err)
		return chromedriver.exists
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
			chromedriver.majorVersion = strings.Split(chromedriver.version, ".")[0]
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

func getDownloadPathForVersionBelow115(majorVersion string) (string, string) {

	major, _ := strconv.Atoi(majorVersion)
	if major >= 106 && osInfo.ARCHForVersionBelow115 == "64_m1" {
		osInfo.ARCHForVersionBelow115 = "_arm64"
	}

	version := getLatestReleaseForSpecificVersionBelow115(majorVersion)

	downloadPath := fmt.Sprintf(
		"https://chromedriver.storage.googleapis.com/%s/chromedriver_%s%s.zip",
		version,
		osInfo.OS,
		osInfo.ARCHForVersionBelow115,
	)

	return downloadPath, version
}

func getLatestReleaseForSpecificVersionBelow115(majorVersion string) string {
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

func getDownloadPathForVersionAboveOrEqual115(majorVersion string) Version {
	var chromeForTesting ChromeForTesting
	var tmpVersions []Version

	response, err := http.Get(
		"https://googlechromelabs.github.io/chrome-for-testing/known-good-versions-with-downloads.json",
	)

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

		err = json.Unmarshal(bodyBytes, &chromeForTesting)
	}

	for _, version := range chromeForTesting.Versions {
		if parseMajorVersion(version.Version) == majorVersion {
			tmpVersions = append(tmpVersions, version)
		}
	}

	return tmpVersions[len(tmpVersions)-1]
}

func (chromedriver *Chromedriver) downloadChromedriver(downloadPath string) *Chromedriver {

	zipFilePath := "/tmp/chromedriver.zip"
	logger.Infof("Downloading from: %s", downloadPath)

	//TODO: Sanitize
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

	logger.Debug("Chromedriver downloader website response status: ", resp.Status)
	if resp.StatusCode != 200 {
		logger.Fatalf("HTML response: %s", err)
		return chromedriver
	}

	// Create the file
	out, err := os.Create(zipFilePath)
	if err != nil {
		logger.Errorf("An error occurred while creating file: %s", err)
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		logger.Errorf("An error occurred while copying file: %s", err)
	}

	chromedriver.exists = true
	//chromedriver.version = version
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
	logger.Infof("Your chromedriver has been updated. %s, %s", chromedriver.version, chromedriver.path)
	return chromedriver
}

func (chromedriver *Chromedriver) removeFile(path string) *Chromedriver {
	_, err := os.Stat(path)
	if err == nil {
		err := os.RemoveAll(path)
		if err != nil {
			logger.Fatal(err)
		}
		return chromedriver
	}
	return chromedriver
}
