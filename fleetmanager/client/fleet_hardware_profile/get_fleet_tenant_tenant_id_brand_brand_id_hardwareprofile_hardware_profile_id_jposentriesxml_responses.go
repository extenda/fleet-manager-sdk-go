// Code generated by go-swagger; DO NOT EDIT.

package fleet_hardware_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlReader is a Reader for the GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxml structure.
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK struct {
	Payload *models.S3PresignedPost
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/jposentriesxml][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdJposentriesxmlOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/jposentriesxml][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdJposentriesxmlBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound creates a GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound() *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound {
	return &GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound{}
}

/*GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound handles this case with default header values.

Not Found
*/
type GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound struct {
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/jposentriesxml][%d] getFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdJposentriesxmlNotFound ", 404)
}

func (o *GetFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposentriesxmlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
