// Code generated by go-swagger; DO NOT EDIT.

package fleet_brand

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// PostFleetTenantTenantIDBrandReader is a Reader for the PostFleetTenantTenantIDBrand structure.
type PostFleetTenantTenantIDBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetTenantTenantIDBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFleetTenantTenantIDBrandCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFleetTenantTenantIDBrandBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetTenantTenantIDBrandCreated creates a PostFleetTenantTenantIDBrandCreated with default headers values
func NewPostFleetTenantTenantIDBrandCreated() *PostFleetTenantTenantIDBrandCreated {
	return &PostFleetTenantTenantIDBrandCreated{}
}

/*PostFleetTenantTenantIDBrandCreated handles this case with default header values.

Created
*/
type PostFleetTenantTenantIDBrandCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *PostFleetTenantTenantIDBrandCreated) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand][%d] postFleetTenantTenantIdBrandCreated ", 201)
}

func (o *PostFleetTenantTenantIDBrandCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostFleetTenantTenantIDBrandBadRequest creates a PostFleetTenantTenantIDBrandBadRequest with default headers values
func NewPostFleetTenantTenantIDBrandBadRequest() *PostFleetTenantTenantIDBrandBadRequest {
	return &PostFleetTenantTenantIDBrandBadRequest{}
}

/*PostFleetTenantTenantIDBrandBadRequest handles this case with default header values.

Bad Request
*/
type PostFleetTenantTenantIDBrandBadRequest struct {
	Payload *models.Error
}

func (o *PostFleetTenantTenantIDBrandBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand][%d] postFleetTenantTenantIdBrandBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetTenantTenantIDBrandBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
