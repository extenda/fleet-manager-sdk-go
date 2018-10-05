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

// GetFleetTenantReader is a Reader for the GetFleetTenant structure.
type GetFleetTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantOK creates a GetFleetTenantOK with default headers values
func NewGetFleetTenantOK() *GetFleetTenantOK {
	return &GetFleetTenantOK{}
}

/*GetFleetTenantOK handles this case with default header values.

OK
*/
type GetFleetTenantOK struct {
	Payload *models.FleetTenants
}

func (o *GetFleetTenantOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant][%d] getFleetTenantOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetTenants)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantBadRequest creates a GetFleetTenantBadRequest with default headers values
func NewGetFleetTenantBadRequest() *GetFleetTenantBadRequest {
	return &GetFleetTenantBadRequest{}
}

/*GetFleetTenantBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant][%d] getFleetTenantBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}