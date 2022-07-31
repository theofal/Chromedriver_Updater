# Chromedriver_Updater

This program automatically verifies and updates your Chromedriver to your Google Chrome version.

/!\ : Still work in progress. Do not use it without verifying compatibility and test in the Troubleshooting section.

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
$ ./Chromedriver_Updater -h                                                                                                                       (main✱)
Usage of ./Chromedriver_Updater:
  -f string
    	Specify the folder where the binary will be installed (default "/usr/local/bin")
  -v int
    	Specify the major version of the chromedriver (default: 0 = Same as installed Google chrome version)
```

## Troubleshooting

[x] Tested on mac amd64

[] Tested on mac arm64

[] Tested on linux

[] Tested on windows