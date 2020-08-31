// Code generated by go-swagger; DO NOT EDIT.

package starter_kit_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOKCode is the HTTP code returned for type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK
const PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOKCode int = 200

/*PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK OK

swagger:response postApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK
*/
type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK struct {
}

// NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK creates PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK with default headers values
func NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK() *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK {

	return &PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK{}
}

// WriteResponse to the client
func (o *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbiddenCode is the HTTP code returned for type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden
const PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbiddenCode int = 403

/*PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden Unauthorized

swagger:response postApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden
*/
type PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden struct {
}

// NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden creates PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden with default headers values
func NewPostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden() *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden {

	return &PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden{}
}

// WriteResponse to the client
func (o *PostApisDevxIbmComV1alpha1NamespacesNamespaceStarterkitsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
