package src

import (
	"fmt"
	"github.com/theofal/Chromedriver_Updater/src/utils/zaplogger"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logger = zaplogger.InitLogger(zapcore.ErrorLevel, zapcore.ErrorLevel).Sugar()
	// exec test and this returns an exit code to pass to os
	retCode := m.Run()
	// If exit code is distinct of zero,
	// the test will be failed (red)
	os.Exit(retCode)
}

func TestFirstArgIsGreater(t *testing.T) {

	got := firstArgIsGreater("1.0.3.0", "0.1.0.3")
	if got != true {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", true, got)
	}

	got = firstArgIsGreater("0.1.0.3", "1.0.3.0")
	if got != false {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", false, got)
	}

	got = firstArgIsGreater("1.1.1.1", "1.1.1.1")
	if got != false {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", false, got)
	}

	got = firstArgIsGreater("1.1.1.1.1", "1.1.1.1")
	if got != true {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", true, got)
	}

	got = firstArgIsGreater("1.1.1.1", "1.1.1.1.1")
	if got != false {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", false, got)
	}

	got = firstArgIsGreater("1", "1.00000.00000")
	if got != false {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", false, got)
	}
}

func TestParseMajorVersion(t *testing.T) {
	got := parseMajorVersion("103.0.0.1")
	want := "103"
	if got != want {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", want, got)
	}
	got = parseMajorVersion("100.")
	want = "100"
	if got != want {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", want, got)
	}
	got = parseMajorVersion("92.098124.129412412")
	want = "92"
	if got != want {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", want, got)
	}
	got = parseMajorVersion("")
	want = ""
	if got != want {
		t.Errorf("TestFirstArgIsGreater FAILED : want %v, got %v.\n", want, got)
	}
}

func BenchmarkPrimeNumbers(b *testing.B) {
	n := 5
	for i := 0; i < b.N; i++ {
		firstArgIsGreater(
			fmt.Sprintf("%v.%v.%v.%v", i*rand.Intn(n), i*(i*rand.Intn(n)), i*i*(i*rand.Intn(n)), i*i*i*(i*rand.Intn(n))),
			fmt.Sprintf("%v.%v.%v.%v", i*(i*rand.Intn(n)), i*i*(i*rand.Intn(n)), i*i*i*(i*rand.Intn(n)), i*i*i*i*(i*rand.Intn(n))))
	}
}
