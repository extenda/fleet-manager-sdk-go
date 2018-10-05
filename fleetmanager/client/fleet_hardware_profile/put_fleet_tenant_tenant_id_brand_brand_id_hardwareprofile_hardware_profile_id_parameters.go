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

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams creates a new PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams object
// with the default values initialized.
func NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams() *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithTimeout creates a new PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithTimeout(timeout time.Duration) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams{

		timeout: timeout,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithContext creates a new PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithContext(ctx context.Context) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams{

		Context: ctx,
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithHTTPClient creates a new PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParamsWithHTTPClient(client *http.Client) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	var ()
	return &PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams{
		HTTPClient: client,
	}
}

/*PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams contains all the parameters to send to the API endpoint
for the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID operation typically these are written to a http.Request
*/
type PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams struct {

	/*Body*/
	Body *models.UpdateFleetHardwareProfile
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

// WithTimeout adds the timeout to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithTimeout(timeout time.Duration) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithContext(ctx context.Context) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithHTTPClient(client *http.Client) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithBody(body *models.UpdateFleetHardwareProfile) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetBody(body *models.UpdateFleetHardwareProfile) {
	o.Body = body
}

// WithBrandID adds the brandID to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithBrandID(brandID string) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithHardwareProfileID adds the hardwareProfileID to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithHardwareProfileID(hardwareProfileID string) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetHardwareProfileID(hardwareProfileID)
	return o
}

// SetHardwareProfileID adds the hardwareProfileId to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetHardwareProfileID(hardwareProfileID string) {
	o.HardwareProfileID = hardwareProfileID
}

// WithTenantID adds the tenantID to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WithTenantID(tenantID string) *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the put fleet tenant tenant ID brand brand ID hardwareprofile hardware profile ID params
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *PutFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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