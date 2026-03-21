package zap

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Setup(req *Params) (*zap.SugaredLogger, error) {
	logPath, err := validateRequest(req)
	if err != nil {
		return nil, err
	}
	
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "message",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	
	// normalize app mode
	mode := strings.ToLower(req.AppMode)
	if mode != "dev" && mode != "prod" {
		mode = "prod"
	}
	
	cores := make([]zapcore.Core, 0, len(req.LogFiles)+1)
	
	// build file cores
	for _, logRequest := range req.LogFiles {
		if logRequest.FileName == "" {
			return nil, errFileNameRequired
		}
		
		core, err := buildZap(logPath, req, logRequest, encoderCfg)
		if err != nil {
			return nil, err
		}
		
		cores = append(cores, core)
	}
	
	// console level based on mode
	var consoleLevel zapcore.Level
	if mode == "dev" {
		consoleLevel = zapcore.DebugLevel
	} else {
		consoleLevel = zapcore.InfoLevel
	}
	
	// console core (stdout only)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	consoleCore := zapcore.NewCore(
		consoleEncoder,
		zapcore.AddSync(os.Stdout),
		consoleLevel,
	)
	
	cores = append(cores, consoleCore)
	
	combined := zapcore.NewTee(cores...)
	
	logger := zap.New(
		combined,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	
	return logger.Sugar(), nil
}

func buildZap(logPath string, request *Params, file File, encoderCfg zapcore.EncoderConfig) (zapcore.Core, error) {
	fullPath := filepath.Join(logPath, file.FileName)
	
	// ensure file can be created with correct permission
	openFile, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, request.FilePermission)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errOpenLogFile, fullPath)
	}
	
	err = openFile.Close()
	if err != nil {
		return nil, fmt.Errorf("%w: %s", errCloseLogFile, fullPath)
	}
	
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    request.MaxSize,
		MaxBackups: request.MaxBackups,
		MaxAge:     request.MaxAge,
		Compress:   request.Compress,
	})
	
	// file encoder always JSON (structured logging)
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	
	if file.MinLevel > file.MaxLevel {
		return nil, errInvalidLevelRange
	}
	
	return zapcore.NewCore(
		encoder,
		writer,
		levelEnabler(file),
	), nil
}

func validateRequest(request *Params) (string, error) {
	if err := validateLogDir(request); err != nil {
		return "", err
	}
	
	if err := os.MkdirAll(request.LogDir, request.DirPermission); err != nil {
		return "", fmt.Errorf("%w: %w", errCreateDir, err)
	}
	
	return request.LogDir, nil
}

func validateLogDir(request *Params) error {
	if request == nil {
		return errNilRequest
	}
	
	if len(request.LogFiles) == 0 {
		return errDirRequired
	}
	
	setDefaultPermission(request)
	
	return nil
}

func setDefaultPermission(request *Params) {
	if request.DirPermission == 0 {
		request.DirPermission = 0755
	}
	
	if request.FilePermission == 0 {
		request.FilePermission = 0644
	}
}

func levelEnabler(file File) zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= file.MinLevel && level <= file.MaxLevel
	}
}
