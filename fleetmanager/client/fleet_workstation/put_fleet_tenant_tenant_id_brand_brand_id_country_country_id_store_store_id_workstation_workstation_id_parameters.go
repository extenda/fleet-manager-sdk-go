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

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams creates a new PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized.
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams() *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithTimeout creates a new PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithTimeout(timeout time.Duration) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		timeout: timeout,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithContext creates a new PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithContext(ctx context.Context) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{

		Context: ctx,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithHTTPClient creates a new PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParamsWithHTTPClient(client *http.Client) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams{
		HTTPClient: client,
	}
}

/*PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams contains all the parameters to send to the API endpoint
for the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID operation typically these are written to a http.Request
*/
type PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams struct {

	/*Body*/
	Body *models.UpdateFleetWorkstation
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

// WithTimeout adds the timeout to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithTimeout(timeout time.Duration) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithContext(ctx context.Context) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithHTTPClient(client *http.Client) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithBody(body *models.UpdateFleetWorkstation) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetBody(body *models.UpdateFleetWorkstation) {
	o.Body = body
}

// WithBrandID adds the brandID to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithBrandID(brandID string) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithCountryID(countryID string) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithStoreID(storeID string) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithTenantID(tenantID string) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WithWorkstationID(workstationID string) *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the put fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID params
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
