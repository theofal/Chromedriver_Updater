package utils

import (
	"runtime"
)

type OSInfo struct {
	GOARCH string
	GOOS   string
}

func GetOSInfo() *OSInfo {
	return &OSInfo{
		GOARCH: runtime.GOARCH,
		GOOS:   runtime.GOOS,
	}
}
