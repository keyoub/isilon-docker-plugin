package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	address := []byte("http://127.0.0.1:8080")
	err := ioutil.WriteFile(
		"/usr/share/docker/plugins/isilon.spec", address, 0777)
	if err != nil {
		log.Fatal(err)
	}

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
