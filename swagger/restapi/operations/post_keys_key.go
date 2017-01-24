package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostKeysKeyHandlerFunc turns a function with the right signature into a post keys key handler
type PostKeysKeyHandlerFunc func(PostKeysKeyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostKeysKeyHandlerFunc) Handle(params PostKeysKeyParams) middleware.Responder {
	return fn(params)
}

// PostKeysKeyHandler interface for that can handle valid post keys key params
type PostKeysKeyHandler interface {
	Handle(PostKeysKeyParams) middleware.Responder
}

// NewPostKeysKey creates a new http.Handler for the post keys key operation
func NewPostKeysKey(ctx *middleware.Context, handler PostKeysKeyHandler) *PostKeysKey {
	return &PostKeysKey{Context: ctx, Handler: handler}
}

/*PostKeysKey swagger:route POST /keys/{key} postKeysKey

PostKeysKey post keys key API

*/
type PostKeysKey struct {
	Context *middleware.Context
	Handler PostKeysKeyHandler
}

func (o *PostKeysKey) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostKeysKeyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}