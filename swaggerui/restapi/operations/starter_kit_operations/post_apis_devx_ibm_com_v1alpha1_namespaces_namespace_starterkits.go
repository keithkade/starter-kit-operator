// Code generated by go-swagger; DO NOT EDIT.

package starter_kit_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandlerFunc turns a function with the right signature into a post apis devx ibm com v1alpha1 namespaces namespace starterkits handler
type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandlerFunc func(PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandlerFunc) Handle(params PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandler interface for that can handle valid post apis devx ibm com v1alpha1 namespaces namespace starterkits params
type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandler interface {
	Handle(PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams, interface{}) middleware.Responder
}

// NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits creates a new http.Handler for the post apis devx ibm com v1alpha1 namespaces namespace starterkits operation
func NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits(ctx *middleware.Context, handler PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandler) *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits {
	return &PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits{Context: ctx, Handler: handler}
}

/*PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits swagger:route POST /apis/devx.ibm.com/v1alpha1/namespaces/{namespace}/starterkits Starter Kit Operations postApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits

Creates a starter kit.

*/
type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits struct {
	Context *middleware.Context
	Handler PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsHandler
}

func (o *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
