// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// ListInventoryManufacturersReader is a Reader for the ListInventoryManufacturers structure.
type ListInventoryManufacturersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInventoryManufacturersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListInventoryManufacturersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListInventoryManufacturersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListInventoryManufacturersOK creates a ListInventoryManufacturersOK with default headers values
func NewListInventoryManufacturersOK() *ListInventoryManufacturersOK {
	return &ListInventoryManufacturersOK{}
}

/*ListInventoryManufacturersOK handles this case with default header values.

OK
*/
type ListInventoryManufacturersOK struct {
	Payload *models.InventoryManufacturers
}

func (o *ListInventoryManufacturersOK) Error() string {
	return fmt.Sprintf("[GET /inventory/manufacturer][%d] listInventoryManufacturersOK  %+v", 200, o.Payload)
}

func (o *ListInventoryManufacturersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryManufacturers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInventoryManufacturersBadRequest creates a ListInventoryManufacturersBadRequest with default headers values
func NewListInventoryManufacturersBadRequest() *ListInventoryManufacturersBadRequest {
	return &ListInventoryManufacturersBadRequest{}
}

/*ListInventoryManufacturersBadRequest handles this case with default header values.

Bad Request
*/
type ListInventoryManufacturersBadRequest struct {
	Payload *models.Error
}

func (o *ListInventoryManufacturersBadRequest) Error() string {
	return fmt.Sprintf("[GET /inventory/manufacturer][%d] listInventoryManufacturersBadRequest  %+v", 400, o.Payload)
}

func (o *ListInventoryManufacturersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
