package utils

import (
	"runtime"
)

type OSInfo struct {
	GOOS   string
	GOARCH string
}

func GetOSInfo() *OSInfo {
	return &OSInfo{
		GOOS:   runtime.GOOS,
		GOARCH: runtime.GOARCH,
	}
}
