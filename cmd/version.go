package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "aws-nuke",
	Long:  "aws-nuke",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Aws-Nuke version 1.0")
	},
}
