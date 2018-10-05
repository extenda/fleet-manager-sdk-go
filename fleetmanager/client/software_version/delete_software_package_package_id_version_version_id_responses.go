// Code generated by go-swagger; DO NOT EDIT.

package software_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteSoftwarePackagePackageIDVersionVersionIDReader is a Reader for the DeleteSoftwarePackagePackageIDVersionVersionID structure.
type DeleteSoftwarePackagePackageIDVersionVersionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSoftwarePackagePackageIDVersionVersionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSoftwarePackagePackageIDVersionVersionIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteSoftwarePackagePackageIDVersionVersionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteSoftwarePackagePackageIDVersionVersionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSoftwarePackagePackageIDVersionVersionIDNoContent creates a DeleteSoftwarePackagePackageIDVersionVersionIDNoContent with default headers values
func NewDeleteSoftwarePackagePackageIDVersionVersionIDNoContent() *DeleteSoftwarePackagePackageIDVersionVersionIDNoContent {
	return &DeleteSoftwarePackagePackageIDVersionVersionIDNoContent{}
}

/*DeleteSoftwarePackagePackageIDVersionVersionIDNoContent handles this case with default header values.

No Content
*/
type DeleteSoftwarePackagePackageIDVersionVersionIDNoContent struct {
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}/version/{versionId}][%d] deleteSoftwarePackagePackageIdVersionVersionIdNoContent ", 204)
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwarePackagePackageIDVersionVersionIDBadRequest creates a DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest with default headers values
func NewDeleteSoftwarePackagePackageIDVersionVersionIDBadRequest() *DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest {
	return &DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest{}
}

/*DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}/version/{versionId}][%d] deleteSoftwarePackagePackageIdVersionVersionIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSoftwarePackagePackageIDVersionVersionIDNotFound creates a DeleteSoftwarePackagePackageIDVersionVersionIDNotFound with default headers values
func NewDeleteSoftwarePackagePackageIDVersionVersionIDNotFound() *DeleteSoftwarePackagePackageIDVersionVersionIDNotFound {
	return &DeleteSoftwarePackagePackageIDVersionVersionIDNotFound{}
}

/*DeleteSoftwarePackagePackageIDVersionVersionIDNotFound handles this case with default header values.

Not Found
*/
type DeleteSoftwarePackagePackageIDVersionVersionIDNotFound struct {
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}/version/{versionId}][%d] deleteSoftwarePackagePackageIdVersionVersionIdNotFound ", 404)
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity creates a DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity with default headers values
func NewDeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity() *DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity {
	return &DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity{}
}

/*DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity struct {
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /software/package/{packageId}/version/{versionId}][%d] deleteSoftwarePackagePackageIdVersionVersionIdUnprocessableEntity ", 422)
}

func (o *DeleteSoftwarePackagePackageIDVersionVersionIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}