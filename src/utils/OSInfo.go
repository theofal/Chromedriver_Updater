package utils

import (
	"runtime"
	"strings"
)

type OSInfo struct {
	OS   string
	ARCH string
}

func GetOSInfo() *OSInfo {
	platform := runtime.GOOS
	arch := runtime.GOARCH
	if strings.ToLower(runtime.GOOS) == "darwin" {
		platform = "mac"
		arch = "64"
		if !strings.Contains(arch, "amd") {
			arch = "64_m1"
		}
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	if strings.ToLower(runtime.GOOS) == "linux" {
		platform = "linux"
		arch = "64"
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	if strings.Contains(strings.ToLower(runtime.GOOS), "win") {
		platform = "win"
		arch = "32"
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	return &OSInfo{
		OS:   platform,
		ARCH: arch,
	}
}
