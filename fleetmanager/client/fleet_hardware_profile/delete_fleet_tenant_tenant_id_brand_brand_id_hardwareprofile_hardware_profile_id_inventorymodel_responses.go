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

// DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelReader is a Reader for the DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodel structure.
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent creates a DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent() *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent {
	return &DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent handles this case with default header values.

No Content
*/
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent struct {
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] deleteFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdInventorymodelNoContent ", 204)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest creates a DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest() *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest {
	return &DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest handles this case with default header values.

Bad Request
*/
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] deleteFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdInventorymodelBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound creates a DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound() *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound {
	return &DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound handles this case with default header values.

Not Found
*/
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound struct {
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] deleteFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdInventorymodelNotFound ", 404)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDInventorymodelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
