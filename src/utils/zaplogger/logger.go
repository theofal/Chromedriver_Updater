package zaplogger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	//Logger instance.
	logger *zap.Logger
)

// InitLogger : zaplogger initialisation.
func InitLogger(consoleLogLevel zapcore.Level, fileLogLevel zapcore.Level) *zap.Logger {
	writeSyncer := getLogWriter()
	encoder, encoderColored := getEncoder(), getColoredEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, fileLogLevel),                          //logfile output
		zapcore.NewCore(encoderColored, zapcore.AddSync(os.Stdout), consoleLogLevel), //console output
	)
	logger = zap.New(core, zap.AddCaller())
	return logger
}

// getEncoder returns an encoder used for logfiles.
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getColoredEncoder returns a colored encoder used for console output.
func getColoredEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Save file log cut.
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./src/logs/chromedriverUpdaterLogs.log", // Log name
		MaxSize:    1,                                        // File content size, MB
		MaxBackups: 5,                                        // Maximum number of backups
		MaxAge:     30,                                       // Maximum number of days to keep old files
		Compress:   false,                                    // Is the file compressed
	}
	return zapcore.AddSync(lumberJackLogger)
}
