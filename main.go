package main

import (
	"flag"
	"log"

	"github.com/calavera/dkvolume"
	"github.west.isilon.com/bkeyoumarsi/isilon-docker-plugin/driver"
)

const (
	socketAddress = "/run/docker/plugins/isilon.sock"
)

func main() {
	var (
		clusterAddress string
		username       string
		password       string
	)

	flag.StringVar(&clusterAddress, "cluster-ip", "",
		"Isilon cluster ip address <x.x.x.x>")

	flag.StringVar(&username, "username", "",
		"Admin username")

	flag.StringVar(&password, "password", "",
		"Admin password")

	flag.Parse()

	d := driver.NewIsilonDriver(clusterAddress, username, password)
	handler := dkvolume.NewHandler(d)
	log.Printf("listening on %s\n", socketAddress)
	log.Fatal(handler.ServeUnix("root", socketAddress))
}
