package main

import "net/http"

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
	return v.UnmarshalHTTP(r)
}
