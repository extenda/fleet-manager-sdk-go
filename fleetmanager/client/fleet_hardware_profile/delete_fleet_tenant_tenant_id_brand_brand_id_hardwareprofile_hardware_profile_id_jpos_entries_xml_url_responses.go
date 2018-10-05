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

// DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLReader is a Reader for the DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURL structure.
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent creates a DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent() *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent {
	return &DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent handles this case with default header values.

No Content
*/
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent struct {
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/jposEntriesXmlUrl][%d] deleteFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdJposEntriesXmlUrlNoContent ", 204)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest creates a DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest() *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest {
	return &DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest handles this case with default header values.

Bad Request
*/
type DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/jposEntriesXmlUrl][%d] deleteFleetTenantTenantIdBrandBrandIdHardwareprofileHardwareProfileIdJposEntriesXmlUrlBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDHardwareprofileHardwareProfileIDJposEntriesXMLURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
