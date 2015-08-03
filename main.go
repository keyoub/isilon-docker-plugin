package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/calavera/dkvolume"
	"github.west.isilon.com/bkeyoumarsi/docker-plugin/driver"
)

const (
	socketAddress = "/run/docker/plugins/isilon.sock"
	mountPoint    = "/tmp/isilon/volumes/"
)

func main() {
	var (
		clusterIP string
	)

	flag.StringVar(&clusterIP, "cluster-ip", "",
		"Isilon cluster ip address <x.x.x.x>")

	flag.Parse()

	err := mountCluster(clusterIP)
	defer unMountCluster()
	if err != nil {
		log.Panic(err.Error())
	}

	d := driver.NewIsilonDriver()
	handler := dkvolume.NewHandler(d)
	log.Printf("listening on %s\n", socketAddress)
	log.Fatal(handler.ServeUnix("root", socketAddress))
}

func mountCluster(clusterIP string) error {
	if _, err := os.Stat(mountPoint); os.IsNotExist(err) {
		log.Println("Creating temp mount point")
		if err := os.MkdirAll(mountPoint, 0755); err != nil {
			return errors.New("Failed to create mountpoint directory")
		}
		mountPath := fmt.Sprintf("%s:/ifs/data/docker/volumes", clusterIP)
		cmd := exec.Command("mount", "-t", "nfs", mountPath, mountPoint)

		err := cmd.Run()
		if err != nil {
			return errors.New("Failed to mount Isilon cluster")
		}
	}
	return nil
}

func unMountCluster() {
	// unmount cluster
	cmd := exec.Command("umount", mountPoint)
	err := cmd.Run()
	if err != nil {
		log.Fatal("Failed to unmount cluster")
		return
	}

	err = os.RemoveAll("/tmp/isilon")
	if err != nil {
		log.Fatal("Failed to delete tmp location for isi-volumes")
	}
}
