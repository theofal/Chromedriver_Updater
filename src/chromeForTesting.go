package src

import "time"

type ChromeForTesting struct {
	Timestamp time.Time `json:"timestamp"`
	Versions  []struct {
		Version   string `json:"version"`
		Revision  string `json:"revision"`
		Downloads struct {
			Chromedriver []struct {
				Platform string `json:"platform"`
				URL      string `json:"url"`
			} `json:"chromedriver"`
		} `json:"downloads"`
	} `json:"versions"`
}

type Version struct {
	Version   string `json:"version"`
	Revision  string `json:"revision"`
	Downloads struct {
		Chromedriver []struct {
			Platform string `json:"platform"`
			URL      string `json:"url"`
		} `json:"chromedriver"`
	} `json:"downloads"`
}
