// Code generated by go-swagger; DO NOT EDIT.

package driver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// ListDriverVersionsReader is a Reader for the ListDriverVersions structure.
type ListDriverVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDriverVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDriverVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListDriverVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDriverVersionsOK creates a ListDriverVersionsOK with default headers values
func NewListDriverVersionsOK() *ListDriverVersionsOK {
	return &ListDriverVersionsOK{}
}

/*ListDriverVersionsOK handles this case with default header values.

OK
*/
type ListDriverVersionsOK struct {
	Payload *models.DriverVersions
}

func (o *ListDriverVersionsOK) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}/version][%d] listDriverVersionsOK  %+v", 200, o.Payload)
}

func (o *ListDriverVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverVersions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDriverVersionsBadRequest creates a ListDriverVersionsBadRequest with default headers values
func NewListDriverVersionsBadRequest() *ListDriverVersionsBadRequest {
	return &ListDriverVersionsBadRequest{}
}

/*ListDriverVersionsBadRequest handles this case with default header values.

Bad Request
*/
type ListDriverVersionsBadRequest struct {
	Payload *models.Error
}

func (o *ListDriverVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}/version][%d] listDriverVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListDriverVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
