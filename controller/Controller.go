package controller

import "net/http"

type Init interface {
	RoutesInit()
	Get(w http.ResponseWriter, req *http.Request)
	Post(w http.ResponseWriter, req *http.Request)
	Put(w http.ResponseWriter, req *http.Request)
	Delete(w http.ResponseWriter, req *http.Request)
}
