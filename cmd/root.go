package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "obsidian-cli",
		Short:   "obsidian-cli - CLI to open, search, move, create and update notes",
		Version: "v0.1.7",
		Long:    "obsidian-cli - CLI to open, search, move, create and update notes",
	}
	cfg interface{}
)

// Execute executes the root command with the provided configuration
func Execute(config interface{}) error {
	cfg = config
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
	return nil
}

// GetConfig returns the current configuration
func GetConfig() interface{} {
	return cfg
}
