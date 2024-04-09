package controller

import "net/http"

type Controller struct {
	*http.ServeMux
}
