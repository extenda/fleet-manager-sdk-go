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

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// NewCreateSoftwareProfileForBrandParams creates a new CreateSoftwareProfileForBrandParams object
// with the default values initialized.
func NewCreateSoftwareProfileForBrandParams() *CreateSoftwareProfileForBrandParams {
	var ()
	return &CreateSoftwareProfileForBrandParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSoftwareProfileForBrandParamsWithTimeout creates a new CreateSoftwareProfileForBrandParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSoftwareProfileForBrandParamsWithTimeout(timeout time.Duration) *CreateSoftwareProfileForBrandParams {
	var ()
	return &CreateSoftwareProfileForBrandParams{

		timeout: timeout,
	}
}

// NewCreateSoftwareProfileForBrandParamsWithContext creates a new CreateSoftwareProfileForBrandParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSoftwareProfileForBrandParamsWithContext(ctx context.Context) *CreateSoftwareProfileForBrandParams {
	var ()
	return &CreateSoftwareProfileForBrandParams{

		Context: ctx,
	}
}

// NewCreateSoftwareProfileForBrandParamsWithHTTPClient creates a new CreateSoftwareProfileForBrandParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSoftwareProfileForBrandParamsWithHTTPClient(client *http.Client) *CreateSoftwareProfileForBrandParams {
	var ()
	return &CreateSoftwareProfileForBrandParams{
		HTTPClient: client,
	}
}

/*CreateSoftwareProfileForBrandParams contains all the parameters to send to the API endpoint
for the create software profile for brand operation typically these are written to a http.Request
*/
type CreateSoftwareProfileForBrandParams struct {

	/*Body*/
	Body *models.CreateFleetSoftwareProfile
	/*BrandID*/
	BrandID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithTimeout(timeout time.Duration) *CreateSoftwareProfileForBrandParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithContext(ctx context.Context) *CreateSoftwareProfileForBrandParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithHTTPClient(client *http.Client) *CreateSoftwareProfileForBrandParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithBody(body *models.CreateFleetSoftwareProfile) *CreateSoftwareProfileForBrandParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetBody(body *models.CreateFleetSoftwareProfile) {
	o.Body = body
}

// WithBrandID adds the brandID to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithBrandID(brandID string) *CreateSoftwareProfileForBrandParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithTenantID adds the tenantID to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) WithTenantID(tenantID string) *CreateSoftwareProfileForBrandParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the create software profile for brand params
func (o *CreateSoftwareProfileForBrandParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSoftwareProfileForBrandParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
