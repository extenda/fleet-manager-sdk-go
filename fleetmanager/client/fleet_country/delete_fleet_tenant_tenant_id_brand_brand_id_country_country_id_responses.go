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

// DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDReader is a Reader for the DeleteFleetTenantTenantIDBrandBrandIDCountryCountryID structure.
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent creates a DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent() *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent {
	return &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent handles this case with default header values.

No Content
*/
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent struct {
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}][%d] deleteFleetTenantTenantIdBrandBrandIdCountryCountryIdNoContent ", 204)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest creates a DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest with default headers values
func NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest() *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest {
	return &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest{}
}

/*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}][%d] deleteFleetTenantTenantIdBrandBrandIdCountryCountryIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
