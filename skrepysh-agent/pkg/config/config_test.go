package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	expected := Config{
		Log: Log{
			Level:      LogLevel_DEBUG,
			OutputPath: []string{"/var/log/skrepysh/skrepysh.log"},
		},
		NodeExporter: NodeExporterConfig{
			Host: "http://localhost",
			Port: 9100,
		},
	}
	actual := &Config{}
	err := ReadYaml("example.yaml", actual)
	require.NoError(t, err)
	require.Equal(t, expected, *actual)
}
