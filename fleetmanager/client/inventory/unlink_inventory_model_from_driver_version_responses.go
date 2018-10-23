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

// UnlinkInventoryModelFromDriverVersionReader is a Reader for the UnlinkInventoryModelFromDriverVersion structure.
type UnlinkInventoryModelFromDriverVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlinkInventoryModelFromDriverVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUnlinkInventoryModelFromDriverVersionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUnlinkInventoryModelFromDriverVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUnlinkInventoryModelFromDriverVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnlinkInventoryModelFromDriverVersionNoContent creates a UnlinkInventoryModelFromDriverVersionNoContent with default headers values
func NewUnlinkInventoryModelFromDriverVersionNoContent() *UnlinkInventoryModelFromDriverVersionNoContent {
	return &UnlinkInventoryModelFromDriverVersionNoContent{}
}

/*UnlinkInventoryModelFromDriverVersionNoContent handles this case with default header values.

No Content
*/
type UnlinkInventoryModelFromDriverVersionNoContent struct {
}

func (o *UnlinkInventoryModelFromDriverVersionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] unlinkInventoryModelFromDriverVersionNoContent ", 204)
}

func (o *UnlinkInventoryModelFromDriverVersionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkInventoryModelFromDriverVersionBadRequest creates a UnlinkInventoryModelFromDriverVersionBadRequest with default headers values
func NewUnlinkInventoryModelFromDriverVersionBadRequest() *UnlinkInventoryModelFromDriverVersionBadRequest {
	return &UnlinkInventoryModelFromDriverVersionBadRequest{}
}

/*UnlinkInventoryModelFromDriverVersionBadRequest handles this case with default header values.

Bad Request
*/
type UnlinkInventoryModelFromDriverVersionBadRequest struct {
	Payload *models.Error
}

func (o *UnlinkInventoryModelFromDriverVersionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] unlinkInventoryModelFromDriverVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UnlinkInventoryModelFromDriverVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnlinkInventoryModelFromDriverVersionNotFound creates a UnlinkInventoryModelFromDriverVersionNotFound with default headers values
func NewUnlinkInventoryModelFromDriverVersionNotFound() *UnlinkInventoryModelFromDriverVersionNotFound {
	return &UnlinkInventoryModelFromDriverVersionNotFound{}
}

/*UnlinkInventoryModelFromDriverVersionNotFound handles this case with default header values.

Not Found
*/
type UnlinkInventoryModelFromDriverVersionNotFound struct {
}

func (o *UnlinkInventoryModelFromDriverVersionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /inventory/manufacturer/{manufacturerId}/model/{modelId}/driverversion][%d] unlinkInventoryModelFromDriverVersionNotFound ", 404)
}

func (o *UnlinkInventoryModelFromDriverVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
