package internal

import (
	"skrepysh-agent/pkg/config"
	"skrepysh-agent/pkg/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configPath = "/etc/skrepysh/config.yaml"
	conf       = &config.Config{}
	log        *zap.Logger
	port       uint16 = 8080
)

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "skrepysh-agent",
	}

	for _, cmd := range commands() {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}

func commands() []*cobra.Command {
	var cmds []*cobra.Command

	serveCmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := config.ReadYaml(configPath, conf)
			if err != nil {
				return err
			}
			log, err = config.InitLogger(&conf.Log)
			if err != nil {
				return err
			}
			return server.Serve(log, conf)
		},
	}
	serveCmd.Flags().StringVarP(&configPath, "config", "c", configPath, "path/to/config")

	cmds = append(cmds, serveCmd)

	return cmds
}
