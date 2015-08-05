package rest

import "testing"

// Tests both CreateVolume and CheckVolume
func TestCreateVolume(t *testing.T) {
	c := NewClient("10.28.102.200", "root", "a")
	c.CreateVolume("scriptTest")

	b, err := c.CheckVolume("scriptTest")
	if err != nil || !b {
		t.Fail()
	}
}
