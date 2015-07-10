package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Handshake(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&HandshakeResp{
		[]string{"VolumeDriver"},
	})
	if err != nil {
		log.Fatal("handshake encode:", err)
		panic(err)
	}
}

func CreateVolume(w http.ResponseWriter, r *http.Request) {
	var volume VolumeReq

	if err := GetEntity(r, &volume); err != nil {
		// TODO: send error response
		log.Fatal("Failed to parse JSON body")
	}

	log.Printf("Volume Name: %s", volume.Name)

	err := json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("createVolume encode:", err)
	}
}

func RemoveVolume(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("removeVolume encode:", err)
		return
	}
}

func MountVolume(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&MountResp{
		"/tmp/testvolume",
		"",
	})
	if err != nil {
		log.Fatal("mountVolume encode:", err)
		return
	}
}

func UnmountVolume(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("unmountVolume encode:", err)
		return
	}
}

func VolumePath(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&MountResp{
		"/tmp/testvolume",
		"",
	})
	if err != nil {
		log.Fatal("volumePath encode:", err)
		return
	}
}
