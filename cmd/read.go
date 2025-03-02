package cmd

import (
	"github.com/HanksJCTsai/XReader/internal/logger"
	"github.com/HanksJCTsai/XReader/internal/read"
	"github.com/spf13/cobra"
)

var ReadCMD = &cobra.Command{
	Use:     "read",
	Aliases: []string{"r", "R"},
	Short:   "Read file contents",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			logger.Log.Error("Please provide a file path!")
			cmd.Usage()
			return
		}
		if err := read.ReadFile(filePath); err != nil {
			logger.Log.Error("File reading failed: %v", err)
		}
	},
}

func init() {
	ReadCMD.Flags().StringP("file", "f", "", "Define file path")
	ReadCMD.MarkFlagRequired("file")
}
