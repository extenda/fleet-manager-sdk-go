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

// NewUploadStoreSystemPropertiesParams creates a new UploadStoreSystemPropertiesParams object
// with the default values initialized.
func NewUploadStoreSystemPropertiesParams() *UploadStoreSystemPropertiesParams {
	var ()
	return &UploadStoreSystemPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadStoreSystemPropertiesParamsWithTimeout creates a new UploadStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadStoreSystemPropertiesParamsWithTimeout(timeout time.Duration) *UploadStoreSystemPropertiesParams {
	var ()
	return &UploadStoreSystemPropertiesParams{

		timeout: timeout,
	}
}

// NewUploadStoreSystemPropertiesParamsWithContext creates a new UploadStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadStoreSystemPropertiesParamsWithContext(ctx context.Context) *UploadStoreSystemPropertiesParams {
	var ()
	return &UploadStoreSystemPropertiesParams{

		Context: ctx,
	}
}

// NewUploadStoreSystemPropertiesParamsWithHTTPClient creates a new UploadStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadStoreSystemPropertiesParamsWithHTTPClient(client *http.Client) *UploadStoreSystemPropertiesParams {
	var ()
	return &UploadStoreSystemPropertiesParams{
		HTTPClient: client,
	}
}

/*UploadStoreSystemPropertiesParams contains all the parameters to send to the API endpoint
for the upload store system properties operation typically these are written to a http.Request
*/
type UploadStoreSystemPropertiesParams struct {

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

// WithTimeout adds the timeout to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithTimeout(timeout time.Duration) *UploadStoreSystemPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithContext(ctx context.Context) *UploadStoreSystemPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithHTTPClient(client *http.Client) *UploadStoreSystemPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithBrandID(brandID string) *UploadStoreSystemPropertiesParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithCountryID(countryID string) *UploadStoreSystemPropertiesParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithStoreID(storeID string) *UploadStoreSystemPropertiesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) WithTenantID(tenantID string) *UploadStoreSystemPropertiesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the upload store system properties params
func (o *UploadStoreSystemPropertiesParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UploadStoreSystemPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
