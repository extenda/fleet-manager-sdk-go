// Code generated by go-swagger; DO NOT EDIT.

package driver_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteDriverPackagePackageIDVersionVersionIDReader is a Reader for the DeleteDriverPackagePackageIDVersionVersionID structure.
type DeleteDriverPackagePackageIDVersionVersionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDriverPackagePackageIDVersionVersionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDriverPackagePackageIDVersionVersionIDNoContent creates a DeleteDriverPackagePackageIDVersionVersionIDNoContent with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDNoContent() *DeleteDriverPackagePackageIDVersionVersionIDNoContent {
	return &DeleteDriverPackagePackageIDVersionVersionIDNoContent{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDNoContent handles this case with default header values.

No Content
*/
type DeleteDriverPackagePackageIDVersionVersionIDNoContent struct {
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}][%d] deleteDriverPackagePackageIdVersionVersionIdNoContent ", 204)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBadRequest creates a DeleteDriverPackagePackageIDVersionVersionIDBadRequest with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDBadRequest() *DeleteDriverPackagePackageIDVersionVersionIDBadRequest {
	return &DeleteDriverPackagePackageIDVersionVersionIDBadRequest{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteDriverPackagePackageIDVersionVersionIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}][%d] deleteDriverPackagePackageIdVersionVersionIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDriverPackagePackageIDVersionVersionIDNotFound creates a DeleteDriverPackagePackageIDVersionVersionIDNotFound with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDNotFound() *DeleteDriverPackagePackageIDVersionVersionIDNotFound {
	return &DeleteDriverPackagePackageIDVersionVersionIDNotFound{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDNotFound handles this case with default header values.

Not Found
*/
type DeleteDriverPackagePackageIDVersionVersionIDNotFound struct {
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}][%d] deleteDriverPackagePackageIdVersionVersionIdNotFound ", 404)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity creates a DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity() *DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity {
	return &DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity struct {
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}][%d] deleteDriverPackagePackageIdVersionVersionIdUnprocessableEntity ", 422)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
