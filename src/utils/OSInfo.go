package utils

import (
	"runtime"
	"strings"

	"go.uber.org/zap"
)

type OSInfo struct {
	OS                     string
	ARCHForVersionBelow115 string
	ARCHForVersionAbove115 string
}

var logger *zap.SugaredLogger

func GetOSInfo(loggerInstance *zap.SugaredLogger) *OSInfo {
	logger = loggerInstance
	platform := runtime.GOOS
	archForVersionBelow115 := runtime.GOARCH
	archForVersionAbove115 := runtime.GOARCH

	logger.Debugf("GOOS: %s, GOARCH: %s", runtime.GOOS, runtime.GOARCH)
	if strings.ToLower(runtime.GOOS) == "darwin" {
		platform = "mac"
		archForVersionBelow115 = "64"
		archForVersionAbove115 = "-x64"
		if strings.Contains(runtime.GOARCH, "arm") {
			archForVersionBelow115 = "64_m1"
			archForVersionAbove115 = "-arm64"
		}
		logger.Debugf("OS: %v, Arch: %v", platform, archForVersionBelow115)
		return &OSInfo{
			OS:                     platform,
			ARCHForVersionBelow115: archForVersionBelow115,
			ARCHForVersionAbove115: archForVersionAbove115,
		}
	}
	if strings.ToLower(runtime.GOOS) == "linux" {
		platform = "linux"
		archForVersionBelow115 = "64"
		archForVersionAbove115 = "64"
		logger.Debugf("OS: %v, Arch: %v", platform, archForVersionBelow115)
		return &OSInfo{
			OS:                     platform,
			ARCHForVersionBelow115: archForVersionBelow115,
			ARCHForVersionAbove115: archForVersionAbove115,
		}
	}
	if strings.Contains(strings.ToLower(runtime.GOOS), "win") {
		platform = "win"
		archForVersionBelow115 = "32"
		archForVersionAbove115 = "32"
		logger.Debugf("OS: %v, Arch: %v", platform, archForVersionBelow115)
		return &OSInfo{
			OS:                     platform,
			ARCHForVersionBelow115: archForVersionBelow115,
			ARCHForVersionAbove115: archForVersionAbove115,
		}
	}
	logger.Debugf("OS: %v, Arch: %v", platform, archForVersionBelow115)
	return &OSInfo{
		OS:                     platform,
		ARCHForVersionBelow115: archForVersionBelow115,
		ARCHForVersionAbove115: archForVersionAbove115,
	}
}
