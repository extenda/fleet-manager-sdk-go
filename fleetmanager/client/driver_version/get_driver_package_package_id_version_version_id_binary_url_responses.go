// Code generated by go-swagger; DO NOT EDIT.

package driver_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

// GetDriverPackagePackageIDVersionVersionIDBinaryURLReader is a Reader for the GetDriverPackagePackageIDVersionVersionIDBinaryURL structure.
type GetDriverPackagePackageIDVersionVersionIDBinaryURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDriverPackagePackageIDVersionVersionIDBinaryURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDriverPackagePackageIDVersionVersionIDBinaryURLOK creates a GetDriverPackagePackageIDVersionVersionIDBinaryURLOK with default headers values
func NewGetDriverPackagePackageIDVersionVersionIDBinaryURLOK() *GetDriverPackagePackageIDVersionVersionIDBinaryURLOK {
	return &GetDriverPackagePackageIDVersionVersionIDBinaryURLOK{}
}

/*GetDriverPackagePackageIDVersionVersionIDBinaryURLOK handles this case with default header values.

OK
*/
type GetDriverPackagePackageIDVersionVersionIDBinaryURLOK struct {
	Payload *models.S3PresignedPost
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLOK) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}/version/{versionId}/binaryUrl][%d] getDriverPackagePackageIdVersionVersionIdBinaryUrlOK  %+v", 200, o.Payload)
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PresignedPost)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest creates a GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest with default headers values
func NewGetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest() *GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest {
	return &GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest{}
}

/*GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest handles this case with default header values.

Bad Request
*/
type GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest struct {
	Payload *models.Error
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}/version/{versionId}/binaryUrl][%d] getDriverPackagePackageIdVersionVersionIdBinaryUrlBadRequest  %+v", 400, o.Payload)
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound creates a GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound with default headers values
func NewGetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound() *GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound {
	return &GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound{}
}

/*GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound handles this case with default header values.

Not Found
*/
type GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound struct {
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound) Error() string {
	return fmt.Sprintf("[GET /driver/package/{packageId}/version/{versionId}/binaryUrl][%d] getDriverPackagePackageIdVersionVersionIdBinaryUrlNotFound ", 404)
}

func (o *GetDriverPackagePackageIDVersionVersionIDBinaryURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
