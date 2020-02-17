package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(auditCmd)
}

var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "aws-nuke",
	Long:  "aws-nuke",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Audting")
	},
}
