// Code generated by go-swagger; DO NOT EDIT.

package driver

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

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// NewUpdateDriverPackageParams creates a new UpdateDriverPackageParams object
// with the default values initialized.
func NewUpdateDriverPackageParams() *UpdateDriverPackageParams {
	var ()
	return &UpdateDriverPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDriverPackageParamsWithTimeout creates a new UpdateDriverPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDriverPackageParamsWithTimeout(timeout time.Duration) *UpdateDriverPackageParams {
	var ()
	return &UpdateDriverPackageParams{

		timeout: timeout,
	}
}

// NewUpdateDriverPackageParamsWithContext creates a new UpdateDriverPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDriverPackageParamsWithContext(ctx context.Context) *UpdateDriverPackageParams {
	var ()
	return &UpdateDriverPackageParams{

		Context: ctx,
	}
}

// NewUpdateDriverPackageParamsWithHTTPClient creates a new UpdateDriverPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDriverPackageParamsWithHTTPClient(client *http.Client) *UpdateDriverPackageParams {
	var ()
	return &UpdateDriverPackageParams{
		HTTPClient: client,
	}
}

/*UpdateDriverPackageParams contains all the parameters to send to the API endpoint
for the update driver package operation typically these are written to a http.Request
*/
type UpdateDriverPackageParams struct {

	/*Body*/
	Body *models.UpdateDriverPackage
	/*PackageID*/
	PackageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update driver package params
func (o *UpdateDriverPackageParams) WithTimeout(timeout time.Duration) *UpdateDriverPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update driver package params
func (o *UpdateDriverPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update driver package params
func (o *UpdateDriverPackageParams) WithContext(ctx context.Context) *UpdateDriverPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update driver package params
func (o *UpdateDriverPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update driver package params
func (o *UpdateDriverPackageParams) WithHTTPClient(client *http.Client) *UpdateDriverPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update driver package params
func (o *UpdateDriverPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update driver package params
func (o *UpdateDriverPackageParams) WithBody(body *models.UpdateDriverPackage) *UpdateDriverPackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update driver package params
func (o *UpdateDriverPackageParams) SetBody(body *models.UpdateDriverPackage) {
	o.Body = body
}

// WithPackageID adds the packageID to the update driver package params
func (o *UpdateDriverPackageParams) WithPackageID(packageID string) *UpdateDriverPackageParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the update driver package params
func (o *UpdateDriverPackageParams) SetPackageID(packageID string) {
	o.PackageID = packageID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDriverPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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