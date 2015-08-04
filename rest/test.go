package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ownership struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type aclRequest struct {
	Authoritative string    `json:"authoritative"`
	Action        string    `json:"action"`
	Owner         ownership `json:"owner"`
	Group         ownership `json:"group"`
}

func main() {
	var u = "https://10.28.102.200:8080/namespace/ifs/data/docker/volumes/sc_test6/"
	fmt.Println(u)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	req.SetBasicAuth("root", "a")
	req.Header.Add("x-isi-ifs-target-type", "container")
	req.Header.Add("x-isi-ifs-access-control", "0755")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode != 200 {
		fmt.Printf("Bad status: %d\n", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("err: %s\n", string(body))
		return
	}

	var data = aclRequest{
		"acl",
		"update",
		ownership{"UID:65534", "nobody", "user"},
		ownership{"GID:65534", "nobody", "group"},
	}

	var u2 = fmt.Sprintf("%s?acl", u)

	b, err := json.Marshal(data)
	fmt.Println(string(b))

	req, err = http.NewRequest("PUT", u2, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(req.URL)

	req.SetBasicAuth("root", "a")

	resp, err = client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Status: %d\n", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("err: %s\n", string(body))
}
