package src

import (
	"github.com/blang/semver/v4"
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
	v1, _ := semver.Make(chromeVersion)
	v2, _ := semver.Make(chromedriverVersion)
	return v1.Compare(v2) == 1
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

	if firstArgIsGreater(chromeVersion, chromedriverVersion) || chromedriverVersion == "" {
		app.chromedriver = app.chromedriver.downloadChromedriver(getLatestReleaseForSpecificVersion(parseMajorVersion(app.chrome.version)))
		return app
	}

	logger.Info("Your chromedriver is up to date.")
	return app
}

/*
[] print every step in console
[] implement flags
*/
