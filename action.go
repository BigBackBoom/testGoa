package main

import (
	"github.com/bigbackboom/testGoa/app"
	"github.com/goadesign/goa"
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
	message := "pong"

	res := &app.Hello{}
	res.ID = 0
	res.Message = message
	return ctx.OK(res)
	// ActionController_Ping: end_implement
}
