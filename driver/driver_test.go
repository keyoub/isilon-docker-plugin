package driver

import (
	"log"
	"testing"
)
import "github.com/calavera/dkvolume"

func TestDriver(t *testing.T) {
	d := NewIsilonDriver("10.28.102.200", "root", "a")
	req := dkvolume.Request{Name: "test_volume"}
	resp := d.Create(req)
	log.Println(resp.Err)
	resp = d.Mount(req)
	log.Println(resp.Err)
}
