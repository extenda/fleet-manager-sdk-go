// Code generated by go-swagger; DO NOT EDIT.

package fleet_workstation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationReader is a Reader for the PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstation structure.
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated creates a PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated with default headers values
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated() *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated {
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated{}
}

/*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated handles this case with default header values.

Created
*/
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/workstation][%d] postFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdWorkstationCreated ", 201)
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest creates a PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest with default headers values
func NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest() *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest {
	return &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest{}
}

/*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest handles this case with default header values.

Bad Request
*/
type PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest struct {
	Payload *models.Error
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/workstation][%d] postFleetTenantTenantIdBrandBrandIdCountryCountryIdStoreStoreIdWorkstationBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDWorkstationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
