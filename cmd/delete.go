package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Sending HTTP GET Request",
	Long:  `Sending HTTP GET Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		body, err := Str2JSON(json_input)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		Request(url, nil, body, DELETE)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&json_input, "json", "", "", "JSON String Input")
}
