package main

import (
	"Chromedriver_Updater/src" // TODO: update to "github.com/theofal/Chromedriver_Updater/src" when public
	"Chromedriver_Updater/src/utils/zapLogger"
)

func main() {
	logger := zapLogger.InitLogger().Sugar()
	app := src.NewApp(logger)
	app.InitApp()
}
