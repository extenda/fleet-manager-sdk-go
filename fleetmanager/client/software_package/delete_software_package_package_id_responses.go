// Code generated by go-swagger; DO NOT EDIT.

package software_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteSoftwarePackagePackageIDReader is a Reader for the DeleteSoftwarePackagePackageID structure.
type DeleteSoftwarePackagePackageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSoftwarePackagePackageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSoftwarePackagePackageIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSoftwarePackagePackageIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSoftwarePackagePackageIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteSoftwarePackagePackageIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSoftwarePackagePackageIDNoContent creates a DeleteSoftwarePackagePackageIDNoContent with default headers values
func NewDeleteSoftwarePackagePackageIDNoContent() *DeleteSoftwarePackagePackageIDNoContent {
	return &DeleteSoftwarePackagePackageIDNoContent{}
}

/*DeleteSoftwarePackagePackageIDNoContent handles this case with default header values.

No Content
*/
type DeleteSoftwarePackagePackageIDNoContent struct {
}

func (o *DeleteSoftwarePackagePackageIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}][%d] deleteSoftwarePackagePackageIdNoContent ", 204)
}

func (o *DeleteSoftwarePackagePackageIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwarePackagePackageIDBadRequest creates a DeleteSoftwarePackagePackageIDBadRequest with default headers values
func NewDeleteSoftwarePackagePackageIDBadRequest() *DeleteSoftwarePackagePackageIDBadRequest {
	return &DeleteSoftwarePackagePackageIDBadRequest{}
}

/*DeleteSoftwarePackagePackageIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSoftwarePackagePackageIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteSoftwarePackagePackageIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}][%d] deleteSoftwarePackagePackageIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSoftwarePackagePackageIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSoftwarePackagePackageIDNotFound creates a DeleteSoftwarePackagePackageIDNotFound with default headers values
func NewDeleteSoftwarePackagePackageIDNotFound() *DeleteSoftwarePackagePackageIDNotFound {
	return &DeleteSoftwarePackagePackageIDNotFound{}
}

/*DeleteSoftwarePackagePackageIDNotFound handles this case with default header values.

Not Found
*/
type DeleteSoftwarePackagePackageIDNotFound struct {
}

func (o *DeleteSoftwarePackagePackageIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}][%d] deleteSoftwarePackagePackageIdNotFound ", 404)
}

func (o *DeleteSoftwarePackagePackageIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwarePackagePackageIDUnprocessableEntity creates a DeleteSoftwarePackagePackageIDUnprocessableEntity with default headers values
func NewDeleteSoftwarePackagePackageIDUnprocessableEntity() *DeleteSoftwarePackagePackageIDUnprocessableEntity {
	return &DeleteSoftwarePackagePackageIDUnprocessableEntity{}
}

/*DeleteSoftwarePackagePackageIDUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type DeleteSoftwarePackagePackageIDUnprocessableEntity struct {
}

func (o *DeleteSoftwarePackagePackageIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}][%d] deleteSoftwarePackagePackageIdUnprocessableEntity ", 422)
}

func (o *DeleteSoftwarePackagePackageIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}