# Chromedriver_Updater

This program automatically verifies and updates your Chromedriver to your Google Chrome version.

## How does it work

On first launch, this program writes the chromedriver path in a config file in `~/.config/Chromedriver_Updater/config.yaml`.
It is then used to compare the existing version of Google Chrome to your Chromedriver's version, then updates it if needed.

If you need to reconfigure the chromedriver path, simply usr the `-i` flag. 

/!\ : Still work in progress. Do not use it without verifying compatibility in the Troubleshooting section. **Windows not supported yet.**

## How to launch it:
- Clone the repository:
  ```bash
  git clone https://github.com/theofal/Chromedriver_Updater.git
  ```
- Go to your app clone folder:
  ```bash
  cd path/to/the/clone/folder
  ```
- Build the binary:
  ```bash
  go build .
  ```
- Execute the binary:
  ```bash
  ./Chromedriver_Updater
  ```

## Flags
```
$ ./Chromedriver_Updater -h
Usage of ./Chromedriver_Updater:
  -f string
    	Specify the folder where the binary will be installed
  -i	
        Configure the chromedriver path
  -v int
    	Specify the major version of the chromedriver 
    	(default: 0 = Same as installed Google chrome version)
```

## Troubleshooting

- [x] Tested on mac amd64

- [x] Tested on mac arm64

- [x] Tested on linux

*Not compatible with windows*