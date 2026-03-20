package zap

import (
	"errors"
	"os"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// validation errors
	errNilRequest       = errors.New("request is nil")
	errDirRequired      = errors.New("logger config: base path is required")
	errFileNameRequired = errors.New("logger config: file name is required")
	
	// runtime errors
	errCreateDir   = errors.New("failed to create directory")
	errOpenLogFile = errors.New("failed to open log file")
)

type Zap struct {
	Sugar *zap.SugaredLogger
}

type File struct {
	FileName string
	Level    zapcore.Level
}

type Params struct {
	LogDir         string
	LogFiles       []File
	DirPermission  os.FileMode
	FilePermission os.FileMode
	MaxSize        int
	MaxBackups     int
	MaxAge         int
	Compress       bool
}
