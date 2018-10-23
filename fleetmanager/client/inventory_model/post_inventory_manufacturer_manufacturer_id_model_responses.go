// Code generated by go-swagger; DO NOT EDIT.

package inventory_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// PostInventoryManufacturerManufacturerIDModelReader is a Reader for the PostInventoryManufacturerManufacturerIDModel structure.
type PostInventoryManufacturerManufacturerIDModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInventoryManufacturerManufacturerIDModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostInventoryManufacturerManufacturerIDModelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInventoryManufacturerManufacturerIDModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInventoryManufacturerManufacturerIDModelCreated creates a PostInventoryManufacturerManufacturerIDModelCreated with default headers values
func NewPostInventoryManufacturerManufacturerIDModelCreated() *PostInventoryManufacturerManufacturerIDModelCreated {
	return &PostInventoryManufacturerManufacturerIDModelCreated{}
}

/*PostInventoryManufacturerManufacturerIDModelCreated handles this case with default header values.

Created
*/
type PostInventoryManufacturerManufacturerIDModelCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *PostInventoryManufacturerManufacturerIDModelCreated) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model][%d] postInventoryManufacturerManufacturerIdModelCreated ", 201)
}

func (o *PostInventoryManufacturerManufacturerIDModelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewPostInventoryManufacturerManufacturerIDModelBadRequest creates a PostInventoryManufacturerManufacturerIDModelBadRequest with default headers values
func NewPostInventoryManufacturerManufacturerIDModelBadRequest() *PostInventoryManufacturerManufacturerIDModelBadRequest {
	return &PostInventoryManufacturerManufacturerIDModelBadRequest{}
}

/*PostInventoryManufacturerManufacturerIDModelBadRequest handles this case with default header values.

Bad Request
*/
type PostInventoryManufacturerManufacturerIDModelBadRequest struct {
	Payload *models.Error
}

func (o *PostInventoryManufacturerManufacturerIDModelBadRequest) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model][%d] postInventoryManufacturerManufacturerIdModelBadRequest  %+v", 400, o.Payload)
}

func (o *PostInventoryManufacturerManufacturerIDModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
