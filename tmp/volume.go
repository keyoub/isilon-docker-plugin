package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type VolumeReq struct {
	Name string
}

func (volume *VolumeReq) UnmarshalHTTP(request *http.Request) error {
	defer request.Body.Close()
	// Parse request body
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal("Failed to read request body:", err)
		return err
	}
	if err = json.Unmarshal(body, &volume); err != nil {
		log.Fatal("Unable to parse JSON body")
		return err
	}
	return nil
}
