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

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// NewPutInventoryManufacturerManufacturerIDParams creates a new PutInventoryManufacturerManufacturerIDParams object
// with the default values initialized.
func NewPutInventoryManufacturerManufacturerIDParams() *PutInventoryManufacturerManufacturerIDParams {
	var ()
	return &PutInventoryManufacturerManufacturerIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutInventoryManufacturerManufacturerIDParamsWithTimeout creates a new PutInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutInventoryManufacturerManufacturerIDParamsWithTimeout(timeout time.Duration) *PutInventoryManufacturerManufacturerIDParams {
	var ()
	return &PutInventoryManufacturerManufacturerIDParams{

		timeout: timeout,
	}
}

// NewPutInventoryManufacturerManufacturerIDParamsWithContext creates a new PutInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutInventoryManufacturerManufacturerIDParamsWithContext(ctx context.Context) *PutInventoryManufacturerManufacturerIDParams {
	var ()
	return &PutInventoryManufacturerManufacturerIDParams{

		Context: ctx,
	}
}

// NewPutInventoryManufacturerManufacturerIDParamsWithHTTPClient creates a new PutInventoryManufacturerManufacturerIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutInventoryManufacturerManufacturerIDParamsWithHTTPClient(client *http.Client) *PutInventoryManufacturerManufacturerIDParams {
	var ()
	return &PutInventoryManufacturerManufacturerIDParams{
		HTTPClient: client,
	}
}

/*PutInventoryManufacturerManufacturerIDParams contains all the parameters to send to the API endpoint
for the put inventory manufacturer manufacturer ID operation typically these are written to a http.Request
*/
type PutInventoryManufacturerManufacturerIDParams struct {

	/*Body*/
	Body *models.UpdateInventoryManufacturer
	/*ManufacturerID*/
	ManufacturerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) WithTimeout(timeout time.Duration) *PutInventoryManufacturerManufacturerIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) WithContext(ctx context.Context) *PutInventoryManufacturerManufacturerIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) WithHTTPClient(client *http.Client) *PutInventoryManufacturerManufacturerIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) WithBody(body *models.UpdateInventoryManufacturer) *PutInventoryManufacturerManufacturerIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) SetBody(body *models.UpdateInventoryManufacturer) {
	o.Body = body
}

// WithManufacturerID adds the manufacturerID to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) WithManufacturerID(manufacturerID string) *PutInventoryManufacturerManufacturerIDParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the put inventory manufacturer manufacturer ID params
func (o *PutInventoryManufacturerManufacturerIDParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WriteToRequest writes these params to a swagger request
func (o *PutInventoryManufacturerManufacturerIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param manufacturerId
	if err := r.SetPathParam("manufacturerId", o.ManufacturerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
