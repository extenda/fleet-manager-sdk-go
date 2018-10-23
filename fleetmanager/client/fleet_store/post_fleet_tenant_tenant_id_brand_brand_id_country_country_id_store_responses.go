// Code generated by go-swagger; DO NOT EDIT.

package fleet_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreReader is a Reader for the PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStore structure.
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated creates a PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated with default headers values
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated() *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated {
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated{}
}

/*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated handles this case with default header values.

Created
*/
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store][%d] postFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreCreated ", 201)
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest creates a PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest with default headers values
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest() *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest {
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest{}
}

/*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest handles this case with default header values.

Bad Request
*/
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest struct {
	Payload *models.Error
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store][%d] postFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
