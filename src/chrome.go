package src

import (
	"sync"
)

type Chrome struct {
	path string
}

var lockChrome = &sync.Mutex{}
var singleChromeInstance *Chrome

func GetChromeInstance() *Chrome {

	if singleChromeInstance == nil {
		lockChrome.Lock()
		defer lockChrome.Unlock()
		if singleChromeInstance == nil {
			logger.Debug("Creating Chrome instance.")
			singleChromeInstance = &Chrome{}
		}
	}
	return singleChromeInstance
}

/*
Verify Chrome exists
Get Chrome version
*/
