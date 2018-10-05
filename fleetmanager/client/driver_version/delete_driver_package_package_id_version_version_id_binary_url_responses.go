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

// DeleteDriverPackagePackageIDVersionVersionIDBinaryURLReader is a Reader for the DeleteDriverPackagePackageIDVersionVersionIDBinaryURL structure.
type DeleteDriverPackagePackageIDVersionVersionIDBinaryURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent creates a DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent() *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent {
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent handles this case with default header values.

No Content
*/
type DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent struct {
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}/binaryUrl][%d] deleteDriverPackagePackageIdVersionVersionIdBinaryUrlNoContent ", 204)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest creates a DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest with default headers values
func NewDeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest() *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest {
	return &DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest{}
}

/*DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest handles this case with default header values.

Bad Request
*/
type DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest struct {
	Payload *models.Error
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /driver/package/{packageId}/version/{versionId}/binaryUrl][%d] deleteDriverPackagePackageIdVersionVersionIdBinaryUrlBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDriverPackagePackageIDVersionVersionIDBinaryURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
