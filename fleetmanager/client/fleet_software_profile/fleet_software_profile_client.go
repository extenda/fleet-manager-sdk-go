// Code generated by go-swagger; DO NOT EDIT.

package fleet_software_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new fleet software profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fleet software profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID deletes fleet software profile
*/
func (a *Client) DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID(params *DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) (*DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID",
		Method:             "DELETE",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent), nil

}

/*
GetFleetTenantTenantIDBrandBrandIDSoftwareprofile lists fleet software profiles
*/
func (a *Client) GetFleetTenantTenantIDBrandBrandIDSoftwareprofile(params *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileParams) (*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFleetTenantTenantIDBrandBrandIDSoftwareprofile",
		Method:             "GET",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileOK), nil

}

/*
GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID gets fleet software profile details
*/
func (a *Client) GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID(params *GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) (*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID",
		Method:             "GET",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDOK), nil

}

/*
PostFleetTenantTenantIDBrandBrandIDSoftwareprofile creates fleet software profile
*/
func (a *Client) PostFleetTenantTenantIDBrandBrandIDSoftwareprofile(params *PostFleetTenantTenantIDBrandBrandIDSoftwareprofileParams) (*PostFleetTenantTenantIDBrandBrandIDSoftwareprofileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetTenantTenantIDBrandBrandIDSoftwareprofileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFleetTenantTenantIDBrandBrandIDSoftwareprofile",
		Method:             "POST",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetTenantTenantIDBrandBrandIDSoftwareprofileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostFleetTenantTenantIDBrandBrandIDSoftwareprofileCreated), nil

}

/*
PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID links fleet software profile with software version
*/
func (a *Client) PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID(params *PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) (*PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID",
		Method:             "POST",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent), nil

}

/*
PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID updates fleet software profile
*/
func (a *Client) PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID(params *PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams) (*PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileID",
		Method:             "PUT",
		PathPattern:        "/fleet/tenant/{tenantId}/brand/{brandId}/softwareprofile/{softwareProfileId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutFleetTenantTenantIDBrandBrandIDSoftwareprofileSoftwareProfileIDNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}