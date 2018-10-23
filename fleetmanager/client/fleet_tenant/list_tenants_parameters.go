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

// NewListTenantsParams creates a new ListTenantsParams object
// with the default values initialized.
func NewListTenantsParams() *ListTenantsParams {
	var ()
	return &ListTenantsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTenantsParamsWithTimeout creates a new ListTenantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTenantsParamsWithTimeout(timeout time.Duration) *ListTenantsParams {
	var ()
	return &ListTenantsParams{

		timeout: timeout,
	}
}

// NewListTenantsParamsWithContext creates a new ListTenantsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListTenantsParamsWithContext(ctx context.Context) *ListTenantsParams {
	var ()
	return &ListTenantsParams{

		Context: ctx,
	}
}

// NewListTenantsParamsWithHTTPClient creates a new ListTenantsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTenantsParamsWithHTTPClient(client *http.Client) *ListTenantsParams {
	var ()
	return &ListTenantsParams{
		HTTPClient: client,
	}
}

/*ListTenantsParams contains all the parameters to send to the API endpoint
for the list tenants operation typically these are written to a http.Request
*/
type ListTenantsParams struct {

	/*NextToken*/
	NextToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tenants params
func (o *ListTenantsParams) WithTimeout(timeout time.Duration) *ListTenantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tenants params
func (o *ListTenantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tenants params
func (o *ListTenantsParams) WithContext(ctx context.Context) *ListTenantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tenants params
func (o *ListTenantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tenants params
func (o *ListTenantsParams) WithHTTPClient(client *http.Client) *ListTenantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tenants params
func (o *ListTenantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextToken adds the nextToken to the list tenants params
func (o *ListTenantsParams) WithNextToken(nextToken *string) *ListTenantsParams {
	o.SetNextToken(nextToken)
	return o
}

// SetNextToken adds the nextToken to the list tenants params
func (o *ListTenantsParams) SetNextToken(nextToken *string) {
	o.NextToken = nextToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListTenantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
