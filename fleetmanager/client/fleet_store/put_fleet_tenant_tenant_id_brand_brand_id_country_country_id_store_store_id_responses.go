// Code generated by go-swagger; DO NOT EDIT.

package fleet_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader is a Reader for the PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID structure.
type PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent creates a PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent with default headers values
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent() *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent {
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent{}
}

/*PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent handles this case with default header values.

No Content
*/
type PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent struct {
}

func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}][%d] putFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdNoContent ", 204)
}

func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest creates a PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest with default headers values
func NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest() *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest {
	return &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest{}
}

/*PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest handles this case with default header values.

Bad Request
*/
type PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest struct {
	Payload *models.Error
}

func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}][%d] putFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}