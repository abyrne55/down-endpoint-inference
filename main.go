package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// readEndpointFile loads a list of test endpoints from a file
// Adapted from https://stackoverflow.com/a/18479916
func readEndpointFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	endpointFilePath := "endpoints"
	if len(os.Args) >= 2 {
		endpointFilePath = os.Args[1]
	}

	fmt.Printf("Loading endpoints file '%s'...\n", endpointFilePath)
	endpoints, err := readEndpointFile(endpointFilePath)
	if err != nil {
		log.Fatal(err)
	}

	httpClient := http.Client{Timeout: 2000 * time.Millisecond}
	for _, endpoint := range endpoints {
		fmt.Printf("Testing %s... ", endpoint)
		resp, err := httpClient.Get("http://" + endpoint)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println(resp.Status)
		}
	}
}
