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
)

// NewGetInventoryManufacturerManufacturerIDModelParams creates a new GetInventoryManufacturerManufacturerIDModelParams object
// with the default values initialized.
func NewGetInventoryManufacturerManufacturerIDModelParams() *GetInventoryManufacturerManufacturerIDModelParams {
	var ()
	return &GetInventoryManufacturerManufacturerIDModelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInventoryManufacturerManufacturerIDModelParamsWithTimeout creates a new GetInventoryManufacturerManufacturerIDModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInventoryManufacturerManufacturerIDModelParamsWithTimeout(timeout time.Duration) *GetInventoryManufacturerManufacturerIDModelParams {
	var ()
	return &GetInventoryManufacturerManufacturerIDModelParams{

		timeout: timeout,
	}
}

// NewGetInventoryManufacturerManufacturerIDModelParamsWithContext creates a new GetInventoryManufacturerManufacturerIDModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInventoryManufacturerManufacturerIDModelParamsWithContext(ctx context.Context) *GetInventoryManufacturerManufacturerIDModelParams {
	var ()
	return &GetInventoryManufacturerManufacturerIDModelParams{

		Context: ctx,
	}
}

// NewGetInventoryManufacturerManufacturerIDModelParamsWithHTTPClient creates a new GetInventoryManufacturerManufacturerIDModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInventoryManufacturerManufacturerIDModelParamsWithHTTPClient(client *http.Client) *GetInventoryManufacturerManufacturerIDModelParams {
	var ()
	return &GetInventoryManufacturerManufacturerIDModelParams{
		HTTPClient: client,
	}
}

/*GetInventoryManufacturerManufacturerIDModelParams contains all the parameters to send to the API endpoint
for the get inventory manufacturer manufacturer ID model operation typically these are written to a http.Request
*/
type GetInventoryManufacturerManufacturerIDModelParams struct {

	/*ManufacturerID*/
	ManufacturerID string
	/*NextToken*/
	NextToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) WithTimeout(timeout time.Duration) *GetInventoryManufacturerManufacturerIDModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) WithContext(ctx context.Context) *GetInventoryManufacturerManufacturerIDModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) WithHTTPClient(client *http.Client) *GetInventoryManufacturerManufacturerIDModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithManufacturerID adds the manufacturerID to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) WithManufacturerID(manufacturerID string) *GetInventoryManufacturerManufacturerIDModelParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) SetManufacturerID(manufacturerID string) {
	o.ManufacturerID = manufacturerID
}

// WithNextToken adds the nextToken to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) WithNextToken(nextToken *string) *GetInventoryManufacturerManufacturerIDModelParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get inventory manufacturer manufacturer ID model params
func (o *GetInventoryManufacturerManufacturerIDModelParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WriteToRequest writes these params to a swagger request
func (o *GetInventoryManufacturerManufacturerIDModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param manufacturerId
	if err := r.SetPathParam("manufacturerId", o.ManufacturerID); err != nil {
		return err
	}

	if o.NextToken != nil {

		// query param nextToken
		var qrNextToken string
		if o.NextToken != nil {
			qrNextToken = *o.NextToken
		}
		qNextToken := qrNextToken
		if qNextToken != "" {
			if err := r.SetQueryParam("nextToken", qNextToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
