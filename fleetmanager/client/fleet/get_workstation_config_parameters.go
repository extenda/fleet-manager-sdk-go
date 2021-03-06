// Code generated by go-swagger; DO NOT EDIT.

package fleet

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

// NewGetWorkstationConfigParams creates a new GetWorkstationConfigParams object
// with the default values initialized.
func NewGetWorkstationConfigParams() *GetWorkstationConfigParams {
	var ()
	return &GetWorkstationConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkstationConfigParamsWithTimeout creates a new GetWorkstationConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWorkstationConfigParamsWithTimeout(timeout time.Duration) *GetWorkstationConfigParams {
	var ()
	return &GetWorkstationConfigParams{

		timeout: timeout,
	}
}

// NewGetWorkstationConfigParamsWithContext creates a new GetWorkstationConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWorkstationConfigParamsWithContext(ctx context.Context) *GetWorkstationConfigParams {
	var ()
	return &GetWorkstationConfigParams{

		Context: ctx,
	}
}

// NewGetWorkstationConfigParamsWithHTTPClient creates a new GetWorkstationConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWorkstationConfigParamsWithHTTPClient(client *http.Client) *GetWorkstationConfigParams {
	var ()
	return &GetWorkstationConfigParams{
		HTTPClient: client,
	}
}

/*GetWorkstationConfigParams contains all the parameters to send to the API endpoint
for the get workstation config operation typically these are written to a http.Request
*/
type GetWorkstationConfigParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*Platform*/
	Platform string
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

// WithTimeout adds the timeout to the get workstation config params
func (o *GetWorkstationConfigParams) WithTimeout(timeout time.Duration) *GetWorkstationConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workstation config params
func (o *GetWorkstationConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workstation config params
func (o *GetWorkstationConfigParams) WithContext(ctx context.Context) *GetWorkstationConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workstation config params
func (o *GetWorkstationConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workstation config params
func (o *GetWorkstationConfigParams) WithHTTPClient(client *http.Client) *GetWorkstationConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workstation config params
func (o *GetWorkstationConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the get workstation config params
func (o *GetWorkstationConfigParams) WithBrandID(brandID string) *GetWorkstationConfigParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the get workstation config params
func (o *GetWorkstationConfigParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the get workstation config params
func (o *GetWorkstationConfigParams) WithCountryID(countryID string) *GetWorkstationConfigParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the get workstation config params
func (o *GetWorkstationConfigParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithPlatform adds the platform to the get workstation config params
func (o *GetWorkstationConfigParams) WithPlatform(platform string) *GetWorkstationConfigParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the get workstation config params
func (o *GetWorkstationConfigParams) SetPlatform(platform string) {
	o.Platform = platform
}

// WithStoreID adds the storeID to the get workstation config params
func (o *GetWorkstationConfigParams) WithStoreID(storeID string) *GetWorkstationConfigParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the get workstation config params
func (o *GetWorkstationConfigParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the get workstation config params
func (o *GetWorkstationConfigParams) WithTenantID(tenantID string) *GetWorkstationConfigParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get workstation config params
func (o *GetWorkstationConfigParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the get workstation config params
func (o *GetWorkstationConfigParams) WithWorkstationID(workstationID string) *GetWorkstationConfigParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the get workstation config params
func (o *GetWorkstationConfigParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkstationConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param platform
	qrPlatform := o.Platform
	qPlatform := qrPlatform
	if qPlatform != "" {
		if err := r.SetQueryParam("platform", qPlatform); err != nil {
			return err
		}
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
