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

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	fmt.Println("Arguments received:", os.Args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define persistent flags or configuration settings if needed.
}
