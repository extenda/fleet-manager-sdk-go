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

// NewListStoresParams creates a new ListStoresParams object
// with the default values initialized.
func NewListStoresParams() *ListStoresParams {
	var ()
	return &ListStoresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListStoresParamsWithTimeout creates a new ListStoresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListStoresParamsWithTimeout(timeout time.Duration) *ListStoresParams {
	var ()
	return &ListStoresParams{

		timeout: timeout,
	}
}

// NewListStoresParamsWithContext creates a new ListStoresParams object
// with the default values initialized, and the ability to set a context for a request
func NewListStoresParamsWithContext(ctx context.Context) *ListStoresParams {
	var ()
	return &ListStoresParams{

		Context: ctx,
	}
}

// NewListStoresParamsWithHTTPClient creates a new ListStoresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListStoresParamsWithHTTPClient(client *http.Client) *ListStoresParams {
	var ()
	return &ListStoresParams{
		HTTPClient: client,
	}
}

/*ListStoresParams contains all the parameters to send to the API endpoint
for the list stores operation typically these are written to a http.Request
*/
type ListStoresParams struct {

	/*BrandID*/
	BrandID string
	/*CountryID*/
	CountryID string
	/*NextToken*/
	NextToken *string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list stores params
func (o *ListStoresParams) WithTimeout(timeout time.Duration) *ListStoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list stores params
func (o *ListStoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list stores params
func (o *ListStoresParams) WithContext(ctx context.Context) *ListStoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list stores params
func (o *ListStoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list stores params
func (o *ListStoresParams) WithHTTPClient(client *http.Client) *ListStoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list stores params
func (o *ListStoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBrandID adds the brandID to the list stores params
func (o *ListStoresParams) WithBrandID(brandID string) *ListStoresParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the list stores params
func (o *ListStoresParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithCountryID adds the countryID to the list stores params
func (o *ListStoresParams) WithCountryID(countryID string) *ListStoresParams {
	o.SetCountryID(countryID)
	return o
}

// SetCountryID adds the countryId to the list stores params
func (o *ListStoresParams) SetCountryID(countryID string) {
	o.CountryID = countryID
}

// WithNextToken adds the nextToken to the list stores params
func (o *ListStoresParams) WithNextToken(nextToken *string) *ListStoresParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list stores params
func (o *ListStoresParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WithTenantID adds the tenantID to the list stores params
func (o *ListStoresParams) WithTenantID(tenantID string) *ListStoresParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list stores params
func (o *ListStoresParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *ListStoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
