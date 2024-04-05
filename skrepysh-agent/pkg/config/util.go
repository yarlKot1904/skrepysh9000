package config

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	defaultLogPath = "/var/log/skrepysh/skrepysh.log"
)

func ReadYaml(filepath string, dst interface{}) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(file, dst)
}

func InitLogger(conf *Log) (*zap.Logger, error) {
	level, ok := MapToZap[conf.Level]
	if !ok {
		return nil, fmt.Errorf("unknown log level")
	}
	outputPaths := conf.OutputPath
	errorOutputPath := conf.ErrorOutputPath
	if outputPaths == nil {
		outputPaths = []string{"stdout", defaultLogPath}
	}
	if errorOutputPath == nil {
		errorOutputPath = []string{"stdout", defaultLogPath}
	}
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(level),
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: errorOutputPath,
	}
	return config.Build()
}
