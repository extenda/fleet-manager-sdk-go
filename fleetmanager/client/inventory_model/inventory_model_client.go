// Code generated by go-swagger; DO NOT EDIT.

package inventory_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new inventory model API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for inventory model API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteInventoryManufacturerManufacturerIDModelModelID deletes inventory model
*/
func (a *Client) DeleteInventoryManufacturerManufacturerIDModelModelID(params *DeleteInventoryManufacturerManufacturerIDModelModelIDParams) (*DeleteInventoryManufacturerManufacturerIDModelModelIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInventoryManufacturerManufacturerIDModelModelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteInventoryManufacturerManufacturerIDModelModelID",
		Method:             "DELETE",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInventoryManufacturerManufacturerIDModelModelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteInventoryManufacturerManufacturerIDModelModelIDNoContent), nil

}

/*
DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL deletes inventory model j p o s entries XML
*/
func (a *Client) DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL(params *DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLParams) (*DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL",
		Method:             "DELETE",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}/jposEntriesXmlUrl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLNoContent), nil

}

/*
GetInventoryManufacturerManufacturerIDModel lists inventory models
*/
func (a *Client) GetInventoryManufacturerManufacturerIDModel(params *GetInventoryManufacturerManufacturerIDModelParams) (*GetInventoryManufacturerManufacturerIDModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryManufacturerManufacturerIDModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInventoryManufacturerManufacturerIDModel",
		Method:             "GET",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryManufacturerManufacturerIDModelReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInventoryManufacturerManufacturerIDModelOK), nil

}

/*
GetInventoryManufacturerManufacturerIDModelModelID gets inventory model details
*/
func (a *Client) GetInventoryManufacturerManufacturerIDModelModelID(params *GetInventoryManufacturerManufacturerIDModelModelIDParams) (*GetInventoryManufacturerManufacturerIDModelModelIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryManufacturerManufacturerIDModelModelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInventoryManufacturerManufacturerIDModelModelID",
		Method:             "GET",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryManufacturerManufacturerIDModelModelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInventoryManufacturerManufacturerIDModelModelIDOK), nil

}

/*
GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL uploads inventory model j p o s entries XML
*/
func (a *Client) GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL(params *GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLParams) (*GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURL",
		Method:             "GET",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}/jposEntriesXmlUrl",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInventoryManufacturerManufacturerIDModelModelIDJposEntriesXMLURLOK), nil

}

/*
PostInventoryManufacturerManufacturerIDModel creates inventory model
*/
func (a *Client) PostInventoryManufacturerManufacturerIDModel(params *PostInventoryManufacturerManufacturerIDModelParams) (*PostInventoryManufacturerManufacturerIDModelCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryManufacturerManufacturerIDModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostInventoryManufacturerManufacturerIDModel",
		Method:             "POST",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryManufacturerManufacturerIDModelReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostInventoryManufacturerManufacturerIDModelCreated), nil

}

/*
PostInventoryManufacturerManufacturerIDModelModelID links inventory model with driver version
*/
func (a *Client) PostInventoryManufacturerManufacturerIDModelModelID(params *PostInventoryManufacturerManufacturerIDModelModelIDParams) (*PostInventoryManufacturerManufacturerIDModelModelIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostInventoryManufacturerManufacturerIDModelModelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostInventoryManufacturerManufacturerIDModelModelID",
		Method:             "POST",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostInventoryManufacturerManufacturerIDModelModelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostInventoryManufacturerManufacturerIDModelModelIDNoContent), nil

}

/*
PutInventoryManufacturerManufacturerIDModelModelID updates inventory model
*/
func (a *Client) PutInventoryManufacturerManufacturerIDModelModelID(params *PutInventoryManufacturerManufacturerIDModelModelIDParams) (*PutInventoryManufacturerManufacturerIDModelModelIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutInventoryManufacturerManufacturerIDModelModelIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutInventoryManufacturerManufacturerIDModelModelID",
		Method:             "PUT",
		PathPattern:        "/inventory/manufacturer/{manufacturerId}/model/{modelId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutInventoryManufacturerManufacturerIDModelModelIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutInventoryManufacturerManufacturerIDModelModelIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}