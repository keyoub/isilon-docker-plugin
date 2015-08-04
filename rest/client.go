package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	volumesPath      = "/ifs/docker/volumes"
	volumeCreatePath = "/ifs/docker/volumes/"
	volumeStopPath   = "/api/1.0/volume/%s/stop"
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

type response struct {
	Ok  bool   `json:"ok"`
	Err string `json:"error,omitempty"`
}

type peerResponse struct {
	//Data []peer `json:"data",omitempty`
	response
}

type volumeResponse struct {
	//Data []volume `json:"data",omitempty`
	response
}

type Client struct {
	addr string
	usr  string
	pass string
}

func NewClient(addr, usr, pass string) *Client {
	return &Client{addr, usr, pass}
}

func (r Client) VolumeExist(name string) (bool, error) {
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

func (r Client) CreateVolume(name string, peers []string) error {
	u := fmt.Sprintf("%s:8080/namespace%s", r.addr,
		fmt.Sprintf("%s%s", volumeCreatePath, name))
	fmt.Println(u)

	req, err := http.NewRequest("PUT", u, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

	return responseCheck(resp)
}

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

func responseCheck(resp *http.Response) error {
	var p response
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return err
	}

	if !p.Ok {
		return fmt.Errorf(p.Err)
	}

	return nil
}
