// Code generated by go-swagger; DO NOT EDIT.

package fleet_tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetFleetTenantTenantIDSystemPropertiesURLReader is a Reader for the GetFleetTenantTenantIDSystemPropertiesURL structure.
type GetFleetTenantTenantIDSystemPropertiesURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFleetTenantTenantIDSystemPropertiesURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFleetTenantTenantIDSystemPropertiesURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetFleetTenantTenantIDSystemPropertiesURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFleetTenantTenantIDSystemPropertiesURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFleetTenantTenantIDSystemPropertiesURLOK creates a GetFleetTenantTenantIDSystemPropertiesURLOK with default headers values
func NewGetFleetTenantTenantIDSystemPropertiesURLOK() *GetFleetTenantTenantIDSystemPropertiesURLOK {
	return &GetFleetTenantTenantIDSystemPropertiesURLOK{}
}

/*GetFleetTenantTenantIDSystemPropertiesURLOK handles this case with default header values.

OK
*/
type GetFleetTenantTenantIDSystemPropertiesURLOK struct {
	Payload *models.S3PresignedPost
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemPropertiesUrl][%d] getFleetTenantTenantIdSystemPropertiesUrlOK  %+v", 200, o.Payload)
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDSystemPropertiesURLBadRequest creates a GetFleetTenantTenantIDSystemPropertiesURLBadRequest with default headers values
func NewGetFleetTenantTenantIDSystemPropertiesURLBadRequest() *GetFleetTenantTenantIDSystemPropertiesURLBadRequest {
	return &GetFleetTenantTenantIDSystemPropertiesURLBadRequest{}
}

/*GetFleetTenantTenantIDSystemPropertiesURLBadRequest handles this case with default header values.

Bad Request
*/
type GetFleetTenantTenantIDSystemPropertiesURLBadRequest struct {
	Payload *models.Error
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemPropertiesUrl][%d] getFleetTenantTenantIdSystemPropertiesUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFleetTenantTenantIDSystemPropertiesURLNotFound creates a GetFleetTenantTenantIDSystemPropertiesURLNotFound with default headers values
func NewGetFleetTenantTenantIDSystemPropertiesURLNotFound() *GetFleetTenantTenantIDSystemPropertiesURLNotFound {
	return &GetFleetTenantTenantIDSystemPropertiesURLNotFound{}
}

/*GetFleetTenantTenantIDSystemPropertiesURLNotFound handles this case with default header values.

Not Found
*/
type GetFleetTenantTenantIDSystemPropertiesURLNotFound struct {
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLNotFound) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/systemPropertiesUrl][%d] getFleetTenantTenantIdSystemPropertiesUrlNotFound ", 404)
}

func (o *GetFleetTenantTenantIDSystemPropertiesURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
