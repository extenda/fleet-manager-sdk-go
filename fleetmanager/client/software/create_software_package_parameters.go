// Code generated by go-swagger; DO NOT EDIT.

package software

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

// NewCreateSoftwarePackageParams creates a new CreateSoftwarePackageParams object
// with the default values initialized.
func NewCreateSoftwarePackageParams() *CreateSoftwarePackageParams {
	var ()
	return &CreateSoftwarePackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSoftwarePackageParamsWithTimeout creates a new CreateSoftwarePackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSoftwarePackageParamsWithTimeout(timeout time.Duration) *CreateSoftwarePackageParams {
	var ()
	return &CreateSoftwarePackageParams{

		timeout: timeout,
	}
}

// NewCreateSoftwarePackageParamsWithContext creates a new CreateSoftwarePackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSoftwarePackageParamsWithContext(ctx context.Context) *CreateSoftwarePackageParams {
	var ()
	return &CreateSoftwarePackageParams{

		Context: ctx,
	}
}

// NewCreateSoftwarePackageParamsWithHTTPClient creates a new CreateSoftwarePackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSoftwarePackageParamsWithHTTPClient(client *http.Client) *CreateSoftwarePackageParams {
	var ()
	return &CreateSoftwarePackageParams{
		HTTPClient: client,
	}
}

/*CreateSoftwarePackageParams contains all the parameters to send to the API endpoint
for the create software package operation typically these are written to a http.Request
*/
type CreateSoftwarePackageParams struct {

	/*Body*/
	Body *models.CreateSoftwarePackage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create software package params
func (o *CreateSoftwarePackageParams) WithTimeout(timeout time.Duration) *CreateSoftwarePackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create software package params
func (o *CreateSoftwarePackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create software package params
func (o *CreateSoftwarePackageParams) WithContext(ctx context.Context) *CreateSoftwarePackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create software package params
func (o *CreateSoftwarePackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create software package params
func (o *CreateSoftwarePackageParams) WithHTTPClient(client *http.Client) *CreateSoftwarePackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create software package params
func (o *CreateSoftwarePackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create software package params
func (o *CreateSoftwarePackageParams) WithBody(body *models.CreateSoftwarePackage) *CreateSoftwarePackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create software package params
func (o *CreateSoftwarePackageParams) SetBody(body *models.CreateSoftwarePackage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSoftwarePackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
