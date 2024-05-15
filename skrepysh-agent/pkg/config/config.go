package config

import "go.uber.org/zap/zapcore"

type Config struct {
	Log             Log                   `yaml:"log"`
	NodeExporter    NodeExporterConfig    `yaml:"node-exporter"`
	ServerPort      uint16                `yaml:"server-port"`
	SkrepyshBackend SkrepyshBackendConfig `yaml:"skrepysh-backend"`
}

type Log struct {
	Level           LogLevel `yaml:"level"`
	OutputPath      []string `yaml:"output-path"`
	ErrorOutputPath []string `yaml:"error-output-path"`
}

type NodeExporterConfig struct {
	Host string `yaml:"host"`
	Port uint16 `json:"port"`
}

type SkrepyshBackendConfig struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type LogLevel string

const (
	LogLevel_NONE  LogLevel = ""
	LogLevel_DEBUG LogLevel = "DEBUG"
	LogLevel_INFO  LogLevel = "INFO"
	LogLevel_WARN  LogLevel = "WARN"
	LogLevel_ERROR LogLevel = "ERROR"
	LogLevel_PANIC LogLevel = "PANIC"
	LogLevel_FATAL LogLevel = "FATAL"
)

var MapToZap = map[LogLevel]zapcore.Level{
	LogLevel_NONE:  zapcore.DebugLevel,
	LogLevel_DEBUG: zapcore.DebugLevel,
	LogLevel_INFO:  zapcore.InfoLevel,
	LogLevel_WARN:  zapcore.WarnLevel,
	LogLevel_ERROR: zapcore.ErrorLevel,
	LogLevel_PANIC: zapcore.PanicLevel,
	LogLevel_FATAL: zapcore.FatalLevel,
}
