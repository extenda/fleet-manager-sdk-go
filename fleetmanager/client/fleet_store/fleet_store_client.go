// Code generated by go-swagger; DO NOT EDIT.

package fleet_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fleet store API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fleet store API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID deletes fleet store
*/
func (a *Client) DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID(params *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) (*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID",
		Method:             "DELETE",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent), nil

}

/*
DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties deletes fleet store system properties
*/
func (a *Client) DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties(params *DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesParams) (*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties",
		Method:             "DELETE",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/systemproperties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesNoContent), nil

}

/*
GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStore lists fleet stores
*/
func (a *Client) GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStore(params *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) (*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStore",
		Method:             "GET",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreOK), nil

}

/*
GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID gets fleet store details
*/
func (a *Client) GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID(params *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) (*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID",
		Method:             "GET",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDOK), nil

}

/*
GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties uploads fleet store system properties
*/
func (a *Client) GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties(params *GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesParams) (*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystemproperties",
		Method:             "GET",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}/systemproperties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDSystempropertiesOK), nil

}

/*
PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStore creates fleet store
*/
func (a *Client) PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStore(params *PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams) (*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStore",
		Method:             "POST",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreCreated), nil

}

/*
PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID updates fleet store
*/
func (a *Client) PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID(params *PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams) (*PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreID",
		Method:             "PUT",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/country/{countryId}/store/{storeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutFleetTenantTenantIDBrandBrandIDCountryCountryIDStoreStoreIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
