package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetKeysKeyHandlerFunc turns a function with the right signature into a get keys key handler
type GetKeysKeyHandlerFunc func(GetKeysKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetKeysKeyHandlerFunc) Handle(params GetKeysKeyParams) middleware.Responder {
	return fn(params)
}

// GetKeysKeyHandler interface for that can handle valid get keys key params
type GetKeysKeyHandler interface {
	Handle(GetKeysKeyParams) middleware.Responder
}

// NewGetKeysKey creates a new http.Handler for the get keys key operation
func NewGetKeysKey(ctx *middleware.Context, handler GetKeysKeyHandler) *GetKeysKey {
	return &GetKeysKey{Context: ctx, Handler: handler}
}

/*GetKeysKey swagger:route GET /keys/{key} getKeysKey

GetKeysKey get keys key API

*/
type GetKeysKey struct {
	Context *middleware.Context
	Handler GetKeysKeyHandler
}

func (o *GetKeysKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetKeysKeyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
