// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigRouteDestinationIPNetIPAddressMaskHandlerFunc turns a function with the right signature into a delete config route destination IP net IP address mask handler
type DeleteConfigRouteDestinationIPNetIPAddressMaskHandlerFunc func(DeleteConfigRouteDestinationIPNetIPAddressMaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigRouteDestinationIPNetIPAddressMaskHandlerFunc) Handle(params DeleteConfigRouteDestinationIPNetIPAddressMaskParams) middleware.Responder {
	return fn(params)
}

// DeleteConfigRouteDestinationIPNetIPAddressMaskHandler interface for that can handle valid delete config route destination IP net IP address mask params
type DeleteConfigRouteDestinationIPNetIPAddressMaskHandler interface {
	Handle(DeleteConfigRouteDestinationIPNetIPAddressMaskParams) middleware.Responder
}

// NewDeleteConfigRouteDestinationIPNetIPAddressMask creates a new http.Handler for the delete config route destination IP net IP address mask operation
func NewDeleteConfigRouteDestinationIPNetIPAddressMask(ctx *middleware.Context, handler DeleteConfigRouteDestinationIPNetIPAddressMaskHandler) *DeleteConfigRouteDestinationIPNetIPAddressMask {
	return &DeleteConfigRouteDestinationIPNetIPAddressMask{Context: ctx, Handler: handler}
}

/*
	DeleteConfigRouteDestinationIPNetIPAddressMask swagger:route DELETE /config/route/destinationIPNet/{ip_address}/{mask} deleteConfigRouteDestinationIpNetIpAddressMask

# Create a new Load balancer service

Create a new load balancer service with .
*/
type DeleteConfigRouteDestinationIPNetIPAddressMask struct {
	Context *middleware.Context
	Handler DeleteConfigRouteDestinationIPNetIPAddressMaskHandler
}

func (o *DeleteConfigRouteDestinationIPNetIPAddressMask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigRouteDestinationIPNetIPAddressMaskParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
