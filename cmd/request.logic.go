package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type RequestType string

func (reqType RequestType) String() string {
	return string(reqType)
}

const (
	GET    RequestType = "GET"
	POST   RequestType = "POST"
	PUT    RequestType = "PUT"
	DELETE RequestType = "DELETE"
)

func Request(url string, query []string, reqBody []byte, method RequestType) error {
	if method == GET {
		if err := get(url, query); err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	} else {
		if err := postputdelete(url, reqBody, method); err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
	}
	return nil
}

func get(url string, query []string) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()

	for _, v := range query {
		s := strings.Split(v, "=")
		q.Add(s[0], s[1])
	}

	req.URL.RawQuery = q.Encode()

	for _, v := range header {
		s := strings.Split(v, "=")
		req.Header.Set(s[0], s[1])
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	h, err := Map2JSON(res.Header)
	if err != nil {
		return err
	}
	fmt.Printf(`{"header":%s,"body":%s}`, string(h), string(b))
	return nil
}

func postputdelete(url string, reqBody []byte, method RequestType) error {

	req, err := http.NewRequest(method.String(), url, bytes.NewBuffer(reqBody))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf(`{"status":"%s","body":%s}`, resp.Status, string(body))
	return nil
}
