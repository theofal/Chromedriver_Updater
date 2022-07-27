package src

import (
	"sync"
)

type Chrome struct {
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
