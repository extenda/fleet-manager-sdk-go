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

// UnlinkHardwareProfileFromInventoryModelReader is a Reader for the UnlinkHardwareProfileFromInventoryModel structure.
type UnlinkHardwareProfileFromInventoryModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnlinkHardwareProfileFromInventoryModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUnlinkHardwareProfileFromInventoryModelNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUnlinkHardwareProfileFromInventoryModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUnlinkHardwareProfileFromInventoryModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUnlinkHardwareProfileFromInventoryModelNoContent creates a UnlinkHardwareProfileFromInventoryModelNoContent with default headers values
func NewUnlinkHardwareProfileFromInventoryModelNoContent() *UnlinkHardwareProfileFromInventoryModelNoContent {
	return &UnlinkHardwareProfileFromInventoryModelNoContent{}
}

/*UnlinkHardwareProfileFromInventoryModelNoContent handles this case with default header values.

No Content
*/
type UnlinkHardwareProfileFromInventoryModelNoContent struct {
}

func (o *UnlinkHardwareProfileFromInventoryModelNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] unlinkHardwareProfileFromInventoryModelNoContent ", 204)
}

func (o *UnlinkHardwareProfileFromInventoryModelNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnlinkHardwareProfileFromInventoryModelBadRequest creates a UnlinkHardwareProfileFromInventoryModelBadRequest with default headers values
func NewUnlinkHardwareProfileFromInventoryModelBadRequest() *UnlinkHardwareProfileFromInventoryModelBadRequest {
	return &UnlinkHardwareProfileFromInventoryModelBadRequest{}
}

/*UnlinkHardwareProfileFromInventoryModelBadRequest handles this case with default header values.

Bad Request
*/
type UnlinkHardwareProfileFromInventoryModelBadRequest struct {
	Payload *models.Error
}

func (o *UnlinkHardwareProfileFromInventoryModelBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] unlinkHardwareProfileFromInventoryModelBadRequest  %+v", 400, o.Payload)
}

func (o *UnlinkHardwareProfileFromInventoryModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnlinkHardwareProfileFromInventoryModelNotFound creates a UnlinkHardwareProfileFromInventoryModelNotFound with default headers values
func NewUnlinkHardwareProfileFromInventoryModelNotFound() *UnlinkHardwareProfileFromInventoryModelNotFound {
	return &UnlinkHardwareProfileFromInventoryModelNotFound{}
}

/*UnlinkHardwareProfileFromInventoryModelNotFound handles this case with default header values.

Not Found
*/
type UnlinkHardwareProfileFromInventoryModelNotFound struct {
}

func (o *UnlinkHardwareProfileFromInventoryModelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/hardwareprofile/{hardwareProfileId}/inventorymodel][%d] unlinkHardwareProfileFromInventoryModelNotFound ", 404)
}

func (o *UnlinkHardwareProfileFromInventoryModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
