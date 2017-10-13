package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Software",
	Long:  `Update Software`,
	Run:   updateRun,
	Args:  cobra.ExactArgs(1),
}

func updateRun(cmd *cobra.Command, args []string) {

	arg := args[0]

	if strings.ToLower(arg) == "docker" {
		fmt.Println("Checking volume available space")
		fmt.Println("Checking Docker version")
		fmt.Println("Updating Docker")
	}

	if strings.ToLower(arg) == "awscli" {
		fmt.Println("Checking AWS CLI version")
		fmt.Println("Updating AWS CLI")
	}

	if strings.ToLower(arg) == "kubernetes" {
		fmt.Println("Checking Kubernetes version")
		fmt.Println("Updating Kubernetes")
	}

}
