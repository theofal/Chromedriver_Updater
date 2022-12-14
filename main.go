// go: generate goversioninfo -icon = icon_YOUR_GO_PROJECT.ico

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"github.com/theofal/Chromedriver_Updater/src"
	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
)

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

func configureAppFile(logger *zap.SugaredLogger) {
	i := false

	for !i {
		err := initViper()
		if err != nil { // Handle errors reading the config file
			fmt.Print("No config file found. Do you want to install it? [y/n]: ")
			// Taking input from user
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			err = scanner.Err()
			if err != nil {
				logger.Fatalf("An error occured while scanning response: %v", err)
			}
			if strings.ToLower(scanner.Text()) == "y" {
				err := installViper()
				if err != nil {
					logger.Fatalf("An error occurred while trying to configure the app: %v", err)
				}
				i = true
			}
			if strings.ToLower(scanner.Text()) == "n" {
				os.Exit(0)
			}
			continue
		}
		i = true
		logger.Infof("Config file found, path: %s", viper.GetString("configPath"))
	}
}

func main() {
	logger := zaplogger.InitLogger(zapcore.DebugLevel, zapcore.DebugLevel).Sugar()

	// -f (--file) set chromedriver path manually (default /usr/local/bin) (string)
	file := flag.String("f", viper.GetString("configPath"), "Specify the folder where the binary will be installed")
	// -i (--install) configure the chromedriver path
	install := flag.Bool("i", false, "Configure the chromedriver path")
	// -v (--version) get the latest version from a given major version (int)
	version := flag.Int("v", 0,
		"Specify the major version of the chromedriver (default: 0 = Same as installed Google chrome version)")

	flag.Parse()

	if *version < 0 {
		logger.Fatalf("Version number cannot be negative.")
	}

	if *file == "" {
		if !*install {
			configureAppFile(logger)
		}
		if *install {
			err := installViper()
			if err != nil {
				logger.Fatalf("An error occurred while configuring file: %v", err)
			}
		}
		*file = viper.GetString("configPath")
		logger.Infof("File path set to %s", *file)
	}

	app := src.NewApp(logger)
	app.InitApp(*version, filepath.Clean(*file))
}
