package src

import (
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type App struct {
	chrome       *Chrome
	chromedriver *Chromedriver
}

var logger *zap.SugaredLogger
var lock = &sync.Mutex{}
var singleAppInstance *App

func GetApp(loggerInstance *zap.SugaredLogger) *App {
	logger = loggerInstance
	if singleAppInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleAppInstance == nil {
			logger.Debug("Creating App instance.")
			singleAppInstance = &App{
				chrome:       GetChromeInstance(),
				chromedriver: GetChromedriverInstance(),
			}
		}
	}
	return singleAppInstance
}

func (app *App) PrintLogger(str string) {
	fmt.Println(GetChromedriverInstance())
	logger.Info(GetChromedriverInstance())
}
