// Code generated by go-swagger; DO NOT EDIT.

package driver_package

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new driver package API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for driver package API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteDriverPackagePackageID deletes driver package
*/
func (a *Client) DeleteDriverPackagePackageID(params *DeleteDriverPackagePackageIDParams) (*DeleteDriverPackagePackageIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDriverPackagePackageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDriverPackagePackageID",
		Method:             "DELETE",
		PathPattern:        "/driver/package/{packageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDriverPackagePackageIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDriverPackagePackageIDNoContent), nil

}

/*
GetDriverPackage lists driver packages
*/
func (a *Client) GetDriverPackage(params *GetDriverPackageParams) (*GetDriverPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDriverPackage",
		Method:             "GET",
		PathPattern:        "/driver/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDriverPackageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverPackageOK), nil

}

/*
GetDriverPackagePackageID gets driver package details
*/
func (a *Client) GetDriverPackagePackageID(params *GetDriverPackagePackageIDParams) (*GetDriverPackagePackageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverPackagePackageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDriverPackagePackageID",
		Method:             "GET",
		PathPattern:        "/driver/package/{packageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDriverPackagePackageIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverPackagePackageIDOK), nil

}

/*
PostDriverPackage creates driver package
*/
func (a *Client) PostDriverPackage(params *PostDriverPackageParams) (*PostDriverPackageCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDriverPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostDriverPackage",
		Method:             "POST",
		PathPattern:        "/driver/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDriverPackageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDriverPackageCreated), nil

}

/*
PutDriverPackagePackageID updates driver package
*/
func (a *Client) PutDriverPackagePackageID(params *PutDriverPackagePackageIDParams) (*PutDriverPackagePackageIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDriverPackagePackageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDriverPackagePackageID",
		Method:             "PUT",
		PathPattern:        "/driver/package/{packageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutDriverPackagePackageIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDriverPackagePackageIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}