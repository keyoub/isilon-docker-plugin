package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	pluginSpecDir = "/usr/share/docker/plugins"
	tcpAddress    = "http://127.0.0.1:8080"
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

	err := writeSpec("isilon", tcpAddress)
	if err != nil {
		log.Fatal("Failed to write spec file.")
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func writeSpec(name, addr string) error {
	spec := filepath.Join(pluginSpecDir, name+".spec")
	url := "tcp://" + addr
	return ioutil.WriteFile(spec, []byte(url), 0644)
}
