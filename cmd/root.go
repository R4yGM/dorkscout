package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "dorkscout",
	DisableSuggestions : false,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {

    }
}