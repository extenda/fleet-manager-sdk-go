// Code generated by go-swagger; DO NOT EDIT.

package driver_package

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

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewPutDriverPackagePackageIDParams creates a new PutDriverPackagePackageIDParams object
// with the default values initialized.
func NewPutDriverPackagePackageIDParams() *PutDriverPackagePackageIDParams {
	var ()
	return &PutDriverPackagePackageIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDriverPackagePackageIDParamsWithTimeout creates a new PutDriverPackagePackageIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDriverPackagePackageIDParamsWithTimeout(timeout time.Duration) *PutDriverPackagePackageIDParams {
	var ()
	return &PutDriverPackagePackageIDParams{

		timeout: timeout,
	}
}

// NewPutDriverPackagePackageIDParamsWithContext creates a new PutDriverPackagePackageIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDriverPackagePackageIDParamsWithContext(ctx context.Context) *PutDriverPackagePackageIDParams {
	var ()
	return &PutDriverPackagePackageIDParams{

		Context: ctx,
	}
}

// NewPutDriverPackagePackageIDParamsWithHTTPClient creates a new PutDriverPackagePackageIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDriverPackagePackageIDParamsWithHTTPClient(client *http.Client) *PutDriverPackagePackageIDParams {
	var ()
	return &PutDriverPackagePackageIDParams{
		HTTPClient: client,
	}
}

/*PutDriverPackagePackageIDParams contains all the parameters to send to the API endpoint
for the put driver package package ID operation typically these are written to a http.Request
*/
type PutDriverPackagePackageIDParams struct {

	/*Body*/
	Body *models.UpdateDriverPackage
	/*PackageID*/
	PackageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) WithTimeout(timeout time.Duration) *PutDriverPackagePackageIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) WithContext(ctx context.Context) *PutDriverPackagePackageIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) WithHTTPClient(client *http.Client) *PutDriverPackagePackageIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) WithBody(body *models.UpdateDriverPackage) *PutDriverPackagePackageIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) SetBody(body *models.UpdateDriverPackage) {
	o.Body = body
}

// WithPackageID adds the packageID to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) WithPackageID(packageID string) *PutDriverPackagePackageIDParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the put driver package package ID params
func (o *PutDriverPackagePackageIDParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *PutDriverPackagePackageIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param packageId
	if err := r.SetPathParam("packageId", o.PackageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
