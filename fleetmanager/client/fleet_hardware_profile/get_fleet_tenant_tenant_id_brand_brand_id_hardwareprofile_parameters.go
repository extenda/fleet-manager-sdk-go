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

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParams creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams object
// with the default values initialized.
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParams() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithTimeout creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams{

		timeout: timeout,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithContext creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams{

		Context: ctx,
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithHTTPClient creates a new GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileParamsWithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	var ()
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams{
		HTTPClient: client,
	}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams contains all the parameters to send to the API endpoint
for the get fleet tenant tenant ID brand brand ID hardwareprofile operation typically these are written to a http.Request
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams struct {

	/*BrandID*/
	BrandID string
	/*NextToken*/
	NextToken *string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithTimeout(timeout time.Duration) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithContext(ctx context.Context) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithHTTPClient(client *http.Client) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithBrandID(brandID string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithNextToken adds the nextToken to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithNextToken(nextToken *string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithTenantID adds the tenantID to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WithTenantID(tenantID string) *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get fleet tenant tenant ID brand brand ID hardwareprofile params
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param brandId
	if err := r.SetPathParam("brandId", o.BrandID); err != nil {
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