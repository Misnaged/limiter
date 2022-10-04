package handlers

import "net/http"

type IHandler interface {
	Registration(login, pass string) func(w http.ResponseWriter, r *http.Request)
}

type Handler struct{}
