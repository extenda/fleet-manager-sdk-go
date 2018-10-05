// Code generated by go-swagger; DO NOT EDIT.

package driver_package

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

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// NewPostDriverPackageParams creates a new PostDriverPackageParams object
// with the default values initialized.
func NewPostDriverPackageParams() *PostDriverPackageParams {
	var ()
	return &PostDriverPackageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDriverPackageParamsWithTimeout creates a new PostDriverPackageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDriverPackageParamsWithTimeout(timeout time.Duration) *PostDriverPackageParams {
	var ()
	return &PostDriverPackageParams{

		timeout: timeout,
	}
}

// NewPostDriverPackageParamsWithContext creates a new PostDriverPackageParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDriverPackageParamsWithContext(ctx context.Context) *PostDriverPackageParams {
	var ()
	return &PostDriverPackageParams{

		Context: ctx,
	}
}

// NewPostDriverPackageParamsWithHTTPClient creates a new PostDriverPackageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDriverPackageParamsWithHTTPClient(client *http.Client) *PostDriverPackageParams {
	var ()
	return &PostDriverPackageParams{
		HTTPClient: client,
	}
}

/*PostDriverPackageParams contains all the parameters to send to the API endpoint
for the post driver package operation typically these are written to a http.Request
*/
type PostDriverPackageParams struct {

	/*Body*/
	Body *models.CreateDriverPackage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post driver package params
func (o *PostDriverPackageParams) WithTimeout(timeout time.Duration) *PostDriverPackageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post driver package params
func (o *PostDriverPackageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post driver package params
func (o *PostDriverPackageParams) WithContext(ctx context.Context) *PostDriverPackageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post driver package params
func (o *PostDriverPackageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post driver package params
func (o *PostDriverPackageParams) WithHTTPClient(client *http.Client) *PostDriverPackageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post driver package params
func (o *PostDriverPackageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post driver package params
func (o *PostDriverPackageParams) WithBody(body *models.CreateDriverPackage) *PostDriverPackageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post driver package params
func (o *PostDriverPackageParams) SetBody(body *models.CreateDriverPackage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostDriverPackageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}