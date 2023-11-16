package cmd

import (
	"crawlnovel/cmd/convert"
	"crawlnovel/cmd/download"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "download",
	Short:             "download API server",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `Start download API server`,
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
}

func init() {
	rootCmd.AddCommand(download.StartCmd)
	rootCmd.AddCommand(convert.StartCmd)
}

// Execute : run commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
