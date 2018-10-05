// Code generated by go-swagger; DO NOT EDIT.

package fleet_tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetFleetTenantTenantIDReader is a Reader for the GetFleetTenantTenantID structure.
type GetFleetTenantTenantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetTenantTenantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDOK creates a GetFleetTenantTenantIDOK with default headers values
func NewGetFleetTenantTenantIDOK() *GetFleetTenantTenantIDOK {
	return &GetFleetTenantTenantIDOK{}
}

/*GetFleetTenantTenantIDOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDOK struct {
	Payload *models.FleetTenant
}

func (o *GetFleetTenantTenantIDOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getFleetTenantTenantIdOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetTenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBadRequest creates a GetFleetTenantTenantIDBadRequest with default headers values
func NewGetFleetTenantTenantIDBadRequest() *GetFleetTenantTenantIDBadRequest {
	return &GetFleetTenantTenantIDBadRequest{}
}

/*GetFleetTenantTenantIDBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getFleetTenantTenantIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDNotFound creates a GetFleetTenantTenantIDNotFound with default headers values
func NewGetFleetTenantTenantIDNotFound() *GetFleetTenantTenantIDNotFound {
	return &GetFleetTenantTenantIDNotFound{}
}

/*GetFleetTenantTenantIDNotFound handles this case with default header values.

Not Found
*/
type GetFleetTenantTenantIDNotFound struct {
}

func (o *GetFleetTenantTenantIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}][%d] getFleetTenantTenantIdNotFound ", 404)
}

func (o *GetFleetTenantTenantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}