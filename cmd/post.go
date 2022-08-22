package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
		body := str2JSON(json_input)
		postputdeleteREQ(url, body, "POST")
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.PersistentFlags().StringVarP(&json_input, "json", "", "", "JSON String Input")
}

func postputdeleteREQ(url string, req_body []byte, method string) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(req_body))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Print(resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	fmt.Printf(`{"status":"%s","body":%s}`, resp.Status, string(body))
}

func str2JSON(str string) []byte {

	var map2 map[string]interface{}
	err := json.Unmarshal([]byte(json_input), &map2)
	if err != nil {
		log.Fatal("Wrong JSON Format!")
		return nil
	}

	b, _ := json.Marshal(&map2)
	return b
}
