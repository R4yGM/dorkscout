package cmd

import (
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "dorkscout",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	DisableSuggestions : false,
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {

    }
}