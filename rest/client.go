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

/*func (r Client) VolumeExist(name string) (bool, error) {
	vols, err := r.volumes()
	if err != nil {
		return false, err
	}

	for _, v := range vols {
		if v.Name == name {
			return true, nil
		}
	}

	return false, nil
}

func (r Client) volumes() ([]volume, error) {
	u := fmt.Sprintf("%s%s", r.addr, volumesPath)

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	var d volumeResponse
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, err
	}

	if !d.Ok {
		return nil, fmt.Errorf(d.Err)
	}
	return d.Data, nil
}
*/
func (r Client) CreateVolume(name string) error {
	u := fmt.Sprintf("https://%s:8080/namespace%s", r.addr,
		fmt.Sprintf("%s%s/", volumesPath, name))
	log.Println(u)

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
	req.Header.Add("x-isi-ifs-access-control", "0765")

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

/*
func (r Client) StopVolume(name string) error {
	u := fmt.Sprintf("%s%s", r.addr, fmt.Sprintf(volumeStopPath, name))

	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return responseCheck(resp)
}
*/
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
