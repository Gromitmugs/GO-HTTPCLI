package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var json_input string

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Sending HTTP POST Request",
	Long:  `Sending HTTP POST Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		body, err := Str2JSON(json_input)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		Request(url, nil, body, POST)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().StringVarP(&json_input, "json", "", "", "JSON String Input")
}
