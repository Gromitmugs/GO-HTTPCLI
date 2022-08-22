package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Sending HTTP GET Request",
	Long:  `Sending HTTP GET Request`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := string(args[0])
		getREQ(url, slice_query)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getREQ(url string, slice_query []string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()

	for _, v := range slice_query {
		s := strings.Split(v, "=")
		q.Add(s[0], s[1])
	}

	req.URL.RawQuery = q.Encode()

	for _, v := range slice_header {
		s := strings.Split(v, "=")
		req.Header.Set(s[0], s[1])
	}

	res, _ := client.Do(req)
	b, _ := io.ReadAll(res.Body)
	h := map2JSON(res.Header)
	fmt.Printf(`{"header":%s,"body":%s}`, string(h), string(b))
}
