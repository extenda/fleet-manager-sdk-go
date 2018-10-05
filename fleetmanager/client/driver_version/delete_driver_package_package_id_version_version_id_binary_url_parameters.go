// Code generated by go-swagger; DO NOT EDIT.

package driver_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams creates a new DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams object
// with the default values initialized.
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams() *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	var ()
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithTimeout creates a new DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithTimeout(timeout time.Duration) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	var ()
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams{

		timeout: timeout,
	}
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithContext creates a new DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithContext(ctx context.Context) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	var ()
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams{

		Context: ctx,
	}
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithHTTPClient creates a new DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLParamsWithHTTPClient(client *http.Client) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	var ()
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams{
		HTTPClient: client,
	}
}

/*DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams contains all the parameters to send to the API endpoint
for the delete driver package package ID version version ID binary URL operation typically these are written to a http.Request
*/
type DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams struct {

	/*PackageID*/
	PackageID string
	/*VersionID*/
	VersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WithTimeout(timeout time.Duration) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WithContext(ctx context.Context) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WithHTTPClient(client *http.Client) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPackageID adds the packageID to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WithPackageID(packageID string) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WithVersionID adds the versionID to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WithVersionID(versionID string) *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete driver package package ID version version ID binary URL params
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) SetVersionID(versionID string) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}