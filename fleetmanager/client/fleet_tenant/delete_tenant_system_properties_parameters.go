// Code generated by go-swagger; DO NOT EDIT.

package fleet_tenant

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

// NewDeleteTenantSystemPropertiesParams creates a new DeleteTenantSystemPropertiesParams object
// with the default values initialized.
func NewDeleteTenantSystemPropertiesParams() *DeleteTenantSystemPropertiesParams {
	var ()
	return &DeleteTenantSystemPropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTenantSystemPropertiesParamsWithTimeout creates a new DeleteTenantSystemPropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTenantSystemPropertiesParamsWithTimeout(timeout time.Duration) *DeleteTenantSystemPropertiesParams {
	var ()
	return &DeleteTenantSystemPropertiesParams{

		timeout: timeout,
	}
}

// NewDeleteTenantSystemPropertiesParamsWithContext creates a new DeleteTenantSystemPropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTenantSystemPropertiesParamsWithContext(ctx context.Context) *DeleteTenantSystemPropertiesParams {
	var ()
	return &DeleteTenantSystemPropertiesParams{

		Context: ctx,
	}
}

// NewDeleteTenantSystemPropertiesParamsWithHTTPClient creates a new DeleteTenantSystemPropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTenantSystemPropertiesParamsWithHTTPClient(client *http.Client) *DeleteTenantSystemPropertiesParams {
	var ()
	return &DeleteTenantSystemPropertiesParams{
		HTTPClient: client,
	}
}

/*DeleteTenantSystemPropertiesParams contains all the parameters to send to the API endpoint
for the delete tenant system properties operation typically these are written to a http.Request
*/
type DeleteTenantSystemPropertiesParams struct {

	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) WithTimeout(timeout time.Duration) *DeleteTenantSystemPropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) WithContext(ctx context.Context) *DeleteTenantSystemPropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) WithHTTPClient(client *http.Client) *DeleteTenantSystemPropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) WithTenantID(tenantID string) *DeleteTenantSystemPropertiesParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete tenant system properties params
func (o *DeleteTenantSystemPropertiesParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTenantSystemPropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
