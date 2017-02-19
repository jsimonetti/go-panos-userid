package main

import (
	"encoding/xml"
	"fmt"
	"local/pa_api/app"

	"github.com/goadesign/goa"
)

// KeygenController implements the keygen resource.
type KeygenController struct {
	*goa.Controller
}

// NewKeygenController creates a keygen controller.
func NewKeygenController(service *goa.Service) *KeygenController {
	return &KeygenController{Controller: service.NewController("KeygenController")}
}

// Get runs the get action.
func (c *KeygenController) Get(ctx *app.GetKeygenContext) error {
	// KeygenController_Get: start_implement

	key := authKey{
		Status: "success",
		Code:   "20",
		Key:    "abcdefghijklmnopqrstuvwxyz123456",
	}

	output, err := xml.Marshal(key)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return ctx.OK(output)

	// KeygenController_Get: end_implement
	return nil
}

// authKey holds our API key.
type authKey struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status,attr"`
	Code    string   `xml:"code,attr"`
	Key     string   `xml:"result>key"`
}
