// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// LinkInventoryModelWithDriverVersionReader is a Reader for the LinkInventoryModelWithDriverVersion structure.
type LinkInventoryModelWithDriverVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LinkInventoryModelWithDriverVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewLinkInventoryModelWithDriverVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewLinkInventoryModelWithDriverVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewLinkInventoryModelWithDriverVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLinkInventoryModelWithDriverVersionNoContent creates a LinkInventoryModelWithDriverVersionNoContent with default headers values
func NewLinkInventoryModelWithDriverVersionNoContent() *LinkInventoryModelWithDriverVersionNoContent {
	return &LinkInventoryModelWithDriverVersionNoContent{}
}

/*LinkInventoryModelWithDriverVersionNoContent handles this case with default header values.

No Content
*/
type LinkInventoryModelWithDriverVersionNoContent struct {
}

func (o *LinkInventoryModelWithDriverVersionNoContent) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] linkInventoryModelWithDriverVersionNoContent ", 204)
}

func (o *LinkInventoryModelWithDriverVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLinkInventoryModelWithDriverVersionBadRequest creates a LinkInventoryModelWithDriverVersionBadRequest with default headers values
func NewLinkInventoryModelWithDriverVersionBadRequest() *LinkInventoryModelWithDriverVersionBadRequest {
	return &LinkInventoryModelWithDriverVersionBadRequest{}
}

/*LinkInventoryModelWithDriverVersionBadRequest handles this case with default header values.

Bad Request
*/
type LinkInventoryModelWithDriverVersionBadRequest struct {
	Payload *models.Error
}

func (o *LinkInventoryModelWithDriverVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] linkInventoryModelWithDriverVersionBadRequest  %+v", 400, o.Payload)
}

func (o *LinkInventoryModelWithDriverVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLinkInventoryModelWithDriverVersionNotFound creates a LinkInventoryModelWithDriverVersionNotFound with default headers values
func NewLinkInventoryModelWithDriverVersionNotFound() *LinkInventoryModelWithDriverVersionNotFound {
	return &LinkInventoryModelWithDriverVersionNotFound{}
}

/*LinkInventoryModelWithDriverVersionNotFound handles this case with default header values.

Not Found
*/
type LinkInventoryModelWithDriverVersionNotFound struct {
}

func (o *LinkInventoryModelWithDriverVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] linkInventoryModelWithDriverVersionNotFound ", 404)
}

func (o *LinkInventoryModelWithDriverVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
