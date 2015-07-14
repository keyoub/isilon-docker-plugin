package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/calavera/dkvolume"
	"github.west.isilon.com/bkeyoumarsi/docker-plugin/driver"
)

const socketAddress = "/usr/share/docker/plugins/isilon.sock"

func main() {
	var clusterPath string

	flag.StringVar(&clusterPath, "path", "",
		"Local directory path where Isilon cluster is mounted on")

	flag.Parse()

	if _, err := os.Stat(clusterPath); os.IsNotExist(err) {
		fmt.Printf("no such file or directory: %s\n", clusterPath)
		os.Exit(1)
	}

	d := driver.NewIsilonDriver(clusterPath)
	handler := dkvolume.NewHandler(d)
	log.Printf("listening on %s\n", socketAddress)
	log.Fatal(handler.ServeUnix("root", socketAddress))
}
