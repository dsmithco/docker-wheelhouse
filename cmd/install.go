package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Software",
	Long:  `Install Software`,
	Run:   installRun,
	Args:  cobra.ExactArgs(1),
}

func installRun(cmd *cobra.Command, args []string) {

	arg := args[0]

	if strings.ToLower(arg) == "docker" {
		fmt.Println("Checking volume available space")
		fmt.Println("Checking Docker version")
		fmt.Println("Installing Docker")
	}

	if strings.ToLower(arg) == "awscli" {
		fmt.Println("Checking AWS CLI version")
		fmt.Println("Installing AWS CLI")
	}

	if strings.ToLower(arg) == "kubernetes" {
		fmt.Println("Checking Kubernetes version")
		fmt.Println("Installing Kubernetes")
	}

}
