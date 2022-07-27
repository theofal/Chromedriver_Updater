package main

import (
	"Chromedriver_Updater/src"
	"Chromedriver_Updater/src/utils"
)

func main() {
	logger := utils.InitLogger().Sugar()
	src.GetApp(logger).PrintLogger("hi")
}
