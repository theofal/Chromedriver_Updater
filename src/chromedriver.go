package src

import (
	"sync"
)

type Chromedriver struct {
}

var lockChromedriver = &sync.Mutex{}
var singleChromedriverInstance *Chromedriver

func GetChromedriverInstance() *Chromedriver {
	if singleChromedriverInstance == nil {
		lockChromedriver.Lock()
		defer lockChromedriver.Unlock()
		if singleChromedriverInstance == nil {
			logger.Debug("Creating Chromedriver instance.")
			singleChromedriverInstance = &Chromedriver{}
		}
	}
	return singleChromedriverInstance
}
