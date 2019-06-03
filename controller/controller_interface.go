package controller

import "net/http"

type Controller interface {
	Index(w http.ResponseWriter, r *http.Request)
}
