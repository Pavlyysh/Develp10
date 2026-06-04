package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: html-loader <url1> <url2>\nExample: html-loader google.com")
		os.Exit(1)
	}

	urls := os.Args[1:]

	for _, url := range urls {
		file, err := os.OpenFile(url+".html", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("unable to create file", err)
			os.Exit(1)
		}
		defer file.Close()

		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("invalid url %s\n", url)
			os.Exit(1)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("status code is not ok %d\n", resp.StatusCode)
			os.Exit(1)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("unable to read the body")
			os.Exit(1)
		}

		file.WriteString(string(body))
	}
}
