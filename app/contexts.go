// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=local/PA_API/design
// --out=$(GOPATH)\src\local\pa_api
// --version=v1.0.0
//
// API "PANOS API": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// GetKeygenContext provides the keygen get action context.
type GetKeygenContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Password *string
	User     *string
}

// NewGetKeygenContext parses the incoming request URL and body, performs validations and creates the
// context used by the keygen controller get action.
func NewGetKeygenContext(ctx context.Context, service *goa.Service) (*GetKeygenContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetKeygenContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPassword := req.Params["password"]
	if len(paramPassword) > 0 {
		rawPassword := paramPassword[0]
		rctx.Password = &rawPassword
	}
	paramUser := req.Params["user"]
	if len(paramUser) > 0 {
		rawUser := paramUser[0]
		rctx.User = &rawUser
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetKeygenContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/xml")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// GetUserIDContext provides the user-id get action context.
type GetUserIDContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Action *string
	Cmd    *string
	Key    *string
}

// NewGetUserIDContext parses the incoming request URL and body, performs validations and creates the
// context used by the user-id controller get action.
func NewGetUserIDContext(ctx context.Context, service *goa.Service) (*GetUserIDContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := GetUserIDContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAction := req.Params["action"]
	if len(paramAction) > 0 {
		rawAction := paramAction[0]
		rctx.Action = &rawAction
	}
	paramCmd := req.Params["cmd"]
	if len(paramCmd) > 0 {
		rawCmd := paramCmd[0]
		rctx.Cmd = &rawCmd
	}
	paramKey := req.Params["key"]
	if len(paramKey) > 0 {
		rawKey := paramKey[0]
		rctx.Key = &rawKey
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetUserIDContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/xml")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
