package utils

import (
	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	// exec setUp function
	//setUp()
	logger = zaplogger.InitLogger(zapcore.ErrorLevel, zapcore.ErrorLevel).Sugar()
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// exec tearDown function
	//tearDown("one")
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

func TestGetOSInfo(t *testing.T) {
	want := &OSInfo{OS: runtime.GOOS, ARCH: runtime.GOARCH}
	got := GetOSInfo(logger)

	if strings.ToLower(runtime.GOOS) == "darwin" {
		want.OS = "mac"
		if strings.Contains(runtime.GOARCH, "arm") {
			want.ARCH = "64_m1"
			if got.OS != want.OS && got.ARCH != want.ARCH {
				t.Errorf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
				return
			}
		}
		want.ARCH = "64"
		if got.OS != want.OS && got.ARCH != want.ARCH {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if strings.ToLower(runtime.GOOS) == "linux" {
		want.OS = "linux"
		want.ARCH = "64"
		if got.OS != want.OS && got.ARCH != want.ARCH {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if strings.Contains(strings.ToLower(runtime.GOOS), "win") {
		want.OS = "win"
		want.ARCH = "64"
		if got.OS != want.OS && got.ARCH != want.ARCH {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if got.OS != want.OS && got.ARCH != want.ARCH {
		t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
		return
	}
}
