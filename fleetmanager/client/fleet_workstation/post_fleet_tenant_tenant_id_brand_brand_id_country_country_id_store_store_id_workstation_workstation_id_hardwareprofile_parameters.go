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

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams creates a new PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams object
// with the default values initialized.
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams() *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	var ()
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithTimeout creates a new PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithTimeout(timeout time.Duration) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	var ()
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams{

		timeout: timeout,
	}
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithContext creates a new PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithContext(ctx context.Context) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	var ()
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams{

		Context: ctx,
	}
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithHTTPClient creates a new PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParamsWithHTTPClient(client *http.Client) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	var ()
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams{
		HTTPClient: client,
	}
}

/*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams contains all the parameters to send to the API endpoint
for the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile operation typically these are written to a http.Request
*/
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams struct {

	/*Body*/
	Body *models.FleetWorkstation2FleetHardwareProfile
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

// WithTimeout adds the timeout to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithTimeout(timeout time.Duration) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithContext(ctx context.Context) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithHTTPClient(client *http.Client) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithBody(body *models.FleetWorkstation2FleetHardwareProfile) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetBody(body *models.FleetWorkstation2FleetHardwareProfile) {
	o.Body = body
}

// WithBrandID adds the brandID to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithBrandID(brandID string) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithCountryID(countryID string) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithStoreID(storeID string) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithTenantID(tenantID string) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WithWorkstationID(workstationID string) *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the post fleet tenant tenant ID brand brand ID country country ID store store ID workstation workstation ID hardwareprofile params
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationWorkstationIDHardwareprofileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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