package rest

import (
	"log"
	"testing"
)

// Tests both CreateVolume and CheckVolume
func TestCreateVolume(t *testing.T) {
	c := NewClient("10.28.102.200", "root", "a")
	err := c.CreateVolume("test")
	if err != nil {
		log.Println(err.Error())
		t.Fail()
	}

	b, err := c.CheckVolume("test")
	if err != nil || !b {
		log.Println(err.Error())
		t.Fail()
	}
}
