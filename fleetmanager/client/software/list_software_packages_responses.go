// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "fleet-manager-sdk-go/fleetmanager/models"
)

// ListSoftwarePackagesReader is a Reader for the ListSoftwarePackages structure.
type ListSoftwarePackagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSoftwarePackagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSoftwarePackagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListSoftwarePackagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSoftwarePackagesOK creates a ListSoftwarePackagesOK with default headers values
func NewListSoftwarePackagesOK() *ListSoftwarePackagesOK {
	return &ListSoftwarePackagesOK{}
}

/*ListSoftwarePackagesOK handles this case with default header values.

OK
*/
type ListSoftwarePackagesOK struct {
	Payload *models.SoftwarePackages
}

func (o *ListSoftwarePackagesOK) Error() string {
	return fmt.Sprintf("[GET /software/package][%d] listSoftwarePackagesOK  %+v", 200, o.Payload)
}

func (o *ListSoftwarePackagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwarePackages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSoftwarePackagesBadRequest creates a ListSoftwarePackagesBadRequest with default headers values
func NewListSoftwarePackagesBadRequest() *ListSoftwarePackagesBadRequest {
	return &ListSoftwarePackagesBadRequest{}
}

/*ListSoftwarePackagesBadRequest handles this case with default header values.

Bad Request
*/
type ListSoftwarePackagesBadRequest struct {
	Payload *models.Error
}

func (o *ListSoftwarePackagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /software/package][%d] listSoftwarePackagesBadRequest  %+v", 400, o.Payload)
}

func (o *ListSoftwarePackagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
