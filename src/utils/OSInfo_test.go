package utils

import (
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
	"go.uber.org/zap/zapcore"
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
	want := &OSInfo{OS: runtime.GOOS, ARCHForVersionBelow115: runtime.GOARCH}
	got := GetOSInfo(logger)

	if strings.ToLower(runtime.GOOS) == "darwin" {
		want.OS = "mac"
		if strings.Contains(runtime.GOARCH, "arm") {
			want.ARCHForVersionBelow115 = "64_m1"
			if got.OS != want.OS && got.ARCHForVersionBelow115 != want.ARCHForVersionBelow115 {
				t.Errorf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
				return
			}
		}
		want.ARCHForVersionBelow115 = "64"
		if got.OS != want.OS && got.ARCHForVersionBelow115 != want.ARCHForVersionBelow115 {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if strings.ToLower(runtime.GOOS) == "linux" {
		want.OS = "linux"
		want.ARCHForVersionBelow115 = "64"
		if got.OS != want.OS && got.ARCHForVersionBelow115 != want.ARCHForVersionBelow115 {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if strings.Contains(strings.ToLower(runtime.GOOS), "win") {
		want.OS = "win"
		want.ARCHForVersionBelow115 = "64"
		if got.OS != want.OS && got.ARCHForVersionBelow115 != want.ARCHForVersionBelow115 {
			t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
			return
		}
	}
	if got.OS != want.OS && got.ARCHForVersionBelow115 != want.ARCHForVersionBelow115 {
		t.Fatalf("TestGetOSInfo FAILED : want %v, got %v.\n", want, got)
		return
	}
}
