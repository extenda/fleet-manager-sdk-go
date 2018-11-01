// Code generated by go-swagger; DO NOT EDIT.

package fleet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// UpdateHardwareProfileReader is a Reader for the UpdateHardwareProfile structure.
type UpdateHardwareProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHardwareProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateHardwareProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateHardwareProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateHardwareProfileNoContent creates a UpdateHardwareProfileNoContent with default headers values
func NewUpdateHardwareProfileNoContent() *UpdateHardwareProfileNoContent {
	return &UpdateHardwareProfileNoContent{}
}

/*UpdateHardwareProfileNoContent handles this case with default header values.

No Content
*/
type UpdateHardwareProfileNoContent struct {
}

func (o *UpdateHardwareProfileNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] updateHardwareProfileNoContent ", 204)
}

func (o *UpdateHardwareProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateHardwareProfileBadRequest creates a UpdateHardwareProfileBadRequest with default headers values
func NewUpdateHardwareProfileBadRequest() *UpdateHardwareProfileBadRequest {
	return &UpdateHardwareProfileBadRequest{}
}

/*UpdateHardwareProfileBadRequest handles this case with default header values.

Bad Request
*/
type UpdateHardwareProfileBadRequest struct {
	Payload *models.Error
}

func (o *UpdateHardwareProfileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}][%d] updateHardwareProfileBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateHardwareProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
