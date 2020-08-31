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

// NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams creates a new GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams object
// no default values defined in spec.
func NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams() GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams {

	return GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams{}
}

// GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams contains all the bound params for the get apis devx ibm com v1alpha1 namespaces namespace starterkits name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsName
type GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The name of the starter kit.
	  Required: true
	  In: path
	*/
	Name string
	/*Namespace where the starter kit is located.
	  Required: true
	  In: path
	*/
	Namespace string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams() beforehand.
func (o *GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindName binds and validates parameter Name from path.
func (o *GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *GetApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsNameParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}
