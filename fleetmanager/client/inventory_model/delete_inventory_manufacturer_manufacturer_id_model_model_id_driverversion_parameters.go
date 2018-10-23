// Code generated by go-swagger; DO NOT EDIT.

package inventory_model

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

// NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams creates a new DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams object
// with the default values initialized.
func NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams() *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithTimeout creates a new DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithTimeout(timeout time.Duration) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams{

		timeout: timeout,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithContext creates a new DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithContext(ctx context.Context) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams{

		Context: ctx,
	}
}

// NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithHTTPClient creates a new DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParamsWithHTTPClient(client *http.Client) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	var ()
	return &DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams{
		HTTPClient: client,
	}
}

/*DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams contains all the parameters to send to the API endpoint
for the delete inventory manufacturer manufacturer ID model model ID driverversion operation typically these are written to a http.Request
*/
type DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams struct {

	/*Body*/
	Body *models.UnlinkInventoryModel2DriverVersion
	/*ManufacturerID*/
	ManufacturerID string
	/*ModelID*/
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithTimeout(timeout time.Duration) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithContext(ctx context.Context) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithHTTPClient(client *http.Client) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithBody(body *models.UnlinkInventoryModel2DriverVersion) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetBody(body *models.UnlinkInventoryModel2DriverVersion) {
	o.Body = body
}

// WithManufacturerID adds the manufacturerID to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithManufacturerID(manufacturerID string) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WithModelID adds the modelID to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WithModelID(modelID string) *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the delete inventory manufacturer manufacturer ID model model ID driverversion params
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInventoryManufacturerManufacturerIDModelModelIDDriverversionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param modelId
	if err := r.SetPathParam("modelId", o.ModelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
