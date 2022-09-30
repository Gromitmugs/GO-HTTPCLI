package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Sending HTTP GET Request",
	Long:  `Sending HTTP GET Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		Request(url, query, nil, GET)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
