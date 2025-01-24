package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godataframe",
	Short: "A CLI tool for dataframe operations in Go",
	Long:  `GoDataFrame allows you to convert files into dataframes and perform statistical analyses directly from the command line.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Define flags here
}
