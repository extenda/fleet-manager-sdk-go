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

// NewDeleteFleetTenantTenantIDSystemPropertiesURLParams creates a new DeleteFleetTenantTenantIDSystemPropertiesURLParams object
// with the default values initialized.
func NewDeleteFleetTenantTenantIDSystemPropertiesURLParams() *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	var ()
	return &DeleteFleetTenantTenantIDSystemPropertiesURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithTimeout creates a new DeleteFleetTenantTenantIDSystemPropertiesURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithTimeout(timeout time.Duration) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	var ()
	return &DeleteFleetTenantTenantIDSystemPropertiesURLParams{

		timeout: timeout,
	}
}

// NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithContext creates a new DeleteFleetTenantTenantIDSystemPropertiesURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithContext(ctx context.Context) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	var ()
	return &DeleteFleetTenantTenantIDSystemPropertiesURLParams{

		Context: ctx,
	}
}

// NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithHTTPClient creates a new DeleteFleetTenantTenantIDSystemPropertiesURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFleetTenantTenantIDSystemPropertiesURLParamsWithHTTPClient(client *http.Client) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	var ()
	return &DeleteFleetTenantTenantIDSystemPropertiesURLParams{
		HTTPClient: client,
	}
}

/*DeleteFleetTenantTenantIDSystemPropertiesURLParams contains all the parameters to send to the API endpoint
for the delete fleet tenant tenant ID system properties URL operation typically these are written to a http.Request
*/
type DeleteFleetTenantTenantIDSystemPropertiesURLParams struct {

	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) WithTimeout(timeout time.Duration) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) WithContext(ctx context.Context) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) WithHTTPClient(client *http.Client) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) WithTenantID(tenantID string) *DeleteFleetTenantTenantIDSystemPropertiesURLParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the delete fleet tenant tenant ID system properties URL params
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFleetTenantTenantIDSystemPropertiesURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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