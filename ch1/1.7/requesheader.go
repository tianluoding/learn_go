package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "http://localhost:8000"
	payload := strings.NewReader("username=value1&password=value2")

	req, err := http.NewRequest("GET", url, payload)
	if err != nil {
		// handle error
		fmt.Fprintf(os.Stderr, "new request err: %v", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;...")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Fprintf(os.Stderr, "Do err: %v", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
