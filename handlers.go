package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Handshake(resp http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(resp).Encode(&HandshakeResp{
		[]string{"VolumeDriver"},
	})
	if err != nil {
		log.Fatal("handshake encode:", err)
		panic(err)
	}
}

func CreateVolume(resp http.ResponseWriter, req *http.Request) {
	var volume VolumeReq

	if err := GetEntity(req, &volume); err != nil {
		// TODO: send error response
		log.Fatal("Failed to parse JSON body")
		return
	}

	VolPath = fmt.Sprintf("%s/%s", ClusterPath, volume.Name)

	err := json.NewEncoder(resp).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("createVolume encode:", err)
	}
}

func RemoveVolume(resp http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(resp).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("removeVolume encode:", err)
		return
	}
}

func MountVolume(resp http.ResponseWriter, req *http.Request) {
	var volume VolumeReq

	if err := GetEntity(req, &volume); err != nil {
		// TODO: send error response
		log.Fatal("Failed to parse JSON body")
		return
	}

	err := os.Mkdir(VolPath, 0777)

	if err != nil {
		log.Fatal("Failed to create volume dir")
		json.NewEncoder(resp).Encode(&MountResp{
			"",
			err.Error(),
		})
		return
	}

	err = json.NewEncoder(resp).Encode(&MountResp{
		VolPath,
		"",
	})
	if err != nil {
		log.Fatal("mountVolume encode:", err)
	}
}

func UnmountVolume(resp http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(resp).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("unmountVolume encode:", err)
		return
	}
}

func VolumePath(resp http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(resp).Encode(&MountResp{
		VolPath,
		"",
	})
	if err != nil {
		log.Fatal("volumePath encode:", err)
		return
	}
}
