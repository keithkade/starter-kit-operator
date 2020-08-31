// Code generated by go-swagger; DO NOT EDIT.

package starter_kit_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams creates a new GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams object
// no default values defined in spec.
func NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams() GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams {

	return GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams{}
}

// GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams contains all the bound params for the get apis devx ibm com v1alpha1 namespaces namespace starterkits operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkits
type GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Namespace where the starter kit is located.
	  Required: true
	  In: path
	*/
	Namespace string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams() beforehand.
func (o *GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}
