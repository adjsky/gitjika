package cmd

import (
	"fmt"
	"os"

	"github.com/adjsky/gitjika/internal/config"
	"github.com/spf13/cobra"
)

func Main() int {
	var cfgFile string

	root := &cobra.Command{
		Use:   "gitjika",
		Short: "",
		Long:  "",
	}

	root.PersistentFlags().StringVar(&cfgFile, "config",
		"", "Config file (default is $HOME/.gitjika/config.toml or /etc/gitjika/config.toml)")

	root.SilenceUsage = true
	root.SilenceErrors = true

	cfg := config.New(cfgFile)

	root.AddCommand(NewServeCmd(cfg))

	err := root.Execute()

	if err == nil {
		return 0
	}

	fmt.Fprintf(os.Stderr, "Error: %s\n", err)

	return 1
}
