// Code generated by go-swagger; DO NOT EDIT.

package driver_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetDriverPackagePackageIDReader is a Reader for the GetDriverPackagePackageID structure.
type GetDriverPackagePackageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriverPackagePackageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverPackagePackageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDriverPackagePackageIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDriverPackagePackageIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverPackagePackageIDOK creates a GetDriverPackagePackageIDOK with default headers values
func NewGetDriverPackagePackageIDOK() *GetDriverPackagePackageIDOK {
	return &GetDriverPackagePackageIDOK{}
}

/*GetDriverPackagePackageIDOK handles this case with default header values.

OK
*/
type GetDriverPackagePackageIDOK struct {
	Payload *models.DriverPackage
}

func (o *GetDriverPackagePackageIDOK) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}][%d] getDriverPackagePackageIdOK  %+v", 200, o.Payload)
}

func (o *GetDriverPackagePackageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverPackage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverPackagePackageIDBadRequest creates a GetDriverPackagePackageIDBadRequest with default headers values
func NewGetDriverPackagePackageIDBadRequest() *GetDriverPackagePackageIDBadRequest {
	return &GetDriverPackagePackageIDBadRequest{}
}

/*GetDriverPackagePackageIDBadRequest handles this case with default header values.

Bad Request
*/
type GetDriverPackagePackageIDBadRequest struct {
	Payload *models.Error
}

func (o *GetDriverPackagePackageIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}][%d] getDriverPackagePackageIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDriverPackagePackageIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverPackagePackageIDNotFound creates a GetDriverPackagePackageIDNotFound with default headers values
func NewGetDriverPackagePackageIDNotFound() *GetDriverPackagePackageIDNotFound {
	return &GetDriverPackagePackageIDNotFound{}
}

/*GetDriverPackagePackageIDNotFound handles this case with default header values.

Not Found
*/
type GetDriverPackagePackageIDNotFound struct {
}

func (o *GetDriverPackagePackageIDNotFound) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}][%d] getDriverPackagePackageIdNotFound ", 404)
}

func (o *GetDriverPackagePackageIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}