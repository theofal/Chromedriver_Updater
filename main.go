package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/theofal/Chromedriver_Updater/src"
	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

func main() {
	logger := zaplogger.InitLogger(zapcore.DebugLevel, zapcore.DebugLevel).Sugar()

	// flags:
	// -v (--version) get the latest version from a given major version (int)
	// -o (--output) set chromedriver path manually (default /usr/local/bin) (string)
	output := flag.String("f", viper.GetString("configPath"), "Specify the folder where the binary will be installed")
	install := flag.Bool("i", false, "App configuration.")
	version := flag.Int("v", 0,
		"Specify the major version of the chromedriver (default: 0 = Same as installed Google chrome version)")
	flag.Parse()

	if !*install {
		err := initViper()
		if err != nil { // Handle errors reading the config file
			logger.Fatalf("Have you done the install part (-i)? An error occurred while reading config file: %v", err)
		}
		logger.Infof("Config file found, path: %s", viper.GetString("configPath"))
	}

	if *install {
		err := installViper()
		if err != nil {
			logger.Fatalf("An error occurred while trying to configure the app: %v", err)
		}
	}

	if *version < 0 {
		logger.Fatalf("Version number cannot be negative.")
	}

	if *output == "" {
		*output = viper.GetString("configPath")
		logger.Infof("Empty flag detected, file path set to %s", *output)
	}

	app := src.NewApp(logger)
	app.InitApp(*version, *output)
}

func initViper() error {
	viper.SetConfigName("config")                                             // name of config file (without extension)
	viper.SetConfigType("yaml")                                               // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(os.Getenv("HOME") + "/.config/Chromedriver_Updater/") // call multiple times to add many search paths
	return viper.ReadInConfig()                                               // Find and read the config file
}

func installViper() error {
	err := os.MkdirAll(os.Getenv("HOME")+"/.config/Chromedriver_Updater/", 0777)
	if err != nil {
		return err
	}
	_, err = os.Create(os.Getenv("HOME") + "/.config/Chromedriver_Updater/config.yaml")
	if err != nil {
		return err
	}
	fmt.Println("Enter the path to your Chromedriver: ")

	// Taking input from user
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		return err
	}
	pathInput := scanner.Text()
	strings.Contains(pathInput, "\n")
	viper.Set("configPath", pathInput)
	return viper.WriteConfigAs(os.Getenv("HOME") + "/.config/Chromedriver_Updater/config.yaml")
}
