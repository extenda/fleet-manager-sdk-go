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

// GetDriverPackageReader is a Reader for the GetDriverPackage structure.
type GetDriverPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriverPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDriverPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverPackageOK creates a GetDriverPackageOK with default headers values
func NewGetDriverPackageOK() *GetDriverPackageOK {
	return &GetDriverPackageOK{}
}

/*GetDriverPackageOK handles this case with default header values.

OK
*/
type GetDriverPackageOK struct {
	Payload *models.DriverPackages
}

func (o *GetDriverPackageOK) Error() string {
	return fmt.Sprintf("[GET /driver/package][%d] getDriverPackageOK  %+v", 200, o.Payload)
}

func (o *GetDriverPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriverPackages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverPackageBadRequest creates a GetDriverPackageBadRequest with default headers values
func NewGetDriverPackageBadRequest() *GetDriverPackageBadRequest {
	return &GetDriverPackageBadRequest{}
}

/*GetDriverPackageBadRequest handles this case with default header values.

Bad Request
*/
type GetDriverPackageBadRequest struct {
	Payload *models.Error
}

func (o *GetDriverPackageBadRequest) Error() string {
	return fmt.Sprintf("[GET /driver/package][%d] getDriverPackageBadRequest  %+v", 400, o.Payload)
}

func (o *GetDriverPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
