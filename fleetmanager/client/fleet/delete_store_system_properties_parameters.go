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

// NewDeleteStoreSystemPropertiesParams creates a new DeleteStoreSystemPropertiesParams object
// with the default values initialized.
func NewDeleteStoreSystemPropertiesParams() *DeleteStoreSystemPropertiesParams {
	var ()
	return &DeleteStoreSystemPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStoreSystemPropertiesParamsWithTimeout creates a new DeleteStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStoreSystemPropertiesParamsWithTimeout(timeout time.Duration) *DeleteStoreSystemPropertiesParams {
	var ()
	return &DeleteStoreSystemPropertiesParams{

		timeout: timeout,
	}
}

// NewDeleteStoreSystemPropertiesParamsWithContext creates a new DeleteStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStoreSystemPropertiesParamsWithContext(ctx context.Context) *DeleteStoreSystemPropertiesParams {
	var ()
	return &DeleteStoreSystemPropertiesParams{

		Context: ctx,
	}
}

// NewDeleteStoreSystemPropertiesParamsWithHTTPClient creates a new DeleteStoreSystemPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStoreSystemPropertiesParamsWithHTTPClient(client *http.Client) *DeleteStoreSystemPropertiesParams {
	var ()
	return &DeleteStoreSystemPropertiesParams{
		HTTPClient: client,
	}
}

/*DeleteStoreSystemPropertiesParams contains all the parameters to send to the API endpoint
for the delete store system properties operation typically these are written to a http.Request
*/
type DeleteStoreSystemPropertiesParams struct {

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

// WithTimeout adds the timeout to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithTimeout(timeout time.Duration) *DeleteStoreSystemPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithContext(ctx context.Context) *DeleteStoreSystemPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithHTTPClient(client *http.Client) *DeleteStoreSystemPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithBrandID(brandID string) *DeleteStoreSystemPropertiesParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithCountryID(countryID string) *DeleteStoreSystemPropertiesParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithStoreID(storeID string) *DeleteStoreSystemPropertiesParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) WithTenantID(tenantID string) *DeleteStoreSystemPropertiesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete store system properties params
func (o *DeleteStoreSystemPropertiesParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStoreSystemPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
