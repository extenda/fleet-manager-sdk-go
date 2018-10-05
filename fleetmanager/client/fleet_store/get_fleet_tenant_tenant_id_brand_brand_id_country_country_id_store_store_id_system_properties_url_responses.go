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

// GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLReader is a Reader for the GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURL structure.
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK creates a GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK {
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK{}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK struct {
	Payload *models.S3PresignedPost
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/systemPropertiesUrl][%d] getFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdSystemPropertiesUrlOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest creates a GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest {
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest{}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/systemPropertiesUrl][%d] getFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdSystemPropertiesUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound creates a GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound with default headers values
func NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound() *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound {
	return &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound{}
}

/*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound handles this case with default header values.

Not Found
*/
type GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound struct {
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/systemPropertiesUrl][%d] getFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdSystemPropertiesUrlNotFound ", 404)
}

func (o *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemPropertiesURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}