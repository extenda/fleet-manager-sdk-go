// Code generated by go-swagger; DO NOT EDIT.

package fleet_software_profile

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

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams creates a new GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams() *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams struct {

	/*BrandID*/
	BrandID string
	/*SoftwareProfileID*/
	SoftwareProfileID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithSoftwareProfileID adds the softwareProfileID to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithSoftwareProfileID(softwareProfileID string) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetSoftwareProfileID(softwareProfileID)
	return o
}

// SetSoftwareProfileID adds the softwareProfileId to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetSoftwareProfileID(softwareProfileID string) {
	o.SoftwareProfileID = softwareProfileID
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID softwareprofile software profile ID params
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
		return err
	}

	// path param softwareProfileId
	if err := r.SetPathParam("softwareProfileId", o.SoftwareProfileID); err != nil {
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
