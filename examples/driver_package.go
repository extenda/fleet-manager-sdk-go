package examples

import (
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/client/driver_package"
	"github.com/extenda/fleet-manager-sdk-go/fleetmanager/models"
)

func fetchDriverPackage() (models.DriverPackage, error) {
	host := "1bcbeyza0g.execute-api.eu-west-1.amazonaws.com"
	c := fleetmanager.NewHTTPClientWithHost(host)

	params := driver_package.NewGetDriverPackagePackageIDParams()
	params.WithPackageID("epson")
	res, err := c.DriverPackage.GetDriverPackagePackageID(params)

	return *res.Payload, err
}
