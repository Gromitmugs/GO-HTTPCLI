package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Sending HTTP GET Request",
	Long:  `Sending HTTP GET Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		body := str2JSON(json_input)
		postputdeleteREQ(url, body, "DELETE")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&json_input, "json", "", "", "JSON String Input")
}
