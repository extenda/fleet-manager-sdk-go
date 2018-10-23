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

// DeleteHardwareProfileReader is a Reader for the DeleteHardwareProfile structure.
type DeleteHardwareProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHardwareProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteHardwareProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteHardwareProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteHardwareProfileNoContent creates a DeleteHardwareProfileNoContent with default headers values
func NewDeleteHardwareProfileNoContent() *DeleteHardwareProfileNoContent {
	return &DeleteHardwareProfileNoContent{}
}

/*DeleteHardwareProfileNoContent handles this case with default header values.

No Content
*/
type DeleteHardwareProfileNoContent struct {
}

func (o *DeleteHardwareProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] deleteHardwareProfileNoContent ", 204)
}

func (o *DeleteHardwareProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardwareProfileBadRequest creates a DeleteHardwareProfileBadRequest with default headers values
func NewDeleteHardwareProfileBadRequest() *DeleteHardwareProfileBadRequest {
	return &DeleteHardwareProfileBadRequest{}
}

/*DeleteHardwareProfileBadRequest handles this case with default header values.

Bad Request
*/
type DeleteHardwareProfileBadRequest struct {
	Payload *models.Error
}

func (o *DeleteHardwareProfileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] deleteHardwareProfileBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteHardwareProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}