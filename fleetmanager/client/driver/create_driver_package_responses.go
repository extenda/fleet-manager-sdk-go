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

// CreateDriverPackageReader is a Reader for the CreateDriverPackage structure.
type CreateDriverPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDriverPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateDriverPackageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDriverPackageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDriverPackageCreated creates a CreateDriverPackageCreated with default headers values
func NewCreateDriverPackageCreated() *CreateDriverPackageCreated {
	return &CreateDriverPackageCreated{}
}

/*CreateDriverPackageCreated handles this case with default header values.

Created
*/
type CreateDriverPackageCreated struct {
	/*Path to created resource
	 */
	Location string
}

func (o *CreateDriverPackageCreated) Error() string {
	return fmt.Sprintf("[POST /driver/package][%d] createDriverPackageCreated ", 201)
}

func (o *CreateDriverPackageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewCreateDriverPackageBadRequest creates a CreateDriverPackageBadRequest with default headers values
func NewCreateDriverPackageBadRequest() *CreateDriverPackageBadRequest {
	return &CreateDriverPackageBadRequest{}
}

/*CreateDriverPackageBadRequest handles this case with default header values.

Bad Request
*/
type CreateDriverPackageBadRequest struct {
	Payload *models.Error
}

func (o *CreateDriverPackageBadRequest) Error() string {
	return fmt.Sprintf("[POST /driver/package][%d] createDriverPackageBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDriverPackageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}