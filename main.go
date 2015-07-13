package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	ClusterPath string
	VolPath     string
)

func main() {

	flag.StringVar(&ClusterPath, "path", "",
		"Local directory path where Isilon cluster is mounted on")

	flag.Parse()

	if _, err := os.Stat(ClusterPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", ClusterPath)
		os.Exit(1)
	}

	address := []byte("http://127.0.0.1:8080")
	err := ioutil.WriteFile(
		"/usr/share/docker/plugins/isilon.spec", address, 0777)
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
