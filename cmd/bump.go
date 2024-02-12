package cmd

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
)

var level string

// bumpCmd represents the bump command
var bumpCmd = &cobra.Command{
	Use:   "bump VERSION",
	Short: "Takes a given version and bumps it by the desired level (major|minor|patch)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		v, err := semver.NewVersion(args[0])
		if err != nil {
			return err
		}
		var bumped semver.Version
		switch strings.ToLower(level) {
		case "major":
			bumped = v.IncMajor()
		case "minor":
			bumped = v.IncMinor()
		case "patch":
			bumped = v.IncPatch()
		default:
			return fmt.Errorf("invalid level: %s", level)
		}
		fmt.Println(bumped.String())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bumpCmd)
	bumpCmd.Flags().StringVarP(&level, "level", "l", "minor", "level to bump - major|minor|patch")
}
