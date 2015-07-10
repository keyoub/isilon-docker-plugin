package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Activate",
		"POST",
		"/Plugin.Activate",
		Handshake,
	},
	Route{
		"CreateVolume",
		"POST",
		"/VolumeDriver.Create",
		CreateVolume,
	},
	Route{
		"RemoveVolume",
		"POST",
		"/VolumeDriver.Remove",
		RemoveVolume,
	},
	Route{
		"MountVolume",
		"POST",
		"/VolumeDriver.Mount",
		MountVolume,
	},
	Route{
		"UnmountVolume",
		"POST",
		"/VolumeDriver.Unmount",
		UnmountVolume,
	},
	Route{
		"VolumePath",
		"POST",
		"/VolumeDriver.Path",
		VolumePath,
	},
}
