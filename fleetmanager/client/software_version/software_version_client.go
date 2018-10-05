// Code generated by go-swagger; DO NOT EDIT.

package software_version

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new software version API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for software version API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteSoftwarePackagePackageIDVersionVersionID deletes software version
*/
func (a *Client) DeleteSoftwarePackagePackageIDVersionVersionID(params *DeleteSoftwarePackagePackageIDVersionVersionIDParams) (*DeleteSoftwarePackagePackageIDVersionVersionIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSoftwarePackagePackageIDVersionVersionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteSoftwarePackagePackageIDVersionVersionID",
		Method:             "DELETE",
		PathPattern:        "/software/package/{packageId}/version/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSoftwarePackagePackageIDVersionVersionIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSoftwarePackagePackageIDVersionVersionIDNoContent), nil

}

/*
GetSoftwarePackagePackageIDVersion lists software versions
*/
func (a *Client) GetSoftwarePackagePackageIDVersion(params *GetSoftwarePackagePackageIDVersionParams) (*GetSoftwarePackagePackageIDVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSoftwarePackagePackageIDVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSoftwarePackagePackageIDVersion",
		Method:             "GET",
		PathPattern:        "/software/package/{packageId}/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSoftwarePackagePackageIDVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSoftwarePackagePackageIDVersionOK), nil

}

/*
GetSoftwarePackagePackageIDVersionVersionID gets software version details
*/
func (a *Client) GetSoftwarePackagePackageIDVersionVersionID(params *GetSoftwarePackagePackageIDVersionVersionIDParams) (*GetSoftwarePackagePackageIDVersionVersionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSoftwarePackagePackageIDVersionVersionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSoftwarePackagePackageIDVersionVersionID",
		Method:             "GET",
		PathPattern:        "/software/package/{packageId}/version/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSoftwarePackagePackageIDVersionVersionIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSoftwarePackagePackageIDVersionVersionIDOK), nil

}

/*
PostSoftwarePackagePackageIDVersion creates software version
*/
func (a *Client) PostSoftwarePackagePackageIDVersion(params *PostSoftwarePackagePackageIDVersionParams) (*PostSoftwarePackagePackageIDVersionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSoftwarePackagePackageIDVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostSoftwarePackagePackageIDVersion",
		Method:             "POST",
		PathPattern:        "/software/package/{packageId}/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSoftwarePackagePackageIDVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSoftwarePackagePackageIDVersionCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}