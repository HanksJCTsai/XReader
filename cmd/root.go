package cmd

import (
	"fmt"
	"github.com/HanksJCTsai/XReader/internal/config"
	"github.com/HanksJCTsai/XReader/internal/logger"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
	rootCMD = &cobra.Command{Use: "XReader", Short: "XReader is a file processing tool!",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if err := config.LoadConfig(cfgFile); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
				os.Exit(1)
			}
			logger.InitLogger()
		},
	}
)

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		os.Exit(1)
	}
}

func Init() {
	rootCMD.PersistentFlags().StringVar(&cfgFile, "Config", "", "Configuration file path (default: ./config/config.yaml)")
	rootCMD.AddCommand(ReadCMD)
	rootCMD.AddCommand(infoCMD)
	rootCMD.AddCommand(versionCMD)
}
