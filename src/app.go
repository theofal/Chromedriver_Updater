package src

import (
	"Chromedriver_Updater/src/utils"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

type App struct {
	chrome       *Chrome
	chromedriver *Chromedriver
}

var logger *zap.SugaredLogger
var osInfo = utils.GetOSInfo()

func NewApp(loggerInstance *zap.SugaredLogger) *App {
	logger = loggerInstance
	return &App{
		chrome:       &Chrome{},
		chromedriver: &Chromedriver{},
	}
}

/*func versionsAreEqual(chromeVersion, chromedriverVersion string) bool {
	v1, _ := semver.Make(chromeVersion)
	v2, _ := semver.Make(chromedriverVersion)
	return v1.Compare(v2) == 0
}

func version1IsGreater(chromeVersion, chromedriverVersion string) bool {
	v1, _ := semver.Make(chromeVersion)
	v2, _ := semver.Make(chromedriverVersion)
	return v1.Compare(v2) == 1
}*/

func parseMajorVersion(version string) string {
	fmt.Println(strings.Split(version, ".")[0])
	return strings.Split(version, ".")[0]
}

func (app *App) PrintOsInfo() {
	fmt.Println(app.chrome.getChromeVersion())
	fmt.Println(app.chromedriver.getChromedriverVersion())
	fmt.Println(getLatestReleaseForSpecificVersion(parseMajorVersion(app.chrome.version)))
}

/*
[] get the path
[] Parse the major version
[x] get the chrome version
[] get the chrome driver version
[] verify if chrome driver version is compatible with chrome
[] if app.get_chromedriver_version() >= app.get_chrome_version():
print("Votre version de chromedriver est à jour.")
else:
print("Votre version de chromedriver n'est pas à jour.")
if not compatible :
[] delete old chromedriver
[] download new chromedriver
[] print every step in console
*/
