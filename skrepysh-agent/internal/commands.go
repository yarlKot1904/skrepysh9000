package internal

import (
	"fmt"
	"os"
	"os/signal"
	"skrepysh-agent/pkg/client"
	"skrepysh-agent/pkg/config"
	"skrepysh-agent/pkg/server"
	"syscall"

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

			c, err := client.New(&conf.SkrepyshBackend)
			fmt.Printf(c.IP)
			fmt.Printf(c.OS)

			if err != nil {
				return err
			}
			err = c.Init()
			if err != nil {
				return err
			}

			sigs := make(chan os.Signal, 1)
			signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

			go func(log *zap.Logger, sigs chan os.Signal, c *client.Client) {
				s := <-sigs
				fmt.Printf(s.String())
				log.Info("deleting vm from database")
				var err error
				for i := 1; i <= 3; i++ {
					err = c.Delete()
					if err != nil {
						log.Error("error deleting vm from database", zap.Int("Attempt", i), zap.Error(err))
					} else {
						os.Exit(0)
					}
				}
				os.Exit(-1)
			}(log, sigs, c)

			return server.Serve(log, conf)
		},
	}
	serveCmd.Flags().StringVarP(&configPath, "config", "c", configPath, "path/to/config")

	cmds = append(cmds, serveCmd)

	return cmds
}
