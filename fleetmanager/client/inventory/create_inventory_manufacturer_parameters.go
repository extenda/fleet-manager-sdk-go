// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewCreateInventoryManufacturerParams creates a new CreateInventoryManufacturerParams object
// with the default values initialized.
func NewCreateInventoryManufacturerParams() *CreateInventoryManufacturerParams {
	var ()
	return &CreateInventoryManufacturerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInventoryManufacturerParamsWithTimeout creates a new CreateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInventoryManufacturerParamsWithTimeout(timeout time.Duration) *CreateInventoryManufacturerParams {
	var ()
	return &CreateInventoryManufacturerParams{

		timeout: timeout,
	}
}

// NewCreateInventoryManufacturerParamsWithContext creates a new CreateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInventoryManufacturerParamsWithContext(ctx context.Context) *CreateInventoryManufacturerParams {
	var ()
	return &CreateInventoryManufacturerParams{

		Context: ctx,
	}
}

// NewCreateInventoryManufacturerParamsWithHTTPClient creates a new CreateInventoryManufacturerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInventoryManufacturerParamsWithHTTPClient(client *http.Client) *CreateInventoryManufacturerParams {
	var ()
	return &CreateInventoryManufacturerParams{
		HTTPClient: client,
	}
}

/*CreateInventoryManufacturerParams contains all the parameters to send to the API endpoint
for the create inventory manufacturer operation typically these are written to a http.Request
*/
type CreateInventoryManufacturerParams struct {

	/*Body*/
	Body *models.CreateInventoryManufacturer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) WithTimeout(timeout time.Duration) *CreateInventoryManufacturerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) WithContext(ctx context.Context) *CreateInventoryManufacturerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) WithHTTPClient(client *http.Client) *CreateInventoryManufacturerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) WithBody(body *models.CreateInventoryManufacturer) *CreateInventoryManufacturerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create inventory manufacturer params
func (o *CreateInventoryManufacturerParams) SetBody(body *models.CreateInventoryManufacturer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInventoryManufacturerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
