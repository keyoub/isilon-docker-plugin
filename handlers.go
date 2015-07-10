package main

import (
	"encoding/json"
	"io/ioutil"
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
	log.Println("Handshake complete!")
}

func CreateVolume(w http.ResponseWriter, r *http.Request) {
	var volume VolumeReq

	// Parse request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Failed to read request body:", err)
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &volume); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("createVolume encode:", err)
		return
	}
	log.Println("createVolume complete!")
}

func RemoveVolume(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("removeVolume encode:", err)
		return
	}
	log.Println("removeVolume complete!")
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
	log.Println("mountVolume complete!")
}

func UnmountVolume(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(&ErrResp{
		"",
	})
	if err != nil {
		log.Fatal("unmountVolume encode:", err)
		return
	}
	log.Println("unmountVolume complete!")
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
	log.Println("volumePath complete!")
}
