// Code generated by go-swagger; DO NOT EDIT.

package fleet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// CreateStoreReader is a Reader for the CreateStore structure.
type CreateStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateStoreCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateStoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateStoreCreated creates a CreateStoreCreated with default headers values
func NewCreateStoreCreated() *CreateStoreCreated {
	return &CreateStoreCreated{}
}

/*CreateStoreCreated handles this case with default header values.

Created
*/
type CreateStoreCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *CreateStoreCreated) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store][%d] createStoreCreated ", 201)
}

func (o *CreateStoreCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewCreateStoreBadRequest creates a CreateStoreBadRequest with default headers values
func NewCreateStoreBadRequest() *CreateStoreBadRequest {
	return &CreateStoreBadRequest{}
}

/*CreateStoreBadRequest handles this case with default header values.

Bad Request
*/
type CreateStoreBadRequest struct {
	Payload *models.Error
}

func (o *CreateStoreBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store][%d] createStoreBadRequest  %+v", 400, o.Payload)
}

func (o *CreateStoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}