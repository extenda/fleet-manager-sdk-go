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

// NewDeleteWorkstationParams creates a new DeleteWorkstationParams object
// with the default values initialized.
func NewDeleteWorkstationParams() *DeleteWorkstationParams {
	var ()
	return &DeleteWorkstationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkstationParamsWithTimeout creates a new DeleteWorkstationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkstationParamsWithTimeout(timeout time.Duration) *DeleteWorkstationParams {
	var ()
	return &DeleteWorkstationParams{

		timeout: timeout,
	}
}

// NewDeleteWorkstationParamsWithContext creates a new DeleteWorkstationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkstationParamsWithContext(ctx context.Context) *DeleteWorkstationParams {
	var ()
	return &DeleteWorkstationParams{

		Context: ctx,
	}
}

// NewDeleteWorkstationParamsWithHTTPClient creates a new DeleteWorkstationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkstationParamsWithHTTPClient(client *http.Client) *DeleteWorkstationParams {
	var ()
	return &DeleteWorkstationParams{
		HTTPClient: client,
	}
}

/*DeleteWorkstationParams contains all the parameters to send to the API endpoint
for the delete workstation operation typically these are written to a http.Request
*/
type DeleteWorkstationParams struct {

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

// WithTimeout adds the timeout to the delete workstation params
func (o *DeleteWorkstationParams) WithTimeout(timeout time.Duration) *DeleteWorkstationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workstation params
func (o *DeleteWorkstationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workstation params
func (o *DeleteWorkstationParams) WithContext(ctx context.Context) *DeleteWorkstationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workstation params
func (o *DeleteWorkstationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workstation params
func (o *DeleteWorkstationParams) WithHTTPClient(client *http.Client) *DeleteWorkstationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workstation params
func (o *DeleteWorkstationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the delete workstation params
func (o *DeleteWorkstationParams) WithBrandID(brandID string) *DeleteWorkstationParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the delete workstation params
func (o *DeleteWorkstationParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the delete workstation params
func (o *DeleteWorkstationParams) WithCountryID(countryID string) *DeleteWorkstationParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the delete workstation params
func (o *DeleteWorkstationParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithStoreID adds the storeID to the delete workstation params
func (o *DeleteWorkstationParams) WithStoreID(storeID string) *DeleteWorkstationParams {
	o.SetStoreID(storeID)
	return o
}

// SetStoreID adds the storeId to the delete workstation params
func (o *DeleteWorkstationParams) SetStoreID(storeID string) {
	o.StoreID = storeID
}

// WithTenantID adds the tenantID to the delete workstation params
func (o *DeleteWorkstationParams) WithTenantID(tenantID string) *DeleteWorkstationParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete workstation params
func (o *DeleteWorkstationParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WithWorkstationID adds the workstationID to the delete workstation params
func (o *DeleteWorkstationParams) WithWorkstationID(workstationID string) *DeleteWorkstationParams {
	o.SetWorkstationID(workstationID)
	return o
}

// SetWorkstationID adds the workstationId to the delete workstation params
func (o *DeleteWorkstationParams) SetWorkstationID(workstationID string) {
	o.WorkstationID = workstationID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkstationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
