package cmd

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
)

var reverse bool

// sortCmd represents the sort command
var sortCmd = &cobra.Command{
	Use:   "sort VERSION VERSION...",
	Short: "Sort semantic versions in accordance with the SemVer 2 specification",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		vs := make([]*semver.Version, len(args))
		for i, r := range args {
			v, err := semver.NewVersion(r)
			if err != nil {
				return err
			}
			vs[i] = v
		}
		if reverse {
			sort.Sort(sort.Reverse(semver.Collection(vs)))
		} else {
			sort.Sort(semver.Collection(vs))
		}
		for _, v := range vs {
			fmt.Println(v.String())
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sortCmd)
	sortCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "reverse sort order (default is oldest to newest)")
}
