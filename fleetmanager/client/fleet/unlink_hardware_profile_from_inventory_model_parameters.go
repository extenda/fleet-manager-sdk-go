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

// NewUnlinkHardwareProfileFromInventoryModelParams creates a new UnlinkHardwareProfileFromInventoryModelParams object
// with the default values initialized.
func NewUnlinkHardwareProfileFromInventoryModelParams() *UnlinkHardwareProfileFromInventoryModelParams {
	var ()
	return &UnlinkHardwareProfileFromInventoryModelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnlinkHardwareProfileFromInventoryModelParamsWithTimeout creates a new UnlinkHardwareProfileFromInventoryModelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnlinkHardwareProfileFromInventoryModelParamsWithTimeout(timeout time.Duration) *UnlinkHardwareProfileFromInventoryModelParams {
	var ()
	return &UnlinkHardwareProfileFromInventoryModelParams{

		timeout: timeout,
	}
}

// NewUnlinkHardwareProfileFromInventoryModelParamsWithContext creates a new UnlinkHardwareProfileFromInventoryModelParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnlinkHardwareProfileFromInventoryModelParamsWithContext(ctx context.Context) *UnlinkHardwareProfileFromInventoryModelParams {
	var ()
	return &UnlinkHardwareProfileFromInventoryModelParams{

		Context: ctx,
	}
}

// NewUnlinkHardwareProfileFromInventoryModelParamsWithHTTPClient creates a new UnlinkHardwareProfileFromInventoryModelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUnlinkHardwareProfileFromInventoryModelParamsWithHTTPClient(client *http.Client) *UnlinkHardwareProfileFromInventoryModelParams {
	var ()
	return &UnlinkHardwareProfileFromInventoryModelParams{
		HTTPClient: client,
	}
}

/*UnlinkHardwareProfileFromInventoryModelParams contains all the parameters to send to the API endpoint
for the unlink hardware profile from inventory model operation typically these are written to a http.Request
*/
type UnlinkHardwareProfileFromInventoryModelParams struct {

	/*Body*/
	Body *models.UnlinkFleetHardwareProfile2InventoryModel
	/*BrandID*/
	BrandID string
	/*HardwareProfileID*/
	HardwareProfileID string
	/*TenantID*/
	TenantID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithTimeout(timeout time.Duration) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithContext(ctx context.Context) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithHTTPClient(client *http.Client) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithBody(body *models.UnlinkFleetHardwareProfile2InventoryModel) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetBody(body *models.UnlinkFleetHardwareProfile2InventoryModel) {
	o.Body = body
}

// WithBrandID adds the brandID to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithBrandID(brandID string) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetBrandID(brandID)
	return o
}

// SetBrandID adds the brandId to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetBrandID(brandID string) {
	o.BrandID = brandID
}

// WithHardwareProfileID adds the hardwareProfileID to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithHardwareProfileID(hardwareProfileID string) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetHardwareProfileID(hardwareProfileID)
	return o
}

// SetHardwareProfileID adds the hardwareProfileId to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetHardwareProfileID(hardwareProfileID string) {
	o.HardwareProfileID = hardwareProfileID
}

// WithTenantID adds the tenantID to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) WithTenantID(tenantID string) *UnlinkHardwareProfileFromInventoryModelParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the unlink hardware profile from inventory model params
func (o *UnlinkHardwareProfileFromInventoryModelParams) SetTenantID(tenantID string) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *UnlinkHardwareProfileFromInventoryModelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param hardwareProfileId
	if err := r.SetPathParam("hardwareProfileId", o.HardwareProfileID); err != nil {
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
