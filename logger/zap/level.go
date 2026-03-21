package zap

import (
	"strings"
	
	"go.uber.org/zap/zapcore"
)

func StringToZapLevel(value string, def zapcore.Level) zapcore.Level {
	switch strings.ToLower(value) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn", "warning":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return def
	}
}
