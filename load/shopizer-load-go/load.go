package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Server points to a hostname, port where shopizer runs
type Server struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

// Endpoint represents a single shopizer endpoint
type Endpoint struct {
	URL    string `json:"url"`
	Body   string `json:"body"`
	Method string `json:"method"`
}

func readEndpoints() *[]Endpoint {
	file := "./endpoints.json"
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}

	var endpoints *[]Endpoint
	err = json.Unmarshal(raw, &endpoints)
	if err != nil {
		panic(err.Error())
	}
	return endpoints
}

func readServers(serverfile string) *[]Server {
	file := serverfile
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}

	var servers *[]Server
	err = json.Unmarshal(raw, &servers)
	if err != nil {
		panic(err.Error())
	}
	return servers
}

func main() {

	serverfile := os.Args[1]

	servers := *readServers(serverfile)
	endpoints := *readEndpoints()

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for {

		for _, server := range servers {
			for _, endpoint := range endpoints {

				time.Sleep(100 * time.Millisecond)
				url := fmt.Sprintf("http://%s:%d/%s", server.Hostname, server.Port, endpoint.URL)

				fmt.Printf("%s:%d - %s\n", server.Hostname, server.Port, endpoint.URL)
				req, err := http.NewRequest(endpoint.Method, url, nil)
				if err != nil {
					fmt.Printf("Error %s\n", err.Error())
				} else {

					resp, err := client.Do(req)

					if err != nil {
						fmt.Printf("Error %s\n", err.Error())
					} else {
						fmt.Printf("%v\n", resp.Status)
					}

				}
			}

		}

	}

}
