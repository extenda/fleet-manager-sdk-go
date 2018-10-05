// Code generated by go-swagger; DO NOT EDIT.

package fleet_workstation

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

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*StoreID*/
	StoreID string
	/*TenantID*/
	TenantID string
	/*WorkstationID*/
	WorkstationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithCountryID(countryID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithStoreID(storeID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithWorkstationID(workstationID string) *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the get fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param workstationId
	if err := r.SetPathParam("workstationId", o.WorkstationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
