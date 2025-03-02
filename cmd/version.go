package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const versionNumber = "0.2.0"

var versionCMD = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "V"},
	Short:   "Display XReader version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("XReader version: %s\n", versionNumber)
	},
}
