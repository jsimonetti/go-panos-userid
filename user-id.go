package main

import (
	"encoding/xml"
	"fmt"
	"local/pa_api/app"
	"strings"

	"github.com/goadesign/goa"
)

// UserIDController implements the user-id resource.
type UserIDController struct {
	*goa.Controller
}

// NewUserIDController creates a user-id controller.
func NewUserIDController(service *goa.Service) *UserIDController {
	return &UserIDController{Controller: service.NewController("UserIDController")}
}

// Get runs the get action.
func (c *UserIDController) Get(ctx *app.GetUserIDContext) error {
	// UserIDController_Get: start_implement

	switch ctx.Params.Get("action") {
	case "set":
		Set(ctx)
	}

	// UserIDController_Get: end_implement
	return nil
}

// uidSetMessage holds the user-id set message
type uidSetMessage struct {
	XMLName xml.Name  `xml:"uid-message"`
	Version string    `xml:"version"`
	Type    string    `xml:"type"`
	Login   *userInfo `xml:"payload>login>entry,omit-empty"`
	Logout  *userInfo `xml:"payload>logout>entry,omit-empty"`
}

type userInfo struct {
	Action string
	Name   string `xml:"name,attr"`
	SSID   string
	IP     string `xml:"ip,attr"`
}

func Set(ctx *app.GetUserIDContext) {

	cmd := ctx.Params.Get("cmd")

	var msg uidSetMessage
	err := xml.Unmarshal([]byte(cmd), &msg)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	user := msg.Login
	if msg.Login == nil {
		user = msg.Logout
		user.Action = "logout"
	} else {
		user.Action = "login"
	}
	name := strings.Split(user.Name, "@")
	user.Name = name[0]
	user.SSID = name[1]
	fmt.Printf("type: %s, action: %s, name: %s, ssid: %s, ip: %s\n", msg.Type, user.Action, user.Name, user.SSID, user.IP)

}
