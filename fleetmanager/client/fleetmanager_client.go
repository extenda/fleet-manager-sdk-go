// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/driver_package"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/driver_version"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_brand"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_country"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_hardware_profile"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_software_profile"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_store"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_tenant"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/fleet_workstation"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/inventory_manufacturer"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/inventory_model"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/software_package"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/software_version"
)

// Default fleetmanager HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new fleetmanager HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Fleetmanager {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new fleetmanager HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Fleetmanager {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new fleetmanager client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Fleetmanager {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Fleetmanager)
	cli.Transport = transport

	cli.DriverPackage = driver_package.New(transport, formats)

	cli.DriverVersion = driver_version.New(transport, formats)

	cli.FleetBrand = fleet_brand.New(transport, formats)

	cli.FleetCountry = fleet_country.New(transport, formats)

	cli.FleetHardwareProfile = fleet_hardware_profile.New(transport, formats)

	cli.FleetSoftwareProfile = fleet_software_profile.New(transport, formats)

	cli.FleetStore = fleet_store.New(transport, formats)

	cli.FleetTenant = fleet_tenant.New(transport, formats)

	cli.FleetWorkstation = fleet_workstation.New(transport, formats)

	cli.InventoryManufacturer = inventory_manufacturer.New(transport, formats)

	cli.InventoryModel = inventory_model.New(transport, formats)

	cli.SoftwarePackage = software_package.New(transport, formats)

	cli.SoftwareVersion = software_version.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Fleetmanager is a client for fleetmanager
type Fleetmanager struct {
	DriverPackage *driver_package.Client

	DriverVersion *driver_version.Client

	FleetBrand *fleet_brand.Client

	FleetCountry *fleet_country.Client

	FleetHardwareProfile *fleet_hardware_profile.Client

	FleetSoftwareProfile *fleet_software_profile.Client

	FleetStore *fleet_store.Client

	FleetTenant *fleet_tenant.Client

	FleetWorkstation *fleet_workstation.Client

	InventoryManufacturer *inventory_manufacturer.Client

	InventoryModel *inventory_model.Client

	SoftwarePackage *software_package.Client

	SoftwareVersion *software_version.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Fleetmanager) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.DriverPackage.SetTransport(transport)

	c.DriverVersion.SetTransport(transport)

	c.FleetBrand.SetTransport(transport)

	c.FleetCountry.SetTransport(transport)

	c.FleetHardwareProfile.SetTransport(transport)

	c.FleetSoftwareProfile.SetTransport(transport)

	c.FleetStore.SetTransport(transport)

	c.FleetTenant.SetTransport(transport)

	c.FleetWorkstation.SetTransport(transport)

	c.InventoryManufacturer.SetTransport(transport)

	c.InventoryModel.SetTransport(transport)

	c.SoftwarePackage.SetTransport(transport)

	c.SoftwareVersion.SetTransport(transport)

}