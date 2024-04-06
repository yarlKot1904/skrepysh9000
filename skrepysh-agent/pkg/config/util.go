package config

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
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
	} else if err := createLogFiles(outputPaths); err != nil {
		return nil, err
	}
	if errorOutputPath == nil {
		errorOutputPath = []string{"stdout", defaultLogPath}
	} else if err := createLogFiles(errorOutputPath); err != nil {
		return nil, err
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

func createLogFiles(filepaths []string) error {
	for _, f := range filepaths {
		if f == "stdout" || f == "stderr" {
			continue
		}
		dir := filepath.Dir(f)
		if err := os.MkdirAll(dir, 0644); err != nil {
			return err
		}
		if _, err := os.OpenFile(f, os.O_RDONLY|os.O_CREATE, 0666); err != nil {
			return err
		}
	}
	return nil
}
