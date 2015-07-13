package main

type HandshakeResp struct {
	Implements []string
}

type ErrResp struct {
	Err string
}

type MountResp struct {
	Mountpoint string
	Err        string
}
