package cmd

import (
	"github.com/HanksJCTsai/XReader/internal/logger"
	"github.com/HanksJCTsai/XReader/internal/read"
	"github.com/spf13/cobra"
)

var infoCMD = &cobra.Command{
	Use:     "info",
	Aliases: []string{"i", "I"},
	Short:   "Display file information",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			logger.Log.Error("Please provide a file path!")
			cmd.Usage()
			return
		}
		if err := read.GetFileInfo(filePath, false); err != nil {
			logger.Log.Error("Got file info failed: %v", err)
		}
	},
}

func init() {
	infoCMD.Flags().StringP("file", "f", "", "Define file path")
	infoCMD.MarkFlagRequired("file")
}
