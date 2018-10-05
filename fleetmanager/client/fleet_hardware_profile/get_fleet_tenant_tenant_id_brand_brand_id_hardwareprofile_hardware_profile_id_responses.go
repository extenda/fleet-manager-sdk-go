// Code generated by go-swagger; DO NOT EDIT.

package fleet_hardware_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDReader is a Reader for the GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileID structure.
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK struct {
	Payload *models.FleetHardwareProfile
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetHardwareProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound handles this case with default header values.

Not Found
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound struct {
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdNotFound ", 404)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
