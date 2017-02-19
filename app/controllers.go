// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=local/PA_API/design
// --out=$(GOPATH)\src\local\pa_api
// --version=v1.0.0
//
// API "PANOS API": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewXMLEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// KeygenController is the controller interface for the Keygen actions.
type KeygenController interface {
	goa.Muxer
	Get(*GetKeygenContext) error
}

// MountKeygenController "mounts" a Keygen resource controller on the given service.
func MountKeygenController(service *goa.Service, ctrl KeygenController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetKeygenContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/api/keygen", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "Keygen", "action", "Get", "route", "GET /api/keygen")
}

// UserIDController is the controller interface for the UserID actions.
type UserIDController interface {
	goa.Muxer
	Get(*GetUserIDContext) error
}

// MountUserIDController "mounts" a UserID resource controller on the given service.
func MountUserIDController(service *goa.Service, ctrl UserIDController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetUserIDContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	service.Mux.Handle("GET", "/api/user-id", ctrl.MuxHandler("Get", h, nil))
	service.LogInfo("mount", "ctrl", "UserID", "action", "Get", "route", "GET /api/user-id")
}
