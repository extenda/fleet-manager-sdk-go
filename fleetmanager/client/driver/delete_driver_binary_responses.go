// Code generated by go-swagger; DO NOT EDIT.

package driver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// DeleteDriverBinaryReader is a Reader for the DeleteDriverBinary structure.
type DeleteDriverBinaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDriverBinaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDriverBinaryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDriverBinaryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDriverBinaryNoContent creates a DeleteDriverBinaryNoContent with default headers values
func NewDeleteDriverBinaryNoContent() *DeleteDriverBinaryNoContent {
	return &DeleteDriverBinaryNoContent{}
}

/*DeleteDriverBinaryNoContent handles this case with default header values.

No Content
*/
type DeleteDriverBinaryNoContent struct {
}

func (o *DeleteDriverBinaryNoContent) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}/binary][%d] deleteDriverBinaryNoContent ", 204)
}

func (o *DeleteDriverBinaryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverBinaryBadRequest creates a DeleteDriverBinaryBadRequest with default headers values
func NewDeleteDriverBinaryBadRequest() *DeleteDriverBinaryBadRequest {
	return &DeleteDriverBinaryBadRequest{}
}

/*DeleteDriverBinaryBadRequest handles this case with default header values.

Bad Request
*/
type DeleteDriverBinaryBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDriverBinaryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}/binary][%d] deleteDriverBinaryBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDriverBinaryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
