/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Sending HTTP PUT Request",
	Long:  `Sending HTTP PUT Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		body := str2JSON(json_input)
		postputdeleteREQ(url, body, "PUT")
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	putCmd.PersistentFlags().StringVarP(&json_input, "json", "", "", "JSON String Input")
}
