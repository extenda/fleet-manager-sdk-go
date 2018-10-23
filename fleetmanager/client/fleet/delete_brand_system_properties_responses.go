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

// DeleteBrandSystemPropertiesReader is a Reader for the DeleteBrandSystemProperties structure.
type DeleteBrandSystemPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBrandSystemPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteBrandSystemPropertiesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteBrandSystemPropertiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBrandSystemPropertiesNoContent creates a DeleteBrandSystemPropertiesNoContent with default headers values
func NewDeleteBrandSystemPropertiesNoContent() *DeleteBrandSystemPropertiesNoContent {
	return &DeleteBrandSystemPropertiesNoContent{}
}

/*DeleteBrandSystemPropertiesNoContent handles this case with default header values.

No Content
*/
type DeleteBrandSystemPropertiesNoContent struct {
}

func (o *DeleteBrandSystemPropertiesNoContent) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/systemproperties][%d] deleteBrandSystemPropertiesNoContent ", 204)
}

func (o *DeleteBrandSystemPropertiesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBrandSystemPropertiesBadRequest creates a DeleteBrandSystemPropertiesBadRequest with default headers values
func NewDeleteBrandSystemPropertiesBadRequest() *DeleteBrandSystemPropertiesBadRequest {
	return &DeleteBrandSystemPropertiesBadRequest{}
}

/*DeleteBrandSystemPropertiesBadRequest handles this case with default header values.

Bad Request
*/
type DeleteBrandSystemPropertiesBadRequest struct {
	Payload *models.Error
}

func (o *DeleteBrandSystemPropertiesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fleet/tenant/{tenantId}/brand/{brandId}/systemproperties][%d] deleteBrandSystemPropertiesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteBrandSystemPropertiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}