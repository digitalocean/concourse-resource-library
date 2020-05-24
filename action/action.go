package action

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "crl",
	Short: "crl bootstraps Concourse Resources",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

// Execute executes the root command.
func Execute() error {
	return root.Execute()
}
