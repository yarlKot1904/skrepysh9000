package internal

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"skrepysh-agent/pkg/config"
)

var (
	configPath string         = "/etc/skrepysh/config.yaml"
	conf       *config.Config = &config.Config{}
	log        *zap.Logger
)

func RootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "skrepysh-agent",
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := config.ReadYaml(configPath, conf)
			if err != nil {
				return err
			}
			log, err = config.InitLogger(&conf.Log)
			return err
		},
	}
	rootCmd.Flags().StringVarP(&configPath, "config", "c", configPath, "path/to/config")

	for _, cmd := range commands() {
		rootCmd.AddCommand(cmd)
	}
	return rootCmd
}

func commands() []*cobra.Command {
	var cmds []*cobra.Command

	initCmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	cmds = append(cmds, initCmd)

	validateCmd := &cobra.Command{
		Use: "validate",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	cmds = append(cmds, validateCmd)

	return cmds
}
