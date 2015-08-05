package rest

import "testing"

func TestCreateVolume(t *testing.T) {
	c := NewClient("10.28.102.200", "root", "a")
	c.CreateVolume("scriptTest")
	return
}
