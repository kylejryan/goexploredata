package cmd

import (
	"fmt"

	"github.com/kylejryan/go-data-explorer/internal/dataframe"
	"github.com/spf13/cobra"
)

var column string

var histCmd = &cobra.Command{
	Use:   "hist [file]",
	Short: "Generate histogram for a specified column in a CSV file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		df, err := dataframe.Load(filePath)
		if err != nil {
			fmt.Println("Error loading file:", err)
			return
		}

		if column == "" {
			fmt.Println("Please specify a column using the --column flag.")
			return
		}

		err = df.Histogram(column)
		if err != nil {
			fmt.Println("Error generating histogram:", err)
			return
		}
	},
}

func init() {
	histCmd.Flags().StringVarP(&column, "column", "c", "", "Column name to generate histogram for")
	rootCmd.AddCommand(histCmd)
}
