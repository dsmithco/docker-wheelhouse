package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

}

var RootCmd = &cobra.Command{
	Use:   "wheelhouse",
	Short: `Everything you need to setup, run, push, pull,
			deploy, & monitor your container environments`,
	Long: `Everything you need to setup, run, push, pull, deploy,
& monitor your container environments`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}
