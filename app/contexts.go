// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "arduino-create-agent": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/arduino/arduino-create-agent/design
// --out=$(GOPATH)/src/github.com/arduino/arduino-create-agent
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
	"strconv"
)

// ExecCommandsV1Context provides the commands_v1 exec action context.
type ExecCommandsV1Context struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      string
	Payload ExecCommandsV1Payload
}

// NewExecCommandsV1Context parses the incoming request URL and body, performs validations and creates the
// context used by the commands_v1 controller exec action.
func NewExecCommandsV1Context(ctx context.Context, r *http.Request, service *goa.Service) (*ExecCommandsV1Context, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ExecCommandsV1Context{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// ExecCommandsV1Payload is the commands_v1 exec action payload.
type ExecCommandsV1Payload []*CommandParam

// OK sends a HTTP response with status code 200.
func (ctx *ExecCommandsV1Context) OK(r *ArduinoAgentExec) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.arduino.agent.exec+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ListCommandsV1Context provides the commands_v1 list action context.
type ListCommandsV1Context struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCommandsV1Context parses the incoming request URL and body, performs validations and creates the
// context used by the commands_v1 controller list action.
func NewListCommandsV1Context(ctx context.Context, r *http.Request, service *goa.Service) (*ListCommandsV1Context, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCommandsV1Context{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCommandsV1Context) OK(r ArduinoAgentCommandCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.arduino.agent.command+json; type=collection")
	if r == nil {
		r = ArduinoAgentCommandCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// WebsocketConnectV1Context provides the connect_v1 websocket action context.
type WebsocketConnectV1Context struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Baud int
	Port string
}

// NewWebsocketConnectV1Context parses the incoming request URL and body, performs validations and creates the
// context used by the connect_v1 controller websocket action.
func NewWebsocketConnectV1Context(ctx context.Context, r *http.Request, service *goa.Service) (*WebsocketConnectV1Context, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := WebsocketConnectV1Context{Context: ctx, ResponseData: resp, RequestData: req}
	paramBaud := req.Params["baud"]
	if len(paramBaud) == 0 {
		rctx.Baud = 9600
	} else {
		rawBaud := paramBaud[0]
		if baud, err2 := strconv.Atoi(rawBaud); err2 == nil {
			rctx.Baud = baud
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("baud", rawBaud, "integer"))
		}
	}
	paramPort := req.Params["port"]
	if len(paramPort) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("port"))
	} else {
		rawPort := paramPort[0]
		rctx.Port = rawPort
	}
	return &rctx, err
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *WebsocketConnectV1Context) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// ListDiscoverV1Context provides the discover_v1 list action context.
type ListDiscoverV1Context struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListDiscoverV1Context parses the incoming request URL and body, performs validations and creates the
// context used by the discover_v1 controller list action.
func NewListDiscoverV1Context(ctx context.Context, r *http.Request, service *goa.Service) (*ListDiscoverV1Context, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListDiscoverV1Context{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListDiscoverV1Context) OK(r *ArduinoAgentDiscover) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.arduino.agent.discover+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
