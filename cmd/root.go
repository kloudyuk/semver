package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var red = color.New(color.FgRed)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "A CLI for working with semantic versions based on masterminds/semver",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		red.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
