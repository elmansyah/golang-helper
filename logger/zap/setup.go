package zap

import (
	"fmt"
	"os"
	"path/filepath"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LoggerSetup(req *Params) (*zap.SugaredLogger, error) {
	logPath, err := validateRequest(req)
	if err != nil {
		return nil, err
	}
	
	cores := make([]zapcore.Core, 0, len(req.LogFiles))
	
	for _, logRequest := range req.LogFiles {
		if logRequest.FileName == "" {
			return nil, errFileNameRequired
		}
		
		core, err := buildZap(logPath, req, logRequest)
		if err != nil {
			return nil, err
		}
		
		cores = append(cores, core)
	}
	
	combined := zapcore.NewTee(cores...)
	logger := zap.New(combined)
	
	return logger.Sugar(), nil
}

func buildZap(logPath string, request *Params, file File) (zapcore.Core, error) {
	fullPath := filepath.Join(logPath, file.FileName)
	
	// By default file permission handled by lumberjack, but we need to check it again
	openFile, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, request.FilePermission)
	if err != nil {
		return nil, errOpenLogFile
	}
	
	openFile.Close()
	
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fullPath,
		MaxSize:    request.MaxSize,
		MaxBackups: request.MaxBackups,
		MaxAge:     request.MaxAge,
		Compress:   request.Compress,
	})
	
	multiWriter := zapcore.NewMultiWriteSyncer(writer, zapcore.AddSync(os.Stdout))
	encoderCfg := zap.NewProductionEncoderConfig()
	
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	
	return zapcore.NewCore(
		encoder,
		multiWriter,
		zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			return level >= file.Level
		}),
	), nil
}

func validateRequest(request *Params) (string, error) {
	if err := validateLogDir(request); err != nil {
		return "", err
	}
	
	if err := os.MkdirAll(request.LogDir, request.DirPermission); err != nil {
		return "", fmt.Errorf("%v: %w", errCreateDir, err)
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
