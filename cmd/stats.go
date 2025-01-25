// /cmd/stats.go
package cmd

import (
	"fmt"

	"github.com/kylejryan/go-data-explorer/internal/dataframe"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats [file]",
	Short: "Get statistical summary of a CSV file",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		df, err := dataframe.Load(filePath)
		if err != nil {
			fmt.Println("Error loading file:", err)
			return
		}

		summary, err := df.Describe()
		if err != nil {
			fmt.Println("Error generating statistics:", err)
			return
		}

		fmt.Println(summary)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
