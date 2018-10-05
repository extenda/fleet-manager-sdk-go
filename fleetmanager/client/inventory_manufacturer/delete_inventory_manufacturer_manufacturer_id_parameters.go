// Code generated by go-swagger; DO NOT EDIT.

package inventory_manufacturer

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

// NewDeleteInventoryManufacturerManufacturerIDParams creates a new DeleteInventoryManufacturerManufacturerIDParams object
// with the default values initialized.
func NewDeleteInventoryManufacturerManufacturerIDParams() *DeleteInventoryManufacturerManufacturerIDParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDParamsWithTimeout creates a new DeleteInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInventoryManufacturerManufacturerIDParamsWithTimeout(timeout time.Duration) *DeleteInventoryManufacturerManufacturerIDParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDParams{

		timeout: timeout,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDParamsWithContext creates a new DeleteInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInventoryManufacturerManufacturerIDParamsWithContext(ctx context.Context) *DeleteInventoryManufacturerManufacturerIDParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDParams{

		Context: ctx,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDParamsWithHTTPClient creates a new DeleteInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInventoryManufacturerManufacturerIDParamsWithHTTPClient(client *http.Client) *DeleteInventoryManufacturerManufacturerIDParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDParams{
		HTTPClient: client,
	}
}

/*DeleteInventoryManufacturerManufacturerIDParams contains all the parameters to send to the API endpoint
for the delete inventory manufacturer manufacturer ID operation typically these are written to a http.Request
*/
type DeleteInventoryManufacturerManufacturerIDParams struct {

	/*ManufacturerID*/
	ManufacturerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) WithTimeout(timeout time.Duration) *DeleteInventoryManufacturerManufacturerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) WithContext(ctx context.Context) *DeleteInventoryManufacturerManufacturerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) WithHTTPClient(client *http.Client) *DeleteInventoryManufacturerManufacturerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManufacturerID adds the manufacturerID to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) WithManufacturerID(manufacturerID string) *DeleteInventoryManufacturerManufacturerIDParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the delete inventory manufacturer manufacturer ID params
func (o *DeleteInventoryManufacturerManufacturerIDParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInventoryManufacturerManufacturerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param manufacturerId
	if err := r.SetPathParam("manufacturerId", o.ManufacturerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}