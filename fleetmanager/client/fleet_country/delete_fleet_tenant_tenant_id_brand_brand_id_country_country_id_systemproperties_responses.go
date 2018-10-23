// Code generated by go-swagger; DO NOT EDIT.

package fleet_country

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesReader is a Reader for the DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystemproperties structure.
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent creates a DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent() *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent {
	return &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent handles this case with default header values.

No Content
*/
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent struct {
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] deleteFleetTenantTenantIdBrandBrandIdCountryCountryIdSystempropertiesNoContent ", 204)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest creates a DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest() *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest {
	return &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest handles this case with default header values.

Bad Request
*/
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/systemproperties][%d] deleteFleetTenantTenantIdBrandBrandIdCountryCountryIdSystempropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDSystempropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
