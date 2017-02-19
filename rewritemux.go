package main

import (
	"net/http"
	"strings"

	"github.com/goadesign/goa"
)

type mux struct {
	goa goa.ServeMux
}

func NewMux() *mux {
	return &mux{
		goa: goa.NewMux(),
	}
}

func (m *mux) Handle(method, path string, handle goa.MuxHandler) {
	m.goa.Handle(method, path, handle)
}

func (m *mux) HandleNotFound(handle goa.MuxHandler) {
	m.goa.HandleNotFound(handle)
}

func (m *mux) Lookup(method, path string) goa.MuxHandler {
	return m.goa.Lookup(method, path)
}

func (m *mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	req.URL.Path = strings.TrimSuffix(req.URL.Path, "/") + "/" + req.FormValue("type")
	req.RequestURI = req.URL.RequestURI()

	m.goa.ServeHTTP(rw, req)
}
