package main

import (
	"Chromedriver_Updater/src"       // TODO: update to "github.com/theofal/Chromedriver_Updater/src" when public
	"Chromedriver_Updater/src/utils" // TODO: update to "github.com/theofal/Chromedriver_Updater/src/utils" when public
)

func main() {
	logger := utils.InitLogger().Sugar()
	src.NewApp(logger).PrintOsInfo()
}
