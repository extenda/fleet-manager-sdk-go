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

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID country country ID store store ID operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*StoreID*/
	StoreID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithCountryID(countryID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithStoreID(storeID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID country country ID store store ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param storeId
	if err := r.SetPathParam("storeId", o.StoreID); err != nil {
		return err
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