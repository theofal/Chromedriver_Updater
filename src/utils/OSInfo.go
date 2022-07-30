package utils

import (
	"go.uber.org/zap"
	"runtime"
	"strings"
)

type OSInfo struct {
	OS   string
	ARCH string
}

var logger *zap.SugaredLogger

func GetOSInfo(loggerInstance *zap.SugaredLogger) *OSInfo {
	logger = loggerInstance
	platform := runtime.GOOS
	arch := runtime.GOARCH
	if strings.ToLower(runtime.GOOS) == "darwin" {
		platform = "mac"
		arch = "64"
		if !strings.Contains(arch, "amd") {
			arch = "64_m1"
		}
		logger.Debugf("OS: %v, Arch: %v", platform, arch)
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	if strings.ToLower(runtime.GOOS) == "linux" {
		platform = "linux"
		arch = "64"
		logger.Debugf("OS: %v, Arch: %v", platform, arch)
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	if strings.Contains(strings.ToLower(runtime.GOOS), "win") {
		platform = "win"
		arch = "32"
		logger.Debugf("OS: %v, Arch: %v", platform, arch)
		return &OSInfo{
			OS:   platform,
			ARCH: arch,
		}
	}
	logger.Debugf("OS: %v, Arch: %v", platform, arch)
	return &OSInfo{
		OS:   platform,
		ARCH: arch,
	}
}
