// Code generated by go-swagger; DO NOT EDIT.

package inventory_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// PostInventoryManufacturerManufacturerIDModelModelIDReader is a Reader for the PostInventoryManufacturerManufacturerIDModelModelID structure.
type PostInventoryManufacturerManufacturerIDModelModelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostInventoryManufacturerManufacturerIDModelModelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostInventoryManufacturerManufacturerIDModelModelIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInventoryManufacturerManufacturerIDModelModelIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostInventoryManufacturerManufacturerIDModelModelIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInventoryManufacturerManufacturerIDModelModelIDNoContent creates a PostInventoryManufacturerManufacturerIDModelModelIDNoContent with default headers values
func NewPostInventoryManufacturerManufacturerIDModelModelIDNoContent() *PostInventoryManufacturerManufacturerIDModelModelIDNoContent {
	return &PostInventoryManufacturerManufacturerIDModelModelIDNoContent{}
}

/*PostInventoryManufacturerManufacturerIDModelModelIDNoContent handles this case with default header values.

No Content
*/
type PostInventoryManufacturerManufacturerIDModelModelIDNoContent struct {
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDNoContent) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}][%d] postInventoryManufacturerManufacturerIdModelModelIdNoContent ", 204)
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostInventoryManufacturerManufacturerIDModelModelIDBadRequest creates a PostInventoryManufacturerManufacturerIDModelModelIDBadRequest with default headers values
func NewPostInventoryManufacturerManufacturerIDModelModelIDBadRequest() *PostInventoryManufacturerManufacturerIDModelModelIDBadRequest {
	return &PostInventoryManufacturerManufacturerIDModelModelIDBadRequest{}
}

/*PostInventoryManufacturerManufacturerIDModelModelIDBadRequest handles this case with default header values.

Bad Request
*/
type PostInventoryManufacturerManufacturerIDModelModelIDBadRequest struct {
	Payload *models.Error
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDBadRequest) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}][%d] postInventoryManufacturerManufacturerIdModelModelIdBadRequest  %+v", 400, o.Payload)
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostInventoryManufacturerManufacturerIDModelModelIDNotFound creates a PostInventoryManufacturerManufacturerIDModelModelIDNotFound with default headers values
func NewPostInventoryManufacturerManufacturerIDModelModelIDNotFound() *PostInventoryManufacturerManufacturerIDModelModelIDNotFound {
	return &PostInventoryManufacturerManufacturerIDModelModelIDNotFound{}
}

/*PostInventoryManufacturerManufacturerIDModelModelIDNotFound handles this case with default header values.

Not Found
*/
type PostInventoryManufacturerManufacturerIDModelModelIDNotFound struct {
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDNotFound) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}][%d] postInventoryManufacturerManufacturerIdModelModelIdNotFound ", 404)
}

func (o *PostInventoryManufacturerManufacturerIDModelModelIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
