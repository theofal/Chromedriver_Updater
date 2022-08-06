package src

import (
	"github.com/theofal/Chromedriver_Updater/src/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type App struct {
	chrome       *Chrome
	chromedriver *Chromedriver
}

var logger *zap.SugaredLogger
var osInfo *utils.OSInfo

func NewApp(loggerInstance *zap.SugaredLogger) *App {
	logger = loggerInstance
	osInfo = utils.GetOSInfo(logger)
	return &App{
		chrome:       &Chrome{},
		chromedriver: &Chromedriver{},
	}
}

func firstArgIsGreater(chromeVersion, chromedriverVersion string) bool {

	if chromedriverVersion == "" {
		return true
	}

	v1, v2 := strings.Split(chromeVersion, "."), strings.Split(chromedriverVersion, ".")
	var version2Int []int
	var version1Int []int

	if len(v1) > len(v2) {
		for i := 1; i <= len(v1)-len(v2); i++ {
			v2 = append(v2, "0")
		}
	}
	if len(v1) < len(v2) {
		for i := 1; i <= len(v2)-len(v1); i++ {
			v1 = append(v1, "0")
		}
	}
	for index := range v1 {
		tmp, err := strconv.Atoi(v1[index])
		if err != nil {
			logger.Fatalf("An error occurred while converting string to int: %s", err)
		}
		version1Int = append(version1Int, tmp)

		tmp, err = strconv.Atoi(v2[index])
		if err != nil {
			logger.Fatalf("An error occurred while converting string to int: %s", err)
		}
		version2Int = append(version2Int, tmp)

		if version1Int[index] > version2Int[index] {
			logger.Infof("Your Chromedriver version (%s) is behind your Google Chrome version (%s).", chromedriverVersion, chromeVersion)
			return true
		}
		if version1Int[index] == version2Int[index] {
			continue
		}
		return false
	}
	return false
}

func parseMajorVersion(version string) string {
	return strings.Split(version, ".")[0]
}

func (app *App) InitApp(version int, strArgs string) *App {
	app.chromedriver.path = strArgs + "/chromedriver" //TODO Renforcer
	logger.Infof("Chromedriver path set to %s.", app.chromedriver.path)

	chromeVersion := app.chrome.getChromeVersion()
	chromedriverVersion := app.chromedriver.getChromedriverVersion()

	if version != 0 {
		vers := strconv.Itoa(version)
		app.chromedriver = app.chromedriver.downloadChromedriver(getLatestReleaseForSpecificVersion(parseMajorVersion(vers)))
		return app
	}

	if firstArgIsGreater(chromeVersion, chromedriverVersion) {
		app.chromedriver = app.chromedriver.downloadChromedriver(getLatestReleaseForSpecificVersion(parseMajorVersion(app.chrome.version)))
		return app
	}

	logger.Infof("Your chromedriver is up to date (version: %s).", chromedriverVersion)
	return app
}
