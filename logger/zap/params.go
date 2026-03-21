package zap

import (
	"errors"
	"os"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	errNilRequest        = errors.New("logger config: request is nil")
	errDirRequired       = errors.New("logger config: base path is required")
	errFileNameRequired  = errors.New("logger config: file name is required")
	errCreateDir         = errors.New("logger config: failed to create directory")
	errOpenLogFile       = errors.New("logger config: failed to open log file")
	errCloseLogFile      = errors.New("logger config: failed to close log file")
	errInvalidLevelRange = errors.New("logger config: min level must be less than or equal to max level")
)

type Zap struct {
	Sugar *zap.SugaredLogger
}

type File struct {
	FileName string
	MinLevel zapcore.Level
	MaxLevel zapcore.Level
}

type Params struct {
	AppMode        string
	LogDir         string
	LogFiles       []File
	DirPermission  os.FileMode
	FilePermission os.FileMode
	MaxSize        int
	MaxBackups     int
	MaxAge         int
	Compress       bool
}
