/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var slice_query []string
var slice_header []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "httpcli",
	Short: "HTTP client command-line program ",
	Long:  `HTTP client command-line program`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		getREQ(url, slice_query)
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
	rootCmd.PersistentFlags().StringSliceVarP(&slice_query, "query", "", []string{}, "Query Input [key1=val1]")
	rootCmd.PersistentFlags().StringSliceVarP(&slice_header, "header", "", []string{}, "Header Input [key1=val1]")
}

func slice2map(slice []string) map[string]string {
	map1 := make(map[string]string)
	for _, v := range slice {
		s := strings.Split(v, "=")
		map1[s[0]] = s[1]
	}
	return map1
}

func slice2JSON(slice []string) []byte {
	map1 := slice2map(slice)
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStr
}

func map2JSON(map1 map[string][]string) []byte {
	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStr
}
