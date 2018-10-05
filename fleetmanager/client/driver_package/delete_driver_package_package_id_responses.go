// Code generated by go-swagger; DO NOT EDIT.

package driver_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteDriverPackagePackageIDReader is a Reader for the DeleteDriverPackagePackageID structure.
type DeleteDriverPackagePackageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDriverPackagePackageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDriverPackagePackageIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDriverPackagePackageIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteDriverPackagePackageIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteDriverPackagePackageIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDriverPackagePackageIDNoContent creates a DeleteDriverPackagePackageIDNoContent with default headers values
func NewDeleteDriverPackagePackageIDNoContent() *DeleteDriverPackagePackageIDNoContent {
	return &DeleteDriverPackagePackageIDNoContent{}
}

/*DeleteDriverPackagePackageIDNoContent handles this case with default header values.

No Content
*/
type DeleteDriverPackagePackageIDNoContent struct {
}

func (o *DeleteDriverPackagePackageIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}][%d] deleteDriverPackagePackageIdNoContent ", 204)
}

func (o *DeleteDriverPackagePackageIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverPackagePackageIDBadRequest creates a DeleteDriverPackagePackageIDBadRequest with default headers values
func NewDeleteDriverPackagePackageIDBadRequest() *DeleteDriverPackagePackageIDBadRequest {
	return &DeleteDriverPackagePackageIDBadRequest{}
}

/*DeleteDriverPackagePackageIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteDriverPackagePackageIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDriverPackagePackageIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}][%d] deleteDriverPackagePackageIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDriverPackagePackageIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDriverPackagePackageIDNotFound creates a DeleteDriverPackagePackageIDNotFound with default headers values
func NewDeleteDriverPackagePackageIDNotFound() *DeleteDriverPackagePackageIDNotFound {
	return &DeleteDriverPackagePackageIDNotFound{}
}

/*DeleteDriverPackagePackageIDNotFound handles this case with default header values.

Not Found
*/
type DeleteDriverPackagePackageIDNotFound struct {
}

func (o *DeleteDriverPackagePackageIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}][%d] deleteDriverPackagePackageIdNotFound ", 404)
}

func (o *DeleteDriverPackagePackageIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverPackagePackageIDUnprocessableEntity creates a DeleteDriverPackagePackageIDUnprocessableEntity with default headers values
func NewDeleteDriverPackagePackageIDUnprocessableEntity() *DeleteDriverPackagePackageIDUnprocessableEntity {
	return &DeleteDriverPackagePackageIDUnprocessableEntity{}
}

/*DeleteDriverPackagePackageIDUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type DeleteDriverPackagePackageIDUnprocessableEntity struct {
}

func (o *DeleteDriverPackagePackageIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}][%d] deleteDriverPackagePackageIdUnprocessableEntity ", 422)
}

func (o *DeleteDriverPackagePackageIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}