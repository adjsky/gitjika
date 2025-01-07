package cmd

import (
	"github.com/adjsky/gitjika/internal/config"
	"github.com/adjsky/gitjika/server"
	"github.com/spf13/cobra"
)

func NewServeCmd(cfg config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run server",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			server := server.New(cfg)
			return server.Listen()
		},
	}

	return cmd
}
