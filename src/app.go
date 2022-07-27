package src

import (
	"Chromedriver_Updater/src/utils"
	"go.uber.org/zap"
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

func (app *App) PrintOsInfo() {
	logger.Info(osInfo)
}

/*
get the OS
get the path
Parse the major version
get the chrome version
get the chrome driver version
verify if chrome driver version is compatible with chrome
if app.get_chromedriver_version() >= app.get_chrome_version():
print("Votre version de chromedriver est à jour.")
else:
print("Votre version de chromedriver n'est pas à jour.")
if not compatible :
delete old chromedriver
download new chromedriver
print every step in console
*/
