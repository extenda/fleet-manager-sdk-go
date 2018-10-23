// Code generated by go-swagger; DO NOT EDIT.

package fleet_software_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// GetFleetTenantTenantIDBrandBrandIDSoftwareprofileReader is a Reader for the GetFleetTenantTenantIDBrandBrandIDSoftwareprofile structure.
type GetFleetTenantTenantIDBrandBrandIDSoftwareprofileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK creates a GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK() *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK {
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK{}
}

/*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK struct {
	Payload *models.FleetSoftwareProfiles
}

func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile][%d] getFleetTenantTenantIdBrandBrandIdSoftwareprofileOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetSoftwareProfiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest creates a GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest() *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest {
	return &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest{}
}

/*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile][%d] getFleetTenantTenantIdBrandBrandIdSoftwareprofileBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
