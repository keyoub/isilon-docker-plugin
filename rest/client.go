package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	volumesPath = "/ifs/data/docker/volumes/"
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

type volume struct {
	Name string `json:"name"`
}

type volumeDirectory struct {
	Children []volume `json:"children"`
}

type Client struct {
	httpClient *http.Client
	addr       string
	usr        string
	pass       string
}

func NewClient(addr, usr, pass string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{&http.Client{Transport: tr}, addr, usr, pass}
}

func (r Client) CheckVolume(name string) (bool, error) {
	volumes, err := r.getVolumes()
	if err != nil {
		return false, err
	}

	for _, v := range volumes.Children {
		if v.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func (r Client) CreateVolume(name string) error {
	u := fmt.Sprintf("https://%s:8080/namespace%s", r.addr,
		fmt.Sprintf("%s%s/", volumesPath, name))

	if err := r.ranCreate(u); err != nil {
		return err
	}
	if err := r.ranUpdatePerm(fmt.Sprintf("%s?acl", u)); err != nil {
		return err
	}

	return nil
}

func (r Client) ranCreate(url string) error {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	req.SetBasicAuth(r.usr, r.pass)
	req.Header.Add("x-isi-ifs-target-type", "container")
	req.Header.Add("x-isi-ifs-access-control", "public_read_write")

	resp, err := r.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return responseCheck(resp)
}

func (r Client) ranUpdatePerm(url string) error {
	var data = aclRequest{
		"acl",
		"update",
		ownership{"UID:65534", "nobody", "user"},
		ownership{"GID:65534", "nobody", "group"},
	}

	b, err := json.Marshal(data)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	req.SetBasicAuth(r.usr, r.pass)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return responseCheck(resp)
}

func (r Client) getVolumes() (volumeDirectory, error) {
	var data volumeDirectory
	u := fmt.Sprintf("https://%s:8080/namespace%s", r.addr, volumesPath)

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}
	req.SetBasicAuth(r.usr, r.pass)

	resp, err := r.httpClient.Do(req)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err.Error())
		return data, err
	}

	return data, nil
}

func responseCheck(resp *http.Response) error {
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf("Status: %d\n", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("err: %s\n", string(body))
		return errors.New(string(body))
	}
	return nil
}
