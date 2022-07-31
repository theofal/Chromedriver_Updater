package main

import (
	"flag"
	"github.com/theofal/Chromedriver_Updater/src"
	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
)

func main() {
	logger := zaplogger.InitLogger().Sugar()

	// flags:
	// -v (--version) get the latest version from a given major version (int)
	// -o (--output) set chromedriver path manually (default /usr/local/bin) (string)
	output := flag.String("f", "/usr/local/bin", "Specify the folder where the binary will be installed")
	version := flag.Int("v", 0,
		"Specify the major version of the chromedriver (default: 0 = Same as installed Google chrome version)")
	flag.Parse()

	if *version < 0 {
		logger.Fatalf("Version number cannot be negative.")
	}

	if *output == "" {
		*output = "/usr/local/bin"
	}

	app := src.NewApp(logger)
	app.InitApp(*version, *output)
}
