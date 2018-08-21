package main

import (
	"github.com/goadesign/goa"
	"github.com/testGoa/app"
)

// ActionController implements the action resource.
type ActionController struct {
	*goa.Controller
}

// NewActionController creates a action controller.
func NewActionController(service *goa.Service) *ActionController {
	return &ActionController{Controller: service.NewController("ActionController")}
}

// Ping runs the ping action.
func (c *ActionController) Ping(ctx *app.PingActionContext) error {
	// ActionController_Ping: start_implement

	// Put your logic here

	res := &app.Hello{}
	return ctx.OK(res)
	// ActionController_Ping: end_implement
}
