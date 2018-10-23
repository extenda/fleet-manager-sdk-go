// Code generated by go-swagger; DO NOT EDIT.

package fleet_tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// GetTenantByIDReader is a Reader for the GetTenantByID structure.
type GetTenantByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTenantByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTenantByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTenantByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTenantByIDOK creates a GetTenantByIDOK with default headers values
func NewGetTenantByIDOK() *GetTenantByIDOK {
	return &GetTenantByIDOK{}
}

/*GetTenantByIDOK handles this case with default header values.

OK
*/
type GetTenantByIDOK struct {
	Payload *models.FleetTenant
}

func (o *GetTenantByIDOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getTenantByIdOK  %+v", 200, o.Payload)
}

func (o *GetTenantByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetTenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantByIDBadRequest creates a GetTenantByIDBadRequest with default headers values
func NewGetTenantByIDBadRequest() *GetTenantByIDBadRequest {
	return &GetTenantByIDBadRequest{}
}

/*GetTenantByIDBadRequest handles this case with default header values.

Bad Request
*/
type GetTenantByIDBadRequest struct {
	Payload *models.Error
}

func (o *GetTenantByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getTenantByIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetTenantByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantByIDNotFound creates a GetTenantByIDNotFound with default headers values
func NewGetTenantByIDNotFound() *GetTenantByIDNotFound {
	return &GetTenantByIDNotFound{}
}

/*GetTenantByIDNotFound handles this case with default header values.

Not Found
*/
type GetTenantByIDNotFound struct {
}

func (o *GetTenantByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getTenantByIdNotFound ", 404)
}

func (o *GetTenantByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
