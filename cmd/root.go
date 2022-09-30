package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var query []string
var header []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "GO-HTTPCLI",
	Short: "HTTP client command-line program ",
	Long:  `HTTP client command-line program`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		fmt.Print(url)
		Request(url, query, nil, GET)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringSliceVarP(&query, "query", "", []string{}, "Query Input [key1=val1]")
	rootCmd.PersistentFlags().StringSliceVarP(&header, "header", "", []string{}, "Header Input [key1=val1]")
}
