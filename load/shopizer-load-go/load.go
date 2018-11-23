package main

import (
		"fmt"
		"io/ioutil"
	"encoding/json"
	"time"
	"net/http"
	"os"
)



type Server struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type Endpoint struct {
	URL    string `json:"url"`
	Body   string `json:"body"`
	Method string `json:"method"`
}

func readEndpoints() *[]Endpoint{
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

func readServers(serverfile string) *[]Server{
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


func main(){


	serverfile := os.Args[1]

	servers := *readServers(serverfile)
	endpoints := *readEndpoints()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}


	for {

		for _, server := range servers {
			for _, endpoint := range endpoints{

				url :=  fmt.Sprintf("http://%s:%d/%s", server.Hostname, server.Port, endpoint.URL)

				fmt.Printf("%s:%d - %s\n", server.Hostname, server.Port, endpoint.URL)
				req, _ := http.NewRequest(endpoint.Method, url, nil )

				time.Sleep(100 * time.Millisecond)
				resp, err := client.Do(req)
				if err != nil {
					fmt.Printf("Error %s\n", err.Error())
				}
				fmt.Printf("%v\n", resp.Status)

			}

		}

	}

}