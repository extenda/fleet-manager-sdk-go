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

// ListBrandsReader is a Reader for the ListBrands structure.
type ListBrandsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBrandsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListBrandsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListBrandsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBrandsOK creates a ListBrandsOK with default headers values
func NewListBrandsOK() *ListBrandsOK {
	return &ListBrandsOK{}
}

/*ListBrandsOK handles this case with default header values.

OK
*/
type ListBrandsOK struct {
	Payload *models.FleetBrands
}

func (o *ListBrandsOK) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand][%d] listBrandsOK  %+v", 200, o.Payload)
}

func (o *ListBrandsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FleetBrands)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBrandsBadRequest creates a ListBrandsBadRequest with default headers values
func NewListBrandsBadRequest() *ListBrandsBadRequest {
	return &ListBrandsBadRequest{}
}

/*ListBrandsBadRequest handles this case with default header values.

Bad Request
*/
type ListBrandsBadRequest struct {
	Payload *models.Error
}

func (o *ListBrandsBadRequest) Error() string {
	return fmt.Sprintf("[GET /fleet/tenant/{tenantId}/brand][%d] listBrandsBadRequest  %+v", 400, o.Payload)
}

func (o *ListBrandsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
