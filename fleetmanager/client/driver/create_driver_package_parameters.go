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

// NewCreateDriverPackageParams creates a new CreateDriverPackageParams object
// with the default values initialized.
func NewCreateDriverPackageParams() *CreateDriverPackageParams {
	var ()
	return &CreateDriverPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDriverPackageParamsWithTimeout creates a new CreateDriverPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateDriverPackageParamsWithTimeout(timeout time.Duration) *CreateDriverPackageParams {
	var ()
	return &CreateDriverPackageParams{

		timeout: timeout,
	}
}

// NewCreateDriverPackageParamsWithContext creates a new CreateDriverPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateDriverPackageParamsWithContext(ctx context.Context) *CreateDriverPackageParams {
	var ()
	return &CreateDriverPackageParams{

		Context: ctx,
	}
}

// NewCreateDriverPackageParamsWithHTTPClient creates a new CreateDriverPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateDriverPackageParamsWithHTTPClient(client *http.Client) *CreateDriverPackageParams {
	var ()
	return &CreateDriverPackageParams{
		HTTPClient: client,
	}
}

/*CreateDriverPackageParams contains all the parameters to send to the API endpoint
for the create driver package operation typically these are written to a http.Request
*/
type CreateDriverPackageParams struct {

	/*Body*/
	Body *models.CreateDriverPackage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create driver package params
func (o *CreateDriverPackageParams) WithTimeout(timeout time.Duration) *CreateDriverPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create driver package params
func (o *CreateDriverPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create driver package params
func (o *CreateDriverPackageParams) WithContext(ctx context.Context) *CreateDriverPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create driver package params
func (o *CreateDriverPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create driver package params
func (o *CreateDriverPackageParams) WithHTTPClient(client *http.Client) *CreateDriverPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create driver package params
func (o *CreateDriverPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create driver package params
func (o *CreateDriverPackageParams) WithBody(body *models.CreateDriverPackage) *CreateDriverPackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create driver package params
func (o *CreateDriverPackageParams) SetBody(body *models.CreateDriverPackage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDriverPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
