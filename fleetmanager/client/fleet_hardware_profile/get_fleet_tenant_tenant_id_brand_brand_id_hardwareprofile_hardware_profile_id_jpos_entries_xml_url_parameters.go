// Code generated by go-swagger; DO NOT EDIT.

package fleet_hardware_profile

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

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams struct {

	/*BrandID*/
	BrandID string
	/*HardwareProfileID*/
	HardwareProfileID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithHardwareProfileID adds the hardwareProfileID to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithHardwareProfileID(hardwareProfileID string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetHardwareProfileID(hardwareProfileID)
	return o
}

// SetHardwareProfileID adds the hardwareProfileId to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetHardwareProfileID(hardwareProfileID string) {
	o.HardwareProfileID = hardwareProfileID
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID jpos entries XML URL params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
		return err
	}

	// path param hardwareProfileId
	if err := r.SetPathParam("hardwareProfileId", o.HardwareProfileID); err != nil {
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
