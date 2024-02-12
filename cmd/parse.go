package cmd

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse VERSION",
	Short: "Parses a version in accordance with the SemVer 2 specification and returns the full version or an error",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := semver.NewVersion(args[0])
		if err != nil {
			return err
		}
		fmt.Println(v.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
