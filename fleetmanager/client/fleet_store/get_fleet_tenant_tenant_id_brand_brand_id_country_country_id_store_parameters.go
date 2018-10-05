// Code generated by go-swagger; DO NOT EDIT.

package fleet_store

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

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID country country ID store operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*NextToken*/
	NextToken *string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithCountryID(countryID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithNextToken adds the nextToken to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithNextToken(nextToken *string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID country country ID store params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
		return err
	}

	// path param countryId
	if err := r.SetPathParam("countryId", o.CountryID); err != nil {
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

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
